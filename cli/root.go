package cli

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"tdict/db"

	"github.com/spf13/cobra"
)

var (
	dbPath     string
	configPath string
	useJSON    bool
	useCSV     bool
	verbose    bool
	// webFS holds the embedded frontend build output (web/dist), provided by main via Execute.
	webFS fs.FS
)

// rootCmd is the base command.
var rootCmd = &cobra.Command{
	Use:   "tdict",
	Short: "TDict - ERP data dictionary query tool",
	Long: `TDict 是一个查询 ERP 数据字典的 CLI 工具。
查询数据源(r.t/r.v/desc/scc/r.q)默认是本地 SQLite 镜像 (erp_data.db,由 tdict db sync 同步);
可用 config.json 顶层 query.source 或 --conn <环境名> 切换为某环境的远程库直查(金仓/Oracle)。
所有输出使用简体中文 (zh_CN)。`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 打开查询数据源:本地 SQLite(--db/TDICT_DB)或远程库(--conn/query.source)
		return openQuerySource()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		closeQuerySource()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db", "d", "erp_data.db", "Path to SQLite database file")
	rootCmd.PersistentFlags().BoolVar(&useJSON, "json", false, "Output in JSON format")
	rootCmd.PersistentFlags().BoolVar(&useCSV, "csv", false, "Output in CSV format")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output (show resolved db path)")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "ERP 数据库连接配置文件路径 (JSON;缺省取统一用户目录 "+defaultConfigPathHint()+")")
}

// resolveDBPath resolves the database file path with the following priority:
//  1. TDICT_DB environment variable
//  2. -d flag as absolute path
//  3. -d flag relative to executable directory
//  4. -d flag relative to current working directory
//
// Returns an error if no existing file can be found at any of these locations.
func resolveDBPath(flagPath string) (string, error) {
	var candidates []string

	// 1. TDICT_DB environment variable (highest priority)
	if env := os.Getenv("TDICT_DB"); env != "" {
		candidates = append(candidates, env)
	}

	// 2. -d flag as-is
	candidates = append(candidates, flagPath)

	// 3. -d flag relative to executable directory
	if !filepath.IsAbs(flagPath) {
		if execPath, err := os.Executable(); err == nil {
			candidates = append(candidates, filepath.Join(filepath.Dir(execPath), flagPath))
		}
	}

	// 4. -d flag relative to CWD
	if !filepath.IsAbs(flagPath) {
		if cwd, err := os.Getwd(); err == nil {
			candidates = append(candidates, filepath.Join(cwd, flagPath))
		}
	}

	// Try each candidate
	var tried []string
	for _, p := range candidates {
		abs, _ := filepath.Abs(p)
		if _, err := os.Stat(abs); err == nil {
			return abs, nil
		}
		tried = append(tried, abs)
	}

	// None found - give a helpful error
	return "", fmt.Errorf(
		"database file not found.\n\nTried these paths:\n%s\n\nSet the TDICT_DB environment variable or use -d to specify the correct path:\n  setx TDICT_DB \"D:\\path\\to\\erp_data.db\"\n  tdict -d \"D:\\path\\to\\erp_data.db\" table dzea_t",
		formatTriedPaths(tried),
	)
}

func formatTriedPaths(paths []string) string {
	var s string
	for _, p := range paths {
		exists := ""
		if _, err := os.Stat(p); err == nil {
			exists = " (found)"
		}
		s += fmt.Sprintf("  - %s%s\n", p, exists)
	}
	return s
}

// ---------- 配置文件位置 ----------
//
// 存放规则(与 TDebug 保持一致,改动请两边同步):
//  1. TDICT_CONFIG 环境变量                           显式指定
//  2. --config <路径>                                 显式指定
//  3. <exe 目录>\.portable 存在                       便携包:配置留在包内
//  4. %APPDATA%\T100\tdict\config.json                默认:固定用户的统一位置
//  5. <exe 目录>\config.json、<当前目录>\config.json   旧位置兜底(首次运行自动迁移到 4)
//
// 统一位置可用 T100_HOME 环境变量整体改写(如 T100_HOME=D:\t100)。
// 只挪配置文件,不挪数据:mirror.dir / bdldoc.dir / sync.target / erp_data.db
// 仍是配置里的绝对路径,搬到哪由用户自己决定。

