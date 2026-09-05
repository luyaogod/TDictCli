package cli

// db 连接管理:tdict db list / ping / discover。
// list    :列出 connections(含 source/viaSsh 标记与默认项)
// ping    :验证给定连接可达(含 viaSsh 隧道;只读 SELECT version)
// discover:SSH 自动发现某 debug 环境(env)的数据库连接要素 → 预览/--save 写入 connections
//          (oracle:chenv zone+ORACLE_HOME+TNS/tnsnames;kingbase:实例发现)

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"tdict/dbconfig"
	"tdict/debug"
	"tdict/live"

	"github.com/spf13/cobra"
)

var (
	discoverEnv  string
	discoverType string
	discoverHost string
	discoverUser string
	discoverSave bool
	discoverConn string
)

// dbListCmd 列出连接。
var dbListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出 ERP 数据库连接(含来源/隧道标记)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if IsJSON() {
			return printJSON(dbConfig.Connections)
		}
		fmt.Printf("连接(%d):\n", len(dbConfig.Connections))
		for i := range dbConfig.Connections {
			c := &dbConfig.Connections[i]
			mark := "  "
			if c.IsDefault {
				mark = " *"
			}
			via := ""
			if c.ViaSSH != nil {
				via = "  [viaSSH→" + c.ViaSSH.Host + "]"
			}
			src := ""
			if c.Source != "" {
				src = "  (source=" + c.Source + ")"
			}
			fmt.Printf("  %s %-12s %-8s %-28s%s%s\n", mark, c.Name, c.Type, c.Address(), via, src)
		}
		return nil
	},
}

// dbPingCmd 验证连接可达。
var dbPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "验证连接可达(只读 SELECT version;支持 viaSsh 隧道)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := resolveConnection(dbSyncConn)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
		defer cancel()
		l, err := live.Open(ctx, *conn)
		if err != nil {
			return err
		}
		defer l.Close()
		ver, err := l.Connector().ServerVersion(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("连接正常: %s (%s) %s\n", conn.Name, conn.Type, firstLine(ver))
		return nil
	},
}

func firstLine(s string) string {
	if i := strings.IndexAny(s, "\r\n"); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return strings.TrimSpace(s)
}

// dbDiscoverCmd SSH 自动发现 → 预览/保存连接。
var dbDiscoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "SSH 自动发现数据库连接要素并保存为连接(--env <环境名>)",
	Long: `登录 debug 环境(env)的 SSH,按区域探测数据库连接要素:
  oracle  → chenv zone + ORACLE_HOME/sqlplus + TNS 别名 + tnsnames 解析 host/port/service
  kingbase→ 实例发现(ps 找 kingbase + ksql + kingbase.conf 端口 + 库名)
输出候选连接(不落盘);--save 写入 config.json connections(source=ssh:<env>)。
注意:自动发现给出的是"服务器视角"要素;客户端直连地址(host)需自行确认,
若 DB 仅服务器可达,保存后再补 viaSsh 配置。`,
	Example: `  tdict db discover --env 主机正式区                 # oracle(默认),预览
  tdict db discover --env 恒烁正式区 --type kingbase
  tdict db discover --env 恒烁正式区 --type kingbase --save`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ssh, zone, envName, err := discoverSSHFromEnv(discoverEnv)
		if err != nil {
			return err
		}
		typ := discoverType
		if typ == "" {
			typ = "oracle"
		}
		req := debug.DBProbeReq{Host: ssh.Host, Port: ssh.Port, User: ssh.User,
			Password: ssh.Password, Zone: zone, Type: typ}
		out, err := debug.ProbeDBConfig(req)
		if err != nil {
			return err
		}
		cand := candidateConn(out, envName)
		if IsJSON() {
			return printJSON(cand)
		}
		fmt.Printf("自动发现(SSH %s, env=%s):\n", ssh.Host, envName)
		fmt.Printf("  类型: %s", cand.Type)
		if cand.Type == "oracle" {
			fmt.Printf("  TNS=%s  Host=%s:%d  Service=%s\n",
				out.TNS, cand.Host, cand.Port, cand.Database)
		} else {
			fmt.Printf("  Host=%s:%d  Database=%s\n", cand.Host, cand.Port, cand.Database)
		}
		if out.Note != "" {
			fmt.Printf("  提示: %s\n", out.Note)
		}
		fmt.Println("  说明: host 为服务器视角候选地址,客户端不可达时请改 host 或补 viaSsh")
		if !discoverSave {
			fmt.Println("  保存: 加 --save 写入 config.json connections")
			return nil
		}
		if err := saveDiscoveredConn(cand, envName); err != nil {
			return err
		}
		fmt.Printf("已保存连接 %q(source=ssh:%s)\n", cand.Name, envName)
		return nil
	},
}

