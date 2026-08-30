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

// Connection describes a single ERP database connection.
type Connection struct {
	Name          string `json:"name"`
	Type          string `json:"type"`            // "kingbase" | "oracle" (后续支持)
	Host          string `json:"host,omitempty"`
	Port          int    `json:"port,omitempty"`
	Database      string `json:"database,omitempty"`
	User          string `json:"user"`
	Password      string `json:"password"`
	ConnectString string `json:"connectString,omitempty"` // oracle 专用: "host:port/service"
	IsDefault     bool   `json:"isDefault"`
	Description   string `json:"description,omitempty"`
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
