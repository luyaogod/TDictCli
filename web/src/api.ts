// tdict 配置服务 REST 客户端
export interface DBAcct { account: string; password: string }
export interface DBConnection {
  type: string // oracle | kingbase
  host: string
  port: number
  service?: string // oracle: SERVICE_NAME
  database?: string // kingbase: 库名
  accounts?: DBAcct[]
}
export interface SshEnv {
  name: string
  host: string
  port: number
  user: string
  password: string
  zone?: string
  topent?: string
  db?: DBConnection | null
}
export interface ConfigView {
  configPath: string
  exists: boolean
  activeEnv?: string
  sshs: SshEnv[]
}
// 服务器侧探测结果(host.ProbeDBConfig)
export interface DBProbeOut {
  type: string
  tns?: string
  port?: number
  database?: string
  sqlplus?: string
  oracleHome?: string
  twoTask?: string
  host?: string
  service?: string
  note?: string
}

async function req<T>(url: string, opts?: RequestInit): Promise<T> {
  const r = await fetch(url, { headers: { 'Content-Type': 'application/json' }, ...opts })
  const data = await r.json().catch(() => ({}))
  if (!r.ok || data.ok === false) throw new Error(data.error || `HTTP ${r.status}`)
  return data as T
}

export const api = {
  config: () => req<ConfigView>('/api/config'),
  saveConfig: (cfg: { activeEnv: string; sshs: SshEnv[] }) =>
    req<{ ok: boolean }>('/api/config', { method: 'PUT', body: JSON.stringify(cfg) }),
  // 服务器侧探测连接要素(登录该环境 SSH 只读执行;note 说明未获取到的原因)
  probeDB: (body: { host: string; port: number; user: string; password: string; zone: string; type: string }) =>
    req<DBProbeOut>('/api/dbprobe', { method: 'POST', body: JSON.stringify(body) }),
  // 客户端直连测试(凭据取账号列表首项)
  connTest: (body: DBConnection) =>
    req<{ ok: boolean; version?: string; error?: string }>('/api/conntest', { method: 'POST', body: JSON.stringify(body) }),
  // 账号清单「验证」:SSH 上服务器以该账号+密码连显式目标库 select 1(只读)
  dbAccVerify: (body: {
    host: string; port: number; user: string; password: string; zone: string; type: string
    account: string; acctPassword: string; dbHost?: string; dbPort?: number; dbSvc?: string; dbDatabase?: string
  }) => req<{ ok: boolean; error?: string }>('/api/dbaccverify', { method: 'POST', body: JSON.stringify(body) }),
}
