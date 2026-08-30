import { create } from 'zustand'
import { api, type Event, type StopInfo, type Breakpoint, type Frame, type VarItem, type WSLogItem, type WSLogContent, type WSTestResult } from './api'

// timeline 条目(人/AI/系统 的操作与事件,可审计)
export interface TimelineItem {
  time: string
  origin: 'system' | 'human' | 'ai'
  text: string
  kind: 'info' | 'stop' | 'command' | 'warn'
}

interface Watch { expr: string; value?: string; error?: string }

interface Store {
  // 连接
  wsConnected: boolean
  // 会话
  sessionId: string | null
  module: string
  prog: string
  state: string // loading|stopped|running|exit|''
  started: boolean // 程序是否已 run 过(入口停站时步进不可用)
  stop: StopInfo | null
  holdingSeconds: number
  launching: boolean
  // UI 面板显隐
  showRight: boolean
  showBottom: boolean
  // 数据
  breakpoints: Breakpoint[]
  adjustedBps: Record<number, number> // 点击行号 → 实际注册断点编号(fgldb 会把非可执行行的断点自动下移)
  frames: Frame[]
  watches: Watch[]
  autovars: VarItem[] // 停站自动变量(当前源码窗内变量的自动求值)
  selectedFrame: number // 当前选中栈帧(-1 = 未选)
  backendDead: string // 后端死亡/程序退出原因(空 = 正常)
  timeline: TimelineItem[]
  rawLog: string[]
  runProg: string // gzzz_t 解析出的实体程序(源码命名/预取用);空 = 与 prog 相同
  // 视图与接口日志(VS Code 活动栏切换)
  view: 'debug' | 'wslogs' | 'wstest'
  // 服务测试(复刻 awsq990 集成服务测试)
  wsTestMode: string // 1/2 awsp900, 3 awsp920, 4 awsp940, 5 awsp930
  wsTestUrl: string
  wsTestBody: string
  wsTestSoap: boolean
  wsTestResult: WSTestResult | null
  wsTestRunning: boolean
  wsTestErr: string
  wsTestHistory: { time: string; url: string; httpCode: number; durationSec: number; response: string }[]
  wsLogs: WSLogItem[]
  wsLogsLoading: boolean
  wsLogsPage: number
  wsLogsHasMore: boolean
  wsLogSel: WSLogItem | null
  wsLogContent: WSLogContent | null
  wsLogTab: 'info' | 'request' | 'response'
  wsLogErr: string
  // 源码
  sourceContent: string
  sourcePath: string
  sourceDVM: string // 当前已加载源码对应的 DVM 模块文件名(无源码模块也记录,避免错文件高亮)
  currentLine: number
  loadingSource: boolean

  // actions
  setWsConnected: (b: boolean) => void
  toggleRight: () => void
  toggleBottom: () => void
  pushRaw: (line: string) => void
  sendRaw: (cmd: string) => Promise<void>
  pushTimeline: (item: Omit<TimelineItem, 'time'>) => void
  onEvent: (ev: Event) => void
  setView: (v: 'debug' | 'wslogs' | 'wstest') => void
  setWsTest: (p: { mode?: string; url?: string; body?: string; soap?: boolean; result?: WSTestResult | null }) => void
  runWsTest: () => Promise<void>
  loadWsLogs: (service: string, onlyFail: boolean, page?: number, startFrom?: string, startTo?: string) => Promise<void>
  selectWsLog: (item: WSLogItem) => Promise<void>
  setWsLogTab: (t: 'info' | 'request' | 'response') => void
  replayDebug: (item: WSLogItem) => Promise<void>
  launch: (module: string, prog: string) => Promise<void>
  refreshSnapshot: () => Promise<void>
  refreshSource: (file?: string) => Promise<void>
  refreshFrames: () => Promise<Frame[]>
  refreshWatches: () => Promise<void>
  addWatch: (expr: string) => Promise<void>
  removeWatch: (expr: string) => void
  selectFrame: (idx: number) => Promise<void>
  toggleBPEnabled: (num: number, enabled: boolean) => Promise<void>
  jumpToBp: (b: Breakpoint) => Promise<void>
  runToCursor: (line: number) => Promise<void>
  control: (action: string, arg?: string) => Promise<void>
  toggleBreakpoint: (line: number) => Promise<void>
  removeBreakpoint: (num: number) => Promise<void>
  quit: () => Promise<void>
  restart: () => Promise<void>
  doPrint: (expr: string) => Promise<void>
  adoptExisting: () => Promise<void>
}

