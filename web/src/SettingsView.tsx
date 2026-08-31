// 设置页:VS Code 式左侧一级分类(外观/环境/高级)。
// 「环境」是核心:SSH 连接 + 该环境专属启动参数(zone/launchArgs/库)成组维护,
// 列表单选「设为生效」→ 后端把该环境合并到顶层字段作为当前默认配置。
import { useEffect, useState, type ReactNode } from 'react'
import { CircleCheck, Circle, Plus, Save, Trash2, Monitor, Server, SlidersHorizontal, Sun } from 'lucide-react'
import { api } from './api'
import { Button, Input } from './ui'
import { useStore } from './store'

interface EnvItem {
  name: string
  host: string
  port: number
  user: string
  password: string
  zone: string
  topDir: string
  launchArgs: string
  watchdogSeconds: number
  dbTns: string
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

  useEffect(() => {
    api.settings().then((c) => {
      setCfg(c)
      setActiveEnv(c.activeEnv || '')
      setEnvs((c.envs || []).map((e: any) => ({
        name: e.name || '', host: e.host || '', port: e.port || 22, user: e.user || '', password: e.password || '',
        zone: e.zone || '', topDir: e.topDir || '', launchArgs: e.launchArgs || '', watchdogSeconds: e.watchdogSeconds || 0,
        dbTns: e.db?.tns || '', dbEnt: e.db?.ent || 0,
      })))
    }).catch((e) => setErr(e.message))
  }, [])

  const save = async (nextActive?: string) => {
    if (!cfg) return
    setSaving(true); setMsg(''); setErr('')
    try {
      const active = nextActive ?? activeEnv
      // 环境名 → db 覆盖重组回嵌套结构
      const next = {
        ...cfg,
        activeEnv: active,
        envs: envs.filter((e) => e.name && e.host).map((e) => ({
          name: e.name, host: e.host, port: e.port || 22, user: e.user, password: e.password,
          zone: e.zone, topDir: e.topDir, launchArgs: e.launchArgs, watchdogSeconds: e.watchdogSeconds || undefined,
          db: e.dbTns ? { tns: e.dbTns, ent: e.dbEnt || 0 } : undefined,
        })),
      }
      setActiveEnv(active)
      await api.saveSettings(next)
      setCfg(next)
      setMsg(nextActive !== undefined ? `已切换生效环境:${nextActive}` : '已保存并热生效(监听地址改端口需重启 serve)')
    } catch (e: any) {
      setErr(e.message)
    } finally {
      setSaving(false)
    }
  }

  const patchEnv = (i: number, patch: Partial<EnvItem>) =>
    setEnvs(envs.map((x, j) => j === i ? { ...x, ...patch } : x))
  const addEnv = () => {
    setEnvs([...envs, { name: '', host: '', port: 22, user: '', password: '', zone: '35', topDir: '', launchArgs: '', watchdogSeconds: 0, dbTns: '', dbEnt: 0 }])
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
                        <span className="block truncate font-medium">{e.name || '(未命名)'}</span>
                        <span className="block truncate text-muted-foreground">{e.host}{e.zone ? ` · ${e.zone}` : ''}</span>
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
                      环境参数{activeEnv === cur.name && cur.name && <span className="ml-2 text-emerald-600 dark:text-emerald-400">(生效中)</span>}
                    </h3>
                    <div className="flex gap-1.5">
                      <Button size="sm" variant="secondary" disabled={!cur.name || activeEnv === cur.name}
                        title={activeEnv === cur.name ? '已是生效环境' : '把该环境设为当前默认(保存后热生效)'}
                        onClick={() => void save(cur.name)}>
                        设为生效
                      </Button>
                      <Button size="sm" variant="ghost" className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400"
                        onClick={() => {
                          setEnvs(envs.filter((_, j) => j !== selEnv))
                          if (activeEnv === cur.name) setActiveEnv('')
                          setSelEnv(Math.max(0, selEnv - 1))
                        }}>
                        <Trash2 className="h-3.5 w-3.5" />
                      </Button>
                    </div>
                  </div>
                  <div className="grid grid-cols-2 gap-2">
                    <Field label="环境名称(start --ssh 按名引用)"><Input className={cell} placeholder="如 35测试" value={cur.name} onChange={(e) => patchEnv(selEnv, { name: e.target.value.trim() })} /></Field>
                    <Field label="区域(31开发/35测试/36正式)"><Input className={cell} value={cur.zone} onChange={(e) => patchEnv(selEnv, { zone: e.target.value.trim() })} /></Field>
                    <Field label="主机"><Input className={cell} value={cur.host} onChange={(e) => patchEnv(selEnv, { host: e.target.value.trim() })} /></Field>
                    <Field label="端口"><Input className={cell} value={cur.port} onChange={(e) => patchEnv(selEnv, { port: Number(e.target.value) || 22 })} /></Field>
                    <Field label="用户"><Input className={cell} value={cur.user} onChange={(e) => patchEnv(selEnv, { user: e.target.value.trim() })} /></Field>
                    <Field label="密码"><Input className={cell} type="password" value={cur.password} onChange={(e) => patchEnv(selEnv, { password: e.target.value })} /></Field>
                    <Field label="启动参数模板({prog} 替换)" className="col-span-2"><Input className={cell} value={cur.launchArgs} onChange={(e) => patchEnv(selEnv, { launchArgs: e.target.value })} /></Field>
                    <Field label="顶级目录(留空按区域推导)"><Input className={cell} placeholder="/u1/t35tst" value={cur.topDir} onChange={(e) => patchEnv(selEnv, { topDir: e.target.value.trim() })} /></Field>
                    <Field label="停站看门狗(秒,0=默认)"><Input className={cell} value={cur.watchdogSeconds || ''} onChange={(e) => patchEnv(selEnv, { watchdogSeconds: Number(e.target.value) || 0 })} /></Field>
                    <Field label="TNS 别名(留空按区域推导)"><Input className={cell} placeholder="t35tst" value={cur.dbTns} onChange={(e) => patchEnv(selEnv, { dbTns: e.target.value.trim() })} /></Field>
                    <Field label="企业(TOPENT)"><Input className={cell} value={cur.dbEnt || ''} onChange={(e) => patchEnv(selEnv, { dbEnt: Number(e.target.value) || 0 })} /></Field>
                  </div>
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
