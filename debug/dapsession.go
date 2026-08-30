package debug

// dapsession.go — DAP 协议调试会话(替代 PTY 刮屏)。
// 经 SSH exec 通道启动 fglrun --da-debugger(Genero 3.21+ 内嵌适配器),
// 被调程序由适配器拉起为 `fglrun --da-wait <适配器pid> <program> <args>`。
// Phase 1 原型已验证:initialize/launch{program,args,cwd,stopAtEntry}/
// setBreakpoints/configurationDone/stackTrace/scopes/variables/evaluate。
// 注意:适配器报告的源码路径在模块的 42m 目录(如 .../asf/42m/asf_xxx.4gl)。

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type DAPSession struct {
	cfg          *Config
	id           string
	module       string
	prog         string
	runProg      string
	argsOverride string
	emit         func(Event)

	conn   *SSHConn
	kaStop func()
	proc   *ssh.Session   // 适配器进程(SSH exec 通道)
	stdin  io.WriteCloser // 适配器 stdin(DAP 帧写入口)
	dap    *dapClient

	mu        sync.Mutex
	state     State
	started   bool
	startAt   time.Time
	cur       StopInfo
	curFrame  int
	frames    []Frame
	frameDap  []int // 与 frames 对齐的 DAP frameId
	bps       []*Breakpoint
	restoring bool
	quitting  bool

	localsRef  int // 最近停站的 Locals scope variablesReference(-1 无)
	globalsRef int
	stopFrame  int // 最近停站的 DAP frameId(断点/变量默认上下文)
	initCh     chan struct{} // initialized 事件信号(pump 独占事件流,经此通知 Launch)
	stopGen    uint64        // 停站代数:每次 stopped 事件 +1(等待"新"停站用)
	autovars   []VarItem

	srcCache map[string]srcCacheEntry
	sftpCl   *sftp.Client
	sftpMu   sync.Mutex
}

// NewDAPSession 建立 SSH 连接并启动 DAP 适配器(登录环境由 Launch 驱动)
func NewDAPSession(cfg *Config, module, prog, runProg string, emit func(Event)) (session, error) {
	conn, err := Dial(cfg.SSH)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %w", err)
	}
	s := &DAPSession{
		cfg:      cfg,
		id:       fmt.Sprintf("d%d", time.Now().UnixMilli()),
		module:   module,
		prog:     prog,
		runProg:  runProg,
		emit:     emit,
		conn:     conn,
		state:    StateLoading,
		localsRef:  -1,
		globalsRef: -1,
		initCh:   make(chan struct{}),
		srcCache: map[string]srcCacheEntry{},
	}
	s.kaStop = conn.StartKeepalive(func(err error) {
		s.emitEvent(Event{Type: "dead", Text: "SSH 连接断开: " + err.Error()})
		s.ForceExit()
	})
	return s, nil
}

// ---------- 标识 ----------

func (s *DAPSession) ID() string            { return s.id }
func (s *DAPSession) Module() string        { return s.module }
func (s *DAPSession) Prog() string          { return s.prog }
func (s *DAPSession) RunProg() string       { return s.runProg }
func (s *DAPSession) SetArgsOverride(a string) { s.argsOverride = a }
func (s *DAPSession) Log(text string)       { s.emitEvent(Event{Type: "log", Text: text}) }

func (s *DAPSession) emitEvent(ev Event) {
	ev.SessionID = s.id
	ev.Time = time.Now()
	s.emit(ev)
}

// ---------- 生命周期 ----------

