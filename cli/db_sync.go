package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"tdict/db"
	"tdict/erpdb"

	"github.com/spf13/cobra"
)

var (
	dbSyncConn  string
	dbSyncTable string
)

// dictTables are the T100 data dictionary tables synced from the ERP database.
var dictTables = []string{
	"dzea_t", "dzeal_t", "dzeb_t", "dzebl_t", "dzec_t", "dzed_t", "dzee_t", "dzef_t", "dzeg_t",
	// 校验带值 (r.v) 体系: 校验定义/多语言/外部参数/参数多语言/判断条件
	"dzcd_t", "dzcdl_t", "dzce_t", "dzcel_t", "dzch_t",
	// 系统分类码 (SCC): 分类码头档/多语言/分类值/值多语言
	"gzca_t", "gzcal_t", "gzcb_t", "gzcbl_t",
	// 字段规格 (画面设计器参考配置): 控件/SCC码/格式/必填等
	"dzep_t",
	// 可复用开窗 (r.q): 开窗主档/多语言/设计参数/参数多语言/显现设定
	"dzca_t", "dzcal_t", "dzcb_t", "dzcbl_t", "dzcc_t",
	// 系统消息档 (azzi920 维护, cl_err/cl_getmsg 取用): 消息文本/建议处理/作业多语言名称
	"gzze_t", "gzzal_t",
	// 参数定义档 (azzi990 系统参数 / azzi991 单据别参数): 定义/参数多语言/单据性质绑定
	"gzsz_t", "gzszl_t", "gzsy_t",
}

type syncResult struct {
	table   string
	rows    int
	elapsed time.Duration
}

var dbSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "从 ERP 刷新本地查询数据",
	Long: `从 ERP 刷新本地查询数据(表字典、校验、分类码、画面规格、开窗、消息、参数等
全部内容,约 85 万行),查询命令读的就是它。
写入 -d/--db 或 TDICT_DB 指向的数据库(默认 ./erp_data.db);原库自动备份为 .bak。
--table 可只刷部分;缺省数据源 = 默认环境(activeEnv,tdict env 查看/切换)的库,
可用 --conn <环境名> 指定。`,
	Example: `  tdict db sync
  tdict db sync --table dzea_t,dzeal_t
  tdict db sync --conn 恒烁正式区`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, envName, err := resolveDbConn(dbSyncConn)
		if err != nil {
			return err
		}
		// 客户端直连凭据 = 账号列表首项(SelectAllSQL 用它限定 user.table schema)
		if err := conn.FillDialCred(); err != nil {
			return err
		}

		tables := dictTables
		if dbSyncTable != "" {
			tables = splitNames(dbSyncTable)
			if len(tables) == 0 {
				return fmt.Errorf("--table 未指定有效表名")
			}
		}

		// 目标 SQLite 路径: TDICT_DB 环境变量 > -d 标志 > 当前工作目录
		target := dbPath
		if env := os.Getenv("TDICT_DB"); env != "" {
			target = env
		}
		if !filepath.IsAbs(target) {
			if cwd, err := os.Getwd(); err == nil {
				target = filepath.Join(cwd, target)
			}
		}

		ctx := context.Background()
		ext, err := erpdb.Open(ctx, *conn)
		if err != nil {
			return err
		}
		defer ext.Close()

		// 写入临时库，全部成功后原子替换目标文件，避免中途失败留下半成品
		tmpPath := target + ".sync.tmp"
		os.Remove(tmpPath)
		sq, err := db.Open(tmpPath)
		if err != nil {
			return fmt.Errorf("创建临时数据库失败 (%s): %w", tmpPath, err)
		}
		if err := sq.SetBulkWriteMode(); err != nil {
			sq.Close()
			os.Remove(tmpPath)
			return fmt.Errorf("配置临时数据库失败: %w", err)
		}
		fail := func(err error) error {
			sq.Close()
			os.Remove(tmpPath)
			return err
		}

		fmt.Printf("正在从 ERP 拉取字典数据 (环境: %s) -> %s\n", envName, target)
		var results []syncResult
		var dzedCols []string
		var dzedRows [][]string
		for _, t := range tables {
			start := time.Now()
			sql, err := erpdb.SelectAllSQL(*conn, t)
			if err != nil {
				return fail(err)
			}
			cols, rows, err := ext.Query(ctx, sql)
			if err != nil {
				return fail(fmt.Errorf("查询表 %s 失败: %w", t, err))
			}
			n, err := sq.RebuildTable(t, cols, rows)
			if err != nil {
				return fail(fmt.Errorf("写入表 %s 失败: %w", t, err))
			}
			if t == "dzed_t" {
				dzedCols, dzedRows = cols, rows
			}
			elapsed := time.Since(start)
			results = append(results, syncResult{t, n, elapsed})
			fmt.Printf("  %-10s %8d 行 (%v)\n", t, n, elapsed.Round(time.Millisecond))
		}

		// 依据 dzed_t 的 PK 定义 (dzed003='P') 为本次同步的表创建唯一索引
		synced := make(map[string]bool, len(tables))
		for _, t := range tables {
			synced[t] = true
		}
		for _, w := range createPKIndexes(sq, dzedCols, dzedRows, synced) {
			fmt.Println(w)
		}

		totalRows := 0
		for _, r := range results {
			totalRows += r.rows
		}

		sq.Close()

		// 原子替换: 原库备份为 .bak, 临时库替换为正式库
		if _, err := os.Stat(target); err == nil {
			bak := target + ".bak"
			os.Remove(bak)
			if err := os.Rename(target, bak); err != nil {
				os.Remove(tmpPath)
				return fmt.Errorf("备份原数据库失败 (%s): %w", bak, err)
			}
		}
		if err := os.Rename(tmpPath, target); err != nil {
			os.Remove(tmpPath)
			return fmt.Errorf("替换数据库失败 (%s): %w", target, err)
		}
		// 清理可能残留的 WAL/SHM 孤儿文件 (旧库为 WAL 模式时)
		os.Remove(target + "-wal")
		os.Remove(target + "-shm")

		fmt.Printf("同步完成: %d 张表, 共 %d 行\n", len(results), totalRows)
		fmt.Printf("数据库已更新: %s\n", target)
		if _, err := os.Stat(target + ".bak"); err == nil {
			fmt.Printf("原数据库已备份到: %s.bak\n", target)
		}
		return nil
	},
}

