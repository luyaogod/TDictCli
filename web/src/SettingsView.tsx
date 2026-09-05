// 设置页:VS Code 式左侧一级分类(外观/环境/高级)。
// 「环境」:左侧环境(SSH 服务器)列表 + 右侧表单区拆两个 Tab —— 「SSH 服务器」与
// 「数据库」一对一编辑同一个环境:SSH Tab 维护登录连接/区域/TOPENT;DB Tab 维护该
// 环境的数据库连接(显式 host/port/service|库名 + 账号列表,无主账号)。
// 账号语义:运行时由 TOPENT 决定(服务器 gzou_t 解析账号名,密码查账号列表);
// 客户端直连(ping/在线查询/连接测试)取账号列表首项。
import { useCallback, useEffect, useRef, useState, type ReactNode } from 'react'
import { CircleCheck, Circle, Plus, Trash2, Monitor, Server, Database, SlidersHorizontal, Sun, Search, Eye, EyeOff, RefreshCw, PlugZap } from 'lucide-react'
import { api } from './api'
import { Button, Input, Separator } from './ui'
import { useStore } from './store'

interface SshDb {
  type: string // oracle | kingbase
  host: string
  port: number
  service: string // oracle: SERVICE_NAME
  database: string // kingbase: 库名
  accounts: { account: string; password: string }[]
}

interface SshItem {
  name: string
  host: string
  port: number
  user: string
  password: string
  zone: string
  topDir: string
  topent: string
  db: SshDb | null
}

type Section = 'appearance' | 'envs' | 'advanced'
type EnvTab = 'ssh' | 'db'

const SECTIONS: { key: Section; label: string; icon: typeof Monitor }[] = [
  { key: 'envs', label: '环境', icon: Server },
  { key: 'appearance', label: '外观', icon: Monitor },
  { key: 'advanced', label: '高级', icon: SlidersHorizontal },
]

const input = 'h-7 text-xs'
const cell = 'h-7 w-full min-w-0 text-xs'

const blankDb = (): SshDb => ({ type: 'oracle', host: '', port: 1521, service: '', database: '', accounts: [] })
const blankSsh = (): SshItem => ({ name: '', host: '', port: 22, user: '', password: '', zone: '36', topDir: '', topent: '', db: null })