let snapTimer: number | undefined
let holdTimer: number | undefined

function now() { return new Date().toLocaleTimeString('zh-CN', { hour12: false }) }

// 跳到 MAIN 语句行(无停站位置时给用户一个可点断点的起点)
function jumpToMain(set: (p: Partial<Store>) => void, get: () => Store) {
  if (get().currentLine > 0) return
  const lines = get().sourceContent.split('\n')
  for (let i = 0; i < lines.length; i++) {
    if (/^\s*MAIN\b/i.test(lines[i])) { set({ currentLine: i + 1 }); break }
  }
}

export const useStore = create<Store>((set, get) => ({
  wsConnected: false,
  sessionId: null, module: '', prog: '', state: '', started: false, stop: null, holdingSeconds: 0,
  launching: false, showRight: true, showBottom: true,
  breakpoints: [], adjustedBps: {}, frames: [], watches: [], autovars: [], selectedFrame: -1, backendDead: '',
  timeline: [], rawLog: [],
  runProg: '',
  view: 'debug',
  wsLogs: [], wsLogsLoading: false, wsLogsPage: 1, wsLogsHasMore: false,
  wsLogSel: null, wsLogContent: null, wsLogTab: 'info', wsLogErr: '',
  wsTestMode: '3', wsTestUrl: '', wsTestBody: '', wsTestSoap: false,
  wsTestResult: null, wsTestRunning: false, wsTestErr: '', wsTestHistory: [],
  sourceContent: '', sourcePath: '', sourceDVM: '', currentLine: 0, loadingSource: false,

  setWsConnected: (b) => set({ wsConnected: b }),
  toggleRight: () => set((st) => ({ showRight: !st.showRight })),
  toggleBottom: () => set((st) => ({ showBottom: !st.showBottom })),

  pushRaw: (line) => set((st) => {
    const log = st.rawLog.length > 3000 ? st.rawLog.slice(-2000) : st.rawLog
    return { rawLog: [...log, line] }
  }),

  // 直接执行 fgldb 命令(复刻原版 fgldeb Ctrl+D「Input Debugger Command」的透传):
  // 命令与输出原样进协议流;状态类命令执行后刷新快照/栈,与原版行为一致
  sendRaw: async (cmd) => {
    const { sessionId } = get()
    if (!sessionId || !cmd.trim()) return
    get().pushRaw(`> ${cmd}`)
    try {
      const r = await api.raw(sessionId, cmd)
      for (const l of r.lines) get().pushRaw(l)
      const head = cmd.trim().split(/\s+/)[0].toLowerCase()
      if (['step', 'next', 'continue', 'run'].includes(head)) {
        // 程序会继续运行:走轮询等下一次停站
        pollUntilStopped(set, get)
      } else if (['break', 'tbreak', 'clear', 'delete', 'enable', 'disable', 'where', 'frame', 'finish', 'return'].includes(head)) {
        void get().refreshSnapshot()
        if (get().state === 'stopped') void get().refreshFrames()
      }
    } catch (e: any) {
      get().pushRaw(`[错误] ${e.message || String(e)}`)
    }
  },

  pushTimeline: (item) => set((st) => ({
    timeline: [...st.timeline.slice(-500), { ...item, time: now() }],
  })),

  onEvent: (ev) => {
    const st = get()
    if (st.sessionId && ev.sessionId && ev.sessionId !== st.sessionId) return
    switch (ev.type) {
      case 'output':
        st.pushRaw(ev.text || '')
        return
      case 'state':
        set({ state: ev.state || '' })
        if (ev.state === 'stopped') {
          // 停站后自动刷新栈与监视值(异步,不阻塞)
          void get().refreshFrames()
          void get().refreshWatches()
          startHoldTimer(set, get)
        }
        if (ev.state === 'running' || ev.state === 'exit') {
          if (snapTimer) { clearInterval(snapTimer); snapTimer = undefined }
          stopHoldTimer()
          if (ev.state === 'running') set({ selectedFrame: -1 })
        }
        if (ev.state === 'exit') {
          st.pushTimeline({ origin: 'system', kind: 'warn', text: '会话结束' })
          if (snapTimer) { clearInterval(snapTimer); snapTimer = undefined }
          stopHoldTimer()
        }
        return
      case 'dead':
        // 后端死亡/SSH 断开:forceExit 随后会发 state=exit
        set({ backendDead: ev.text || '调试后端连接断开' })
        st.pushTimeline({ origin: 'system', kind: 'warn', text: ev.text || '调试后端连接断开' })
        return
      case 'autovars':
        set({ autovars: ev.vars || [] })
        return
      case 'stopped':
        set({ stop: ev.stop || null, state: 'stopped', currentLine: ev.stop?.line || get().currentLine, selectedFrame: -1 })
        st.pushTimeline({
          origin: 'system', kind: 'stop',
          text: `停站[${ev.stop?.reason}] ${ev.stop?.file || ''}:${ev.stop?.line ?? ''} ${ev.stop?.func || ''}`,
        })
        // 停站文件缺失(如 SIGINT 中断块无文件头)时,用栈顶帧的真实文件跟随
        void (async () => {
          let file = ev.stop?.file
          if (!file) {
            const frames = await get().refreshFrames()
            file = frames[0]?.file
            if (file) set({ currentLine: frames[0]?.line ?? 0 })
          }
          await get().refreshSource(file)
          await get().refreshWatches()
          await get().refreshSnapshot()
        })()
        startHoldTimer(set, get)
        return
      case 'watchdog':
        st.pushTimeline({ origin: 'system', kind: 'warn', text: ev.text || '看门狗触发' })
        return
      case 'ai_action':
        st.pushTimeline({ origin: 'ai', kind: 'command', text: ev.text || '' })
        return
      case 'log':
        st.pushTimeline({ origin: 'system', kind: 'info', text: ev.text || '' })
        return
    }
  },

  setView: (v) => set({ view: v }),

  setWsTest: (p) => {
    const patch: Partial<Store> = {}
    if (p.mode !== undefined) patch.wsTestMode = p.mode
    if (p.url !== undefined) patch.wsTestUrl = p.url
    if (p.body !== undefined) patch.wsTestBody = p.body
    if (p.soap !== undefined) patch.wsTestSoap = p.soap
    if (p.result !== undefined) patch.wsTestResult = p.result
    set(patch)
  },

  runWsTest: async () => {
    const st = get()
    set({ wsTestRunning: true, wsTestErr: '' })
    try {
      const { result } = await api.wsTest(st.wsTestMode, st.wsTestUrl, st.wsTestBody, st.wsTestSoap)
      set({ wsTestResult: result })
      if (result.httpCode > 0) {
        const entry = {
          time: new Date().toLocaleTimeString('zh-CN', { hour12: false }),
          url: st.wsTestUrl || '(默认地址)',
          httpCode: result.httpCode, durationSec: result.durationSec, response: result.response,
        }
        set((s) => ({ wsTestHistory: [entry, ...s.wsTestHistory].slice(0, 10) }))
      }
    } catch (e: any) {
      set({ wsTestErr: e.message || String(e) })
    } finally {
      set({ wsTestRunning: false })
    }
  },

  loadWsLogs: async (service, onlyFail, page = 1, startFrom = '', startTo = '') => {
    set({ wsLogsLoading: true, wsLogErr: '' })
    try {
      const r = await api.wsLogs(service.trim(), onlyFail, page, startFrom, startTo)
      set({ wsLogs: r.items || [], wsLogsHasMore: !!r.hasMore, wsLogsPage: page, wsLogSel: null, wsLogContent: null })
    } catch (e: any) {
      set({ wsLogErr: e.message || String(e), wsLogs: [] })
    } finally {
      set({ wsLogsLoading: false })
    }
  },

  selectWsLog: async (item) => {
    set({ wsLogSel: item, wsLogContent: null, wsLogErr: '' })
    try {
      const { content } = await api.wsLogContent(item.rowid)
      set({ wsLogContent: content })
    } catch (e: any) {
      set({ wsLogErr: e.message || String(e) })
    }
  },

  setWsLogTab: (t) => set({ wsLogTab: t }),

  replayDebug: async (item) => {
    set({ wsLogErr: '' })
    try {
      const r = await api.wsLogDebug(item.rowid)
      const mod = r.module || ''
      const rp = r.runProg || r.prog || item.job
      set({
        view: 'debug', sessionId: r.sessionId, module: mod, prog: r.prog || item.job,
        runProg: rp, state: 'loading', timeline: [], rawLog: [], watches: [], autovars: [],
        backendDead: '', selectedFrame: -1, stop: null, breakpoints: [], frames: [],
        sourceContent: '', sourcePath: '', sourceDVM: '', currentLine: 0,
      })
      // 入口停站前预取源码(实体程序命名)
      if (mod && rp) {
        void (async () => {
          try {
            const pr = await api.sourcePreview(mod, rp)
            const cur = get()
            if (cur.sourceContent || cur.sourceDVM) return
            set({ sourceContent: pr.source.content, sourcePath: pr.source.path, sourceDVM: `${mod}_${rp}.4gl` })
            jumpToMain(set, get)
          } catch { /* 停站后由会话路径加载 */ }
        })()
      }
      pollUntilStopped(set, get)
    } catch (e: any) {
      set({ wsLogErr: e.message || String(e) })
    }
  },

  launch: async (module, prog) => {
    set({ launching: true, timeline: [], rawLog: [], watches: [], autovars: [], backendDead: '', selectedFrame: -1 })
    const dvm = `${module}_${prog}.4gl`
    // 更换作业时清空旧代码;同作业重启则保留(用户可能正在翻看)
    if (get().sourceDVM !== dvm) {
      set({ sourceContent: '', sourcePath: '', sourceDVM: '', currentLine: 0 })
    } else {
      set({ currentLine: 0 })
    }
    try {
      const r = await api.launch(module, prog)
      // 作业编号解析:后端连 gzzz_t 后回读模块与实体程序(aint301_wf → aint302_wf)
      const mod = r.module || module
      const rp = r.runProg || prog
      set({ sessionId: r.sessionId, module: mod, prog, runProg: rp, state: 'loading' })
      get().pushTimeline({ origin: 'human', kind: 'command', text: `启动调试会话 ${mod}/${prog}` })
      // 入口停站前预取源码(用解析出的实体程序命名,消除空白)
      if (!get().sourceContent) {
        void (async () => {
          try {
            const pr = await api.sourcePreview(mod, rp)
            const cur = get()
            if (cur.sourceContent || cur.sourceDVM) return // 会话路径已先行加载
            set({ sourceContent: pr.source.content, sourcePath: pr.source.path, sourceDVM: `${mod}_${rp}.4gl` })
            jumpToMain(set, get)
            cur.pushTimeline({ origin: 'system', kind: 'info', text: '已预取源码(会话建立中)' })
          } catch { /* 预取失败不致命,停站后由会话路径加载 */ }
        })()
      }
      // 轮询直到入口停站
      pollUntilStopped(set, get)
    } catch (e: any) {
      get().pushTimeline({ origin: 'system', kind: 'warn', text: `启动失败: ${e.message}` })
      throw e
    } finally {
      set({ launching: false })
    }
  },

  refreshSnapshot: async () => {
    const { sessionId } = get()
    if (!sessionId) return
    try {
      const snap = await api.snapshot(sessionId)
      // 注意:这里不更新 currentLine——下断点等操作也会触发快照刷新,
      // 焦点行只能由真正的停站事件/步进响应驱动,否则浏览位置会被拽回运行行
      set({
        state: snap.state, stop: snap.stop, breakpoints: snap.breakpoints || [],
        started: !!snap.started,
        holdingSeconds: snap.holdingSeconds || 0,
        // 免模块启动时后端会按作业名解析模块,回读给前端(源码兜底路径依赖它)
        module: snap.module || get().module,
        runProg: snap.runProg || get().runProg,
      })
      if (snap.state === 'stopped' && !snapTimer) startHoldTimer(set, get)
      if (snap.state !== 'stopped' && snap.state !== 'loading' && snapTimer) {
        clearInterval(snapTimer); snapTimer = undefined
      }
    } catch { /* 会话可能已结束 */ }
  },

  refreshSource: async (file) => {
    const { sessionId, module, prog, runProg, sourceDVM } = get()
    if (!sessionId) return
    let f = file || get().stop?.file
    let entryMode = false
    if (!f) {
      // 入口停站:DVM 未上报文件名,按 T100 命名规则加载母版源码,便于直接点行号下断点
      // 注意:作业编号可能解析出不同名的实体程序(gzzz_t),母版跟实体程序走
      const rp = runProg || prog
      if (!rp) return
      f = `${module ? module + '_' : ''}${rp}.4gl`
      entryMode = true
    }
    if (f === sourceDVM) return // 已加载(或已确认无源码),避免重复拉取
    set({ loadingSource: true })
    try {
      const { source } = await api.sourceByFile(sessionId, f, module)
      set({ sourceContent: source.content, sourcePath: source.path, sourceDVM: f })
      const st = get()
      if (entryMode && st.currentLine === 0) {
        jumpToMain(set, get)
        st.pushTimeline({ origin: 'system', kind: 'info', text: '入口停站:已显示源码,点击行号下断点后点「继续 F5」开始' })
      }
    } catch {
      // 该模块无源码(如 com 公共库只有 42m):明确置空,避免在旧文件上标错停站行
      set({ sourceContent: '', sourcePath: '', sourceDVM: f })
    } finally {
      set({ loadingSource: false })
    }
  },

  refreshFrames: async () => {
    const { sessionId, state } = get()
    if (!sessionId || state !== 'stopped') return []
    try {
      const { frames } = await api.where(sessionId)
      set({ frames })
      return frames
    } catch { set({ frames: [] }); return [] }
  },

  refreshWatches: async () => {
    const { sessionId, watches, state } = get()
    if (!sessionId || state !== 'stopped') return
    const out: Watch[] = []
    for (const w of watches) {
      try {
        const { value } = await api.print(sessionId, w.expr)
        out.push({ expr: w.expr, value })
      } catch (e: any) {
        out.push({ expr: w.expr, error: e.message || String(e) })
      }
    }
    set({ watches: out })
  },

  addWatch: async (expr) => {
    const st = get()
    if (!expr.trim() || st.watches.some((w) => w.expr === expr)) return
    set({ watches: [...st.watches, { expr }] })
    await st.refreshWatches()
    st.pushTimeline({ origin: 'human', kind: 'command', text: `监视 ${expr}` })
  },

  removeWatch: (expr) => set((st) => ({ watches: st.watches.filter((w) => w.expr !== expr) })),

  control: async (action, arg) => {
    const { sessionId } = get()
    if (!sessionId) return
    get().pushTimeline({ origin: 'human', kind: 'command', text: arg ? `${action} ${arg}` : action })
    try {
      const resp = await api.control(sessionId, action, arg)
      if (action !== 'interrupt') {
        await get().refreshSnapshot()
        const st = get()
        // 步进可能跨模块:停站文件与当前显示源码不同时,跟随切换
        if (st.stop?.file && st.stop.file !== st.sourceDVM) void st.refreshSource(st.stop.file)
        // 停站后同步调用栈与监视取值
        if (st.state === 'stopped') {
          void st.refreshFrames()
          void st.refreshWatches()
        }
      }
      if (action === 'run' || action === 'continue') pollUntilStopped(set, get)
      if (resp?.stop && (action === 'next' || action === 'step' || action === 'finish' || action === 'until')) {
        set({ currentLine: resp.stop.line ?? 0, selectedFrame: -1 })
      }
    } catch (e: any) {
      get().pushTimeline({ origin: 'system', kind: 'warn', text: `${action} 失败: ${e.message}` })
    }
  },

  toggleBreakpoint: async (line) => {
    const st = get()
    if (!st.sessionId) return
    // 当前视图仅显示一个文件的源码:按行号匹配;另查"点击行→注册断点"映射
    // (fgldb 会把空行/注释行上的断点自动调整到下一条可执行语句,点原行也要能取消)
    const existing =
      st.breakpoints.find((b) => b.line === line) ||
      (st.adjustedBps[line] !== undefined
        ? st.breakpoints.find((b) => b.num === st.adjustedBps[line])
        : undefined)
    try {
      let newBp: Breakpoint | undefined
      if (existing) {
        await api.bpDel(st.sessionId, existing.num)
        st.pushTimeline({ origin: 'human', kind: 'command', text: `删除断点 ${existing.file}:${existing.line}` })
      } else {
        const loc = st.stop?.file ? `${st.stop.file}:${line}` : String(line)
        const { breakpoint } = await api.bpAdd(st.sessionId, loc)
        newBp = breakpoint
        st.pushTimeline({ origin: 'human', kind: 'command', text: `断点 ${breakpoint.file}:${breakpoint.line}` })
      }
      await st.refreshSnapshot()
      // 维护"点击行→注册断点"映射:实际注册行号与点击行不同时记录;失效映射清理
      const next: Record<number, number> = {}
      for (const [clickLine, num] of Object.entries(st.adjustedBps)) {
        if (get().breakpoints.some((b) => b.num === num)) next[Number(clickLine)] = num
      }
      if (newBp && newBp.line !== line) next[line] = newBp.num
      set({ adjustedBps: next })
    } catch (e: any) {
      st.pushTimeline({ origin: 'system', kind: 'warn', text: `断点操作失败: ${e.message}` })
    }
  },

  removeBreakpoint: async (num) => {
    const st = get()
    if (!st.sessionId) return
    try {
      await api.bpDel(st.sessionId, num)
      await st.refreshSnapshot()
    } catch (e: any) {
      st.pushTimeline({ origin: 'system', kind: 'warn', text: `删除断点失败: ${e.message}` })
    }
  },

  // 选择栈帧:切换 print/locals 求值上下文并跳转该帧源码位置(仅停站时可用)
  selectFrame: async (idx) => {
    const st = get()
    if (!st.sessionId || st.state !== 'stopped') return
    const frame = st.frames.find((f) => f.idx === idx)
    try {
      await api.frame(st.sessionId, idx)
      set({ selectedFrame: idx })
      if (frame) {
        st.pushTimeline({ origin: 'human', kind: 'command', text: `选帧 #${idx} ${frame.func} ${frame.file}:${frame.line}` })
        if (frame.file && frame.file !== st.sourceDVM) await st.refreshSource(frame.file)
        if (frame.line) set({ currentLine: frame.line })
      }
      void st.refreshWatches()
    } catch (e: any) {
      st.pushTimeline({ origin: 'system', kind: 'warn', text: `选帧失败: ${e.message}` })
    }
  },

  toggleBPEnabled: async (num, enabled) => {
    const st = get()
    if (!st.sessionId) return
    try {
      await api.bpEnabled(st.sessionId, num, enabled)
      st.pushTimeline({ origin: 'human', kind: 'command', text: `${enabled ? '启用' : '禁用'}断点 #${num}` })
      await st.refreshSnapshot()
    } catch (e: any) {
      st.pushTimeline({ origin: 'system', kind: 'warn', text: `断点启停失败: ${e.message}` })
    }
  },

  // 断点列表点击跳转:必要时切换到断点所在文件,再定位到断点行(不改变运行上下文)
  jumpToBp: async (b) => {
    const st = get()
    if (b.file && b.file !== st.sourceDVM) await st.refreshSource(b.file)
    set({ currentLine: b.line })
  },

  // 运行到光标:当前文件即停站文件时用行号,否则带文件名(fgldb until [file:]line)
  runToCursor: async (line) => {
    const st = get()
    if (!st.sessionId || st.state !== 'stopped') return
    const cur = st.sourceDVM
    const top = st.stop?.file || st.frames[0]?.file || ''
    const arg = cur && top && cur !== top ? `${cur}:${line}` : String(line)
    await st.control('until', arg)
  },

  quit: async () => {
    const st = get()
    if (!st.sessionId) return
    try {
      await api.quit(st.sessionId)
      st.pushTimeline({ origin: 'human', kind: 'command', text: '结束会话' })
    } catch { /* ignore */ }
    if (snapTimer) { clearInterval(snapTimer); snapTimer = undefined }
    stopHoldTimer()
    // 保留源码:会话结束后用户可能仍想翻看代码(重启/换作业时才重载)
    set({ sessionId: null, state: '', started: false, stop: null, breakpoints: [], frames: [], adjustedBps: {}, currentLine: 0, autovars: [], selectedFrame: -1, backendDead: '' })
  },

  restart: async () => {
    const st = get()
    if (!st.sessionId) return
    const { module, prog } = st
    st.pushTimeline({ origin: 'human', kind: 'command', text: `重新开始 ${module}/${prog}` })
    try { await api.quit(st.sessionId) } catch { /* 忽略,直接重启 */ }
    if (snapTimer) { clearInterval(snapTimer); snapTimer = undefined }
    stopHoldTimer()
    set({ sessionId: null, state: '', started: false, stop: null, breakpoints: [], frames: [], adjustedBps: {}, autovars: [], selectedFrame: -1, backendDead: '' })
    await get().launch(module, prog)
  },

  doPrint: async (expr) => {
    const st = get()
    if (!st.sessionId || !expr.trim()) return
    try {
      const { value } = await api.print(st.sessionId, expr)
      st.pushTimeline({ origin: 'human', kind: 'info', text: `print ${expr} → ${value}` })
      await st.addWatch(expr)
    } catch (e: any) {
      st.pushTimeline({ origin: 'system', kind: 'warn', text: `print ${expr} → ${e.message}` })
    }
  },

  adoptExisting: async () => {
    try {
      const { sessions } = await api.list()
      if (!sessions.length) return
      const s = sessions[0]
      set({ sessionId: s.id, module: s.module, prog: s.prog, runProg: s.runProg || s.prog, state: s.state })
      get().pushTimeline({ origin: 'system', kind: 'info', text: `接管已存在的会话 ${s.module}/${s.prog}` })
      const snap = await api.snapshot(s.id)
      set({
        stop: snap.stop, breakpoints: snap.breakpoints || [],
        started: !!snap.started,
        currentLine: snap.stop?.line ?? 0, holdingSeconds: snap.holdingSeconds || 0,
      })
      if (snap.state === 'stopped') await get().refreshSource(snap.stop?.file)
      void get().refreshFrames()
      void get().refreshWatches()
    } catch { /* ignore */ }
  },
}))