// Launch 启动适配器与被调程序,直到入口停站
func (s *DAPSession) Launch(ctx context.Context) error {
	launchProg := s.runProg
	if launchProg == "" {
		launchProg = s.prog
	}
	if s.module == "" {
		mod, err := s.findModule(launchProg)
		if err != nil {
			return err
		}
		s.module = mod
	}
	s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("DAP 模式启动:%s(模块 %s)", launchProg, s.module)})
	if s.runProg != "" && s.runProg != s.prog {
		s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("作业编号 %s → 实体程序 %s(gzzz_t)", s.prog, s.runProg)})
	}

	// exec 通道:登录环境(.profile <zone> 参数化跳过交互菜单)+ FGLSERVER 兜底
	// (exec 无 tty,who -m 取不到 GDC 来源 IP,退回 SSH_CLIENT)
	moduleDir := s.cfg.ModuleDir(s.module)
	execCmd := fmt.Sprintf(
		"ksh -c '. /u1/usr/youruser/.profile %s >/dev/null 2>&1; "+
			"[ -z \"$FGLSERVER\" ] && export FGLSERVER=$(echo $SSH_CLIENT | cut -d\" \" -f1); "+
			"cd %s; exec fglrun --da-debugger' 2>/tmp/tdict_dap_%s.err",
		s.cfg.Zone, moduleDir, s.id)
	sess, err := s.conn.cli.NewSession()
	if err != nil {
		return fmt.Errorf("启动适配器失败: %w", err)
	}
	stdin, err := sess.StdinPipe()
	if err != nil {
		sess.Close()
		return fmt.Errorf("适配器管道失败: %w", err)
	}
	stdout, err := sess.StdoutPipe()
	if err != nil {
		sess.Close()
		return fmt.Errorf("适配器管道失败: %w", err)
	}
	stderr, err := sess.StderrPipe()
	if err != nil {
		sess.Close()
		return fmt.Errorf("适配器管道失败: %w", err)
	}
	if err := sess.Start(execCmd); err != nil {
		sess.Close()
		return fmt.Errorf("启动适配器失败: %w", err)
	}
	// stderr 已由命令本身重定向到服务器临时文件,这里兜底丢弃
	go func() { buf := make([]byte, 4096); for { if _, err := stderr.Read(buf); err != nil { return } } }()
	s.proc = sess
	s.stdin = stdin
	s.dap = newDAPClient(stdout, stdin)
	go s.pump()

	// initialize → launch
	if _, err := s.dap.request("initialize", map[string]any{"adapterID": "tdict"}, 15*time.Second); err != nil {
		return err
	}
	progPath := moduleDir + "/42r/" + launchProg + ".42r"
	argsSrc := s.argsOverride
	if argsSrc == "" {
		argsSrc = strings.ReplaceAll(s.cfg.LaunchArgs, "{prog}", s.prog)
	}
	launchArgs := map[string]any{
		"program":     progPath,
		"args":        splitShellArgs(argsSrc),
		"cwd":         moduleDir,
		"stopAtEntry": true,
	}
	if _, err := s.dap.request("launch", launchArgs, 30*time.Second); err != nil {
		return err
	}

	// 等 initialized 事件 → 恢复断点 → configurationDone
	if err := s.waitInitialized(60 * time.Second); err != nil {
		return err
	}
	s.mu.Lock()
	s.restoring = true
	s.mu.Unlock()
	if saved := s.restoreBPs(); len(saved) > 0 {
		s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("断点恢复:%d 个(持久化存档)", len(saved))})
	}
	s.mu.Lock()
	s.restoring = false // 断点已下发,此后停站对前端可见
	s.mu.Unlock()
	if _, err := s.dap.request("configurationDone", map[string]any{}, 15*time.Second); err != nil {
		return err
	}
	// 等入口停站(stopped 事件由 pump 处理;restoring 标志保证停站宣告时断点已就绪)
	if err := s.waitStopped(150 * time.Second); err != nil {
		return fmt.Errorf("等待入口停站超时: %w", err)
	}
	s.emitEvent(Event{Type: "log", Text: "调试会话就绪: " + s.prog + "@" + s.cfg.Zone + "(DAP)"})
	return nil
}

// findModule 模块为空时按 42r 文件名搜索兜底(与 PTY 版语义一致)
func (s *DAPSession) findModule(prog string) (string, error) {
	conn, err := Dial(s.cfg.SSH)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	if cands := searchModule42r(conn, s.cfg.ModuleRoots, prog); len(cands) == 1 {
		return cands[0], nil
	}
	return "", fmt.Errorf("无法定位程序 %s 的模块目录(可显式指定模块)", prog)
}

// waitInitialized 等待 initialized 事件(由 pump 发信号,避免双消费者)
func (s *DAPSession) waitInitialized(timeout time.Duration) error {
	select {
	case <-s.initCh:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("等待 initialized 事件超时")
	}
}

