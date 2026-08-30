import { useEffect, useRef, useState } from 'react'
import { RotateCcw, WifiOff, Bug, Globe, FlaskConical, Settings, type LucideIcon } from 'lucide-react'
import { connectWS, useStore } from './store'
import { Toolbar } from './Toolbar'
import { SourceView } from './SourceView'
import { RightPanels, TimelinePanel } from './Panels'
import { WsLogView } from './WsLogView'
import { WsTestView } from './WsTestView'
import { SettingsView } from './SettingsView'
import { StatusBar } from './StatusBar'
import { api } from './api'

// VS Code 风格活动栏图标按钮:选中态灰色底色(无左侧蓝条),未选中透明显灰
function ActivityIcon({ icon: Icon, label, active, onClick }: {
  icon: LucideIcon; label: string; active: boolean; onClick: () => void
}) {
  return (
    <button title={label} onClick={onClick}
      className={`flex h-8 w-8 items-center justify-center rounded-md transition-colors ${
        active ? 'bg-white/15 text-white' : 'text-zinc-500 hover:text-zinc-300'
      }`}>
      <Icon className="h-5 w-5" />
    </button>
  )
}

// 浏览器标签页标题反馈:停站时标题闪烁(仅后台),运行/启动改前缀。
// 停站标题在「⚠ 停站 file:line」与原名之间每秒交替,后台标签会被浏览器高亮提醒。
function useDocTitleBlink() {
  const sessionId = useStore((s) => s.sessionId)
  const state = useStore((s) => s.state)
  const file = useStore((s) => s.stop?.file)
  const line = useStore((s) => s.stop?.line)
  useEffect(() => {
    const base = 'TDict Debug'
    const setT = (t: string) => { document.title = t }
    if (!sessionId) { setT(base); return }
    let timer: number | undefined
    if (state === 'stopped') {
      const loc = file && line ? `${file}:${line}` : '已停站'
      const flashA = `⚠ 停站 ${loc}`
      setT(flashA)
      let toggle = true
      const onVis = () => {
        if (document.hidden) {
          toggle = true
          timer = window.setInterval(() => {
            document.title = toggle ? base : flashA
            toggle = !toggle
          }, 1000)
        } else {
          if (timer) { clearInterval(timer); timer = undefined }
          setT(flashA)
        }
      }
      document.addEventListener('visibilitychange', onVis)
      if (document.hidden) onVis()
      return () => {
        document.removeEventListener('visibilitychange', onVis)
        if (timer) clearInterval(timer)
        setT(base)
      }
    }
    if (state === 'running') setT(`▶ 运行中 · ${base}`)
    else if (state === 'loading') setT(`⏳ 启动中 · ${base}`)
    else setT(base)
  }, [sessionId, state, file, line])
}

