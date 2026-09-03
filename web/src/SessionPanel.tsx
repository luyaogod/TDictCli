// 「会话」sheet:管理全局单一常驻会话(每个环境一条,通常就是默认环境那一条)。
// 会话 = 与某 T100 环境的 SSH+登录态,debug 只是在其上跑一轮;结束调试/程序退出都只回空闲。
// 本面板负责 切换(换到另一环境并重连)/ 重启(断开重连同环境)/ 结束(彻底释放连接);
// 这些操作都会先结束当前 debug(由后端收口)。行只显示环境名 + 状态,保持精简。
import { useCallback, useEffect, useState } from 'react'
import { ArrowRightLeft, RefreshCw, RotateCcw, Power } from 'lucide-react'
import { useStore } from './store'
import { api } from './api'

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

export function SessionPanel() {
  const sessionId = useStore((s) => s.sessionId)
  const sessionEnv = useStore((s) => s.sessionEnv)
  const state = useStore((s) => s.state)
  const sessionOp = useStore((s) => s.sessionOp)

  const [envNames, setEnvNames] = useState<string[]>([])
  const [defaultEnv, setDefaultEnv] = useState('')
  const [busy, setBusy] = useState<string | null>(null)
  const [err, setErr] = useState('')

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

  return (
    <div className="flex h-full min-h-0 w-full flex-col overflow-hidden bg-background">
      <div className="flex h-8 shrink-0 items-center justify-between border-b border-border px-2.5">
        <span className="text-xs font-medium text-muted-foreground">会话</span>
        <button title="刷新环境/会话状态" onClick={load}
          className="p-0.5 text-muted-foreground transition-colors hover:text-foreground">
          <RefreshCw className="h-3.5 w-3.5" />
        </button>
      </div>

      <div className="min-h-0 flex-1 overflow-auto py-1">
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
        <p className="px-2.5 pb-1 pt-2 text-[11px] leading-relaxed text-muted-foreground">
          会话常驻复用:「结束调试」只收口本轮运行,SSH/登录态保留;再次启动或换作业免重新登录。
          {connectedEnv.length > 0 && !STATE_LABEL[state] && state === 'loading' ? '正在连接…' : ''}
        </p>
        {err && <div className="px-2.5 pb-1 text-[11px] text-red-600 dark:text-red-400">{err}</div>}
      </div>
    </div>
  )
}
