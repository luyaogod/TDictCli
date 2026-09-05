package debug

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// EntValue 企业编号(TOPENT):数字或文本均可,兼容旧配置的 JSON 数字。
// 数据库探测等需要真实编号的场景用 Int()(非数字返回 false)
type EntValue string

// UnmarshalJSON 同时接受 JSON 数字与字符串(旧配置 "ent": 99 / 新文本 "ent": "99x")
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

// DBConfig 数据库连接探查配置(config.json debug.db 节,全部可选)
type DBConfig struct {
	Type       string   `json:"type"`           // 数据库类型:"oracle"(默认)| "kingbase"(人大金仓,PG 引擎)
	Ent        EntValue `json:"ent"`            // 默认企业编号(TOPENT);空=不指定;数字或文本均可
	SQLPlus    string   `json:"sqlplus"`        // oracle: sqlplus 路径,留空自动探测
	OracleHome string   `json:"oracleHome"`     // oracle: ORACLE_HOME,留空自动探测
	TNS        string   `json:"tns"`            // oracle: TNS 别名(如 t35prd)/ kingbase: 库名,留空自动发现实例
	Port       int      `json:"port,omitempty"` // kingbase: 实例端口,0=自动发现(默认 54321)
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
	SSHConfig                 // 匿名嵌入:host/port/user/password 提升到 env 层
	Zone            string    `json:"zone,omitempty"`
	TopDir          string    `json:"topDir,omitempty"`
	LaunchArgs      string    `json:"launchArgs,omitempty"`
	WatchdogSeconds int       `json:"watchdogSeconds,omitempty"`
	DB              *DBConfig `json:"db,omitempty"` // TNS/企业覆盖(留空按 zone 推导)
}

// DefaultListen 是 serve 未显式配置监听地址时的默认地址。
// 选用不常用端口(28670),降低与其它开发服务(8000/8080/3000 等)冲突的概率;
// 若仍被占用,serve 启动时会自动顺延到下一个空闲端口。
const DefaultListen = "127.0.0.1:28670"

