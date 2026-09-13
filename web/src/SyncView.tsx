// 数据同步页:在浏览器里选环境执行 tdict db sync(远程 ERP 字典 → 本地 SQLite),并显示进度。
//
// 后端 POST /api/dbsync 在服务端后台执行(dbsync.Run),前端轮询 /api/dbsync 画进度:
// 逐表拉取(第 x/y 张表 + 当前表行数/累计行数)→ 建索引 → 原子替换本地库(原库备份 .bak)。
import { useCallback, useEffect, useState } from 'react'
import { RefreshCw, AlertCircle, CheckCircle2, Database } from 'lucide-react'
import { api, type DBSyncResp } from './api'
import { Button, cn, SectionTitle } from './ui'

const PHASE_LABEL: Record<string, string> = {
  open: '连接远程数据库…',
  table: '拉取数据表…',
  index: '创建主键索引…',
  replace: '替换本地数据库…',
  done: '完成',
  error: '失败',
}

export function SyncView() {
  const [data, setData] = useState<DBSyncResp | null>(null)
  const [env, setEnv] = useState('')
  const [err, setErr] = useState('')
  const [busy, setBusy] = useState(false)

  const refresh = useCallback(async () => {
    try {
      const d = await api.dbsync()
      setData(d)
      setErr('')
      setEnv((prev) => {
        if (prev && d.envs.some((e) => e.name === prev)) return prev
        if (d.activeEnv && d.envs.some((e) => e.name === d.activeEnv)) return d.activeEnv
        return d.envs[0]?.name || ''
      })
    } catch (e) {
      setErr(e instanceof Error ? e.message : String(e))
    }
  }, [])

  useEffect(() => { void refresh() }, [refresh])

  const running = !!data?.job?.running
  // 同步进行中每 800ms 刷新进度;结束后停止轮询
  useEffect(() => {
    if (!running) return
    const t = window.setInterval(() => { void refresh() }, 800)
    return () => window.clearInterval(t)
  }, [running, refresh])

  const run = async () => {
    if (!env) { setErr('请先选择环境'); return }
    if (!window.confirm(`从环境「${env}」拉取字典数据到:\n${data?.target || ''}\n\n将覆盖本地 SQLite(原库自动备份为 .bak),约 85 万行,需数分钟。继续?`)) return
    setBusy(true); setErr('')
    try {
      await api.dbsyncRun(env)
      await refresh()
    } catch (e) {
      setErr(e instanceof Error ? e.message : String(e))
    } finally { setBusy(false) }
  }

  const job = data?.job
  const pct = !job ? 0
    : job.phase === 'done' ? 100
      : job.tableTotal > 0 ? Math.min(100, Math.round((job.tableIndex * 100) / job.tableTotal)) : 0
  const cur = data?.envs.find((e) => e.name === env)

  return (
    <div className="flex h-full min-h-0 flex-col bg-zinc-50 text-zinc-800 dark:bg-zinc-950 dark:text-zinc-100">
      <header className="flex shrink-0 items-center gap-3 border-b border-zinc-200 bg-white px-4 py-2 dark:border-zinc-800 dark:bg-zinc-900">
        <h1 className="shrink-0 text-sm font-semibold">数据同步</h1>
        <span className="min-w-0 flex-1 truncate text-[11px] text-zinc-500 dark:text-zinc-400" title={data?.target}>
          从远程 ERP 拉取数据字典(表字典/校验/分类码/画面规格/开窗/消息/参数)到本地 SQLite
        </span>
        {err && <span className="flex shrink-0 items-center gap-1 text-[11px] text-red-600 dark:text-red-400"><AlertCircle className="h-3.5 w-3.5" />{err}</span>}
      </header>

      <main className="min-h-0 flex-1 overflow-auto p-4">
        <div className="mx-auto max-w-3xl space-y-4">
          {/* 目标库 */}
          <section>
            <SectionTitle>目标数据库</SectionTitle>
            <div className="mt-2 flex items-center gap-2">
              <Database className="h-3.5 w-3.5 shrink-0 text-zinc-500" />
              <span className="min-w-0 truncate text-xs text-zinc-600 dark:text-zinc-300" title={data?.target}>{data?.target || '(解析中…)'}</span>
            </div>
            <p className="mt-1 text-[11px] text-zinc-500 dark:text-zinc-400">
              同步写入该文件(查询命令读的就是它);原库自动备份为 <code>.bak</code>,失败不影响原库。
            </p>
          </section>

          {/* 环境 */}
          <section>
            <SectionTitle>环境</SectionTitle>
            {!data && <div className="mt-2 text-xs text-zinc-500">加载中…</div>}
            {data && data.envs.length === 0 && (
              <div className="mt-2 border border-dashed border-zinc-300 p-3 text-xs text-zinc-500 dark:border-zinc-700">
                没有可同步的环境(需先在「环境配置」为环境挂上数据库)。
              </div>
            )}
            {data && data.envs.length > 0 && (
              <>
                <div className="mt-2 flex items-center gap-2">
                  <span className="shrink-0 text-[11px] text-zinc-500 dark:text-zinc-400">环境</span>
                  <select value={env} disabled={running} onChange={(e) => setEnv(e.target.value)}
                    className="h-7 min-w-0 flex-1 border border-zinc-300 bg-white px-2 text-xs text-zinc-800 outline-none focus:border-sky-500 disabled:opacity-60 dark:border-zinc-700 dark:bg-zinc-950 dark:text-zinc-100">
                    {data.envs.map((e) => (
                      <option key={e.name} value={e.name}>
                        {e.name}{e.name === data.activeEnv ? '（默认）' : ''}
                      </option>
                    ))}
                  </select>
                </div>
                {cur && (
                  <div className="mt-1 text-[11px] text-zinc-500 dark:text-zinc-400">{cur.type} · {cur.address}</div>
                )}
              </>
            )}
            <div className="mt-2">
              <Button variant="primary" disabled={running || busy || !env || !data?.envs.length} onClick={() => void run()}>
                <RefreshCw className={cn('h-3.5 w-3.5', running && 'animate-spin')} />{running ? '同步中…' : '开始同步'}
              </Button>
            </div>
          </section>

          {/* 进度 */}
          {job && (job.running || job.done) && (
            <section>
              <SectionTitle>同步进度</SectionTitle>
              <div className="mt-2 border border-zinc-200 p-3 dark:border-zinc-800">
                <div className="mb-2 flex items-center gap-2 text-xs">
                  <span className="font-medium">{job.env}</span>
                  <span className="text-zinc-500 dark:text-zinc-400">{PHASE_LABEL[job.phase] || job.phase}</span>
                  <span className="ml-auto text-[11px] text-zinc-500 dark:text-zinc-400">
                    {job.running ? job.elapsed : (job.elapsed ? '用时 ' + job.elapsed : '')}
                  </span>
                </div>
                <div className="h-2 w-full overflow-hidden bg-zinc-200 dark:bg-zinc-800">
                  <div className={cn('h-full transition-[width] duration-300', job.phase === 'error' ? 'bg-red-500' : 'bg-sky-500')}
                    style={{ width: pct + '%' }} />
                </div>
                <div className="mt-2 grid grid-cols-3 gap-2 text-[11px] text-zinc-500 dark:text-zinc-400">
                  <span>数据表:{job.tableTotal > 0 ? `${job.tableIndex}/${job.tableTotal}` : '—'}{job.tables > 0 ? `(完成 ${job.tables})` : ''}</span>
                  <span>当前表:{(job.table || '—') + (job.tableRows > 0 ? ` ${job.tableRows} 行` : '')}</span>
                  <span>累计行数:{job.totalRows > 0 ? job.totalRows.toLocaleString() : '—'}</span>
                </div>
                {job.message && (
                  <p className={cn('mt-2 flex items-center gap-1 text-xs',
                    job.phase === 'error' ? 'text-red-600 dark:text-red-400'
                      : job.phase === 'done' ? 'text-emerald-600 dark:text-emerald-400' : 'text-zinc-600 dark:text-zinc-300')}>
                    {job.phase === 'done' && <CheckCircle2 className="h-3.5 w-3.5" />}
                    {job.phase === 'error' && <AlertCircle className="h-3.5 w-3.5" />}
                    {job.error || job.message}
                  </p>
                )}
                {job.backup && <p className="mt-1 text-[11px] text-zinc-500 dark:text-zinc-400">原库已备份:{job.backup}</p>}
                {job.warning && <pre className="mt-1 whitespace-pre-wrap text-[11px] text-amber-600 dark:text-amber-400">{job.warning}</pre>}
              </div>
            </section>
          )}
        </div>
      </main>
    </div>
  )
}
