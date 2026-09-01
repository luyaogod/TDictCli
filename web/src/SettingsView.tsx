// 设置页:VS Code 式左侧一级分类(外观/环境/高级)。
// 「环境」是核心:SSH 连接 + 该环境专属启动参数(zone/launchArgs/库)成组维护,
// 列表单选「设为生效」→ 后端把该环境合并到顶层字段作为当前默认配置。
import { useEffect, useState, type ReactNode } from 'react'
import { CircleCheck, Circle, Plus, Save, Trash2, Monitor, Server, SlidersHorizontal, Sun, Search } from 'lucide-react'
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
  const [msg, setMsg] = useState('')
  const [err, setErr] = useState('')
  const [saving, setSaving] = useState(false)
  const [probing, setProbing] = useState(false)
  const [probeNote, setProbeNote] = useState('')

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
    }).catch((e) => setErr(e.message))
  }, [])

  const save = async (nextActive?: string) => {
    if (!cfg) return
    setSaving(true); setMsg(''); setErr('')
    try {
      const active = nextActive ?? activeEnv
      // 环境名自动 = 主机-区域(用户只填 IP/端口/区域/账号/密码/ENT)
      const envName = (e: EnvItem) => e.name || `${e.host}-${e.zone}`.replace(/-$/, '')
      const next = {
        ...cfg,
        activeEnv: active,
        envs: envs.filter((e) => e.host).map((e) => ({
          name: envName(e), host: e.host, port: e.port || 22, user: e.user, password: e.password,
          zone: e.zone,
          db: (e.dbEnt > 0 || e.dbType === 'kingbase') ? { type: e.dbType || 'oracle', ent: e.dbEnt || 0 } : undefined,
        })),
      }
      setActiveEnv(active)
      // 自动生成的环境名(host-zone)写回列表,保持显示一致
      setEnvs(envs.filter((e) => e.host).map((e) => ({ ...e, name: envName(e) })))
      await api.saveSettings(next)
      setCfg(next)
      setMsg(nextActive !== undefined ? `已切换生效环境:${nextActive}` : '已保存并热生效(监听地址改端口需重启 serve)')
    } catch (e: any) {
      setErr(e.message)
    } finally {
      setSaving(false)
    }
  }

  const patchEnv = (i: number, patch: Partial<EnvItem>) => {
    // 生效环境改名时同步 activeEnv,否则合并失联
    if (patch.name !== undefined && envs[i]?.name === activeEnv) setActiveEnv(patch.name)
    setEnvs(envs.map((x, j) => j === i ? { ...x, ...patch } : x))
  }

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
  const addEnv = () => {
    setEnvs([...envs, { name: '', host: '', port: 22, user: '', password: '', zone: '35', launchArgs: '', watchdogSeconds: 0, dbType: 'oracle', dbEnt: 0 }])
    setSelEnv(envs.length)
  }

  if (!cfg) return <div className="p-6 text-sm text-muted-foreground">{err || '加载配置中…'}</div>
  const cur = envs[selEnv]
  return (
    <div className="flex h-full min-h-0 text-xs">
      {/* VS Code 式左侧一级分类 */}
      <div className="w-36 shrink-0 space-y-0.5 overflow-auto border-r border-border p-2">
        {SECTIONS.map(({ key, label, icon: Icon }) => (
          <button key={key} onClick={() => setSection(key)}
            className={`flex w-full items-center gap-2 rounded-md px-2 py-1.5 text-left transition-colors ${
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
            <Button size="sm" disabled={saving} onClick={() => void save()}>
              <Save className="mr-1 h-3.5 w-3.5" />{saving ? '保存中…' : '保存'}
            </Button>
          </div>
          {msg && <div className="mb-3 rounded bg-emerald-500/10 px-3 py-2 text-emerald-700 dark:text-emerald-300">{msg}</div>}
          {err && <div className="mb-3 rounded bg-red-500/10 px-3 py-2 text-red-600 dark:text-red-400">{err}</div>}

          {/* 环境:列表单选生效 + 表单维护 */}
          {section === 'envs' && (
            <div className="flex gap-3">
              <div className="w-44 shrink-0">
                <div className="rounded-sm border border-border bg-card/60 p-1.5">
                  {envs.length === 0 && <div className="p-2 text-muted-foreground">(空)</div>}
                  {envs.map((e, i) => (
                    <button key={i} onClick={() => setSelEnv(i)}
                      className={`flex w-full items-center gap-1.5 rounded-md px-2 py-1.5 text-left transition-colors ${
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
                <section className="min-w-0 flex-1 rounded-sm border border-border bg-card/60 p-3">
                  <div className="mb-2 flex items-center justify-between">
                    <h3 className="font-medium text-foreground">
                      环境参数{cur.host && activeEnv === (cur.name || `${cur.host}-${cur.zone}`.replace(/-$/, '')) && <span className="ml-2 text-emerald-600 dark:text-emerald-400">(生效中)</span>}
                    </h3>
                    <div className="flex gap-1.5">
                      <Button size="sm" variant="secondary" disabled={!cur.host || (!cur.name && `${cur.host}-${cur.zone}` === activeEnv)}
                        title={activeEnv === (cur.name || `${cur.host}-${cur.zone}`) ? '已是生效环境' : '把该环境设为当前默认(保存后热生效)'}
                        onClick={() => void save(cur.name || `${cur.host}-${cur.zone}`.replace(/-$/, ''))}>
                        设为生效
                      </Button>
                      <Button size="sm" variant="ghost" className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400"
                        onClick={() => {
                          setEnvs(envs.filter((_, j) => j !== selEnv))
                          if (activeEnv === (cur.name || `${cur.host}-${cur.zone}`)) setActiveEnv('')
                          setSelEnv(Math.max(0, selEnv - 1))
                        }}>
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
                    <Field label="密码"><Input className={cell} type="password" value={cur.password} onChange={(e) => patchEnv(selEnv, { password: e.target.value })} /></Field>

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
                  {probeNote && <div className="mt-2 rounded bg-sky-500/10 px-3 py-2 text-sky-700 dark:text-sky-300">{probeNote}</div>}
                  <p className="mt-2 text-muted-foreground">
                    T100 目录与源码路径按「登录区域」在服务器上自动获取(与标准调试同源),无需配置。数据库连接要素(TNS/实例/库名)也完全自动探测,点「测试连接」可验证连通。start --ssh 按环境名引用,清空名称恢复自动「主机-区域」。
                  </p>
                </section>
              )}
            </div>
          )}

          {/* 外观 */}
          {section === 'appearance' && (
            <section className="rounded-sm border border-border bg-card/60 p-3">
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
            <section className="rounded-sm border border-border bg-card/60 p-3">
              <h3 className="mb-2 font-medium text-foreground">本机参数</h3>
              <div className="grid grid-cols-2 gap-2 md:grid-cols-3">
                <Field label="启动参数模板({prog} 替换)" className="col-span-2 md:col-span-3"><Input className={input} value={cfg.launchArgs || ''} onChange={(e) => setCfg({ ...cfg, launchArgs: e.target.value })} /></Field>
                <Field label="监听地址"><Input className={input} value={cfg.listen || ''} onChange={(e) => setCfg({ ...cfg, listen: e.target.value })} /></Field>
                <Field label="停站看门狗默认(秒)"><Input className={input} value={cfg.watchdogSeconds || 0} onChange={(e) => setCfg({ ...cfg, watchdogSeconds: Number(e.target.value) || 0 })} /></Field>
                <Field label="print 数组元素上限"><Input className={input} value={cfg.printElements || 0} onChange={(e) => setCfg({ ...cfg, printElements: Number(e.target.value) || 0 })} /></Field>
                <Field label="终端宽"><Input className={input} value={cfg.termWidth || 200} onChange={(e) => setCfg({ ...cfg, termWidth: Number(e.target.value) || 200 })} /></Field>
                <Field label="终端高"><Input className={input} value={cfg.termHeight || 50} onChange={(e) => setCfg({ ...cfg, termHeight: Number(e.target.value) || 50 })} /></Field>
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
