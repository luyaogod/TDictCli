package debug

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/coder/websocket"
)

// Server 本地调试服务:REST + WebSocket + Web 前端
type Server struct {
	cfg        *Config
	cfgPath    string // config.json 路径(设置页写回用)
	web        fs.FS
	webSub     fs.FS // web/dist 子文件系统(未构建前端时为 nil)
	mgr        *Manager
}

// NewServer 创建服务实例;cfgPath 为 config.json 路径(设置页写回用,可为空=只读)
func NewServer(cfg *Config, web fs.FS, cfgPath string) *Server {
	s := &Server{cfg: cfg, cfgPath: cfgPath, web: web, mgr: NewManager(cfg)}
	if web != nil {
		if sub, err := fs.Sub(web, "web/dist"); err == nil {
			if f, err := sub.Open("index.html"); err == nil {
				f.Close()
				s.webSub = sub
			}
		}
	}
	return s
}

// Run 启动 HTTP 服务(阻塞到 ctx 取消)
func (s *Server) Run(ctx context.Context) error {
	mux := http.NewServeMux()
	s.routes(mux)
	srv := &http.Server{Addr: s.cfg.Listen, Handler: mux}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	fmt.Printf("[tdict debug] 服务已启动  前端+API: http://%s\n", s.cfg.Listen)
	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		s.mgr.CloseAll()
		return nil
	}
	return err
}

func (s *Server) routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/status", s.hStatus)
	mux.HandleFunc("POST /api/sessions", s.hLaunch)
	mux.HandleFunc("GET /api/sessions", s.hList)
	mux.HandleFunc("GET /api/sessions/{id}", s.hSnapshot)
	mux.HandleFunc("DELETE /api/sessions/{id}", s.hQuit)
	mux.HandleFunc("GET /api/sessions/{id}/breakpoints", s.hBPList)
	mux.HandleFunc("POST /api/sessions/{id}/breakpoints", s.hBPAdd)
	mux.HandleFunc("DELETE /api/sessions/{id}/breakpoints/{num}", s.hBPDel)
	mux.HandleFunc("POST /api/sessions/{id}/control", s.hControl)
	mux.HandleFunc("POST /api/sessions/{id}/print", s.hPrint)
	mux.HandleFunc("POST /api/sessions/{id}/where", s.hWhere)
	mux.HandleFunc("POST /api/sessions/{id}/raw", s.hRaw)
	mux.HandleFunc("GET /api/sessions/{id}/locals", s.hLocals)
	mux.HandleFunc("GET /api/sessions/{id}/globals", s.hGlobals)
	mux.HandleFunc("GET /api/sessions/{id}/sources", s.hSources)
	mux.HandleFunc("GET /api/sessions/{id}/functions", s.hFunctions)
	mux.HandleFunc("GET /api/sessions/{id}/autovars", s.hAutovars)
	mux.HandleFunc("POST /api/sessions/{id}/frame", s.hFrame)
	mux.HandleFunc("POST /api/sessions/{id}/breakpoints/{num}/enabled", s.hBPEnabled)
	mux.HandleFunc("GET /api/sessions/{id}/source", s.hSource)
	mux.HandleFunc("GET /api/source-preview", s.hSourcePreview)
	mux.HandleFunc("POST /api/wstest", s.hWSTest)
	mux.HandleFunc("GET /api/wslogs", s.hWSLogs)
	mux.HandleFunc("GET /api/wslogs/content", s.hWSLogContent)
	mux.HandleFunc("POST /api/wslogs/debug", s.hWSLogDebug)
	mux.HandleFunc("GET /api/settings", s.hSettingsGet)
	mux.HandleFunc("PUT /api/settings", s.hSettingsPut)
	mux.HandleFunc("POST /api/dbprobe", s.hDBProbe)
	mux.HandleFunc("GET /api/ws", s.hWS)
	mux.HandleFunc("/", s.hStatic)
}

// ---------- 工具 ----------

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, code int, err error) {
	writeJSON(w, code, map[string]any{"ok": false, "error": err.Error()})
}

func readBody[T any](w http.ResponseWriter, r *http.Request, out *T) bool {
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		fail(w, 400, fmt.Errorf("请求体解析失败: %v", err))
		return false
	}
	return true
}

