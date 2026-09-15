// Package server 提供本地可视化配置服务:静态前端 + config.json(hosts 节)读写
// + SSH/数据库连接探测。只覆盖「SSH 环境与数据库连接」的配置管理,不含任何调试能力。
//
// SSH 连接与服务器侧数据库探测复用 tdict/host(与 CLI 的 db discover 同源):
//   - 连接:host.Dial(密码认证)→ 先按 zone 加载 T100 环境(chenv 链)再执行只读命令;
//   - oracle:探测 ORACLE_HOME/sqlplus/TWO_TASK,再解析 tnsnames.ora 得到 host/port/service;
//   - kingbase:发现运行中实例的数据目录/端口/库名(ksql 路径一并探测)。
//
// 客户端直连测试走 tdict/erpdb(与 db ping / 在线查询同链路)。
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DefaultListen 是未显式配置监听地址时的默认地址。
const DefaultListen = "127.0.0.1:28670"

// maxPortTries 端口被占用时最多顺延尝试的次数。
const maxPortTries = 50

// Server 本地配置服务:静态前端 + REST 配置接口。
type Server struct {
	cfgPath string // config.json 路径(读写;文件不存在时首次保存创建)
	listen  string // 配置的监听地址
	webSub  fs.FS  // web/dist 子文件系统(未构建前端时为 nil)
	addr    string // 实际监听地址(端口顺延时与 listen 不同;不写回配置)
	dbPath  string // 数据同步目标 SQLite(serve 启动时由 CLI 注入;空=erp_data.db)

	// 源码镜像拉取任务(单实例:同一时间只允许一个)
	mirrorMu    sync.Mutex
	mirror      mirrorJob
	mirrorStart time.Time

	// 数据库同步任务(单实例:同一时间只允许一个)
	syncMu    sync.Mutex
	sync      dbSyncJob
	syncStart time.Time
}

// SetDBTarget 注入数据同步的默认目标(便携版:exe 同目录的 erp_data.db,由 cli/serve 传入)。
func (s *Server) SetDBTarget(p string) { s.dbPath = p }

// defaultDBTarget 返回默认同步目标(未注入时用当前目录下 erp_data.db)。
func (s *Server) defaultDBTarget() string {
	if s.dbPath != "" {
		return s.dbPath
	}
	abs, err := filepath.Abs("erp_data.db")
	if err != nil {
		return "erp_data.db"
	}
	return abs
}

// syncTarget 返回当前同步目标:优先 config.json 顶层 sync.target(页面可改),
// 否则用默认目标(exe 同目录)。文件不存在时由同步过程创建(含父目录)。
func (s *Server) syncTarget() string {
	if root, err := s.readRoot(); err == nil {
		if d := syncTargetFromRoot(root); d != "" {
			return d
		}
	}
	return s.defaultDBTarget()
}

// syncTargetFromRoot 读 config.json 顶层 sync.target(相对路径转绝对)。
func syncTargetFromRoot(root map[string]any) string {
	sec, _ := root["sync"].(map[string]any)
	if sec == nil {
		return ""
	}
	d, _ := sec["target"].(string)
	d = strings.TrimSpace(d)
	if d == "" {
		return ""
	}
	if abs, err := filepath.Abs(d); err == nil {
		return abs
	}
	return d
}

