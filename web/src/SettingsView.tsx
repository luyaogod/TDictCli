// 设置页:VS Code 式左侧一级分类(外观/环境/高级)。
// 「环境」是核心:SSH 连接 + 该环境专属启动参数(zone/launchArgs/库)成组维护,
// 列表单选「设为生效」→ 后端把该环境合并到顶层字段作为当前默认配置。
import { useEffect, useRef, useState, type ReactNode } from 'react'
import { CircleCheck, Circle, Plus, Trash2, Monitor, Server, SlidersHorizontal, Sun, Search, Eye, EyeOff } from 'lucide-react'
import { api } from './api'
import { Button, Input, Separator } from './ui'
import { useStore } from './store'

interface EnvItem {
  name: string
  host: string
  port: number
  user: string
  password: string
  zone: string
  launchArgs: string
  watchdogSeconds: number
  dbType: string // oracle | kingbase
  dbEnt: number
}

type Section = 'appearance' | 'envs' | 'advanced'

const SECTIONS: { key: Section; label: string; icon: typeof Monitor }[] = [
  { key: 'envs', label: '环境', icon: Server },
  { key: 'appearance', label: '外观', icon: Monitor },
  { key: 'advanced', label: '高级', icon: SlidersHorizontal },
]

const input = 'h-7 text-xs'
const cell = 'h-7 w-full min-w-0 text-xs'

