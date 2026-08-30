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
	Ent      int    `json:"ent"`
	Account  string `json:"account"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Service  string `json:"service"`
	Connect  bool   `json:"connect"` // 用 账号/账号@tns 能否连通(密码规则:账号=密码)
	Error    string `json:"error,omitempty"`
}

// DBReport 完整探查报告
type DBReport struct {
	Zone     string          `json:"zone"`
	TNS      string          `json:"tns"`
	Mappings []EntMapping    `json:"mappings"`           // 全部企业映射
	Probe    *DBProbeResult  `json:"probe,omitempty"`    // 指定企业的连接验证
	Env      map[string]string `json:"env,omitempty"`    // 探测到的环境(sqlplus/oracleHome)
}

var (
	reGzouMap  = regexp.MustCompile(`^\s*(\d+)\|(\S+)\s*$`)
	reEnvKV    = regexp.MustCompile(`^(ORA|SQLP|TNSADM|TWOTASK)=(\S*)\s*$`)
	reTNSField = regexp.MustCompile(`(HOST|PORT|SERVICE_NAME)\s*=\s*([^)\s]+)`)
	reGzzzRow  = regexp.MustCompile(`^\s*(\S+)\|(\S+)\s*$`)
)

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

// sqlplusCmd 拼出在服务器上执行 sqlplus 的命令(T100 环境 + 静默)
func sqlplusCmd(zone, connStr, sql string) string {
	q := strings.ReplaceAll(sql, "'", "'\\''")
	return fmt.Sprintf(`bash -lc '%s; echo "%s" | sqlplus -S %s'`, chenvCmd(zone), q, connStr)
}

// dbAllMappings 查询 gzou_t 全部企业→账号映射(用 ds 系统账号)
func dbAllMappings(conn *SSHConn, zone, tns string) ([]EntMapping, error) {
	sql := "set heading off\nset feedback off\nselect gzou001||'|'||nvl(gzou003,'-') from gzou_t where gzoustus='Y' order by gzou001;"
	out, err := conn.Output(sqlplusCmd(zone, fmt.Sprintf("ds/ds@%s", tns), sql), 30*time.Second)
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

// dbResolveJob 按作业编号查 gzzz_t,返回实体程序编号(gzzz002)与模块码(gzzz005)。
// 对齐 gendbg.4gl 的解析:SELECT ... FROM gzzz_t INNER JOIN gzza_t ON gzza001=gzzz002
// WHERE gzzz001=作业编号;一个程序(gzzz002)可被多个作业编号共用。
// 无记录返回 ("","",nil)——调用方回退到 42r 文件搜索。
func dbResolveJob(conn *SSHConn, zone, tns, job string) (prog, module string, err error) {
	sql := "set heading off\nset feedback off\nselect gzzz002||'|'||nvl(gzzz005,'-') from gzzz_t where gzzz001='" + job + "';"
	out, err := conn.Output(sqlplusCmd(zone, fmt.Sprintf("ds/ds@%s", tns), sql), 25*time.Second)
	if err != nil {
		return "", "", fmt.Errorf("查询 gzzz_t 失败: %w (%s)", err, firstLines(out, 3))
	}
	for _, ln := range strings.Split(out, "\n") {
		if strings.Contains(ln, "ORA-") {
			return "", "", fmt.Errorf("gzzz_t 查询出错: %s", firstLines(out, 3))
		}
		if m := reGzzzRow.FindStringSubmatch(ln); m != nil {
			return m[1], m[2], nil
		}
	}
	return "", "", nil
}

// dbConnectTest 用 账号/账号@tns 尝试连接(密码规则:账号=密码,fglprofile 明文)
func dbConnectTest(conn *SSHConn, zone, tns, account string) error {
	sql := "set heading off\nset feedback off\nselect 'OK' from dual;"
	out, err := conn.Output(sqlplusCmd(zone, fmt.Sprintf("%s/%s@%s", account, account, tns), sql), 30*time.Second)
	if err != nil {
		// sqlplus 连接失败也返回退出码;取输出中的 ORA- 错误
		msg := firstLines(out, 4)
		return fmt.Errorf("连接失败: %s", msg)
	}
	if strings.Contains(out, "ORA-") || strings.Contains(out, "ERROR") {
		return fmt.Errorf("连接失败: %s", firstLines(out, 4))
	}
	return nil
}

// ProbeDB 登录服务器并探查数据库连接:
// ent>0 时验证该企业账号连通性;ent<=0 时仅列出全部企业映射。
// password 可覆盖默认"账号=密码"规则。
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
	tns := cfg.TNSName()
	env, err := probeDBEnv(conn, zone)
	if err != nil {
		return nil, err
	}
	maps, err := dbAllMappings(conn, zone, tns)
	if err != nil {
		return nil, err
	}
	report := &DBReport{
		Zone:     zone,
		TNS:      tns,
		Mappings: maps,
		Env: map[string]string{
			"oracleHome": env["ORA"],
			"sqlplus":    env["SQLP"],
			"twoTask":    env["TWOTASK"],
		},
	}
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
	if h, p, s, e := parseTNS(conn, env["ORA"], tns); e == nil {
		res.Host, res.Port, res.Service = h, p, s
	}
	if err := dbConnectTest(conn, zone, tns, account); err != nil {
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