function pollUntilStopped(set: (p: Partial<Store>) => void, get: () => Store) {
  if (snapTimer) clearInterval(snapTimer)
  snapTimer = window.setInterval(async () => {
    const st = get()
    if (!st.sessionId) { if (snapTimer) clearInterval(snapTimer); return }
    try {
      const snap = await api.snapshot(st.sessionId)
      set({
        state: snap.state, stop: snap.stop, breakpoints: snap.breakpoints || [],
        started: !!snap.started,
        holdingSeconds: snap.holdingSeconds || 0,
        module: snap.module || get().module,
        runProg: snap.runProg || get().runProg,
        currentLine: snap.stop?.line || get().currentLine,
      })
      // 源码尽早显示:不等停站,会话就绪(模块已解析)就按 entryMode 拉取,断点也随之可见
      if (!get().sourceContent && snap.module && snap.state !== 'exit') {
        void get().refreshSource()
      }
      if (snap.state === 'stopped') {
        clearInterval(snapTimer!)
        snapTimer = undefined
        startHoldTimer(set, get)
        void get().refreshFrames()
        void get().refreshWatches()
        void get().refreshSource(snap.stop?.file)
        // 兜底:停站宣告与后端断点恢复收尾之间仍有微小窗口(事件先于快照到达),
        // 稍后补一次快照,保证缓存的断点无需步进就能显示
        window.setTimeout(() => {
          const cur = get()
          if (cur.sessionId && cur.state === 'stopped') void cur.refreshSnapshot()
        }, 2500)
        if (snap.stop?.reason === 'breakpoint') {
          st.pushTimeline({ origin: 'system', kind: 'stop', text: `命中断点 ${snap.stop.file}:${snap.stop.line}` })
        }
      }
      if (snap.state === 'exit') { clearInterval(snapTimer!); snapTimer = undefined }
    } catch { /* ignore */ }
  }, 1000)
}

