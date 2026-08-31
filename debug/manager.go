package debug

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Manager 调试会话管理器:持有会话、向订阅者(WS/MCP)广播事件
type Manager struct {
	cfg *Config

	mu       sync.Mutex
	sessions map[string]*Session

	subMu sync.Mutex
	subs  map[chan Event]string
}

func NewManager(cfg *Config) *Manager {
	return &Manager{
		cfg:      cfg,
		sessions: map[string]*Session{},
		subs:     map[chan Event]string{},
	}
}

// Subscribe 订阅事件流(WS 用),返回取消函数
func (m *Manager) Subscribe(tag string) (<-chan Event, func()) {
	ch := make(chan Event, 4096)
	m.subMu.Lock()
	m.subs[ch] = tag
	m.subMu.Unlock()
	return ch, func() {
		m.subMu.Lock()
		delete(m.subs, ch)
		m.subMu.Unlock()
	}
}

// emit 会话事件回调:广播给所有订阅者(非阻塞,满则丢弃)
func (m *Manager) emit(ev Event) {
	m.subMu.Lock()
	defer m.subMu.Unlock()
	for ch := range m.subs {
		select {
		case ch <- ev:
		default: // 订阅者消费太慢,丢弃(前端靠 state/stopped 事件对齐,丢原始行无碍)
		}
	}
}

// Launch 创建并启动一个调试会话(同一时间仅允许一个会话,保证生产安全)。
// 启动前清扫已结束的残留会话(程序退出/后端死亡),重启无需先 DELETE。
func (m *Manager) Launch(module, prog string) (*Session, error) {
	return m.launchWith(m.cfg, module, prog)
}

// LaunchWith 按指定配置启动(多 SSH/多区域:调用方克隆好 cfg 再传入)
func (m *Manager) LaunchWith(cfg *Config, module, prog string) (*Session, error) {
	return m.launchWith(cfg, module, prog)
}

func (m *Manager) launchWith(cfg *Config, module, prog string) (*Session, error) {
	m.mu.Lock()
	for id, s := range m.sessions {
		if s.State() == StateExit {
			delete(m.sessions, id)
		}
	}
	if len(m.sessions) > 0 {
		m.mu.Unlock()
		return nil, fmt.Errorf("已存在调试会话,请先结束当前会话(M4 将支持多会话)")
	}
	m.mu.Unlock()

	// 对齐 T100 gendbg:启动前连库把作业编号解析成实体程序+模块+启动引用(gzzz_t JOIN gzza_t),
	// 源码与 42r 都跟实体程序走;未配置 db/查询失败/未命中 → 会话内按名称文件搜索兜底
	runProg, launchRef, extra := "", "", ""
	if mod2, prog2, ref2, extra2 := m.resolveJobWith(cfg, module, prog); prog2 != "" {
		runProg = prog2
		launchRef, extra = ref2, extra2
		if module == "" && mod2 != "" {
			module = mod2
		}
	}

	sess, err := NewSession(cfg, module, prog, runProg, launchRef, extra, m.emit)
	if err != nil {
		return nil, err
	}
	if runProg != "" && runProg != prog {
		sess.emitEvent(Event{Type: "log", Text: fmt.Sprintf("作业编号 %s → 实体程序 %s(gzzz_t)", prog, runProg)})
	}
	m.mu.Lock()
	m.sessions[sess.ID] = sess
	m.mu.Unlock()
	return sess, nil
}

// resolveJob 连库按 gendbg 语义解析作业编号(gzzz_t:gzzz001 → gzzz002 实体程序 + gzzz005 模块)。
// 返回 (模块, 实体程序);未配置 db/查询失败/未命中 → ("",""),由调用方回退文件搜索。
func (m *Manager) resolveJob(module, job string) (mod, prog, launchRef, extra string) {
	return m.resolveJobWith(m.cfg, module, job)
}