func (s *Server) sessOf(w http.ResponseWriter, r *http.Request) *Session {
	id := r.PathValue("id")
	sess := s.mgr.Get(id)
	if sess == nil {
		fail(w, 404, fmt.Errorf("会话不存在: %s", id))
	}
	return sess
}

// ---------- 会话 ----------

func (s *Server) hStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{
		"server":  "tdict-debug",
		"ssh":     s.cfg.SSH.Host,
		"zone":    s.cfg.Zone,
		"topDir":  s.cfg.TopDir,
		"listen":  s.cfg.Listen,
		"watchdog": s.cfg.WatchdogSeconds,
	})
}

func (s *Server) hLaunch(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Module string `json:"module"`
		Prog   string `json:"prog"`
		SSH    string `json:"ssh"` // 多 SSH 配置名(设置页维护);空 = 默认连接
		Zone   string `json:"zone"` // 区域覆盖(31/35/36/39/t);空 = 默认区域
	}
	if !readBody(w, r, &req) {
		return
	}
	if req.Prog == "" {
		fail(w, 400, fmt.Errorf("prog 必填(module 可留空,自动按作业名解析)"))
		return
	}
	// 按名换 SSH / 覆盖区域:克隆配置,默认链路零影响
	cfg := s.cfg
	if req.SSH != "" || req.Zone != "" {
		c2 := *s.cfg
		if req.SSH != "" {
			c2.SSH = s.cfg.SSHByName(req.SSH)
		}
		c2 = *c2.CloneWithZone(req.Zone)
		cfg = &c2
	}
	sess, err := s.mgr.LaunchWith(cfg, req.Module, req.Prog)
	if err != nil {
		fail(w, 409, err)
		return
	}
	go func() {
		if err := sess.Launch(r.Context()); err != nil {
			log.Printf("[debug] 会话 %s(%s) 启动失败: %v", sess.ID, sess.Prog, err)
			s.mgr.emit(Event{Type: "log", SessionID: sess.ID, Text: "启动失败: " + err.Error()})
			// 失败会话终结并移除,避免卡在 loading 且挡住下一次启动
			sess.ForceExit()
			s.mgr.Remove(sess.ID)
		}
	}()
	// runProg: gzzz_t 解析出的实体程序(前端源码命名/预取要用它,而不是作业编号)
	writeJSON(w, 200, map[string]any{"ok": true, "sessionId": sess.ID,
		"module": sess.Module, "prog": sess.Prog, "runProg": sess.RunProg})
}

func (s *Server) hList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"ok": true, "sessions": s.mgr.Snapshot()})
}

func (s *Server) hSnapshot(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	cur := sess.Cur()
	writeJSON(w, 200, map[string]any{
		"ok": true,
		"id": sess.ID, "module": sess.Module, "prog": sess.Prog, "runProg": sess.RunProg,
		"state": string(sess.State()), "stop": cur,
		"started":         sess.Started(),
		"breakpoints":     sess.Breakpoints(),
		"holdingSeconds":  sess.HoldingSeconds(),
		"watchdogSeconds": s.cfg.WatchdogSeconds,
	})
}

func (s *Server) hQuit(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	id := sess.ID
	if err := sess.Quit(); err != nil {
		fail(w, 500, err)
		return
	}
	s.mgr.Remove(id)
	writeJSON(w, 200, map[string]any{"ok": true})
}

// ---------- 断点 ----------

func (s *Server) hBPList(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "breakpoints": sess.Breakpoints()})
}

func (s *Server) hBPAdd(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	var req struct {
		Location string `json:"location"` // 行号 / 函数名 / file:line
	}
	if !readBody(w, r, &req) {
		return
	}
	log.Println("[bp-add] session=" + sess.ID + " loc=" + req.Location)
	bp, err := sess.Break(req.Location)
	if err != nil {
		fail(w, 400, err)
		return
	}
	log.Printf("[bp-add] result #%d %s:%d", bp.Num, bp.File, bp.Line)
	writeJSON(w, 200, map[string]any{"ok": true, "breakpoint": bp})
}

