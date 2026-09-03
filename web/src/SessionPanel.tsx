// 「会话」sheet:管理全局单一常驻会话(每个环境一条,通常就是默认环境那一条)。
// 会话 = 与某 T100 环境的 SSH+登录态,debug 只是在其上跑一轮;结束调试/程序退出都只回空闲。
// 本面板负责 切换(换到另一环境并重连)/ 重启(断开重连同环境)/ 结束(彻底释放连接);
// 这些操作都会先结束当前 debug(由后端收口)。行只显示环境名 + 状态,保持精简。
// 两个可折叠区域(仿 debug 运行区手风琴):「会话」(环境行列表)、「环境变量」(TOPENT 设置)。
import { useCallback, useEffect, useState, type ReactNode } from 'react'
import { ArrowRightLeft, Check, RefreshCw, RotateCcw, Power } from 'lucide-react'
import { useStore } from './store'
import { api } from './api'
import { Input } from './ui'

const STATE_LABEL: Record<string, string> = {
  idle: '空闲', loading: '连接中…', stopped: '已停站', running: '运行中', exit: '已断开',
}
const STATE_DOT: Record<string, string> = {
  idle: 'bg-emerald-500',
  loading: 'bg-amber-400',
  stopped: 'bg-[#cc6633]',
  running: 'bg-[#0078d4]',
  exit: 'bg-muted-foreground/50',
}

// 折叠区头(仿 RightPanels 手风琴):标题 + 旋转箭头;extra 放右侧附加控件
function SectionHead({ title, open, onToggle, extra }: {
  title: string; open: boolean; onToggle: () => void; extra?: ReactNode
}) {
  return (
    <div className="flex h-8 shrink-0 items-center border-b border-border px-2.5 text-xs font-medium text-muted-foreground">
      <button onClick={onToggle}
        title={open ? `收起${title}` : `展开${title}`}
        className="flex min-w-0 flex-1 items-center justify-between gap-1 hover:text-foreground">
        <span className="truncate">{title}</span>
        <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none"
          stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"
          className={`shrink-0 text-muted-foreground transition-transform duration-200 ${open ? 'rotate-180' : ''}`}>
          <path d="m6 9 6 6 6-6" />
        </svg>
      </button>
      {extra}
    </div>
  )
}