export function App() {
  useDocTitleBlink()
  useEffect(() => {
    const close = connectWS()
    void api.status().catch(() => {})
    void useStore.getState().adoptExisting()
    // 调试快捷键:F5 继续 / F10 步过 / F11 步入 / Shift+F11 步出
    const onKey = (e: KeyboardEvent) => {
      const st = useStore.getState()
      if (!st.sessionId) return
      const stopped = st.state === 'stopped'
      if (e.key === 'F5') { e.preventDefault(); if (stopped) void st.control('continue') }
      else if (e.key === 'F10') { e.preventDefault(); if (stopped) void st.control('next') }
      else if (e.key === 'F11') {
        e.preventDefault()
        if (stopped) void st.control(e.shiftKey ? 'finish' : 'step')
      }
    }
    window.addEventListener('keydown', onKey)
    return () => { window.removeEventListener('keydown', onKey); close() }
  }, [])

  const showRight = useStore((s) => s.showRight)
  const showBottom = useStore((s) => s.showBottom)
  const view = useStore((s) => s.view)
  const setView = useStore((s) => s.setView)
  const backendDead = useStore((s) => s.backendDead)
  const wsConnected = useStore((s) => s.wsConnected)
  const launching = useStore((s) => s.launching)
  const sessionId = useStore((s) => s.sessionId)
  // 注意:sessionId 必须是独立 hook 调用——嵌在 || 条件里会在短路时少调用一次,
  // 违反 React Hooks 规则导致 hook 队列错乱(Should have a queue / 无限重渲染)
  const banner = backendDead || (sessionId && !wsConnected ? '服务连接断开,自动重连中…' : '') 
  const restart = useStore((s) => s.restart)

  // 侧边栏/代码区宽度拖拽(记忆到 localStorage)
  const [panelW, setPanelW] = useState(() => {
    const v = Number(localStorage.getItem('tdict.panelW'))
    return v >= 240 && v <= 640 ? v : 320
  })
  const onResizeDown = (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault()
    const el = e.currentTarget
    el.classList.add('dragging')
    const startX = e.clientX
    const startW = panelW
    let latest = startW
    const move = (ev: MouseEvent) => {
      latest = Math.min(640, Math.max(240, startW - (ev.clientX - startX)))
      setPanelW(latest)
    }
    const up = () => {
      el.classList.remove('dragging')
      localStorage.setItem('tdict.panelW', String(latest))
      window.removeEventListener('mousemove', move)
      window.removeEventListener('mouseup', up)
    }
    window.addEventListener('mousemove', move)
    window.addEventListener('mouseup', up)
  }

  // 布局:上 Toolbar / 下 StatusBar 整条;中部 = 活动栏 + [中间列(编辑区/时间线上下) + 整高右面板]
  return (
    <div className="flex h-screen flex-col bg-zinc-950 text-zinc-200">
      <Toolbar />
      {banner && (
        <div className="flex shrink-0 items-center gap-2 border-b border-red-900/60 bg-red-950/40 px-3 py-1.5 text-xs text-red-300">
          <WifiOff className="h-3.5 w-3.5" />
          <span>{banner}</span>
          {backendDead && (
            <button
              className="ml-auto inline-flex items-center gap-1 rounded border border-red-800/80 px-2 py-0.5 hover:bg-red-900/40 disabled:opacity-40"
              disabled={launching}
              onClick={() => void restart()}
            >
              <RotateCcw className="h-3 w-3" />
              {launching ? '重启中…' : '重新启动'}
            </button>
          )}
        </div>
      )}
      <div className="flex min-h-0 flex-1">
        {/* VS Code 风格活动栏:背景与编辑区同色,仅图标高亮区分 */}
        <div className="flex w-10 shrink-0 flex-col items-center gap-1 bg-zinc-950 py-2">
          <ActivityIcon icon={Bug} label="调试" active={view === 'debug'} onClick={() => setView('debug')} />
          <ActivityIcon icon={Globe} label="接口日志" active={view === 'wslogs'} onClick={() => setView('wslogs')} />
          <ActivityIcon icon={FlaskConical} label="服务测试" active={view === 'wstest'} onClick={() => setView('wstest')} />
          <ActivityIcon icon={Settings} label="设置" active={view === 'settings'} onClick={() => setView('settings')} />
        </div>
        {view === 'wslogs' ? (
          <WsLogView />
        ) : view === 'wstest' ? (
          <WsTestView />
        ) : view === 'settings' ? (
          <SettingsView />
        ) : (
          /* 中间列(编辑区 + 时间线)与整高右面板左右并排
             min-w-0 + overflow-hidden:Monaco 会给编辑器写内联像素宽度,
             否则 flex 最小宽度被钉死,收起再展开时编辑区不回缩、右面板被挤出屏幕 */
          <div className="flex min-h-0 min-w-0 flex-1 overflow-hidden p-2 pt-1">
            <div className="flex min-h-0 min-w-0 flex-1 flex-col overflow-hidden">
              <div className="min-h-0 flex-1 overflow-hidden rounded-sm border border-zinc-800">
                <SourceView />
              </div>
              {showBottom && (
                <div className="mt-2 h-48 shrink-0">
                  <TimelinePanel />
                </div>
              )}
            </div>
            {showRight && (
              <>
                  <div className="col-resizer mx-0.5" onMouseDown={onResizeDown} title="拖拽调整代码区与侧边栏宽度" />
                <div style={{ width: panelW }} className="min-h-0 shrink-0">
                  <RightPanels />
                </div>
              </>
            )}
          </div>
        )}
      </div>
      <StatusBar />
    </div>
  )
}