func (s *Server) hBPDel(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	num, err := strconv.Atoi(r.PathValue("num"))
	if err != nil {
		fail(w, 400, fmt.Errorf("断点编号无效"))
		return
	}
	log.Println("[bp-del] session=" + sess.ID + " num=" + strconv.Itoa(num))
	if err := sess.DeleteBreakpoint(num); err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

// ---------- 控制 / 求值 ----------

func (s *Server) hControl(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	var req struct {
		Action string `json:"action"` // continue|run|next|step|finish|until|interrupt
		Arg    string `json:"arg,omitempty"`
	}
	if !readBody(w, r, &req) {
		return
	}
	var stop *StopInfo
	var err error
	switch req.Action {
	case "continue":
		_, err = sess.Continue()
	case "run":
		_, err = sess.Run()
	case "next", "step", "finish":
		stop, err = sess.Step(req.Action)
	case "until":
		loc := req.Arg
		if loc == "" {
			stop, err = sess.Step("until")
		} else {
			stop, err = sess.Step("until " + loc)
		}
	case "interrupt":
		err = sess.Interrupt()
	default:
		fail(w, 400, fmt.Errorf("未知 action: %s", req.Action))
		return
	}
	if err != nil {
		fail(w, 409, err)
		return
	}
	resp := map[string]any{"ok": true, "state": string(sess.State())}
	if stop != nil {
		resp["stop"] = stop
	}
	writeJSON(w, 200, resp)
}

func (s *Server) hPrint(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	var req struct {
		Expr string `json:"expr"`
	}
	if !readBody(w, r, &req) {
		return
	}
	v, err := sess.Print(req.Expr)
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "value": v})
}

func (s *Server) hWhere(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	frames, err := sess.Where()
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "frames": frames})
}

func (s *Server) hRaw(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	var req struct {
		Command string `json:"command"`
	}
	if !readBody(w, r, &req) {
		return
	}
	cmd := strings.TrimSpace(req.Command)
	if cmd == "" || strings.ContainsAny(cmd, "\r\n") {
		fail(w, 400, fmt.Errorf("命令不能为空或含换行"))
		return
	}
	lower := strings.ToLower(cmd)
	if lower == "quit" || lower == "run" || strings.HasPrefix(lower, "run ") {
		fail(w, 400, fmt.Errorf("请使用会话控制接口执行 quit/run"))
		return
	}
	lines, err := sess.Raw(cmd, 30*time.Second)
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "lines": lines})
}

// ---------- 上下文查询 ----------

func (s *Server) hLocals(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	vars, err := sess.Locals()
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "vars": vars})
}

func (s *Server) hGlobals(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	vars, total, err := sess.Globals(limit)
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "vars": vars, "total": total})
}

func (s *Server) hSources(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	sources, err := sess.Sources()
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "sources": sources})
}

func (s *Server) hFunctions(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	fns, total, err := sess.Functions(limit)
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "functions": fns, "total": total})
}

func (s *Server) hAutovars(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "vars": sess.Autovars()})
}

func (s *Server) hFrame(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	var req struct {
		Num int `json:"num"`
	}
	if !readBody(w, r, &req) {
		return
	}
	if err := sess.Frame(req.Num); err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "frame": sess.CurFrame()})
}

func (s *Server) hBPEnabled(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	num, err := strconv.Atoi(r.PathValue("num"))
	if err != nil {
		fail(w, 400, fmt.Errorf("断点编号无效"))
		return
	}
	var req struct {
		Enabled *bool `json:"enabled"`
	}
	if !readBody(w, r, &req) || req.Enabled == nil {
		return
	}
	if err := sess.SetBPEnabled(num, *req.Enabled); err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

// ---------- 源码 ----------

// hSourcePreview 会话建立前预取源码(消除启动期空白):GET /api/source-preview?module=&prog=
func (s *Server) hSourcePreview(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	module, prog := q.Get("module"), q.Get("prog")
	if module == "" || prog == "" {
		fail(w, 400, fmt.Errorf("需要 module 与 prog 参数"))
		return
	}
	sf, err := s.mgr.SourcePreview(module, prog)
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "source": sf})
}

func (s *Server) hSource(w http.ResponseWriter, r *http.Request) {
	sess := s.sessOf(w, r)
	if sess == nil {
		return
	}
	q := r.URL.Query()
	var sf *SourceFile
	var err error
	switch {
	case q.Get("path") != "":
		sf, err = sess.ReadPath(q.Get("path"))
	case q.Get("file") != "":
		sf, err = sess.ResolveSource(q.Get("file"), q.Get("module"))
	default:
		err = fmt.Errorf("需要 file 或 path 参数")
	}
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "source": sf})
}

