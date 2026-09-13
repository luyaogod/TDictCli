package server

import (
	"os"
	"path/filepath"
)

// pathInstallStatus 命令行安装状态(GET /api/install 返回)。
type pathInstallStatus struct {
	Supported  bool   `json:"supported"`  // 本平台是否支持自动写入 PATH
	ExePath    string `json:"exePath"`    // 当前可执行文件
	ExeDir     string `json:"exeDir"`     // 将被加入 PATH 的目录
	InUserPath bool   `json:"inUserPath"` // 该目录是否已在用户 PATH 中
	UserPath   string `json:"userPath,omitempty"`
	Manual     string `json:"manual,omitempty"` // 不支持时的等价手动命令
	Note       string `json:"note,omitempty"`
}

// exePath 当前运行的可执行文件绝对路径(解析符号链接)。
func exePath() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	if r, err := filepath.EvalSymlinks(exe); err == nil {
		exe = r
	}
	return exe
}
