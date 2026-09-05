package debug

// 服务器侧数据库执行:登录服务器 → 用显式连接(dbconfig.Connection:
// host/port/service|database + 账号清单)连库 —— Oracle 走 EZCONNECT
// (sqlplus 账号/密码@//host:port/service),金仓走 ksql -h host。
// gzou_t/gzzz_t 等 T100 字典查询仍按 TOPENT 语义:企业→账号由 gzou_t 现查,
// 密码查连接的账号清单(含主账号),未收录回退 T100 惯例 账号=密码。
// tnsnames/chenv/ORA 环境读取仅保留给「从服务器获取」辅助(ProbeDBConfig)。

import (
	"fmt"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"tdict/dbconfig"
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
	Connect bool   `json:"connect"` // 用账号+密码(清单/惯例)能否连通
	Error   string `json:"error,omitempty"`
}

// DBReport 完整探查报告
type DBReport struct {
	Zone     string            `json:"zone"`
	TNS      string            `json:"tns"` // 显示用:oracle service / kingbase 库名
	Mappings []EntMapping      `json:"mappings"`
	Probe    *DBProbeResult    `json:"probe,omitempty"`
	Env      map[string]string `json:"env,omitempty"` // 探测到的环境(sqlplus/oracleHome 等)
}

var (
	reGzouMap  = regexp.MustCompile(`^\s*(\d+)\|(\S+)\s*$`)
	reEnvKV    = regexp.MustCompile(`^(ORA|SQLP|TNSADM|TWOTASK)=(\S*)\s*$`)
	reKBEnvKV  = regexp.MustCompile(`^(KSQL|KPORT|KDB)=(.*)\s*$`)
	reTNSField = regexp.MustCompile(`(HOST|PORT|SERVICE_NAME)\s*=\s*([^)\s]+)`)
)

// chenvCmd 拼出"加载 zone 环境"的 bash 前缀(带参数不弹菜单,静默)
func chenvCmd(zone string) string {
	return fmt.Sprintf("source /u3/pub/bin/chenv %s >/dev/null 2>&1", zone)
}

// probeDBEnv 在服务器上探测 ORACLE_HOME / sqlplus / TWO_TASK(chenv zone 后)
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

// probeKsqlPath 在服务器上找 ksql 客户端(command -v,退化 find 常见安装根)
func probeKsqlPath(conn *SSHConn) (string, error) {
	out, err := conn.Output(`
KSQL=$(command -v ksql || true)
[ -z "$KSQL" ] && KSQL=$(find /u2 /home /opt /u1 -maxdepth 8 -name ksql -type f 2>/dev/null | head -1)
echo KSQL=$KSQL
`, 40*time.Second)
	if err != nil && out == "" {
		return "", fmt.Errorf("探测金仓 ksql 失败: %w", err)
	}
	p := ""
	for _, ln := range strings.Split(out, "\n") {
		if m := reKBEnvKV.FindStringSubmatch(strings.TrimSpace(ln)); m != nil && m[1] == "KSQL" {
			p = m[2]
		}
	}
	if p == "" {
		return "", fmt.Errorf("未找到 ksql 客户端(金仓未安装或路径非标准)")
	}
	return p, nil
}

// probeKBEnv 金仓自动探测(「从服务器获取」辅助:实例/库名/端口;不改服务器状态)
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

// kbCmd 拼出在服务器上执行 ksql 的命令(显式 host/port/db;KINGBASE_PASSWORD 传密)
func kbCmd(ksqlPath, host, port, db, connStr, sql string) string {
	q := strings.ReplaceAll(sql, `"`, `\"`)
	user, pass := connStr, connStr
	if i := strings.Index(connStr, "/"); i >= 0 {
		user, pass = connStr[:i], connStr[i+1:]
	}
	return fmt.Sprintf(`KINGBASE_PASSWORD=%s %s -w -h %s -p %s -U %s -d %s -t -A -F '|' -c "%s"`,
		pass, ksqlPath, host, port, user, db, q)
}

// sqlplusRun 拼出在服务器上执行 sqlplus 的命令:chenv zone 提供运行环境,
// 连接串为显式 EZCONNECT "账号/密码@//host:port/service";sqlplusPath 空则用 PATH 内 sqlplus
func sqlplusRun(zone, sqlplusPath, connStr, sql string) string {
	q := strings.ReplaceAll(sql, "'", "'\\''")
	p := sqlplusPath
	if p == "" {
		p = "sqlplus"
	}
	return fmt.Sprintf(`bash -lc '%s; echo "%s" | %s -S %s'`, chenvCmd(zone), q, p, connStr)
}

// ---- 运行时执行上下文(显式连接) ----

