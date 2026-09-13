// 包 host:远程服务器能力共享层 —— 环境模型/SSH 连接/登录动态路径探测/源码镜像。
// 供 mirror/db/source/env 等 CLI 命令共同依赖,命令之间互不 import。
// tdict 均为一次性 CLI:本包不维护跨进程状态,探测结果只在单进程内使用。

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

// NamedSsh 服务器环境:SSH 连接 + 登录区域 + 默认企业(TOPENT)。
// 数据库连接经内嵌 DB 与之一对一挂载。
// T100 路径不允许静态配置:登录后按 zone 动态获取,失败即报错。
type NamedSsh struct {
	Name      string               `json:"name"`
	SSHConfig                      // 匿名嵌入:host/port/user/password 提升到 ssh 层
	Zone      string               `json:"zone,omitempty"`   // 登录后区域菜单代码:31开发 35测试 36正式 39PATCH t出货
	Topent    EntValue             `json:"topent,omitempty"` // 默认企业编号(TOPENT)
	DB        *dbconfig.Connection `json:"db,omitempty"`     // 该环境的数据库连接(与 SSH 一对一;显式 host/port/service|库名+账号列表)
}

// Hosts 服务器环境清单:config.json 顶层 hosts 节(兼容旧键 debug)的 sshs/activeEnv 子集。
type Hosts struct {
	SSHs      []NamedSsh
	ActiveEnv string
}

// hostsSection config.json 顶层环境节的持久化字段(hosts 与旧键 debug 同构)。
type hostsSection struct {
	SSHs      []NamedSsh `json:"sshs"`
	ActiveEnv string     `json:"activeEnv"`
}

// LoadHosts 读取 config.json 顶层的环境清单(sshs/activeEnv),
// 优先取新键 "hosts",缺失时兼容旧键 "debug";activeEnv 缺省取首条。
func LoadHosts(path string) (*Hosts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}
	var wrapper struct {
		Hosts *hostsSection `json:"hosts"`
		Debug *hostsSection `json:"debug"` // 旧键兼容:读取不回写,写路径由 cfgfile.Hosts 迁移
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	sec := wrapper.Hosts
	if sec == nil {
		sec = wrapper.Debug
	}
	if sec == nil {
		return nil, fmt.Errorf("config.json 缺少 \"hosts\" 配置节")
	}
	h := &Hosts{SSHs: sec.SSHs, ActiveEnv: sec.ActiveEnv}
	if len(h.SSHs) == 0 {
		return nil, fmt.Errorf("尚未配置 SSH 环境:请运行 tdict serve 打开配置页添加,或编辑 config.json 的 hosts.sshs")
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
