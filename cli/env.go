package cli

// tdict env:离线查看/设置默认环境(直接读写 config.json 的 debug.activeEnv,
// 不需要 debug serve 运行)——与前端设置页「设为默认」等效。
// 与 tdict debug env(需 serve 运行,且会把正在进行的调试会话一并切换过去)
// 互补:本命令只动配置文件,不影响运行中的服务/会话内存态
// (serve 下次保存设置时以其内存态为准)。

import (
	"fmt"

	"tdict/host"
	"tdict/output"
	"tdict/cfgfile"

	"github.com/spf13/cobra"
)

var (
	envCfg     *host.Hosts
	envCfgPath string
)

var envCmd = &cobra.Command{
	Use:   "env [<环境名>]",
	Short: "查看/设置默认环境",
	Long: `列出 config.json 中全部 SSH 环境与当前默认(activeEnv),或直接把默认环境写入配置。
无参数 = 列出全部环境(带 * 为当前默认);带环境名 = 设为默认并持久化(校验名称存在)。
本命令离线执行(同 mirror/db 组,只读改 config.json),不依赖 debug serve;
要连同正在运行的调试会话一起切换,请用 tdict debug env <环境名>(需 serve)。`,
	Example: `  tdict env                     # 列出全部环境与当前默认
  tdict env 恒烁正式区            # 把默认环境设为恒烁正式区(写入 config.json)
  tdict env --json`,
	Args: cobra.MaximumNArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 覆盖 root 的 SQLite 钩子:本命令只操作 config.json
		path, err := resolveConfigPath(configPath)
		if err != nil {
			return err
		}
		cfg, err := host.LoadHosts(path)
		if err != nil {
			return err
		}
		envCfgPath = path
		envCfg = cfg
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return envSetDefault(args[0])
		}
		return envList()
	},
}

// envListItem JSON 列表项(只含非敏感字段)。
type envListItem struct {
	Name    string `json:"名称"`
	Host    string `json:"主机"`
	Port    int    `json:"端口"`
	Zone    string `json:"区域,omitempty"`
	Topent  string `json:"topent,omitempty"`
	DBType  string `json:"dbType,omitempty"`
	DBAddr  string `json:"db地址,omitempty"`
	Default bool   `json:"默认"`
}

// envList 列出全部环境与当前默认。
func envList() error {
	items := make([]envListItem, 0, len(envCfg.SSHs))
	for i := range envCfg.SSHs {
		e := &envCfg.SSHs[i]
		it := envListItem{
			Name:   e.Name,
			Host:   e.Host,
			Port:   e.Port,
			Zone:   e.Zone,
			Topent: string(e.Topent),
			Default: e.Name == envCfg.ActiveEnv,
		}
		if e.DB != nil {
			it.DBType = e.DB.Type
			it.DBAddr = e.DB.Address()
		}
		items = append(items, it)
	}
	if IsJSON() {
		out := struct {
			ActiveEnv string         `json:"activeEnv"`
			Envs      []envListItem  `json:"环境"`
		}{envCfg.ActiveEnv, items}
		return output.PrintJSON(out)
	}

	fmt.Printf("当前默认(activeEnv): %s\n", orDefault(envCfg.ActiveEnv, "未设置"))
	fmt.Printf("环境 (%d):\n", len(items))
	for _, it := range items {
		mark := "  "
		if it.Default {
			mark = " *"
		}
		port := it.Port
		if port == 0 {
			port = 22
		}
		line := fmt.Sprintf("  %s %-12s %-20s zone=%-4s TOPENT=%s", mark, it.Name,
			fmt.Sprintf("%s:%d", it.Host, port), orDefault(it.Zone, "-"), orDefault(it.Topent, "-"))
		if it.DBType != "" {
			line += fmt.Sprintf("  db=%s %s", it.DBType, it.DBAddr)
		}
		fmt.Println(line)
	}
	return nil
}

// envSetDefault 校验环境名并把默认环境写入 config.json(原子;其余键原样保留)。
func envSetDefault(name string) error {
	hit := false
	for i := range envCfg.SSHs {
		if envCfg.SSHs[i].Name == name {
			hit = true
			break
		}
	}
	if !hit {
		return fmt.Errorf("未找到环境 %q(可用: tdict env 列出全部环境)", name)
	}
	if name == envCfg.ActiveEnv {
		fmt.Printf("当前默认已是 %q。\n", name)
		return nil
	}

	if err := cfgfile.Edit(envCfgPath, nil, func(root map[string]any) error {
		dbg, err := cfgfile.Debug(root)
		if err != nil {
			return err
		}
		dbg["activeEnv"] = name
		return nil
	}); err != nil {
		return err
	}
	fmt.Printf("默认环境已切换为: %s(已写入 %s)\n", name, envCfgPath)
	fmt.Printf("提示: 运行中的 debug serve 不受影响;新调试会话/缺省 db sync、mirror 将使用该环境。\n")
	return nil
}


func init() {
	rootCmd.AddCommand(envCmd)
}
