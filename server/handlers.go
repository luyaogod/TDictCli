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

// settingsPayload 前端保存的配置。
// SSHs 为 nil 表示"本次不动 hosts 节"(设置页只改查询数据源时用),
// 非 nil 时按原逻辑整块替换 hosts;QuerySource 为 nil 表示不动 query.source。
type settingsPayload struct {
	ActiveEnv   string           `json:"activeEnv"`
	SSHs        *[]host.NamedSsh `json:"sshs"`
	QuerySource *string          `json:"querySource,omitempty"`
}

// configResp GET /api/config 的应答:配置来源 + 当前环境清单。
type configResp struct {
	ConfigPath string          `json:"configPath"`
	Exists     bool            `json:"exists"`
	ActiveEnv  string          `json:"activeEnv,omitempty"`
	SSHs       []host.NamedSsh `json:"sshs"`
	// query.source:空=在线(缺省)/"local"=本地库/环境名=该环境远程(见「设置」页)
	QuerySource string `json:"querySource,omitempty"`
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
	if q, ok := root["query"].(map[string]any); ok {
		if src, ok := q["source"].(string); ok {
			resp.QuerySource = src
		}
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
	if p.SSHs == nil && p.QuerySource == nil {
		writeJSON(w, 200, map[string]any{"ok": false, "error": "没有要保存的内容"})
		return
	}

	root, err := s.readRoot()
	if err != nil {
		fail(w, 500, err)
		return
	}

	// ① hosts 节:环境配置页保存时整块替换;设置页只改查询数据源时不带 sshs,这里就不动它
	// (避免设置页用陈旧快照覆盖刚在环境配置页改好的环境)
	if p.SSHs != nil {
		list := *p.SSHs
		if len(list) == 0 {
			writeJSON(w, 200, map[string]any{"ok": false, "error": "至少保留一个 SSH 环境"})
			return
		}
		for i := range list {
			e := &list[i]
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
		if !hasEnv(list, p.ActiveEnv) {
			p.ActiveEnv = list[0].Name
		}
		b, err := json.Marshal(list)
		if err != nil {
			fail(w, 500, err)
			return
		}
		var arr []any
		if err := json.Unmarshal(b, &arr); err != nil {
			fail(w, 500, err)
			return
		}
		root["hosts"] = map[string]any{"activeEnv": p.ActiveEnv, "sshs": arr}
		delete(root, "debug") // 旧键迁移:写路径统一落到 hosts
	}

	// ② 查询数据源:顶层 query.source("local"/环境名;空串=清除,回到缺省的"在线")
	if p.QuerySource != nil {
		src := strings.TrimSpace(*p.QuerySource)
		q, _ := root["query"].(map[string]any)
		if q == nil {
			q = map[string]any{}
		}
		if src == "" {
			delete(q, "source")
		} else {
			q["source"] = src
		}
		if len(q) == 0 {
			delete(root, "query")
		} else {
			root["query"] = q
		}
	}

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
