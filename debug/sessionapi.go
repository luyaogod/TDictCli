package debug

import (
	"context"
	"time"
)

// 调试会话统一接口:PTY 刮屏(Session)与 DAP 协议(DAPSession)两种实现。
// Manager/API/MCP 只依赖该接口,协议切换由 config.json debug.mode 决定。
type session interface {
	// 标识
	ID() string
	Module() string
	Prog() string
	RunProg() string
	Mode() string // "pty" | "dap"
	SetArgsOverride(args string)
	Log(text string)

	// 生命周期
	Launch(ctx context.Context) error
	Quit() error
	ForceExit()
	Interrupt() error

	// 控制
	Continue() (*execResult, error)
	Run() (*execResult, error)
	Step(cmd string) (*StopInfo, error)
	WaitForStop(timeout time.Duration) (*StopInfo, error)

	// 状态
	Started() bool
	State() State
	Cur() StopInfo
	CurFrame() int
	Frame(n int) error
	Breakpoints() []Breakpoint
	HoldingSeconds() float64
	Restoring() bool

	// 断点
	Break(loc string) (*Breakpoint, error)
	DeleteBreakpoint(num int) error
	SetBPEnabled(num int, enabled bool) error

	// 求值 / 上下文
	Print(expr string) (string, error)
	Where() ([]Frame, error)
	Locals() ([]VarItem, error)
	Globals(limit int) ([]VarDecl, int, error)
	Sources() ([]string, error)
	Functions(limit int) ([]string, int, error)
	InfoLine(loc string) (string, int, error)
	Autovars() []VarItem
	Raw(cmd string, timeout time.Duration) ([]string, error)

	// 源码
	ResolveSource(dvmFile, module string) (*SourceFile, error)
	ReadPath(path string) (*SourceFile, error)
}

// ---------- Session(PTY 实现)的访问器:字段已私有化,接口方法取值 ----------

func (s *Session) ID() string           { return s.id }
func (s *Session) Module() string       { return s.module }
func (s *Session) Prog() string         { return s.prog }
func (s *Session) RunProg() string      { return s.runProg }
func (s *Session) Mode() string         { return "pty" }
func (s *Session) SetArgsOverride(a string) { s.argsOverride = a }
func (s *Session) Restoring() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.restoring
}

func (s *Session) Log(text string) { s.emitEvent(Event{Type: "log", Text: text}) }