// ---------- 接口服务测试(复刻 awsq990 集成服务测试) ----------

// hWSTest POST /api/wstest {mode,url,body,soap} → 经服务器 curl 调用接口
func (s *Server) hWSTest(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Mode string `json:"mode"`
		URL  string `json:"url"`
		Body string `json:"body"`
		Soap bool   `json:"soap"`
	}
	if !readBody(w, r, &req) {
		return
	}
	if req.URL == "" {
		req.URL = WSDefaultURLFor(s.cfg, req.Mode)
	}
	conn, err := Dial(s.cfg.SSH)
	if err != nil {
		fail(w, 500, fmt.Errorf("SSH 连接失败: %w", err))
		return
	}
	defer conn.Close()
	res, err := WSTest(conn, req.URL, req.Body, req.Soap, 60)
	if err != nil {
		fail(w, 400, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "result": res})
}

// ---------- 接口日志(wsfa_t) ----------

// wsZone 返回查询用 zone(空配置默认 36)
func (s *Server) wsZone() string {
	if s.cfg.Zone == "" {
		return "36"
	}
	return s.cfg.Zone
}

// hWSLogs GET /api/wslogs?service=&onlyFail=1&limit=200
func (s *Server) hWSLogs(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("pageSize"))
	conn, err := Dial(s.cfg.SSH)
	if err != nil {
		fail(w, 500, fmt.Errorf("SSH 连接失败: %w", err))
		return
	}
	defer conn.Close()
	items, hasMore, err := listWSLogs(conn, s.wsZone(), s.cfg.TNSName(), WSLogFilter{
		Service:   q.Get("service"),
		OnlyFail:  q.Get("onlyFail") == "1",
		StartFrom: q.Get("startFrom"),
		StartTo:   q.Get("startTo"),
		Page:      page,
		PageSize:  pageSize,
	})
	if err != nil {
		fail(w, 500, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "items": items, "hasMore": hasMore})
}

// hWSLogContent GET /api/wslogs/content?rowid=
func (s *Server) hWSLogContent(w http.ResponseWriter, r *http.Request) {
	rowid := r.URL.Query().Get("rowid")
	conn, err := Dial(s.cfg.SSH)
	if err != nil {
		fail(w, 500, fmt.Errorf("SSH 连接失败: %w", err))
		return
	}
	defer conn.Close()
	item, content, err := WSLogDetail(conn, s.wsZone(), s.cfg.TNSName(), rowid)
	if err != nil {
		fail(w, 500, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "item": item, "content": content})
}

// hWSLogDebug POST /api/wslogs/debug {rowid} → 用该日志的报文重放调试
func (s *Server) hWSLogDebug(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RowID string `json:"rowid"`
	}
	if !readBody(w, r, &req) {
		return
	}
	conn, err := Dial(s.cfg.SSH)
	if err != nil {
		fail(w, 500, fmt.Errorf("SSH 连接失败: %w", err))
		return
	}
	item, content, err := WSLogDetail(conn, s.wsZone(), s.cfg.TNSName(), req.RowID)
	conn.Close()
	if err != nil {
		fail(w, 500, err)
		return
	}
	// 已有会话先结束(重放调试独占通道,旧会话不再有用)
	if cur := s.mgr.Current(); cur != nil {
		s.mgr.emit(Event{Type: "log", SessionID: cur.ID, Text: "重放调试启动,自动结束当前会话"})
		_ = cur.Quit()
		s.mgr.Remove(cur.ID)
	}
	sess, err := s.mgr.LaunchReplay(item, content)
	if err != nil {
		fail(w, 409, err)
		return
	}
	go func() {
		if err := sess.Launch(r.Context()); err != nil {
			s.mgr.emit(Event{Type: "log", SessionID: sess.ID, Text: "启动失败: " + err.Error()})
			sess.ForceExit()
			s.mgr.Remove(sess.ID)
		}
	}()
	writeJSON(w, 200, map[string]any{"ok": true, "sessionId": sess.ID,
		"module": sess.Module, "prog": sess.Prog, "runProg": sess.RunProg})
}