func (m *Manager) resolveJobWith(cfg *Config, module, job string) (mod, prog, launchRef, extra string) {
	if cfg.DB == nil || !reProgName.MatchString(job) {
		return "", "", "", ""
	}
	conn, err := Dial(cfg.SSH)
	if err != nil {
		return "", "", "", ""
	}
	defer conn.Close()
	zone := cfg.Zone
	if zone == "" {
		zone = "36"
	}
	// 金仓:探测实例要素后连库解析;Oracle:直接用 TNS
	var kb *kbCtx
	tns := cfg.TNSName()
	if cfg.DBType() == "kingbase" {
		env, err := probeKBEnv(conn)
		if err != nil {
			return "", "", "", ""
		}
		kb = &kbCtx{ksql: env["KSQL"], port: env["KPORT"], db: env["KDB"]}
		if cfg.DB != nil && cfg.DB.TNS != "" {
			kb.db = cfg.DB.TNS
		}
		if cfg.DB != nil && cfg.DB.Port > 0 {
			kb.port = strconv.Itoa(cfg.DB.Port)
		}
		tns = kb.db
	}
	jr, err := dbResolveJob(conn, zone, tns, job, kb)
	if err != nil || jr.Prog == "" {
		return "", "", "", ""
	}
	p := jr.Prog
	// gendbg 原版语义:gzza004 的 $变量 由选区后 shell 展开为权威 42r 路径(标准/客制都覆盖)
	if !strings.HasPrefix(jr.LaunchRef, "$") {
		jr.LaunchRef, jr.Extra = "", ""
	}
	launchRef, extra = jr.LaunchRef, jr.Extra
	if module != "" {
		return module, p, launchRef, extra // 用户显式指定模块:尊重指定,仅采纳实体程序
	}
	mod = strings.ToLower(jr.Module)
	if jr.Module == "" || jr.Module == "-" || !modHas42r(conn, cfg.TopDir, mod, p) {
		// 模块码缺失或目录对不上:全模块搜索实体程序的 42r;0/多命中留给会话内兜底报错
		if cands := searchModule42r(conn, cfg.ModuleRoots, p); len(cands) == 1 {
			mod = cands[0]
		} else {
			mod = ""
		}
	}
	return mod, p, launchRef, extra
}

// LaunchReplay 接口日志重放调试:等价 T100 日志内嵌的 `r.dg <作业> '<req>' '<rsp>'`。
// 作业取 wsfa012(gzja_t 服务→程序的解析结果),再走 gzzz_t 解析实体程序+模块;
// 报文文件被清理时用 CLOB 内容落到服务器临时文件再重放。
func (m *Manager) LaunchReplay(item *WSLogItem, content *WSLogContent) (*Session, error) {
	m.mu.Lock()
	for id, s := range m.sessions {
		if s.State() == StateExit {
			delete(m.sessions, id)
		}
	}
	if len(m.sessions) > 0 {
		m.mu.Unlock()
		return nil, fmt.Errorf("已存在调试会话,请先结束当前会话")
	}
	m.mu.Unlock()

	job := strings.TrimSpace(item.Job)
	if job == "" {
		return nil, fmt.Errorf("该日志没有关联作业编号(wsfa012 为空),无法重放")
	}
	module, runProg, launchRef, extra := "", "", "", ""
	if mod2, prog2, ref2, extra2 := m.resolveJob("", job); prog2 != "" {
		module, runProg, launchRef, extra = mod2, prog2, ref2, extra2
	}
	conn, err := Dial(m.cfg.SSH)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %w", err)
	}
	defer conn.Close()
	reqPath, rspPath, err := WriteReplayFiles(conn, item, content)
	if err != nil {
		return nil, err
	}
	// 路径进 shell 单引号,剔除单引号防注入。
	// 对齐 r.dg:报文文件追加在标准参数(BBDL 会话标记等)之后,
	// cl_wss_init 依 BBDL 标记识别调试模式并重建 g_argv;只传报文会导致 g_argv 为空 → -1
	q := func(p string) string { return "'" + strings.ReplaceAll(p, "'", "") + "'" }
	args := strings.ReplaceAll(m.cfg.LaunchArgs, "{prog}", job)
	args += " " + q(reqPath)
	if rspPath != "" {
		args += " " + q(rspPath)
	}
	sess, err := NewSession(m.cfg, module, job, runProg, launchRef, extra, m.emit)
	if err != nil {
		return nil, err
	}
	sess.ArgsOverride = args
	if runProg != "" && runProg != job {
		sess.emitEvent(Event{Type: "log", Text: fmt.Sprintf("重放调试:作业 %s → 实体程序 %s(gzzz_t)", job, runProg)})
	}
	sess.emitEvent(Event{Type: "log", Text: fmt.Sprintf("重放 %s:fglrun -d %s %s", item.Service, runProgOr(job, runProg), args)})
	m.mu.Lock()
	m.sessions[sess.ID] = sess
	m.mu.Unlock()
	return sess, nil
}

