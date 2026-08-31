package debug

// 数据库连接探查:登录服务器 → 查 gzou_t 拿到企业对应账号(gzou003)
// → 用该账号尝试连接数据库,输出连接信息。
// 机制来自对 T100 的调研:cl_connect_to_ds 按 TOPENT 查 gzou_t.gzou003 决定 schema/账号,
// 密码规则为 fglprofile 明文(账号=密码,如 dsdemo/dsdemo)。

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// EntMapping 企业 → 账号映射
type EntMapping struct {
	Ent     int    `json:"ent"`
	Account string `json:"account"` // gzou003:账号/schema 名
}

// DBProbeResult 单账号连接探查结果
type DBProbeResult struct {
	Ent     int    `json:"ent"`
	Account string `json:"account"`
	Host    string `json:"host"`
	Port    string `json:"port"`
	Service string `json:"service"`
	Connect bool   `json:"connect"` // 用 账号/账号@tns 能否连通(密码规则:账号=密码)
	Error   string `json:"error,omitempty"`
}

// DBReport 完整探查报告
type DBReport struct {
	Zone     string            `json:"zone"`
	TNS      string            `json:"tns"`
	Mappings []EntMapping      `json:"mappings"`        // 全部企业映射
	Probe    *DBProbeResult    `json:"probe,omitempty"` // 指定企业的连接验证
	Env      map[string]string `json:"env,omitempty"`   // 探测到的环境(sqlplus/oracleHome)
}

var (
	reGzouMap  = regexp.MustCompile(`^\s*(\d+)\|(\S+)\s*$`)
	reEnvKV    = regexp.MustCompile(`^(ORA|SQLP|TNSADM|TWOTASK)=(\S*)\s*$`)
	reKBEnvKV  = regexp.MustCompile(`^(KSQL|KPORT|KDB)=(.*)\s*$`)
	reTNSField = regexp.MustCompile(`(HOST|PORT|SERVICE_NAME)\s*=\s*([^)\s]+)`)
	reGzzzRow  = regexp.MustCompile(`^\s*(\S+)\|(\S+)\s*$`)
	reKBData   = regexp.MustCompile(`kingbase\s+-D\s+(\S+)`)
)

// DBType 判断数据库类型("oracle" 默认 | "kingbase")
func (c *Config) DBType() string {
	if c.DB != nil && c.DB.Type == "kingbase" {
		return "kingbase"
	}
	return "oracle"
}

// chenvCmd 拼出"加载 zone 环境"的 bash 前缀(带参数不弹菜单,静默)
func chenvCmd(zone string) string {
	return fmt.Sprintf("source /u3/pub/bin/chenv %s >/dev/null 2>&1", zone)
}

// probeDBEnv 在服务器上探测 ORACLE_HOME / sqlplus / TWO_TASK
func probeDBEnv(conn *SSHConn, zone string) (map[string]string, error) {
	cmd := fmt.Sprintf(`bash -lc '%s; echo ORA=$ORACLE_HOME; echo SQLP=$(command -v sqlplus); echo TNSADM=$TNS_ADMIN; echo TWOTASK=$TWO_TASK'`, chenvCmd(zone))
	out, err := conn.Output(cmd, 25*time.Second)
	if err != nil && out == "" {
		return nil, fmt.Errorf("探测数据库环境失败: %w", err)
	}
	env := map[string]string{}
	for _, ln := range strings.Split(out, "\n") {
		if m := reEnvKV.FindStringSubmatch(ln); m != nil {
			env[m[1]] = m[2]
		}
	}
	if env["ORA"] == "" {
		return nil, fmt.Errorf("未探测到 ORACLE_HOME(zone %s 环境未加载?)", zone)
	}
	return env, nil
}

