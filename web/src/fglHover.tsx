// 调试编辑器变量悬浮取值(仅 debug model 生效):
// - Monaco IContentWidget 承载悬浮卡片,锚定在变量词下方,随内容滚动
// - 卡片用 React 渲染:Record/ARRAY 复用侧边栏同款 VarTreeNodes 树(可展开),
//   标量显示原始值;头部右上角复制按钮复制 fgldb print 原始文本
// - 仅停站(stopped)时取值:走 api.print(fgldb print),按表达式缓存;
//   停站行变化(步进/落站)即清缓存保证取值新鲜
import { useState } from 'react'
import { createRoot, type Root } from 'react-dom/client'
import { Copy } from 'lucide-react'
import * as monaco from 'monaco-editor'
import { useStore } from './store'
import { api } from './api'
import { parseFglTree } from './fglparse'
import { VarTreeNodes } from './VarTreeUi'

interface HoverData { expr: string; loading?: boolean; v?: string; e?: string }

// 悬浮卡片:头部(表达式 + 类型 + 右上角复制)+ 树/原始值
function HoverCard({ expr, loading, v, e: err }: HoverData) {
  const [copied, setCopied] = useState(false)
  const kids = v !== undefined ? parseFglTree(v) : null
  const isRecord = !!kids && kids.some((k) => !k.name.startsWith('['))
  const copy = () => {
    if (v === undefined) return
    navigator.clipboard.writeText(v).then(() => {
      setCopied(true)
      setTimeout(() => setCopied(false), 1200)
    })
  }
  return (
    <div className="fgl-hover" onMouseDown={(ev) => ev.stopPropagation()}>
      <div className="fgl-hover-head">
        <span className="fgl-hover-name">{expr}</span>
        {kids && <span className="fgl-hover-type">{isRecord ? 'RECORD' : `ARRAY[${kids.length}]`}</span>}
        <button className="fgl-hover-copy" title="复制原始值" disabled={v === undefined} onClick={copy}>
          {copied ? '已复制' : <Copy className="h-3 w-3" />}
        </button>
      </div>
      {loading ? (
        <div className="fgl-hover-value text-muted-foreground">fgldb print 中…</div>
      ) : err ? (
        <div className="fgl-hover-value text-red-400">{err}</div>
      ) : kids ? (
        <VarTreeNodes nodes={[{ name: expr, open: true, children: kids, type: isRecord ? 'RECORD' : `ARRAY[${kids!.length}]` }]} />
      ) : (
        <div className="fgl-hover-value">{v}</div>
      )}
    </div>
  )
}

// 扩取光标处表达式:支持 record.field / a.b.c 链与 cust.* 结尾
function exprAt(model: monaco.editor.ITextModel, pos: monaco.Position): string | null {
  const w = model.getWordAtPosition(pos)
  if (!w) return null
  const line = model.getLineContent(pos.lineNumber)
  const isW = (ch: string) => /[\w]/.test(ch)
  let s = w.startColumn - 1
  let e = w.endColumn - 1
  while (e < line.length && line[e] === '.') {
    let j = e + 1
    if (line[j] === '*') { e = j + 1; break }
    if (!isW(line[j] ?? '')) break
    while (j < line.length && isW(line[j])) j++
    e = j
  }
  while (s > 0 && line[s - 1] === '.') {
    let i = s - 2
    while (i >= 0 && isW(line[i])) i--
    if (i + 1 >= s - 1) break
    s = i + 1
  }
  const expr = line.slice(s, e)
  return expr || null
}

let attachedTo: monaco.editor.IStandaloneCodeEditor | null = null

