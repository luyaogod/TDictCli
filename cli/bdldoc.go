package cli

// tdict bdldoc:BDL(4GL)编程语言参考文档目录的查看/设置。
// 文档(markdown,无图片)随仓库放在 docs/bdl 下;实际路径记录在 config.json
// 顶层 "bdldoc": {"dir": ...}——AI/工具按此路径定位语言文档。
// 本命令只查看/设置路径(改 config.json),不移动任何文件;与 mirror.dir 同风格。

import (
	"fmt"
	"path/filepath"

	"tdict/cfgfile"

	"github.com/spf13/cobra"
)

var bdldocCmd = &cobra.Command{
	Use:   "bdldoc",
	Short: "BDL(4GL)语言文档目录(config.json 顶层 bdldoc.dir 查看/设置)",
	Long: `管理 BDL(4GL)编程语言参考文档(Genero BDL markdown)的存放目录,
供 AI/工具读取语言文档时定位。文档本体随仓库在 docs/bdl 下;本命令只
查看或设置 config.json 顶层 bdldoc.dir,不移动任何文件。

子命令:
  dir [<目录>]  不带参数显示当前目录;带目录则写入 config.json(转绝对路径)

未设置时 AI 应提示先执行: tdict bdldoc dir <目录>`,
	Example: `  tdict bdldoc dir                    # 显示当前文档目录
  tdict bdldoc dir D:\dev\bdl-docs     # 设置文档目录(写入 config.json)`,
	// 覆盖 root 的 SQLite 钩子:本命令只操作 config.json
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
}

// bdldocDir 读 config.json 顶层 bdldoc.dir;返回绝对路径;ok=false=未配置。
func bdldocDir() (string, bool) {
	path, err := resolveConfigPath(configPath)
	if err != nil {
		return "", false
	}
	root, err := cfgfile.Open(path)
	if err != nil {
		return "", false
	}
	sec, _ := root["bdldoc"].(map[string]any)
	dir, _ := sec["dir"].(string)
	if dir == "" {
		return "", false
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return dir, true
	}
	return abs, true
}

// setBdldocDir 写回 config.json 顶层 bdldoc.dir(原子;其余顶层键原样保留)。
func setBdldocDir(dir string) error {
	path, err := resolveConfigPath(configPath)
	if err != nil {
		return err
	}
	return cfgfile.Edit(path, nil, func(root map[string]any) error {
		root["bdldoc"] = map[string]any{"dir": dir}
		return nil
	})
}

var bdldocDirCmd = &cobra.Command{
	Use:   "dir [<目录>]",
	Short: "查看或设置 BDL 文档目录(config.json 顶层 bdldoc.dir)",
	Long: `不带参数显示当前 BDL 文档目录(第一行即路径,便于 AI/脚本取用);
带参数把目录写入 config.json(相对路径转绝对)。仅改设置,不移动文档文件。`,
	Example: `  tdict bdldoc dir                      # 显示当前目录
  tdict bdldoc dir D:\dev\bdl-docs       # 设置目录(写入 config.json)`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			abs, err := filepath.Abs(args[0])
			if err != nil {
				return fmt.Errorf("解析目录路径失败: %w", err)
			}
			if err := setBdldocDir(abs); err != nil {
				return err
			}
			fmt.Printf("BDL 文档目录已设置: %s\n", abs)
			return nil
		}
		dir, ok := bdldocDir()
		if !ok {
			fmt.Println("BDL 文档目录未设置。请执行: tdict bdldoc dir <目录>")
			return nil
		}
		fmt.Println(dir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bdldocCmd)
	bdldocCmd.AddCommand(bdldocDirCmd)
}