// dbRun 服务器侧数据库执行上下文:绑定一条显式连接 + 服务器工具路径。
type dbRun struct {
	conn *dbconfig.Connection // 显式连接(host/port/service|database/accounts/sqlplus…);nil=未挂库
	zone string               // 服务器登录区域(chenv 环境前缀)
	ksql string               // kingbase: ksql 绝对路径(探测缓存)
	oraT string               // oracle: sqlplus 路径(conn.Sqlplus 或探测;空=PATH 找)
}

// connAddrOracle 返回 oracle 显式地址 host:port/service(EZCONNECT 用)
func connAddrOracle(c *dbconfig.Connection) (string, error) {
	if c == nil || c.Host == "" {
		return "", fmt.Errorf("数据库连接缺少主机地址(host)")
	}
	if c.Svc() == "" {
		return "", fmt.Errorf("Oracle 连接缺少服务名(service)")
	}
	port := c.Port
	if port == 0 {
		port = 1521
	}
	return net.JoinHostPort(c.Host, strconv.Itoa(port)) + "/" + c.Svc(), nil
}

// acctConnStr 组装 "账号/密码":密码优先连接账号清单(含主账号),未收录回退 账号=密码 惯例
func acctConnStr(c *dbconfig.Connection, account string) string {
	pass := account
	if p, ok := c.PasswordFor(account); ok {
		pass = p
	}
	return account + "/" + pass
}

// resolveDBRun 解析当前生效连接并探测服务器工具路径;未挂库返回错误
func resolveDBRun(sshConn *SSHConn, cfg *Config) (*dbRun, error) {
	if cfg.DB == nil {
		return nil, fmt.Errorf("当前环境未挂数据库连接(设置-环境-SSH 页选择「数据库连接」)")
	}
	d := &dbRun{conn: cfg.DB, zone: cfg.Zone}
	var err error
	if cfg.DB.Type == "kingbase" {
		d.ksql, err = probeKsqlPath(sshConn)
		if err != nil {
			return nil, err
		}
	} else {
		// oracle:sqlplus 路径一律自动探测(chenv zone 后 command -v;失败留空由 PATH 兜底)
		zone := cfg.Zone
		if zone == "" {
			zone = "36"
		}
		if env, e := probeDBEnv(sshConn, zone); e == nil {
			d.oraT = env["SQLP"]
		}
	}
	return d, nil
}

// dbConnStr 取该库下账号的完整连接串:kingbase=账号/密码;oracle=账号/密码@//host:port/service
func (d *dbRun) dbConnStr(account string) (string, error) {
	creds := acctConnStr(d.conn, account)
	if d.conn.Type != "kingbase" {
		addr, err := connAddrOracle(d.conn)
		if err != nil {
			return "", err
		}
		return creds + "@//" + addr, nil
	}
	return creds, nil
}

// kbTarget 金仓显式目标(host/port/db),host 空回退服务器本机
func (d *dbRun) kbTarget() (host, port, db string, err error) {
	if d.conn == nil || d.conn.Database == "" {
		return "", "", "", fmt.Errorf("金仓连接缺少库名(database)")
	}
	host, port, db = d.conn.Host, strconv.Itoa(d.conn.Port), d.conn.Database
	if host == "" {
		host = "127.0.0.1"
	}
	if d.conn.Port == 0 {
		port = "54321"
	}
	return
}

// exec 执行查询:金仓走 ksql(kbSQL),Oracle 走 sqlplus(oracleSQL)。
// connStr 由 dbConnStr(账号)生成,须与方言匹配;ds 系统账号为 gzou_t/gzzz_t 解析入口。
func (d *dbRun) exec(conn *SSHConn, connStr, oracleSQL, kbSQL string, timeout time.Duration) (string, error) {
	if d.conn.Type == "kingbase" {
		h, p, db, err := d.kbTarget()
		if err != nil {
			return "", err
		}
		return conn.Output(kbCmd(d.ksql, h, p, db, connStr, kbSQL), timeout)
	}
	return conn.Output(sqlplusRun(d.zone, d.oraT, connStr, oracleSQL), timeout)
}

