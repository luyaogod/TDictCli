package live

// 远程查询实现:让 *Live 实现 db.Source —— 与本地 SQLite 镜像(db 包)同一组
// 语义查询与 DTO,直接连 ERP 库执行。rt/rv/desc/scc/rq 的 SQL 以 db 包为蓝本
// (SELECT 列序一致),值经 QuoteLit 内联(erpdb simple protocol 无绑定参数)。
//
// 方言差异的集中处理:
//   - LIKE 过滤:SQLite 的 LIKE 默认不区分大小写,PG/金仓与 Oracle 区分;
//     远程统一 UPPER(列) LIKE UPPER('%kw%') 对齐本地体验;关键字为空时不加过滤子句。
//   - 排序的 CAST(x AS INTEGER):金仓已随 rt 实测通过;Oracle 真库验证见 README 说明。
//   - Oracle 空串即 NULL:COALESCE 兜底值在文本化扫描后仍是空串,表现一致。

import (
	"context"
	"fmt"
	"strconv"

	"tdict/db"
	"tdict/erpdb"
)

// q 执行单条只读 SQL(查询命令无需外部 ctx,沿用连接自身生命周期)。
func (l *Live) q(sql string) ([][]string, error) {
	_, rows, err := l.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	return rows, nil
}

// get 取行内第 i 列,越界/为空返回空串(与本地全 TEXT 镜像的 NULL→'' 一致)。
func get(row []string, i int) string {
	if i < len(row) {
		return row[i]
	}
	return ""
}

// lit 把用户值转 SQL 字面量(单引号翻倍);表名等标识符先过 ValidIdent 防注入。
func lit(s string) string { return erpdb.QuoteLit(s) }

// tableLit 表名(标识符形态,如 dzea_t)校验后转字面量;非法返回 "''"(查不到,不报错)。
func tableLit(t string) string {
	if !erpdb.ValidIdent(t) {
		return "''"
	}
	return lit(t)
}

// kwWhere 生成大小写不敏感的关键字过滤(本地 LIKE 不区分大小写;远程统一 UPPER)。
// 返回空串表示不加过滤。kw 为空与本地 `?='' OR ... LIKE '%%'` 语义等价。
func kwWhere(idExpr, descExpr, kw string) string {
	if kw == "" {
		return ""
	}
	k := lit("%" + kw + "%")
	return fmt.Sprintf(" WHERE (UPPER(%s) LIKE UPPER(%s) OR UPPER(COALESCE(%s, '')) LIKE UPPER(%s))",
		idExpr, k, descExpr, k)
}

// ---- rt:表字典 ----

// QueryTableMeta 表档 dzea_t:未收录返回 (nil, nil)。
func (l *Live) QueryTableMeta(table string) (*db.TableMeta, error) {
	t := tableLit(table)
	sql := `SELECT a.dzea001, COALESCE(al.dzeal003, a.dzea002, ''), a.dzea003, a.dzea004
FROM dzea_t a
LEFT JOIN dzeal_t al ON al.dzeal001 = a.dzea001 AND al.dzeal002 = 'zh_CN'
WHERE a.dzea001 = ` + t
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询表元数据 %s: %w", table, err)
	}
	if len(rows) == 0 {
		return nil, nil
	}
	r := rows[0]
	m := &db.TableMeta{TableName: table}
	m.TableDesc, m.Module, m.TableType = get(r, 1), get(r, 2), get(r, 3)
	if name := get(r, 0); name != "" {
		m.TableName = name
	}
	return m, nil
}