// waitStopped 等待下一次 stopped 事件(pump 已完成栈采集)
func (s *DAPSession) waitStopped(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for {
		s.mu.Lock()
		st := s.state
		s.mu.Unlock()
		if st == StateStopped {
			return nil
		}
		if st == StateExit {
			return fmt.Errorf("程序已退出")
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("超时")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// restoreBPs 从持久化存档恢复断点(setBreakpoints 按源码文件整批下发)
func (s *DAPSession) restoreBPs() []StoredBP {
	if !s.cfg.BPsPersisted() {
		return nil
	}
	st, err := loadBPs(s.cfg.DataDir, s.module, s.prog)
	if err != nil || len(st.Breakpoints) == 0 {
		return nil
	}
	lines := make([]int, 0, len(st.Breakpoints))
	for _, sb := range st.Breakpoints {
		lines = append(lines, sb.Line)
	}
	bps := s.setBreakpointsAt(s.dapSourcePath(), lines)
	s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("断点恢复完成:%d 成功(共 %d)", len(bps), len(st.Breakpoints))})
	return st.Breakpoints
}

// Quit 结束会话:disconnect 终止被调程序并释放连接
func (s *DAPSession) Quit() error {
	s.mu.Lock()
	if s.quitting {
		s.mu.Unlock()
		return nil
	}
	s.quitting = true
	s.mu.Unlock()
	if s.dap != nil {
		_, _ = s.dap.request("disconnect", map[string]any{"terminateDebuggee": true}, 8*time.Second)
		s.dap.close()
	}
	s.cleanup()
	return nil
}

// ForceExit 强制终结(启动失败等场景)
func (s *DAPSession) ForceExit() { s.cleanup() }

func (s *DAPSession) cleanup() {
	s.mu.Lock()
	already := s.state == StateExit
	s.state = StateExit
	s.mu.Unlock()
	if s.kaStop != nil {
		s.kaStop()
	}
	if s.stdin != nil {
		_ = s.stdin.Close()
	}
	if s.proc != nil {
		_ = s.proc.Close()
	}
	if s.conn != nil {
		s.conn.Close()
	}
	// 清理服务器侧 stderr 文件与可能的残留进程
	if s.conn != nil || s.id != "" {
		if conn, err := Dial(s.cfg.SSH); err == nil {
			_, _ = conn.Output("rm -f /tmp/tdict_dap_"+s.id+".err; true", 5*time.Second)
			conn.Close()
		}
	}
	if !already {
		s.emitEvent(Event{Type: "state", State: string(StateExit)})
		s.emitEvent(Event{Type: "log", Text: "调试会话已结束,SSH 连接已释放(可启动其它作业)"})
	}
}

// Interrupt 运行中请求中断(pause)
func (s *DAPSession) Interrupt() error {
	s.mu.Lock()
	st := s.state
	started := s.started
	s.mu.Unlock()
	if !started {
		return fmt.Errorf("程序尚未启动,无需中断:直接「继续」即可运行,或点行号下断点")
	}
	if st != StateRunning {
		return fmt.Errorf("仅运行中可中断(当前 %s)", st)
	}
	s.emitEvent(Event{Type: "log", Text: "请求中断(pause)"})
	_, err := s.dap.request("pause", map[string]any{"threadId": 1}, 10*time.Second)
	return err
}

// ---------- 控制 ----------

func (s *DAPSession) Continue() (*execResult, error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return nil, ErrNotStopped
	}
	s.mu.Unlock()
	if _, err := s.dap.request("continue", map[string]any{"threadId": 1}, 15*time.Second); err != nil {
		return nil, err
	}
	s.setRunning()
	return &execResult{Cmd: "continue", Continuing: true}, nil
}

func (s *DAPSession) Run() (*execResult, error) { return s.Continue() }

// Step 单步:next→步过,step→步入,finish→步出;until 暂不支持。
// 入口停站(pre-run)直接 next 无效(原型验证):仿 PTY stepFromEntry,
// 在当前行下临时断点后 continue,命中即删。
func (s *DAPSession) Step(cmd string) (*StopInfo, error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return nil, ErrNotStopped
	}
	entry := !s.started
	s.mu.Unlock()

	var dcmd string
	switch {
	case cmd == "next":
		dcmd = "next"
	case cmd == "step":
		dcmd = "stepIn"
	case cmd == "finish":
		dcmd = "stepOut"
	default:
		return nil, fmt.Errorf("DAP 模式暂不支持 %s", cmd)
	}

	if entry {
		cur := s.Cur()
		line := cur.Line
		if line <= 0 {
			return nil, fmt.Errorf("入口步进失败:未知当前行")
		}
		s.emitEvent(Event{Type: "log", Text: "入口步进:临时断点后启动程序"})
		if len(s.setBreakpointsAt(s.dapSourcePath(), []int{line})) == 0 {
			return nil, fmt.Errorf("入口步进失败:临时断点未通过")
		}
		if _, err := s.Continue(); err != nil {
			return nil, err
		}
		stop, err := s.WaitForStop(60 * time.Second)
		// 删临时断点(重新下发空列表)
		s.setBreakpointsAt(s.dapSourcePath(), nil)
		if err != nil || stop == nil {
			return nil, fmt.Errorf("入口步进失败:未在预期时间内停站")
		}
		s.mu.Lock()
		s.started = true
		stop.Reason = "step"
		s.mu.Unlock()
		return stop, nil
	}

	if _, err := s.dap.request(dcmd, map[string]any{"threadId": 1}, 15*time.Second); err != nil {
		return nil, err
	}
	s.setRunning()
	stop, err := s.WaitForStop(60 * time.Second)
	if err != nil || stop == nil {
		return nil, fmt.Errorf("步进后未停站(程序可能已运行/退出)")
	}
	s.mu.Lock()
	s.started = true
	s.mu.Unlock()
	return stop, nil
}

// WaitForStop 等待"新的一次"停站(以停站代数为准,避免把上一次的停站当新鲜数据)
func (s *DAPSession) WaitForStop(timeout time.Duration) (*StopInfo, error) {
	gen := s.curStopGen()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if s.curStopGen() != gen {
			s.mu.Lock()
			st := s.state
			s.mu.Unlock()
			if st == StateStopped {
				cur := s.Cur()
				return &cur, nil
			}
		}
		s.mu.Lock()
		st := s.state
		s.mu.Unlock()
		if st == StateExit {
			return nil, fmt.Errorf("程序已退出")
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil, fmt.Errorf("等待停站超时(%s)", timeout)
}

func (s *DAPSession) curStopGen() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stopGen
}

// ---------- 状态 ----------

func (s *DAPSession) Started() bool   { s.mu.Lock(); defer s.mu.Unlock(); return s.started }
func (s *DAPSession) CurFrame() int   { s.mu.Lock(); defer s.mu.Unlock(); return s.curFrame }
func (s *DAPSession) HoldingSeconds() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.state == StateStopped && !s.startAt.IsZero() {
		return time.Since(s.startAt).Seconds()
	}
	return 0
}

