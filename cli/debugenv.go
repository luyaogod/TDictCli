package cli

// debug 环境/TOPENT 切换命令:
//   tdict debug env           列出全部已配置环境与当前生效项(基于 /api/settings + /api/sessions)
//   tdict debug env <名称>     把该环境设为默认(持久化 activeEnv)并切换当前会话到它(等会话回 idle)
//   tdict debug topent        显示当前会话的 TOPENT override 与配置级默认
//   tdict debug topent <值>    设置会话级 TOPENT override(仅 idle;下一轮调试生效)
//   tdict debug topent --clear 清除会话级 TOPENT override(回退配置默认)
//
// 与 debugctl 其它命令一致:均经 tdict debug serve 的 REST 接口执行。

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	dbgEnvTimeout int
	dbgTopClear   bool
)

// dbgSnapshot 会话快照中本命令关心的字段。
type dbgSnapshot struct {
	ID        string `json:"id"`
	Env       string `json:"env"`
	State     string `json:"state"`
	Topent    string `json:"topent"`    // 会话内手动设置的 TOPENT override(空=未设置)
	TopentCfg string `json:"topentCfg"` // 配置级企业 TOPENT 默认(该环境 db.ent)
}

// dbgEnvItem 环境清单中的一项(取自 /api/settings 返回的完整配置)。
type dbgEnvItem struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Zone string `json:"zone"`
	DB   *struct {
		Type string `json:"type"`
		Ent  string `json:"ent"`
	} `json:"db"`
}

// dbgSettingsView /api/settings 返回的顶层 debug 配置(仅本命令需要)。
type dbgSettingsView struct {
	ActiveEnv string       `json:"activeEnv"`
	SSH       dbgEnvSSH    `json:"ssh"` // 未配置 envs 时顶层的直接连接
	Envs      []dbgEnvItem `json:"envs"`
}

type dbgEnvSSH struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
}

// debugEnvCmd 查看/切换当前生效的调试环境(SSH 配置)。
var debugEnvCmd = &cobra.Command{
	Use:   "env [环境名]",
	Short: "查看/切换当前生效的调试环境(SSH 配置,需 serve 在运行)",
	Long: `查看或切换当前生效的调试环境(SSH 连接配置)。

无参数:列出 config.json debug.envs 中的全部环境、当前默认(activeEnv)与当前会话所在环境/状态。
带环境名:把该环境设为默认(写入 config.json activeEnv,持久化)并切换会话到它——
先结束当前调试(若有),再按新环境的 SSH/区域重连到空闲(idle)。

需要 tdict debug serve 在运行;控制端命令会自动发现其地址。`,
	Example: `  tdict debug env                 # 列出环境与当前生效项
  tdict debug env 恒烁测试区       # 切换默认/会话到该环境(等会话回空闲)
  tdict debug env --json          # 环境清单 JSON 输出`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return dbgEnvSwitch(args[0])
		}
		return dbgEnvList()
	},
}

// debugTopentCmd 查看/设置会话级 TOPENT override。
var debugTopentCmd = &cobra.Command{
	Use:   "topent [值]",
	Short: "查看/设置会话 TOPENT override(需 serve 在运行,会话空闲可设)",
	Long: `查看或设置当前会话的 TOPENT(企业编号)override。

无参数:显示当前会话所在环境/状态、会话内 TOPENT override 与配置级默认(该环境 db.ent)。
带值:仅在会话空闲(idle)时可设置,下一轮调试启动时采用;值不限数字/文本,
服务端会剔除两侧空白。--clear 等价于传空值:清除会话级 override,回退配置默认。

需要 tdict debug serve 在运行且已有空闲会话(无会话时先 tdict debug env <环境名> 连接)。`,
	Example: `  tdict debug topent          # 查看当前会话 TOPENT
  tdict debug topent 99       # 设置会话 TOPENT=99(仅 idle)
  tdict debug topent --clear  # 清除会话级 override`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		value := ""
		if len(args) == 1 {
			value = strings.TrimSpace(args[0])
		}
		if dbgTopClear {
			value = ""
		}
		if len(args) == 0 && !dbgTopClear {
			return dbgTopentShow()
		}
		return dbgTopentSet(value)
	},
}

// dbgSettings 拉取 /api/settings 并解出环境清单。
func dbgSettings() (*dbgSettingsView, error) {
	data, err := dbgAPI("GET", "/api/settings", nil)
	if err != nil {
		return nil, err
	}
	var s dbgSettingsView
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("解析设置失败: %w", err)
	}
	return &s, nil
}

