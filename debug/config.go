package debug

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// SSHConfig 远程服务器连接配置
type SSHConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// NamedSSH 命名 SSH 连接(设置页维护的多服务器列表;start --ssh 按名引用)
type NamedSSH struct {
	Name string `json:"name"`
	SSHConfig
}

// DBConfig 数据库连接探查配置(config.json debug.db 节,全部可选)
type DBConfig struct {
	Ent        int    `json:"ent"`        // 默认企业编号(TOPENT);0=不指定
	SQLPlus    string `json:"sqlplus"`    // sqlplus 路径,留空自动探测
	OracleHome string `json:"oracleHome"` // ORACLE_HOME,留空自动探测
	TNS        string `json:"tns"`        // TNS 别名(如 t35prd),留空按 zone 推导
}

// NamedDB 命名数据库连接(设置页维护的多数据库列表)
type NamedDB struct {
	Name     string `json:"name"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBConfig
}

// NamedEnv 服务器环境:SSH 连接 + 该环境专属的启动参数。
// 一个环境对应一台 T100 服务器(开发/测试/正式),zone/launchArgs/数据库随环境走;
// 设置页以列表维护,activeEnv 单选生效——生效 = 合并覆盖到顶层字段。
type NamedEnv struct {
	Name            string    `json:"name"`
	SSHConfig       // 匿名嵌入:host/port/user/password 提升到 env 层
	Zone            string    `json:"zone,omitempty"`
	TopDir          string    `json:"topDir,omitempty"`
	LaunchArgs      string    `json:"launchArgs,omitempty"`
	WatchdogSeconds int       `json:"watchdogSeconds,omitempty"`
	DB              *DBConfig `json:"db,omitempty"` // TNS/企业覆盖(留空按 zone 推导)
}

// Config debug 功能配置,存放在 config.json 顶层 "debug" 键。
// 与 dbconfig 使用同一文件但互不干扰(各取所需键)。
type Config struct {
	SSH             SSHConfig `json:"ssh"`
	Zone            string    `json:"zone"`            // 登录后区域菜单代码:31开发 35测试 36正式 39PATCH t出货
	Listen          string    `json:"listen"`          // HTTP 监听地址
	LaunchArgs      string    `json:"launchArgs"`      // T100 作业启动参数模板,{prog} 替换为作业名
	WatchdogSeconds int       `json:"watchdogSeconds"` // 停站停留超时(秒),超时自动 continue;0=禁用
	ModuleRoots     []string  `json:"moduleRoots"`     // 源码查找根目录
	FGLServer       string    `json:"fglserver"`       // 留空使用 T100 按 SSH 来源 IP 自动设置
	TopDir          string    `json:"topDir"`          // 区域顶级目录如 /u1/t35prd;留空按 zone 推导
	TermWidth       int       `json:"termWidth"`
	TermHeight      int       `json:"termHeight"`
	PrintElements   int       `json:"printElements"`     // fgldb 单次 print 的数组元素上限(防大数组刷爆);0=默认 1000
	PersistBPs      *bool     `json:"persistBreakpoints"` // 断点持久化开关(nil 视为 true)
	DataDir         string    `json:"-"`               // 数据目录(断点持久化等);由 serve 注入 config.json 所在目录,空=禁用
	DB              *DBConfig `json:"db,omitempty"`    // 数据库连接探查配置(debug db 命令用)
	SSHS            []NamedSSH `json:"sshs,omitempty"` // (兼容保留)旧多 SSH 列表;新配置用 envs
	DBS             []NamedDB  `json:"dbs,omitempty"`  // (兼容保留)旧多数据库列表
	Envs            []NamedEnv `json:"envs,omitempty"` // 服务器环境列表(SSH+启动参数,设置页维护)
	ActiveEnv       string     `json:"activeEnv,omitempty"` // 当前生效的环境名;空=用顶层默认字段
}

// SSHByName 按名取 SSH 连接:先查旧 sshs 列表,再查 envs(取其 SSH 部分);
// 空名/未命中返回默认 ssh
func (c *Config) SSHByName(name string) SSHConfig {
	if name != "" {
		for _, s := range c.SSHS {
			if s.Name == name {
				if s.Port == 0 {
					s.Port = 22
				}
				return s.SSHConfig
			}
		}
		for _, e := range c.Envs {
			if e.Name == name {
				if e.Port == 0 {
					e.Port = 22
				}
				return e.SSHConfig
			}
		}
	}
	return c.SSH
}

// ApplyActiveEnv 把当前生效环境(activeEnv)的连接与启动参数合并覆盖到顶层字段,
// 再重推默认值(TopDir/ModuleRoots 等)。顶层字段即「当前生效配置」,
// manager/hLaunch/CLI 无需感知 envs 的存在;activeEnv 为空或未命中时不做任何事。
func (c *Config) ApplyActiveEnv() {
	if c.ActiveEnv == "" {
		return
	}
	for _, e := range c.Envs {
		if e.Name != c.ActiveEnv {
			continue
		}
		if e.Host != "" {
			c.SSH = e.SSHConfig
			if c.SSH.Port == 0 {
				c.SSH.Port = 22
			}
		}
		if e.Zone != "" {
			c.Zone = e.Zone
		}
		if e.TopDir != "" {
			c.TopDir = e.TopDir
		}
		if e.LaunchArgs != "" {
			c.LaunchArgs = e.LaunchArgs
		}
		if e.WatchdogSeconds > 0 {
			c.WatchdogSeconds = e.WatchdogSeconds
		}
		if e.DB != nil {
			db := *e.DB
			c.DB = &db
		}
		c.fillDefaults()
		return
	}
}

// TNSName 返回数据库 TNS 别名(zone 36→t35prd,35→t35tst,31→t35dev,39→t35pth,t→topprd)
func (c *Config) TNSName() string {
	if c.DB != nil && c.DB.TNS != "" {
		return c.DB.TNS
	}
	switch c.Zone {
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

// DBEnt 返回默认企业编号(配置值)
func (c *Config) DBEnt() int {
	if c.DB != nil {
		return c.DB.Ent
	}
	return 0
}

// BPsPersisted 断点持久化是否启用
func (c *Config) BPsPersisted() bool { return c.DataDir != "" && (c.PersistBPs == nil || *c.PersistBPs) }

// zoneTopDir 区域代码 → T100 顶级目录(默认推导)
var zoneTopDir = map[string]string{
	"31": "/u1/t35dev",
	"35": "/u1/t35tst",
	"36": "/u1/t35prd",
	"39": "/u1/t35pth",
	"t":  "/u1/topprd",
}

func (c *Config) fillDefaults() {
	if c.Zone == "" {
		c.Zone = "36"
	}
	if c.TopDir == "" {
		c.TopDir = zoneTopDir[c.Zone]
	}
	if c.Listen == "" {
		c.Listen = "127.0.0.1:8000"
	}
	if c.LaunchArgs == "" {
		c.LaunchArgs = "BBDL512840855a 2 12345 'N' {prog}"
	}
	if c.WatchdogSeconds == 0 {
		c.WatchdogSeconds = 180
	}
	if len(c.ModuleRoots) == 0 && c.TopDir != "" {
		// com/wss 是 WebService 程序(wssp* / awsp*)的专用模块目录
		c.ModuleRoots = []string{c.TopDir + "/erp", c.TopDir + "/com", c.TopDir + "/com/wss"}
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

// Top 返回区域顶级目录(去尾斜杠)
func (c *Config) Top() string { return c.TopDir }

// CloneWithZone 复制配置并覆盖区域(启动参数 --zone 用):
// 连带推导 TopDir 与 ModuleRoots,其余字段原样共享
func (c *Config) CloneWithZone(zone string) *Config {
	c2 := *c
	if zone != "" && zone != c.Zone {
		c2.Zone = zone
		if td := zoneTopDir[zone]; td != "" {
			c2.TopDir = td
			// com/wss 是 WebService 程序的专用模块目录(与 fillDefaults 同规则)
			c2.ModuleRoots = []string{td + "/erp", td + "/com", td + "/com/wss"}
		}
	}
	return &c2
}

// ModuleDir 返回模块主目录,如 /u1/t35prd/erp/asf;wss 模块挂载在 com 下
func (c *Config) ModuleDir(module string) string {
	if module == "wss" {
		return c.TopDir + "/com/wss"
	}
	return c.TopDir + "/erp/" + module
}

// FGLSOURCEPath 返回 launch 前 export 的源码搜索路径
func (c *Config) FGLSOURCEPath(module string) string {
	base := c.TopDir + "/erp/" + module
	if module == "wss" {
		base = c.TopDir + "/com/wss"
	}
	dirs := []string{
		base + "/4gl",
		base + "/42m",
		c.TopDir + "/com/lib/42m",
		c.TopDir + "/com/sub/42m",
		c.TopDir + "/com/qry/42m",
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

// LoadConfig 从 config.json 读取顶层 "debug" 键并填充默认值
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
		return nil, fmt.Errorf("config.json 缺少 \"debug\" 配置节,请补充 ssh/zone 等字段")
	}
	cfg := wrapper.Debug
	cfg.fillDefaults()
	// 顶层 ssh 允许为空,前提是 envs 里有可用连接(ApplyActiveEnv 合并后自会填充)
	if len(cfg.Envs) == 0 && (cfg.SSH.Host == "" || cfg.SSH.User == "") {
		return nil, fmt.Errorf("debug.ssh.host / debug.ssh.user 未配置")
	}
	return cfg, nil
}

// Addr 返回 SSH 地址 host:port
func (c *SSHConfig) Addr() string { return c.Host + ":" + strconv.Itoa(c.Port) }