// 在 SourceView onMount 里调用(单编辑器实例)
export function attachHover(editor: monaco.editor.IStandaloneCodeEditor) {
  if (attachedTo === editor) return
  attachedTo = editor

  const dom = document.createElement('div')
  dom.style.display = 'none'
  const root: Root = createRoot(dom)
  let wpos: { line: number; column: number } | null = null
  let visible = false
  let current: { expr: string; token: number } | null = null
  let seq = 0
  const cache = new Map<string, { v?: string; e?: string }>()

  const widget: monaco.editor.IContentWidget = {
    getId: () => 'fgl.hover.card',
    getDomNode: () => dom,
    getPosition: () => {
      if (!wpos) return null
      const lineCount = editor.getModel()?.getLineCount() ?? 1
      return {
        position: { lineNumber: Math.min(wpos.line + 1, lineCount), column: wpos.column },
        preference: [monaco.editor.ContentWidgetPositionPreference.EXACT],
      }
    },
  }
  editor.addContentWidget(widget)

  const render = (d: HoverData) => root.render(<HoverCard {...d} />)
  const show = (expr: string, pos: monaco.Position) => {
    wpos = { line: pos.lineNumber, column: pos.column }
    const token = ++seq
    current = { expr, token }
    const cached = cache.get(expr)
    dom.style.display = 'block'
    visible = true
    render({ expr, loading: !cached, v: cached?.v, e: cached?.e })
    editor.layoutContentWidget(widget)
    if (!cached) {
      const st = useStore.getState()
      api.print(st.sessionId!, expr)
        .then(({ value }) => {
          cache.set(expr, { v: value })
          if (current?.token === token) render({ expr, v: value })
        })
        .catch((err: Error) => {
          cache.set(expr, { e: err.message })
          if (current?.token === token) render({ expr, e: err.message })
        })
    }
  }
  const hide = () => {
    if (!visible) return
    visible = false
    current = null
    wpos = null
    dom.style.display = 'none'
    editor.layoutContentWidget(widget)
  }
  // 鼠标在卡片内(含 8px 缓冲)时不隐藏——移向卡片点击复制/展开树途中不被判为换词
  const inCard = (mx: number, my: number) => {
    if (!visible) return false
    const r = dom.getBoundingClientRect()
    return mx >= r.left - 8 && mx <= r.right + 8 && my >= r.top - 8 && my <= r.bottom + 8
  }
  const hideIfOutside = (mx: number, my: number) => {
    if (visible && !inCard(mx, my)) hide()
  }

  editor.onMouseMove((e) => {
    const mx = e.event.posx
    const my = e.event.posy
    if (e.target.type !== monaco.editor.MouseTargetType.CONTENT_TEXT) { hideIfOutside(mx, my); return }
    const model = editor.getModel()
    if (!model || model.uri.scheme !== 'debug') { hide(); return }
    const st = useStore.getState()
    if (!st.sessionId || st.state !== 'stopped') { hide(); return }
    const pos = e.target.position
    if (!pos) { hideIfOutside(mx, my); return }
    const expr = exprAt(model, pos)
    if (!expr) { hideIfOutside(mx, my); return }
    if (visible && current?.expr === expr) return
    if (visible && inCard(mx, my)) return // 移向卡片途中不切换
    hide()
    show(expr, pos)
  })
  editor.onMouseLeave(() => hide())
  // 兜底:鼠标移出编辑器区域(右侧面板/顶栏等)也收卡——不完全依赖 Monaco 的 leave 事件
  const onDocMove = (ev: MouseEvent) => {
    if (!visible) return
    const edom = editor.getDomNode()
    if (!edom) return
    if (edom.contains(ev.target as Node)) return // 编辑器内部交给 editor.onMouseMove
    if (inCard(ev.clientX, ev.clientY)) return // 移向卡片途中
    hide()
  }
  document.addEventListener('mousemove', onDocMove)
  editor.onMouseDown((e) => {
    if (e.target.type !== monaco.editor.MouseTargetType.CONTENT_TEXT) hide()
  })
  editor.onKeyDown((e) => { if (e.keyCode === monaco.KeyCode.Escape) hide() })
  editor.onDidChangeModel(() => hide())
  // 离开停站即收卡;停站行变化(步进/换帧)清缓存,保证下次悬浮取到新值
  let lastStop = ''
  useStore.subscribe((s) => {
    if (s.state !== 'stopped') hide()
    const k = s.stop ? `${s.stop.file}:${s.stop.line}` : ''
    if (k !== lastStop) {
      lastStop = k
      cache.clear()
    }
  })
}
