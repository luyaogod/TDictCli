package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"tdict/debug"

	"github.com/spf13/cobra"
)

// webFS 由 main.go 通过 //go:embed web/dist 注入;未构建前端时为 nil
var webFS fs.FS

var (
	dbgModule string
	dbgProg   string
	dbgLine   int
	dbgListen string
	dbEnt     int
	dbJSON    bool
)

// debugCmd 覆盖 PersistentPreRunE:debug 功能不需要本地 SQLite
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "AI 人机协同调试 T100 作业(fgldb 协议驱动)",
	Long: `通过 SSH 在 T100 服务器上驱动 fglrun -d 的 (fgldb) 文本调试协议,
提供本地 Web 调试界面与命令行控制端(tdict debug start/exec/...),实现
"人操作 GDC 界面 + AI 借助命令行检查分析"的人机协同调试。

需要 config.json 中的 "debug" 配置节(ssh/zone 等)。`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil // 跳过 root 的 SQLite 打开逻辑
	},
}

// debugProbeCmd M0 尖刺:全自动跑通 登录→启动→下断点→步进→求值→退出
var debugProbeCmd = &cobra.Command{
	Use:   "probe",
	Short: "M0 尖刺:验证协议驱动器(全自动,不依赖 GDC 交互)",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath, err := resolveConfigPath(configPath)
		if err != nil {
			return err
		}
		cfg, err := debug.LoadConfig(cfgPath)
		if err != nil {
			return err
		}
		fmt.Printf("[probe] 服务器 %s:%d 区域 %s\n", cfg.SSH.Host, cfg.SSH.Port, cfg.Zone)

		sess, err := debug.NewSession(cfg, dbgModule, dbgProg, "", "", "", func(ev debug.Event) {
			switch ev.Type {
			case "state":
				fmt.Printf("[state] %s\n", ev.State)
			case "watchdog":
				fmt.Printf("[watchdog] %s\n", ev.Text)
			case "output":
				if verbose {
					fmt.Printf("[out] %q\n", ev.Text)
				}
			}
		})
		if err != nil {
			return err
		}
		defer sess.Quit()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() { // Ctrl-C 优雅退出
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
			<-sig
			cancel()
		}()

		fmt.Println("[probe] 启动调试会话(Launch)...")
		if err := sess.Launch(ctx); err != nil {
			return fmt.Errorf("Launch 失败: %w", err)
		}
		fmt.Printf("[probe] 已到 (fgldb) 入口停站 ✓\n\n")

		fmt.Println("[probe] break main(函数断点):")
		bp, err := sess.Break("main")
		if err != nil {
			return err
		}
		fmt.Printf("  Breakpoint %d at %s:%d ✓\n", bp.Num, bp.File, bp.Line)

		fmt.Println("\n[probe] run(等待命中断点)...")
		if _, err := sess.Run(); err != nil {
			return err
		}
		stop, err := sess.WaitForStop(60 * time.Second)
		if err != nil {
			return err
		}
		fmt.Printf("  命中: %s:%d (%s) ✓\n", stop.File, stop.Line, stop.Func)
		for _, sl := range stop.Source {
			mark := "   "
			if sl.IsCur {
				mark = "-> "
			}
			fmt.Printf("  %s%-6d %s\n", mark, sl.Num, sl.Text)
		}

		fmt.Println("\n[probe] where:")
		frames, err := sess.Where()
		if err != nil {
			return err
		}
		for _, f := range frames {
			fmt.Printf("  #%d %s at %s:%d\n", f.Idx, f.Func, f.File, f.Line)
		}

		fmt.Println("\n[probe] next:")
		stop2, err := sess.Step("next")
		if err != nil {
			return err
		}
		if stop2 != nil {
			fmt.Printf("  -> %s:%d (%s)\n", stop2.File, stop2.Line, stop2.Func)
			for _, sl := range stop2.Source {
				mark := "   "
				if sl.IsCur {
					mark = "-> "
				}
				fmt.Printf("  %s%-6d %s\n", mark, sl.Num, sl.Text)
			}
		}

		fmt.Println("\n[probe] print num_args():")
		v, err := sess.Print("num_args()")
		fmt.Printf("  %s (err=%v)\n", v, err)

		fmt.Println("\n[probe] 全流程验证完成,退出会话...")
		return nil
	},
}

