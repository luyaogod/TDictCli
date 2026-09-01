// 顶栏(会话徽章 + 面板收展)+ 浮动调试工具条(脱离文档流,手柄拖拽,位置记忆)
import { useState, type ComponentType, type ReactNode } from 'react'
import {
  RedoDot, StepForward, ArrowDownToDot, ArrowUpFromDot,
  RotateCcw, Square, GripVertical,
  SquareChevronRight, SquareChevronLeft, SquareChevronDown, SquareChevronUp,
} from 'lucide-react'
import { useStore } from './store'

// 工具条内图标按钮(无独立边框,悬停浮起,禁用半透明)
function ToolIcon({ icon: Icon, label, onClick, disabled, color = 'text-sky-600 dark:text-sky-400' }: {
  icon: ComponentType<{ className?: string }>
  label: string
  onClick: () => void
  disabled?: boolean
  color?: string
}) {
  return (
    <button
      title={label}
      disabled={disabled}
      onClick={onClick}
      className={`rounded p-1 transition-colors hover:bg-accent/60 disabled:pointer-events-none disabled:opacity-30 ${color}`}
    >
      <Icon className="h-4 w-4" />
    </button>
  )
}

// 浮动工具条容器:脱离文档流,拖左侧手柄移动,位置记忆到 localStorage
function FloatingToolbar({ children }: { children: ReactNode }) {
  const showRight = useStore((s) => s.showRight)
  const [pos, setPos] = useState(() => {
    try {
      const v = JSON.parse(localStorage.getItem('tdict.toolbarPos') || '')
      if (typeof v?.top === 'number' && typeof v?.left === 'number') return v
    } catch { /* 忽略 */ }
    return null // null = 默认定位(代码编辑器右上角,由 CSS 类给值)
  })

  const onDragStart = (e: React.MouseEvent) => {
    e.preventDefault()
    const el = (e.currentTarget.parentElement as HTMLElement)
    const rect = el.getBoundingClientRect()
    const dx = e.clientX - rect.left
    const dy = e.clientY - rect.top
    const move = (ev: MouseEvent) => {
      const left = Math.max(4, Math.min(window.innerWidth - rect.width - 4, ev.clientX - dx))
      const top = Math.max(4, Math.min(window.innerHeight - rect.height - 4, ev.clientY - dy))
      setPos({ top, left })
    }
    const up = () => {
      setPos((p: { top: number; left: number } | null) => {
        localStorage.setItem('tdict.toolbarPos', JSON.stringify(p))
        return p
      })
      window.removeEventListener('mousemove', move)
      window.removeEventListener('mouseup', up)
    }
    window.addEventListener('mousemove', move)
    window.addEventListener('mouseup', up)
  }

  // 默认:代码编辑器右上角(避开右侧栏;右侧栏收起时贴视口右缘)
  const panelW = Number(localStorage.getItem('tdict.panelW')) || 320
  const style = pos
    ? { top: pos.top, left: pos.left }
    : { top: 44, right: showRight ? panelW + 24 : 8 }

  return (
    <div
      className="fixed z-50 inline-flex items-center gap-0.5 rounded-sm border border-border/70 bg-card/95 px-1 py-0.5 shadow-xl"
      style={style}
    >
      <span
        title="拖拽移动工具条"
        onMouseDown={onDragStart}
        className="cursor-grab rounded p-0.5 text-muted-foreground hover:bg-accent/60 hover:text-muted-foreground active:cursor-grabbing"
      >
        <GripVertical className="h-3.5 w-3.5" />
      </span>
      {children}
    </div>
  )
}

export function Toolbar() {
  const launching = useStore((s) => s.launching)
  const quit = useStore((s) => s.quit)
  const restart = useStore((s) => s.restart)
  const control = useStore((s) => s.control)
  const showRight = useStore((s) => s.showRight)
  const showBottom = useStore((s) => s.showBottom)
  const toggleRight = useStore((s) => s.toggleRight)
  const toggleBottom = useStore((s) => s.toggleBottom)
  const view = useStore((s) => s.view)

  const stopped = useStore((s) => s.state) === 'stopped'

  return (
    <>
      {/* 顶栏(融入背景,无边框,低高度;启动控件已移至右侧面板顶部)
          左内边距 48px = 活动栏 40 + 内容区 p-2 8,与代码编辑器左缘对齐 */}
      <div className="flex h-8 shrink-0 items-center gap-2 bg-background pl-12 pr-3">
        {/* 作业编号已由源码区调试页签展示,不再重复 */}

        {/* 右上角:收展右方/下方面板 */}
        <div className="ml-auto inline-flex items-center gap-0.5">
          <ToolIcon icon={showRight ? SquareChevronRight : SquareChevronLeft}
            label={showRight ? '收起右方面板' : '展开右方面板'} color="text-muted-foreground"
            onClick={() => void toggleRight()} />
          <ToolIcon icon={showBottom ? SquareChevronDown : SquareChevronUp}
            label={showBottom ? '收起下方面板' : '展开下方面板'} color="text-muted-foreground"
            onClick={() => void toggleBottom()} />
        </div>
      </div>

      {/* 浮动调试工具条(有会话且在调试视图时显示) */}
      {view === 'debug' && (
        <FloatingToolbar>
          <ToolIcon icon={StepForward} label="继续 (F5) — 运行到下一个断点" disabled={!stopped}
            onClick={() => void control('continue')} />
          <ToolIcon icon={RedoDot} label="步过 (F10)" disabled={!stopped}
            onClick={() => void control('next')} />
          <ToolIcon icon={ArrowDownToDot} label="步入 (F11)" disabled={!stopped}
            onClick={() => void control('step')} />
          <ToolIcon icon={ArrowUpFromDot} label="步出 (finish)" disabled={!stopped}
            onClick={() => void control('finish')} />
          <ToolIcon icon={RotateCcw} label="重新开始 — 结束并重启同一作业" color="text-green-600 dark:text-green-400" disabled={launching}
            onClick={() => void restart()} />
          <ToolIcon icon={Square} label="结束会话 — quit(作业窗口随之关闭)" color="text-red-600 dark:text-red-400"
            onClick={() => void quit()} />
        </FloatingToolbar>
      )}
    </>
  )
}
