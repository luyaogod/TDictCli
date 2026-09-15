package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"tdict/cfgfile"
)

// bdldocStatus BDL(4GL)语言文档目录的当前配置(GET/PUT /api/bdldoc)。
type bdldocStatus struct {
	OK         bool   `json:"ok"`
	Dir        string `json:"dir"`        // 当前配置的目录(绝对路径;空=未设置)
	Exists     bool   `json:"exists"`     // 该目录是否存在于本机
	ConfigPath string `json:"configPath"` // 写入的配置文件
}

// bdldocDirFromRoot 读 config.json 顶层 bdldoc.dir(相对路径转绝对)。
func bdldocDirFromRoot(root map[string]any) string {
	sec, _ := root["bdldoc"].(map[string]any)
	if sec == nil {
		return ""
	}
	d, _ := sec["dir"].(string)
	d = strings.TrimSpace(d)
	if d == "" {
		return ""
	}
	if abs, err := filepath.Abs(d); err == nil {
		return abs
	}
	return d
}

func (s *Server) bdldocStatus() bdldocStatus {
	st := bdldocStatus{OK: true, ConfigPath: s.cfgPath}
	if root, err := s.readRoot(); err == nil {
		st.Dir = bdldocDirFromRoot(root)
	}
	if st.Dir != "" {
		if fi, err := os.Stat(st.Dir); err == nil && fi.IsDir() {
			st.Exists = true
		}
	}
	return st
}

// hBdldocGet 返回当前 BDL 文档目录与是否存在于本机(等价 tdict bdldoc dir)。
func (s *Server) hBdldocGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.bdldocStatus())
}

// hBdldocPut 写入 config.json 顶层 bdldoc.dir(其余键保留;只改设置,不移动文档文件)。
func (s *Server) hBdldocPut(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Dir string `json:"dir"`
	}
	if !readBody(w, r, &req) {
		return
	}
	dir := strings.TrimSpace(req.Dir)
	if dir == "" {
		writeJSON(w, 200, map[string]any{"ok": false, "error": "请填写 BDL 文档目录"})
		return
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		fail(w, 400, fmt.Errorf("解析目录路径失败: %w", err))
		return
	}
	root, err := s.readRoot()
	if err != nil {
		fail(w, 500, err)
		return
	}
	root["bdldoc"] = map[string]any{"dir": abs}
	if err := cfgfile.Save(s.cfgPath, root); err != nil {
		fail(w, 500, err)
		return
	}
	writeJSON(w, 200, s.bdldocStatus())
}