// probeKBEnv 金仓自动探测(不依赖环境脚本,走实例发现):
//   - ksql 客户端:command -v,退化到 find 常见安装根(/u2 /home /opt)
//   - 实例与库名:ps 里 kingbase -D <数据目录> → 库名取目录名(如 /u2/kingbase/topprd/data → topprd)
//   - 端口:实例 kingbase.conf 的 port=,读不到用 54321(金仓默认)
//
// 返回 env: KSQL=ksql 绝对路径 KPORT=端口 KDB=库名
func probeKBEnv(conn *SSHConn) (map[string]string, error) {
	out, err := conn.Output(`
KSQL=$(command -v ksql || true)
[ -z "$KSQL" ] && KSQL=$(find /u2 /home /opt /u1 -maxdepth 8 -name ksql -type f 2>/dev/null | head -1)
echo KSQL=$KSQL
DATA=$(ps -ef | grep 'kingbase' | grep -v grep | grep -o 'kingbase -D [^ ]*' | head -1 | awk '{print $3}')
echo KDB=$(basename "$(dirname "$DATA")")
PORT=$(grep -E '^[[:space:]]*port[[:space:]]*=' "$DATA/kingbase.conf" 2>/dev/null | head -1 | grep -o '[0-9]\+')
echo KPORT=${PORT:-54321}
`, 40*time.Second)
	if err != nil && out == "" {
		return nil, fmt.Errorf("探测金仓环境失败: %w", err)
	}
	env := map[string]string{}
	for _, ln := range strings.Split(out, "\n") {
		if m := reKBEnvKV.FindStringSubmatch(strings.TrimSpace(ln)); m != nil {
			env[m[1]] = m[2]
		}
	}
	if env["KSQL"] == "" {
		return nil, fmt.Errorf("未找到 ksql 客户端(金仓未安装或路径非标准)")
	}
	if env["KDB"] == "" || env["KDB"] == "data" {
		return nil, fmt.Errorf("未发现运行中的金仓实例(kingbase -D 数据目录)")
	}
	return env, nil
}

// kbCmd 拼出在服务器上执行 ksql 的命令(金仓:KINGBASE_PASSWORD 传密,-w 禁交互,-t -A 裸输出)
func kbCmd(ksqlPath, port, db, connStr, sql string) string {
	q := strings.ReplaceAll(sql, `"`, `\"`)
	user, pass := connStr, connStr // 连接串复用为 账号/密码(账号=密码规则)
	if i := strings.Index(connStr, "/"); i >= 0 {
		user, pass = connStr[:i], connStr[i+1:]
	}
	return fmt.Sprintf(`KINGBASE_PASSWORD=%s %s -w -h 127.0.0.1 -p %s -U %s -d %s -t -A -F '|' -c "%s"`,
		pass, ksqlPath, port, user, db, q)
}

// sqlplusCmd 拼出在服务器上执行 sqlplus 的命令(T100 环境 + 静默)
func sqlplusCmd(zone, connStr, sql string) string {
	q := strings.ReplaceAll(sql, "'", "'\\''")
	return fmt.Sprintf(`bash -lc '%s; echo "%s" | sqlplus -S %s'`, chenvCmd(zone), q, connStr)
}

// kbCtx 金仓连接上下文(探测结果;nil 表示走 Oracle)
type kbCtx struct{ ksql, port, db string }

// dbRun 数据库执行上下文:金仓(kb 非 nil,ksql 直连)或 Oracle(tns + sqlplus)
type dbRun struct {
	kb   *kbCtx
	zone string
	tns  string
}

// resolveDBRun 按配置探测并构造执行上下文(金仓实例发现失败返回 nil 让调用方报错)
func resolveDBRun(conn *SSHConn, cfg *Config) (*dbRun, error) {
	if cfg.DBType() != "kingbase" {
		return &dbRun{zone: cfg.Zone, tns: cfg.TNSName()}, nil
	}
	env, err := probeKBEnv(conn)
	if err != nil {
		return nil, err
	}
	kb := &kbCtx{ksql: env["KSQL"], port: env["KPORT"], db: env["KDB"]}
	if cfg.DB != nil && cfg.DB.TNS != "" {
		kb.db = cfg.DB.TNS
	}
	if cfg.DB != nil && cfg.DB.Port > 0 {
		kb.port = strconv.Itoa(cfg.DB.Port)
	}
	return &dbRun{kb: kb, zone: cfg.Zone, tns: kb.db}, nil
}