// createPKIndexes builds UNIQUE indexes for each synced table's primary key
// defined in dzed_t (dzed003 = 'P'). Returns non-fatal warning strings on failure.
func createPKIndexes(sq *db.DB, dzedCols []string, dzedRows [][]string, synced map[string]bool) []string {
	ti, ki, ty, fi := indexOf(dzedCols, "dzed001"), indexOf(dzedCols, "dzed002"),
		indexOf(dzedCols, "dzed003"), indexOf(dzedCols, "dzed004")
	if ti < 0 || ki < 0 || ty < 0 || fi < 0 || len(dzedRows) == 0 {
		return nil
	}
	var warnings []string
	for _, r := range dzedRows {
		if r[ty] != "P" {
			continue
		}
		if !synced[r[ti]] {
			continue // 该表本次未同步
		}
		table, key := r[ti], r[ki]
		var cols []string
		for _, f := range strings.Split(r[fi], ",") {
			f = strings.Trim(strings.TrimSpace(f), `"`)
			if f != "" {
				cols = append(cols, f)
			}
		}
		if len(cols) == 0 {
			continue
		}
		idxName := fmt.Sprintf("idx_%s_%s", table, key)
		if err := sq.CreateUniqueIndex(idxName, table, cols); err != nil {
			warnings = append(warnings, fmt.Sprintf("  [警告] 索引 %s 创建失败: %v", idxName, err))
		}
	}
	return warnings
}

func indexOf(cols []string, target string) int {
	for i, c := range cols {
		if c == target {
			return i
		}
	}
	return -1
}

func init() {
	dbSyncCmd.Flags().StringVarP(&dbSyncConn, "conn", "c", "", "连接名称 (默认使用 isDefault 连接)")
	dbSyncCmd.Flags().StringVar(&dbSyncTable, "table", "", "仅同步指定的表 (逗号分隔, 默认全部字典表)")
	dbCmd.AddCommand(dbSyncCmd)
}
