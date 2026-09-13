package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"tdict/cfgfile"
	"tdict/dbconfig"
	"tdict/erpdb"
	"tdict/host"
)

// settingsPayload 前端保存的配置(仅 hosts 节的 SSH 环境与数据库连接)。
type settingsPayload struct {
	ActiveEnv string          `json:"activeEnv"`
	SSHs      []host.NamedSsh `json:"sshs"`
}

// configResp GET /api/config 的应答:配置来源 + 当前环境清单。
type configResp struct {
	ConfigPath string          `json:"configPath"`
	Exists     bool            `json:"exists"`
	ActiveEnv  string          `json:"activeEnv,omitempty"`
	SSHs       []host.NamedSsh `json:"sshs"`
}

// readRoot 读取 config.json 为可变 map;文件不存在时返回空配置(首次保存创建)。
func (s *Server) readRoot() (map[string]any, error) {
	if _, err := os.Stat(s.cfgPath); err != nil {
		if os.IsNotExist(err) {
			return map[string]any{}, nil
		}
		return nil, err
	}
	return cfgfile.Open(s.cfgPath)
}

// hConfigGet 返回 hosts 节的 SSH 环境与数据库连接(兼容旧键 debug)。
func (s *Server) hConfigGet(w http.ResponseWriter, r *http.Request) {
	root, err := s.readRoot()
	if err != nil {
		fail(w, 500, err)
		return
	}
	resp := configResp{ConfigPath: s.cfgPath, SSHs: []host.NamedSsh{}}
	if _, err := os.Stat(s.cfgPath); err == nil {
		resp.Exists = true
	}
	if h, err := cfgfile.Hosts(root); err == nil {
		b, _ := json.Marshal(h)
		_ = json.Unmarshal(b, &resp)
	}
	if resp.SSHs == nil {
		resp.SSHs = []host.NamedSsh{}
	}
	writeJSON(w, 200, resp)
}

// hConfigPut 保存配置:整块替换顶层 "hosts" 键(其余顶层键如 query/mirror/bdldoc 保留),
// 并迁移旧键 "debug";文件不存在时创建。
func (s *Server) hConfigPut(w http.ResponseWriter, r *http.Request) {
	var p settingsPayload
	if !readBody(w, r, &p) {
		return
	}
	if len(p.SSHs) == 0 {
		writeJSON(w, 200, map[string]any{"ok": false, "error": "至少保留一个 SSH 环境"})
		return
	}
	for i := range p.SSHs {
		e := &p.SSHs[i]
		e.Name = strings.TrimSpace(e.Name)
		e.Host = strings.TrimSpace(e.Host)
		e.User = strings.TrimSpace(e.User)
		e.Zone = strings.TrimSpace(e.Zone)
		if e.Port == 0 {
			e.Port = 22
		}
		if e.Host == "" {
			writeJSON(w, 200, map[string]any{"ok": false, "error": fmt.Sprintf("第 %d 个环境缺少主机地址", i+1)})
			return
		}
		if e.Name == "" {
			e.Name = e.Host
			if e.Zone != "" {
				e.Name = e.Host + "-" + e.Zone
			}
		}
	}
	// 默认环境必须指向列表中存在的环境,否则回退首条(与 host.LoadHosts 语义一致)
	if !hasEnv(p.SSHs, p.ActiveEnv) {
		p.ActiveEnv = p.SSHs[0].Name
	}

	root, err := s.readRoot()
	if err != nil {
		fail(w, 500, err)
		return
	}
	b, err := json.Marshal(p)
	if err != nil {
		fail(w, 500, err)
		return
	}
	var section map[string]any
	if err := json.Unmarshal(b, &section); err != nil {
		fail(w, 500, err)
		return
	}
	root["hosts"] = section
	delete(root, "debug") // 旧键迁移:写路径统一落到 hosts
	if err := cfgfile.Save(s.cfgPath, root); err != nil {
		fail(w, 500, err)
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

func hasEnv(list []host.NamedSsh, name string) bool {
	if name == "" {
		return false
	}
	for i := range list {
		if list[i].Name == name {
			return true
		}
	}
	return false
}

// hDBProbe 「从服务器获取数据库配置」:登录该环境 SSH,按区域只读探测连接要素。
func (s *Server) hDBProbe(w http.ResponseWriter, r *http.Request) {
	var req host.DBProbeReq
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
	out, err := host.ProbeDBConfig(req)
	if err != nil {
		fail(w, 502, err)
		return
	}
	writeJSON(w, 200, out)
}

// hDbAccVerify 账号清单「验证」:服务器侧以 账号/密码 连显式目标库 select 1(只读)。
func (s *Server) hDbAccVerify(w http.ResponseWriter, r *http.Request) {
	var req host.DBAccVerifyReq
	if !readBody(w, r, &req) {
		return
	}
	if req.Host == "" || req.User == "" {
		fail(w, 400, fmt.Errorf("请先填写 SSH 主机与账号"))
		return
	}
	if req.Account == "" {
		fail(w, 400, fmt.Errorf("账号不能为空"))
		return
	}
	if req.Type == "" {
		req.Type = "oracle"
	}
	if err := host.VerifyDBAcct(req); err != nil {
		writeJSON(w, 200, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

// hConnTest 客户端直连测试:按表单显式字段连库(凭据取账号列表首项,与 db ping 同链路)。
func (s *Server) hConnTest(w http.ResponseWriter, r *http.Request) {
	var req dbconfig.Connection
	if !readBody(w, r, &req) {
		return
	}
	if req.Host == "" {
		fail(w, 400, fmt.Errorf("请填写主机地址"))
		return
	}
	if len(req.Accounts) == 0 || req.Accounts[0].Account == "" {
		fail(w, 400, fmt.Errorf("请先在账号列表中添加账号(直连取列表首项)"))
		return
	}
	if req.Type == "" {
		req.Type = "oracle"
	}
	if req.Type != "oracle" && req.Type != "kingbase" {
		fail(w, 400, fmt.Errorf("类型非法: %s(仅 oracle/kingbase)", req.Type))
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	conn, err := erpdb.Open(ctx, req)
	if err != nil {
		writeJSON(w, 200, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	defer conn.Close()
	ver, err := conn.ServerVersion(ctx)
	if err != nil {
		writeJSON(w, 200, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "version": ver})
}
