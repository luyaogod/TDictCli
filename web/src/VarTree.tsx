// 变量树:DAP 模式专属,局部/全局两棵树,节点懒下钻。
// variablesReference 跨停站失效:每次停站(stopSeq 变化)整体重建,陈旧 ref 求值失败时也整体收起。
import { useCallback, useEffect, useRef, useState } from 'react'
import { api, type VarNode } from './api'
import { useStore } from './store'

interface TNode extends VarNode {
  open: boolean
  loading?: boolean
  children?: TNode[]
}

function toNodes(vars: VarNode[]): TNode[] {
  return vars.map((v) => ({ ...v, open: false }))
}

export function VarTreePanel() {
  const sessionId = useStore((s) => s.sessionId)
  const state = useStore((s) => s.state)
  const mode = useStore((s) => s.mode)
  const stopSeq = useStore((s) => s.stopSeq)
  const selectedFrame = useStore((s) => s.selectedFrame)
  const [roots, setRoots] = useState<TNode[]>([])
  const [, bump] = useState(0)
  const rerender = () => bump((v) => v + 1)
  const stopKeyRef = useRef('')

  const stopped = state === 'stopped' && mode === 'dap' && !!sessionId

  // 停站/切帧后重建:局部变量默认展开,全局变量默认收起(T100 全局动辄上千)
  useEffect(() => {
    if (!sessionId || state !== 'stopped' || mode !== 'dap') {
      if (stopKeyRef.current !== '') {
        stopKeyRef.current = ''
        setRoots([])
      }
      return
    }
    const key = `${sessionId}:${stopSeq}:${selectedFrame}`
    if (key === stopKeyRef.current) return
    stopKeyRef.current = key
    setRoots([
      { name: '局部变量', open: true, loading: true },
      { name: '全局变量', open: false },
    ])
    ;(async () => {
      try {
        const { localsRef, globalsRef } = await api.varRoots(sessionId)
        const locals: TNode = { name: '局部变量', ref: localsRef, open: true, loading: !!localsRef }
        const globals: TNode = { name: '全局变量', ref: globalsRef, open: false }
        if (localsRef > 0) {
          try {
            locals.children = toNodes((await api.varChildren(sessionId, localsRef)).vars)
          } catch { locals.children = [] }
          locals.loading = false
        }
        setRoots([locals, globals])
      } catch {
        setRoots([]) // 会话已不在停站态
      }
    })()
  }, [sessionId, state, mode, stopSeq, selectedFrame])

  const toggle = useCallback(
    async (node: TNode) => {
      if (!sessionId || !node.ref) return
      if (node.open) {
        node.open = false
        rerender()
        return
      }
      node.open = true
      if (!node.children) {
        node.loading = true
        rerender()
        try {
          node.children = toNodes((await api.varChildren(sessionId, node.ref)).vars)
        } catch {
          // 陈旧 ref(已跨停站):整体收起重建
          node.open = false
          node.loading = false
          stopKeyRef.current = ''
          setRoots([])
          return
        }
        node.loading = false
      }
      rerender()
    },
    [sessionId],
  )

  if (!mode) return null
  if (mode !== 'dap') {
    return <div className="px-3 py-2 text-xs text-zinc-500">变量树仅 DAP 模式可用(config.debug.mode = "dap")</div>
  }
  if (!stopped) {
    return <div className="px-3 py-2 text-xs text-zinc-500">{sessionId ? '停站后可查看变量' : '启动调试会话后可用'}</div>
  }
  if (roots.length === 0) return <div className="px-3 py-2 text-xs text-zinc-500">加载中…</div>
  return (
    <div className="px-1 py-1 text-xs overflow-auto max-h-[45vh]">
      {roots.map((r, i) => (
        <TreeNode key={i} node={r} depth={0} toggle={toggle} />
      ))}
    </div>
  )
}

function TreeNode({ node, depth, toggle }: { node: TNode; depth: number; toggle: (n: TNode) => void }) {
  const expandable = !!node.ref
  return (
    <>
      <div
        className={`flex items-baseline gap-1 rounded px-1 hover:bg-white/5 ${expandable ? 'cursor-pointer' : 'cursor-default'}`}
        style={{ paddingLeft: depth * 14 + 4 }}
        onClick={() => expandable && toggle(node)}
      >
        <span className={`w-3 shrink-0 text-zinc-500 ${expandable ? '' : 'opacity-0'}`}>
          {node.loading ? '…' : node.open ? '▾' : '▸'}
        </span>
        <span className="text-sky-300 whitespace-pre">{node.name}</span>
        {node.value !== undefined && (
          <>
            <span className="text-zinc-500 whitespace-pre"> = </span>
            <span className="text-zinc-200 break-all">{node.value || '(空)'}</span>
          </>
        )}
        {node.type && <span className="ml-1 shrink-0 text-zinc-600">{node.type}</span>}
      </div>
      {node.open && node.children?.map((c, j) => (
        <TreeNode key={j} node={c} depth={depth + 1} toggle={toggle} />
      ))}
    </>
  )
}
