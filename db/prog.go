package db

import (
	"database/sql"
	"fmt"
)

// 程序与作业字典(azzi900 程式基本資料設定作業 / azzi910 作業基本資料維護):
//
//	gzza_t  程序档:gzza001 程序编号(PK)、gzza002 程序类别、gzza003 归属模块、
//	        gzza008 引用主程序编号、gzza011 客制
//	gzzal_t 程序名称多语言:gzzal001 程序编号 + gzzal002 语言别 → gzzal003 程序名称
//	gzzz_t  作业编号设置表:gzzz001 作业编号(PK)、gzzz002 程序编号(作业挂的程序,
//	        一个程序可被多个作业使用)、gzzz003 应用参数组编号、gzzz005 归属模块
//	gzzk_t  程序应用参数组设置表:gzzk001 程序编号 + gzzk002 参数组编号 → gzzk003 说明
//
// 作业的显示名称既不是独立字段也不是独立表:工具链(azzi910)用
// gzzal_t.gzzal003 按 gzzz002 关联取得,即"作业名称 = 它挂的程序名称"。

// ProgInfo 一个编号作为"程序"的登记信息;若该编号同时是"作业"也一并给出它挂的程序。
type ProgInfo struct {
	Code      string `json:"编号"`
	IsProg    bool   `json:"是程序"`
	Name      string `json:"程序名称"`   // gzzal003(指定语言)
	ShortName string `json:"程序简称"`   // gzzal005
	Category  string `json:"程序类别"`   // gzza002
	Module    string `json:"归属模块"`   // gzza003
	Cust      string `json:"客制"`     // gzza011
	RefMain   string `json:"引用主程序"`  // gzza008
	RunCmd    string `json:"系统运行指令"` // gzza004
	Status    string `json:"状态码"`    // gzzastus
	IsJob     bool   `json:"是作业"`
	JobProg   string `json:"作业挂的程序"` // gzzz002(该编号作为作业登记时)
}

// ProgJob 一个使用某程序的作业(gzzz_t 按 gzzz002 关联到程序)。
type ProgJob struct {
	JobCode   string `json:"作业编号"`    // gzzz001
	JobName   string `json:"作业名称"`    // gzzal003(经 gzzz002 关联的程序名称)
	Module    string `json:"归属模块"`    // gzzz005
	ParamGrp  string `json:"应用参数组"`   // gzzz003
	ParamDesc string `json:"参数组说明"`   // gzzk003(经 gzzk001/gzzk002)
	DocType   string `json:"默认单据性质"`  // gzzz006
	Status    string `json:"状态码"`     // gzzzstus
}

// ProgListItem 程序列表/搜索(prog --kw)的一行。
type ProgListItem struct {
	Code     string `json:"程序编号"`
	Name     string `json:"程序名称"`
	Category string `json:"程序类别"`
	Module   string `json:"归属模块"`
	Cust     string `json:"客制"`
	JobCount int    `json:"作业数"`
}

// QueryProgInfo 按编号查程序登记信息(附"该编号是否也是作业、挂的哪个程序")。
// 程序与作业都未收录时返回 (nil, nil)。
func (d *DB) QueryProgInfo(code, lang string) (*ProgInfo, error) {
	info := &ProgInfo{Code: code}

	row := d.conn.QueryRow(`
		SELECT COALESCE(a.gzza002, ''), COALESCE(a.gzza003, ''), COALESCE(a.gzza004, ''),
		       COALESCE(a.gzza008, ''), COALESCE(a.gzza011, ''), COALESCE(a.gzzastus, ''),
		       COALESCE(l.gzzal003, ''), COALESCE(l.gzzal005, '')
		FROM gzza_t a
		LEFT JOIN gzzal_t l ON l.gzzal001 = a.gzza001 AND l.gzzal002 = ?
		WHERE a.gzza001 = ?`, lang, code)
	var category, module, runCmd, refMain, cust, status, name, shortName string
	switch err := row.Scan(&category, &module, &runCmd, &refMain, &cust, &status, &name, &shortName); err {
	case nil:
		info.IsProg = true
		info.Category, info.Module, info.RunCmd = category, module, runCmd
		info.RefMain, info.Cust, info.Status = refMain, cust, status
		info.Name, info.ShortName = name, shortName
	case sql.ErrNoRows:
		// 不是程序,再看看是不是作业
	default:
		return nil, fmt.Errorf("query prog info %s: %w", code, err)
	}

	if err := d.conn.QueryRow(
		`SELECT COALESCE(gzzz002, '') FROM gzzz_t WHERE gzzz001 = ?`, code).Scan(&info.JobProg); err == nil {
		info.IsJob = true
	} else if err != sql.ErrNoRows {
		return nil, fmt.Errorf("query job info %s: %w", code, err)
	}

	if !info.IsProg && !info.IsJob {
		return nil, nil
	}
	return info, nil
}