// QueryTable 字段档 dzeb_t(按显示顺序 dzeb021)。
func (l *Live) QueryTable(table string) ([]db.TableInfo, error) {
	t := tableLit(table)
	sql := `SELECT b.dzeb021, b.dzeb002,
       COALESCE(bl.dzebl003, b.dzeb003, ''),
       b.dzeb007, b.dzeb008,
       CASE WHEN b.dzeb004 = 'Y' THEN 'PK' ELSE '' END,
       CASE WHEN b.dzeb005 = 'Y' THEN 'Y' ELSE '' END,
       b.dzeb024
FROM dzeb_t b
LEFT JOIN dzebl_t bl ON bl.dzebl001 = b.dzeb002 AND bl.dzebl002 = 'zh_CN'
WHERE b.dzeb001 = ` + t + `
ORDER BY CAST(b.dzeb021 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询字段档 %s: %w", table, err)
	}
	out := make([]db.TableInfo, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.TableInfo{Seq: get(r, 0), FieldName: get(r, 1), FieldDesc: get(r, 2),
			DataType: get(r, 3), Length: get(r, 4), IsPK: get(r, 5), Required: get(r, 6), Remark: get(r, 7)})
	}
	return out, nil
}

// QueryKeys 键值档 dzed_t。
func (l *Live) QueryKeys(table string) ([]db.KeyInfo, error) {
	t := tableLit(table)
	sql := `SELECT dzed002, dzed003, dzed004,
       dzed005, dzed006
FROM dzed_t WHERE dzed001 = ` + t + ` ORDER BY dzed002`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询键值档 %s: %w", table, err)
	}
	out := make([]db.KeyInfo, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.KeyInfo{KeyName: get(r, 0), KeyType: get(r, 1), KeyField: get(r, 2),
			RefTable: get(r, 3), RefField: get(r, 4)})
	}
	return out, nil
}

// QueryIndexes 索引档 dzec_t。
func (l *Live) QueryIndexes(table string) ([]db.IndexInfo, error) {
	t := tableLit(table)
	sql := `SELECT dzec002, dzec003, dzec004
FROM dzec_t WHERE dzec001 = ` + t + ` ORDER BY dzec002`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询索引档 %s: %w", table, err)
	}
	out := make([]db.IndexInfo, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.IndexInfo{IndexName: get(r, 0), IndexType: get(r, 1), IndexField: get(r, 2)})
	}
	return out, nil
}

// ---- rv:校验带值 dzcd_t/dzce_t/dzch_t ----

// QueryCheckList 校验定义列表(可 --kw 过滤)。
func (l *Live) QueryCheckList(lang, keyword string) ([]db.CheckListRow, error) {
	sql := `SELECT c.dzcd001, c.dzcd002, cl.dzcdl003,
       c.dzcd005, c.dzcd004, cl.dzcdl004
FROM dzcd_t c
LEFT JOIN dzcdl_t cl ON cl.dzcdl001 = c.dzcd001 AND cl.dzcdl002 = ` + lit(lang) +
		kwWhere("c.dzcd001", "cl.dzcdl003", keyword) +
		` ORDER BY c.dzcd001, CASE c.dzcd002 WHEN 's' THEN 0 ELSE 1 END`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询校验列表: %w", err)
	}
	out := make([]db.CheckListRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.CheckListRow{ID: get(r, 0), Cust: get(r, 1), Desc: get(r, 2),
			TypeCode: get(r, 3), ErrMsg: get(r, 4), Remark: get(r, 5)})
	}
	return out, nil
}

// QueryCheckHeaders 单个校验的全部头(标准/客制)。
func (l *Live) QueryCheckHeaders(id, lang string) ([]db.CheckHeaderRow, error) {
	sql := `SELECT c.dzcd001, c.dzcd002, c.dzcd005,
       c.dzcd004, c.dzcd006, c.dzcdstus,
       cl.dzcdl003, cl.dzcdl004, c.dzcd003
FROM dzcd_t c
LEFT JOIN dzcdl_t cl ON cl.dzcdl001 = c.dzcd001 AND cl.dzcdl002 = ` + lit(lang) + `
WHERE c.dzcd001 = ` + lit(id) + `
ORDER BY CASE c.dzcd002 WHEN 's' THEN 0 ELSE 1 END`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询校验头 %s: %w", id, err)
	}
	out := make([]db.CheckHeaderRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.CheckHeaderRow{ID: get(r, 0), Cust: get(r, 1), TypeCode: get(r, 2),
			ErrMsg: get(r, 3), Industry: get(r, 4), Status: get(r, 5),
			Desc: get(r, 6), Remark: get(r, 7), SQL: get(r, 8)})
	}
	return out, nil
}

// QueryCheckParams 校验的外部参数(argN)定义。
func (l *Live) QueryCheckParams(id, lang string) ([]db.CheckParamRow, error) {
	sql := `SELECT e.dzce003, e.dzce002, e.dzce004,
       e.dzcestus, el.dzcel004, el.dzcel005
FROM dzce_t e
LEFT JOIN dzcel_t el ON el.dzcel001 = e.dzce001 AND el.dzcel002 = e.dzce002 AND el.dzcel003 = ` + lit(lang) + `
WHERE e.dzce001 = ` + lit(id) + `
ORDER BY CAST(e.dzce002 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询校验参数 %s: %w", id, err)
	}
	out := make([]db.CheckParamRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.CheckParamRow{Cust: get(r, 0), Seq: get(r, 1), Name: get(r, 2),
			DateType: get(r, 3), Desc: get(r, 4), Remark: get(r, 5)})
	}
	return out, nil
}

