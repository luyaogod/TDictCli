package cli

import (
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
	database   *db.DB
	// skillFS holds the embedded Claude Code skill files (.claude/skills),
	// provided by main via Execute. Used by `tdict install`.
	skillFS fs.FS
)

// rootCmd is the base command.
var rootCmd = &cobra.Command{
	Use:   "tdict",
	Short: "TDict - ERP data dictionary query tool",
	Long: `TDict 是一个查询 ERP 数据字典的 CLI 工具。
查询本地 SQLite (erp_data.db) 中的数据字典，并通过 tdict db sync 从 ERP 实时刷新。
所有输出使用简体中文 (zh_CN)。`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if useOnline {
			// 在线直查:打开远程 ERP 连接(而非本地 SQLite)
			return openOnline(cmd.Name())
		}
		resolvedPath, err := resolveDBPath(dbPath)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "[tdict] database: %s\n", resolvedPath)
		}

		database, err = db.Open(resolvedPath)
		if err != nil {
			return fmt.Errorf("failed to open database at %s: %w", resolvedPath, err)
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if useOnline {
			closeOnline()
			return
		}
		if database != nil {
			database.Close()
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db", "d", "erp_data.db", "Path to SQLite database file")
	rootCmd.PersistentFlags().BoolVar(&useJSON, "json", false, "Output in JSON format")
	rootCmd.PersistentFlags().BoolVar(&useCSV, "csv", false, "Output in CSV format")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output (show resolved db path)")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config.json", "ERP 数据库连接配置文件路径 (JSON)")
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

// resolveConfigPath resolves the connection config file path with the following priority:
//  1. TDICT_CONFIG environment variable
//  2. --config flag as absolute path
//  3. --config flag relative to executable directory
//  4. --config flag relative to current working directory
//
// Returns an error if no existing file can be found at any of these locations.
func resolveConfigPath(flagPath string) (string, error) {
	var candidates []string

	// 1. TDICT_CONFIG environment variable (highest priority)
	if env := os.Getenv("TDICT_CONFIG"); env != "" {
		candidates = append(candidates, env)
	}

	// 2. --config flag as-is
	candidates = append(candidates, flagPath)

	// 3. --config flag relative to executable directory
	if !filepath.IsAbs(flagPath) {
		if execPath, err := os.Executable(); err == nil {
			candidates = append(candidates, filepath.Join(filepath.Dir(execPath), flagPath))
		}
	}

	// 4. --config flag relative to CWD
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
		"配置文件未找到。\n\n尝试了以下路径:\n%s\n\n设置 TDICT_CONFIG 环境变量或使用 --config 指定正确路径:\n  setx TDICT_CONFIG \"D:\\path\\to\\config.json\"\n  tdict --config \"D:\\path\\to\\config.json\" db list",
		formatTriedPaths(tried),
	)
}

// Execute runs the root command. skills carries the embedded Claude Code
// skill files (used by `tdict install`); it may be nil when unavailable.
// web carries the embedded debug web frontend (web/dist); may be nil/empty.
func Execute(skills fs.FS, web fs.FS) {
	skillFS = skills
	webFS = web
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// GetDB returns the current database connection.
func GetDB() *db.DB {
	return database
}

// IsJSON returns true if JSON output is requested.
func IsJSON() bool { return useJSON }

// IsCSV returns true if CSV output is requested.
func IsCSV() bool { return useCSV }