export function SettingsView() {
  const theme = useStore((s) => s.theme)
  const setTheme = useStore((s) => s.setTheme)
  const [section, setSection] = useState<Section>('envs')
  const [cfg, setCfg] = useState<any>(null)
  const [envs, setEnvs] = useState<EnvItem[]>([])
  const [activeEnv, setActiveEnv] = useState('')
  const [selEnv, setSelEnv] = useState(0) // 列表浏览选中项(≠ 生效项)
  const [err, setErr] = useState('')
  const [saveState, setSaveState] = useState<'idle' | 'saving' | 'saved'>('idle')
  const [showPwd, setShowPwd] = useState(false) // 密码明文/密文切换
  const [probing, setProbing] = useState(false)
  const [probeNote, setProbeNote] = useState('')
  // 自动保存:待写快照 ref + 串行落盘 + 短防抖(避免连续键入时并发 PUT 乱序)
  const latestRef = useRef<{ cfg: any; envs: EnvItem[]; active: string } | null>(null)
  const busyRef = useRef(false)
  const timerRef = useRef<number | undefined>(undefined)

  useEffect(() => {
    api.settings().then((c) => {
      setCfg(c)
      setActiveEnv(c.activeEnv || '')
      setEnvs((c.envs || []).map((e: any) => ({
        name: e.name || '', host: e.host || '', port: e.port || 22, user: e.user || '', password: e.password || '',
        zone: e.zone || '', launchArgs: e.launchArgs || '', watchdogSeconds: e.watchdogSeconds || 0,
        dbType: e.db?.type || 'oracle', dbEnt: e.db?.ent || 0,
      })))
      // activeEnv 指向的环境不存在(历史脏数据/已删除)时视为未设置
      const names = (c.envs || []).map((e: any) => e.name)
      if (c.activeEnv && !names.includes(c.activeEnv)) setActiveEnv('')
      // 默认选中生效环境的明细(进入设置页即展示当前生效配置)
      const idx = names.indexOf(c.activeEnv || '')
      if (idx >= 0) setSelEnv(idx)
    }).catch((e) => setErr(e.message))
  }, [])

  // ---- 即时保存:字段修改后短防抖自动整包写回并热生效,无需「保存」按钮 ----
  const envName = (e: EnvItem) => e.name || `${e.host}-${e.zone}`.replace(/-$/, '')
  const persist = async (c: any, e: EnvItem[], a: string) => {
    const next = {
      ...c,
      activeEnv: a,
      envs: e.filter((x) => x.host).map((x) => ({
        name: envName(x), host: x.host, port: x.port || 22, user: x.user, password: x.password,
        zone: x.zone,
        db: (x.dbEnt > 0 || x.dbType === 'kingbase') ? { type: x.dbType || 'oracle', ent: x.dbEnt || 0 } : undefined,
      })),
    }
    await api.saveSettings(next)
    setCfg(next)
    // 自动名(host-zone)回写列表,让生效对勾与显示名称和存储一致;手填名保留
    setEnvs((prev) => prev.map((x) => ({ ...x, name: envName(x) })))
    setActiveEnv(a)
  }
  const drain = async () => {
    if (busyRef.current) return
    const snap = latestRef.current
    if (!snap) return
    latestRef.current = null
    busyRef.current = true
    try {
      await persist(snap.cfg, snap.envs, snap.active)
      setSaveState('saved')
    } catch (ex: any) {
      setErr(ex.message)
      setSaveState('idle')
    } finally {
      busyRef.current = false
      if (latestRef.current) void drain() // 落盘期间又有修改,继续写最新快照
    }
  }
  const scheduleSave = (c: any, e: EnvItem[], a: string, immediate = false) => {
    latestRef.current = { cfg: c, envs: e, active: a }
    setSaveState('saving')
    window.clearTimeout(timerRef.current)
    if (immediate) { void drain(); return }
    timerRef.current = window.setTimeout(() => void drain(), 600)
  }
  // 修改统一入口:先更新本地状态,再登记一次自动保存(用修改后的快照,避免闭包旧值)
  const change = (patch: { cfg?: any; envs?: EnvItem[]; active?: string }, immediate = false) => {
    const nc = patch.cfg ?? cfg
    const ne = patch.envs ?? envs
    const na = patch.active ?? activeEnv
    if (patch.cfg !== undefined) setCfg(nc)
    if (patch.envs !== undefined) setEnvs(ne)
    if (patch.active !== undefined) setActiveEnv(na)
    scheduleSave(nc, ne, na, immediate)
  }
  const patchEnv = (i: number, patch: Partial<EnvItem>) => {
    const next = envs.map((x, j) => (j === i ? { ...x, ...patch } : x))
    let na: string | undefined
    // 生效环境改名时同步 activeEnv,否则合并失联
    if (patch.name !== undefined && envs[i]?.name === activeEnv) na = patch.name
    change({ envs: next, active: na })
  }
  const patchCfg = (patch: any) => change({ cfg: { ...cfg, ...patch } })
  const addEnv = () => {
    setSelEnv(envs.length)
    change({ envs: [...envs, { name: '', host: '', port: 22, user: '', password: '', zone: '35', launchArgs: '', watchdogSeconds: 0, dbType: 'oracle', dbEnt: 0 }] })
  }
  const delEnv = () => {
    const cur2 = envs[selEnv]
    const activeKey = cur2 ? cur2.name || `${cur2.host}-${cur2.zone}`.replace(/-$/, '') : ''
    setSelEnv(Math.max(0, selEnv - 1))
    change({ envs: envs.filter((_, j) => j !== selEnv), active: activeKey === activeEnv ? '' : undefined }, true)
  }
  // 卸载时清掉未落盘的防抖计时
  useEffect(() => () => window.clearTimeout(timerRef.current), [])
  // 「已自动保存」提示短暂停留后恢复默认文案
  useEffect(() => {
    if (saveState !== 'saved') return
    const t = window.setTimeout(() => setSaveState('idle'), 2000)
    return () => window.clearTimeout(t)
  }, [saveState])

  // 测试数据库连接:SSH 上服务器自动探测连接要素并验证连通(只读命令),结果仅展示
  const autoProbe = async () => {
    const e = envs[selEnv]
    if (!e?.host || !e?.user) { setProbeNote('请先填写 IP 主机与账号'); return }
    setProbing(true); setProbeNote('')
    try {
      const r = await api.probeDB({ host: e.host, port: e.port || 22, user: e.user, password: e.password, zone: e.zone, type: e.dbType || 'oracle' })
      if (r.type === 'kingbase') {
        setProbeNote(r.note || `连接正常:金仓 ${r.database || '?'} @ 127.0.0.1:${r.port || '?'}`)
      } else {
        const extra = r.host ? `(${r.host}:${r.port || '?'} / ${r.service || '?'})` : ''
        setProbeNote(r.note || `连接正常:TNS ${r.tns || '?'} ${extra} ORACLE_HOME ${r.oracleHome || '?'}`)
      }
    } catch (err: any) {
      setProbeNote('连接失败:' + err.message)
    } finally {
      setProbing(false)
    }
  }
  if (!cfg) return <div className="p-6 text-sm text-muted-foreground">{err || '加载配置中…'}</div>
  const cur = envs[selEnv]
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
          <div className="mb-3 flex items-center justify-between">
            <h2 className="text-sm font-medium text-foreground">{SECTIONS.find((s) => s.key === section)?.label}设置</h2>
            {/* 即时保存状态:修改自动落盘并热生效,无需保存按钮 */}
            <span className={`text-[11px] ${saveState === 'saved' ? 'text-emerald-600 dark:text-emerald-400' : saveState === 'saving' ? 'text-muted-foreground' : ''}`}>
              {saveState === 'saving' ? '保存中…' : saveState === 'saved' ? '已自动保存 ✓' : '修改即时生效'}
            </span>
          </div>
          {err && <div className="mb-3 bg-red-500/10 px-3 py-2 text-red-600 dark:text-red-400">{err}</div>}

          {/* 环境:列表单选生效 + 表单维护(扁平布局,列表与表单以分割线相连) */}
          {section === 'envs' && (
            <div className="flex">
              <div className="w-44 shrink-0 border-r border-border pr-1.5">
                <div className="py-1">
                  {envs.length === 0 && <div className="p-2 text-muted-foreground">(空)</div>}
                  {envs.map((e, i) => (
                    <button key={i} onClick={() => setSelEnv(i)}
                      className={`flex w-full items-center gap-1.5 px-2 py-1.5 text-left transition-colors ${
                        selEnv === i ? 'bg-accent text-accent-foreground' : 'hover:bg-accent/60'
                      }`}>
                      {activeEnv === e.name
                        ? <CircleCheck className="h-3.5 w-3.5 shrink-0 text-emerald-600 dark:text-emerald-400" />
                        : <Circle className="h-3.5 w-3.5 shrink-0 text-muted-foreground/40" />}
                      <span className="min-w-0 flex-1">
                        <span className="block truncate font-medium">{e.name || e.host || '(新环境)'}</span>
                        <span className="block truncate text-muted-foreground">{e.zone ? `${e.zone} · ` : ''}{e.port || 22}</span>
                      </span>
                    </button>
                  ))}
                </div>
                <Button size="sm" variant="outline" className="mt-2 w-full" onClick={addEnv}>
                  <Plus className="mr-1 h-3 w-3" />新增环境
                </Button>
              </div>

              {cur && (
                <section className="min-w-0 flex-1 pl-3">
                  <div className="mb-2 flex items-center justify-between">
                    <h3 className="font-medium text-foreground">
                      环境参数{cur.host && activeEnv === (cur.name || `${cur.host}-${cur.zone}`.replace(/-$/, '')) && <span className="ml-2 text-emerald-600 dark:text-emerald-400">(生效中)</span>}
                    </h3>
                    <div className="flex gap-1.5">
                      <Button size="sm" variant="secondary" disabled={!cur.host || (!cur.name && `${cur.host}-${cur.zone}` === activeEnv)}
                        title={activeEnv === (cur.name || `${cur.host}-${cur.zone}`) ? '已是生效环境' : '设为当前默认并立即热生效'}
                        onClick={() => change({ active: cur.name || `${cur.host}-${cur.zone}`.replace(/-$/, '') }, true)}>
                        设为生效
                      </Button>
                      <Button size="sm" variant="ghost" className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400"
                        onClick={delEnv}>
                        <Trash2 className="h-3.5 w-3.5" />
                      </Button>
                    </div>
                  </div>
                  <div className="grid grid-cols-2 gap-2">
                    {/* SSH 连接配置 */}
                    <Field label="环境名称(留空自动为主机-区域)" className="col-span-2"><Input className={cell} placeholder={`${cur.host || 'IP'}-${cur.zone || '区域'}`} value={cur.name} onChange={(e) => patchEnv(selEnv, { name: e.target.value.trim() })} /></Field>
                    <Field label="IP 主机"><Input className={cell} value={cur.host} onChange={(e) => patchEnv(selEnv, { host: e.target.value.trim() })} /></Field>
                    <Field label="端口"><Input className={cell} value={cur.port} onChange={(e) => patchEnv(selEnv, { port: Number(e.target.value) || 22 })} /></Field>
                    <Field label="登录区域"><Input className={cell} value={cur.zone} onChange={(e) => patchEnv(selEnv, { zone: e.target.value.trim() })} /></Field>
                    <Field label="账号"><Input className={cell} value={cur.user} onChange={(e) => patchEnv(selEnv, { user: e.target.value.trim() })} /></Field>
                    <Field label="密码">
                      <div className="relative">
                        <Input className={`${cell} pr-8`} type={showPwd ? 'text' : 'password'} value={cur.password}
                          onChange={(e) => patchEnv(selEnv, { password: e.target.value })} />
                        <button type="button" title={showPwd ? '隐藏密码' : '显示密码'}
                          onClick={() => setShowPwd((v) => !v)}
                          className="absolute inset-y-0 right-0.5 flex w-6 items-center justify-center text-muted-foreground hover:text-foreground">
                          {showPwd ? <EyeOff className="h-3.5 w-3.5" /> : <Eye className="h-3.5 w-3.5" />}
                        </button>
                      </div>
                    </Field>

                    {/* 数据库配置:连接要素(TNS/实例)完全自动探测,无需手填 */}
                    <GroupLabel title="数据库配置" />
                    <Field label="数据库类型" className="col-span-2">
                      <div className="flex items-center gap-1.5">
                        <select className={cell} value={cur.dbType || 'oracle'}
                          onChange={(e) => patchEnv(selEnv, { dbType: e.target.value })}>
                          <option value="oracle">Oracle</option>
                          <option value="kingbase">人大金仓(PG 引擎)</option>
                        </select>
                        <Button size="sm" variant="outline" className="h-7 shrink-0 text-xs" disabled={probing || !cur.host || !cur.user}
                          title="SSH 上服务器自动探测连接要素并验证连通(只读命令)"
                          onClick={() => void autoProbe()}>
                          <Search className="mr-0.5 h-3 w-3" />{probing ? '测试中…' : '测试连接'}
                        </Button>
                      </div>
                    </Field>

                    {/* 环境变量 */}
                    <GroupLabel title="环境变量" />
                    <Field label="企业 TOPENT(留空用选区默认)" className="col-span-2"><Input className={cell} value={cur.dbEnt || ''} onChange={(e) => patchEnv(selEnv, { dbEnt: Number(e.target.value) || 0 })} /></Field>
                  </div>
                  {probeNote && <div className="mt-2 bg-sky-500/10 px-3 py-2 text-sky-700 dark:text-sky-300">{probeNote}</div>}
                  <p className="mt-2 text-muted-foreground">
                    T100 目录与源码路径按「登录区域」在服务器上自动获取(与标准调试同源),无需配置。数据库连接要素(TNS/实例/库名)也完全自动探测,点「测试连接」可验证连通。start --ssh 按环境名引用,清空名称恢复自动「主机-区域」。
                  </p>
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
                <Field label="启动参数模板({prog} 替换)" className="col-span-2 md:col-span-3"><Input className={input} value={cfg.launchArgs || ''} onChange={(e) => patchCfg({ launchArgs: e.target.value })} /></Field>
                <Field label="监听地址"><Input className={input} value={cfg.listen || ''} onChange={(e) => patchCfg({ listen: e.target.value })} /></Field>
                <Field label="停站看门狗默认(秒)"><Input className={input} value={cfg.watchdogSeconds || 0} onChange={(e) => patchCfg({ watchdogSeconds: Number(e.target.value) || 0 })} /></Field>
                <Field label="print 数组元素上限"><Input className={input} value={cfg.printElements || 0} onChange={(e) => patchCfg({ printElements: Number(e.target.value) || 0 })} /></Field>
                <Field label="终端宽"><Input className={input} value={cfg.termWidth || 200} onChange={(e) => patchCfg({ termWidth: Number(e.target.value) || 200 })} /></Field>
                <Field label="终端高"><Input className={input} value={cfg.termHeight || 50} onChange={(e) => patchCfg({ termHeight: Number(e.target.value) || 50 })} /></Field>
              </div>
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

// 分组标题 + 分隔线(VS Code 风格):占满一行,标题左侧、分隔线右侧
function GroupLabel({ title }: { title: string }) {
  return (
    <div className="col-span-2 mt-1 flex items-center gap-2">
      <span className="shrink-0 text-[11px] font-medium uppercase tracking-wide text-muted-foreground">{title}</span>
      <Separator className="flex-1" />
    </div>
  )
}
