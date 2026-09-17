package db

import (
	"database/sql"
	"path/filepath"
	"testing"
)

// buildProgDB 建一个只含"程序与作业"四张表的临时库,列全 TEXT(与本地镜像同构)。
func buildProgDB(t *testing.T) *DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "prog.db")
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("open temp db: %v", err)
	}
	stmts := []string{
		`CREATE TABLE gzza_t (gzza001 TEXT, gzza002 TEXT, gzza003 TEXT, gzza004 TEXT,
		   gzza008 TEXT, gzza011 TEXT, gzzastus TEXT)`,
		`CREATE TABLE gzzal_t (gzzal001 TEXT, gzzal002 TEXT, gzzal003 TEXT, gzzal005 TEXT)`,
		`CREATE TABLE gzzz_t (gzzz001 TEXT, gzzz002 TEXT, gzzz003 TEXT, gzzz005 TEXT,
		   gzzz006 TEXT, gzzzstus TEXT)`,
		`CREATE TABLE gzzk_t (gzzk001 TEXT, gzzk002 TEXT, gzzk003 TEXT)`,
		// aooi301:被两个作业使用的共用维护程序;aapi011:作业编号=程序编号
		`INSERT INTO gzza_t VALUES ('aooi301','i','AOO','$FGLRUN $AOOi/aooi301','','s','Y')`,
		`INSERT INTO gzza_t VALUES ('aapi011','i','AAP','$FGLRUN $AAPi/aapi011','','s','Y')`,
		`INSERT INTO gzzal_t VALUES ('aooi301','zh_CN','应用分类码维护作业(单档多栏)','YYFLMWH')`,
		`INSERT INTO gzzal_t VALUES ('aapi011','zh_CN','应付账款类别依账套设置科目作业','YFZKLB')`,
		`INSERT INTO gzzal_t VALUES ('aooi301','zh_TW','應用分類碼維護作業','YYFLMWH')`,
		`INSERT INTO gzzz_t VALUES ('aapi011','aapi011','0','AAP','','Y')`,
		`INSERT INTO gzzz_t VALUES ('aooi701','aooi301','0','AOO','','Y')`,
		`INSERT INTO gzzz_t VALUES ('aooi707','aooi301','1','AOO','','Y')`,
		`INSERT INTO gzzk_t VALUES ('aooi301','0','预设参数组')`,
		`INSERT INTO gzzk_t VALUES ('aooi301','1','单档多栏参数组')`,
	}
	for _, s := range stmts {
		if _, err := conn.Exec(s); err != nil {
			t.Fatalf("setup (%s): %v", s, err)
		}
	}
	if err := conn.Close(); err != nil {
		t.Fatalf("close setup conn: %v", err)
	}

	d, err := Open(path)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(d.Close)
	return d
}

// TestQueryProgInfo 程序编号与作业编号两种形态,以及未收录。
func TestQueryProgInfo(t *testing.T) {
	d := buildProgDB(t)

	// 程序:aooi301
	info, err := d.QueryProgInfo("aooi301", "zh_CN")
	if err != nil {
		t.Fatalf("QueryProgInfo(aooi301): %v", err)
	}
	if info == nil || !info.IsProg {
		t.Fatalf("aooi301 应登记为程序, got %+v", info)
	}
	if info.Name != "应用分类码维护作业(单档多栏)" || info.Module != "AOO" || info.Category != "i" {
		t.Errorf("aooi301 字段不符: %+v", info)
	}
	if info.IsJob {
		t.Errorf("aooi301 不是作业: %+v", info)
	}

	// 繁体名称:换 --lang 应取到繁体那一行
	tw, err := d.QueryProgInfo("aooi301", "zh_TW")
	if err != nil {
		t.Fatalf("QueryProgInfo(zh_TW): %v", err)
	}
	if tw == nil || tw.Name != "應用分類碼維護作業" {
		t.Errorf("zh_TW 名称不符: %+v", tw)
	}

	// 作业编号 != 程序编号:aooi701 挂 aooi301,自身不是程序
	job, err := d.QueryProgInfo("aooi701", "zh_CN")
	if err != nil {
		t.Fatalf("QueryProgInfo(aooi701): %v", err)
	}
	if job == nil || job.IsProg || !job.IsJob || job.JobProg != "aooi301" {
		t.Errorf("aooi701 应只是作业且挂 aooi301: %+v", job)
	}

	// 未收录:返回 nil,nil(不报错)
	none, err := d.QueryProgInfo("nosuch01", "zh_CN")
	if err != nil || none != nil {
		t.Errorf("未收录应返回 (nil,nil), got %+v, %v", none, err)
	}
}

