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

// 源码镜像(host.MirrorPullProgress 的进度经 /api/mirror 轮询返回)
export interface MirrorEnv {
  name: string
  zone: string
  path: string
  ready: boolean // 本地已有完整基线(增量前提)
}
export interface MirrorJob {
  running: boolean
  env: string
  full: boolean
  phase: string // connect|probe|pack|download|done|error
  message: string
  bytes: number
  total: number // 下载阶段为归档总字节;0=未知(pack 阶段)
  files: number
  elapsed: string
  error?: string
  note?: string
  done: boolean
  startedAt?: string
}
export interface MirrorResp {
  mirrorDir: string
  activeEnv: string
  envs: MirrorEnv[]
  job: MirrorJob
}

// 数据库同步(dbsync.Run 的进度经 /api/dbsync 轮询返回)
export interface DBSyncEnv {
  name: string
  type: string
  address: string
}
export interface DBSyncJob {
  running: boolean
  env: string
  phase: string // open|table|index|replace|done|error
  message: string
  table: string
  tableIndex: number
  tableTotal: number
  tableRows: number
  totalRows: number
  tables: number
  elapsed: string
  target: string
  backup?: string
  warning?: string
  error?: string
  done: boolean
  startedAt?: string
}
export interface DBSyncResp {
  target: string
  configured: string
  defaultTarget: string
  exists: boolean
  activeEnv: string
  envs: DBSyncEnv[]
  job: DBSyncJob
}

// 命令行安装:把 tdict 所在目录加入用户 PATH
export interface InstallStatus {
  supported: boolean // 本平台是否支持自动写入
  exePath: string
  exeDir: string
  inUserPath: boolean
  userPath?: string
  manual?: string
  note?: string
}

// BDL(4GL)语言文档目录(等价 tdict bdldoc dir)
export interface BdldocStatus {
  ok: boolean
  dir: string
  exists: boolean
  configPath: string
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
  // 源码镜像:根目录读写 + 拉取任务(单实例)
  mirror: () => req<MirrorResp>('/api/mirror'),
  saveMirrorDir: (dir: string) =>
    req<{ ok: boolean; mirrorDir?: string }>('/api/mirror', { method: 'PUT', body: JSON.stringify({ dir }) }),
  mirrorPull: (env: string, full: boolean) =>
    req<{ ok: boolean }>('/api/mirror/pull', { method: 'POST', body: JSON.stringify({ env, full }) }),
  // 数据库同步:远程库字典 → 本地 SQLite(单实例任务)
  dbsync: () => req<DBSyncResp>('/api/dbsync'),
  dbsyncRun: (env: string) =>
    req<{ ok: boolean }>('/api/dbsync', { method: 'POST', body: JSON.stringify({ env }) }),
  // 设置同步目标(config.json 顶层 sync.target;空串=清除,回到默认 exe 同目录)
  saveDBSyncTarget: (target: string) =>
    req<DBSyncResp>('/api/dbsync', { method: 'PUT', body: JSON.stringify({ target }) }),
  status: () => req<{ ok: boolean; server: string; listen: string }>('/api/status'),
  // 命令行安装:查看/加入/移出用户 PATH(用户级,无需管理员)
  installStatus: () => req<InstallStatus>('/api/install'),
  installAdd: () => req<InstallStatus>('/api/install', { method: 'POST' }),
  installRemove: () => req<InstallStatus>('/api/install', { method: 'DELETE' }),
  // BDL 语言文档目录(config.json 顶层 bdldoc.dir)
  bdldoc: () => req<BdldocStatus>('/api/bdldoc'),
  saveBdldocDir: (dir: string) =>
    req<BdldocStatus>('/api/bdldoc', { method: 'PUT', body: JSON.stringify({ dir }) }),
}