func (s *DAPSession) State() State {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 断点恢复期间对外仍报 loading(与 PTY 版一致:前端见到 stopped 即停轮询)
	if s.state == StateStopped && s.restoring {
		return StateLoading
	}
	return s.state
}

func (s *DAPSession) Cur() StopInfo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.cur
}

func (s *DAPSession) Restoring() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.restoring
}

func (s *DAPSession) setRunning() {
	s.mu.Lock()
	s.state = StateRunning
	s.started = true
	s.cur = StopInfo{}
	s.localsRef, s.globalsRef = -1, -1
	s.autovars = nil
	s.mu.Unlock()
	s.emitEvent(Event{Type: "state", State: string(StateRunning)})
}

// ---------- 事件泵 ----------

func (s *DAPSession) pump() {
	for {
		select {
		case ev, ok := <-s.dap.events():
			if !ok {
				return
			}
			s.handleEvent(ev)
		case <-s.dap.closed:
			s.mu.Lock()
			st := s.state
			s.mu.Unlock()
			if st != StateExit {
				s.emitEvent(Event{Type: "log", Text: "适配器连接关闭"})
				s.ForceExit()
			}
			return
		}
	}
}

func (s *DAPSession) handleEvent(ev *dapFrame) {
	switch ev.Event {
	case "initialized":
		select {
		case <-s.initCh:
		default:
			close(s.initCh)
		}
	case "stopped":
		reason := "unknown"
		var bpNum int
		var body struct {
			Reason           string `json:"reason"`
			ThreadID         int    `json:"threadId"`
			HitBreakpointIDs []int  `json:"hitBreakpointIds"`
		}
		_ = json.Unmarshal(ev.Body, &body)
		switch body.Reason {
		case "entry":
			reason = "entry"
		case "breakpoint":
			reason = "breakpoint"
			if len(body.HitBreakpointIDs) > 0 {
				bpNum = body.HitBreakpointIDs[0]
			}
		case "step":
			reason = "step"
		case "pause":
			reason = "interrupt"
		}
		// 采集调用栈(停站后帧 id 变化,必须重新取)
		s.refreshStack(body.ThreadID)
		s.mu.Lock()
		s.state = StateStopped
		s.startAt = time.Now()
		cur := s.cur
		cur.Reason = reason
		cur.BPNum = bpNum
		s.cur = cur
		s.curFrame = 0
		s.localsRef, s.globalsRef = -1, -1
		s.autovars = nil
		s.stopGen++
		s.mu.Unlock()
		s.emitEvent(Event{Type: "state", State: string(StateStopped)})
		s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("停站:%s %s:%d", reason, cur.File, cur.Line)})
	case "continued":
		s.setRunning()
	case "exited", "terminated":
		s.emitEvent(Event{Type: "log", Text: "程序已退出(" + ev.Event + ")"})
		s.cleanup()
	case "output":
		var body struct {
			Output   string `json:"output"`
			Category string `json:"category"`
		}
		_ = json.Unmarshal(ev.Body, &body)
		if out := strings.TrimSpace(body.Output); out != "" {
			s.emitEvent(Event{Type: "log", Text: "适配器: " + out})
		}
	case "process":
		var body struct {
			Name       string `json:"name"`
			SystemPID  int    `json:"systemProcessId"`
			IsLocalUse bool   `json:"isLocalProcess"`
		}
		_ = json.Unmarshal(ev.Body, &body)
	}
}

// refreshStack 拉取调用栈并更新 frames/cur
func (s *DAPSession) refreshStack(threadID int) {
	resp, err := s.dap.request("stackTrace", map[string]any{"threadId": threadID, "levels": 20}, 15*time.Second)
	if err != nil {
		return
	}
	var body struct {
		StackFrames []struct {
			ID     int `json:"id"`
			Name   string `json:"name"`
			Line   int    `json:"line"`
			Source struct {
				Name string `json:"name"`
				Path string `json:"path"`
			} `json:"source"`
		} `json:"stackFrames"`
		TotalFrames int `json:"totalFrames"`
	}
	if err := json.Unmarshal(resp.Body, &body); err != nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.frames = s.frames[:0]
	s.frameDap = s.frameDap[:0]
	for i, f := range body.StackFrames {
		file := f.Source.Name
		if file == "" && f.Source.Path != "" {
			file = filepath.Base(f.Source.Path)
		}
		s.frames = append(s.frames, Frame{Idx: i, Func: f.Name, File: file, Line: f.Line})
		s.frameDap = append(s.frameDap, f.ID)
	}
	if len(s.frames) > 0 {
		top := s.frames[0]
		s.cur = StopInfo{Func: top.Func, File: top.File, Line: top.Line}
	}
}

// ---------- 断点 ----------