// QueryProgJobs 查"哪些作业用了这个程序"(gzzz_t.gzzz002 = code)。
// 作业名称取它所挂程序的名称(gzzal_t),与 azzi910 的取法一致。
func (d *DB) QueryProgJobs(code, lang string) ([]ProgJob, error) {
	rows, err := d.conn.Query(`
		SELECT z.gzzz001,
		       COALESCE(l.gzzal003, ''),
		       COALESCE(z.gzzz005, ''),
		       COALESCE(z.gzzz003, ''),
		       COALESCE(k.gzzk003, ''),
		       COALESCE(z.gzzz006, ''),
		       COALESCE(z.gzzzstus, '')
		FROM gzzz_t z
		LEFT JOIN gzzal_t l ON l.gzzal001 = z.gzzz002 AND l.gzzal002 = ?
		LEFT JOIN gzzk_t k ON k.gzzk001 = z.gzzz002 AND k.gzzk002 = z.gzzz003
		WHERE z.gzzz002 = ?
		ORDER BY z.gzzz001`, lang, code)
	if err != nil {
		return nil, fmt.Errorf("query prog jobs %s: %w", code, err)
	}
	defer rows.Close()

	var result []ProgJob
	for rows.Next() {
		var j ProgJob
		if err := rows.Scan(&j.JobCode, &j.JobName, &j.Module, &j.ParamGrp,
			&j.ParamDesc, &j.DocType, &j.Status); err != nil {
			return nil, fmt.Errorf("scan prog job: %w", err)
		}
		result = append(result, j)
	}
	return result, rows.Err()
}

// QueryProgList 列出/搜索程序(--kw 按程序编号或程序名称过滤),附各自的作业数。
func (d *DB) QueryProgList(lang, keyword string) ([]ProgListItem, error) {
	like := "%" + keyword + "%"
	rows, err := d.conn.Query(`
		SELECT a.gzza001,
		       COALESCE(l.gzzal003, ''),
		       COALESCE(a.gzza002, ''),
		       COALESCE(a.gzza003, ''),
		       COALESCE(a.gzza011, ''),
		       (SELECT COUNT(*) FROM gzzz_t z WHERE z.gzzz002 = a.gzza001)
		FROM gzza_t a
		LEFT JOIN gzzal_t l ON l.gzzal001 = a.gzza001 AND l.gzzal002 = ?
		WHERE (? = '' OR a.gzza001 LIKE ? OR COALESCE(l.gzzal003, '') LIKE ?)
		ORDER BY a.gzza001`, lang, keyword, like, like)
	if err != nil {
		return nil, fmt.Errorf("query prog list: %w", err)
	}
	defer rows.Close()

	var result []ProgListItem
	for rows.Next() {
		var p ProgListItem
		if err := rows.Scan(&p.Code, &p.Name, &p.Category, &p.Module, &p.Cust, &p.JobCount); err != nil {
			return nil, fmt.Errorf("scan prog row: %w", err)
		}
		result = append(result, p)
	}
	return result, rows.Err()
}
