package cli

// tdict install skills:把与 tdict 可执行文件同目录的 skills/ 目录复制到当前工作目录。
// skills/ 是普通可编辑的 markdown 文件(不再二进制内嵌),用户按自己使用的 AI 工具
// 改名/移动到对应位置即可(如 Claude Code 的 .claude/skills/)。

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// skillsDirName 技能目录名:可执行文件同目录下的 skills/(复制到当前目录时同名)。
const skillsDirName = "skills"

// validateInstallArg 只接受可选的 "skills" 参数(最多一个)。
func validateInstallArg(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("用法: tdict install skills")
	}
	if len(args) == 1 && !strings.EqualFold(strings.TrimSpace(args[0]), skillsDirName) {
		return fmt.Errorf("只支持安装 %s(用法: tdict install skills)", skillsDirName)
	}
	return nil
}

// installSkills 把 <exe目录>/skills 复制到 <当前目录>/skills。
// 源码开发场景(如 go run,exe 在同目录没有 skills/)回退到当前目录已有的 skills/。
// 源与目标为同一目录时不做复制,返回 0。
func installSkills() (string, int, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", 0, fmt.Errorf("定位可执行文件失败: %w", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", 0, fmt.Errorf("获取当前目录失败: %w", err)
	}
	dst := filepath.Join(cwd, skillsDirName)
	src := filepath.Join(filepath.Dir(exe), skillsDirName)
	if !isDir(src) {
		if !isDir(dst) {
			where := src
			if !sameDir(src, dst) {
				where = src + " 与 " + dst
			}
			return "", 0, fmt.Errorf("未找到 %s 目录(%s):请把它与 tdict 可执行文件放在同一目录",
				skillsDirName, where)
		}
		src = dst // 回退:当前目录已有 skills/
	}
	if sameDir(src, dst) {
		return dst, 0, nil
	}
	n, err := copyTree(src, dst)
	if err != nil {
		return "", 0, err
	}
	return dst, n, nil
}

func isDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}

// sameDir 判断两个路径是否指同一目录(绝对化 + 清理;Windows 忽略大小写)。
func sameDir(a, b string) bool {
	pa, err1 := filepath.Abs(a)
	pb, err2 := filepath.Abs(b)
	if err1 != nil || err2 != nil {
		return false
	}
	return strings.EqualFold(filepath.Clean(pa), filepath.Clean(pb))
}

// copyTree 递归复制 src 目录内容到 dst,返回复制的文件数。
func copyTree(src, dst string) (int, error) {
	n := 0
	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.Create(target)
		if err != nil {
			return err
		}
		if _, err := io.Copy(out, in); err != nil {
			out.Close()
			return err
		}
		if err := out.Close(); err != nil {
			return err
		}
		n++
		return nil
	})
	return n, err
}

var installCmd = &cobra.Command{
	Use:   "install [skills]",
	Short: "把 skills/ 目录复制到当前目录",
	Long: `把与 tdict 可执行文件同目录的 skills/ 目录复制到当前工作目录(用法: tdict install skills)。

skills/ 是普通可编辑的 markdown 文件(不再内嵌二进制);安装后由你按所用 AI 工具
改名/移动到对应位置,例如 Claude Code 的 .claude/skills/。`,
	Example: `  tdict install skills     # 把 skills/ 复制到当前目录
  tdict install            # 同上(默认)`,
	Args: cobra.MaximumNArgs(1),
	// 覆盖 root 的 SQLite 钩子:本命令只复制文件,不依赖本地 erp_data.db
	// (否则在没有字典库的目录/便携版首跑时会报"数据库文件未找到")。
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateInstallArg(args); err != nil {
			return err
		}
		dst, n, err := installSkills()
		if err != nil {
			return err
		}
		if n == 0 {
			fmt.Printf("%s 已在当前目录: %s\n", skillsDirName, dst)
			return nil
		}
		fmt.Printf("技能安装完成: %d 个文件 -> %s\n", n, dst)
		fmt.Printf("提示: %s/ 是普通文件,按所用 AI 工具改名/移动即可(如 .claude/skills)\n", skillsDirName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
