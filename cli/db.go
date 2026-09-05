package cli

import (
	"fmt"
	"os"

	"tdict/dbconfig"

	"github.com/spf13/cobra"
)

var (
	dbConfig *dbconfig.Config
)

// dbCmd manages the ERP database connection used by `tdict db sync`.
//
// It defines its own PersistentPreRunE, which overrides root's SQLite-open
// hook (cobra runs the first persistent pre-run found walking leaf→root),
// so `tdict db *` does not depend on the local erp_data.db.
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "管理 ERP 数据库连接与同步/在线查询",
	Long: `从 config.json 读取 ERP 数据库连接配置 (--config / TDICT_CONFIG)。
子命令: sync (从 ERP 拉取字典数据写入 SQLite) / list (列连接) /
ping (验证连接可达) / discover (SSH 自动发现连接要素并保存)。
支持连接类型: kingbase (金仓, PostgreSQL 协议)、oracle (go-ora)。`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		path, err := resolveConfigPath(configPath)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "[tdict] 配置文件: %s\n", path)
		}
		cfg, err := dbconfig.Load(path)
		if err != nil {
			return err
		}
		dbConfig = cfg
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}

// resolveConnection returns the named connection, or the default one when name is empty.
func resolveConnection(name string) (*dbconfig.Connection, error) {
	if dbConfig == nil {
		return nil, fmt.Errorf("连接配置未加载")
	}
	if name != "" {
		return dbConfig.FindByName(name)
	}
	return dbConfig.FindDefault()
}
