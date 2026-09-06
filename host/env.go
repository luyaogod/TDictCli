// 包 host:远程服务器能力共享层 —— 环境模型/SSH 连接/登录动态路径探测/源码镜像。
// 供 debug(会话调度)与 mirror/db/source/env 等 CLI 命令共同依赖,命令之间互不 import。
// tdict 除 debug serve 外均为一次性 CLI:本包不维护跨进程状态,探测结果只在单进程内使用。

package host

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tdict/dbconfig"
)

// SSHConfig 远程服务器连接配置
type SSHConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Addr 返回 SSH 地址 host:port
func (c *SSHConfig) Addr() string { return c.Host + ":" + strconv.Itoa(c.Port) }

// EntValue 企业编号(TOPENT):数字或文本均可,兼容 JSON 数字。
// 需真实编号的场景用 Int()(非数字返回 false)
type EntValue string

// UnmarshalJSON 同时接受 JSON 数字与字符串
func (e *EntValue) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "null" {
		*e = ""
		return nil
	}
	*e = EntValue(strings.Trim(s, `"`))
	return nil
}

// MarshalJSON 统一序列化为字符串
func (e EntValue) MarshalJSON() ([]byte, error) { return json.Marshal(string(e)) }

// Int 解析为数字(供数据库探测按企业编号匹配;非数字内容返回 false)
func (e EntValue) Int() (int, bool) {
	n, err := strconv.Atoi(strings.TrimSpace(string(e)))
	return n, err == nil
}

// NamedSsh 服务器/调试环境:SSH 连接 + 登录区域 + 默认企业(TOPENT)。
// 数据库连接经 DB 引用顶层 connections 条目(名称),详情在设置页 DB 页维护。
// T100 路径不允许静态配置:登录后按 zone 动态获取,失败即报错。
type NamedSsh struct {
	Name            string `json:"name"`
	SSHConfig                // 匿名嵌入:host/port/user/password 提升到 ssh 层
	Zone            string `json:"zone,omitempty"` // 登录后区域菜单代码:31开发 35测试 36正式 39PATCH t出货
	Topent          EntValue `json:"topent,omitempty"` // 默认企业编号(TOPENT);调试会话 export 用
	LaunchArgs      string              `json:"launchArgs,omitempty"`
	WatchdogSeconds int                 `json:"watchdogSeconds,omitempty"`
	DB              *dbconfig.Connection `json:"db,omitempty"` // 该环境的数据库连接(与 SSH 一对一;显式 host/port/service|库名+账号列表)
}

// Hosts 服务器环境清单:config.json 顶层 debug 节的 sshs/activeEnv 子集。
// debug 专属参数(断点/终端/启动模板等)由 tdict/debug 的 Config 承载,本包不解析。
type Hosts struct {
	SSHs      []NamedSsh
	ActiveEnv string
}

// LoadHosts 读取 config.json 顶层 debug 节中的环境清单(sshs/activeEnv),
// 供无需调试会话的命令(mirror/db/env/source)取环境;activeEnv 缺省取首条。
func LoadHosts(path string) (*Hosts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}
	var wrapper struct {
		Debug *struct {
			SSHs      []NamedSsh `json:"sshs"`
			ActiveEnv string     `json:"activeEnv"`
		} `json:"debug"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	if wrapper.Debug == nil {
		return nil, fmt.Errorf("config.json 缺少 \"debug\" 配置节")
	}
	h := &Hosts{SSHs: wrapper.Debug.SSHs, ActiveEnv: wrapper.Debug.ActiveEnv}
	if len(h.SSHs) == 0 {
		return nil, fmt.Errorf("debug.sshs 未配置任何服务器环境(请在设置-环境-SSH 页添加)")
	}
	if h.ActiveEnv == "" {
		h.ActiveEnv = h.SSHs[0].Name
	}
	return h, nil
}

// ByName 按名取环境;name 为空取 activeEnv;未命中返回 nil
func (h *Hosts) ByName(name string) *NamedSsh {
	target := name
	if target == "" {
		target = h.ActiveEnv
	}
	for i := range h.SSHs {
		if h.SSHs[i].Name == target {
			return &h.SSHs[i]
		}
	}
	return nil
}