export function SessionPanel() {
  const sessionId = useStore((s) => s.sessionId)
  const sessionEnv = useStore((s) => s.sessionEnv)
  const state = useStore((s) => s.state)
  const sessionOp = useStore((s) => s.sessionOp)

  const [envNames, setEnvNames] = useState<string[]>([])
  const [defaultEnv, setDefaultEnv] = useState('')
  const [busy, setBusy] = useState<string | null>(null)
  const [err, setErr] = useState('')

  // 两个可折叠区域(默认展开,仿 debug 运行区)
  const [open, setOpen] = useState({ session: true, env: true })
  const toggle = (k: 'session' | 'env') => setOpen((o) => ({ ...o, [k]: !o[k] }))

  // 环境清单来自设置(名称即会话目标);默认环境只影响"尚无会话时的自动建立"
  const load = useCallback(() => {
    void api.settings().then((c: any) => {
      const names = (c.envs || [])
        .map((e: any) => e.name || `${e.host || ''}-${e.zone || ''}`.replace(/-$/, ''))
        .filter(Boolean)
      setEnvNames(names)
      setDefaultEnv(c.activeEnv || '')
    }).catch(() => {})
  }, [])
  useEffect(() => { load() }, [load])
  // 连接中 loading→idle 等状态跟随;每 3s 轻量刷新一次
  useEffect(() => {
    const t = window.setInterval(load, 3000)
    return () => window.clearInterval(t)
  }, [load])

  const curEnv = sessionEnv || (sessionId ? defaultEnv : '')
  // 行 = 已配置环境;当前会话不在列表(顶层直连等)时补一行
  const connectedEnv = sessionId && curEnv ? [curEnv] : []
  const rows = Array.from(new Set([...envNames, ...(curEnv ? [curEnv] : [])]))

  const doOp = async (op: 'close' | 'restart' | 'switch', env?: string) => {
    const hasRun = !!sessionId && state !== 'idle' && state !== '' && state !== 'exit'
    const envLabel = env || curEnv || ''
    if (op === 'close') {
      if (!window.confirm('结束会话将断开该环境的 SSH 连接并结束当前调试,确认?')) return
    } else if (op === 'restart') {
      if (!window.confirm(`重启会话将断开并重新连接 ${envLabel},会先结束当前调试,确认?`)) return
    } else if (op === 'switch') {
      if (hasRun) {
        if (!window.confirm(`切换会话到 ${envLabel} 会先结束当前调试并断开当前连接,确认?`)) return
      }
    }
    setBusy(`${op}:${env || ''}`)
    setErr('')
    try {
      await sessionOp(op, env)
    } catch (e: any) {
      setErr(e.message || String(e))
    } finally {
      setBusy(null)
    }
  }

  // ---- 环境变量:TOPENT(会话内,仅 idle 可设置,下一轮调试生效) ----
  const [topent, setTopent] = useState('')
  const [topentSaved, setTopentSaved] = useState(false)
  const [topentErr, setTopentErr] = useState('')
  const canSetTopent = !!sessionId && state === 'idle' && !busy
  const loadTopent = useCallback(() => {
    if (!sessionId || state !== 'idle') return
    void api.snapshot(sessionId).then((snap: any) => {
      setTopent(snap.topent || '')
    }).catch(() => {})
  }, [sessionId, state])
  useEffect(() => { loadTopent() }, [loadTopent])
  const saveTopent = async () => {
    if (!canSetTopent || !sessionId) return
    const v = topent.trim()
    if (v && !/^\d{1,3}$/.test(v)) {
      setTopentErr('TOPENT 须为 1~3 位数字(留空 = 清除手动设置)')
      return
    }
    setTopentErr('')
    try {
      await api.topent(sessionId, v)
      setTopentSaved(true)
      useStore.getState().pushTimeline({ origin: 'human', kind: 'command', text: `设置 TOPENT${v ? '=' + v : '(清除)'}` })
      window.setTimeout(() => setTopentSaved(false), 1500)
    } catch (e: any) {
      setTopentErr(e.message || String(e))
    }
  }

  return (
    <div className="flex h-full min-h-0 w-full flex-col overflow-hidden bg-background">
      <div className="min-h-0 flex-1 overflow-auto">
        {/* 区域一:会话(环境行列表) */}
        <SectionHead title="会话" open={open.session} onToggle={() => toggle('session')}
          extra={
            <button title="刷新环境/会话状态" onClick={load}
              className="ml-1 p-0.5 text-muted-foreground transition-colors hover:text-foreground">
              <RefreshCw className="h-3.5 w-3.5" />
            </button>
          } />
        {open.session && (
          <div className="py-1">
            {rows.length === 0 && (
              <div className="space-y-1 p-2 text-xs text-muted-foreground">
                <div>尚未配置环境(设置 → 环境)。</div>
                <div>直接「启动调试」将使用默认连接并自动建立会话。</div>
              </div>
            )}
            {rows.map((name) => {
              const connected = connectedEnv.length > 0 && name === curEnv
              const isDefault = !!defaultEnv && name === defaultEnv
              const st = connected ? (state || '') : ''
              return (
                <div key={name}
                  className={`group flex h-7 items-center gap-1.5 px-2.5 text-xs ${
                    connected ? 'bg-accent/40' : 'hover:bg-accent/30'
                  }`}
                  title={connected ? `当前会话 · ${STATE_LABEL[st] || st}` : `环境 ${name}${isDefault ? '(默认)' : ''}`}>
                  <span className={`h-2 w-2 shrink-0 rounded-full ${connected ? (STATE_DOT[st] || 'bg-accent') : isDefault ? 'bg-emerald-500/70' : 'bg-muted-foreground/25'}`} />
                  <span className="min-w-0 flex-1 truncate font-medium">{name}</span>
                  {connected && (
                    <span className="shrink-0 text-[11px] text-muted-foreground">{STATE_LABEL[st] || st}</span>
                  )}
                  {connected ? (
                    <span className="flex shrink-0 items-center gap-0.5 opacity-70 group-hover:opacity-100">
                      <button title="重启会话(断开重连同环境)" disabled={!!busy}
                        onClick={() => void doOp('restart', name)}
                        className="p-1 text-muted-foreground transition-colors hover:text-foreground disabled:pointer-events-none disabled:opacity-30">
                        <RotateCcw className="h-3.5 w-3.5" />
                      </button>
                      <button title="结束会话(彻底断开连接)" disabled={!!busy}
                        onClick={() => void doOp('close')}
                        className="p-1 text-muted-foreground transition-colors hover:text-red-600 dark:hover:text-red-400 disabled:pointer-events-none disabled:opacity-30">
                        <Power className="h-3.5 w-3.5" />
                      </button>
                    </span>
                  ) : (
                    <span className="flex shrink-0 items-center gap-0.5 opacity-70 group-hover:opacity-100">
                      <button title={`切换会话到 ${name}(结束当前 debug 并连接该环境)`} disabled={!!busy}
                        onClick={() => void doOp('switch', name)}
                        className="p-1 text-muted-foreground transition-colors hover:text-foreground disabled:pointer-events-none disabled:opacity-30">
                        <ArrowRightLeft className="h-3.5 w-3.5" />
                      </button>
                    </span>
                  )}
                </div>
              )
            })}
            {err && <div className="px-2.5 pb-1 text-[11px] text-red-600 dark:text-red-400">{err}</div>}
          </div>
        )}

        {/* 区域二:环境变量(TOPENT 设置,空闲会话可改) */}
        <SectionHead title="环境变量" open={open.env} onToggle={() => toggle('env')} />
        {open.env && (
          <div className="py-1">
            <div className="flex items-center gap-1.5 px-2 py-0.5">
              <span className="w-[52px] shrink-0 text-xs text-muted-foreground"
                title="TOPENT(企业编号):T100 运行时的企业环境,作业运行与数据库连接以此为准">TOPENT</span>
              <Input value={topent} disabled={!canSetTopent}
                onChange={(e) => setTopent(e.target.value)}
                onKeyDown={(e) => { if (e.key === 'Enter') void saveTopent() }}
                title="1~3 位数字;留空 = 清除手动设置"
                className="h-6 min-w-0 flex-1 px-1.5 font-mono text-xs" />
              <button disabled={!canSetTopent} onClick={() => void saveTopent()}
                title="保存 TOPENT(下一轮调试启动时采用)"
                className={`flex h-6 shrink-0 items-center px-1.5 text-xs transition-colors disabled:pointer-events-none disabled:opacity-30 ${
                  topentSaved ? 'text-emerald-600 dark:text-emerald-400' : 'hover:bg-accent/60'
                }`}>
                {topentSaved ? <Check className="h-3.5 w-3.5" /> : '保存'}
              </button>
            </div>
            {topentErr && <div className="px-2.5 pb-1 text-[11px] text-red-600 dark:text-red-400">{topentErr}</div>}
          </div>
        )}
      </div>
    </div>
  )
}
