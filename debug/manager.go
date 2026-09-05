package debug

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

// Manager 调试会话管理器:持有会话、向订阅者(WS/MCP)广播事件、保留最近事件供查询
type Manager struct {
	cfg *Config

	mu       sync.Mutex
	sessions map[string]*Session

	subMu sync.Mutex
	subs  map[chan Event]string

	// 最近事件环形缓冲(供 AI/日志 tail 查询;只存结构化事件,output/autovars 不入)
	evMu sync.Mutex
	evs  []Event

	// T100 动态环境缓存(登录区域 → 路径),TTL 5 分钟;探针失败不缓存
	envMu    sync.Mutex
	envCache map[string]cachedEnv
}

// maxBufferedEvents 环形缓冲上限。
const maxBufferedEvents = 4000

type cachedEnv struct {
	env *RuntimeEnv
	at  time.Time
}

func NewManager(cfg *Config) *Manager {
	return &Manager{
		cfg:      cfg,
		sessions: map[string]*Session{},
		subs:     map[chan Event]string{},
		envCache: map[string]cachedEnv{},
	}
}

// getRuntimeEnv 取 T100 动态环境:缓存命中直接返回;未命中(或过期)探针一次并缓存。
// 失败返回 nil,调用方回退静态配置。key 按 服务器+账号+区域 区分(用连接的真实 host,
// 覆盖/多环境时不受 active 环境影响),换环境互不污染。
func (m *Manager) getRuntimeEnv(conn *SSHConn, zone string) *RuntimeEnv {
	key := conn.cfg.Host + "|" + conn.cfg.User + "|" + zone
	m.envMu.Lock()
	if c, ok := m.envCache[key]; ok && time.Since(c.at) < 5*time.Minute && c.env.valid() {
		m.envMu.Unlock()
		return c.env
	}
	m.envMu.Unlock()
	env, err := probeTEnv(conn, zone)
	if err != nil {
		log.Printf("[tenv] probe zone=%s failed: %v", zone, err)
		return nil
	}
	m.envMu.Lock()
	m.envCache[key] = cachedEnv{env: env, at: time.Now()}
	m.envMu.Unlock()
	return env
}