// exec 执行查询:金仓走 ksql(kbSQL),Oracle 走 sqlplus(oracleSQL)
func (d *dbRun) exec(conn *SSHConn, connStr, oracleSQL, kbSQL string, timeout time.Duration) (string, error) {
	if d.kb != nil {
		return conn.Output(kbCmd(d.kb.ksql, d.kb.port, d.kb.db, connStr, kbSQL), timeout)
	}
	return conn.Output(sqlplusCmd(d.zone, connStr, oracleSQL), timeout)
}

// dbAllMappings 查询 gzou_t 全部企业→账号映射(用 ds 系统账号)
func dbAllMappings(conn *SSHConn, zone, tns string, kb *kbCtx) ([]EntMapping, error) {
	var out string
	var err error
	if kb != nil {
		out, err = conn.Output(kbCmd(kb.ksql, kb.port, kb.db,
			"ds/ds", `select gzou001,coalesce(gzou003,'-') from gzou_t where gzoustus='Y' order by gzou001`), 30*time.Second)
	} else {
		sql := "set heading off\nset feedback off\nselect gzou001||'|'||nvl(gzou003,'-') from gzou_t where gzoustus='Y' order by gzou001;"
		out, err = conn.Output(sqlplusCmd(zone, fmt.Sprintf("ds/ds@%s", tns), sql), 30*time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("查询 gzou_t 失败: %w (%s)", err, firstLines(out, 3))
	}
	var maps []EntMapping
	for _, ln := range strings.Split(out, "\n") {
		if m := reGzouMap.FindStringSubmatch(ln); m != nil {
			ent, _ := strconv.Atoi(m[1])
			maps = append(maps, EntMapping{Ent: ent, Account: m[2]})
		}
	}
	if len(maps) == 0 {
		return nil, fmt.Errorf("gzou_t 无有效企业记录(输出: %s)", firstLines(out, 5))
	}
	sort.Slice(maps, func(i, j int) bool { return maps[i].Ent < maps[j].Ent })
	return maps, nil
}

// JobResolve 作业解析结果(gendbg 原版语义)
type JobResolve struct {
	Prog      string // 实体程序编号(gzza001/gzzz002)
	Module    string // 模块代码(gzzz005)
	LaunchRef string // gzza004 去 "$FGLRUN" 前缀的启动引用(如 $CINi/ainq120_wf),由选区 shell 展开权威路径
	Extra     string // gzzz004 额外参数(gendbg 拼在程序名后)
}

// dbResolveJob 按作业编号查实体程序与启动引用(对齐 gendbg.4gl:
// SELECT gzza001,gzzz004,gzza004,gzzz005 FROM gzzz_t INNER JOIN gzza_t ON gzza001=gzzz002)。
// 一个程序(gzzz002)可被多个作业编号共用。无记录返回零值。
func dbResolveJob(conn *SSHConn, zone, tns, job string, kb *kbCtx) (JobResolve, error) {
	var out string
	var err error
	if kb != nil {
		out, err = conn.Output(kbCmd(kb.ksql, kb.port, kb.db, "ds/ds",
			fmt.Sprintf(`select gzza001,coalesce(gzzz004,' '),coalesce(gzza004,' '),coalesce(gzzz005,'-') from gzzz_t inner join gzza_t on gzza001=gzzz002 where gzzz001='%s'`, job)), 25*time.Second)
	} else {
		sql := "set heading off\nset feedback off\nselect gzza001||'|'||nvl(gzzz004,' ')||'|'||nvl(gzza004,' ')||'|'||nvl(gzzz005,'-') from gzzz_t inner join gzza_t on gzza001=gzzz002 where gzzz001='" + job + "';"
		out, err = conn.Output(sqlplusCmd(zone, fmt.Sprintf("ds/ds@%s", tns), sql), 25*time.Second)
	}
	if err != nil {
		return JobResolve{}, fmt.Errorf("查询 gzzz_t 失败: %w (%s)", err, firstLines(out, 3))
	}
	var jr JobResolve
	for _, ln := range strings.Split(out, "\n") {
		if strings.Contains(ln, "ORA-") || strings.Contains(ln, "ERROR") {
			return JobResolve{}, fmt.Errorf("gzzz_t 查询出错: %s", firstLines(out, 3))
		}
		ln = strings.TrimSpace(ln)
		parts := strings.Split(ln, "|")
		if len(parts) < 4 || parts[0] == "" {
			continue
		}
		// gendbg 拼装时 gzza004 = "$FGLRUN $<模块变量>/<程序>";去掉前缀留启动引用($变量 交 shell 展开)
		jr = JobResolve{Prog: parts[0], Extra: strings.TrimSpace(parts[1]), Module: parts[3]}
		jr.LaunchRef = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(parts[2]), "$FGLRUN"))
		break
	}
	return jr, nil
}

