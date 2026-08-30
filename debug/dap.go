package debug

// dap.go — DAP(Debug Adapter Protocol)客户端最小实现。
// 服务器侧适配器:fglrun --da-debugger(Genero 3.21+ 内嵌,stdio 上的
// Content-Length 帧 + JSON 消息);本文件经 SSH exec 通道与其对话。
// 协议参考 https://microsoft.github.io/debug-adapter-protocol/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

// dapFrame DAP 消息帧(request/response/event 共用)
type dapFrame struct {
	Seq        int             `json:"seq"`
	Type       string          `json:"type"` // request | response | event
	Command    string          `json:"command,omitempty"`
	Arguments  json.RawMessage `json:"arguments,omitempty"`
	Body       json.RawMessage `json:"body,omitempty"`
	Event      string          `json:"event,omitempty"`
	RequestSeq int             `json:"request_seq"`
	Success    bool            `json:"success"`
	Message    string          `json:"message,omitempty"`
}

// dapClient 单条适配器连接的请求/事件复用客户端
type dapClient struct {
	w       io.Writer
	pending map[int]chan *dapFrame
	evCh    chan *dapFrame
	mu      sync.Mutex
	seq     int
	closed  chan struct{}
}

func newDAPClient(r io.Reader, w io.Writer) *dapClient {
	c := &dapClient{
		w:       w,
		pending: map[int]chan *dapFrame{},
		evCh:    make(chan *dapFrame, 64),
		closed:  make(chan struct{}),
	}
	go c.readLoop(r)
	return c
}

// request 发送请求并等待响应(超时返回错误)
func (c *dapClient) request(command string, args any, timeout time.Duration) (*dapFrame, error) {
	c.mu.Lock()
	c.seq++
	seq := c.seq
	ch := make(chan *dapFrame, 1)
	c.pending[seq] = ch
	c.mu.Unlock()

	body, err := json.Marshal(map[string]any{
		"seq": seq, "type": "request", "command": command, "arguments": args,
	})
	if err != nil {
		return nil, err
	}
	head := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	c.mu.Lock()
	_, err = io.WriteString(c.w, head+string(body))
	c.mu.Unlock()
	if err != nil {
		return nil, fmt.Errorf("发送 %s 失败: %w", command, err)
	}
	select {
	case resp := <-ch:
		if !resp.Success {
			msg := resp.Message
			if msg == "" {
				msg = "未知错误"
			}
			return resp, fmt.Errorf("%s: %s", command, msg)
		}
		return resp, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("%s 响应超时(%s)", command, timeout)
	case <-c.closed:
		return nil, fmt.Errorf("适配器连接已关闭(%s)", command)
	}
}

// events 事件流
func (c *dapClient) events() <-chan *dapFrame { return c.evCh }

func (c *dapClient) close() {
	select {
	case <-c.closed:
	default:
		close(c.closed)
	}
}

func (c *dapClient) readLoop(r io.Reader) {
	br := bufio.NewReaderSize(r, 64*1024)
	for {
		// 读头部(空行结束)
		clen := 0
		for {
			line, err := br.ReadString('\n')
			if err != nil {
				c.shutdown()
				return
			}
			line = strings.TrimRight(line, "\r\n")
			if line == "" {
				break
			}
			if i := strings.Index(strings.ToLower(line), "content-length:"); i == 0 {
				if n, err := strconv.Atoi(strings.TrimSpace(line[15:])); err == nil {
					clen = n
				}
			}
		}
		if clen <= 0 {
			c.shutdown()
			return
		}
		body := make([]byte, clen)
		if _, err := io.ReadFull(br, body); err != nil {
			c.shutdown()
			return
		}
		var f dapFrame
		if err := json.Unmarshal(body, &f); err != nil {
			continue
		}
		switch f.Type {
		case "response":
			c.mu.Lock()
			ch := c.pending[f.RequestSeq]
			delete(c.pending, f.RequestSeq)
			c.mu.Unlock()
			if ch != nil {
				ch <- &f
			}
		case "event":
			select {
			case c.evCh <- &f:
			case <-c.closed:
				return
			}
		}
	}
}

func (c *dapClient) shutdown() {
	select {
	case <-c.closed:
	default:
		close(c.closed)
	}
	c.mu.Lock()
	for _, ch := range c.pending {
		close(ch)
	}
	c.pending = map[int]chan *dapFrame{}
	c.mu.Unlock()
}

// dapVar DAP variables 请求返回的变量
type dapVar struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	Type               string `json:"type,omitempty"`
	VariablesReference int    `json:"variablesReference"`
}

// splitShellArgs 按空格切分启动参数,单引号内保留空格(T100 参数如 'N')
func splitShellArgs(s string) []string {
	var out []string
	var cur strings.Builder
	inQ := false
	for _, r := range s {
		switch {
		case r == '\'':
			inQ = !inQ
		case r == ' ' && !inQ:
			if cur.Len() > 0 {
				out = append(out, cur.String())
				cur.Reset()
			}
		default:
			cur.WriteRune(r)
		}
	}
	if cur.Len() > 0 {
		out = append(out, cur.String())
	}
	return out
}
