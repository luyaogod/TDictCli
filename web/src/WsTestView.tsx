// 服务测试视图:复刻 awsq990「集成服务测试」——接口方式/网址/请求报文,直接执行看响应。
// 后端经 SSH 在服务器上以 curl 调用(与 awsq990 同网络位置)。
import { useState } from 'react'
import { FlaskConical, Play } from 'lucide-react'
import { useStore } from './store'

// 接口方式选项(wsfc001 映射,与 awsq990 一致)
const MODES: [string, string, string][] = [
  ['1', 'awsp900', 'Web service (SOAP)'],
  ['2', 'awsp900', 'Web service (SOAP) 备选'],
  ['3', 'awsp920', 'RESTful'],
  ['4', 'awsp940', 'OpenApi restful'],
  ['5', 'awsp930', 'OpenApi Web service'],
]

export function WsTestView() {
  const mode = useStore((s) => s.wsTestMode)
  const url = useStore((s) => s.wsTestUrl)
  const body = useStore((s) => s.wsTestBody)
  const soap = useStore((s) => s.wsTestSoap)
  const result = useStore((s) => s.wsTestResult)
  const running = useStore((s) => s.wsTestRunning)
  const err = useStore((s) => s.wsTestErr)
  const history = useStore((s) => s.wsTestHistory)
  const setWsTest = useStore((s) => s.setWsTest)
  const runWsTest = useStore((s) => s.runWsTest)
  const wsLogSel = useStore((s) => s.wsLogSel)
  const wsLogContent = useStore((s) => s.wsLogContent)
  const [showHist, setShowHist] = useState(false)

  const isSoap = mode === '1' || mode === '2' || mode === '5'
  const ep = MODES.find((m) => m[0] === mode)?.[1] || 'awsp920'
  const defaultUrl = `http://127.0.0.1/wt35prd/ws/r/${ep}`

  const doRun = () => void runWsTest()

  return (
    <div className="flex min-h-0 flex-1 flex-col gap-2 p-2 pt-1">
      {/* 工具条 */}
      <div className="flex shrink-0 flex-wrap items-center gap-2">
        {/* <FlaskConical className="h-4 w-4 text-zinc-500" /> */}
        <select value={mode} onChange={(e) => setWsTest({ mode: e.target.value, url: '' })}
          className="h-7 rounded-md border border-zinc-700 bg-zinc-950 px-1.5 text-xs text-zinc-300 focus:border-zinc-500 focus:outline-none">
          {MODES.map(([v, ep2, label]) => (
            <option key={v} value={v}>{label} ({ep2})</option>
          ))}
        </select>
        <input
          value={url}
          onChange={(e) => setWsTest({ url: e.target.value })}
          placeholder={defaultUrl}
          className="h-7 min-w-0 flex-1 rounded-md border border-zinc-700 bg-zinc-950 px-2 font-mono text-xs text-zinc-200 placeholder:text-zinc-600 focus:border-zinc-500 focus:outline-none"
          title={defaultUrl}
        />
        {isSoap && (
          <label className="flex cursor-pointer items-center gap-1.5 text-xs text-zinc-400" title="SOAP 报文(SOAPAction 空头)">
            <input type="checkbox" checked={soap} onChange={(e) => setWsTest({ soap: e.target.checked })}
              className="h-3.5 w-3.5 accent-sky-500" />
            SOAP
          </label>
        )}
        <button onClick={doRun} disabled={running || !body.trim()}
          title="执行接口调用(服务器侧 curl POST)"
          className="inline-flex items-center gap-1 rounded border border-emerald-800/70 px-2.5 py-1 text-xs text-emerald-400 hover:bg-emerald-950/40 disabled:opacity-40">
          <Play className="h-3.5 w-3.5" fill="currentColor" />
          {running ? '执行中…' : '执行'}
        </button>
        {wsLogSel && (
          <button
            onClick={() => setWsTest({ body: wsLogContent?.request || wsLogSel.reqPath, soap: false })}
            title={`带入日志报文:${wsLogSel.service}`}
            className="rounded border border-zinc-700 px-2 py-1 text-xs text-zinc-400 hover:bg-zinc-800"
          >
            从日志带入({wsLogSel.service.length > 16 ? wsLogSel.service.slice(0, 16) + '…' : wsLogSel.service})
          </button>
        )}
        {history.length > 0 && (
          <button onClick={() => setShowHist(!showHist)}
            className="rounded border border-zinc-700 px-2 py-1 text-xs text-zinc-400 hover:bg-zinc-800">
            历史({history.length})
          </button>
        )}
      </div>

      {err && <div className="shrink-0 rounded border border-red-900/60 bg-red-950/40 px-3 py-1.5 text-xs text-red-300">{err}</div>}

      <div className="flex min-h-0 flex-1 flex-col gap-2 lg:flex-row">
        {/* 请求报文 */}
        <div className="flex min-h-40 flex-1 flex-col overflow-hidden rounded-sm border border-zinc-800 bg-zinc-900/60">
          <div className="flex h-8 shrink-0 items-center border-b border-zinc-800 px-2.5 text-xs font-medium text-zinc-400">
            请求报文(JSON / XML)
          </div>
          <textarea
            value={body}
            onChange={(e) => setWsTest({ body: e.target.value })}
            spellCheck={false}
            placeholder={'{\n  "key": "...",\n  "type": "sync",\n  "host": { "prod": "T100", ... },\n  "service": { "name": "xxx" },\n  "payload": { ... }\n}'}
            className="min-h-0 flex-1 resize-none bg-transparent p-2 font-mono text-xs leading-5 text-zinc-200 placeholder:text-zinc-700 focus:outline-none"
          />
        </div>

        {/* 响应 */}
        <div className="flex min-h-40 flex-1 flex-col overflow-hidden rounded-sm border border-zinc-800 bg-zinc-900/60">
          <div className="flex h-8 shrink-0 items-center gap-3 border-b border-zinc-800 px-2.5 text-xs font-medium text-zinc-400">
            响应
            {result && result.httpCode > 0 && (
              <>
                <span className={`font-mono ${result.httpCode === 200 ? 'text-emerald-500' : 'text-red-500'}`}>
                  HTTP {result.httpCode}
                </span>
                <span className="font-mono text-zinc-500">{result.durationSec.toFixed(3)}s</span>
              </>
            )}
            {result?.error && <span className="text-red-400">{result.error}</span>}
          </div>
          <div className="min-h-0 flex-1 overflow-auto">
            {result?.response ? (
              <pre className="whitespace-pre-wrap break-all p-2 font-mono text-[11px] leading-5 text-zinc-300">{result.response}</pre>
            ) : (
              <div className="p-3 text-xs text-zinc-600">{running ? '请求中…' : '执行后在 此显示响应报文'}</div>
            )}
          </div>
        </div>
      </div>

      {/* 执行历史 */}
      {showHist && history.length > 0 && (
        <div className="max-h-44 shrink-0 overflow-auto rounded-sm border border-zinc-800 bg-zinc-900/60">
          {history.map((h, i) => (
            <div key={i} onClick={() => setWsTest({ result: { httpCode: h.httpCode, durationSec: h.durationSec, response: h.response } })}
              title="点击回看该次响应"
              className="flex h-7 cursor-pointer items-center gap-2 border-b border-zinc-800/60 px-2 text-xs last:border-0 hover:bg-zinc-800/40">
              <span className="w-20 shrink-0 font-mono text-[11px] text-zinc-500">{h.time}</span>
              <span className={`w-14 shrink-0 font-mono ${h.httpCode === 200 ? 'text-emerald-500' : 'text-red-500'}`}>{h.httpCode}</span>
              <span className="w-20 shrink-0 font-mono text-[11px] text-zinc-500">{h.durationSec.toFixed(3)}s</span>
              <span className="min-w-0 flex-1 truncate text-zinc-500" title={h.url}>{h.url}</span>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