// QueryCheckConds 校验的存在条件明细。
func (l *Live) QueryCheckConds(id string) ([]db.CheckCondRow, error) {
	sql := `SELECT h.dzch005, h.dzch002, h.dzch003, h.dzch004
FROM dzch_t h
WHERE h.dzch001 = ` + lit(id) + `
ORDER BY CAST(h.dzch002 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询校验条件 %s: %w", id, err)
	}
	out := make([]db.CheckCondRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.CheckCondRow{Cust: get(r, 0), Seq: get(r, 1), Cond: get(r, 2), ErrMsg: get(r, 3)})
	}
	return out, nil
}

// ---- desc:字段规格 dzep_t ----

// QuerySpecRows 单表全部字段规格行(按 dzeb021 顺序)。
func (l *Live) QuerySpecRows(tableName, lang string) ([]db.SpecRow, error) {
	t := tableLit(tableName)
	// dzeb_t 的排序号(dzeb021)在 Oracle 是真 NUMBER、金仓/镜像为文本,
	// COALESCE 兜底字面量须与列类型一致,否则 Oracle 报 ORA-00932。
	seq := "COALESCE(eb.dzeb021, '0')"
	if l.conn.Type() == "oracle" {
		seq = "COALESCE(eb.dzeb021, 0)"
	}
	sql := `SELECT eb.dzeb021, ep.dzep002,
       ebl.dzebl003, ebl.dzebl004,
       ep.dzep010, ep.dzep011,
       ep.dzep005, ep.dzep009,
       ep.dzep021, ep.dzep023,
       ep.dzep012, ep.dzep025, ep.dzep013,
       ep.dzep026, ep.dzep014,
       ep.dzep017, ep.dzep018,
       ep.dzep019, ep.dzep022, ep.dzep020,
       ep.dzep027, ep.dzep028,
       ep.dzepstus
FROM dzep_t ep
LEFT JOIN dzeb_t eb ON eb.dzeb001 = ep.dzep001 AND eb.dzeb002 = ep.dzep002
LEFT JOIN dzebl_t ebl ON ebl.dzebl001 = ep.dzep002 AND ebl.dzebl002 = ` + lit(lang) + `
WHERE ep.dzep001 = ` + t + `
ORDER BY CAST(` + seq + ` AS INTEGER), ep.dzep002`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询字段规格 %s: %w", tableName, err)
	}
	out := make([]db.SpecRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.SpecRow{Seq: get(r, 0), Field: get(r, 1), FieldName: get(r, 2),
			FieldDesc: get(r, 3), Widget: get(r, 4), Scc: get(r, 5), Required: get(r, 6),
			Width: get(r, 7), Format: get(r, 8), CaseConv: get(r, 9), DefaultVal: get(r, 10),
			MaxSym: get(r, 11), MaxVal: get(r, 12), MinSym: get(r, 13), MinVal: get(r, 14),
			OpenEdit: get(r, 15), OpenQuery: get(r, 16), ChkVal: get(r, 17),
			LookupType: get(r, 18), LookupProg: get(r, 19), RepWidth: get(r, 20),
			RepDigit: get(r, 21), Env: get(r, 22)})
	}
	return out, nil
}

// ---- scc:系统分类码 gzca_t/gzcb_t ----

// QuerySccList 分类码列表(可 --kw 过滤),附值数量(COUNT 标量子查询)。
func (l *Live) QuerySccList(lang, keyword string) ([]db.SccListRow, error) {
	sql := `SELECT a.gzca001, a.gzca004, a.gzcastus, al.gzcal003,
       (SELECT COUNT(*) FROM gzcb_t WHERE gzcb001 = a.gzca001)
FROM gzca_t a
LEFT JOIN gzcal_t al ON al.gzcal001 = a.gzca001 AND al.gzcal002 = ` + lit(lang) +
		kwWhere("a.gzca001", "al.gzcal003", keyword) +
		` ORDER BY CAST(a.gzca001 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询分类码列表: %w", err)
	}
	out := make([]db.SccListRow, 0, len(rows))
	for _, r := range rows {
		cnt, err := strconv.Atoi(get(r, 4))
		if err != nil {
			return nil, fmt.Errorf("解析分类值数量 %q: %w", get(r, 4), err)
		}
		out = append(out, db.SccListRow{ID: get(r, 0), Group: get(r, 1), Status: get(r, 2),
			Name: get(r, 3), ValCnt: cnt})
	}
	return out, nil
}