// dapSourcePath 适配器视角的当前源码路径(42m 目录)
func (s *DAPSession) dapSourcePath() string {
	s.mu.Lock()
	cur := s.cur
	s.mu.Unlock()
	if cur.File != "" {
		if len(s.frames) > 0 {
			// stackTrace 的 source.path 在刷新时未保存,从 frames 取不到 path;
			// 用模块目录推 42m 路径
			_ = cur
		}
	}
	dvm := s.runProg
	if dvm == "" {
		dvm = s.prog
	}
	if s.module != "" {
		dvm = s.module + "_" + dvm
	}
	return s.cfg.ModuleDir(s.module) + "/42m/" + dvm + ".4gl"
}

// setBreakpointsAt 对指定源码整批下发断点(DAP 语义:同源码全量替换),
// 返回本地断点账本
func (s *DAPSession) setBreakpointsAt(sourcePath string, lines []int) []*Breakpoint {
	src := map[string]any{"name": filepath.Base(sourcePath), "path": sourcePath}
	bpArgs := make([]map[string]any, 0, len(lines))
	for _, ln := range lines {
		bpArgs = append(bpArgs, map[string]any{"line": ln})
	}
	args := map[string]any{"source": src, "breakpoints": bpArgs}
	if len(lines) > 0 {
		args["lines"] = lines
	}
	resp, err := s.dap.request("setBreakpoints", args, 20*time.Second)
	if err != nil {
		s.emitEvent(Event{Type: "log", Text: "下断点失败: " + err.Error()})
		return nil
	}
	var body struct {
		Breakpoints []struct {
			ID       int    `json:"id"`
			Line     int    `json:"line"`
			Verified bool   `json:"verified"`
			Message  string `json:"message"`
		} `json:"breakpoints"`
	}
	_ = json.Unmarshal(resp.Body, &body)
	s.mu.Lock()
	defer s.mu.Unlock()
	newBPs := make([]*Breakpoint, 0, len(body.Breakpoints))
	for _, b := range body.Breakpoints {
		if !b.Verified {
			s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("断点 %d 未生效:%s", b.Line, b.Message)})
			continue
		}
		newBPs = append(newBPs, &Breakpoint{Num: b.ID, File: filepath.Base(sourcePath), Line: b.Line, Enabled: true})
	}
	s.bps = newBPs
	out := make([]*Breakpoint, len(s.bps))
	for i, b := range s.bps {
		cp := *b
		out[i] = &cp
	}
	return out
}

func (s *DAPSession) Breakpoints() []Breakpoint {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Breakpoint, 0, len(s.bps))
	for _, b := range s.bps {
		out = append(out, *b)
	}
	return out
}

// Break 下断点:支持行号 / file:line / 函数名(函数名走函数断点能力)
func (s *DAPSession) Break(loc string) (*Breakpoint, error) {
	loc = strings.TrimSpace(loc)
	if loc == "" {
		return nil, fmt.Errorf("断点位置为空")
	}
	// 现有断点(同位置)直接返回
	for _, b := range s.Breakpoints() {
		if fmt.Sprintf("%s:%d", b.File, b.Line) == loc || fmt.Sprintf("%d", b.Line) == loc {
			cp := b
			cp.Note = fmt.Sprintf("该位置已有断点 #%d(%s:%d)", b.Num, b.File, b.Line)
			return &cp, nil
		}
	}
	lines := s.bpLines()
	if n := lineOnly(loc); n > 0 {
		lines = append(lines, n)
	} else if strings.Contains(loc, ":") {
		parts := strings.SplitN(loc, ":", 2)
		if n := lineOnly(parts[1]); n > 0 {
			lines = append(lines, n)
		}
	} else {
		return nil, fmt.Errorf("DAP 模式暂不支持函数名断点,请使用行号")
	}
	bps := s.setBreakpointsAt(s.dapSourcePath(), lines)
	if len(bps) > 0 {
		cp := *bps[len(bps)-1]
		return &cp, nil
	}
	return nil, fmt.Errorf("断点未生效(位置无效?)")
}

func (s *DAPSession) bpLines() []int {
	var out []int
	for _, b := range s.Breakpoints() {
		if b.Enabled {
			out = append(out, b.Line)
		}
	}
	return out
}

func lineOnly(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}

func (s *DAPSession) DeleteBreakpoint(num int) error {
	var kept []int
	found := false
	for _, b := range s.Breakpoints() {
		if b.Num == num {
			found = true
			continue
		}
		if b.Enabled {
			kept = append(kept, b.Line)
		}
	}
	if !found {
		return fmt.Errorf("断点 #%d 不存在", num)
	}
	s.setBreakpointsAt(s.dapSourcePath(), kept)
	return nil
}

func (s *DAPSession) SetBPEnabled(num int, enabled bool) error {
	var lines []int
	found := false
	for _, b := range s.Breakpoints() {
		if b.Num == num {
			found = true
			b.Enabled = enabled
		}
		if b.Enabled {
			lines = append(lines, b.Line)
		}
	}
	if !found {
		return fmt.Errorf("断点 #%d 不存在", num)
	}
	s.setBreakpointsAt(s.dapSourcePath(), lines)
	return nil
}

// ---------- 求值 / 上下文 ----------