// debugServeCmd M1+:启动本地服务(REST+WS+MCP+Web 前端)
var debugServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动本地调试服务(HTTP :8000,前端 + REST + WS)",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath, err := resolveConfigPath(configPath)
		if err != nil {
			return err
		}
		cfg, err := debug.LoadConfig(cfgPath)
		if err != nil {
			return err
		}
		cfg.ApplyActiveEnv() // 生效环境的连接/参数合并到顶层
		if dbgListen != "" {
			cfg.Listen = dbgListen
		}
		// 数据目录 = config.json 所在目录(断点持久化等)
		cfg.DataDir = filepath.Dir(cfgPath)
		srv := debug.NewServer(cfg, webFS, cfgPath)
		return srv.Run(context.Background())
	},
}

// debugDBCmd 数据库连接探查:登录服务器 → 查 gzou_t 企业→账号映射 → 尝试用对应账号连库
var debugDBCmd = &cobra.Command{
	Use:   "db",
	Short: "数据库连接探查:按企业(TOPENT)查账号并尝试连接",
	Long: `登录服务器,读取 gzou_t 中企业编号与数据库账号(schema)的映射,
并尝试用对应账号连接数据库(密码规则:账号=密码,取自 fglprofile 明文)。

  tdict debug db            列出全部企业→账号映射
  tdict debug db --ent 99   验证企业 99 对应账号的连接(默认企业取 config.json debug.db.ent)
  tdict debug db --ent 99 --json  输出 JSON`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath, err := resolveConfigPath(configPath)
		if err != nil {
			return err
		}
		cfg, err := debug.LoadConfig(cfgPath)
		if err != nil {
			return err
		}
		cfg.ApplyActiveEnv() // 生效环境的连接/参数合并到顶层
		ent := dbEnt
		if ent <= 0 {
			ent = cfg.DBEnt()
		}
		rep, err := debug.ProbeDB(cfg, ent)
		if err != nil {
			return err
		}
		if dbJSON {
			return printJSON(rep)
		}
		// 人类可读输出
		if kdb := rep.Env["database"]; kdb != "" {
			fmt.Printf("区域: %s   数据库: 人大金仓 %s@127.0.0.1:%s\n", rep.Zone, kdb, rep.Env["port"])
			if ks := rep.Env["ksql"]; ks != "" {
				fmt.Printf("ksql: %s\n", ks)
			}
		} else {
			fmt.Printf("区域: %s   TNS: %s   Oracle: %s\n",
				rep.Zone, rep.TNS, rep.Env["oracleHome"])
			if sp := rep.Env["sqlplus"]; sp != "" {
				fmt.Printf("sqlplus: %s\n", sp)
			}
		}
		fmt.Printf("\n企业(TOPENT) → 账号(schema),共 %d 个:\n", len(rep.Mappings))
		for _, m := range rep.Mappings {
			mark := " "
			if rep.Probe != nil && m.Ent == rep.Probe.Ent {
				mark = "*"
			}
			fmt.Printf("  %s %-6d → %s\n", mark, m.Ent, m.Account)
		}
		if rep.Probe != nil {
			p := rep.Probe
			fmt.Printf("\n企业 %d 连接验证: %s/%s@%s\n", p.Ent, p.Account, p.Account, rep.TNS)
			if p.Host != "" {
				fmt.Printf("  主机: %s:%s  服务名: %s\n", p.Host, p.Port, p.Service)
			}
			if p.Connect {
				fmt.Println("  结果: 连接成功 ✓")
			} else {
				fmt.Printf("  结果: 连接失败 ✗  %s\n", p.Error)
			}
			fmt.Println("  (密码规则为 fglprofile 明文:账号=密码;若失败说明该账号密码已改,需另行确认)")
		} else {
			fmt.Println("\n(未指定企业,仅列出映射。用 --ent <企业号> 验证连接)")
		}
		return nil
	},
}

// printJSON 以 JSON 输出结果
func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func init() {
	debugCmd.PersistentFlags().StringVarP(&dbgModule, "module", "m", "asf", "T100 模块目录名(如 asf)")
	debugProbeCmd.Flags().StringVarP(&dbgProg, "prog", "p", "bsft001_wf", "作业名(如 bsft001_wf)")
	debugProbeCmd.Flags().IntVarP(&dbgLine, "line", "l", 4452, "探针断点行号")
	debugServeCmd.Flags().StringVar(&dbgListen, "listen", "", "覆盖监听地址(默认取配置)")
	debugDBCmd.Flags().IntVar(&dbEnt, "ent", 0, "企业编号(TOPENT),验证该企业账号连接;0=取 config.json debug.db.ent")
	debugDBCmd.Flags().BoolVar(&dbJSON, "json", false, "输出 JSON")
	debugCmd.AddCommand(debugProbeCmd, debugServeCmd, debugDBCmd)
	rootCmd.AddCommand(debugCmd)
}
