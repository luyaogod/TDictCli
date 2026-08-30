// 右侧面板:运行/调试(VS Code 风格)+ 调用栈 / 变量监视 / 断点(一体化容器,分割线分区)
import * as React from 'react'
import * as AccordionPrimitive from '@radix-ui/react-accordion'
import { useEffect, useRef, useState } from 'react'
import { Play } from 'lucide-react'
import { useStore } from './store'
import { Badge, Button, Input } from './ui'
import { parseFglTree, type TNode } from './fglparse'
import { VarTreeNodes } from './VarTreeUi'

// 运行/调试区(VS Code Run and Debug 同款):绿色运行按钮 + 目标输入框。
// 输入作业编号或程序名(模块自动解析);也支持「模块/作业」显式指定模块。
function LaunchSection() {
  const launch = useStore((s) => s.launch)
  const launching = useStore((s) => s.launching)
  const sessionId = useStore((s) => s.sessionId)
  const [v, setV] = useState(() => localStorage.getItem('tdict.launchTarget') || 'bsft001_wf')
  if (sessionId) return null // 会话进行中隐藏(重启走工具条「重新开始」)
  const doLaunch = () => {
    const t = v.trim()
    if (!t || launching) return
    localStorage.setItem('tdict.launchTarget', t)
    const i = t.indexOf('/')
    if (i > 0) void launch(t.slice(0, i).trim(), t.slice(i + 1).trim())
    else void launch('', t)
  }
  return (
    <div className="shrink-0 border-b border-zinc-800">
      <div className="flex h-8 items-center px-2.5 text-xs font-medium text-zinc-400">运行</div>
      <div className="flex items-center gap-1 px-1.5 pb-1.5">
        <button
          title="启动调试会话(Enter 同效)"
          disabled={launching || !v.trim()}
          onClick={doLaunch}
          className="rounded p-1 transition-colors hover:bg-zinc-700/60 disabled:pointer-events-none disabled:opacity-30"
        >
          <Play className="h-4 w-4 text-green-500" fill="currentColor" />
        </button>
        <Input
          value={v}
          onChange={(e) => setV(e.target.value)}
          onKeyDown={(e) => { if (e.key === 'Enter') doLaunch() }}
          placeholder="作业编号,如 bsft001_wf 或 asf/bsft001_wf"
          className="h-6 flex-1 px-1.5 text-xs"
        />
      </div>
    </div>
  )
}

// shadcn 风格 Accordion 基础组件(Radix)
const Accordion = AccordionPrimitive.Root

const AccordionItem = React.forwardRef<
  React.ElementRef<typeof AccordionPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof AccordionPrimitive.Item>
>(({ className, ...props }, ref) => (
  <AccordionPrimitive.Item
    ref={ref}
    className={`flex min-h-0 flex-col overflow-hidden border-b border-zinc-800 data-[state=closed]:flex-none last:border-b-0 ${className || ''}`}
    {...props}
  />
))
AccordionItem.displayName = 'AccordionItem'

const AccordionTrigger = React.forwardRef<
  React.ElementRef<typeof AccordionPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof AccordionPrimitive.Trigger>
>(({ className, children, ...props }, ref) => (
  <AccordionPrimitive.Trigger
    ref={ref}
    className={`flex h-8 shrink-0 items-center justify-between px-2.5 text-xs font-medium text-zinc-400 transition-colors hover:bg-zinc-800/50 hover:text-zinc-300 [&[data-state=open]>svg]:rotate-180 ${className || ''}`}
    {...props}
  >
    {children}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="14"
      height="14"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      className="shrink-0 text-zinc-600 transition-transform duration-200"
    >
      <path d="m6 9 6 6 6-6" />
    </svg>
  </AccordionPrimitive.Trigger>
))
AccordionTrigger.displayName = 'AccordionTrigger'

const AccordionContent = React.forwardRef<
  React.ElementRef<typeof AccordionPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof AccordionPrimitive.Content>
>(({ className, children, ...props }, ref) => (
  <AccordionPrimitive.Content
    ref={ref}
    className={`min-h-0 flex-1 overflow-auto ${className || ''}`}
    {...props}
  >
    {children}
  </AccordionPrimitive.Content>
))
AccordionContent.displayName = 'AccordionContent'

