// REST API 客户端
export interface Frame { idx: number; func: string; file: string; line: number }
export interface SourceLine { num: number; text: string; isCur: boolean }
export interface StopInfo {
  reason: string; bpNum?: number; func?: string; file?: string; line?: number
  frames?: Frame[]; source?: SourceLine[]
}
export interface Breakpoint { num: number; file: string; line: number; func?: string; enabled: boolean; note?: string }
export interface VarItem { expr: string; value?: string }
export interface VarDecl { name: string; type: string }
export interface SessionBrief {
  id: string; module: string; prog: string; runProg?: string; state: string
  file?: string; line?: number; func?: string; reason?: string
  holdingSeconds: number; breakpoints: number; watchdogSeconds: number
}
export interface WSLogItem {
  rowid: string; service: string; pid: string
  start: string; end: string; duration: string
  code: string; job: string
  reqPath: string; rspPath: string; reqSize: string; rspSize: string; errMsg: string
}
export interface WSLogContent {
  request: string; response: string
}

export interface WSTestResult {
  httpCode: number; durationSec: number; response: string; error?: string
}

export interface Event {
  type: string; sessionId: string; time: string
  state?: string; stop?: StopInfo; vars?: VarItem[]; text?: string
}

async function req<T>(url: string, opts?: RequestInit): Promise<T> {
  const r = await fetch(url, {
    headers: { 'Content-Type': 'application/json' },
    ...opts,
  })
  const data = await r.json().catch(() => ({}))
  if (!r.ok || data.ok === false) throw new Error(data.error || `HTTP ${r.status}`)
  return data as T
}

export const api = {
  status: () => req<any>('/api/status'),
  list: () => req<{ sessions: SessionBrief[] }>('/api/sessions'),
  launch: (module: string, prog: string) =>
    req<{ sessionId: string; module?: string; prog?: string; runProg?: string }>('/api/sessions', { method: 'POST', body: JSON.stringify({ module, prog }) }),
  snapshot: (id: string) => req<any>(`/api/sessions/${id}`),
  quit: (id: string) => req<any>(`/api/sessions/${id}`, { method: 'DELETE' }),
  bpAdd: (id: string, location: string) =>
    req<{ breakpoint: Breakpoint }>(`/api/sessions/${id}/breakpoints`, { method: 'POST', body: JSON.stringify({ location }) }),
  bpDel: (id: string, num: number) => req<any>(`/api/sessions/${id}/breakpoints/${num}`, { method: 'DELETE' }),
  control: (id: string, action: string, arg?: string) =>
    req<any>(`/api/sessions/${id}/control`, { method: 'POST', body: JSON.stringify({ action, arg }) }),
  print: (id: string, expr: string) => req<{ value: string }>(`/api/sessions/${id}/print`, { method: 'POST', body: JSON.stringify({ expr }) }),
  where: (id: string) => req<{ frames: Frame[] }>(`/api/sessions/${id}/where`, { method: 'POST' }),
  raw: (id: string, command: string) => req<{ lines: string[] }>(`/api/sessions/${id}/raw`, { method: 'POST', body: JSON.stringify({ command }) }),
  locals: (id: string) => req<{ vars: VarItem[] }>(`/api/sessions/${id}/locals`),
  globals: (id: string, limit?: number) => req<{ vars: VarDecl[]; total: number }>(`/api/sessions/${id}/globals${limit ? `?limit=${limit}` : ''}`),
  sources: (id: string) => req<{ sources: string[] }>(`/api/sessions/${id}/sources`),
  functions: (id: string, limit?: number) => req<{ functions: string[]; total: number }>(`/api/sessions/${id}/functions${limit ? `?limit=${limit}` : ''}`),
  autovars: (id: string) => req<{ vars: VarItem[] }>(`/api/sessions/${id}/autovars`),
  frame: (id: string, num: number) => req<{ frame: number }>(`/api/sessions/${id}/frame`, { method: 'POST', body: JSON.stringify({ num }) }),
  bpEnabled: (id: string, num: number, enabled: boolean) =>
    req<any>(`/api/sessions/${id}/breakpoints/${num}/enabled`, { method: 'POST', body: JSON.stringify({ enabled }) }),
  sourceByFile: (id: string, file: string, module: string) =>
    req<{ source: { path: string; content: string; dvmFile?: string } }>(`/api/sessions/${id}/source?file=${encodeURIComponent(file)}&module=${encodeURIComponent(module)}`),
  wsTest: (mode: string, url: string, body: string, soap: boolean) =>
    req<{ result: WSTestResult }>('/api/wstest', { method: 'POST', body: JSON.stringify({ mode, url, body, soap }) }),
  wsLogs: (service: string, onlyFail: boolean, page: number, startFrom: string, startTo: string) =>
    req<{ items: WSLogItem[]; hasMore: boolean }>(`/api/wslogs?service=${encodeURIComponent(service)}&onlyFail=${onlyFail ? 1 : 0}&page=${page}&pageSize=200&startFrom=${encodeURIComponent(startFrom)}&startTo=${encodeURIComponent(startTo)}`),
  wsLogContent: (rowid: string) =>
    req<{ item: WSLogItem; content: WSLogContent }>(`/api/wslogs/content?rowid=${encodeURIComponent(rowid)}`),
  wsLogDebug: (rowid: string) =>
    req<{ sessionId: string; module?: string; prog?: string; runProg?: string }>('/api/wslogs/debug', { method: 'POST', body: JSON.stringify({ rowid }) }),
  sourcePreview: (module: string, prog: string) =>
    req<{ source: { path: string; content: string; dvmFile?: string } }>(`/api/source-preview?module=${encodeURIComponent(module)}&prog=${encodeURIComponent(prog)}`),
}