const (
	// toolDirName 统一用户目录下本工具的子目录。
	toolDirName = "tdict"
	// defaultConfigName 缺省配置文件名。
	defaultConfigName = "config.json"
	// portableMark 便携标记:exe 同目录存在该文件即视为便携包。
	portableMark = ".portable"
)

// toolsHome 固定用户的统一工具目录:T100_HOME 优先,否则 %APPDATA%\T100
// (os.UserConfigDir() 在 Windows 上即 %AppData%)。
func toolsHome() string {
	if env := os.Getenv("T100_HOME"); env != "" {
		if abs, err := filepath.Abs(env); err == nil {
			return abs
		}
		return env
	}
	if dir, err := os.UserConfigDir(); err == nil {
		return filepath.Join(dir, "T100")
	}
	return ""
}

// userConfigDir 统一用户目录下的本工具目录;定位不到时返回空串。
func userConfigDir() string {
	if home := toolsHome(); home != "" {
		return filepath.Join(home, toolDirName)
	}
	return ""
}

// userConfigPath 统一用户目录下的配置路径;定位不到时返回空串。
func userConfigPath() string {
	if dir := userConfigDir(); dir != "" {
		return filepath.Join(dir, defaultConfigName)
	}
	return ""
}

// isPortable 是否便携包(决定配置留在包内还是进统一用户目录)。
func isPortable() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	_, err = os.Stat(filepath.Join(filepath.Dir(exe), portableMark))
	return err == nil
}

// legacyConfigPaths 旧位置,按优先级排列。
func legacyConfigPaths() []string {
	var out []string
	if exe, err := os.Executable(); err == nil {
		out = append(out, filepath.Join(filepath.Dir(exe), defaultConfigName))
	}
	if cwd, err := os.Getwd(); err == nil {
		out = append(out, filepath.Join(cwd, defaultConfigName))
	}
	return out
}

// looksLikeOwnConfig 判断一段 JSON 是否像本工具的配置 —— 只认自己那几节的顶层键,
// 避免把别的项目(current 目录下恰好存在的)config.json 误迁移过来。
func looksLikeOwnConfig(b []byte) bool {
	var root map[string]json.RawMessage
	if json.Unmarshal(b, &root) != nil {
		return false
	}
	for _, k := range []string{"hosts", "query", "mirror", "bdldoc", "sync", "debug"} {
		if _, ok := root[k]; ok {
			return true
		}
	}
	return false
}

// migrateLegacyConfigTo 首次运行迁移:把旧位置的配置复制到 dst(仅在 dst 尚不存在时)。
// 只复制、不删除源文件 —— 由用户自己确认后再清理。返回迁移来源;未迁移返回空串。
func migrateLegacyConfigTo(dst string) string {
	if dst == "" {
		return ""
	}
	if _, err := os.Stat(dst); err == nil {
		return "" // 已有配置,不动
	}
	for _, src := range legacyConfigPaths() {
		if samePath(src, dst) {
			continue
		}
		b, err := os.ReadFile(src)
		if err != nil || !looksLikeOwnConfig(b) {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return ""
		}
		if err := os.WriteFile(dst, b, 0o600); err != nil {
			return ""
		}
		return src
	}
	return ""
}

// migrateLegacyConfig 迁移到统一用户目录(CLI 路径用)。
func migrateLegacyConfig() string {
	if dst := userConfigPath(); dst != "" {
		return migrateLegacyConfigTo(dst)
	}
	return ""
}

// samePath 两个路径是否指向同一个文件(比较绝对路径,失败则退回字面比较)。
func samePath(a, b string) bool {
	aa, err1 := filepath.Abs(a)
	bb, err2 := filepath.Abs(b)
	if err1 != nil || err2 != nil {
		return a == b
	}
	return aa == bb
}

