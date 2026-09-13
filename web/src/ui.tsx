// 轻量 UI 基元(纯 Tailwind,无第三方组件库):配置页只需要输入框/按钮/字段/表格。
import type { ButtonHTMLAttributes, InputHTMLAttributes, ReactNode } from 'react'

export function cn(...xs: (string | false | null | undefined)[]) {
  return xs.filter(Boolean).join(' ')
}

type Variant = 'primary' | 'secondary' | 'outline' | 'ghost'

const VARIANTS: Record<Variant, string> = {
  primary: 'bg-sky-600 text-white hover:bg-sky-500 disabled:bg-sky-600/50',
  secondary: 'bg-zinc-800 text-zinc-100 hover:bg-zinc-700 dark:bg-zinc-100 dark:text-zinc-900 dark:hover:bg-white',
  outline: 'border border-zinc-300 text-zinc-700 hover:bg-zinc-100 dark:border-zinc-700 dark:text-zinc-200 dark:hover:bg-zinc-800',
  ghost: 'text-zinc-500 hover:bg-zinc-100 hover:text-zinc-800 dark:hover:bg-zinc-800 dark:hover:text-zinc-100',
}

export function Button({
  variant = 'outline', size = 'sm', className, ...props
}: ButtonHTMLAttributes<HTMLButtonElement> & { variant?: Variant; size?: 'xs' | 'sm' }) {
  return (
    <button
      type="button"
      className={cn(
        'inline-flex shrink-0 items-center justify-center gap-1 whitespace-nowrap transition-colors disabled:cursor-not-allowed disabled:opacity-50',
        size === 'xs' ? 'h-6 px-1.5 text-[11px]' : 'h-7 px-2.5 text-xs',
        VARIANTS[variant],
        className,
      )}
      {...props}
    />
  )
}

export function Input({ className, ...props }: InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input
      className={cn(
        'h-7 w-full min-w-0 border border-zinc-300 bg-white px-2 text-xs text-zinc-800 outline-none',
        'placeholder:text-zinc-400 focus:border-sky-500 dark:border-zinc-700 dark:bg-zinc-950 dark:text-zinc-100 dark:placeholder:text-zinc-600',
        className,
      )}
      {...props}
    />
  )
}

export function Field({ label, children, className = '' }: { label: string; children: ReactNode; className?: string }) {
  return (
    <label className={cn('block', className)}>
      <div className="mb-1 text-[11px] text-zinc-500 dark:text-zinc-400">{label}</div>
      {children}
    </label>
  )
}

export function SectionTitle({ children }: { children: ReactNode }) {
  return (
    <div className="flex items-center gap-2">
      <span className="shrink-0 text-[11px] font-medium uppercase tracking-wide text-zinc-500 dark:text-zinc-400">{children}</span>
      <span className="h-px flex-1 bg-zinc-200 dark:bg-zinc-800" />
    </div>
  )
}