// Config debug 功能配置,存放在 config.json 顶层 "debug" 键。
// 与 dbconfig 使用同一文件但互不干扰(各取所需键)。
type Config struct {
	SSH             SSHConfig   `json:"ssh"`
	Zone            string      `json:"zone"`            // 登录后区域菜单代码:31开发 35测试 36正式 39PATCH t出货
	Listen          string      `json:"listen"`          // HTTP 监听地址
	LaunchArgs      string      `json:"launchArgs"`      // T100 作业启动参数模板,{prog} 替换为作业名
	WatchdogSeconds int         `json:"watchdogSeconds"` // 停站停留超时(秒),超时自动 continue;0=禁用
	ModuleRoots     []string    `json:"moduleRoots"`     // 源码查找根目录
	FGLServer       string      `json:"fglserver"`       // 留空使用 T100 按 SSH 来源 IP 自动设置
	TopDir          string      `json:"topDir"`          // 区域顶级目录如 /u1/t35prd;留空按 zone 推导
	TermWidth       int         `json:"termWidth"`
	TermHeight      int         `json:"termHeight"`
	PrintElements   int         `json:"printElements"`       // fgldb 单次 print 的数组元素上限(防大数组刷爆);0=默认 1000
	PersistBPs      *bool       `json:"persistBreakpoints"`  // 断点持久化开关(nil 视为 true)
	DataDir         string      `json:"-"`                   // 数据目录(断点持久化等);由 serve 注入 config.json 所在目录,空=禁用
	DB              *DBConfig   `json:"db,omitempty"`        // 数据库连接探查配置(debug db 命令用)
	SSHS            []NamedSSH  `json:"sshs,omitempty"`      // (兼容保留)旧多 SSH 列表;新配置用 envs
	DBS             []NamedDB   `json:"dbs,omitempty"`       // (兼容保留)旧多数据库列表
	Envs            []NamedEnv  `json:"envs,omitempty"`      // 服务器环境列表(SSH+启动参数,设置页维护)
	ActiveEnv       string      `json:"activeEnv,omitempty"` // 当前生效的环境名;空=用顶层默认字段
	Runtime         *RuntimeEnv `json:"-"`                   // 登录后动态获取的 T100 路径(探针/选区回显);nil=用静态配置兜底
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
		// 换环境 = 换服务器/区域:清动态路径,下次探针/选区回显重新获取
		c.Runtime = nil
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

// EnvName 会话所属环境名:优先设置页 envs 的 activeEnv,缺省用 host-zone 推导名
func (c *Config) EnvName() string {
	if c.ActiveEnv != "" {
		return c.ActiveEnv
	}
	return c.SSH.Host + "-" + c.Zone
}

// CloneEnv 复制配置并切换到指定环境(name 命中 Envs 之一):应用该环境的
// SSH/zone/启动参数/看门狗/库配置,用于会话「切换/重启」按目标环境重连;未命中返回 nil。
func (c *Config) CloneEnv(name string) *Config {
	if name == "" {
		return nil
	}
	var hit *NamedEnv
	for i := range c.Envs {
		if c.Envs[i].Name == name {
			hit = &c.Envs[i]
			break
		}
	}
	if hit == nil {
		return nil
	}
	c2 := *c
	if hit.Host != "" {
		c2.SSH = hit.SSHConfig
		if c2.SSH.Port == 0 {
			c2.SSH.Port = 22
		}
	}
	if hit.Zone != "" {
		c2.Zone = hit.Zone
	}
	c2.ActiveEnv = name
	c2.Runtime = nil // 环境/服务器变了,动态路径需重新获取
	if hit.LaunchArgs != "" {
		c2.LaunchArgs = hit.LaunchArgs
	}
	if hit.WatchdogSeconds > 0 {
		c2.WatchdogSeconds = hit.WatchdogSeconds
	}
	if hit.DB != nil {
		db := *hit.DB
		c2.DB = &db
	}
	c2.fillDefaults()
	return &c2
}

// TNSName 返回数据库 TNS 别名(zone 36→t35prd,35→t35tst,31→t35dev,39→t35pth,t→topprd)。
// 完全自动:按登录区域推导,不接受手填覆盖(T100 环境约定)
func (c *Config) TNSName() string {
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

// DBEnt 返回默认企业编号(数据库探测用,须为数字;文本/未配置返回 0=仅列映射)
func (c *Config) DBEnt() int {
	if c.DB == nil {
		return 0
	}
	n, _ := c.DB.Ent.Int()
	return n
}

// BPsPersisted 断点持久化是否启用
func (c *Config) BPsPersisted() bool {
	return c.DataDir != "" && (c.PersistBPs == nil || *c.PersistBPs)
}

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
		c.Listen = DefaultListen
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

// TopDirActual 返回生效的区域顶级目录:动态获取(登录后 TOP)优先,静态配置兜底
func (c *Config) TopDirActual() string {
	if c.Runtime != nil && c.Runtime.TOP != "" {
		return c.Runtime.TOP
	}
	return c.TopDir
}

// ModuleRootsActual 返回生效的源码查找根目录:动态获取(ERP/COM)优先,静态配置兜底
func (c *Config) ModuleRootsActual() []string {
	if c.Runtime != nil && c.Runtime.ERP != "" {
		// com/wss 是 WebService 程序(wssp* / awsp*)的专用模块目录
		return []string{c.Runtime.ERP, c.Runtime.COM, c.Runtime.COM + "/wss"}
	}
	return c.ModuleRoots
}

// Top 返回区域顶级目录(去尾斜杠)
func (c *Config) Top() string { return c.TopDirActual() }

// CloneWithZone 复制配置并覆盖区域(启动参数 --zone 用):
// 连带推导 TopDir 与 ModuleRoots,其余字段原样共享
func (c *Config) CloneWithZone(zone string) *Config {
	c2 := *c
	if zone != "" && zone != c.Zone {
		c2.Zone = zone
		c2.Runtime = nil // 区域变了,动态路径需重新获取
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