// defaultConfigPath 未经显式指定时的默认落点:便携包内,否则统一用户目录。
func defaultConfigPath() string {
	if isPortable() {
		if exe, err := os.Executable(); err == nil {
			return filepath.Join(filepath.Dir(exe), defaultConfigName)
		}
	}
	if p := userConfigPath(); p != "" {
		return p
	}
	if abs, err := filepath.Abs(defaultConfigName); err == nil {
		return abs
	}
	return defaultConfigName
}

// defaultConfigPathHint 给 --help 用的一行提示(取不到用户目录时退回文件名)。
func defaultConfigPathHint() string {
	if isPortable() {
		return "<exe 目录>\\" + defaultConfigName
	}
	if dir := userConfigDir(); dir != "" {
		return filepath.Join(dir, defaultConfigName)
	}
	return defaultConfigName
}

// resolveConfigPath resolves the connection config file path with the following priority:
//  1. TDICT_CONFIG environment variable
//  2. --config flag (explicit): as-is, relative to exe dir, relative to CWD
//  3. portable package (exe dir has .portable): <exe dir>\config.json
//  4. unified user dir: %APPDATA%\T100\tdict\config.json (migrates legacy on first run)
//  5. legacy locations (see legacyConfigPaths)
//
// Returns an error if no existing file can be found at any of these locations.
func resolveConfigPath(flagPath string) (string, error) {
	var candidates []string

	// 1. TDICT_CONFIG environment variable (highest priority)
	if env := os.Getenv("TDICT_CONFIG"); env != "" {
		candidates = append(candidates, env)
	}

	if flagPath != "" {
		// 2. --config 显式指定:原样 / 相对 exe 同目录 / 相对当前目录
		candidates = append(candidates, flagPath)
		if !filepath.IsAbs(flagPath) {
			if execPath, err := os.Executable(); err == nil {
				candidates = append(candidates, filepath.Join(filepath.Dir(execPath), flagPath))
			}
			if cwd, err := os.Getwd(); err == nil {
				candidates = append(candidates, filepath.Join(cwd, flagPath))
			}
		}
	} else {
		if isPortable() {
			// 3. 便携包:配置留在包内,不参与统一用户目录与迁移
			if exe, err := os.Executable(); err == nil {
				candidates = append(candidates, filepath.Join(filepath.Dir(exe), defaultConfigName))
			}
		} else if p := userConfigPath(); p != "" {
			// 4. 统一用户目录;首次运行先把旧配置迁移过来
			if src := migrateLegacyConfig(); src != "" {
				fmt.Fprintf(os.Stderr, "[tdict] 已迁移配置到统一位置: %s -> %s\n", src, p)
			}
			candidates = append(candidates, p)
		}
		// 5. 旧位置兜底(用户目录不可用或迁移失败时仍能用)
		candidates = append(candidates, legacyConfigPaths()...)
	}

	// Try each candidate
	var tried []string
	for _, p := range candidates {
		abs, _ := filepath.Abs(p)
		if _, err := os.Stat(abs); err == nil {
			return abs, nil
		}
		tried = append(tried, abs)
	}

	// None found - give a helpful error
	return "", fmt.Errorf(
		"配置文件未找到。\n\n尝试了以下路径:\n%s\n\n缺省位置: %s\n设置 TDICT_CONFIG 环境变量或使用 --config 指定正确路径:\n  setx TDICT_CONFIG \"D:\\path\\to\\config.json\"\n  tdict --config \"D:\\path\\to\\config.json\" db list",
		formatTriedPaths(tried), defaultConfigPath(),
	)
}

// Execute runs the root command. web carries the embedded frontend build output
// (web/dist); may be nil/empty. AI 技能文件不再内嵌:以 skills/ 目录随发行包分发,
// 由 `tdict install skills` 复制到当前目录。
func Execute(web fs.FS) {
	webFS = web
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// printJSON 以 JSON 输出结果(缩进两格)。
func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// GetDB returns the current query data source (local SQLite or remote ERP DB).
// It is set by root's PersistentPreRunE (see source.go).
func GetDB() db.Source {
	return dataSrc
}

// IsJSON returns true if JSON output is requested.
func IsJSON() bool { return useJSON }

// IsCSV returns true if CSV output is requested.
func IsCSV() bool { return useCSV }
