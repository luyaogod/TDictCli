package cli

import (
	"fmt"
	"os"

	"tdict/dbconfig"
	"tdict/host"

	"github.com/spf13/cobra"
)

var (
	dbCfg *host.Hosts
)

// dbCmd 管理 ERP 数据库连接(每个 SSH 环境一对一挂载的 db)与同步/在线查询。
//
// It defines its own PersistentPreRunE, which overrides root's SQLite-open
// hook (cobra runs the first persistent pre-run found walking leaf→root),
// so `tdict db *` does not depend on the local erp_data.db.
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "管理 ERP 数据库连接与数据同步",
	Long: `从 config.json 读取配置 (--config / TDICT_CONFIG)。数据库连接按环境一对一挂载:
debug.sshs[].db 即该环境的库(显式 host/port/service|库名 + 账号列表)。
子命令: sync (从 ERP 拉取字典数据写入 SQLite) / list (列各环境的库) /
ping (验证连接可达) / discover (SSH 自动发现连接要素并写入环境 db)。
支持连接类型: kingbase (金仓, PostgreSQL 协议)、oracle (go-ora)。`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		path, err := resolveConfigPath(configPath)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "[tdict] 配置文件: %s\n", path)
		}
		cfg, err := host.LoadHosts(path)
		if err != nil {
			return err
		}
		dbCfg = cfg
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}

// resolveDbConn 按环境名解析其 db(--conn <环境名>);name 为空取活跃环境(activeEnv,
// 未设置自动取首条)。返回该连接的深拷贝与所属环境名。
func resolveDbConn(name string) (*dbconfig.Connection, string, error) {
	if dbCfg == nil {
		return nil, "", fmt.Errorf("配置未加载")
	}
	target := name
	if target == "" {
		target = dbCfg.ActiveEnv
	}
	if target == "" && len(dbCfg.SSHs) > 0 {
		target = dbCfg.SSHs[0].Name
	}
	if target == "" {
		return nil, "", fmt.Errorf("debug.sshs 未配置服务器环境")
	}
	for i := range dbCfg.SSHs {
		e := &dbCfg.SSHs[i]
		if e.Name != target {
			continue
		}
		if e.DB == nil {
			return nil, target, fmt.Errorf("环境 %q 未配置数据库(设置-环境-数据库页添加)", target)
		}
		cc := *e.DB
		cc.Accounts = append([]dbconfig.DBAcct(nil), e.DB.Accounts...)
		return &cc, target, nil
	}
	return nil, name, fmt.Errorf("未找到环境 %q(可用: tdict debug env 查看)", name)
}