function startHoldTimer(set: (p: Partial<Store>) => void, get: () => Store) {
  stopHoldTimer()
  holdTimer = window.setInterval(() => {
    const st = get()
    if (st.state !== 'stopped') { stopHoldTimer(); return }
    set({ holdingSeconds: st.holdingSeconds + 1 })
  }, 1000)
}

function stopHoldTimer() {
  if (holdTimer) { clearInterval(holdTimer); holdTimer = undefined }
}

// 连接 WebSocket(自动重连)
export function connectWS() {
  const setWs = useStore.getState().setWsConnected
  const onEvent = useStore.getState().onEvent
  const pushRaw = useStore.getState().pushRaw
  let closed = false
  let ws: WebSocket | undefined

  function connect() {
    const proto = location.protocol === 'https:' ? 'wss' : 'ws'
    ws = new WebSocket(`${proto}://${location.host}/api/ws`)
    ws.onopen = () => setWs(true)
    ws.onclose = () => {
      setWs(false)
      if (!closed) setTimeout(connect, 2000)
    }
    ws.onmessage = (m) => {
      try {
        const ev: Event = JSON.parse(m.data)
        onEvent(ev)
      } catch { pushRaw(String(m.data)) }
    }
  }
  connect()
  return () => { closed = true; ws?.close() }
}

// 调试句柄:浏览器控制台可用 __store.getState() 检查状态
if (typeof window !== 'undefined') {
  ;(window as unknown as Record<string, unknown>).__store = useStore
}
