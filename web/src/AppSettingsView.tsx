// 设置页:把 tdict 可执行文件所在目录加入「用户 PATH」,之后任意位置都能直接运行 tdict。
// 用户级修改(注册表 HKCU\Environment\Path),无需管理员;可一键移除。另附运行信息。
import { useCallback, useEffect, useState } from 'react'
import { Terminal, CheckCircle2, AlertCircle, Plus, Trash2 } from 'lucide-react'
import { api, type InstallStatus } from './api'
import { Button, cn, SectionTitle } from './ui'

export function AppSettingsView() {
  const [st, setSt] = useState<InstallStatus | null>(null)
  const [configPath, setConfigPath] = useState('')
  const [listen, setListen] = useState('')
  const [err, setErr] = useState('')
  const [notice, setNotice] = useState('')
  const [busy, setBusy] = useState('')

  const refresh = useCallback(async () => {
    try {
      const [s, cfg, status] = await Promise.all([api.installStatus(), api.config(), api.status()])
      setSt(s)
      setConfigPath(cfg.configPath)
      setListen(status.listen)
      setErr('')
    } catch (e) {
      setErr(e instanceof Error ? e.message : String(e))
    }
  }, [])

  useEffect(() => { void refresh() }, [refresh])

  const add = async () => {
    setBusy('add'); setErr(''); setNotice('')
    try {
      setSt(await api.installAdd())
      setNotice('已加入用户 PATH。新开的终端可直接运行 tdict;已打开的终端需重开。')
    } catch (e) {
      setErr(e instanceof Error ? e.message : String(e))
    } finally { setBusy('') }
  }

  const remove = async () => {
    if (!window.confirm('从用户 PATH 中移除该目录?')) return
    setBusy('remove'); setErr(''); setNotice('')
    try {
      setSt(await api.installRemove())
      setNotice('已从用户 PATH 移除(需重开终端生效)。')
    } catch (e) {
      setErr(e instanceof Error ? e.message : String(e))
    } finally { setBusy('') }
  }

  const supported = !!st?.supported
  const inPath = !!st?.inUserPath

  return (
    <div className="flex h-full min-h-0 flex-col bg-zinc-50 text-zinc-800 dark:bg-zinc-950 dark:text-zinc-100">
      <header className="flex shrink-0 items-center gap-3 border-b border-zinc-200 bg-white px-4 py-2 dark:border-zinc-800 dark:bg-zinc-900">
        <h1 className="shrink-0 text-sm font-semibold">设置</h1>
        <span className="min-w-0 flex-1 truncate text-[11px] text-zinc-500 dark:text-zinc-400">
          命令行安装与运行信息
        </span>
        {err && <span className="flex shrink-0 items-center gap-1 text-[11px] text-red-600 dark:text-red-400"><AlertCircle className="h-3.5 w-3.5" />{err}</span>}
      </header>

      <main className="min-h-0 flex-1 overflow-auto p-4">
        <div className="mx-auto max-w-3xl space-y-5">
          {/* 命令行安装 */}
          <section>
            <SectionTitle>命令行安装</SectionTitle>
            <div className="mt-2 border border-zinc-200 p-3 dark:border-zinc-800">
              <div className="flex items-start gap-2">
                <Terminal className="mt-0.5 h-4 w-4 shrink-0 text-zinc-500" />
                <div className="min-w-0 flex-1">
                  <p className="text-xs text-zinc-600 dark:text-zinc-300">
                    把 tdict 所在目录加入<strong>用户 PATH</strong>(当前用户,无需管理员),之后任意位置都能直接运行 <code>tdict</code>。
                  </p>
                  <div className="mt-2 space-y-1 text-[11px] text-zinc-500 dark:text-zinc-400">
                    <div className="min-w-0 truncate" title={st?.exePath}>可执行文件:{st?.exePath || '…'}</div>
                    <div className="min-w-0 truncate" title={st?.exeDir}>将加入的目录:{st?.exeDir || '…'}</div>
                  </div>
                  <div className="mt-2 flex items-center gap-2">
                    <span className={cn('flex items-center gap-1 text-[11px]',
                      inPath ? 'text-emerald-600 dark:text-emerald-400' : 'text-amber-600 dark:text-amber-400')}>
                      {inPath ? <CheckCircle2 className="h-3.5 w-3.5" /> : <AlertCircle className="h-3.5 w-3.5" />}
                      {inPath ? '已在用户 PATH 中' : '尚未加入用户 PATH'}
                    </span>
                  </div>
                  <div className="mt-3 flex items-center gap-2">
                    <Button variant="primary" disabled={!supported || inPath || busy !== ''} onClick={() => void add()}>
                      <Plus className="h-3.5 w-3.5" />{busy === 'add' ? '添加中…' : '添加到用户 PATH'}
                    </Button>
                    <Button variant="outline" disabled={!supported || !inPath || busy !== ''} onClick={() => void remove()}>
                      <Trash2 className="h-3.5 w-3.5" />{busy === 'remove' ? '移除中…' : '从用户 PATH 移除'}
                    </Button>
                  </div>
                  {notice && <p className="mt-2 text-[11px] text-emerald-600 dark:text-emerald-400">{notice}</p>}
                  {st?.note && <p className="mt-2 text-[11px] text-zinc-500 dark:text-zinc-400">{st.note}</p>}
                  {!supported && st?.manual && (
                    <pre className="mt-2 overflow-auto bg-zinc-100 px-3 py-2 text-[11px] text-zinc-700 dark:bg-zinc-900 dark:text-zinc-200">{st.manual}</pre>
                  )}
                  <details className="mt-2">
                    <summary className="cursor-pointer text-[11px] text-zinc-500 dark:text-zinc-400">查看当前用户 PATH</summary>
                    <pre className="mt-1 max-h-40 overflow-auto whitespace-pre-wrap break-all bg-zinc-100 px-3 py-2 text-[11px] text-zinc-700 dark:bg-zinc-900 dark:text-zinc-200">{st?.userPath || '(空)'}</pre>
                  </details>
                </div>
              </div>
            </div>
          </section>

          {/* 运行信息 */}
          <section>
            <SectionTitle>运行信息</SectionTitle>
            <div className="mt-2 grid grid-cols-1 gap-2 text-xs sm:grid-cols-2">
              <div className="border border-zinc-200 p-3 dark:border-zinc-800">
                <div className="text-[11px] text-zinc-500 dark:text-zinc-400">配置文件</div>
                <div className="mt-0.5 break-all font-mono text-[11px]" title={configPath}>{configPath || '—'}</div>
              </div>
              <div className="border border-zinc-200 p-3 dark:border-zinc-800">
                <div className="text-[11px] text-zinc-500 dark:text-zinc-400">服务地址</div>
                <div className="mt-0.5 break-all font-mono text-[11px]">{listen ? `http://${listen}` : '—'}</div>
              </div>
            </div>
          </section>
        </div>
      </main>
    </div>
  )
}
