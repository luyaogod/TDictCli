// 源码视图:VS Code 式多页签 = 调试页(锁定第一个,跟随停站)+ 浏览页(Ctrl+点击函数等静态打开)
// 单 Editor 实例,path 切换复用/重建 monaco model(@monaco-editor/react 自动保存恢复视口)
import { useEffect, useMemo, useRef, useState } from 'react'
import Editor, { loader, type OnMount } from '@monaco-editor/react'
import { Bug, Loader2, X } from 'lucide-react'
import * as monaco from 'monaco-editor'
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
// codicon 图标映射表:ESM 用法要求宿主显式引入,否则 Ctrl+F 查找控件等只有空按钮
import 'monaco-editor/esm/vs/base/browser/ui/codicons/codiconStyles.js'
import { useStore } from './store'
import { cn } from './lib/utils'

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
  monaco.editor.defineTheme('tdict-light', {
    base: 'vs',
    inherit: true,
    rules: [
      { token: 'keyword', foreground: '1d4ed8' },
      { token: 'type', foreground: '15803d' },
      { token: 'comment', foreground: '9ca3af' },
      { token: 'string', foreground: 'b45309' },
      { token: 'number', foreground: 'c2410c' },
    ],
    colors: {
      'editor.background': '#ffffff',
      'scrollbarSlider.background': '#d4d4d880',
      'scrollbarSlider.hoverBackground': '#a1a1aab0',
      'scrollbarSlider.activeBackground': '#71717ac0',
    },
  })
}

// 模块加载时立即配置(必须先于 Editor 挂载,否则 loader 可能走 CDN 导致主题/渲染不稳)
setupMonaco()

// 归一化出可比的程序主名:去路径、去 .4gl、去模块前缀。
// fgldb 报的断点文件是 `bsft001_wf.4gl`(或全路径),而 sourceDVM 可能是
// `${模块}_${程序}.4gl`(启动预取)——严格比较会漏画,红点要停站一次才出现
function progKey(f: string, module?: string): string {
  let b = f.split(/[\\/]/).pop() || f
  b = b.replace(/\.4gl$/i, '')
  if (module && b.toLowerCase().startsWith(module.toLowerCase() + '_')) b = b.slice(module.length + 1)
  return b.toLowerCase()
}