export function RightPanels() {
  return (
    // 一体化面板:单个圆角容器,各区块间用分割线区分(VS Code 风格)
    <div className="flex h-full min-h-0 w-full flex-col overflow-hidden rounded-sm border border-zinc-800 bg-zinc-900/60">
      <LaunchSection />
      <Accordion
        type="multiple"
        defaultValue={['stack', 'autovars', 'watches', 'bps']}
        className="flex h-full min-h-0 w-full flex-col"
      >
        <AccordionItem value="stack">
          <AccordionTrigger>
            <StackTitle />
          </AccordionTrigger>
          <AccordionContent>
            <StackBody />
          </AccordionContent>
        </AccordionItem>

        <AccordionItem value="autovars">
          <AccordionTrigger>
            <AutovarsTitle />
          </AccordionTrigger>
          <AccordionContent>
            <AutovarsBody />
          </AccordionContent>
        </AccordionItem>

        <AccordionItem value="watches">
          <AccordionTrigger>
            <WatchesTitle />
          </AccordionTrigger>
          <AccordionContent>
            <WatchesBody />
          </AccordionContent>
        </AccordionItem>

        <AccordionItem value="bps">
          <AccordionTrigger>
            <BpsTitle />
          </AccordionTrigger>
          <AccordionContent>
            <BpsBody />
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  )
}

// ---- 各面板标题(折叠时也始终可见,计数实时) ----

function StackTitle() {
  const n = useStore((s) => s.frames.length)
  return <span className="pl-0.5">调用栈 ({n})</span>
}

function AutovarsTitle() {
  const n = useStore((s) => s.autovars.length)
  return <span className="pl-0.5">自动变量 ({n})</span>
}

function WatchesTitle() {
  const n = useStore((s) => s.watches.length)
  return <span className="pl-0.5">变量监视 ({n})</span>
}

function BpsTitle() {
  const n = useStore((s) => s.breakpoints.length)
  return <span className="pl-0.5">断点 ({n})</span>
}

// ---- 各面板内容 ----

function StackBody() {
  const frames = useStore((s) => s.frames)
  const selected = useStore((s) => s.selectedFrame)
  const selectFrame = useStore((s) => s.selectFrame)
  const stopped = useStore((s) => s.state === 'stopped')
  if (frames.length === 0) return null
  return (
    <div>
      {frames.map((f) => (
        <div
          key={f.idx}
          className={`cursor-pointer border-b border-zinc-800/60 px-2 py-1 text-xs hover:bg-zinc-800/40 ${
            selected === f.idx ? 'bg-sky-950/60' : ''
          } ${!stopped ? 'opacity-60' : ''}`}
          title={stopped ? `点击切到该帧上下文(print/locals 随之切换)` : '停站后可切换栈帧'}
          onClick={() => stopped && void selectFrame(f.idx)}
        >
          <span className="mr-1.5 text-zinc-600">#{f.idx}</span>
          <span className="text-sky-400">{f.func}</span>
          <span className="ml-1.5 text-zinc-500">{f.file}:{f.line}</span>
        </div>
      ))}
    </div>
  )
}

// 自动变量:停站后从当前源码窗自动提取变量并求值(只读,可一键转为监视)
function AutovarsBody() {
  const autovars = useStore((s) => s.autovars)
  const addWatch = useStore((s) => s.addWatch)
  if (autovars.length === 0) {
    return null
  }
  return (
    <div>
      {autovars.map((v) => (
        <div key={v.expr} className="group flex items-start gap-1 border-b border-zinc-800/60 px-2 py-1 text-xs">
          <button
            className="shrink-0 text-zinc-700 opacity-0 transition-opacity hover:text-emerald-400 group-hover:opacity-100"
            title="加入变量监视"
            onClick={() => void addWatch(v.expr)}
          >
            +
          </button>
          <div className="min-w-0 flex-1">
            <div className="text-zinc-400">{v.expr}</div>
            <div className="whitespace-pre-wrap break-all text-emerald-400">{v.value}</div>
          </div>
        </div>
      ))}
    </div>
  )
}

// 监视条目视图:watches 刷新时重建;RECORD/ARRAY 文本解析成树,标量保持原文本
interface WatchView { expr: string; error?: string; text?: string; root?: TNode }