// dbAllMappings 查询 gzou_t 全部企业→账号映射(用 ds 系统账号)
func dbAllMappings(conn *SSHConn, d *dbRun) ([]EntMapping, error) {
	connStr, err := d.dbConnStr("ds")
	if err != nil {
		return nil, err
	}
	kbSQL := `select gzou001,coalesce(gzou003,'-') from gzou_t where gzoustus='Y' order by gzou001`
	oraSQL := "set heading off\nset feedback off\nselect gzou001||'|'||nvl(gzou003,'-') from gzou_t where gzoustus='Y' order by gzou001;"
	out, err := d.exec(conn, connStr, oraSQL, kbSQL, 30*time.Second)
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
func dbResolveJob(conn *SSHConn, d *dbRun, job string) (JobResolve, error) {
	connStr, err := d.dbConnStr("ds")
	if err != nil {
		return JobResolve{}, err
	}
	kbSQL := fmt.Sprintf(`select gzza001,coalesce(gzzz004,' '),coalesce(gzza004,' '),coalesce(gzzz005,'-') from gzzz_t inner join gzza_t on gzza001=gzzz002 where gzzz001='%s'`, job)
	oraSQL := "set heading off\nset feedback off\nselect gzza001||'|'||nvl(gzzz004,' ')||'|'||nvl(gzza004,' ')||'|'||nvl(gzzz005,'-') from gzzz_t inner join gzza_t on gzza001=gzzz002 where gzzz001='" + job + "';"
	out, err := d.exec(conn, connStr, oraSQL, kbSQL, 25*time.Second)
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

// dbConnectTest 用账号连接验证:密码优先连接账号清单,未收录回退 账号=密码
func dbConnectTest(conn *SSHConn, d *dbRun, account string) error {
	connStr, err := d.dbConnStr(account)
	if err != nil {
		return err
	}
	out, err := d.exec(conn, connStr,
		"set heading off\nset feedback off\nselect 'OK' from dual;", `select 'OK'`, 30*time.Second)
	if err != nil {
		msg := firstLines(out, 4)
		return fmt.Errorf("连接失败: %s", msg)
	}
	if strings.Contains(out, "ORA-") || strings.Contains(out, "ERROR") || strings.Contains(out, "error") {
		return fmt.Errorf("连接失败: %s", firstLines(out, 4))
	}
	return nil
}

// ProbeDB 登录服务器按当前环境探查:ent>0 验证该企业账号连通性;ent<=0 仅列映射。
// 连接完全使用显式 dbconfig.Connection(服务器侧可达的 host/port/service|库名)。
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
	if cfg.DB == nil {
		return nil, fmt.Errorf("当前环境未挂数据库连接(设置-环境-SSH 页选择「数据库连接」)")
	}
	d, err := resolveDBRun(conn, cfg)
	if err != nil {
		return nil, err
	}
	report := &DBReport{Zone: zone, TNS: cfg.DB.Svc()}
	report.Env = map[string]string{}
	if cfg.DB.Type == "kingbase" {
		h, p, db, _ := d.kbTarget()
		report.TNS = db
		report.Env = map[string]string{"ksql": d.ksql, "host": h, "port": p, "database": db}
	} else {
		report.Env = map[string]string{
			"host":    cfg.DB.Host,
			"port":    strconv.Itoa(cfg.DB.Port),
			"service": cfg.DB.Svc(),
			"sqlplus": d.oraT,
		}
	}
	maps, err := dbAllMappings(conn, d)
	if err != nil {
		return nil, err
	}
	report.Mappings = maps
	if ent <= 0 {
		return report, nil // 仅列映射
	}
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
	if cfg.DB.Type == "kingbase" {
		h, p, db, _ := d.kbTarget()
		res.Host, res.Port, res.Service = h, p, db
	} else {
		res.Host, res.Port, res.Service = cfg.DB.Host, strconv.Itoa(cfg.DB.Port), cfg.DB.Svc()
	}
	if err := dbConnectTest(conn, d, account); err != nil {
		res.Error = err.Error()
	} else {
		res.Connect = true
	}
	report.Probe = res
	return report, nil
}

// parseTNS 解析服务器 tnsnames.ora 里 TNS 别名的地址/服务(「从服务器获取」辅助用)
func parseTNS(conn *SSHConn, oracleHome, tns string) (host, port, service string, err error) {
	if oracleHome == "" {
		return "", "", "", fmt.Errorf("缺少 ORACLE_HOME,无法定位 tnsnames.ora")
	}
	path := oracleHome + "/network/admin/tnsnames.ora"
	upper := strings.ToUpper(tns)
	out, _ := conn.Output(fmt.Sprintf(
		`awk -v A=%q 'toupper($0) ~ "^" A "[ \t]*=" {f=1; next}
f && ($0 ~ /^[ \t]*$/ || $0 ~ /^[A-Za-z0-9_]+[ \t]*=/) {exit}
f {print}' %s 2>/dev/null | grep -E 'HOST|PORT|SERVICE_NAME' | head -6`, upper, path), 15*time.Second)
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

// DBProbeReq 设置页「从服务器获取」请求:SSH 连接 + 数据库类型
type DBProbeReq struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Zone     string `json:"zone"`
	Type     string `json:"type"` // oracle | kingbase
}

// DBProbeOut 探测结果:拿到的字段回填 DB 表单,拿不到的留空由用户手填
type DBProbeOut struct {
	Type       string `json:"type"`
	TNS        string `json:"tns,omitempty"`      // oracle: TNS 别名(按区域推导/服务器 TWO_TASK)
	Port       int    `json:"port,omitempty"`     // oracle: tnsnames 端口 / kingbase: 实例端口
	Database   string `json:"database,omitempty"` // kingbase: 库名
	Sqlplus    string `json:"sqlplus,omitempty"`  // oracle: 服务器 sqlplus 路径
	OracleHome string `json:"oracleHome,omitempty"`
	TwoTask    string `json:"twoTask,omitempty"`
	Host       string `json:"host,omitempty"` // oracle tnsnames 解析出的实际库地址
	Service    string `json:"service,omitempty"`
	Note       string `json:"note,omitempty"` // 未获取到时的说明
}

// ProbeDBConfig 登录服务器自动获取数据库连接要素(「从服务器获取」辅助,只读)。
// oracle → chenv 环境 + tnsnames 解析;kingbase → 实例发现(ksql/端口/库名)。
// 解析结果仅作回填参考;运行时连接一律使用显式 host/port/service。
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
	tns := zoneTNSName(zone)
	env, err := probeDBEnv(conn, zone)
	if err != nil {
		out.TNS = tns // chenv 拿不到也给出按区域推导的 TNS
		out.Note = err.Error()
		return out, nil
	}
	out.TNS = env["TWOTASK"]
	if out.TNS == "" {
		out.TNS = tns
	}
	out.Sqlplus = env["SQLP"]
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

// DBAccVerifyReq 账号验证请求(设置页账号清单「验证」;只读)。
// 内嵌 DBProbeReq 提供 SSH 登录要素(password=SSH 密码);目标库要素显式给出:
// oracle 用 dbHost/dbPort/dbSvc;kingbase 用 dbHost/dbPort/dbDatabase。
type DBAccVerifyReq struct {
	DBProbeReq
	Account    string `json:"account"`            // 待验证账号(schema 名)
	AcctPass   string `json:"acctPassword"`       // 待验证账号密码
	DBHost     string `json:"dbHost,omitempty"`   // 库主机(服务器侧可达;空=127.0.0.1)
	DBPort     int    `json:"dbPort,omitempty"`   // 库端口(0=默认 1521/54321)
	DBSvc      string `json:"dbSvc,omitempty"`    // oracle: 服务名
	DBDatabase string `json:"dbDatabase,omitempty"` // kingbase: 库名
}

// VerifyDBAcct 服务器侧以显式 账号/密码 连库执行 select 1(账号清单行内验证用)。
// oracle: chenv zone + sqlplus 账号/密码@//host:port/service;kingbase: ksql -h host。
func VerifyDBAcct(req DBAccVerifyReq) error {
	ssh := SSHConfig{Host: req.Host, User: req.User, Password: req.Password}
	if req.Port > 0 {
		ssh.Port = req.Port
	} else {
		ssh.Port = 22
	}
	conn, err := Dial(ssh)
	if err != nil {
		return fmt.Errorf("SSH 连接失败: %w", err)
	}
	defer conn.Close()

	creds := req.Account + "/" + req.AcctPass
	var out string
	if req.Type == "kingbase" {
		ksql, err := probeKsqlPath(conn)
		if err != nil {
			return err
		}
		host, port, db := req.DBHost, strconv.Itoa(req.DBPort), req.DBDatabase
		if host == "" {
			host = "127.0.0.1"
		}
		if req.DBPort == 0 {
			port = "54321"
		}
		if db == "" {
			return fmt.Errorf("缺少库名(dbDatabase)")
		}
		out, err = conn.Output(kbCmd(ksql, host, port, db, creds, `select 'OK'`), 30*time.Second)
	} else {
		zone := req.Zone
		if zone == "" {
			zone = "36"
		}
		if req.DBSvc == "" {
			return fmt.Errorf("缺少服务名(dbSvc)")
		}
		host := req.DBHost
		if host == "" {
			host = "127.0.0.1"
		}
		port := req.DBPort
		if port == 0 {
			port = 1521
		}
		sql := "set heading off\nset feedback off\nselect 'OK' from dual;"
		addr := net.JoinHostPort(host, strconv.Itoa(port)) + "/" + req.DBSvc
		out, err = conn.Output(sqlplusRun(zone, "", creds+"@//"+addr, sql), 30*time.Second)
	}
	if err != nil {
		return fmt.Errorf("连接失败: %s", firstLines(out, 4))
	}
	if strings.Contains(out, "ORA-") || strings.Contains(out, "ERROR") || strings.Contains(out, "error") {
		return fmt.Errorf("连接失败: %s", firstLines(out, 4))
	}
	return nil
}
