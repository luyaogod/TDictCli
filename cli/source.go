package cli

// 查询数据源(rt/rv/desc/scc/rq 共用):本地 SQLite 镜像(*db.DB)与远程 ERP 库
// (*live.Live,金仓/Oracle)实现同一 db.Source 接口,数据源由配置/CLI 决定,
// 命令层不感知差异 —— 不再有 --online 开关。
//
// 选择优先级:
//  1. --conn <环境名|local>(覆盖本次调用);
//  2. config.json 顶层 query.source(与 debug 键平级,serve 保存设置不动它);
//  3. 默认 "local"(纯本地用法与旧版一致,无需任何配置)。
//
// "local" = 现有 erp_data.db(--db / TDICT_DB);环境名 = debug.sshs 中该环境的
// db(客户端直连,凭据取账号列表首项)。

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"tdict/db"
	"tdict/dbconfig"
	"tdict/debug"
	"tdict/live"
)

var (
	srcConn  string // root --conn:数据源(环境名或 local;空=按配置)
	dataSrc  db.Source
	srcLocal bool // 当前源为本地 SQLite(错误提示文案分流)
)

// openQuerySource 在 root PersistentPreRunE 打开数据源(本地库或远程连接)。
func openQuerySource() error {
	target := srcConn
	if target == "" {
		target = queryCfgSource()
	}
	if target == "" || target == "local" {
		return openLocalSource()
	}
	return openRemoteSource(target)
}

// closeQuerySource 在 root PersistentPostRun 关闭数据源。
func closeQuerySource() {
	if dataSrc != nil {
		dataSrc.Close()
		dataSrc = nil
	}
}

// queryCfgSource 读 config.json 顶层 query.source("local" 或环境名);
// 配置文件缺失/未配置返回 ""(即默认本地)。
func queryCfgSource() string {
	path, err := resolveConfigPath(configPath)
	if err != nil {
		return ""
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	var root struct {
		Query *struct {
			Source string `json:"source"`
		} `json:"query"`
	}
	if json.Unmarshal(data, &root) != nil || root.Query == nil {
		return ""
	}
	return root.Query.Source
}

// openLocalSource 打开本地 SQLite 镜像(-d/TDICT_DB 定位,要求文件已存在)。
func openLocalSource() error {
	resolvedPath, err := resolveDBPath(dbPath)
	if err != nil {
		return err
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "[tdict] 数据源: 本地 SQLite (%s)\n", resolvedPath)
	}
	database, err := db.Open(resolvedPath)
	if err != nil {
		return fmt.Errorf("failed to open database at %s: %w", resolvedPath, err)
	}
	dataSrc = database
	srcLocal = true
	return nil
}

// openRemoteSource 打开指定 SSH 环境的远程库(客户端直连,首账号)。
func openRemoteSource(target string) error {
	path, err := resolveConfigPath(configPath)
	if err != nil {
		return err
	}
	cfg, err := debug.LoadConfig(path)
	if err != nil {
		return err
	}
	cc, err := envDBByName(cfg, target)
	if err != nil {
		return err
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "[tdict] 数据源: 远程 %s (%s %s)\n", target, cc.Type, cc.Address())
	}
	l, err := live.Open(context.Background(), *cc)
	if err != nil {
		return fmt.Errorf("打开远程数据源 %q: %w", target, err)
	}
	dataSrc = l
	srcLocal = false
	return nil
}

// envDBByName 在 debug.sshs 中按环境名取该环境 db 的深拷贝。
func envDBByName(cfg *debug.Config, name string) (*dbconfig.Connection, error) {
	for i := range cfg.SSHs {
		e := &cfg.SSHs[i]
		if e.Name != name {
			continue
		}
		if e.DB == nil {
			return nil, fmt.Errorf("环境 %q 未配置数据库(设置-环境-数据库页添加后可用 --conn 直查)", name)
		}
		cc := *e.DB
		cc.Accounts = append([]dbconfig.DBAcct(nil), e.DB.Accounts...)
		return &cc, nil
	}
	return nil, fmt.Errorf("未找到环境 %q(可用: tdict debug env 查看环境名)", name)
}

// missingHint 主字典表缺失(IsMissingTable)时的提示:本地源建议先 db sync;
// 远程源说明库的问题与回退方式。subject 如 "校验定义 (dzcd_t 等表)"。
func missingHint(subject string) string {
	if srcLocal {
		return fmt.Sprintf("本地库尚未包含%s数据。请先在有数据库的环境执行: tdict db sync", subject)
	}
	return fmt.Sprintf("远程库缺少%s相关字典表(该环境数据库未同步或账号权限不足);可用 --conn local 切回本地库", subject)
}

// emptyHint 查询零结果时的提示:本地源沿用"先 db sync"建议,远程源说明语义。
// subject 如 "校验定义"。
func emptyHint(subject string) string {
	if srcLocal {
		return fmt.Sprintf("暂无%s数据。请先执行: tdict db sync", subject)
	}
	return fmt.Sprintf("未查询到%s数据(库中无记录或 --kw 无匹配;可换 --conn <其他环境> 再试)", subject)
}

func init() {
	rootCmd.PersistentFlags().StringVar(&srcConn, "conn", "",
		"查询数据源: local(本地 SQLite)或 SSH 环境名(远程直查该环境数据库);默认取 config.json query.source,缺省 local")
}