// stripQuotedDisplay 剥掉适配器对 STRING 型结果的展示引号("..."),
// 让截断特征("..." 结尾)在 print/悬停/监视/变量树之间全局一致
func stripQuotedDisplay(v string) string {
	if len(v) >= 2 && strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) {
		return v[1 : len(v)-1]
	}
	return v
}

func (s *DAPSession) Print(expr string) (string, error) {
	v, err := s.evaluate(expr)
	if err != nil {
		return "", err
	}
	return stripQuotedDisplay(v), nil
}

func (s *DAPSession) evaluate(expr string) (string, error) {
	fid := s.currentDapFrame()
	args := map[string]any{"expression": expr, "context": "watch"}
	if fid > 0 {
		args["frameId"] = fid
	}
	resp, err := s.dap.request("evaluate", args, 20*time.Second)
	if err != nil {
		return "", err
	}
	var body struct {
		Result  string `json:"result"`
		Type    string `json:"type"`
		Ref     int    `json:"variablesReference"`
		Named   int    `json:"namedVariables"`
		Indexed int    `json:"indexedVariables"`
	}
	_ = json.Unmarshal(resp.Body, &body)
	out := body.Result
	if body.Indexed > 0 || body.Named > 0 {
		out = fmt.Sprintf("%s (%s, %d 项)", body.Result, body.Type, body.Indexed+body.Named)
	}
	return out, nil
}

func (s *DAPSession) currentDapFrame() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.curFrame < len(s.frameDap) {
		return s.frameDap[s.curFrame]
	}
	return s.stopFrame
}

func (s *DAPSession) Where() ([]Frame, error) {
	s.mu.Lock()
	st := s.state
	frames := append([]Frame(nil), s.frames...)
	s.mu.Unlock()
	if st != StateStopped {
		return nil, fmt.Errorf("仅停站可查看调用栈")
	}
	return frames, nil
}

// refreshScopes 取当前帧的 Locals/Globals variablesReference
func (s *DAPSession) refreshScopes() error {
	fid := s.currentDapFrame()
	resp, err := s.dap.request("scopes", map[string]any{"frameId": fid}, 15*time.Second)
	if err != nil {
		return err
	}
	var body struct {
		Scopes []struct {
			Name              string `json:"name"`
			VariablesReference int   `json:"variablesReference"`
		} `json:"scopes"`
	}
	if err := json.Unmarshal(resp.Body, &body); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sc := range body.Scopes {
		switch sc.Name {
		case "Locals":
			s.localsRef = sc.VariablesReference
		case "Globals":
			s.globalsRef = sc.VariablesReference
		}
	}
	return nil
}

func (s *DAPSession) dapVariables(ref, count int) ([]dapVar, error) {
	args := map[string]any{"variablesReference": ref}
	if count > 0 {
		args["count"] = count
	}
	resp, err := s.dap.request("variables", args, 30*time.Second)
	if err != nil {
		return nil, err
	}
	var body struct {
		Variables []dapVar `json:"variables"`
	}
	_ = json.Unmarshal(resp.Body, &body)
	for i := range body.Variables {
		body.Variables[i].Value = stripQuotedDisplay(body.Variables[i].Value)
	}
	return body.Variables, nil
}

// VarNode 变量树节点(Ref>0 表示可继续下钻)
type VarNode struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
	Ref   int    `json:"ref,omitempty"`
}

func (s *DAPSession) Mode() string { return "dap" }

// VarRoots 返回局部/全局两个 scope 的变量引用(停站后有效;跨停站失效,需重取)
func (s *DAPSession) VarRoots() (int, int, error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return 0, 0, fmt.Errorf("仅停站可查看变量")
	}
	lref, gref := s.localsRef, s.globalsRef
	s.mu.Unlock()
	if lref < 0 || gref < 0 {
		if err := s.refreshScopes(); err != nil {
			return 0, 0, err
		}
		s.mu.Lock()
		lref, gref = s.localsRef, s.globalsRef
		s.mu.Unlock()
	}
	return lref, gref, nil
}

// VarChildren 下钻一层变量(ref 引用仅在当前停站有效,陈旧引用由适配器报错)
func (s *DAPSession) VarChildren(ref int) ([]VarNode, error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return nil, fmt.Errorf("仅停站可查看变量")
	}
	s.mu.Unlock()
	vars, err := s.dapVariables(ref, 500)
	if err != nil {
		return nil, err
	}
	out := make([]VarNode, 0, len(vars))
	for _, v := range vars {
		out = append(out, VarNode{Name: v.Name, Value: v.Value, Type: v.Type, Ref: v.VariablesReference})
	}
	return out, nil
}

// EvalRef 求值并附带适配器的类型与变量引用(ref>0 可经 /variables 下钻,供监视树展开)。
// 引用跨停站失效,前端每次停站后重建。返回值与 Print 同样剥展示引号。
func (s *DAPSession) EvalRef(expr string) (value string, typeStr string, ref int, err error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return "", "", 0, fmt.Errorf("仅停站可求值")
	}
	s.mu.Unlock()
	fid := s.currentDapFrame()
	args := map[string]any{"expression": strings.TrimSpace(expr), "context": "watch"}
	if fid > 0 {
		args["frameId"] = fid
	}
	resp, err := s.dap.request("evaluate", args, 20*time.Second)
	if err != nil {
		return "", "", 0, err
	}
	var body struct {
		Result             string `json:"result"`
		Type               string `json:"type"`
		VariablesReference int    `json:"variablesReference"`
	}
	_ = json.Unmarshal(resp.Body, &body)
	return stripQuotedDisplay(body.Result), body.Type, body.VariablesReference, nil
}

