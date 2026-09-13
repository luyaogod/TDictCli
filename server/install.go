package server

import "net/http"

// ---------- 命令行安装(把可执行文件目录加入用户 PATH) ----------

// hInstallGet 返回安装状态:可执行文件/目录、是否已在用户 PATH。
func (s *Server) hInstallGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, getInstallStatus())
}

// hInstallAdd 把当前可执行文件所在目录加入用户 PATH(幂等,用户级无需管理员)。
func (s *Server) hInstallAdd(w http.ResponseWriter, r *http.Request) {
	st, err := addExeDirToUserPath()
	if err != nil {
		fail(w, 500, err)
		return
	}
	writeJSON(w, 200, st)
}

// hInstallRemove 从用户 PATH 移除该目录(幂等)。
func (s *Server) hInstallRemove(w http.ResponseWriter, r *http.Request) {
	st, err := removeExeDirFromUserPath()
	if err != nil {
		fail(w, 500, err)
		return
	}
	writeJSON(w, 200, st)
}