// dbgSessions 拉取会话列表。
func dbgSessions() ([]struct {
	ID    string `json:"id"`
	Env   string `json:"env"`
	State string `json:"state"`
}, error) {
	data, err := dbgAPI("GET", "/api/sessions", nil)
	if err != nil {
		return nil, err
	}
	var r struct {
		Sessions []struct {
			ID    string `json:"id"`
			Env   string `json:"env"`
			State string `json:"state"`
		} `json:"sessions"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("解析会话列表失败: %w", err)
	}
	return r.Sessions, nil
}

// dbgEnvList 列出环境与当前生效项。
func dbgEnvList() error {
	s, err := dbgSettings()
	if err != nil {
		return err
	}
	sessions, err := dbgSessions()
	if err != nil {
		return err
	}
	// 会话为单一常驻会话(通常一条)
	curEnv, curState := "", ""
	if len(sessions) > 0 {
		curEnv, curState = sessions[0].Env, sessions[0].State
	}
	if IsJSON() {
		type out struct {
			ActiveEnv string `json:"activeEnv"`
			Envs      []any  `json:"envs"`
			Session   any    `json:"session,omitempty"`
		}
		envs := make([]any, 0, len(s.Envs))
		for _, e := range s.Envs {
			envs = append(envs, e)
		}
		var sess any
		if curEnv != "" {
			sess = map[string]any{"env": curEnv, "state": curState}
		}
		return printJSON(out{ActiveEnv: s.ActiveEnv, Envs: envs, Session: sess})
	}

	fmt.Printf("当前生效(activeEnv): %s\n", orDefault(s.ActiveEnv, "未设置(用顶层 ssh)"))
	if curEnv != "" {
		fmt.Printf("当前会话:           %s (%s)\n", curEnv, dbgStateLabel(curState))
	}
	if len(s.Envs) == 0 {
		fmt.Printf("环境: 未配置 envs;直连 %s:%d (user %s)\n", s.SSH.Host, s.SSH.Port, s.SSH.User)
		return nil
	}
	fmt.Printf("环境(%d):\n", len(s.Envs))
	for _, e := range s.Envs {
		mark := "  "
		if s.ActiveEnv == e.Name {
			mark = " *"
		}
		name := e.Name
		if curEnv == e.Name {
			name += " ← 当前会话"
		}
		dbType, dbEnt := "", ""
		if e.DB != nil {
			dbType = e.DB.Type
			dbEnt = e.DB.Ent
		}
		extra := dbType
		if dbEnt != "" {
			extra += " (TOPENT默认=" + dbEnt + ")"
		}
		fmt.Printf("  %s %-12s %s:%d  zone=%-4s %s\n", mark, name, e.Host, e.Port, e.Zone, extra)
	}
	return nil
}

// dbgEnvSwitch 切换默认环境并把会话切过去(等会话回到 idle/exit)。
func dbgEnvSwitch(name string) error {
	// 先校验环境存在(给出更友好的错误)
	s, err := dbgSettings()
	if err != nil {
		return err
	}
	found := false
	for _, e := range s.Envs {
		if e.Name == name {
			found = true
			break
		}
	}
	if !found && s.ActiveEnv != name {
		return fmt.Errorf("环境 %q 不存在(可先运行 tdict debug env 查看已配置环境)", name)
	}
	// 目标不是当前会话环境、且会话不在空闲时,切换会结束当前调试——先提示(Web 端同样要求确认)
	if sessions, e2 := dbgSessions(); e2 == nil && len(sessions) > 0 {
		cur := sessions[0]
		if cur.Env != name && cur.State != "idle" && cur.State != "exit" {
			fmt.Printf("注意:切换会结束当前调试(会话状态 %s)并重连到 %s …\n", dbgStateLabel(cur.State), name)
		}
	}

	data, err := dbgAPI("POST", "/api/sessions/switch", map[string]any{"env": name})
	if err != nil {
		return err
	}
	var r struct {
		SessionID string `json:"sessionId"`
		Env       string `json:"env"`
		State     string `json:"state"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return fmt.Errorf("解析切换结果失败: %w", err)
	}
	fmt.Printf("切换会话到 %s …\n", r.Env)

	// 服务端异步登录到 idle;已在目标环境则幂等返回当前状态,不等 idle
	if r.State == "loading" || r.State == "" {
		deadline := time.Now().Add(time.Duration(dbgEnvTimeout) * time.Second)
		for {
			snap, err := dbgSnap(r.SessionID)
			if err == nil && snap != nil {
				if snap.State == "idle" {
					fmt.Printf("已切换:环境 %s 会话空闲(idle),可启动调试\n", r.Env)
					return nil
				}
				if snap.State == "exit" {
					return fmt.Errorf("切换失败:会话 %s 已断开", r.Env)
				}
			} else if err != nil {
				// 快照失败可能是会话被移除(登录失败):确认真不在列表里再提前报错
				gone := true
				if sessions, e2 := dbgSessions(); e2 == nil {
					for _, s2 := range sessions {
						if s2.ID == r.SessionID {
							gone = false
							break
						}
					}
				}
				if gone {
					return fmt.Errorf("切换失败:会话 %s 未建立成功(环境 %s 连接失败)", r.SessionID, r.Env)
				}
			}
			if time.Now().After(deadline) {
				return fmt.Errorf("切换会话超时(%ds),请用 tdict debug status 查看状态", dbgEnvTimeout)
			}
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Printf("已在环境 %s(会话状态 %s),无需切换\n", r.Env, dbgStateLabel(r.State))
	return nil
}

// dbgSnap 取会话快照。
func dbgSnap(id string) (*dbgSnapshot, error) {
	if id == "" {
		return nil, fmt.Errorf("会话 id 为空")
	}
	data, err := dbgAPI("GET", "/api/sessions/"+id, nil)
	if err != nil {
		return nil, err
	}
	var s dbgSnapshot
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("解析会话快照失败: %w", err)
	}
	return &s, nil
}

// dbgTopentShow 显示当前会话 TOPENT override 与配置默认。
func dbgTopentShow() error {
	sessions, err := dbgSessions()
	if err != nil {
		return err
	}
	if len(sessions) == 0 {
		return fmt.Errorf("没有活动会话;先运行 tdict debug env <环境名> 连接,或 tdict debug start 启动调试")
	}
	id, env, st := sessions[0].ID, sessions[0].Env, sessions[0].State
	snap, err := dbgSnap(id)
	if err != nil {
		return err
	}
	fmt.Printf("会话: %s (%s)\n", env, dbgStateLabel(st))
	if snap.Topent != "" {
		fmt.Printf("会话内 TOPENT override: %s\n", snap.Topent)
	} else {
		fmt.Println("会话内 TOPENT override: (未设置)")
	}
	fmt.Printf("配置级默认(该环境 db.ent): %s\n", orDefault(snap.TopentCfg, "(未配置,沿用登录默认)"))
	return nil
}

// dbgTopentSet 设置/清除会话 TOPENT override(仅 idle)。
func dbgTopentSet(value string) error {
	sessions, err := dbgSessions()
	if err != nil {
		return err
	}
	if len(sessions) == 0 {
		return fmt.Errorf("没有活动会话;先运行 tdict debug env <环境名> 建立空闲会话,再设置 TOPENT")
	}
	id, st := sessions[0].ID, sessions[0].State
	if st != "idle" {
		return fmt.Errorf("仅会话空闲(idle)时可设置 TOPENT(当前 %s);请先结束调试:tdict debug quit", dbgStateLabel(st))
	}
	_, err = dbgAPI("POST", "/api/sessions/"+id+"/topent", map[string]any{"value": value})
	if err != nil {
		return err
	}
	if value == "" {
		fmt.Println("已清除会话 TOPENT override(下一轮调试回退配置默认)")
	} else {
		fmt.Printf("已设置会话 TOPENT = %q(下一轮调试生效)\n", value)
	}
	return nil
}

// dbgStateLabel 状态中文名。
func dbgStateLabel(s string) string {
	switch s {
	case "idle":
		return "空闲"
	case "loading":
		return "连接中…"
	case "stopped":
		return "已停站"
	case "running":
		return "运行中"
	case "exit":
		return "已断开"
	default:
		return s
	}
}

func orDefault(v, d string) string {
	if v == "" {
		return d
	}
	return v
}

func init() {
	debugEnvCmd.Flags().IntVar(&dbgEnvTimeout, "timeout", 120, "切换等待会话空闲超时(秒)")
	debugTopentCmd.Flags().BoolVar(&dbgTopClear, "clear", false, "清除会话级 TOPENT override(等价传空值)")
	debugCmd.AddCommand(debugEnvCmd, debugTopentCmd)
}
