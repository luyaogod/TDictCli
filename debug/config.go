package debug

// 服务器环境(SSH)与数据库连接配置。config.json 顶层 "debug" 键。
// 数据库连接不内嵌在环境中:统一存于顶层 "connections"(dbconfig 包),
// SSH 环境经 DBConn 按名称引用 —— 调试/服务器侧连库用被引用连接的显式
// host/port/service(库名)+ 账号清单;客户端直连(ping/--online/sync)用同一份列表。
// Oracle 全部显式 host:port/service(EZCONNECT),tnsnames/chenv/ORA 环境读取
// 仅作为「从服务器获取」的辅助手段,不参与运行时连接。

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
// 数据库连接经 DBConn 引用顶层 connections 条目(名称),详情在设置页 DB 页维护。
// T100 路径(topDir/moduleRoots)不允许静态配置:登录后按 zone 动态获取,失败即报错。
type NamedSsh struct {
	Name            string `json:"name"`
	SSHConfig                // 匿名嵌入:host/port/user/password 提升到 ssh 层
	Zone            string `json:"zone,omitempty"` // 登录后区域菜单代码:31开发 35测试 36正式 39PATCH t出货
	Topent          EntValue `json:"topent,omitempty"` // 默认企业编号(TOPENT);调试会话 export 用
	LaunchArgs      string              `json:"launchArgs,omitempty"`
	WatchdogSeconds int                 `json:"watchdogSeconds,omitempty"`
	DB              *dbconfig.Connection `json:"db,omitempty"` // 该环境的数据库连接(与 SSH 一对一;显式 host/port/service|库名+账号列表)
}

// DefaultListen 是 serve 未显式配置监听地址时的默认地址。
const DefaultListen = "127.0.0.1:28670"

// Config debug 功能配置,存放在 config.json 顶层 "debug" 键。
// 持久化字段 = 全局调试参数 + sshs 列表;SSH/Zone/Topent/DB 为运行时合并结果
// (ApplyActiveEnv 由活跃 ssh + dbConn 引用生成),不参与持久化。
type Config struct {
	SSHs            []NamedSsh `json:"sshs,omitempty"`      // 服务器环境列表(设置页 SSH 页维护)
	ActiveEnv       string     `json:"activeEnv,omitempty"` // 当前生效的 ssh 名;空=取 sshs 首条
	Listen          string     `json:"listen"`              // HTTP 监听地址
	LaunchArgs      string     `json:"launchArgs"`          // 作业启动参数默认模板,{prog} 替换为作业名(ssh 可覆盖)
	WatchdogSeconds int        `json:"watchdogSeconds"`     // 停站停留超时默认(秒);ssh 可覆盖
	FGLServer       string     `json:"fglserver"`           // 留空使用 T100 按 SSH 来源 IP 自动设置
	TermWidth       int        `json:"termWidth"`
	TermHeight      int        `json:"termHeight"`
	PrintElements   int        `json:"printElements"`      // fgldb 单次 print 的数组元素上限;0=默认 1000
	PersistBPs      *bool      `json:"persistBreakpoints"` // 断点持久化开关(nil 视为 true)
	DataDir         string     `json:"-"`                  // 数据目录(断点持久化等);serve 注入,空=禁用

	// ---- 运行时(合并结果,json:"-" 不持久化) ----
	SSH      SSHConfig             `json:"-"` // 生效 SSH(activeEnv 合并)
	Zone     string                `json:"-"` // 生效登录区域
	Topent   EntValue              `json:"-"` // 生效默认企业(调试会话 export TOPENT)
	DB       *dbconfig.Connection `json:"-"` // 生效数据库连接(activeEnv 的 ssh.db 深拷贝;nil=该 ssh 未挂库)
	Runtime  *RuntimeEnv          `json:"-"` // 登录后动态获取的 T100 路径(探针/选区回显);nil=尚未获取,需登录探测
}

// zoneTNSName 区域代码 → 数据库 TNS 别名推导(31→t35dev,35→t35tst,36→t35prd,
// 39→t35pth,t→topprd)。仅「从服务器获取」辅助探测用——运行时连接一律显式 host/port/service。
func zoneTNSName(zone string) string {
	switch zone {
	case "31":
		return "t35dev"
	case "35":
		return "t35tst"
	case "39":
		return "t35pth"
	case "t":
		return "topprd"
	default:
		return "t35prd"
	}
}

// applySsh 把某 ssh 环境的连接与启动参数合并到运行时字段(activeEnv 与 CloneEnv 共用)
func (c *Config) applySsh(e *NamedSsh) {
	if e.Host != "" {
		c.SSH = e.SSHConfig
		if c.SSH.Port == 0 {
			c.SSH.Port = 22
		}
	}
	if e.Zone != "" {
		c.Zone = e.Zone
	}
	c.Topent = e.Topent
	// 换环境 = 换服务器/区域:清动态路径,登录后重新获取
	c.Runtime = nil
	if e.LaunchArgs != "" {
		c.LaunchArgs = e.LaunchArgs
	}
	if e.WatchdogSeconds > 0 {
		c.WatchdogSeconds = e.WatchdogSeconds
	}
	// 该环境一对一挂载的数据库连接(深拷贝,防共享底层切片)
	c.DB = nil
	if e.DB != nil {
		x := *e.DB
		x.Accounts = append([]dbconfig.DBAcct(nil), e.DB.Accounts...)
		c.DB = &x
	}
	c.fillDefaults()
}