// simpleChainRe 简单变量链(g_xxx / g_qryparam.cond):完整值分段求值只对这类表达式可用
// (4GL 子串下标只能作用在变量上,任意表达式的结果没有载体)
var simpleChainRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*(\.[A-Za-z0-9_]+)*$`)

// FullValue 取被适配器截断(250 字符 + "...")的字符串完整值:
// 适配器写死截断且不提供字符串下钻引用,靠表达式分段绕:
//   先 getLength() 拿精确长度(STRING 类方法,CHAR/VARCHAR 也可用),
//   CHAR/VARCHAR 用 4GL 子串下标 expr[起,止] 分段;STRING 下标报 "not an array",改 String 类方法 expr.subString(起,止)。
// 分段端点按长度夹紧——越界的 subString 返回空而非截断,不能靠探测收尾。每段 200 字符,最多 100 段防失控。
func (s *DAPSession) FullValue(expr string) (string, error) {
	expr = strings.TrimSpace(expr)
	if !simpleChainRe.MatchString(expr) {
		return "", fmt.Errorf("完整值仅支持简单变量链(如 ls_sql / g_qryparam.cond),任意表达式无子串载体")
	}
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return "", fmt.Errorf("仅停站可求值")
	}
	s.mu.Unlock()
	// 注意:currentDapFrame 内部也拿 s.mu,必须在解锁后调用(否则自死锁,整个会话卡死)
	fid := s.currentDapFrame()

	evalOne := func(e string) (string, error) {
		args := map[string]any{"expression": e, "context": "watch"}
		if fid > 0 {
			args["frameId"] = fid
		}
		resp, err := s.dap.request("evaluate", args, 15*time.Second)
		if err != nil {
			return "", err
		}
		var body struct {
			Result string `json:"result"`
		}
		_ = json.Unmarshal(resp.Body, &body)
		return body.Result, nil
	}
	// STRING 求值结果带引号展示,剥掉;CHAR 型不带
	stripQuote := func(r string) string {
		if len(r) >= 2 && strings.HasPrefix(r, `"`) && strings.HasSuffix(r, `"`) {
			return r[1 : len(r)-1]
		}
		return r
	}

	// 精确长度(getLength 是 String 类方法,CHAR/VARCHAR 也响应)
	lenStr, err := evalOne(fmt.Sprintf("%s.getLength()", expr))
	if err != nil {
		return "", fmt.Errorf("getLength 失败: %w", err)
	}
	total := 0
	if n, convErr := fmt.Sscanf(strings.TrimSpace(lenStr), "%d", &total); convErr != nil || n != 1 || total < 0 {
		return "", fmt.Errorf("getLength 返回异常: %s", lenStr)
	}
	if total == 0 {
		return "", fmt.Errorf("变量为空")
	}

	// 分段模式:第一段先试子串下标,STRING 型会报 not an array,自动切 subString
	mode := "idx"
	const chunk = 200
	var sb strings.Builder
	for start := 1; start <= total && start <= 100*chunk; start += chunk {
		end := start + chunk - 1
		if end > total {
			end = total
		}
		var e string
		if mode == "sub" {
			e = fmt.Sprintf("%s.subString(%d,%d)", expr, start, end)
		} else {
			e = fmt.Sprintf("%s[%d,%d]", expr, start, end)
		}
		r, err := evalOne(e)
		if err != nil {
			if mode == "idx" && strings.Contains(err.Error(), "not an array") {
				mode = "sub"
				start -= chunk // 本轮 for 会 +chunk,重取第一段
				continue
			}
			return "", fmt.Errorf("分段 %d-%d 失败: %w", start, end, err)
		}
		sb.WriteString(stripQuote(r))
	}
	out := sb.String()
	if len([]rune(out)) < total {
		return out, fmt.Errorf("分段不完整:取到 %d / %d 字符", len([]rune(out)), total)
	}
	return out, nil
}

func (s *DAPSession) Locals() ([]VarItem, error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return nil, fmt.Errorf("仅停站可查看变量")
	}
	ref := s.localsRef
	s.mu.Unlock()
	if ref < 0 {
		if err := s.refreshScopes(); err != nil {
			return nil, err
		}
		s.mu.Lock()
		ref = s.localsRef
		s.mu.Unlock()
	}
	if ref < 0 {
		return nil, nil
	}
	vars, err := s.dapVariables(ref, 0)
	if err != nil {
		return nil, err
	}
	out := make([]VarItem, 0, len(vars))
	for _, v := range vars {
		out = append(out, VarItem{Expr: v.Name, Value: v.Value})
	}
	s.mu.Lock()
	s.autovars = out
	s.mu.Unlock()
	return out, nil
}

func (s *DAPSession) Globals(limit int) ([]VarDecl, int, error) {
	s.mu.Lock()
	if s.state != StateStopped {
		s.mu.Unlock()
		return nil, 0, fmt.Errorf("仅停站可查看变量")
	}
	ref := s.globalsRef
	s.mu.Unlock()
	if ref < 0 {
		if err := s.refreshScopes(); err != nil {
			return nil, 0, err
		}
		s.mu.Lock()
		ref = s.globalsRef
		s.mu.Unlock()
	}
	if ref < 0 {
		return nil, 0, nil
	}
	vars, err := s.dapVariables(ref, limit)
	if err != nil {
		return nil, 0, err
	}
	out := make([]VarDecl, 0, len(vars))
	for _, v := range vars {
		out = append(out, VarDecl{Name: v.Name, Type: v.Type})
	}
	return out, len(out), nil
}

func (s *DAPSession) Autovars() []VarItem {
	s.mu.Lock()
	av := s.autovars
	s.mu.Unlock()
	if av != nil {
		return av
	}
	// 懒加载:首次访问取一次 Locals
	if _, err := s.Locals(); err != nil {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.autovars
}

func (s *DAPSession) Frame(n int) error {
	s.mu.Lock()
	if n < 0 || n >= len(s.frames) {
		s.mu.Unlock()
		return fmt.Errorf("栈帧 #%d 不存在(共 %d 帧)", n, len(s.frames))
	}
	s.curFrame = n
	fr := s.frames[n]
	s.mu.Unlock()
	s.emitEvent(Event{Type: "log", Text: fmt.Sprintf("切换到栈帧 #%d %s:%d", n, fr.Func, fr.Line)})
	return nil
}

func (s *DAPSession) Sources() ([]string, error) {
	return nil, fmt.Errorf("DAP 模式暂不支持源码列表")
}

func (s *DAPSession) Functions(limit int) ([]string, int, error) {
	return nil, 0, fmt.Errorf("DAP 模式暂不支持函数列表")
}

func (s *DAPSession) InfoLine(loc string) (string, int, error) {
	return "", 0, fmt.Errorf("DAP 模式暂不支持 info line")
}

func (s *DAPSession) Raw(cmd string, timeout time.Duration) ([]string, error) {
	return nil, fmt.Errorf("DAP 模式不支持原始 fgldb 命令(协议直通不可用)")
}

// ---------- 源码读取(SFTP,与 PTY 版同语义) ----------

func (s *DAPSession) sftpClient() (*sftp.Client, error) {
	s.sftpMu.Lock()
	defer s.sftpMu.Unlock()
	if s.sftpCl != nil {
		return s.sftpCl, nil
	}
	cl, err := s.conn.SFTP()
	if err != nil {
		return nil, err
	}
	s.sftpCl = cl
	return cl, nil
}

func (s *DAPSession) ReadPath(path string) (*SourceFile, error) {
	s.mu.Lock()
	cached, ok := s.srcCache[path]
	s.mu.Unlock()
	if ok {
		return &SourceFile{DVMFile: filepath.Base(path), Path: path, Content: cached.file.Content, ModTime: cached.modTime}, nil
	}
	cl, err := s.sftpClient()
	if err != nil {
		return nil, err
	}
	f, err := cl.Open(path)
	if err != nil {
		return nil, err
	}
	var mt time.Time
	if fi, err := f.Stat(); err == nil {
		mt = fi.ModTime()
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		return nil, err
	}
	sf := &SourceFile{DVMFile: filepath.Base(path), Path: path, Content: string(data), ModTime: mt}
	s.mu.Lock()
	s.srcCache[path] = srcCacheEntry{file: *sf, modTime: mt}
	s.mu.Unlock()
	return sf, nil
}

func (s *DAPSession) ResolveSource(dvmFile, module string) (*SourceFile, error) {
	s.mu.Lock()
	cached, ok := s.srcCache[dvmFile]
	s.mu.Unlock()
	if ok {
		return &SourceFile{DVMFile: dvmFile, Path: cached.file.Path, Content: cached.file.Content, ModTime: cached.modTime}, nil
	}
	if module == "" {
		module = s.module
	}
	cl, err := s.sftpClient()
	if err != nil {
		return nil, err
	}
	var lastErr error
	for _, p := range sourceCandidatePaths(s.cfg.ModuleRoots, module, dvmFile) {
		f, err := cl.Open(p)
		if err != nil {
			lastErr = err
			continue
		}
		var mt time.Time
		if fi, err := f.Stat(); err == nil {
			mt = fi.ModTime()
		}
		data, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			lastErr = err
			continue
		}
		s.mu.Lock()
		s.srcCache[dvmFile] = srcCacheEntry{file: SourceFile{DVMFile: dvmFile, Path: p, Content: string(data)}, modTime: mt}
		s.mu.Unlock()
		return &SourceFile{DVMFile: dvmFile, Path: p, Content: string(data), ModTime: mt}, nil
	}
	return nil, fmt.Errorf("源码未找到(%s): %w", dvmFile, lastErr)
}