// ensureRuntimeEnv 确保 cfg.Runtime 有动态路径(探针+缓存);失败静默,用静态配置兜底
func (m *Manager) ensureRuntimeEnv(conn *SSHConn) {
	if m.cfg.Runtime != nil && m.cfg.Runtime.valid() {
		return
	}
	zone := m.cfg.Zone
	if zone == "" {
		zone = "36"
	}
	if env := m.getRuntimeEnv(conn, zone); env != nil {
		m.cfg.Runtime = env
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

// emit 会话事件回调:入环形缓冲并广播给所有订阅者(非阻塞,满则丢弃)
func (m *Manager) emit(ev Event) {
	// 只保留对"发生了什么"有信息量的结构化事件(输出流/自动变量列表不入缓冲,
	// 会占用大量内存且 AI 用 tail 查询时噪声过大)
	switch ev.Type {
	case "state", "stopped", "log", "dead", "watchdog", "ai_action":
		m.evMu.Lock()
		m.evs = append(m.evs, ev)
		if len(m.evs) > maxBufferedEvents {
			m.evs = m.evs[len(m.evs)-maxBufferedEvents:]
		}
		m.evMu.Unlock()
	}
	m.subMu.Lock()
	defer m.subMu.Unlock()
	for ch := range m.subs {
		select {
		case ch <- ev:
		default: // 订阅者消费太慢,丢弃(前端靠 state/stopped 事件对齐,丢原始行无碍)
		}
	}
}

// Events 返回最近 n 条事件(按时间升序;n<=0 返回全部缓冲)。
func (m *Manager) Events(n int) []Event {
	m.evMu.Lock()
	defer m.evMu.Unlock()
	if n <= 0 || n >= len(m.evs) {
		out := make([]Event, len(m.evs))
		copy(out, m.evs)
		return out
	}
	out := make([]Event, n)
	copy(out, m.evs[len(m.evs)-n:])
	return out
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
	sess, err := m.prepareSession(cfg, module, prog, runProg, launchRef, extra, "")
	if err != nil {
		return nil, err
	}
	if runProg != "" && runProg != prog {
		sess.emitEvent(Event{Type: "log", Text: fmt.Sprintf("作业编号 %s → 实体程序 %s(gzzz_t)", prog, runProg)})
	}
	return sess, nil
}

// prepareSession 启动一轮运行前的会话准备(单一常驻会话):
// - 已有空闲(idle)且目标相同的会话 → 复用宿主,仅刷新本轮运行参数(免重新登录);
// - 已有会话但目标不同(默认环境切换/按名启动其它服务器)→ 断开旧的再新建;
// - 已有活跃运行且目标相同 → 报错(请先结束当前调试);
// - 无会话 → 新建。
// 返回已登记的会话,由调用方驱动 Launch 登录/启动。
func (m *Manager) prepareSession(cfg *Config, module, prog, runProg, launchRef, extra, argsOverride string) (*Session, error) {
	m.mu.Lock()
	for id, s := range m.sessions {
		if s.State() == StateExit {
			delete(m.sessions, id) // 已断开(close/切换/EOF)的残留,资源已释放
		}
	}
	var live *Session
	for _, s := range m.sessions {
		live = s
		break
	}
	m.mu.Unlock()

	if live != nil {
		if live.State() == StateIdle && live.sameTarget(cfg) {
			live.SetRun(cfg, module, prog, runProg, launchRef, extra)
			live.ArgsOverride = argsOverride
			return live, nil
		}
		if !live.sameTarget(cfg) {
			log.Printf("[debug] 目标切换为 %s,断开旧会话 %s 后重建", cfg.EnvName(), live.ID)
			_ = live.Close()
			m.Remove(live.ID)
		} else {
			return nil, fmt.Errorf("已存在调试会话(%s 运行中/已停站),请先结束当前调试", live.Prog)
		}
	}

	sess, err := NewSession(cfg, module, prog, runProg, launchRef, extra, m.emit)
	if err != nil {
		return nil, err
	}
	if argsOverride != "" {
		sess.ArgsOverride = argsOverride
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
	// 动态路径(登录区域 → 环境脚本):按本次启动的区域探针+缓存(覆盖/多环境时不用
	// 全局 active 区域),结果写入 cfg.Runtime——预会话模块解析/42r 校验直接用真实路径。
	// 失败静默,后续用静态配置兜底
	if cfg.Runtime == nil || !cfg.Runtime.valid() {
		if env := m.getRuntimeEnv(conn, zone); env != nil {
			cfg.Runtime = env
		}
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
	if jr.Module == "" || jr.Module == "-" || !modHas42r(conn, cfg.TopDirActual(), mod, p) {
		// 模块码缺失或目录对不上:全模块搜索实体程序的 42r;0/多命中留给会话内兜底报错
		if cands := searchModule42r(conn, cfg.ModuleRootsActual(), p); len(cands) == 1 {
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
// 会话复用规则与普通启动一致:同目标空闲宿主直接复用。
func (m *Manager) LaunchReplay(item *WSLogItem, content *WSLogContent) (*Session, error) {
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
	sess, err := m.prepareSession(m.cfg, module, job, runProg, launchRef, extra, args)
	if err != nil {
		return nil, err
	}
	if runProg != "" && runProg != job {
		sess.emitEvent(Event{Type: "log", Text: fmt.Sprintf("重放调试:作业 %s → 实体程序 %s(gzzz_t)", job, runProg)})
	}
	sess.emitEvent(Event{Type: "log", Text: fmt.Sprintf("重放 %s:fglrun -d %s %s", item.Service, runProgOr(job, runProg), args)})
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
	Env      string  `json:"env,omitempty"` // 会话所属环境名(设置页 envs)
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
			State: string(s.State()), Env: s.EnvName(), Started: s.Started(), File: cur.File, Line: cur.Line,
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

// CreateSessionOn 新建并登记一个纯宿主会话(module/prog 为空,登录后回到 idle),
// 供「会话重启/切换/连接」在指定目标环境上重连;调用方随后驱动 BootIdle。
func (m *Manager) CreateSessionOn(cfg *Config) (*Session, error) {
	m.mu.Lock()
	for id, s := range m.sessions {
		if s.State() == StateExit {
			delete(m.sessions, id)
		}
	}
	m.mu.Unlock()
	sess, err := NewSession(cfg, "", "", "", "", "", m.emit)
	if err != nil {
		return nil, err
	}
	m.mu.Lock()
	m.sessions[sess.ID] = sess
	m.mu.Unlock()
	return sess, nil
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
	// 动态路径(登录区域 → 环境脚本):探针+缓存,失败静默用静态配置兜底
	m.ensureRuntimeEnv(conn)
	cl, err := conn.SFTP()
	if err != nil {
		return nil, err
	}
	dvmFile := module + "_" + prog + ".4gl"
	var lastErr error
	for _, p := range sourceCandidatePaths(m.cfg.ModuleRootsActual(), module, dvmFile) {
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

// ---------- 会话外只读能力(AI/CLI 信息通道;不经会话,独立短连接,白名单限 moduleRoots) ----------

// SourceResult 会话外源码读取结果(带行段裁剪,避免大文件整读灌满 AI 上下文)。
type SourceResult struct {
	SourceFile
	From int `json:"from"` // 1-based 返回行段起点
	To   int `json:"to"`   // 1-based 返回行段终点(含)
	All  int `json:"all"`  // 文件总行数
}

// maxSourceFileBytes 会话外读取的文件大小上限(源码一般远小于此)。
const maxSourceFileBytes = 4 << 20

// ReadSourceStandalone 用独立短连接读取 moduleRoots 内的源码文件。
//   - path 非空:直接按路径读(白名单校验 moduleRoots 前缀);
//   - path 为空:file(如 asf_bsft001_wf.4gl)按 sourceCandidatePaths 候选解析,module 可空(仅查公共目录);
//   - from/to(1-based):返回行段(from<=0 从头,to<=0 到底),行文本超长截断防注入超长行。
//
// 返回裁剪后的内容与文件总行数。
func (m *Manager) ReadSourceStandalone(module, file, path string, from, to int) (*SourceResult, error) {
	if path == "" && file == "" {
		return nil, fmt.Errorf("需要 file 或 path 参数")
	}
	conn, err := Dial(m.cfg.SSH)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %w", err)
	}
	defer conn.Close()
	m.ensureRuntimeEnv(conn)
	cl, err := conn.SFTP()
	if err != nil {
		return nil, fmt.Errorf("打开 SFTP 失败: %w", err)
	}
	roots := m.cfg.ModuleRootsActual()

	readOne := func(p string) ([]byte, time.Time, error) {
		f, err := cl.Open(p)
		if err != nil {
			return nil, time.Time{}, err
		}
		defer f.Close()
		var mt time.Time
		if fi, err := f.Stat(); err == nil {
			mt = fi.ModTime()
		}
		data, err := io.ReadAll(f)
		return data, mt, err
	}

	// 路径白名单:必须在 moduleRoots 之内
	inRoots := func(p string) bool {
		for _, root := range roots {
			if strings.HasPrefix(p, root+"/") {
				return true
			}
		}
		return false
	}

	var (
		data []byte
		mt   time.Time
	)
	if path != "" {
		if !inRoots(path) {
			return nil, fmt.Errorf("路径不在 moduleRoots 内: %s", path)
		}
		data, mt, err = readOne(path)
		if err != nil {
			return nil, err
		}
	} else {
		var lastErr error
		for _, p := range sourceCandidatePaths(roots, module, file) {
			if !inRoots(p) {
				continue
			}
			d, t, e := readOne(p)
			if e == nil {
				data, mt, file = d, t, file
				path = p
				break
			}
			lastErr = e
		}
		if data == nil {
			return nil, fmt.Errorf("源码未找到(%s): %v", file, lastErr)
		}
	}
	if len(data) > maxSourceFileBytes {
		data = data[:maxSourceFileBytes]
	}
	lines := strings.Split(string(data), "\n")
	all := len(lines)
	// 去尾部空行(文件常见尾换行)
	for all > 0 && strings.TrimSpace(lines[all-1]) == "" {
		all--
	}
	if from <= 0 {
		from = 1
	}
	if to <= 0 || to > all {
		to = all
	}
	if from > to {
		from, to = 1, all
	}
	slice := lines[from-1 : to]
	// 单行超长截断(记录行/极长字符串防爆)
	for i, ln := range slice {
		if len(ln) > 4096 {
			slice[i] = ln[:4096] + "…(截断)"
		}
	}
	res := &SourceResult{
		SourceFile: SourceFile{DVMFile: file, Path: path, ModTime: mt},
		From:       from, To: to, All: all,
	}
	if from <= all {
		res.Content = strings.Join(slice, "\n")
	}
	return res, nil
}

// JobInfo 作业编号 → 实体程序/模块解析结果(信息通道,不启动会话)。
type JobInfo struct {
	Prog      string `json:"prog"`                // 用户作业编号
	Module    string `json:"module"`              // 解析出的模块目录(空=未解析)
	RunProg   string `json:"runProg"`             // 实体程序(gzzz_t 解析;空=未解析)
	LaunchRef string `json:"launchRef,omitempty"` //
	Extra     string `json:"extra,omitempty"`     //
	Note      string `json:"note,omitempty"`      // 未解析/降级时的说明
}

// ResolveJob 尽力解析作业:未配置 db 或连库失败/未命中时返回 note,不报错(启动时会再次解析/兜底)。
func (m *Manager) ResolveJob(module, prog string) *JobInfo {
	out := &JobInfo{Prog: prog, Module: module}
	if m.cfg.DB == nil {
		out.Note = "config.json 未配置 debug.db,无法按 gzzz_t 解析实体程序;将按作业名直接启动"
		return out
	}
	if !reProgName.MatchString(prog) {
		out.Note = "作业名非标准标识符,跳过库解析"
		return out
	}
	// resolveJobWith 内部含 连库→gzzz_t→42r 校验,失败时返回空串(不抛)
	mod, run, ref, extra := m.resolveJobWith(m.cfg, module, prog)
	if run == "" {
		out.Note = "未能按 gzzz_t 解析(连库失败或未命中);将按作业名直接启动"
		return out
	}
	if mod != "" {
		out.Module = mod
	}
	out.RunProg, out.LaunchRef, out.Extra = run, ref, extra
	return out
}

// ---------- 会话生命周期 ----------

// CloseAll 断开全部会话(服务退出时调用)
func (m *Manager) CloseAll() {
	m.mu.Lock()
	ss := make([]*Session, 0, len(m.sessions))
	for _, s := range m.sessions {
		ss = append(ss, s)
	}
	m.sessions = map[string]*Session{}
	m.mu.Unlock()
	for _, s := range ss {
		_ = s.Close()
	}
}