func runProgOr(job, runProg string) string {
	if runProg != "" {
		return runProg
	}
	return job
}

// Get 取会话
func (m *Manager) Get(id string) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.sessions[id]
}

// Current 返回当前唯一会话(无则 nil)
func (m *Manager) Current() *Session {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.sessions {
		return s
	}
	return nil
}

// Snapshot 会话列表快照
type SessionBrief struct {
	ID       string  `json:"id"`
	Module   string  `json:"module"`
	Prog     string  `json:"prog"`
	RunProg  string  `json:"runProg,omitempty"`
	State    string  `json:"state"`
	Started  bool    `json:"started"`
	File     string  `json:"file,omitempty"`
	Line     int     `json:"line,omitempty"`
	Func     string  `json:"func,omitempty"`
	Reason   string  `json:"reason,omitempty"`
	Holding  float64 `json:"holdingSeconds"`
	Breaks   int     `json:"breakpoints"`
	Watchdog int     `json:"watchdogSeconds"`
}

func (m *Manager) Snapshot() []SessionBrief {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]SessionBrief, 0, len(m.sessions))
	for _, s := range m.sessions {
		cur := s.Cur()
		out = append(out, SessionBrief{
			ID: s.ID, Module: s.Module, Prog: s.Prog, RunProg: s.RunProg,
			State: string(s.State()), Started: s.Started(), File: cur.File, Line: cur.Line,
			Func: cur.Func, Reason: cur.Reason,
			Holding:  s.HoldingSeconds(),
			Breaks:   len(s.Breakpoints()),
			Watchdog: m.cfg.WatchdogSeconds,
		})
	}
	return out
}

// Remove 移除已结束的会话
func (m *Manager) Remove(id string) {
	m.mu.Lock()
	delete(m.sessions, id)
	m.mu.Unlock()
}

// SourcePreview 会话建立前预取源码:用独立短连接读母版,消除前端启动期空白。
// module 为空(留待会话内自动解析)时直接跳过。
func (m *Manager) SourcePreview(module, prog string) (*SourceFile, error) {
	if module == "" {
		return nil, fmt.Errorf("module 为空,跳过源码预取")
	}
	conn, err := Dial(m.cfg.SSH)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %w", err)
	}
	defer conn.Close()
	cl, err := conn.SFTP()
	if err != nil {
		return nil, err
	}
	dvmFile := module + "_" + prog + ".4gl"
	var lastErr error
	for _, p := range sourceCandidatePaths(m.cfg.ModuleRoots, module, dvmFile) {
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
		return &SourceFile{DVMFile: dvmFile, Path: p, Content: string(data), ModTime: mt}, nil
	}
	return nil, fmt.Errorf("源码未找到(%s): %w", dvmFile, lastErr)
}

// CloseAll 结束全部会话(服务退出时调用)
func (m *Manager) CloseAll() {
	m.mu.Lock()
	ss := make([]*Session, 0, len(m.sessions))
	for _, s := range m.sessions {
		ss = append(ss, s)
	}
	m.sessions = map[string]*Session{}
	m.mu.Unlock()
	for _, s := range ss {
		_ = s.Quit()
	}
}