export function SourceView() {
  // 调试页数据
  const sourceContent = useStore((s) => s.sourceContent)
  const sourceDVM = useStore((s) => s.sourceDVM)
  const currentLine = useStore((s) => s.currentLine)
  const breakpoints = useStore((s) => s.breakpoints)
  const state = useStore((s) => s.state)
  const stop = useStore((s) => s.stop)
  const module = useStore((s) => s.module)
  const theme = useStore((s) => s.theme)
  const loadingSource = useStore((s) => s.loadingSource)
  const hasContent = useStore((s) => !!s.sourceContent)
  const prog = useStore((s) => s.prog)
  // 页签
  const tabs = useStore((s) => s.tabs)
  const activeTab = useStore((s) => s.activeTab)
  const setActiveTab = useStore((s) => s.setActiveTab)
  const closeTab = useStore((s) => s.closeTab)

  const active = tabs.find((t) => t.key === activeTab)
  const isDebug = !active

  // 当前编辑器展示内容(调试页 vs 浏览页)
  const content = isDebug ? sourceContent : (active!.content || (active!.missing ? '' : ''))
  const modelPath = isDebug ? 'debug:' + (sourceDVM || prog) : 'tab:' + active!.file
  const cursorLine = isDebug ? currentLine : (active!.line ?? 0)

  const [editorReady, setEditorReady] = useState(false)
  const [modelTick, setModelTick] = useState(0) // 页签切换 = model 切换完成后重画装饰/滚动
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
    // 切换页签(@monaco-editor/react 换 model)后装饰集合要重新应用到新 model。
    // 此处不做滚动补偿:视口滚动只由 revealReq 定位信号驱动,纯页签切换保持离开时视口
    editor.onDidChangeModel(() => setModelTick((t) => t + 1))

    // Ctrl+悬停/点击跳函数(类 VS Code,全自管理):
    // 悬停 = 光标下的词变蓝+下划线+手形光标,同时异步预定位(fgldb info line)缓存结果;
    // 点击 = 命中缓存的词才跳转。不用 Monaco definitionProvider——它在按下 Ctrl 的
    // 检测阶段就会被调用,会造成「还没点左键就跳走」
    const fnDecos = editor.createDecorationsCollection([])
    let ctrlDown = false
    let lastWord = ''
    let pendingDef: { word: string; file: string; line: number } | null = null
    const locateCache = new Map<string, { file: string; line: number } | null>()
    const clearHover = () => { lastWord = ''; fnDecos.set([]) }
    const trackKey = (e: KeyboardEvent) => {
      const down = e.ctrlKey
      if (down !== ctrlDown) {
        ctrlDown = down
        if (!down) clearHover()
      }
    }
    window.addEventListener('keydown', trackKey)
    window.addEventListener('keyup', trackKey)
    window.addEventListener('blur', () => { ctrlDown = false; clearHover() })
    const hoverLocate = (word: string) => {
      const st = useStore.getState()
      if (!st.sessionId || st.state !== 'stopped') return
      if (locateCache.has(word)) { pendingDef = locateCache.get(word) ? { word, ...locateCache.get(word)! } : null; return }
      void st.locate(word).then((r) => {
        locateCache.set(word, r.file ? { file: r.file, line: r.line } : null)
        if (lastWord === word) pendingDef = r.file ? { word, file: r.file, line: r.line } : null
      }).catch(() => locateCache.set(word, null))
    }
    editor.onMouseMove((e) => {
      if (!ctrlDown || e.target.position == null) { if (lastWord) clearHover(); return }
      const w = editor.getModel()?.getWordAtPosition(e.target.position)
      if (!w) { if (lastWord) clearHover(); return }
      if (w.word !== lastWord) {
        lastWord = w.word
        pendingDef = null
        const ln = e.target.position.lineNumber
        fnDecos.set([{
          range: new monaco.Range(ln, w.startColumn, ln, w.endColumn),
          options: { inlineClassName: 'fn-link' },
        }])
        hoverLocate(w.word)
      }
    })
    editor.onMouseLeave(() => clearHover())
    editor.onMouseDown((e) => {
      // 内容区 Ctrl+左键:命中悬停预定位的词才跳转(行号/边栏点击不触发)
      if (!ctrlDown || e.target.type !== monaco.editor.MouseTargetType.CONTENT_TEXT || !pendingDef) return
      const w = editor.getModel()?.getWordAtPosition(e.target.position)
      if (w && w.word === pendingDef.word) {
        const d = pendingDef
        pendingDef = null
        void useStore.getState().openSourceTab(d.file, d.line)
      }
    })
  }

  // 装饰:断点圆点按文件归属过滤;停站/定位光标只画在归属文件上
  useEffect(() => {
    const ed = editorRef.current
    const decos = decosRef.current
    if (!ed || !decos) return
    const viewFile = isDebug ? sourceDVM : active!.file
    const list: monaco.editor.IModelDeltaDecoration[] = []
    for (const b of breakpoints) {
      // 断点归属过滤:跨文件跳转后,其它文件的断点行号不能画在当前文件上
      if (viewFile && b.file && progKey(b.file, module) !== progKey(viewFile, module)) continue
      list.push({
        range: new monaco.Range(b.line, 1, b.line, 1),
        options: {
          isWholeLine: true,
          glyphMarginClassName: 'bp-dot',
          overviewRuler: { color: '#ef4444', position: monaco.editor.OverviewRulerLane.Center },
        },
      })
    }
    if (cursorLine > 0 && (isDebug ? state === 'stopped' : true)) {
      list.push({
        range: new monaco.Range(cursorLine, 1, cursorLine, 1),
        options: {
          isWholeLine: true,
          className: isDebug ? 'cur-line-hl' : 'cur-line-hl',
          glyphMarginClassName: isDebug ? 'cur-arrow' : 'cur-arrow',
          overviewRuler: { color: '#eab308', position: monaco.editor.OverviewRulerLane.Center },
        },
      })
    }
    decos.set(list)
  }, [breakpoints, cursorLine, state, content, isDebug, sourceDVM, active?.file, active?.line, editorReady, modelTick, module])

  // 视口跟随:仅停站行号"值变化"时滚动到当前行(调试页)。
  // 绝不能放进装饰 effect——它依赖 breakpoints,加断点重跑会把视口拽回运行行
  // 视口滚动唯一驱动 = revealReq 定位信号(停站落位/步进/跳函数/选帧/断点跳转)。
  // 纯页签切换不发信号:切回页签恢复上次离开的视口,阅读连续性不受光标位置影响。
  // content 入依赖:定位信号先于内容到达时(新开浏览页签),内容就绪后本 effect 重跑完成滚动;
  // 延迟 80ms 等换 model/setValue 完成,避免被 viewState 恢复覆盖
  const revealSeq = useStore((s) => s.revealReq?.seq ?? 0)
  const revealLine = useStore((s) => s.revealReq?.line ?? 0)
  const revealKey = useStore((s) => s.revealReq?.key ?? '')
  useEffect(() => {
    const targetKey = isDebug ? 'debug' : active?.key
    if (!(revealLine > 0) || revealKey !== targetKey) return
    if (isDebug && state !== 'stopped') return
    const t = setTimeout(() => editorRef.current?.revealLineInCenter(revealLine), 80)
    return () => clearTimeout(t)
  }, [revealSeq, revealLine, revealKey, content, isDebug, state, active?.key])

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

  const busy = isDebug
    ? loadingSource && hasContent
    : !!active!.loading

  return (
    <div className="flex h-full min-h-0 flex-col">
      {/* 页签栏:调试页锁定第一个,浏览页可关 */}
      <div className="flex h-8 shrink-0 items-stretch overflow-x-auto border-b border-border bg-background">
        <button onClick={() => setActiveTab('debug')}
          className={cn('flex shrink-0 items-center gap-1.5 border-r border-border px-3 text-xs transition-colors',
            isDebug ? 'bg-accent font-medium text-accent-foreground' : 'text-muted-foreground hover:bg-accent/60')}>
          <Bug className="h-3 w-3" />
          {prog || '调试'}
        </button>
        {tabs.map((t) => (
          <div key={t.key}
            className={cn('group flex shrink-0 items-center border-r border-border transition-colors',
              activeTab === t.key ? 'bg-accent text-accent-foreground' : 'text-muted-foreground hover:bg-accent/60')}>
            <button className="max-w-45 truncate px-3 text-xs" title={t.path || t.file}
              onClick={() => setActiveTab(t.key)}>
              {t.loading ? <Loader2 className="mr-1 inline h-3 w-3 animate-spin" /> : null}
              {t.file}
              {t.missing ? ' (无源码)' : ''}
            </button>
            <button className="mr-1 rounded p-0.5 opacity-40 transition-opacity hover:bg-accent hover:opacity-100"
              onClick={() => closeTab(t.key)}>
              <X className="h-3 w-3" />
            </button>
          </div>
        ))}
      </div>
      {/* 编辑器 */}
      <div className="relative min-h-0 flex-1">
        <Editor
          language="4gl"
          theme={theme === 'light' ? 'tdict-light' : 'tdict-dark'}
          path={modelPath}
          value={content}
          beforeMount={setupMonaco}
          onMount={onMount}
          options={editorOptions}
          loading={
            <div className="flex h-full items-center justify-center">
              <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
            </div>
          }
        />
        {/* 调试页跨文件切换:旧文件保持显示但加遮罩,停站光标等源码到位再落位 */}
        {busy && (
          <div className="pointer-events-none absolute inset-0 flex items-center justify-center bg-background/40">
            <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
          </div>
        )}
        {isDebug && !sourceContent && (
          <div className="pointer-events-none absolute inset-0 flex items-center justify-center">
            <div className="rounded-sm border border-border bg-card/90 px-4 py-3 text-sm text-muted-foreground">
              {state === 'stopped' && stop?.file
                ? <>该模块无源码(仅 42m),当前停站:<span className="text-foreground">{stop.file}:{stop.line}</span><br /><span className="text-xs">可继续用变量监视/调用栈分析,或继续运行回到有源码的模块</span></>
                : '等待停站后加载源码…'}
            </div>
          </div>
        )}
        {!isDebug && active!.missing && (
          <div className="pointer-events-none absolute inset-0 flex items-center justify-center">
            <div className="rounded-sm border border-border bg-card/90 px-4 py-3 text-sm text-muted-foreground">
              该文件无源码(仅 42m 编译产物),无法静态浏览
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
