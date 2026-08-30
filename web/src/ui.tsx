// shadcn 风格基础组件(深色 zinc 主题,免依赖手写)
import React from 'react'

export function Button({ className = '', variant = 'default', size = 'default', ...props }: React.ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: 'default' | 'outline' | 'ghost' | 'destructive' | 'secondary'
  size?: 'default' | 'sm' | 'icon'
}) {
  const variants: Record<string, string> = {
    default: 'bg-zinc-100 text-zinc-900 hover:bg-zinc-300',
    outline: 'border border-zinc-700 bg-transparent hover:bg-zinc-800 text-zinc-200',
    ghost: 'hover:bg-zinc-800 text-zinc-300',
    destructive: 'bg-red-600 text-white hover:bg-red-500',
    secondary: 'bg-zinc-800 text-zinc-200 hover:bg-zinc-700',
  }
  const sizes: Record<string, string> = {
    default: 'h-8 px-3 text-sm',
    sm: 'h-7 px-2 text-xs',
    icon: 'h-8 w-8',
  }
  return (
    <button
      className={`inline-flex items-center justify-center gap-1 rounded-md font-medium transition-colors disabled:pointer-events-none disabled:opacity-40 ${variants[variant]} ${sizes[size]} ${className}`}
      {...props}
    />
  )
}

export function Badge({ className = '', tone = 'default', ...props }: React.HTMLAttributes<HTMLSpanElement> & {
  tone?: 'default' | 'green' | 'yellow' | 'red' | 'blue' | 'gray'
}) {
  const tones: Record<string, string> = {
    default: 'bg-zinc-800 text-zinc-300 border-zinc-700',
    green: 'bg-emerald-950 text-emerald-400 border-emerald-800',
    yellow: 'bg-yellow-950 text-yellow-400 border-yellow-800',
    red: 'bg-red-950 text-red-400 border-red-800',
    blue: 'bg-sky-950 text-sky-400 border-sky-800',
    gray: 'bg-zinc-900 text-zinc-500 border-zinc-800',
  }
  return (
    <span
      className={`inline-flex items-center rounded-md border px-1.5 py-0.5 text-[11px] font-medium ${tones[tone]} ${className}`}
      {...props}
    />
  )
}

export function Panel({ title, right, children, className = '' }: {
  title: React.ReactNode; right?: React.ReactNode; children: React.ReactNode; className?: string
}) {
  return (
    <div className={`flex min-h-0 flex-col rounded-sm border border-zinc-800 bg-zinc-900/60 ${className}`}>
      <div className="flex h-8 shrink-0 items-center justify-between border-b border-zinc-800 px-2.5">
        <span className="text-xs font-medium text-zinc-400">{title}</span>
        {right}
      </div>
      <div className="min-h-0 flex-1 overflow-auto">{children}</div>
    </div>
  )
}

export function Input({ className = '', ...props }: React.InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input
      className={`h-8 rounded-md border border-zinc-700 bg-zinc-950 px-2 text-sm text-zinc-200 placeholder:text-zinc-600 focus:border-zinc-500 focus:outline-none ${className}`}
      {...props}
    />
  )
}

export const stateTone = (s: string): 'green' | 'yellow' | 'red' | 'gray' | 'blue' => {
  if (s === 'stopped') return 'yellow'
  if (s === 'running') return 'green'
  if (s === 'exit') return 'red'
  if (s === 'loading') return 'blue'
  return 'gray'
}

export const stateLabel = (s: string) => {
  const m: Record<string, string> = { stopped: '已停站', running: '运行中', exit: '已退出', loading: '启动中' }
  return m[s] || '未连接'
}
