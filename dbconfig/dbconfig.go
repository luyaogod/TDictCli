package dbconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Config holds the ERP database connection configuration.
type Config struct {
	Connections []Connection `json:"connections"`
}

// ViaSSH SSH 端口转发隧道:客户端不可达 DB、但 DB 对 SSH 服务器(或同网)可达时,
// 客户端先建到 SSH 的隧道(本地 bindPort → remoteHost:port),驱动改连 127.0.0.1:bindPort。
type ViaSSH struct {
	Host       string `json:"host"`                 // SSH 服务器地址
	Port       int    `json:"port,omitempty"`       // SSH 端口(0=22)
	User       string `json:"user"`                 // SSH 账号
	Password   string `json:"password,omitempty"`   // SSH 密码
	BindPort   int    `json:"bindPort,omitempty"`   // 本地监听端口(0=自动选空闲)
	RemoteHost string `json:"remoteHost,omitempty"` // DB 在 SSH 侧的真实地址(默认取 connection.host)
	RemotePort int    `json:"remotePort,omitempty"` // DB 在 SSH 侧的真实端口(默认取 connection.port)
}

// EffectiveRemote 返回隧道远端转发目标(缺省回退到连接自身 host/port)。
func (v *ViaSSH) EffectiveRemote(fallbackHost string, fallbackPort int) (string, int) {
	h := v.RemoteHost
	if h == "" {
		h = fallbackHost
	}
	p := v.RemotePort
	if p <= 0 {
		p = fallbackPort
	}
	return h, p
}

// Connection describes a single ERP database connection.
type Connection struct {
	Name          string `json:"name"`
	Type          string `json:"type"` // "kingbase" | "oracle"
	Host          string `json:"host,omitempty"`
	Port          int    `json:"port,omitempty"`
	Database      string `json:"database,omitempty"`
	User          string `json:"user"`
	Password      string `json:"password"`
	ConnectString string `json:"connectString,omitempty"` // oracle 专用: "host:port/service"
	IsDefault     bool   `json:"isDefault"`
	Description   string `json:"description,omitempty"`
	// Source 来源标记(可省):manual | ssh:<envName> | ssh:<envName>-<ent>。
	// 仅用于标注"此连接由哪个 SSH 环境自动发现生成",不参与连接逻辑。
	Source string  `json:"source,omitempty"`
	ViaSSH *ViaSSH `json:"viaSsh,omitempty"` // 可省:经 SSH 隧道转发后再直连
}

// Address returns a human-readable connection address for display.
func (c *Connection) Address() string {
	if c.Type == "oracle" {
		return c.ConnectString
	}
	if c.Host == "" {
		return ""
	}
	if c.Port > 0 {
		return fmt.Sprintf("%s:%d/%s", c.Host, c.Port, c.Database)
	}
	return c.Host
}

// Load reads and parses the JSON configuration file at the given path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败 (%s): %w", path, err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败 (%s): %w", path, err)
	}
	return &cfg, nil
}

// FindByName returns the connection with the given name.
func (c *Config) FindByName(name string) (*Connection, error) {
	for i := range c.Connections {
		if c.Connections[i].Name == name {
			return &c.Connections[i], nil
		}
	}
	return nil, fmt.Errorf("未找到连接 \"%s\"。可用连接: %s", name, c.names())
}

// FindDefault returns the connection marked as default (isDefault: true).
func (c *Config) FindDefault() (*Connection, error) {
	for i := range c.Connections {
		if c.Connections[i].IsDefault {
			return &c.Connections[i], nil
		}
	}
	return nil, fmt.Errorf("未设置默认连接 (isDefault: true)。可用连接: %s", c.names())
}

func (c *Config) names() string {
	if len(c.Connections) == 0 {
		return "无"
	}
	var names []string
	for _, conn := range c.Connections {
		names = append(names, conn.Name)
	}
	return strings.Join(names, ", ")
}