// QuerySccHeader 单个分类码的头;未命中返回 (nil, nil)。
func (l *Live) QuerySccHeader(id, lang string) (*db.SccHeaderRow, error) {
	sql := `SELECT a.gzca001, a.gzca004, a.gzcastus,
       a.gzca002, a.gzca003,
       al.gzcal003, al.gzcal004, al.gzcal005
FROM gzca_t a
LEFT JOIN gzcal_t al ON al.gzcal001 = a.gzca001 AND al.gzcal002 = ` + lit(lang) + `
WHERE a.gzca001 = ` + lit(id)
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询分类码头 %s: %w", id, err)
	}
	if len(rows) == 0 {
		return nil, nil
	}
	r := rows[0]
	return &db.SccHeaderRow{ID: get(r, 0), Group: get(r, 1), Status: get(r, 2),
		Aux2: get(r, 3), Aux3: get(r, 4), Name: get(r, 5), Desc: get(r, 6), Desc2: get(r, 7)}, nil
}

// QuerySccValues 单个分类码的全部值(按 gzcb012 顺序)。
func (l *Live) QuerySccValues(id, lang string) ([]db.SccValueRow, error) {
	sql := `SELECT b.gzcb002, b.gzcb012, b.gzcb013,
       bl.gzcbl004, bl.gzcbl005,
       b.gzcb003, b.gzcb004, b.gzcb005,
       b.gzcb006, b.gzcb007, b.gzcb008,
       b.gzcb009, b.gzcb010, b.gzcb011,
       b.gzcb014, b.gzcb015
FROM gzcb_t b
LEFT JOIN gzcbl_t bl ON bl.gzcbl001 = b.gzcb001 AND bl.gzcbl002 = b.gzcb002 AND bl.gzcbl003 = ` + lit(lang) + `
WHERE b.gzcb001 = ` + lit(id) + `
ORDER BY CAST(b.gzcb012 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询分类值 %s: %w", id, err)
	}
	out := make([]db.SccValueRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.SccValueRow{Val: get(r, 0), Sort: get(r, 1), Cust: get(r, 2),
			Desc: get(r, 3), Desc2: get(r, 4), B3: get(r, 5), B4: get(r, 6), B5: get(r, 7),
			B6: get(r, 8), B7: get(r, 9), B8: get(r, 10), B9: get(r, 11), B10: get(r, 12),
			B11: get(r, 13), B14: get(r, 14), B15: get(r, 15)})
	}
	return out, nil
}

// ---- rq:可复用开窗 dzca_t/dzcb_t/dzcc_t ----

// QueryWinList 开窗列表(可 --kw 过滤)。
func (l *Live) QueryWinList(lang, keyword string) ([]db.WinListRow, error) {
	sql := `SELECT a.dzca001, a.dzca002, al.dzcal003,
       a.dzcastus, a.dzca004,
       a.dzca006, a.dzca007
FROM dzca_t a
LEFT JOIN dzcal_t al ON al.dzcal001 = a.dzca001 AND al.dzcal002 = ` + lit(lang) +
		kwWhere("a.dzca001", "al.dzcal003", keyword) +
		` ORDER BY a.dzca001, CASE a.dzca002 WHEN 's' THEN 0 ELSE 1 END`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询开窗列表: %w", err)
	}
	out := make([]db.WinListRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.WinListRow{ID: get(r, 0), Cust: get(r, 1), Desc: get(r, 2),
			Status: get(r, 3), PageSize: get(r, 4), HardCode: get(r, 5), Industry: get(r, 6)})
	}
	return out, nil
}

// QueryWinHeaders 单个开窗的全部头(标准/客制)。
func (l *Live) QueryWinHeaders(id, lang string) ([]db.WinHeaderRow, error) {
	sql := `SELECT a.dzca001, a.dzca002, a.dzcastus,
       a.dzca003, a.dzca004, a.dzca005,
       a.dzca006, a.dzca007,
       al.dzcal003, al.dzcal004
FROM dzca_t a
LEFT JOIN dzcal_t al ON al.dzcal001 = a.dzca001 AND al.dzcal002 = ` + lit(lang) + `
WHERE a.dzca001 = ` + lit(id) + `
ORDER BY CASE a.dzca002 WHEN 's' THEN 0 ELSE 1 END`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询开窗头 %s: %w", id, err)
	}
	out := make([]db.WinHeaderRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.WinHeaderRow{ID: get(r, 0), Cust: get(r, 1), Status: get(r, 2),
			SQL: get(r, 3), PageSize: get(r, 4), Serial: get(r, 5), HardCode: get(r, 6),
			Industry: get(r, 7), Desc: get(r, 8), Memo: get(r, 9)})
	}
	return out, nil
}

// QueryWinParams 开窗的外部参数(argN)定义。
func (l *Live) QueryWinParams(id, lang string) ([]db.WinParamRow, error) {
	sql := `SELECT b.dzcb004, b.dzcb002, b.dzcb003,
       b.dzcbstus, bl.dzcbl004, bl.dzcbl005
FROM dzcb_t b
LEFT JOIN dzcbl_t bl ON bl.dzcbl001 = b.dzcb001 AND bl.dzcbl002 = b.dzcb002 AND bl.dzcbl003 = ` + lit(lang) + `
WHERE b.dzcb001 = ` + lit(id) + `
ORDER BY CAST(b.dzcb002 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询开窗参数 %s: %w", id, err)
	}
	out := make([]db.WinParamRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.WinParamRow{Cust: get(r, 0), Seq: get(r, 1), Name: get(r, 2),
			DateType: get(r, 3), Desc: get(r, 4), Memo: get(r, 5)})
	}
	return out, nil
}

// QueryWinCols 开窗的显现列设置。
func (l *Live) QueryWinCols(id string) ([]db.WinColRow, error) {
	sql := `SELECT c.dzcc009, c.dzcc002, c.dzcc003,
       c.dzcc008, c.dzcc004, c.dzcc005,
       c.dzcc006, c.dzcc010, c.dzcc007
FROM dzcc_t c
WHERE c.dzcc001 = ` + lit(id) + `
ORDER BY CAST(c.dzcc002 AS INTEGER)`
	rows, err := l.q(sql)
	if err != nil {
		return nil, fmt.Errorf("查询开窗列 %s: %w", id, err)
	}
	out := make([]db.WinColRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, db.WinColRow{Cust: get(r, 0), Seq: get(r, 1), Field: get(r, 2),
			Alias: get(r, 3), Widget: get(r, 4), IsRet: get(r, 5), CaseConv: get(r, 6),
			Format: get(r, 7), Label: get(r, 8)})
	}
	return out, nil
}
