package cli

// --online 在线直查:root 打开远程 ERP 连接(live)而非本地 SQLite。
// 当前支持 rt(表字典);rv/scc/desc/rq/win 的在线化属后续阶段。
// 连接解析优先级:--conn <名称> > debug activeEnv 对应环境的 dbConn > 默认(isDefault)连接。
// 使用方式: tdict rt dzea_t --online [--conn 恒烁dsdata]

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"tdict/db"
	"tdict/dbconfig"
	"tdict/erpdb"
	"tdict/live"
)

var (
	useOnline  bool   // root --online
	onlineConn string // root --conn(在线连接名;空=按优先级解析)
	liveDB     *live.Live
)

// onlineSupported 记录当前命令是否已支持在线查询(防止普通命令在 online 下误用 nil 数据库)。
var onlineSupported = map[string]bool{
	"rt": true, "table": true,
}

// openOnline 在 root PersistentPreRunE(useOnline=true)时打开远程连接。
func openOnline(cmdName string) error {
	if cmdName != "" && !onlineSupported[cmdName] {
		return fmt.Errorf("命令 %s 暂不支持 --online(当前支持: rt);离线模式请去掉 --online", cmdName)
	}
	conn, err := resolveOnlineConnection(onlineConn)
	if err != nil {
		return err
	}
	l, err := live.Open(context.Background(), *conn)
	if err != nil {
		return err
	}
	liveDB = l
	return nil
}

func closeOnline() {
	if liveDB != nil {
		liveDB.Close()
		liveDB = nil
	}
}

// resolveOnlineConnection 解析在线查询连接。
func resolveOnlineConnection(name string) (*dbconfig.Connection, error) {
	cfgPath, err := resolveConfigPath(configPath)
	if err != nil {
		return nil, err
	}
	cfg, err := dbconfig.Load(cfgPath)
	if err != nil {
		return nil, err
	}
	if name != "" {
		return cfg.FindByName(name)
	}
	// 未指定:取 debug activeEnv 对应环境的 dbConn(若有)
	if connName := dbConnOfActiveEnv(cfgPath); connName != "" {
		if c, e := cfg.FindByName(connName); e == nil {
			return c, nil
		}
	}
	return cfg.FindDefault()
}

// dbConnOfActiveEnv 轻量读 config.json:debug.activeEnv → envs[].dbConn。
func dbConnOfActiveEnv(cfgPath string) string {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return ""
	}
	var c struct {
		Debug struct {
			ActiveEnv string `json:"activeEnv"`
			Envs      []struct {
				Name   string `json:"name"`
				DBConn string `json:"dbConn"`
			} `json:"envs"`
		} `json:"debug"`
	}
	if json.Unmarshal(data, &c) != nil {
		return ""
	}
	for _, e := range c.Debug.Envs {
		if e.Name == c.Debug.ActiveEnv {
			return e.DBConn
		}
	}
	return ""
}

// queryTableDictOnline rt 在线:直查远程 ERP 库的 24 张字典表(与离线 SQL 语义一致)。
// 当前实现金仓(PG)方言;oracle 在线查询属 P2,明确报错以免执行方言不兼容 SQL。
func queryTableDictOnline(name string) (*db.TableDict, error) {
	if liveDB == nil {
		return nil, errors.New("在线连接未就绪;请先执行 tdict rt --online(或检查连接配置)")
	}
	if liveDB.Dialect() == "oracle" {
		return nil, errors.New("oracle 在线查询尚属后续阶段;当前 --online 支持 kingbase(人大金仓)。" +
			"如需 oracle,请用 tdict db sync 同步到 SQLite 后离线查询,或等待后续版本")
	}
	ctx := context.Background()
	conn := liveDB.Connector()
	d := &db.TableDict{TableName: name}

	meta, err := liveQueryMeta(ctx, conn, name)
	if err != nil {
		return nil, err
	}
	if meta != nil {
		d.TableName = meta.TableName
		d.TableDesc = meta.TableDesc
		d.Module = meta.Module
		d.TableType = meta.TableType
	}
	fields, err := liveQueryFields(ctx, conn, name)
	if err != nil {
		return nil, err
	}
	d.Fields = fields
	keys, err := liveQueryKeys(ctx, conn, name)
	if err != nil && !isRemoteMissing(err) {
		return nil, err
	}
	if err == nil {
		d.Keys = keys
	}
	indexes, err := liveQueryIndexes(ctx, conn, name)
	if err != nil && !isRemoteMissing(err) {
		return nil, err
	}
	if err == nil {
		d.Indexes = indexes
	}
	return d, nil
}

// isRemoteMissing 判断"表/视图不存在"(键/索引字典表缺失时容忍,与离线 no such table 对齐)。
func isRemoteMissing(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	for _, kw := range []string{"does not exist", "ORA-00942", "no such table", "not found", "不存在"} {
		if containsFold(s, kw) {
			return true
		}
	}
	return false
}

func containsFold(s, sub string) bool {
	if len(sub) > len(s) {
		return false
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if eqFold(s[i:i+len(sub)], sub) {
			return true
		}
	}
	return false
}