// ---------- 设置(多 SSH / 多数据库) ----------

// hSettingsGet 返回完整 debug 配置节
func (s *Server) hSettingsGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.cfg)
}

// hDBProbe 设置页「自动获取数据库配置」:SSH 上服务器探测连接要素(只读)
func (s *Server) hDBProbe(w http.ResponseWriter, r *http.Request) {
	var req DBProbeReq
	if !readBody(w, r, &req) {
		return
	}
	if req.Host == "" || req.User == "" {
		fail(w, 400, fmt.Errorf("请先填写 SSH 主机与账号"))
		return
	}
	if req.Type == "" {
		req.Type = "oracle"
	}
	out, err := ProbeDBConfig(req)
	if err != nil {
		fail(w, 502, err)
		return
	}
	writeJSON(w, 200, out)
}

// hSettingsPut 写回设置:仅替换 config.json 顶层 "debug" 键(其它键如 dbconfig 原样保留),
// 并原位热更新运行中的 s.cfg(Manager 等持有同一指针,立即生效)
func (s *Server) hSettingsPut(w http.ResponseWriter, r *http.Request) {
	if s.cfgPath == "" {
		fail(w, 500, fmt.Errorf("服务未挂接 config.json 路径,无法保存"))
		return
	}
	var nc Config
	if !readBody(w, r, &nc) {
		return
	}
	nc.DataDir = s.cfg.DataDir // 运行时注入字段,请求体不带
	nc.ApplyActiveEnv()        // 生效环境的连接/参数合并到顶层(顶层即生效配置)
	nc.fillDefaults()

	raw, err := os.ReadFile(s.cfgPath)
	if err != nil {
		fail(w, 500, fmt.Errorf("读取 config.json 失败: %w", err))
		return
	}
	var root map[string]any
	if err := json.Unmarshal(raw, &root); err != nil || root == nil {
		root = map[string]any{}
	}
	root["debug"] = &nc
	out, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		fail(w, 500, err)
		return
	}
	tmp := s.cfgPath + ".tmp"
	if err := os.WriteFile(tmp, out, 0o644); err != nil {
		fail(w, 500, fmt.Errorf("写入 config.json 失败: %w", err))
		return
	}
	if err := os.Rename(tmp, s.cfgPath); err != nil {
		fail(w, 500, err)
		return
	}
	*s.cfg = nc // 热生效(Listen 需重启服务才换端口,其余字段即时)
	writeJSON(w, 200, map[string]any{"ok": true})
}

// ---------- WebSocket ----------

func (s *Server) hWS(w http.ResponseWriter, r *http.Request) {
	ch, cancel := s.mgr.Subscribe("ws")
	defer cancel()
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")
	ctx := r.Context()
	for {
		select {
		case ev := <-ch:
			data, err := json.Marshal(ev)
			if err != nil {
				continue
			}
			wctx, wcancel := context.WithTimeout(ctx, 5*time.Second)
			err = conn.Write(wctx, websocket.MessageText, data)
			wcancel()
			if err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *Server) hStatic(w http.ResponseWriter, r *http.Request) {
	if s.webSub == nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<!doctype html><html lang="zh"><meta charset="utf-8">
<title>tdict debug</title><body style="font-family:system-ui;padding:40px;line-height:1.8">
<h2>tdict debug server 已运行</h2>
<p>前端尚未构建。请在 web/ 目录执行 <code>npm install &amp;&amp; npm run build</code> 后重启服务。</p>
<p>API 已可用:<code>/api/status</code>、<code>/api/ws</code>、<code>/api/sessions</code> …</p></body>`)
		return
	}
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	f, err := s.webSub.Open(strings.TrimPrefix(path, "/"))
	if err != nil {
		// SPA 兜底:未命中路径回 index.html
		f2, err2 := s.webSub.Open("index.html")
		if err2 != nil {
			http.NotFound(w, r)
			return
		}
		f = f2
		path = "/index.html"
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if st.IsDir() {
		http.NotFound(w, r)
		return
	}
	http.ServeContent(w, r, path, st.ModTime(), f.(readSeeker))
}

type readSeeker interface {
	Read(p []byte) (int, error)
	Seek(offset int64, whence int) (int64, error)
}