// dbConnectTest 用 账号/账号 尝试连接(密码规则:账号=密码,fglprofile 明文)
func dbConnectTest(conn *SSHConn, zone, tns, account string, kb *kbCtx) error {
	var out string
	var err error
	if kb != nil {
		out, err = conn.Output(kbCmd(kb.ksql, kb.port, kb.db, account+"/", `select 'OK'`), 30*time.Second)
	} else {
		sql := "set heading off\nset feedback off\nselect 'OK' from dual;"
		out, err = conn.Output(sqlplusCmd(zone, fmt.Sprintf("%s/%s@%s", account, account, tns), sql), 30*time.Second)
	}
	if err != nil {
		msg := firstLines(out, 4)
		return fmt.Errorf("连接失败: %s", msg)
	}
	if strings.Contains(out, "ORA-") || strings.Contains(out, "ERROR") || strings.Contains(out, "error") {
		return fmt.Errorf("连接失败: %s", firstLines(out, 4))
	}
	return nil
}

// ProbeDB 登录服务器并探查数据库连接:
// ent>0 时验证该企业账号连通性;ent<=0 时仅列出全部企业映射。
// 按配置的数据库类型(oracle/kingbase)自动探测服务器上的连接要素。
func ProbeDB(cfg *Config, ent int) (*DBReport, error) {
	conn, err := Dial(cfg.SSH)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %w", err)
	}
	defer conn.Close()

	zone := cfg.Zone
	if zone == "" {
		zone = "36"
	}
	report := &DBReport{Zone: zone}
	var kb *kbCtx
	tns := ""
	if cfg.DBType() == "kingbase" {
		env, err := probeKBEnv(conn)
		if err != nil {
			return nil, err
		}
		kb = &kbCtx{ksql: env["KSQL"], port: env["KPORT"], db: env["KDB"]}
		if cfg.DB != nil && cfg.DB.TNS != "" {
			kb.db = cfg.DB.TNS
		}
		if cfg.DB != nil && cfg.DB.Port > 0 {
			kb.port = strconv.Itoa(cfg.DB.Port) // 配置显式指定端口
		}
		tns = kb.db
		report.TNS = kb.db
		report.Env = map[string]string{"ksql": kb.ksql, "port": kb.port, "database": kb.db}
	} else {
		tns = cfg.TNSName()
		env, err := probeDBEnv(conn, zone)
		if err != nil {
			return nil, err
		}
		report.TNS = tns
		report.Env = map[string]string{
			"oracleHome": env["ORA"],
			"sqlplus":    env["SQLP"],
			"twoTask":    env["TWOTASK"],
		}
	}
	maps, err := dbAllMappings(conn, zone, tns, kb)
	if err != nil {
		return nil, err
	}
	report.Mappings = maps
	if ent <= 0 {
		return report, nil // 仅列映射
	}
	// 找到该企业的账号
	account := ""
	for _, m := range maps {
		if m.Ent == ent {
			account = m.Account
			break
		}
	}
	if account == "" {
		return nil, fmt.Errorf("企业 %d 不存在于 gzou_t(可用 tdict debug db 查看全部)", ent)
	}
	res := &DBProbeResult{Ent: ent, Account: account}
	if kb != nil {
		res.Host, res.Port, res.Service = "127.0.0.1", kb.port, kb.db
	} else if h, p, s, e := parseTNS(conn, report.Env["oracleHome"], tns); e == nil {
		res.Host, res.Port, res.Service = h, p, s
	}
	if err := dbConnectTest(conn, zone, tns, account, kb); err != nil {
		res.Error = err.Error()
	} else {
		res.Connect = true
	}
	report.Probe = res
	return report, nil
}