export function SettingsView() {
  const theme = useStore((s) => s.theme)
  const setTheme = useStore((s) => s.setTheme)
  const [section, setSection] = useState<Section>('envs')
  const [envTab, setEnvTab] = useState<EnvTab>('ssh')
  const [cfg, setCfg] = useState<any>(null)
  const [sshs, setSshs] = useState<SshItem[]>([])
  const [activeEnv, setActiveEnv] = useState('')
  const [selSsh, setSelSsh] = useState(0)
  const [err, setErr] = useState('')
  const [saveState, setSaveState] = useState<'idle' | 'saving' | 'saved'>('idle')
  const dirtyRef = useRef(false)
  const [dirty, setDirty] = useState(false)
  const markDirty = () => { dirtyRef.current = true; setDirty(true) }
  const cleanDirty = () => { dirtyRef.current = false; setDirty(false) }
  const [showPwd, setShowPwd] = useState(false)
  // DB Tab 操作状态
  const [busy, setBusy] = useState<string>('')
  const [note, setNote] = useState<{ kind: 'ok' | 'err' | 'info'; text: string } | null>(null)
  const [accProbe, setAccProbe] = useState<{ i: number; state: 'testing' | 'ok' | 'err'; msg: string } | null>(null)
  const [addAcct, setAddAcct] = useState({ account: '', password: '' })

  const sshName = (e: SshItem) => e.name || `${e.host}-${e.zone}`.replace(/-$/, '')

  const loadSettings = useCallback(() => {
    api.settings().then((c) => {
      setCfg(c)
      setActiveEnv(c.activeEnv || '')
      setSshs((c.sshs || []).map((e: any) => {
        const d = e.db
        return {
          name: e.name || '', host: e.host || '', port: e.port || 22, user: e.user || '', password: e.password || '',
          zone: e.zone || '', topDir: e.topDir || '', topent: e.topent != null ? String(e.topent) : '',
          db: d ? {
            type: d.type || 'oracle', host: d.host || '', port: d.port || 0,
            service: d.service || '', database: d.database || '',
            accounts: (d.accounts || []).map((a: any) => ({ account: a.account || '', password: a.password || '' })),
          } : null,
        }
      }))
      const names = (c.sshs || []).map((e: any) => e.name)
      if (c.activeEnv && !names.includes(c.activeEnv)) setActiveEnv('')
      const idx = names.indexOf(c.activeEnv || '')
      if (idx >= 0) setSelSsh(idx)
      cleanDirty(); setSaveState('idle'); setErr('')
    }).catch((e) => setErr(e.message))
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => { loadSettings() }, [loadSettings])
  // keep-alive 常驻挂载:再次切到「环境」时回读(无未保存修改时)
  useEffect(() => {
    if (section !== 'envs' || dirtyRef.current) return
    void loadSettings()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [section])

  // ---- 统一保存(整环境:SSH 字段 + db) ----
  const saveAll = async () => {
    if (saveState === 'saving' || !cfg) return
    setErr(''); setSaveState('saving')
    try {
      const next = {
        ...cfg,
        activeEnv: activeEnv || (sshs[0] ? sshName(sshs[0]) : ''),
        sshs: sshs.filter((x) => x.host).map((x) => {
          const o: Record<string, any> = {
            name: sshName(x), host: x.host, port: x.port || 22, user: x.user, password: x.password,
          }
          if (x.zone.trim()) o.zone = x.zone.trim()
          if (x.topDir.trim()) o.topDir = x.topDir.trim()
          if (x.topent.trim()) o.topent = x.topent.trim()
          if (x.db) {
            const db: Record<string, any> = { type: x.db.type || 'oracle', host: x.db.host.trim(), port: x.db.port || 0 }
            if (x.db.type === 'oracle') { if (x.db.service.trim()) db.service = x.db.service.trim() }
            else { if (x.db.database.trim()) db.database = x.db.database.trim() }
            const accounts = x.db.accounts.map((a) => ({ account: a.account.trim(), password: a.password })).filter((a) => a.account)
            if (accounts.length) db.accounts = accounts
            o.db = db
          }
          return o
        }),
      }
      await api.saveSettings(next)
      setCfg(next)
      setSshs((prev) => prev.map((x) => ({ ...x, name: sshName(x) })))
      cleanDirty(); setSaveState('saved')
    } catch (ex: any) {
      setErr(ex.message || String(ex)); setSaveState('idle')
    }
  }
  const patchSsh = (i: number, patch: Partial<SshItem>) => {
    const next = sshs.map((x, j) => (j === i ? { ...x, ...patch } : x))
    let na: string | undefined
    if (patch.name !== undefined && sshs[i]?.name === activeEnv) na = patch.name
    if (na !== undefined) setActiveEnv(na)
    setSshs(next); markDirty()
  }
  const patchDb = (i: number, patch: Partial<SshDb>) => {
    const cur = sshs[i]
    if (!cur) return
    patchSsh(i, { db: cur.db ? { ...cur.db, ...patch } : { ...blankDb(), ...patch } })
  }
  const patchAcct = (i: number, j: number, patch: Partial<{ account: string; password: string }>) => {
    const cur = sshs[i]
    if (!cur?.db) return
    patchDb(i, { accounts: cur.db.accounts.map((a, k) => (k === j ? { ...a, ...patch } : a)) })
  }
  const delAcct = (i: number, j: number) => {
    const cur = sshs[i]
    if (!cur?.db) return
    patchDb(i, { accounts: cur.db.accounts.filter((_, k) => k !== j) })
  }
  const pushAcct = () => {
    const cur = sshs[selSsh]
    if (!cur?.db) return
    const account = addAcct.account.trim()
    if (!account) return
    patchDb(selSsh, { accounts: [...cur.db.accounts, { account, password: addAcct.password }] })
    setAddAcct({ account: '', password: '' })
  }
  const addSsh = () => { setSelSsh(sshs.length); setSshs([...sshs, blankSsh()]); markDirty() }
  const delSsh = () => {
    const cur = sshs[selSsh]
    const k = cur ? sshName(cur) : ''
    setSelSsh(Math.max(0, selSsh - 1))
    setSshs(sshs.filter((_, j) => j !== selSsh))
    if (k === activeEnv) setActiveEnv('')
    markDirty()
  }

  // ---- DB Tab 操作(针对当前环境的 db) ----
  // 客户端直连测试:凭据 = 账号列表首项
  const testConn = async () => {
    const d = sshs[selSsh]?.db
    if (!d?.host.trim()) { setNote({ kind: 'err', text: '请先填写数据库主机地址' }); return }
    if (!d.accounts.length || !d.accounts[0].account.trim()) { setNote({ kind: 'err', text: '请先在账号列表添加账号(直连取列表首项)' }); return }
    if (d.type === 'oracle' && !d.service.trim()) { setNote({ kind: 'err', text: 'Oracle 需填写服务名(service)' }); return }
    if (d.type === 'kingbase' && !d.database.trim()) { setNote({ kind: 'err', text: '金仓需填写库名(database)' }); return }
    setBusy('test'); setNote(null)
    try {
      const r = await api.connTest({
        type: d.type, host: d.host.trim(), port: d.port || 0,
        service: d.type === 'oracle' ? d.service.trim() : undefined,
        database: d.type === 'kingbase' ? d.database.trim() : undefined,
        accounts: d.accounts,
      })
      setNote({ kind: 'ok', text: `连接正常 ✓ ${r.version || ''}(账号 ${d.accounts[0].account})` })
    } catch (ex: any) {
      setNote({ kind: 'err', text: '连接失败: ' + (ex.message || String(ex)) })
    } finally { setBusy('') }
  }
  // 从服务器获取:登录当前环境的 SSH 自动探测连接要素回填(辅助)
  const fetchFromServer = async () => {
    const s = sshs[selSsh]
    const d = s?.db
    if (!s?.host || !s?.user) { setNote({ kind: 'err', text: '请先填写 SSH 主机与账号' }); return }
    if (!d) return
    setBusy('fetch'); setNote(null)
    try {
      const r = await api.probeDB({ host: s.host, port: s.port || 22, user: s.user, password: s.password, zone: s.zone, type: d.type || 'oracle' })
      const patch: Partial<SshDb> = {}
      if (r.type === 'kingbase') {
        if (r.database) patch.database = r.database
        if (r.port) patch.port = r.port
        setNote(r.note ? { kind: 'info', text: r.note } : { kind: 'ok', text: `获取成功:金仓 ${r.database || '?'} @ ${s.host}:${r.port || '?'}(服务器视角,地址按需调整)` })
      } else {
        if (r.service || r.tns) patch.service = r.service || r.tns || ''
        if (r.port) patch.port = r.port
        setNote({ kind: r.note ? 'err' : 'info', text: (r.note ? r.note + ';' : '') + `服务器解析地址 ${r.host || '未知'}:${r.port || '?'}(service ${r.service || r.tns || '?'})。host 通常与 SSH 主机一致,客户端不可达时手工调整` })
      }
      if (Object.keys(patch).length) patchDb(selSsh, patch)
    } catch (ex: any) {
      setNote({ kind: 'err', text: '获取失败: ' + (ex.message || String(ex)) })
    } finally { setBusy('') }
  }
  // 附加账号行验证:服务器侧以 账号+密码 连当前环境的显式目标库
  const verifyAcct = async (j: number) => {
    const i = selSsh
    const s = sshs[i]
    const d = s?.db
    const row = d?.accounts?.[j]
    if (!s?.host || !s?.user) { setAccProbe({ i: j, state: 'err', msg: '请先填写 SSH 主机与账号' }); return }
    if (!d?.host.trim()) { setAccProbe({ i: j, state: 'err', msg: '请先填写数据库主机地址' }); return }
    if (!row?.account.trim()) { setAccProbe({ i: j, state: 'err', msg: '账号不能为空' }); return }
    setAccProbe({ i: j, state: 'testing', msg: '' })
    try {
      await api.dbAccVerify({
        host: s.host, port: s.port || 22, user: s.user, password: s.password, zone: s.zone,
        type: d.type || 'oracle', account: row.account.trim(), acctPassword: row.password,
        dbHost: d.host.trim(), dbPort: d.port || 0,
        ...(d.type === 'kingbase' ? { dbDatabase: d.database.trim() } : { dbSvc: d.service.trim() }),
      })
      setAccProbe({ i: j, state: 'ok', msg: '连接正常 ✓' })
    } catch (ex: any) {
      setAccProbe({ i: j, state: 'err', msg: ex.message || String(ex) })
    }
  }
  // 「已保存」提示短暂停留后回到空闲
  useEffect(() => {
    if (saveState !== 'saved') return
    const t = window.setTimeout(() => setSaveState('idle'), 2000)
    return () => window.clearTimeout(t)
  }, [saveState])

  if (!cfg) return <div className="p-6 text-sm text-muted-foreground">{err || '加载配置中…'}</div>
  const cur = sshs[selSsh]
  const curDb = cur?.db
  return (
    <div className="flex h-full min-h-0 text-xs">
      {/* VS Code 式左侧一级分类 */}
      <div className="w-36 shrink-0 space-y-0.5 overflow-auto border-r border-border p-2">
        {SECTIONS.map(({ key, label, icon: Icon }) => (
          <button key={key} onClick={() => setSection(key)}
            className={`flex w-full items-center gap-2 px-2 py-1.5 text-left transition-colors ${
              section === key ? 'bg-accent text-accent-foreground font-medium' : 'text-muted-foreground hover:bg-accent/60'
            }`}>
            <Icon className="h-3.5 w-3.5" />{label}
          </button>
        ))}
      </div>

      {/* 右侧内容区 */}
      <div className="min-h-0 flex-1 overflow-auto">
        <div className="mx-auto max-w-3xl p-4">
          <div className="mb-2 flex items-center justify-between">
            <h2 className="text-sm font-medium text-foreground">环境设置</h2>
            <span className="flex items-center gap-2">
              {err && <span className="text-[11px] text-red-600 dark:text-red-400">{err}</span>}
              <span className={`text-[11px] ${saveState === 'saved' ? 'text-emerald-600 dark:text-emerald-400' : dirty ? 'text-amber-600 dark:text-amber-400' : ''}`}>
                {saveState === 'saving' ? '保存中…' : saveState === 'saved' ? '已保存 ✓' : dirty ? '有未保存的修改' : ''}
              </span>
            </span>
          </div>

          {section === 'envs' && (
            <div className="flex">
              {/* 环境列表(左侧;右侧表单区分 SSH/DB 两 Tab 编辑同一环境) */}
              <div className="w-44 shrink-0 border-r border-border pr-1.5">
                <div className="py-1">
                  {sshs.length === 0 && <div className="p-2 text-muted-foreground">(空)</div>}
                  {sshs.map((e, i) => (
                    <button key={i} onClick={() => setSelSsh(i)}
                      className={`flex w-full items-center gap-1.5 px-2 py-1.5 text-left transition-colors ${
                        selSsh === i ? 'bg-accent text-accent-foreground' : 'hover:bg-accent/60'
                      }`}>
                      {activeEnv === e.name
                        ? <CircleCheck className="h-3.5 w-3.5 shrink-0 text-emerald-600 dark:text-emerald-400" />
                        : <Circle className="h-3.5 w-3.5 shrink-0 text-muted-foreground/40" />}
                      <span className="min-w-0 flex-1">
                        <span className="block truncate font-medium">{e.name || e.host || '(新环境)'}</span>
                        <span className="block truncate text-muted-foreground">
                          {e.zone ? `${e.zone} · ` : ''}{e.port || 22}{e.db ? ' · 库' : ''}
                        </span>
                      </span>
                    </button>
                  ))}
                </div>
                <Button size="sm" variant="outline" className="mt-2 w-full" onClick={addSsh}>
                  <Plus className="mr-1 h-3 w-3" />新增环境
                </Button>
              </div>

              {cur && (
                <section className="min-w-0 flex-1 pl-3">
                  <div className="mb-2 flex items-center justify-between">
                    <h3 className="font-medium text-foreground">
                      环境参数{activeEnv === (cur.name || sshName(cur)) && <span className="ml-2 text-emerald-600 dark:text-emerald-400">(默认)</span>}
                    </h3>
                    <div className="flex gap-1.5">
                      <Button size="sm" variant="ghost" className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400"
                        title="删除该环境(保存后生效)" onClick={delSsh}>
                        <Trash2 className="h-3.5 w-3.5" />
                      </Button>
                      <Button size="sm" variant="secondary" disabled={!cur.host || cur.name === activeEnv}
                        title="设为默认(新会话的自动建立目标)"
                        onClick={() => { setActiveEnv(cur.name || sshName(cur)); markDirty() }}>
                        设为默认
                      </Button>
                      {(dirty || saveState === 'saving') && (
                        <Button size="sm" variant="secondary" className="h-7" disabled={saveState === 'saving'} onClick={() => void saveAll()}>
                          {saveState === 'saving' ? '保存中…' : '保存'}
                        </Button>
                      )}
                    </div>
                  </div>

                  {/* 右侧表单区:SSH 服务器 | 数据库(同一环境的一对一两组字段) */}
                  <div className="mb-2 flex items-center gap-1 border-b border-border">
                    {([
                      { key: 'ssh', label: 'SSH 服务器', icon: Server },
                      { key: 'db', label: '数据库', icon: Database },
                    ] as { key: EnvTab; label: string; icon: typeof Server }[]).map(({ key, label, icon: Icon }) => (
                      <button key={key} onClick={() => setEnvTab(key)}
                        className={`relative flex items-center gap-1.5 px-3 py-1.5 text-xs transition-colors ${
                          envTab === key ? 'text-foreground' : 'text-muted-foreground hover:text-foreground'
                        }`}>
                        <Icon className="h-3.5 w-3.5" />{label}
                        {envTab === key && <span className="absolute inset-x-0 bottom-0 h-px bg-primary" />}
                      </button>
                    ))}
                  </div>

                  {envTab === 'ssh' && (
                    <div className="grid grid-cols-2 gap-2">
                      <Field label="环境名称(留空自动为主机-区域)" className="col-span-2">
                        <Input className={cell} value={cur.name} onChange={(e) => patchSsh(selSsh, { name: e.target.value.trim() })} />
                      </Field>
                      <Field label="IP 主机"><Input className={cell} value={cur.host} onChange={(e) => patchSsh(selSsh, { host: e.target.value.trim() })} /></Field>
                      <Field label="端口"><Input className={cell} type="number" value={cur.port || ''} onChange={(e) => patchSsh(selSsh, { port: Number(e.target.value) || 22 })} /></Field>
                      <Field label="登录区域"><Input className={cell} value={cur.zone} onChange={(e) => patchSsh(selSsh, { zone: e.target.value.trim() })} /></Field>
                      <Field label="账号"><Input className={cell} value={cur.user} onChange={(e) => patchSsh(selSsh, { user: e.target.value.trim() })} /></Field>
                      <Field label="密码">
                        <div className="relative">
                          <Input className={`${cell} pr-8`} type={showPwd ? 'text' : 'password'} value={cur.password}
                            onChange={(e) => patchSsh(selSsh, { password: e.target.value })} />
                          <button type="button" title={showPwd ? '隐藏密码' : '显示密码'}
                            onClick={() => setShowPwd((v) => !v)}
                            className="absolute inset-y-0 right-0.5 flex w-6 items-center justify-center text-muted-foreground hover:text-foreground">
                            {showPwd ? <EyeOff className="h-3.5 w-3.5" /> : <Eye className="h-3.5 w-3.5" />}
                          </button>
                        </div>
                      </Field>
                      <Field label="区域顶级目录(留空按区域推导)" className="col-span-2"><Input className={cell} placeholder="如 /u1/topprd" value={cur.topDir} onChange={(e) => patchSsh(selSsh, { topDir: e.target.value.trim() })} /></Field>
                      <Field label="TOPENT(默认企业;调试会话 export,数字或文本)" className="col-span-2"><Input className={cell} value={cur.topent} onChange={(e) => patchSsh(selSsh, { topent: e.target.value })} /></Field>
                      <p className="col-span-2 text-muted-foreground">
                        调试会话按该服务器登录(区域/TOPENT)。该环境的数据库连接在「数据库」Tab 维护(一对一)。
                      </p>
                    </div>
                  )}

                  {envTab === 'db' && (
                    <div>
                      {!curDb ? (
                        <div className="flex flex-col items-start gap-2 border border-dashed border-border p-4 text-muted-foreground">
                          该环境未配置数据库(调试的作业解析/gzou_t 查询需要)。
                          <Button size="sm" variant="outline" className="h-7" onClick={() => patchDb(selSsh, {})}>
                            <Plus className="mr-1 h-3 w-3" />添加数据库
                          </Button>
                        </div>
                      ) : (
                        <div className="grid grid-cols-2 gap-2">
                          <Field label="类型">
                            <select className={cell} value={curDb.type || 'oracle'}
                              onChange={(e) => patchDb(selSsh, { type: e.target.value })}>
                              <option value="oracle">Oracle</option>
                              <option value="kingbase">人大金仓(PG 引擎)</option>
                            </select>
                          </Field>
                          <div className="flex items-end justify-end pb-0.5">
                            <Button size="sm" variant="ghost" className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400"
                              title="移除该环境的数据库配置(保存后生效)" onClick={() => patchSsh(selSsh, { db: null })}>
                              <Trash2 className="h-3.5 w-3.5" />移除
                            </Button>
                          </div>
                          <Field label="主机地址"><Input className={cell} placeholder="客户端与服务器均可达" value={curDb.host} onChange={(e) => patchDb(selSsh, { host: e.target.value.trim() })} /></Field>
                          <Field label="端口"><Input className={cell} type="number" placeholder={curDb.type === 'oracle' ? '1521' : '54321'} value={curDb.port || ''} onChange={(e) => patchDb(selSsh, { port: Number(e.target.value) || 0 })} /></Field>
                          {curDb.type === 'oracle' ? (
                            <Field label="服务名 (SERVICE_NAME)" className="col-span-2"><Input className={cell} placeholder="如 t35prd" value={curDb.service} onChange={(e) => patchDb(selSsh, { service: e.target.value.trim() })} /></Field>
                          ) : (
                            <Field label="库名 (database)" className="col-span-2"><Input className={cell} placeholder="如 topprd" value={curDb.database} onChange={(e) => patchDb(selSsh, { database: e.target.value.trim() })} /></Field>
                          )}

                          {/* 操作行:测试连接(直连=账号列表首项)/ 从服务器获取 */}
                          <div className="col-span-2 mt-1 flex flex-wrap items-center gap-1.5">
                            <Button size="sm" variant="outline" className="h-7 text-xs" disabled={busy === 'test'} onClick={() => void testConn()}>
                              <PlugZap className="mr-0.5 h-3 w-3" />{busy === 'test' ? '测试中…' : '测试连接(客户端直连)'}
                            </Button>
                            <Button size="sm" variant="outline" className="h-7 text-xs" disabled={busy === 'fetch' || !cur.host || !cur.user} onClick={() => void fetchFromServer()}>
                              <RefreshCw className={`mr-0.5 h-3 w-3 ${busy === 'fetch' ? 'animate-spin' : ''}`} />从服务器获取
                            </Button>
                          </div>

                          {/* 账号列表(无主账号;TOPENT 决定账号,密码查本表;客户端直连取首项) */}
                          <div className="col-span-2 mt-1">
                            <div className="mb-1 flex items-center gap-2">
                              <span className="shrink-0 text-[11px] font-medium uppercase tracking-wide text-muted-foreground">账号列表(账号=schema)</span>
                              <Separator className="flex-1" />
                            </div>
                            <div className="space-y-1">
                              <div className="flex items-center gap-1.5 text-[11px] text-muted-foreground">
                                <span className="w-36 shrink-0">账号</span>
                                <span className="flex-1">密码(缺省=账号)</span>
                                <span className="w-16 shrink-0 text-right">操作</span>
                              </div>
                              {(curDb.accounts || []).map((a, j) => (
                                <div key={j} className="flex items-center gap-1.5">
                                  <Input className={`${cell} w-36 shrink-0`} value={a.account} placeholder="如 ds"
                                    onChange={(e) => patchAcct(selSsh, j, { account: e.target.value })} />
                                  <Input className={`${cell} flex-1`} type={showPwd ? 'text' : 'password'} value={a.password}
                                    onChange={(e) => patchAcct(selSsh, j, { password: e.target.value })} />
                                  <div className="flex w-16 shrink-0 justify-end gap-1">
                                    <button type="button"
                                      className={`flex h-6 w-6 items-center justify-center ${accProbe?.i === j && accProbe.state === 'testing' ? 'animate-pulse text-sky-600 dark:text-sky-400' : 'text-muted-foreground hover:text-foreground'}`}
                                      title="服务器上以该账号+密码连显式目标库验证(只读)"
                                      onClick={() => void verifyAcct(j)}>
                                      {accProbe?.i === j && accProbe.state === 'testing' ? <Search className="h-3 w-3" /> : <CircleCheck className="h-3.5 w-3.5" />}
                                    </button>
                                    <button type="button" className="flex h-6 w-6 items-center justify-center text-muted-foreground hover:text-red-600 dark:hover:text-red-400"
                                      title="删除该账号" onClick={() => delAcct(selSsh, j)}>
                                      <Trash2 className="h-3.5 w-3.5" />
                                    </button>
                                  </div>
                                </div>
                              ))}
                              {accProbe && accProbe.i < (curDb.accounts || []).length && curDb.accounts[accProbe.i] && (
                                <div className={`text-[11px] ${accProbe.state === 'ok' ? 'text-emerald-600 dark:text-emerald-400' : accProbe.state === 'err' ? 'text-red-600 dark:text-red-400' : 'text-sky-600 dark:text-sky-400'}`}>
                                  账号 {curDb.accounts[accProbe.i].account}: {accProbe.msg}
                                </div>
                              )}
                              <div className="flex items-center gap-1.5">
                                <Input className={`${cell} w-36 shrink-0`} value={addAcct.account} placeholder="账号(如 ds)"
                                  onChange={(e) => setAddAcct({ ...addAcct, account: e.target.value })} />
                                <Input className={`${cell} flex-1`} type={showPwd ? 'text' : 'password'} value={addAcct.password} placeholder="密码(缺省=账号)"
                                  onChange={(e) => setAddAcct({ ...addAcct, password: e.target.value })}
                                  onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); pushAcct() } }} />
                                <Button size="sm" variant="outline" className="h-7 w-16 shrink-0 text-xs" onClick={pushAcct} disabled={!addAcct.account.trim()}>
                                  <Plus className="mr-0.5 h-3 w-3" />添加
                                </Button>
                              </div>
                              {!(curDb.accounts || []).length && (
                                <div className="text-[11px] text-muted-foreground">
                                  账号无主次之分:调试时按 TOPENT 经服务器 gzou_t 解析出账号,密码查本表(未收录按 账号=密码 惯例);
                                  客户端直连(ping/在线查询/测试连接)取列表首项。列表顺序可调整(直连账号放首位)。
                                </div>
                              )}
                            </div>
                          </div>
                          {note && (
                            <div className={`col-span-2 mt-1 px-3 py-2 text-xs ${note.kind === 'ok' ? 'bg-emerald-500/10 text-emerald-700 dark:text-emerald-300' : note.kind === 'err' ? 'bg-red-500/10 text-red-600 dark:text-red-400' : 'bg-sky-500/10 text-sky-700 dark:text-sky-300'}`}>
                              {note.text}
                            </div>
                          )}
                          <p className="col-span-2 mt-1 text-muted-foreground">
                            显式统一模型(主机/端口/服务名或库名 + 账号列表),客户端直连与服务器侧调试共用;服务器执行工具(sqlplus/ksql)自动探测,无需配置。
                          </p>
                        </div>
                      )}
                    </div>
                  )}
                </section>
              )}
            </div>
          )}

          {/* 外观 */}
          {section === 'appearance' && (
            <section>
              <h3 className="mb-2 font-medium text-foreground">主题</h3>
              <div className="flex gap-2">
                <Button size="sm" variant={theme === 'dark' ? 'secondary' : 'outline'} onClick={() => setTheme('dark')}>
                  <Monitor className="mr-1 h-3.5 w-3.5" />暗色
                </Button>
                <Button size="sm" variant={theme === 'light' ? 'secondary' : 'outline'} onClick={() => setTheme('light')}>
                  <Sun className="mr-1 h-3.5 w-3.5" />亮色
                </Button>
              </div>
            </section>
          )}

          {/* 高级:本机参数(不随环境走) */}
          {section === 'advanced' && (
            <section>
              <h3 className="mb-2 font-medium text-foreground">本机参数</h3>
              <div className="grid grid-cols-2 gap-2 md:grid-cols-3">
                <Field label="启动参数模板({prog} 替换)" className="col-span-2 md:col-span-3">
                  <Input className={input} value={cfg.launchArgs || ''} onChange={(e) => { setCfg({ ...cfg, launchArgs: e.target.value }); markDirty() }} />
                </Field>
                <Field label="监听地址"><Input className={input} value={cfg.listen || ''} onChange={(e) => { setCfg({ ...cfg, listen: e.target.value }); markDirty() }} /></Field>
                <Field label="停站看门狗默认(秒)"><Input className={input} value={cfg.watchdogSeconds || 0} onChange={(e) => { setCfg({ ...cfg, watchdogSeconds: Number(e.target.value) || 0 }); markDirty() }} /></Field>
                <Field label="print 数组元素上限"><Input className={input} value={cfg.printElements || 0} onChange={(e) => { setCfg({ ...cfg, printElements: Number(e.target.value) || 0 }); markDirty() }} /></Field>
                <Field label="终端宽"><Input className={input} value={cfg.termWidth || 200} onChange={(e) => { setCfg({ ...cfg, termWidth: Number(e.target.value) || 200 }); markDirty() }} /></Field>
                <Field label="终端高"><Input className={input} value={cfg.termHeight || 50} onChange={(e) => { setCfg({ ...cfg, termHeight: Number(e.target.value) || 50 }); markDirty() }} /></Field>
              </div>
              {(dirty || saveState === 'saving') && sshs.length === 0 && (
                <div className="mt-3">
                  <Button size="sm" variant="secondary" className="h-7" disabled={saveState === 'saving'} onClick={() => void saveAll()}>
                    {saveState === 'saving' ? '保存中…' : '保存'}
                  </Button>
                </div>
              )}
            </section>
          )}
        </div>
      </div>
    </div>
  )
}

function Field({ label, children, className = '' }: { label: string; children: ReactNode; className?: string }) {
  return (
    <label className={`block ${className}`}>
      <div className="mb-1 text-muted-foreground">{label}</div>
      {children}
    </label>
  )
}