function WatchesBody() {
  const watches = useStore((s) => s.watches)
  const addWatch = useStore((s) => s.addWatch)
  const removeWatch = useStore((s) => s.removeWatch)
  const doPrint = useStore((s) => s.doPrint)
  const [expr, setExpr] = useState('')
  const [views, setViews] = useState<WatchView[]>([])
  useEffect(() => {
    setViews(watches.map((w) => {
      if (w.error) return { expr: w.expr, error: w.error }
      const kids = parseFglTree(w.value || '')
      if (kids) {
        const record = kids.some((k) => !k.name.startsWith('['))
        return { expr: w.expr, root: { name: w.expr, open: false, children: kids, type: record ? 'RECORD' : `ARRAY[${kids.length}]` } }
      }
      return { expr: w.expr, text: w.value }
    }))
  }, [watches])
  return (
    <div>
      <div className="flex gap-1 p-1.5">
        <Input
          value={expr}
          onChange={(e) => setExpr(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && expr.trim()) { void addWatch(expr.trim()); setExpr('') }
          }}
          placeholder="表达式,如 lp_str / g_qryparam.*"
          className="h-7 flex-1 text-xs"
        />
        <Button size="sm" variant="outline" disabled={!expr.trim()}
          onClick={() => { void doPrint(expr.trim()); setExpr('') }}>
          求值
        </Button>
      </div>
      {views.map((v) => (
        <div key={v.expr} className="flex items-start gap-1 border-b border-zinc-800/60 px-2 py-1 text-xs">
          <button className="shrink-0 text-zinc-600 hover:text-red-400" onClick={() => removeWatch(v.expr)}>×</button>
          <div className="min-w-0 flex-1">
            {v.root ? (
              <VarTreeNodes nodes={[v.root]} />
            ) : (
              <>
                <div className="text-zinc-400">{v.expr}</div>
                <div className={`whitespace-pre-wrap break-all ${v.error ? 'text-red-400' : 'text-emerald-400'}`}>
                  {v.error || v.text}
                </div>
              </>
            )}
          </div>
        </div>
      ))}
    </div>
  )
}

function BpsBody() {
  const breakpoints = useStore((s) => s.breakpoints)
  const removeBreakpoint = useStore((s) => s.removeBreakpoint)
  const toggleBPEnabled = useStore((s) => s.toggleBPEnabled)
  const jumpToBp = useStore((s) => s.jumpToBp)
  if (breakpoints.length === 0) return null
  return (
    <div>
      {breakpoints.map((b) => (
        <div key={b.num} className={`flex items-center gap-2 border-b border-zinc-800/60 px-2 py-1 text-xs ${b.enabled ? '' : 'opacity-50'}`}>
          <input
            type="checkbox"
            checked={b.enabled}
            onChange={(e) => void toggleBPEnabled(b.num, e.target.checked)}
            className="h-3.5 w-3.5 shrink-0 cursor-pointer accent-red-500"
            title={b.enabled ? '取消勾选禁用断点' : '勾选启用断点'}
          />
          <span
            className="min-w-0 flex-1 cursor-pointer truncate text-zinc-300 hover:text-sky-300 hover:underline"
            title={`跳转到 ${b.file}:${b.line}`}
            onClick={() => void jumpToBp(b)}
          >
            {b.func ? <span className="text-sky-400">{b.func} </span> : null}
            {b.file}:{b.line}
            {!b.enabled && <span className="ml-1 text-[10px] text-zinc-500">(已禁用)</span>}
          </span>
          <button className="text-zinc-600 hover:text-red-400" onClick={() => void removeBreakpoint(b.num)}>×</button>
        </div>
      ))}
    </div>
  )
}