// parseTNS 解析 tnsnames.ora 里的 TNS 别名,拿主机/端口/服务名
func parseTNS(conn *SSHConn, oracleHome, tns string) (host, port, service string, err error) {
	path := oracleHome + "/network/admin/tnsnames.ora"
	// 别名在 tnsnames.ora 里通常大写,grep -i 忽略大小写;-A12 覆盖一个连接块
	out, _ := conn.Output(fmt.Sprintf(
		"grep -A12 -i '^%s[ \\t]*=' %s 2>/dev/null | grep -E 'HOST|PORT|SERVICE_NAME' | head -6",
		strings.ToUpper(tns), path), 15*time.Second)
	for _, ln := range strings.Split(out, "\n") {
		for _, m := range reTNSField.FindAllStringSubmatch(ln, -1) {
			switch m[1] {
			case "HOST":
				host = m[2]
			case "PORT":
				port = m[2]
			case "SERVICE_NAME":
				service = m[2]
			}
		}
	}
	if host == "" {
		err = fmt.Errorf("tnsnames.ora 未解析到 %s", tns)
	}
	return
}

// DBProbeReq 设置页「自动获取数据库配置」请求:SSH 连接 + 数据库类型
type DBProbeReq struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Zone     string `json:"zone"`
	Type     string `json:"type"` // oracle | kingbase
}

// DBProbeOut 探测结果:拿到的字段回填表单,拿不到的留空由用户手填
type DBProbeOut struct {
	Type       string `json:"type"`
	TNS        string `json:"tns,omitempty"`      // oracle: TNS 别名(按区域推导)
	Port       int    `json:"port,omitempty"`     // kingbase: 实例端口
	Database   string `json:"database,omitempty"` // kingbase: 库名
	OracleHome string `json:"oracleHome,omitempty"`
	TwoTask    string `json:"twoTask,omitempty"`
	Host       string `json:"host,omitempty"` // oracle tnsnames 解析出的实际库地址
	Service    string `json:"service,omitempty"`
	Note       string `json:"note,omitempty"` // 未获取到时的说明
}

// ProbeDBConfig 登录服务器自动获取数据库连接要素(只读命令,不改服务器状态)。
// oracle → chenv 环境 + tnsnames 解析;kingbase → 实例发现(ksql/端口/库名)。
func ProbeDBConfig(req DBProbeReq) (*DBProbeOut, error) {
	ssh := SSHConfig{Host: req.Host, User: req.User, Password: req.Password}
	if req.Port > 0 {
		ssh.Port = req.Port
	} else {
		ssh.Port = 22
	}
	conn, err := Dial(ssh)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %w", err)
	}
	defer conn.Close()

	out := &DBProbeOut{Type: req.Type}
	if req.Type == "kingbase" {
		env, err := probeKBEnv(conn)
		if err != nil {
			out.Note = err.Error()
			return out, nil // 拿不到也让用户手填,不报错
		}
		out.Port, _ = strconv.Atoi(env["KPORT"])
		out.Database = env["KDB"]
		return out, nil
	}
	// oracle
	zone := req.Zone
	if zone == "" {
		zone = "36"
	}
	cfg := &Config{Zone: zone}
	env, err := probeDBEnv(conn, zone)
	if err != nil {
		out.TNS = cfg.TNSName() // chenv 拿不到也给出按区域推导的 TNS
		out.Note = err.Error()
		return out, nil
	}
	out.TNS = cfg.TNSName()
	out.OracleHome = env["ORA"]
	out.TwoTask = env["TWOTASK"]
	if h, p, s, e := parseTNS(conn, env["ORA"], out.TNS); e == nil {
		out.Host, out.Service = h, s
		out.Port, _ = strconv.Atoi(p)
	}
	return out, nil
}

// firstLines 取输出前 n 行(清理空行)
func firstLines(s string, n int) string {
	var out []string
	for _, ln := range strings.Split(s, "\n") {
		ln = strings.TrimSpace(ln)
		if ln != "" {
			out = append(out, ln)
		}
		if len(out) >= n {
			break
		}
	}
	return strings.Join(out, "; ")
}
