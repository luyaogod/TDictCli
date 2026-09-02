// 大纲面板:解析当前显示的 4GL 源码(FUNCTION/MAIN + DIALOG/CONSTRUCT/INPUT/INPUT ARRAY 及其子块),
// 树形展示并点击跳行。跳转直接操作编辑器视口(setPosition + revealLineInCenterIfOutsideViewport),
// 不走 revealReq——不受"仅停站可定位"的调试门限制,运行中也可浏览跳转,且不产生任何调试信号。
// 调试页行号需补偿 lineOffset(Monaco 顶部前插空行对齐 DVM 行号)。
import { useMemo, useState, type ReactNode } from 'react'
import { useStore } from './store'
import { editorRef } from './SourceView'
import { parseOutline, type OutlineNode } from './fgloutline'

export function OutlinePanel() {
  const sourceContent = useStore((s) => s.sourceContent)
  const lineOffset = useStore((s) => s.lineOffset)
  const tabs = useStore((s) => s.tabs)
  const activeTab = useStore((s) => s.activeTab)
  const active = tabs.find((t) => t.key === activeTab)
  const isDebug = !active
  const content = isDebug ? (sourceContent || '') : (active!.content || '')
  const offset = isDebug ? lineOffset : 0
  const nodes = useMemo(() => parseOutline(content), [content])
  // 折叠集合:默认全展开(函数直接列出其交互块),点击箭头收起
  const [collapsed, setCollapsed] = useState<Set<string>>(new Set())

  const jump = (line: number) => {
    const ed = editorRef.current
    if (!ed) return
    const L = line + offset
    ed.setPosition({ lineNumber: L, column: 1 })
    ed.revealLineInCenterIfOutsideViewport(L)
  }
  const toggleKey = (key: string) => {
    setCollapsed((prev) => {
      const n = new Set(prev)
      if (n.has(key)) n.delete(key)
      else n.add(key)
      return n
    })
  }

  const row = (node: OutlineNode, depth: number, path: string): ReactNode => {
    const key = `${path}/${node.line}:${node.label}`
    const kids = node.children
    const hasKids = !!kids && kids.length > 0
    const isCollapsed = collapsed.has(key)
    return (
      <div key={key}>
        <div onClick={() => jump(node.line)}
          className={`flex h-6 cursor-pointer items-center gap-1 pr-2 text-xs hover:bg-accent/40 ${
            depth === 0 ? 'font-medium text-foreground' : depth === 1 ? 'text-sky-600 dark:text-sky-400' : 'text-muted-foreground'
          }`}
          style={{ paddingLeft: depth * 14 + 4 }}
          title={`${node.label} (第 ${node.line} 行)`}>
          {hasKids ? (
            <button className="w-3 shrink-0 text-muted-foreground"
              title={isCollapsed ? '展开' : '收起'}
              onClick={(e) => { e.stopPropagation(); toggleKey(key) }}>
              {isCollapsed ? '▸' : '▾'}
            </button>
          ) : (
            <span className="w-3 shrink-0" />
          )}
          <span className="min-w-0 truncate">{node.label}</span>
        </div>
        {hasKids && !isCollapsed && kids!.map((k, i) => row(k, depth + 1, key))}
      </div>
    )
  }

  return (
    <div className="flex h-full min-h-0 w-full flex-col overflow-hidden bg-background">
      <div className="flex h-8 shrink-0 items-center justify-between border-b border-border px-2.5">
        <span className="text-xs font-medium text-muted-foreground">大纲</span>
        <span className="text-[11px] text-muted-foreground">{nodes.length} 个函数</span>
      </div>
      <div className="min-h-0 flex-1 overflow-auto py-1">
        {nodes.length === 0 && (
          <div className="p-2 text-xs text-muted-foreground">当前源码无可识别的函数/交互块</div>
        )}
        {nodes.map((n, i) => row(n, 0, String(i)))}
      </div>
    </div>
  )
}