export function TimelinePanel() {
  const [tab, setTab] = useState<'timeline' | 'raw'>('timeline')
  const timeline = useStore((s) => s.timeline)
  const rawLog = useStore((s) => s.rawLog)
  const sessionId = useStore((s) => s.sessionId)
  const state = useStore((s) => s.state)
  const sendRaw = useStore((s) => s.sendRaw)
  const bodyRef = useRef<HTMLDivElement | null>(null)
  const [cmd, setCmd] = useState('')
  const [hist, setHist] = useState<string[]>([])
  const [histIdx, setHistIdx] = useState(-1)
  // 新条目到达时自动滚动到底部(最新操作始终可见)
  useEffect(() => {
    const el = bodyRef.current
    if (el) el.scrollTop = el.scrollHeight
  }, [tab, timeline.length, rawLog.length])
  const kindTone: Record<string, string> = {
    stop: 'text-yellow-400', warn: 'text-red-400', command: 'text-sky-400', info: 'text-zinc-400',
  }
  // fgldb 命令直通输入(复刻原版 fgldeb Ctrl+D 子画面);仅停站时可发
  const canSend = !!sessionId && state === 'stopped'
  const submit = () => {
    const c = cmd.trim()
    if (!c || !canSend) return
    setHist((h) => [...h.filter((x) => x !== c), c].slice(-50))
    setHistIdx(-1)
    setCmd('')
    void sendRaw(c)
  }
  const onCmdKey = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') { e.preventDefault(); submit(); return }
    // ↑/↓ 翻命令历史(最新在末尾,↓ 可回到空)
    if (e.key === 'ArrowUp' || e.key === 'ArrowDown') {
      e.preventDefault()
      if (hist.length === 0) return
      let idx = histIdx
      if (e.key === 'ArrowUp') idx = histIdx < 0 ? hist.length - 1 : Math.max(0, histIdx - 1)
      else idx = histIdx < 0 ? -1 : (histIdx + 1 >= hist.length ? -1 : histIdx + 1)
      setHistIdx(idx)
      setCmd(idx >= 0 ? hist[idx] : '')
    }
  }
  return (
    <div className="flex h-full min-h-0 flex-col rounded-sm border border-zinc-800 bg-zinc-900/60">
      <div className="flex h-8 shrink-0 items-center gap-2 border-b border-zinc-800 px-2">
        <button
          className={`rounded px-2 py-0.5 text-xs ${tab === 'timeline' ? 'bg-zinc-800 text-zinc-100' : 'text-zinc-500 hover:text-zinc-300'}`}
          onClick={() => setTab('timeline')}
        >
          操作时间线
        </button>
        <button
          className={`rounded px-2 py-0.5 text-xs ${tab === 'raw' ? 'bg-zinc-800 text-zinc-100' : 'text-zinc-500 hover:text-zinc-300'}`}
          onClick={() => setTab('raw')}
        >
          原始协议流 ({rawLog.length})
        </button>
      </div>
      <div ref={bodyRef} className="min-h-0 flex-1 overflow-auto p-1 font-mono text-[11px] leading-5">
        {tab === 'timeline' && (
          <>
            {timeline.length === 0 && <div className="p-2 text-zinc-600">暂无记录</div>}
            {timeline.map((t, i) => (
              <div key={i} className="flex gap-2">
                <span className="shrink-0 text-zinc-600">{t.time}</span>
                <Badge tone={t.origin === 'ai' ? 'blue' : t.origin === 'human' ? 'green' : 'gray'} className="mt-0.5 h-4 shrink-0">
                  {t.origin === 'ai' ? 'AI' : t.origin === 'human' ? '人' : '系统'}
                </Badge>
                <span className={`min-w-0 ${kindTone[t.kind] || 'text-zinc-400'}`}>{t.text}</span>
              </div>
            ))}
          </>
        )}
        {tab === 'raw' && (
          <>
            {rawLog.map((l, i) => (
              <div key={i} className="whitespace-pre-wrap break-all text-zinc-500">{l || ' '}</div>
            ))}
          </>
        )}
      </div>
      {tab === 'raw' && (
        <div className="flex h-8 shrink-0 items-center gap-1 border-t border-zinc-800 px-2">
          <span className="shrink-0 font-mono text-xs text-sky-400">{canSend ? '(fgldb)' : '—'}</span>
          <input
            value={cmd}
            onChange={(e) => setCmd(e.target.value)}
            onKeyDown={onCmdKey}
            disabled={!canSend}
            placeholder={canSend ? 'fgldb 命令,如 print lp_str / info breakpoints(↑↓ 历史)' : '需停站后才能发送命令'}
            className="h-6 min-w-0 flex-1 rounded-sm border border-zinc-800 bg-zinc-950 px-2 font-mono text-xs text-zinc-200 placeholder:text-zinc-600 focus:border-zinc-600 focus:outline-none disabled:opacity-50"
          />
          <button
            className="shrink-0 rounded-sm border border-zinc-700 px-2 py-0.5 text-xs text-zinc-300 transition-colors hover:bg-zinc-800 disabled:pointer-events-none disabled:opacity-40"
            disabled={!canSend || !cmd.trim()}
            onClick={submit}
          >
            发送
          </button>
        </div>
      )}
    </div>
  )
}
