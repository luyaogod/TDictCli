package main

import (
	"embed"

	"tdict/cli"
)

//go:embed .claude/skills/*.md
var skillFS embed.FS

// webFS 前端构建产物(web/dist):可视化配置页(SSH 环境与数据库),由 tdict serve 提供。
// 构建方式:cd web && npm install && npm run build
//
//go:embed all:web/dist
var webFS embed.FS

func main() {
	cli.Execute(skillFS, webFS)
}
