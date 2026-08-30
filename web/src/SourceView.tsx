// Monaco 源码视图:4gl 简易语法高亮 + 断点 gutter + 停站行高亮
import { useEffect, useMemo, useRef, useState } from 'react'
import Editor, { loader, type OnMount } from '@monaco-editor/react'
import { Loader2 } from 'lucide-react'
import * as monaco from 'monaco-editor'
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
// codicon 图标映射表:ESM 用法要求宿主显式引入,否则 Ctrl+F 查找控件等只有空按钮
import 'monaco-editor/esm/vs/base/browser/ui/codicons/codiconStyles.js'
import { useStore } from './store'

export const editorRef = { current: null as monaco.editor.IStandaloneCodeEditor | null }

let registered = false
function setupMonaco() {
  if (registered) return
  registered = true
  self.MonacoEnvironment = { getWorker: () => new editorWorker() }
  loader.config({ monaco })
  monaco.languages.register({ id: '4gl' })
  monaco.languages.setMonarchTokensProvider('4gl', {
    keywords: [
      'main', 'function', 'return', 'if', 'then', 'else', 'elif', 'for', 'to', 'step',
      'while', 'case', 'when', 'otherwise', 'define', 'record', 'array', 'dynamic', 'like',
      'type', 'constant', 'let', 'call', 'display', 'input', 'construct', 'by', 'name',
      'on', 'from', 'menu', 'command', 'continue', 'exit', 'next', 'field', 'before',
      'after', 'row', 'key', 'options', 'defer', 'interrupt', 'whenever', 'error',
      'warning', 'open', 'window', 'form', 'close', 'current', 'dialog', 'attributes',
      'unbuffered', 'without', 'defaults', 'accept', 'cancel', 'insert', 'delete',
      'update', 'select', 'foreach', 'execute', 'immediate', 'prepare', 'declare',
      'fetch', 'free', 'rollback', 'work', 'commit', 'run', 'sleep', 'import', 'fgl',
      'public', 'private', 'returns', 'returning', 'null', 'true', 'false', 'not', 'and',
      'or', 'is', 'in', 'goto', 'label', 'end',
    ],
    typeKeywords: ['integer', 'int', 'smallint', 'char', 'varchar', 'string', 'decimal', 'date', 'datetime', 'interval', 'byte', 'text', 'boolean', 'float', 'smallfloat', 'money'],
    tokenizer: {
      root: [
        [/#.*$/, 'comment'],
        [/--.*$/, 'comment'],
        [/"/, { token: 'string', next: '@string' }],
        [/'/, { token: 'string', next: '@sstring' }],
        [/[a-zA-Z_][\w]*/, {
          cases: {
            '@keywords': 'keyword',
            '@typeKeywords': 'type',
            '@default': 'identifier',
          },
        }],
        [/\d+(\.\d+)?/, 'number'],
        [/[()[\],.:=+\-*/|<>]/, 'delimiter'],
      ],
      string: [
        [/[^"]/, 'string'],
        [/"/, { token: 'string', next: '@pop' }],
      ],
      sstring: [
        [/[^']/, 'string'],
        [/'/, { token: 'string', next: '@pop' }],
      ],
    },
  } as any)
  monaco.editor.defineTheme('tdict-dark', {
    base: 'vs-dark',
    inherit: true,
    rules: [
      { token: 'keyword', foreground: '7aa2f7' },
      { token: 'type', foreground: '9ece6a' },
      { token: 'comment', foreground: '6b7280' },
      { token: 'string', foreground: 'e0af68' },
      { token: 'number', foreground: 'ff9e64' },
    ],
    colors: {
      'editor.background': '#101013',
      // 滚动条适配暗色
      'scrollbarSlider.background': '#3f3f4680',
      'scrollbarSlider.hoverBackground': '#52525bb0',
      'scrollbarSlider.activeBackground': '#71717ac0',
    },
  })
}

// 模块加载时立即配置(必须先于 Editor 挂载,否则 loader 可能走 CDN 导致主题/渲染不稳)
setupMonaco()

export function SourceView() {
  const sourceContent = useStore((s) => s.sourceContent)
  const sourceDVM = useStore((s) => s.sourceDVM)
  const currentLine = useStore((s) => s.currentLine)
  const breakpoints = useStore((s) => s.breakpoints)
  const state = useStore((s) => s.state)
  const stop = useStore((s) => s.stop)
  const module = useStore((s) => s.module)
  const [editorReady, setEditorReady] = useState(false)
  const decosRef = useRef<monaco.editor.IEditorDecorationsCollection | null>(null)

  const onMount: OnMount = (editor) => {
    editorRef.current = editor
    decosRef.current = editor.createDecorationsCollection([])
    setEditorReady(true)
    // 点击行号/边栏切换断点
    editor.onMouseDown((e) => {
      const t = e.target.type
      if (t !== monaco.editor.MouseTargetType.GUTTER_GLYPH_MARGIN && t !== monaco.editor.MouseTargetType.GUTTER_LINE_NUMBERS) return
      const line = e.target.position?.lineNumber
      if (line) useStore.getState().toggleBreakpoint(line)
    })
  }

  // 归一化出可比的程序主名:去路径、去 .4gl、去模块前缀。
// fgldb 报的断点文件是 `bsft001_wf.4gl`(或全路径),而 sourceDVM 可能是
// `${模块}_${程序}.4gl`(启动预取)——严格比较会漏画,红点要停站一次才出现
function progKey(f: string, module?: string): string {
  let b = f.split(/[\\/]/).pop() || f
  b = b.replace(/\.4gl$/i, '')
  if (module && b.toLowerCase().startsWith(module.toLowerCase() + '_')) b = b.slice(module.length + 1)
  return b.toLowerCase()
}

// 断点圆点 + 停站行高亮(editorReady 入依赖:编辑器晚于数据就绪时补跑一次)
  useEffect(() => {
    const ed = editorRef.current
    const decos = decosRef.current
    if (!ed || !decos) return
    const list: monaco.editor.IModelDeltaDecoration[] = []
    for (const b of breakpoints) {
      // 断点归属过滤:跨文件跳转后,其它文件的断点行号不能画在当前文件上
      if (sourceDVM && b.file && progKey(b.file, module) !== progKey(sourceDVM, module)) continue
      list.push({
        range: new monaco.Range(b.line, 1, b.line, 1),
        options: {
          isWholeLine: true,
          glyphMarginClassName: 'bp-dot',
          overviewRuler: { color: '#ef4444', position: monaco.editor.OverviewRulerLane.Center },
        },
      })
    }
    if (currentLine > 0 && state === 'stopped') {
      list.push({
        range: new monaco.Range(currentLine, 1, currentLine, 1),
        options: {
          isWholeLine: true,
          className: 'cur-line-hl',
          glyphMarginClassName: 'cur-arrow',
          overviewRuler: { color: '#eab308', position: monaco.editor.OverviewRulerLane.Center },
        },
      })
    }
    decos.set(list)
  }, [breakpoints, currentLine, state, sourceContent, sourceDVM, editorReady, module])

  // 视口跟随:仅在停站行号"值变化"时滚动到当前行。
  // 绝不能放进上面的装饰 effect——它依赖 breakpoints,加断点重跑会把视口拽回运行行
  const prevLineRef = useRef(0)
  useEffect(() => {
    const ed = editorRef.current
    if (!ed) return
    const changed = prevLineRef.current !== currentLine
    prevLineRef.current = currentLine
    if (changed && currentLine > 0 && state === 'stopped') {
      ed.revealLineInCenter(currentLine)
    }
  }, [currentLine, state, sourceDVM, sourceContent, editorReady])

  // 编辑器 options 必须稳定:字面量每次渲染都是新对象,会触发 @monaco-editor/react
  // 反复 updateOptions(minimap 重建),加断点等重渲染时会把滚动位置复位
  const editorOptions = useMemo(
    () => ({
      readOnly: true,
      glyphMargin: true,
      fontSize: 13,
      minimap: { enabled: true, renderCharacters: false, maxColumn: 120 },
      stickyScroll: { enabled: false }, // 关掉滚动时固定在顶部的函数/段头
      scrollBeyondLastLine: false,
      renderLineHighlight: 'none' as const,
      lineNumbersMinChars: 5,
      folding: false,
      automaticLayout: true,
      scrollbar: { verticalScrollbarSize: 10, horizontalScrollbarSize: 10 },
    }),
    [],
  )

  return (
    <div className="relative h-full min-h-0">
      <Editor
        language="4gl"
        theme="tdict-dark"
        value={sourceContent}
        beforeMount={setupMonaco}
        onMount={onMount}
        options={editorOptions}
        loading={
          <div className="flex h-full items-center justify-center">
            <Loader2 className="h-6 w-6 animate-spin text-zinc-500" />
          </div>
        }
      />
      {!sourceContent && (
        <div className="pointer-events-none absolute inset-0 flex items-center justify-center">
          <div className="rounded-sm border border-zinc-800 bg-zinc-900/90 px-4 py-3 text-sm text-zinc-500">
            {state === 'stopped' && stop?.file
              ? <>该模块无源码(仅 42m),当前停站:<span className="text-zinc-300">{stop.file}:{stop.line}</span><br /><span className="text-xs">可继续用变量监视/调用栈分析,或继续运行回到有源码的模块</span></>
              : '等待停站后加载源码…'}
          </div>
        </div>
      )}
    </div>
  )
}
