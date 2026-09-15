package main

import (
	"embed"

	"tdict/cli"
)

// AI 技能文件不再内嵌:以仓库根目录的 skills/ 随发行包(便携版)一起分发,
// 用户可直接编辑;`tdict install skills` 把它复制到当前目录。

// webFS 前端构建产物(web/dist):可视化配置页(SSH 环境与数据库),由 tdict serve 提供。
// 构建方式:cd web && npm install && npm run build
//
//go:embed all:web/dist
var webFS embed.FS

func main() {
	cli.Execute(webFS)
}