// ApplyActiveEnv 把当前生效环境(activeEnv)的 SSH/zone/topent/库引用合并到运行时字段。
// activeEnv 为空时自动取 sshs 首条;无任何 ssh 时保持空(LoadConfig 已校验不允许)。
func (c *Config) ApplyActiveEnv() {
	if c.ActiveEnv == "" && len(c.SSHs) > 0 {
		c.ActiveEnv = c.SSHs[0].Name
	}
	if c.ActiveEnv == "" {
		return
	}
	for i := range c.SSHs {
		if c.SSHs[i].Name == c.ActiveEnv {
			c.applySsh(&c.SSHs[i])
			return
		}
	}
}

// SSHByName 按名取 SSH 连接;空名/未命中返回生效 SSH(合并后的 c.SSH)
func (c *Config) SSHByName(name string) SSHConfig {
	if name != "" {
		for _, s := range c.SSHs {
			if s.Name == name {
				if s.Port == 0 {
					s.Port = 22
				}
				return s.SSHConfig
			}
		}
	}
	return c.SSH
}

// EnvName 会话所属环境名:优先 activeEnv,缺省用 host-zone 推导名
func (c *Config) EnvName() string {
	if c.ActiveEnv != "" {
		return c.ActiveEnv
	}
	return c.SSH.Host + "-" + c.Zone
}

// CloneEnv 复制配置并切换到指定 ssh 环境(name 命中 SSHs 之一):
// 用于会话「切换/重启」按目标环境重连;未命中返回 nil。
func (c *Config) CloneEnv(name string) *Config {
	if name == "" {
		return nil
	}
	var hit *NamedSsh
	for i := range c.SSHs {
		if c.SSHs[i].Name == name {
			hit = &c.SSHs[i]
			break
		}
	}
	if hit == nil {
		return nil
	}
	c2 := *c
	c2.ActiveEnv = name
	c2.applySsh(hit)
	return &c2
}

// TopentInt 返回生效默认企业的数字编号(数据库探测用;文本/未配置返回 0=仅列映射)
func (c *Config) TopentInt() int {
	n, _ := c.Topent.Int()
	return n
}

// BPsPersisted 断点持久化是否启用
func (c *Config) BPsPersisted() bool {
	return c.DataDir != "" && (c.PersistBPs == nil || *c.PersistBPs)
}

func (c *Config) fillDefaults() {
	if c.Listen == "" {
		c.Listen = DefaultListen
	}
	if c.LaunchArgs == "" {
		c.LaunchArgs = "BBDL512840855a 2 12345 'N' {prog}"
	}
	if c.WatchdogSeconds == 0 {
		c.WatchdogSeconds = 180
	}
	if c.TermWidth == 0 {
		c.TermWidth = 200
	}
	if c.TermHeight == 0 {
		c.TermHeight = 50
	}
	if c.SSH.Port == 0 {
		c.SSH.Port = 22
	}
	if c.PrintElements == 0 {
		c.PrintElements = 1000
	}
}

// TopDirActual 返回登录后动态获取的区域顶级目录(TOP)。
// 未获取(Runtime 为 nil)时返回空串——调用方须先确保动态环境已获取,不得回退静态路径。
func (c *Config) TopDirActual() string {
	if c.Runtime != nil {
		return c.Runtime.TOP
	}
	return ""
}

// ModuleRootsActual 返回登录后动态获取的源码查找根目录(ERP/COM/wss)。
// 未获取时返回 nil——调用方须先确保动态环境已获取,不得回退静态路径。
func (c *Config) ModuleRootsActual() []string {
	if c.Runtime != nil && c.Runtime.ERP != "" {
		// com/wss 是 WebService 程序(wssp* / awsp*)的专用模块目录
		return []string{c.Runtime.ERP, c.Runtime.COM, c.Runtime.COM + "/wss"}
	}
	return nil
}

// Top 返回区域顶级目录(去尾斜杠)
func (c *Config) Top() string { return c.TopDirActual() }

// CloneWithZone 复制配置并覆盖区域(启动参数 --zone 用):
// 区域变了,登录后的动态路径需重新获取(Runtime 置空),不做任何静态推导。
func (c *Config) CloneWithZone(zone string) *Config {
	c2 := *c
	if zone != "" && zone != c.Zone {
		c2.Zone = zone
		c2.Runtime = nil // 区域变了,动态路径需重新获取
	}
	return &c2
}

// ModuleDir 返回模块主目录,如 /u1/t35prd/erp/asf;wss 模块挂载在 com 下
func (c *Config) ModuleDir(module string) string {
	top := c.TopDirActual()
	if module == "wss" {
		return top + "/com/wss"
	}
	return top + "/erp/" + module
}

// FGLSOURCEPath 返回 launch 前 export 的源码搜索路径
func (c *Config) FGLSOURCEPath(module string) string {
	top := c.TopDirActual()
	base := top + "/erp/" + module
	if module == "wss" {
		base = top + "/com/wss"
	}
	dirs := []string{
		base + "/4gl",
		base + "/42m",
		top + "/com/lib/42m",
		top + "/com/sub/42m",
		top + "/com/qry/42m",
	}
	out := ""
	for i, d := range dirs {
		if i > 0 {
			out += ":"
		}
		out += d
	}
	return out
}

// LoadConfig 从 config.json 读取顶层 "debug" 键并填充默认值;
// 同时读入同文件顶层 "connections" 快照(供 dbConn 引用解析)。
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}
	var wrapper struct {
		Debug *Config `json:"debug"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	if wrapper.Debug == nil {
		return nil, fmt.Errorf("config.json 缺少 \"debug\" 配置节")
	}
	cfg := wrapper.Debug
	cfg.fillDefaults()
	if len(cfg.SSHs) == 0 {
		return nil, fmt.Errorf("debug.sshs 未配置任何服务器环境(请在设置-环境-SSH 页添加)")
	}
	return cfg, nil
}

// Addr 返回 SSH 地址 host:port
func (c *SSHConfig) Addr() string { return c.Host + ":" + strconv.Itoa(c.Port) }