// discoverSSHFromEnv 从 debug.envs 找到匹配环境并返回 SSH+zone。
func discoverSSHFromEnv(name string) (debug.SSHConfig, string, string, error) {
	cfgPath, err := resolveConfigPath(configPath)
	if err != nil {
		return debug.SSHConfig{}, "", "", err
	}
	cfg, err := debug.LoadConfig(cfgPath)
	if err != nil {
		return debug.SSHConfig{}, "", "", err
	}
	if name != "" {
		for _, e := range cfg.Envs {
			if e.Name == name {
				s := e.SSHConfig
				if s.Port == 0 {
					s.Port = 22
				}
				return s, e.Zone, e.Name, nil
			}
		}
		return debug.SSHConfig{}, "", "", fmt.Errorf("未找到环境 %q(可 tdict debug env 查看)", name)
	}
	if cfg.SSH.Host != "" {
		s := cfg.SSH
		if s.Port == 0 {
			s.Port = 22
		}
		return s, cfg.Zone, cfg.EnvName(), nil
	}
	return debug.SSHConfig{}, "", "", fmt.Errorf("未指定 --env 且无顶层 ssh 配置")
}

// candidateConn 由探测结果构造候选连接(host 以探测/ssh host 兜底)。
// user 按 T100 惯例(账号=密码)填充;未指定则留空待手填。
func candidateConn(out *debug.DBProbeOut, envName string) *dbconfig.Connection {
	name := discoverConn
	if name == "" {
		name = "auto-" + envName + "-" + out.Type
	}
	c := &dbconfig.Connection{Name: name, Type: out.Type, Source: "ssh:" + envName}
	if discoverUser != "" {
		c.User = discoverUser
		c.Password = discoverUser // T100 fglprofile 明文规则:账号=密码
	}
	if out.Type == "kingbase" {
		c.Host = discoverHost
		c.Port = out.Port
		c.Database = out.Database
		if c.Port == 0 {
			c.Port = 54321
		}
	} else {
		c.Host = out.Host
		c.Port = out.Port
		c.Database = out.Service // oracle: service_name
		c.ConnectString = fmt.Sprintf("%s:%d/%s", out.Host, out.Port, out.Service)
	}
	if c.Host == "" {
		c.Host = discoverHost
	}
	return c
}

// saveDiscoveredConn 把连接写回 config.json top.connections(保留其它顶层键)。
func saveDiscoveredConn(c *dbconfig.Connection, envName string) error {
	cfgPath, err := resolveConfigPath(configPath)
	if err != nil {
		return err
	}
	raw, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}
	var root map[string]any
	if err := json.Unmarshal(raw, &root); err != nil || root == nil {
		root = map[string]any{}
	}
	var conns []any
	if old, ok := root["connections"].([]any); ok {
		conns = old
	} else if old2, ok := root["connections"].([]map[string]any); ok {
		for _, m := range old2 {
			conns = append(conns, m)
		}
	}
	// 同 source 同 env 覆盖,否则追加
	replaced := false
	for i := range conns {
		if m, ok := conns[i].(map[string]any); ok && m["source"] == "ssh:"+envName {
			conns[i] = connToAny(c)
			replaced = true
			break
		}
	}
	if !replaced {
		conns = append(conns, connToAny(c))
	}
	root["connections"] = conns
	out, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return err
	}
	tmp := cfgPath + ".tmp"
	if err := os.WriteFile(tmp, out, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, cfgPath)
}

func connToAny(c *dbconfig.Connection) map[string]any {
	b, _ := json.Marshal(c)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	return m
}

func init() {
	dbPingCmd.Flags().StringVarP(&dbSyncConn, "conn", "c", "", "连接名称(默认 isDefault 连接)")
	dbDiscoverCmd.Flags().StringVar(&discoverEnv, "env", "", "debug 环境名(envs;空=顶层 ssh)")
	dbDiscoverCmd.Flags().StringVar(&discoverType, "type", "", "数据库类型(oracle/kingbase;默认 oracle)")
	dbDiscoverCmd.Flags().StringVar(&discoverHost, "host", "", "客户端可达 DB 地址覆盖(默认取探测/ssh host)")
	dbDiscoverCmd.Flags().StringVar(&discoverUser, "user", "", "数据库账号(默认套用 T100 规则:账号=密码;空则保存后手填)")
	dbDiscoverCmd.Flags().StringVar(&discoverConn, "conn", "", "保存的连接名(默认 auto-<env>-<type>)")
	dbDiscoverCmd.Flags().BoolVar(&discoverSave, "save", false, "写入 config.json connections")
	dbCmd.AddCommand(dbListCmd, dbPingCmd, dbDiscoverCmd)
}