// TestQueryProgJobs 一个程序被多个作业使用,且参数组说明按 gzzk 关联带出。
func TestQueryProgJobs(t *testing.T) {
	d := buildProgDB(t)

	jobs, err := d.QueryProgJobs("aooi301", "zh_CN")
	if err != nil {
		t.Fatalf("QueryProgJobs: %v", err)
	}
	if len(jobs) != 2 {
		t.Fatalf("aooi301 应有 2 个作业, got %d: %+v", len(jobs), jobs)
	}
	if jobs[0].JobCode != "aooi701" || jobs[1].JobCode != "aooi707" {
		t.Errorf("作业应按编号升序: %+v", jobs)
	}
	// 作业名称取所挂程序的名称(azzi910 的取法)
	if jobs[0].JobName != "应用分类码维护作业(单档多栏)" {
		t.Errorf("作业名称应取程序名称: %+v", jobs[0])
	}
	if jobs[0].ParamGrp != "0" || jobs[0].ParamDesc != "预设参数组" {
		t.Errorf("应用参数组关联有误: %+v", jobs[0])
	}
	if jobs[1].ParamGrp != "1" || jobs[1].ParamDesc != "单档多栏参数组" {
		t.Errorf("第二个作业的参数组关联有误: %+v", jobs[1])
	}

	only, err := d.QueryProgJobs("aapi011", "zh_CN")
	if err != nil {
		t.Fatalf("QueryProgJobs(aapi011): %v", err)
	}
	if len(only) != 1 || only[0].JobCode != "aapi011" {
		t.Errorf("aapi011 应只有 1 个作业: %+v", only)
	}
}

// TestQueryProgList 列表带作业数,--kw 按编号或中文名称过滤。
func TestQueryProgList(t *testing.T) {
	d := buildProgDB(t)

	all, err := d.QueryProgList("zh_CN", "")
	if err != nil {
		t.Fatalf("QueryProgList: %v", err)
	}
	if len(all) != 2 {
		t.Fatalf("应有 2 个程序, got %d", len(all))
	}
	if all[0].Code != "aapi011" || all[0].JobCount != 1 {
		t.Errorf("aapi011 作业数应为 1: %+v", all[0])
	}
	if all[1].Code != "aooi301" || all[1].JobCount != 2 {
		t.Errorf("aooi301 作业数应为 2(一个程序被多个作业使用): %+v", all[1])
	}

	byName, err := d.QueryProgList("zh_CN", "分类码")
	if err != nil {
		t.Fatalf("QueryProgList(按名称): %v", err)
	}
	if len(byName) != 1 || byName[0].Code != "aooi301" {
		t.Errorf("按中文名称搜索应命中 aooi301: %+v", byName)
	}

	byCode, err := d.QueryProgList("zh_CN", "aapi")
	if err != nil {
		t.Fatalf("QueryProgList(按编号): %v", err)
	}
	if len(byCode) != 1 || byCode[0].Code != "aapi011" {
		t.Errorf("按编号搜索应命中 aapi011: %+v", byCode)
	}

	if none, err := d.QueryProgList("zh_CN", "无此程序"); err != nil || len(none) != 0 {
		t.Errorf("无匹配应返回空: %+v, %v", none, err)
	}
}
