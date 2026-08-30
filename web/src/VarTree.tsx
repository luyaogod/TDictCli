// 变量树:DAP 模式专属,局部/全局两棵树,节点懒下钻。
// variablesReference 跨停站失效:每次停站(stopSeq 变化)整体重建,陈旧 ref 求值失败时也整体收起。
// 被 250 字符截断("..."结尾)的字符串节点可点"全文",按 4GL 子串下标分段求值拼接。
import { useCallback, useEffect, useRef, useState } from 'react'
import { api, type VarNode } from './api'
import { useStore } from './store'

export interface TNode extends VarNode {
  path?: string // 相对求值路径(g_qryparam.cond / g_argv[1]),全文求值用
  open: boolean
  loading?: boolean
  children?: TNode[]
  openFull?: boolean
  full?: string
  fullErr?: string
}
const simpleChainRe = /^[A-Za-z_]\w*(\.\w+)*$/

// 子节点路径:record 字段用点连接,数组元素名形如 "[1]" 直接拼接
function childPath(parent: string, name: string): string {
  if (!parent) return name
  return name.startsWith('[') ? parent + name : parent + '.' + name
}

function toNodes(vars: VarNode[], parent: string): TNode[] {
  return vars.map((v) => ({ ...v, open: false, path: childPath(parent, v.name) }))
}

// useVarTree 变量树下钻/取全文的通用操作(变量树与变量监视共用)。
// 就地改节点对象后 rerender;陈旧 ref(跨停站)由调用方重建节点。
export function useVarTree(sessionId: string | null) {
  const [, bump] = useState(0)
  const rerender = useCallback(() => bump((v) => v + 1), [])

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
          node.children = toNodes((await api.varChildren(sessionId, node.ref)).vars, node.path || '')
        } catch {
          node.open = false // 陈旧 ref:调用方负责整体重建
        }
        node.loading = false
      }
      rerender()
    },
    [sessionId, rerender],
  )

  const toggleFull = useCallback(
    async (node: TNode) => {
      if (!sessionId) return
      if (node.openFull) {
        node.openFull = false
        rerender()
        return
      }
      node.openFull = true
      if (node.full || node.fullErr) rerender()
      try {
        const { value } = await api.printFull(sessionId, node.path || node.name)
        node.full = value
      } catch (e: any) {
        node.fullErr = e.message
      }
      rerender()
    },
    [sessionId, rerender],
  )

  return { toggle, toggleFull, rerender }
}

export function VarTreePanel() {
  const sessionId = useStore((s) => s.sessionId)
  const state = useStore((s) => s.state)
  const mode = useStore((s) => s.mode)
  const stopSeq = useStore((s) => s.stopSeq)
  const selectedFrame = useStore((s) => s.selectedFrame)
  const { toggle, toggleFull } = useVarTree(sessionId)
  const [roots, setRoots] = useState<TNode[]>([])
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
            locals.children = toNodes((await api.varChildren(sessionId, localsRef)).vars, '')
          } catch { locals.children = [] }
          locals.loading = false
        }
        setRoots([locals, globals])
      } catch {
        setRoots([]) // 会话已不在停站态
      }
    })()
  }, [sessionId, state, mode, stopSeq, selectedFrame])

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
        <TreeNode key={i} node={r} depth={0} toggle={toggle} toggleFull={toggleFull} />
      ))}
    </div>
  )
}

export function TreeNode({ node, depth, toggle, toggleFull }: {
  node: TNode
  depth: number
  toggle: (n: TNode) => void
  toggleFull: (n: TNode) => void
}) {
  const expandable = !!node.ref
  // 值被适配器截断(250 字符 + "...")且是可子串求值的变量链:提供全文展开
  const truncatable = !!node.value?.endsWith('...') && !!node.path && simpleChainRe.test(node.path)
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
        {truncatable && (
          <button
            className="shrink-0 rounded bg-sky-500/15 px-1 text-[10px] text-sky-300 hover:bg-sky-500/25"
            title="适配器截断了长字符串,点击分段取完整值"
            onClick={(e) => { e.stopPropagation(); toggleFull(node) }}
          >
            {node.openFull ? '收起' : '全文'}
          </button>
        )}
        {node.type && <span className="ml-1 shrink-0 text-zinc-600">{node.type}</span>}
      </div>
      {node.openFull && (node.full || node.fullErr) && (
        <div
          className={`mx-1 my-0.5 max-h-56 overflow-auto whitespace-pre-wrap break-all rounded bg-black/40 p-1.5 ${node.fullErr ? 'text-red-400' : 'text-emerald-300'}`}
          style={{ marginLeft: depth * 14 + 8 }}
          onClick={(e) => e.stopPropagation()}
        >
          {node.fullErr || node.full}
        </div>
      )}
      {node.open && node.children?.map((c, j) => (
        <TreeNode key={j} node={c} depth={depth + 1} toggle={toggle} toggleFull={toggleFull} />
      ))}
    </>
  )
}
