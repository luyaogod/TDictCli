package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"tdict/dbsync"

	"github.com/spf13/cobra"
)

var (
	dbSyncConn  string
	dbSyncTable string
)

var dbSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "从 ERP 刷新本地查询数据",
	Long: `从 ERP 刷新本地查询数据(表字典、校验、分类码、画面规格、开窗、消息、参数等
全部内容,约 85 万行),查询命令读的就是它。
写入 -d/--db 或 TDICT_DB 指向的数据库(默认 ./erp_data.db);原库自动备份为 .bak。
--table 可只刷部分;缺省数据源 = 默认环境(activeEnv,tdict env 查看/切换)的库,
可用 --conn <环境名> 指定。
也可在 tdict serve 的「数据同步」页面里选环境执行(带进度)。`,
	Example: `  tdict db sync
  tdict db sync --table dzea_t,dzeal_t
  tdict db sync --conn 正式区`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, envName, err := resolveDbConn(dbSyncConn)
		if err != nil {
			return err
		}
		tables := dbsync.DictTables
		if dbSyncTable != "" {
			tables = splitNames(dbSyncTable)
			if len(tables) == 0 {
				return fmt.Errorf("--table 未指定有效表名")
			}
		}

		// 目标 SQLite:优先已存在的库(与查询命令读的是同一个),不存在再按 TDICT_DB/exe 目录/当前目录创建
		target := resolveSyncTarget()

		fmt.Printf("正在从 ERP 拉取字典数据 (环境: %s) -> %s\n", envName, target)
		st, err := dbsync.Run(context.Background(), *conn, target, tables, func(p dbsync.Progress) {
			// 每张表完成时打印一行(开始回调 TableDur=0,跳过)
			if p.Phase == "table" && p.TableDur > 0 {
				fmt.Printf("  %-10s %8d 行 (%v)\n", p.Table, p.TableRows, p.TableDur.Round(time.Millisecond))
			}
		})
		if err != nil {
			return err
		}
		for _, w := range st.Warnings {
			fmt.Println(w)
		}
		fmt.Printf("同步完成: %d 张表, 共 %d 行\n", st.Tables, st.Rows)
		fmt.Printf("数据库已更新: %s\n", st.Target)
		if st.Backup != "" {
			fmt.Printf("原数据库已备份到: %s\n", st.Backup)
		}
		return nil
	},
}

// resolveSyncTarget 解析数据同步(写库)的目标路径,与查询命令的查找规则保持一致:
// 优先返回已存在的库(TDICT_DB > -d 绝对 > exe 同目录 > 当前目录),使写入的正是查询要读的那个;
// 都不存在时按 TDICT_DB > exe 同目录 > 当前目录 决定新建位置(便携版首跑:exe 同目录)。
func resolveSyncTarget() string {
	if p, err := resolveDBPath(dbPath); err == nil {
		return p
	}
	if env := os.Getenv("TDICT_DB"); env != "" {
		if abs, err := filepath.Abs(env); err == nil {
			return abs
		}
		return env
	}
	if exe, err := os.Executable(); err == nil {
		return filepath.Join(filepath.Dir(exe), dbPath)
	}
	if abs, err := filepath.Abs(dbPath); err == nil {
		return abs
	}
	return dbPath
}

func init() {
	dbSyncCmd.Flags().StringVarP(&dbSyncConn, "conn", "c", "", "连接名称 (默认使用 isDefault 连接)")
	dbSyncCmd.Flags().StringVar(&dbSyncTable, "table", "", "仅同步指定的表 (逗号分隔, 默认全部字典表)")
	dbCmd.AddCommand(dbSyncCmd)
}