// New 创建服务实例;web 由 main 经 //go:embed 注入,可为 nil(此时返回引导页)。
func New(cfgPath, listen string, web fs.FS) *Server {
	if strings.TrimSpace(listen) == "" {
		listen = DefaultListen
	}
	s := &Server{cfgPath: cfgPath, listen: listen}
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

// ListenAddr 返回实际监听地址(未启动时返回配置值)。
func (s *Server) ListenAddr() string {
	if s.addr != "" {
		return s.addr
	}
	return s.listen
}

// Listen 建立监听;配置端口被占用时顺延到下一个空闲端口(最多 maxPortTries 个)。
// 返回 listener 与实际监听地址(host:port);s.listen 保持原值不变。
func (s *Server) Listen() (net.Listener, string, error) {
	host, portStr, err := net.SplitHostPort(s.listen)
	if err != nil {
		return nil, "", fmt.Errorf("非法监听地址 %q: %w", s.listen, err)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 0 || port > 65535 {
		return nil, "", fmt.Errorf("非法监听端口 %q: %w", portStr, err)
	}
	var lastErr error
	for i := 0; i < maxPortTries; i++ {
		addr := net.JoinHostPort(host, strconv.Itoa(port+i))
		ln, err := net.Listen("tcp", addr)
		if err == nil {
			real := addr
			if tcp, ok := ln.Addr().(*net.TCPAddr); ok {
				h := host
				if h == "" {
					h = "127.0.0.1"
				}
				real = net.JoinHostPort(h, strconv.Itoa(tcp.Port))
			}
			s.addr = real
			return ln, real, nil
		}
		lastErr = err
	}
	return nil, "", fmt.Errorf("监听 %s 失败(尝试 %d 个端口均不可用): %w", s.listen, maxPortTries, lastErr)
}

// Serve 用既有 listener 提供 HTTP 服务,阻塞到 ctx 取消或连接关闭。
func (s *Server) Serve(ctx context.Context, ln net.Listener) error {
	mux := http.NewServeMux()
	s.routes(mux)
	srv := &http.Server{Handler: mux}

	runCtx, cancelRun := context.WithCancel(ctx)
	defer cancelRun()
	go func() {
		<-runCtx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	err := srv.Serve(ln)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func (s *Server) routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/status", s.hStatus)
	mux.HandleFunc("GET /api/config", s.hConfigGet)
	mux.HandleFunc("PUT /api/config", s.hConfigPut)
	mux.HandleFunc("POST /api/dbprobe", s.hDBProbe)
	mux.HandleFunc("POST /api/dbaccverify", s.hDbAccVerify)
	mux.HandleFunc("POST /api/conntest", s.hConnTest)
	mux.HandleFunc("GET /api/mirror", s.hMirrorGet)
	mux.HandleFunc("PUT /api/mirror", s.hMirrorPut)
	mux.HandleFunc("POST /api/mirror/pull", s.hMirrorPull)
	mux.HandleFunc("GET /api/dbsync", s.hDBSyncGet)
	mux.HandleFunc("POST /api/dbsync", s.hDBSyncPost)
	mux.HandleFunc("PUT /api/dbsync", s.hDBSyncPut)
	mux.HandleFunc("GET /api/install", s.hInstallGet)
	mux.HandleFunc("POST /api/install", s.hInstallAdd)
	mux.HandleFunc("DELETE /api/install", s.hInstallRemove)
	mux.HandleFunc("GET /api/bdldoc", s.hBdldocGet)
	mux.HandleFunc("PUT /api/bdldoc", s.hBdldocPut)
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

func (s *Server) hStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{
		"ok":     true,
		"server": "tdict",
		"listen": s.ListenAddr(),
	})
}

type readSeeker interface {
	Read(p []byte) (int, error)
	Seek(offset int64, whence int) (int64, error)
}

// hStatic 提供前端构建产物;未构建时返回引导页。
func (s *Server) hStatic(w http.ResponseWriter, r *http.Request) {
	if s.webSub == nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<!doctype html><html lang="zh"><meta charset="utf-8">
<title>tdict</title><body style="font-family:system-ui;padding:40px;line-height:1.8">
<h2>tdict 配置服务已运行</h2>
<p>前端尚未构建。请在 web/ 目录执行 <code>npm install &amp;&amp; npm run build</code> 后重启服务。</p>
<p>API 已可用:<code>/api/config</code>、<code>/api/dbprobe</code>、<code>/api/conntest</code>。</p></body>`)
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
	if err != nil || st.IsDir() {
		http.NotFound(w, r)
		return
	}
	http.ServeContent(w, r, path, st.ModTime(), f.(readSeeker))
}
