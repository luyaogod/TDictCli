// 设置页:多 SSH / 多数据库连接管理 + 全局调试参数(config.json debug 节)
import { useEffect, useState, type ReactNode } from 'react'
import { Plus, Save, Trash2 } from 'lucide-react'
import { api } from './api'
import { Button, Input } from './ui'
import { useStore } from './store'
import { Moon, Sun } from 'lucide-react'

interface SSHItem { name: string; host: string; port: number; user: string; password: string }
interface DBItem { name: string; tns: string; ent: number; user: string; password: string }

const input = 'h-7 flex-1 text-xs'
const cell = 'h-7 w-full min-w-0 text-xs'

export function SettingsView() {
  const theme = useStore((s) => s.theme)
  const setTheme = useStore((s) => s.setTheme)
  const [cfg, setCfg] = useState<any>(null)
  const [sshs, setSshs] = useState<SSHItem[]>([])
  const [dbs, setDbs] = useState<DBItem[]>([])
  const [msg, setMsg] = useState('')
  const [err, setErr] = useState('')
  const [saving, setSaving] = useState(false)

  useEffect(() => {
    api.settings().then((c) => {
      setCfg(c)
      setSshs((c.sshs || []).map((s: any) => ({ name: s.name || '', host: s.host || '', port: s.port || 22, user: s.user || '', password: s.password || '' })))
      setDbs((c.dbs || []).map((d: any) => ({ name: d.name || '', tns: d.tns || '', ent: d.ent || 0, user: d.user || '', password: d.password || '' })))
    }).catch((e) => setErr(e.message))
  }, [])

  const save = async () => {
    if (!cfg) return
    setSaving(true); setMsg(''); setErr('')
    try {
      const next = {
        ...cfg,
        sshs: sshs.filter((s) => s.name && s.host),
        dbs: dbs.filter((d) => d.name && d.tns),
      }
      await api.saveSettings(next)
      setCfg(next)
      setMsg('已保存并热生效(listen 改端口需重启 serve)')
    } catch (e: any) {
      setErr(e.message)
    } finally {
      setSaving(false)
    }
  }

  if (!cfg) return <div className="p-6 text-sm text-muted-foreground">{err || '加载配置中…'}</div>
  return (
    <div className="h-full min-h-0 overflow-auto p-4 text-xs">
      <div className="mx-auto max-w-3xl space-y-5">
        <div className="flex items-center justify-between">
          <h2 className="text-sm font-medium text-foreground">设置</h2>
          <Button size="sm" disabled={saving} onClick={() => void save()}>
            <Save className="mr-1 h-3.5 w-3.5" />{saving ? '保存中…' : '保存'}
          </Button>
        </div>
        {msg && <div className="rounded bg-emerald-500/10 px-3 py-2 text-emerald-700 dark:text-emerald-300">{msg}</div>}
        {err && <div className="rounded bg-red-500/10 px-3 py-2 text-red-600 dark:text-red-400">{err}</div>}

        {/* 外观 */}
        <section className="rounded-sm border border-border bg-card/60 p-3">
          <h3 className="mb-2 font-medium text-foreground">外观</h3>
          <div className="flex gap-2">
            <Button size="sm" variant={theme === 'dark' ? 'secondary' : 'outline'} onClick={() => setTheme('dark')}>
              <Moon className="mr-1 h-3.5 w-3.5" />暗色
            </Button>
            <Button size="sm" variant={theme === 'light' ? 'secondary' : 'outline'} onClick={() => setTheme('light')}>
              <Sun className="mr-1 h-3.5 w-3.5" />亮色
            </Button>
          </div>
        </section>

        {/* 默认连接 */}
        <section className="rounded-sm border border-border bg-card/60 p-3">
          <h3 className="mb-2 font-medium text-foreground">默认 SSH 连接</h3>
          <div className="grid grid-cols-2 gap-2 md:grid-cols-4">
            <Field label="主机"><Input className={input} value={cfg.ssh?.host || ''} onChange={(e) => setCfg({ ...cfg, ssh: { ...cfg.ssh, host: e.target.value } })} /></Field>
            <Field label="端口"><Input className={input} value={cfg.ssh?.port || 22} onChange={(e) => setCfg({ ...cfg, ssh: { ...cfg.ssh, port: Number(e.target.value) || 22 } })} /></Field>
            <Field label="用户"><Input className={input} value={cfg.ssh?.user || ''} onChange={(e) => setCfg({ ...cfg, ssh: { ...cfg.ssh, user: e.target.value } })} /></Field>
            <Field label="密码"><Input className={input} type="password" value={cfg.ssh?.password || ''} onChange={(e) => setCfg({ ...cfg, ssh: { ...cfg.ssh, password: e.target.value } })} /></Field>
          </div>
        </section>

        {/* 多 SSH 列表 */}
        <section className="rounded-sm border border-border bg-card/60 p-3">
          <div className="mb-2 flex items-center justify-between">
            <h3 className="font-medium text-foreground">SSH 连接列表(start --ssh 按名引用)</h3>
            <Button size="sm" variant="outline" onClick={() => setSshs([...sshs, { name: '', host: '', port: 22, user: '', password: '' }])}>
              <Plus className="mr-1 h-3 w-3" />新增
            </Button>
          </div>
          {sshs.length === 0 && <div className="text-muted-foreground">(空)</div>}
          {sshs.map((s, i) => (
            <div key={i} className="mb-1.5 grid grid-cols-[110px_1fr_70px_110px_110px_28px] items-center gap-1.5">
              <Input className={cell} placeholder="名称" value={s.name} onChange={(e) => setSshs(sshs.map((x, j) => j === i ? { ...x, name: e.target.value } : x))} />
              <Input className={cell} placeholder="主机" value={s.host} onChange={(e) => setSshs(sshs.map((x, j) => j === i ? { ...x, host: e.target.value } : x))} />
              <Input className={cell} placeholder="端口" value={s.port} onChange={(e) => setSshs(sshs.map((x, j) => j === i ? { ...x, port: Number(e.target.value) || 22 } : x))} />
              <Input className={cell} placeholder="用户" value={s.user} onChange={(e) => setSshs(sshs.map((x, j) => j === i ? { ...x, user: e.target.value } : x))} />
              <Input className={cell} placeholder="密码" type="password" value={s.password} onChange={(e) => setSshs(sshs.map((x, j) => j === i ? { ...x, password: e.target.value } : x))} />
              <button className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400" onClick={() => setSshs(sshs.filter((_, j) => j !== i))}><Trash2 className="h-3.5 w-3.5" /></button>
            </div>
          ))}
        </section>

        {/* 默认数据库 */}
        <section className="rounded-sm border border-border bg-card/60 p-3">
          <h3 className="mb-2 font-medium text-foreground">默认数据库</h3>
          <div className="grid grid-cols-2 gap-2 md:grid-cols-4">
            <Field label="TNS 别名"><Input className={input} value={cfg.db?.tns || ''} onChange={(e) => setCfg({ ...cfg, db: { ...(cfg.db || {}), tns: e.target.value } })} /></Field>
            <Field label="企业(TOPENT)"><Input className={input} value={cfg.db?.ent || ''} onChange={(e) => setCfg({ ...cfg, db: { ...(cfg.db || {}), ent: Number(e.target.value) || 0 } })} /></Field>
          </div>
        </section>

        {/* 多数据库列表 */}
        <section className="rounded-sm border border-border bg-card/60 p-3">
          <div className="mb-2 flex items-center justify-between">
            <h3 className="font-medium text-foreground">数据库连接列表</h3>
            <Button size="sm" variant="outline" onClick={() => setDbs([...dbs, { name: '', tns: '', ent: 0, user: '', password: '' }])}>
              <Plus className="mr-1 h-3 w-3" />新增
            </Button>
          </div>
          {dbs.length === 0 && <div className="text-muted-foreground">(空)</div>}
          {dbs.map((d, i) => (
            <div key={i} className="mb-1.5 grid grid-cols-[110px_1fr_90px_110px_110px_28px] items-center gap-1.5">
              <Input className={cell} placeholder="名称" value={d.name} onChange={(e) => setDbs(dbs.map((x, j) => j === i ? { ...x, name: e.target.value } : x))} />
              <Input className={cell} placeholder="TNS 别名" value={d.tns} onChange={(e) => setDbs(dbs.map((x, j) => j === i ? { ...x, tns: e.target.value } : x))} />
              <Input className={cell} placeholder="企业" value={d.ent || ''} onChange={(e) => setDbs(dbs.map((x, j) => j === i ? { ...x, ent: Number(e.target.value) || 0 } : x))} />
              <Input className={cell} placeholder="用户" value={d.user} onChange={(e) => setDbs(dbs.map((x, j) => j === i ? { ...x, user: e.target.value } : x))} />
              <Input className={cell} placeholder="密码" type="password" value={d.password} onChange={(e) => setDbs(dbs.map((x, j) => j === i ? { ...x, password: e.target.value } : x))} />
              <button className="text-muted-foreground hover:text-red-600 dark:hover:text-red-400" onClick={() => setDbs(dbs.filter((_, j) => j !== i))}><Trash2 className="h-3.5 w-3.5" /></button>
            </div>
          ))}
        </section>

        {/* 全局参数 */}
        <section className="rounded-sm border border-border bg-card/60 p-3">
          <h3 className="mb-2 font-medium text-foreground">全局参数</h3>
          <div className="grid grid-cols-2 gap-2 md:grid-cols-3">
            <Field label="区域(31开发/35测试/36正式)">
              <Input className={input} value={cfg.zone || ''} onChange={(e) => setCfg({ ...cfg, zone: e.target.value.trim() })} />
            </Field>
            <Field label="监听地址">
              <Input className={input} value={cfg.listen || ''} onChange={(e) => setCfg({ ...cfg, listen: e.target.value })} />
            </Field>
            <Field label="停站看门狗(秒)">
              <Input className={input} value={cfg.watchdogSeconds || 0} onChange={(e) => setCfg({ ...cfg, watchdogSeconds: Number(e.target.value) || 0 })} />
            </Field>
            <Field label="作业启动参数模板({prog} 替换)">
              <Input className={input} value={cfg.launchArgs || ''} onChange={(e) => setCfg({ ...cfg, launchArgs: e.target.value })} />
            </Field>
          </div>
        </section>
      </div>
    </div>
  )
}

function Field({ label, children }: { label: string; children: ReactNode }) {
  return (
    <label className="block">
      <div className="mb-1 text-muted-foreground">{label}</div>
      {children}
    </label>
  )
}