func eqFold(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		ca, cb := a[i], b[i]
		if ca >= 'A' && ca <= 'Z' {
			ca += 'a' - 'A'
		}
		if cb >= 'A' && cb <= 'Z' {
			cb += 'a' - 'A'
		}
		if ca != cb {
			return false
		}
	}
	return true
}

// ---- 在线 SQL(金仓/PG 方言) ----
// erpdb 走 simple protocol(无绑定参数):所有值经 QuoteIdent/QuoteLit 白名单后内联。
// 表/列不在这些固定查询里内插用户自由文本,仅表名字面量需校验。
// schema 暂不限定(跟随连接当前 schema;与 db sync 的 user.table 一致时可后补 --schema)。

func liveQueryMeta(ctx context.Context, c erpdb.Connector, table string) (*db.TableMeta, error) {
	t := safeTableLit(table)
	sql := `SELECT a.dzea001, COALESCE(al.dzeal003, a.dzea002, ''), COALESCE(a.dzea003, ''), COALESCE(a.dzea004, '')
FROM dzea_t a
LEFT JOIN dzeal_t al ON al.dzeal001 = a.dzea001 AND al.dzeal002 = 'zh_CN'
WHERE a.dzea001 = ` + t
	_, rows, err := c.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	r := rows[0]
	m := &db.TableMeta{}
	if len(r) > 0 {
		m.TableName = r[0]
	}
	if len(r) > 1 {
		m.TableDesc = r[1]
	}
	if len(r) > 2 {
		m.Module = r[2]
	}
	if len(r) > 3 {
		m.TableType = r[3]
	}
	if m.TableName == "" {
		m.TableName = table
	}
	return m, nil
}

func liveQueryFields(ctx context.Context, c erpdb.Connector, table string) ([]db.TableInfo, error) {
	t := safeTableLit(table)
	sql := `SELECT COALESCE(b.dzeb021, ''), b.dzeb002,
       COALESCE(bl.dzebl003, b.dzeb003, ''),
       COALESCE(b.dzeb007, ''), COALESCE(b.dzeb008, ''),
       CASE WHEN COALESCE(b.dzeb004,'') = 'Y' THEN 'PK' ELSE '' END,
       CASE WHEN COALESCE(b.dzeb005,'') = 'Y' THEN 'Y' ELSE '' END,
       COALESCE(b.dzeb024, '')
FROM dzeb_t b
LEFT JOIN dzebl_t bl ON bl.dzebl001 = b.dzeb002 AND bl.dzebl002 = 'zh_CN'
WHERE b.dzeb001 = ` + t + `
ORDER BY CAST(b.dzeb021 AS INTEGER)`
	_, rows, err := c.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	out := make([]db.TableInfo, 0, len(rows))
	for _, r := range rows {
		var f db.TableInfo
		get := func(i int) string {
			if i < len(r) {
				return r[i]
			}
			return ""
		}
		f.Seq = get(0)
		f.FieldName = get(1)
		f.FieldDesc = get(2)
		f.DataType = get(3)
		f.Length = get(4)
		f.IsPK = get(5)
		f.Required = get(6)
		f.Remark = get(7)
		out = append(out, f)
	}
	return out, nil
}

func liveQueryKeys(ctx context.Context, c erpdb.Connector, table string) ([]db.KeyInfo, error) {
	t := safeTableLit(table)
	sql := `SELECT dzed002, COALESCE(dzed003, ''), COALESCE(dzed004, ''), COALESCE(dzed005, ''), COALESCE(dzed006, '')
FROM dzed_t WHERE dzed001 = ` + t + ` ORDER BY dzed002`
	_, rows, err := c.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	out := make([]db.KeyInfo, 0, len(rows))
	for _, r := range rows {
		var k db.KeyInfo
		get := func(i int) string {
			if i < len(r) {
				return r[i]
			}
			return ""
		}
		k.KeyName = get(0)
		k.KeyType = get(1)
		k.KeyField = get(2)
		k.RefTable = get(3)
		k.RefField = get(4)
		out = append(out, k)
	}
	return out, nil
}

func liveQueryIndexes(ctx context.Context, c erpdb.Connector, table string) ([]db.IndexInfo, error) {
	t := safeTableLit(table)
	sql := `SELECT dzec002, COALESCE(dzec003, ''), COALESCE(dzec004, '')
FROM dzec_t WHERE dzec001 = ` + t + ` ORDER BY dzec002`
	_, rows, err := c.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	out := make([]db.IndexInfo, 0, len(rows))
	for _, r := range rows {
		var ix db.IndexInfo
		get := func(i int) string {
			if i < len(r) {
				return r[i]
			}
			return ""
		}
		ix.IndexName = get(0)
		ix.IndexType = get(1)
		ix.IndexField = get(2)
		out = append(out, ix)
	}
	return out, nil
}

// safeTableLit 表名转 SQL 字面量:白名单校验通过则原样内联(单引号包裹),否则报错值。
func safeTableLit(t string) string {
	if !erpdb.ValidIdent(t) {
		return "''"
	}
	return "'" + t + "'"
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&useOnline, "online", false, "在线直查 ERP 数据库(当前支持 rt;需连接配置)")
	rootCmd.PersistentFlags().StringVar(&onlineConn, "conn", "", "在线查询连接名(--online 用;默认取 activeEnv.dbConn 或 isDefault 连接)")
}
