// 接口日志视图:Chrome DevTools Network 风格 —— 左侧列表(状态/服务/作业/时间/耗时)
// + 右侧选中行详情(基本信息 / Request / Response)+ 一键重放调试
// 查询条件对齐 awsq990 主查询 QBE:服务名(wsfa001)+ 开始时间范围(wsfa003);仅失败为本工具扩展
import { useEffect, useState } from 'react'
import { Bug, ChevronLeft, ChevronRight, RefreshCw } from 'lucide-react'
import { useStore } from './store'
import type { WSLogItem } from './api'

export function WsLogView() {
  const wsLogs = useStore((s) => s.wsLogs)
  const loading = useStore((s) => s.wsLogsLoading)
  const page = useStore((s) => s.wsLogsPage)
  const hasMore = useStore((s) => s.wsLogsHasMore)
  const sel = useStore((s) => s.wsLogSel)
  const content = useStore((s) => s.wsLogContent)
  const tab = useStore((s) => s.wsLogTab)
  const err = useStore((s) => s.wsLogErr)
  const loadWsLogs = useStore((s) => s.loadWsLogs)
  const selectWsLog = useStore((s) => s.selectWsLog)
  const setWsLogTab = useStore((s) => s.setWsLogTab)
  const replayDebug = useStore((s) => s.replayDebug)
  const sessionId = useStore((s) => s.sessionId)
  const [service, setService] = useState('')
  const [onlyFail, setOnlyFail] = useState(false)
  // 默认过滤条件:当天(awsq990 查当天日志是最常用场景)
  const today = () => {
    const d = new Date()
    return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
  }
  const [from, setFrom] = useState(today)
  const [to, setTo] = useState(today)

  useEffect(() => { void loadWsLogs('', false, 1, from, to) }, [loadWsLogs, from, to])

  const doLoad = (p = 1) => void loadWsLogs(service, onlyFail, p, from, to)

  // 详情面板宽度拖拽(记忆到 localStorage)
  const [detailW, setDetailW] = useState(() => {
    const v = Number(localStorage.getItem('tdict.wslogW'))
    return v >= 320 && v <= 720 ? v : 460
  })
  const onDetailResizeDown = (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault()
    const el = e.currentTarget
    el.classList.add('dragging')
    const startX = e.clientX
    const startW = detailW
    let latest = startW
    const move = (ev: MouseEvent) => {
      latest = Math.min(720, Math.max(320, startW + (ev.clientX - startX)))
      setDetailW(latest)
    }
    const up = () => {
      el.classList.remove('dragging')
      localStorage.setItem('tdict.wslogW', String(latest))
      window.removeEventListener('mousemove', move)
      window.removeEventListener('mouseup', up)
    }
    window.addEventListener('mousemove', move)
    window.addEventListener('mouseup', up)
  }

  return (
    <div className="flex min-h-0 flex-1 flex-col p-2 pt-1">
      {/* 工具条:服务名 + 时间范围(awsq990 QBE 同款条件)+ 仅失败 + 刷新 + 翻页 */}
      <div className="mb-2 flex shrink-0 flex-wrap items-center gap-2">
        <input
          value={service}
          onChange={(e) => setService(e.target.value)}
          onKeyDown={(e) => { if (e.key === 'Enter') doLoad(1) }}
          placeholder="按服务名过滤,如 icd.erp.wo*(回车生效)"
          className="h-7 w-56 rounded-md border border-border bg-background px-2 text-xs text-foreground placeholder:text-muted-foreground focus:border-border focus:outline-none"
        />
        <label className="flex items-center gap-1 text-xs text-muted-foreground">
          开始时间
          <input type="date" value={from} onChange={(e) => setFrom(e.target.value)}
            className="h-7 rounded-md border border-border bg-background px-1.5 text-xs text-foreground focus:border-border focus:outline-none" />
          ~
          <input type="date" value={to} onChange={(e) => setTo(e.target.value)}
            className="h-7 rounded-md border border-border bg-background px-1.5 text-xs text-foreground focus:border-border focus:outline-none" />
        </label>
        <label className="flex cursor-pointer items-center gap-1.5 text-xs text-muted-foreground">
          <input type="checkbox" checked={onlyFail} onChange={(e) => { setOnlyFail(e.target.checked); setTimeout(() => doLoad(1), 0) }}
            className="h-3.5 w-3.5 accent-red-500" />
          仅失败
        </label>
        <button onClick={() => doLoad(1)} disabled={loading} title="重新加载"
          className="rounded p-1.5 text-muted-foreground transition-colors hover:bg-accent hover:text-foreground disabled:opacity-40">
          <RefreshCw className={`h-4 w-4 ${loading ? 'animate-spin' : ''}`} />
        </button>
        {/* 分页 */}
        <div className="ml-auto flex items-center gap-1 text-xs text-muted-foreground">
          <span>{wsLogs.length} 条</span>
          <button onClick={() => doLoad(page - 1)} disabled={loading || page <= 1} title="上一页"
            className="rounded p-1 hover:bg-accent hover:text-foreground disabled:opacity-30">
            <ChevronLeft className="h-4 w-4" />
          </button>
          <span className="w-10 text-center font-mono">第 {page} 页</span>
          <button onClick={() => doLoad(page + 1)} disabled={loading || !hasMore} title="下一页"
            className="rounded p-1 hover:bg-accent hover:text-foreground disabled:opacity-30">
            <ChevronRight className="h-4 w-4" />
          </button>
        </div>
        {sessionId && <span className="text-xs text-amber-600/80 dark:text-amber-500/80">调试会话进行中(重放前需先结束)</span>}
      </div>

      {err && (
        <div className="mb-2 shrink-0 rounded border border-red-500/20 bg-red-500/10 px-3 py-1.5 text-xs text-red-600 dark:text-red-600 dark:text-red-400">{err}</div>
      )}

      {/* 左右布局:左列列表,右列详情(列宽固定,容器变窄时列表内部横向滚动) */}
      <div className="flex min-h-0 flex-1 gap-2">
        {/* 列表:表头与行同处一个滚动容器,横向滚动时表头跟着滚不错位 */}
        <div className="flex min-w-0 flex-1 flex-col overflow-hidden rounded-sm border border-border bg-card/60">
          <div className="min-h-0 flex-1 overflow-auto">
            <div className="sticky top-0 flex h-7 items-center gap-2 border-b border-border bg-card px-2 text-[11px] font-medium text-muted-foreground">
              <span className="w-12 shrink-0">状态</span>
              <span className="w-44 shrink-0">服务</span>
              <span className="w-28 shrink-0">作业</span>
              <span className="w-32 shrink-0">开始时间</span>
              <span className="w-16 shrink-0">耗时(s)</span>
              <span className="min-w-0 flex-1">错误描述</span>
            </div>
            {wsLogs.length === 0 && !loading && (
              <div className="p-4 text-center text-xs text-muted-foreground">暂无日志记录</div>
            )}
            {wsLogs.map((it, i) => (
              <div key={it.rowid} onClick={() => void selectWsLog(it)}
                title="点击在右侧查看请求/响应报文"
                className={`flex h-7 cursor-pointer items-center gap-2 px-2 text-xs ${
                  sel?.rowid === it.rowid ? 'bg-sky-500/10' : i % 2 ? 'bg-card/40 hover:bg-accent/40' : 'hover:bg-accent/40'
                }`}>
                <span className={`w-12 shrink-0 font-mono text-[11px] ${it.code === '000' ? 'text-emerald-600 dark:text-emerald-500' : it.code ? 'text-red-500' : 'text-muted-foreground'}`}>
                  {it.code || '-'}
                </span>
                <span className="w-44 shrink-0 truncate text-foreground" title={it.service}>{it.service}</span>
                <span className="w-28 shrink-0 truncate text-sky-600 dark:text-sky-400" title={it.job}>{it.job}</span>
                <span className="w-32 shrink-0 font-mono text-[11px] text-muted-foreground">{it.start}</span>
                <span className="w-16 shrink-0 font-mono text-[11px] text-muted-foreground">{it.duration}</span>
                <span className="min-w-0 flex-1 truncate text-red-600 dark:text-red-400/90" title={it.errMsg}>{it.errMsg}</span>
              </div>
            ))}
          </div>
        </div>

        {/* 分隔条:拖拽调整详情面板宽度 */}
        <div className="col-resizer mx-0.5 self-stretch" onMouseDown={onDetailResizeDown} title="拖拽调整详情面板宽度" />

        {/* 详情(右列) */}
        <div style={{ width: detailW }} className="flex shrink-0 flex-col overflow-hidden rounded-sm border border-border bg-card/60">
          {sel ? (
            <>
              <div className="flex h-8 shrink-0 items-center gap-1 border-b border-border px-2">
                {([['info', '基本信息'], ['request', 'Request'], ['response', 'Response']] as const).map(([k, label]) => (
                  <button key={k} onClick={() => setWsLogTab(k)}
                    className={`rounded px-2 py-0.5 text-xs ${tab === k ? 'bg-accent text-foreground' : 'text-muted-foreground hover:text-foreground'}`}>
                    {label}
                  </button>
                ))}
                <button
                  onClick={() => void replayDebug(sel)}
                  disabled={!!sessionId}
                  title="用该日志的报文重放此接口调用并进入调试(T100 r.dg 同款)"
                  className="ml-auto inline-flex items-center gap-1 rounded border border-emerald-500/20 px-2 py-0.5 text-xs text-emerald-600 dark:text-emerald-400 hover:bg-emerald-500/10 disabled:opacity-40"
                >
                  <Bug className="h-3.5 w-3.5" />
                  调试此调用
                </button>
              </div>
              <div className="min-h-0 flex-1 overflow-auto">
                {tab === 'info' && <InfoBody item={sel} />}
                {tab === 'request' && <PayloadBody text={content?.request} loading={!content} />}
                {tab === 'response' && <PayloadBody text={content?.response} loading={!content} />}
              </div>
            </>
          ) : (
            <div className="flex flex-1 items-center justify-center p-4 text-center text-xs text-muted-foreground">
              点击左侧日志行
              <br />
              在此处查看基本信息 / 请求 / 响应报文
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

// 基本信息(DevTools Headers 风格的两列键值)
function InfoBody({ item }: { item: WSLogItem }) {
  const rows: [string, string][] = [
    ['服务', item.service],
    ['作业', item.job],
    ['返回码', item.code],
    ['开始时间', item.start],
    ['耗时(s)', item.duration],
    ['进程 PID', item.pid],
    ['错误描述', item.errMsg],
    ['请求报文文件', item.reqPath],
    ['响应报文文件', item.rspPath],
    ['请求大小(字节)', item.reqSize],
    ['响应大小(字节)', item.rspSize],
  ]
  return (
    <div className="p-2 text-xs">
      {rows.map(([k, v]) => (
        <div key={k} className="flex gap-2 border-b border-border/60 py-1">
          <span className="w-28 shrink-0 text-muted-foreground">{k}</span>
          <span className="min-w-0 flex-1 break-all text-foreground">{v || '-'}</span>
        </div>
      ))}
    </div>
  )
}

// 报文内容(monospace pre)
function PayloadBody({ text, loading }: { text?: string; loading: boolean }) {
  if (loading) return <div className="p-3 text-xs text-muted-foreground">加载报文中…</div>
  if (!text) return <div className="p-3 text-xs text-muted-foreground">无报文(超过入库大小上限且源文件已清理)</div>
  return (
    <pre className="whitespace-pre-wrap break-all p-2 font-mono text-[11px] leading-5 text-foreground">{text}</pre>
  )
}
