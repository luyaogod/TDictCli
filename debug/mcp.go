package debug

// MCP Server:把调试会话以工具形式暴露给 AI(ZCode 等任意支持 MCP 的客户端)。
// 每次工具调用都会产生 ai_action 事件,同步显示在前端时间线上(可审计)。
import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// initMCP 构建 MCP server 并挂载到 /mcp
func (s *Server) initMCP() {
	srv := mcp.NewServer(&mcp.Implementation{Name: "tdict-debug", Version: "0.1.0"}, nil)

	// ---- 状态 ----
	mcp.AddTool(srv, &mcp.Tool{Name: "debug_status", Description: "查看调试服务状态:SSH 服务器/区域/当前会话摘要"}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		return nil, map[string]any{
			"ssh": s.cfg.SSH.Host, "zone": s.cfg.Zone, "topDir": s.cfg.TopDir,
			"watchdogSeconds": s.cfg.WatchdogSeconds, "sessions": s.mgr.Snapshot(),
		}, nil
	})

	// ---- 会话生命周期 ----
	mcp.AddTool(srv, &mcp.Tool{Name: "debug_launch", Description: "启动一个新的调试会话(fglrun -d)。作业的 GUI 界面会弹到用户的 GDC 上;同一时间只允许一个会话"}, func(ctx context.Context, req *mcp.CallToolRequest, in inLaunch) (*mcp.CallToolResult, any, error) {
		sess, err := s.mgr.Launch(in.Module, in.Prog)
		if err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI 启动调试会话 %s/%s", in.Module, in.Prog)
		if err := sess.Launch(ctx); err != nil {
			return nil, nil, err
		}
		return nil, map[string]any{"ok": true, "sessionId": sess.ID, "state": string(sess.State())}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_snapshot", Description: "获取当前会话快照:状态、停站位置、断点列表、停留时长"}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		cur := sess.Cur()
		return nil, map[string]any{
			"id": sess.ID, "module": sess.Module, "prog": sess.Prog,
			"state": string(sess.State()), "stop": cur,
			"breakpoints": sess.Breakpoints(), "holdingSeconds": sess.HoldingSeconds(),
			"autovars": sess.Autovars(),
		}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_quit", Description: "结束当前调试会话(quit,作业窗口随之关闭)"}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		s.aiLog(sess.ID, "AI 结束会话")
		id := sess.ID
		if err := sess.Quit(); err != nil {
			return nil, nil, err
		}
		s.mgr.Remove(id)
		return nil, map[string]any{"ok": true}, nil
	})

	// ---- 断点 ----
	mcp.AddTool(srv, &mcp.Tool{Name: "debug_break", Description: "下断点。位置支持:行号、函数名、file:line。必须在停站状态调用"}, func(ctx context.Context, req *mcp.CallToolRequest, in inLocation) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		bp, err := sess.Break(in.Location)
		if err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI 下断点 %s → #%d %s:%d", in.Location, bp.Num, bp.File, bp.Line)
		return nil, map[string]any{"breakpoint": bp}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_break_delete", Description: "删除断点"}, func(ctx context.Context, req *mcp.CallToolRequest, in inNum) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		if err := sess.DeleteBreakpoint(in.Num); err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI 删除断点 #%d", in.Num)
		return nil, map[string]any{"ok": true}, nil
	})

	// ---- 执行控制 ----
	mcp.AddTool(srv, &mcp.Tool{Name: "debug_control", Description: "执行控制:run(从入口启动)/continue(继续)/next(步过)/step(步入)/finish(步出)/until/interrupt(SIGINT 中断)。run/continue 后请用 debug_wait_stop 等待停站"}, func(ctx context.Context, req *mcp.CallToolRequest, in inAction) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		var stop *StopInfo
		var err error
		switch in.Action {
		case "continue":
			_, err = sess.Continue()
		case "run":
			_, err = sess.Run()
		case "next", "step", "finish":
			stop, err = sess.Step(in.Action)
		case "until":
			if in.Arg == "" {
				stop, err = sess.Step("until")
			} else {
				stop, err = sess.Step("until " + in.Arg)
			}
		case "interrupt":
			err = sess.Interrupt()
		default:
			return nil, nil, fmt.Errorf("未知 action: %s", in.Action)
		}
		if err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI 执行 %s%s", in.Action, argSuffix(in.Arg))
		resp := map[string]any{"ok": true, "state": string(sess.State())}
		if stop != nil {
			resp["stop"] = stop
		}
		return nil, resp, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_wait_stop", Description: "等待程序停站(断点命中或人工中断)。在下断点并 run/continue 后调用;人操作 GDC 触发断点后此工具返回停站现场"}, func(ctx context.Context, req *mcp.CallToolRequest, in inWait) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		timeout := 120 * time.Second
		if in.TimeoutSeconds > 0 && in.TimeoutSeconds <= 1800 {
			timeout = time.Duration(in.TimeoutSeconds * float64(time.Second))
		}
		stop, err := sess.WaitForStop(timeout)
		if err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI 等到停站: %s:%d (%s)", stop.File, stop.Line, stop.Func)
		return nil, map[string]any{"stop": stop, "state": string(sess.State()), "autovars": sess.Autovars()}, nil
	})

	// ---- 检查 ----
	mcp.AddTool(srv, &mcp.Tool{Name: "debug_print", Description: "求值表达式(print):变量、记录、数组、函数调用如 num_args()、arr.getLength()。必须在停站状态"}, func(ctx context.Context, req *mcp.CallToolRequest, in inExpr) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		v, err := sess.Print(in.Expr)
		if err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI print %s → %s", in.Expr, oneLine(v))
		return nil, map[string]any{"expr": in.Expr, "value": v}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_where", Description: "查看调用栈"}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		frames, err := sess.Where()
		if err != nil {
			return nil, nil, err
		}
		return nil, map[string]any{"frames": frames}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_source", Description: "读取作业源码(.4gl)。file 填停站现场里的源文件名(如 asf_bsft001_wf.4gl),自动解析到母版源文件"}, func(ctx context.Context, req *mcp.CallToolRequest, in inFile) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		sf, err := sess.ResolveSource(in.File, in.Module)
		if err != nil {
			return nil, nil, err
		}
		return nil, map[string]any{"path": sf.Path, "lines": strings.Count(sf.Content, "\n") + 1, "content": sf.Content}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_raw", Description: "透传任意 fgldb 调试命令(tbreak/watch/info/list/ptype/whatis/display/...)。quit 与 run 除外"}, func(ctx context.Context, req *mcp.CallToolRequest, in inCmd) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		cmd := strings.TrimSpace(in.Command)
		lower := strings.ToLower(cmd)
		if cmd == "" || strings.ContainsAny(cmd, "\r\n") {
			return nil, nil, fmt.Errorf("命令不能为空或含换行")
		}
		if lower == "quit" || lower == "run" || strings.HasPrefix(lower, "run ") {
			return nil, nil, fmt.Errorf("请使用 debug_control/debug_quit")
		}
		lines, err := sess.Raw(cmd, 30*time.Second)
		if err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI raw: %s", cmd)
		return nil, map[string]any{"lines": lines}, nil
	})

	// ---- 上下文查询(fgldeb 同款命令面) ----
	mcp.AddTool(srv, &mcp.Tool{Name: "debug_locals", Description: "获取当前帧的全部局部变量(info locals,名字+当前值)。必须停站;随 debug_frame 切换上下文"}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		vars, err := sess.Locals()
		if err != nil {
			return nil, nil, err
		}
		return nil, map[string]any{"vars": vars}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_globals", Description: "列出作业全部全局变量声明(info variables:名字+类型,T100 可达数千条)。用 filter 子串过滤、limit 截取"}, func(ctx context.Context, req *mcp.CallToolRequest, in inFilter) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		vars, total, err := sess.Globals(0)
		if err != nil {
			return nil, nil, err
		}
		vars = filterNames(vars, func(i int) string { return vars[i].Name }, in.Filter, in.Limit)
		return nil, map[string]any{"total": total, "matched": len(vars), "vars": vars}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_sources", Description: "列出作业已加载符号的全部模块源文件名(info sources)"}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		sources, err := sess.Sources()
		if err != nil {
			return nil, nil, err
		}
		return nil, map[string]any{"sources": sources}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_functions", Description: "列出作业全部函数名(info functions,T100 大作业上万条)。用 filter 子串过滤、limit 截取"}, func(ctx context.Context, req *mcp.CallToolRequest, in inFilter) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		fns, total, err := sess.Functions(0)
		if err != nil {
			return nil, nil, err
		}
		fns2 := filterNames(fns, func(i int) string { return fns[i] }, in.Filter, in.Limit)
		return nil, map[string]any{"total": total, "matched": len(fns2), "functions": fns2}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_frame", Description: "选择栈帧(frame N),此后 print/locals 在该帧上下文求值;新停站自动回到栈顶帧 0"}, func(ctx context.Context, req *mcp.CallToolRequest, in inFrame) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		if err := sess.Frame(in.Num); err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI 选择栈帧 #%d", in.Num)
		return nil, map[string]any{"ok": true, "frame": sess.CurFrame()}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_break_enable", Description: "启用/禁用断点(disable/enable),不必删除;禁用的断点不生效但保留"}, func(ctx context.Context, req *mcp.CallToolRequest, in inEnable) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		if err := sess.SetBPEnabled(in.Num, in.Enabled); err != nil {
			return nil, nil, err
		}
		s.aiLog(sess.ID, "AI %s 断点 #%d", enableWord(in.Enabled), in.Num)
		return nil, map[string]any{"ok": true}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{Name: "debug_info_line", Description: "让 fgldb 解析位置为源文件与行号(info line):传函数名或 file:line,返回权威的 DVM 源文件名+行号;用于下断点前定位"}, func(ctx context.Context, req *mcp.CallToolRequest, in inLocation) (*mcp.CallToolResult, any, error) {
		sess := s.current()
		if sess == nil {
			return nil, nil, errNoSession
		}
		file, line, err := sess.InfoLine(in.Location)
		if err != nil {
			return nil, nil, err
		}
		return nil, map[string]any{"file": file, "line": line}, nil
	})

	s.mcpHandler = mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return srv }, nil)
}

