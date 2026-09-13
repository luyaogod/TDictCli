package cli

// tdict serve:本地可视化配置服务(SSH 环境与数据库连接)。
// 只做配置读写与只读探测,不含调试能力;Ctrl+C 停止。

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"tdict/server"

	"github.com/spf13/cobra"
)

var serveListen string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动可视化配置页面(维护 SSH 连接与数据库)",
	Long: `启动本地配置服务:浏览器打开后可视化维护 config.json 的 SSH 环境与数据库连接。

页面支持:
  - 环境列表增删,编辑 SSH 主机/端口/账号/密码/登录区域/TOPENT;
  - 每个环境一对一挂载的数据库(oracle 服务名 / 金仓库名 + 账号列表);
  - 「从服务器获取数据库配置」:登录该环境 SSH 只读探测连接要素并回填;
  - 账号逐行服务器侧验证、数据库客户端直连测试。

配置文件取 --config / TDICT_CONFIG(缺省 config.json);文件不存在时在页面首次保存时创建。
监听地址缺省 127.0.0.1:28670,被占用自动顺延;Ctrl+C 停止。`,
	Example: `  tdict serve
  tdict serve --listen 127.0.0.1:9123`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil // 覆盖 root 的本地 SQLite 打开逻辑:配置服务不依赖字典库
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// 配置文件不存在也允许启动(首次在页面里创建)
		cfgPath, err := resolveConfigPath(configPath)
		if err != nil {
			abs, aerr := filepath.Abs(configPath)
			if aerr != nil {
				return err
			}
			cfgPath = abs
		}
		srv := server.New(cfgPath, serveListen, webFS)
		// 数据同步目标库:与查询命令同一解析规则(优先已存在的库;否则 TDICT_DB/exe 同目录/当前目录)
		dbTarget := resolveSyncTarget()
		srv.SetDBTarget(dbTarget)
		ln, addr, err := srv.Listen()
		if err != nil {
			return err
		}
		fmt.Printf("[tdict] 配置服务已启动: http://%s\n", addr)
		fmt.Printf("  配置文件: %s\n", cfgPath)
		fmt.Printf("  数据同步目标: %s\n", dbTarget)
		fmt.Println("  按 Ctrl+C 停止")
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()
		return srv.Serve(ctx, ln)
	},
}

func init() {
	serveCmd.Flags().StringVar(&serveListen, "listen", "", "监听地址(默认 127.0.0.1:28670;被占用自动顺延)")
	rootCmd.AddCommand(serveCmd)
}
