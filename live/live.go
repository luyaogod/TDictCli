// Package live 提供"在线直查 ERP 数据库"的数据访问层。
// 与 db(SQLite 镜像) 同语义:同一套字典表(dzea_t 等)与 JOIN SQL,但直接连远程库执行。
// 目前实现 rt(表字典) 的金仓(PG)方言;oracle 方言与 rv/scc/desc/rq/win 属后续阶段。
//
// 设计要点:
//   - erpdb.Connector 走简单协议(simple protocol),不支持绑定参数 → 所有值经
//     QuoteLit/ValidIdent 白名单后内联,仍保持单条只读 SELECT;
//   - viaSsh 隧道:客户端不可达 DB 时先建 SSH 端口转发,连接 host 换成 127.0.0.1:本地端口。
package live

import (
	"context"
	"fmt"
	"net"

	"tdict/dbconfig"
	"tdict/erpdb"
	"tdict/sshtun"
)

// Live 一个远程数据源:底层连接器 + 可选 SSH 隧道(Close 一并释放)。
type Live struct {
	conn erpdb.Connector
	tun  *sshtun.Tunnel
}

// Dialect 返回底层连接方言("kingbase"|"oracle")。
func (l *Live) Dialect() string { return l.conn.Type() }

// Connector 暴露底层连接器(SelectAllSQL 等通用能力)。
func (l *Live) Connector() erpdb.Connector { return l.conn }

// Close 释放连接;若开了 SSH 隧道一并关闭。
func (l *Live) Close() {
	if l.conn != nil {
		l.conn.Close()
	}
	if l.tun != nil {
		l.tun.Close()
	}
}

// Open 按连接配置建立远程数据源。若配置了 viaSsh,先起 SSH 隧道再连本地转发端口。
func Open(ctx context.Context, c dbconfig.Connection) (*Live, error) {
	effective := c
	if c.ViaSSH != nil {
		remoteHost, remotePort := c.ViaSSH.EffectiveRemote(c.Host, c.Port)
		t, err := sshtun.Dial(sshtun.Options{
			Host:       c.ViaSSH.Host,
			Port:       c.ViaSSH.Port,
			User:       c.ViaSSH.User,
			Password:   c.ViaSSH.Password,
			RemoteHost: remoteHost,
			RemotePort: remotePort,
			LocalPort:  c.ViaSSH.BindPort,
		})
		if err != nil {
			return nil, err
		}
		// 连接器连本地转发端口(host 固定 127.0.0.1)
		effective.Host = "127.0.0.1"
		effective.Port = t.LocalPort
		if effective.Type == "oracle" && effective.ConnectString != "" {
			// oracle 用 connectString 时同样改写:host:port/service → 本地转发
			effective.ConnectString = fmt.Sprintf("%s:%d/%s", "127.0.0.1", t.LocalPort,
				serviceOf(c.ConnectString))
		}
		conn, err := erpdb.Open(ctx, effective)
		if err != nil {
			t.Close()
			return nil, err
		}
		return &Live{conn: conn, tun: t}, nil
	}
	conn, err := erpdb.Open(ctx, effective)
	if err != nil {
		return nil, err
	}
	return &Live{conn: conn}, nil
}

// serviceOf 从 oracle connectString "host:port/service" 中取 service 部分。
func serviceOf(cs string) string {
	for i := len(cs) - 1; i >= 0; i-- {
		if cs[i] == '/' {
			return cs[i+1:]
		}
	}
	return cs
}

// resolveHostPort 备用:把 host:port 拆开(未用时可删)。
func resolveHostPort(addr string) (string, int) {
	h, p, err := net.SplitHostPort(addr)
	if err != nil {
		return addr, 0
	}
	var port int
	fmt.Sscanf(p, "%d", &port)
	return h, port
}