// ---------- 输入结构 ----------

type inLaunch struct {
	Module string `json:"module,omitempty" jsonschema:"T100 模块目录名,如 asf;留空则按作业名自动查找"`
	Prog   string `json:"prog" jsonschema:"作业名,如 bsft001_wf"`
}
type inLocation struct {
	Location string `json:"location" jsonschema:"断点位置:行号、函数名或 file:line"`
}
type inNum struct {
	Num int `json:"num" jsonschema:"断点编号"`
}
type inAction struct {
	Action string `json:"action" jsonschema:"continue|run|next|step|finish|until|interrupt"`
	Arg    string `json:"arg,omitempty" jsonschema:"until 的目标位置(可选)"`
}
type inWait struct {
	TimeoutSeconds float64 `json:"timeoutSeconds" jsonschema:"最长等待秒数,默认120,上限1800"`
}
type inExpr struct {
	Expr string `json:"expr" jsonschema:"4gl 表达式"`
}
type inCmd struct {
	Command string `json:"command" jsonschema:"fgldb 命令"`
}
type inFile struct {
	File   string `json:"file" jsonschema:"DVM 报告的源文件名,如 asf_bsft001_wf.4gl"`
	Module string `json:"module,omitempty" jsonschema:"模块目录名,如 asf"`
}
type inFilter struct {
	Filter string `json:"filter,omitempty" jsonschema:"名字子串过滤(不区分大小写)"`
	Limit  int    `json:"limit,omitempty" jsonschema:"返回条数上限,默认 500"`
}
type inFrame struct {
	Num int `json:"num" jsonschema:"栈帧编号(where 输出的 #N)"`
}
type inEnable struct {
	Num     int  `json:"num" jsonschema:"断点编号"`
	Enabled bool `json:"enabled" jsonschema:"true 启用 / false 禁用"`
}

var errNoSession = fmt.Errorf("没有活跃的调试会话,请先 debug_launch")

// ---------- 小工具 ----------

func (s *Server) current() *Session { return s.mgr.Current() }

func (s *Server) aiLog(sessionID, format string, args ...any) {
	s.mgr.emit(Event{Type: "ai_action", SessionID: sessionID, Text: fmt.Sprintf(format, args...)})
}

func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\n", " ⏎ ")
	if len(s) > 300 {
		s = s[:300] + "…"
	}
	return s
}

func argSuffix(arg string) string {
	if arg == "" {
		return ""
	}
	return " " + arg
}

func enableWord(on bool) string {
	if on {
		return "启用"
	}
	return "禁用"
}

// filterNames 通用子串过滤 + 截取(不区分大小写)
func filterNames[T any](items []T, nameOf func(int) string, filter string, limit int) []T {
	if limit <= 0 {
		limit = 500
	}
	lf := strings.ToLower(filter)
	out := make([]T, 0, len(items))
	for i := range items {
		if lf != "" && !strings.Contains(strings.ToLower(nameOf(i)), lf) {
			continue
		}
		out = append(out, items[i])
		if len(out) >= limit {
			break
		}
	}
	return out
}
