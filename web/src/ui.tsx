// shadcn/ui 规范基础组件:语义 token + cn/cva,颜色一律不写死(亮暗由 index.css token 决定)
import * as React from 'react'
import { Slot } from '@radix-ui/react-slot'
import { cva, type VariantProps } from 'class-variance-authority'
import { cn } from './lib/utils'

const buttonVariants = cva(
  "inline-flex shrink-0 items-center justify-center gap-1.5 whitespace-nowrap text-sm font-medium transition-colors outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-40 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4",
  {
    variants: {
      variant: {
        default: 'bg-primary text-primary-foreground hover:bg-primary/90',
        destructive: 'bg-destructive text-white hover:bg-destructive/90',
        outline: 'border border-input bg-transparent hover:bg-accent hover:text-accent-foreground',
        secondary: 'bg-secondary text-secondary-foreground hover:bg-secondary/80',
        ghost: 'hover:bg-accent hover:text-accent-foreground',
        link: 'text-primary underline-offset-4 hover:underline',
      },
      size: {
        default: 'h-8 px-3',
        sm: 'h-7 gap-1 px-2 text-xs',
        lg: 'h-10 px-6',
        icon: 'h-8 w-8',
      },
    },
    defaultVariants: { variant: 'default', size: 'default' },
  },
)

export function Button({ className, variant, size, asChild = false, ...props }:
  React.ButtonHTMLAttributes<HTMLButtonElement> &
  VariantProps<typeof buttonVariants> & { asChild?: boolean }) {
  const Comp = asChild ? Slot : 'button'
  return <Comp className={cn(buttonVariants({ variant, size }), className)} {...props} />
}

const badgeVariants = cva(
  'inline-flex items-center border px-1.5 py-0.5 text-[11px] font-medium',
  {
    variants: {
      tone: {
        default: 'border-transparent bg-secondary text-secondary-foreground',
        green: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
        yellow: 'border-yellow-500/20 bg-yellow-500/10 text-yellow-600 dark:text-yellow-400',
        red: 'border-red-500/20 bg-red-500/10 text-red-600 dark:text-red-400',
        blue: 'border-sky-500/20 bg-sky-500/10 text-sky-600 dark:text-sky-400',
        gray: 'border-transparent bg-muted text-muted-foreground',
      },
    },
    defaultVariants: { tone: 'default' },
  },
)

export function Badge({ className, tone, ...props }:
  React.HTMLAttributes<HTMLSpanElement> & VariantProps<typeof badgeVariants>) {
  return <span className={cn(badgeVariants({ tone }), className)} {...props} />
}

export function Panel({ title, right, children, className }: {
  title: React.ReactNode; right?: React.ReactNode; children: React.ReactNode; className?: string
}) {
  return (
    <div className={cn('flex min-h-0 flex-col border border-border bg-card', className)}>
      <div className="flex h-8 shrink-0 items-center justify-between border-b border-border px-2.5">
        <span className="text-xs font-medium text-muted-foreground">{title}</span>
        {right}
      </div>
      <div className="min-h-0 flex-1 overflow-auto">{children}</div>
    </div>
  )
}

export function Input({ className, type, ...props }: React.InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input
      type={type}
      className={cn(
        'flex h-8 w-full min-w-0 border border-input bg-transparent px-2 py-1 text-sm transition-[color,box-shadow] outline-none placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 dark:bg-input/30',
        className,
      )}
      {...props}
    />
  )
}

export function Separator({ className, orientation = 'horizontal' }: {
  className?: string; orientation?: 'horizontal' | 'vertical'
}) {
  return (
    <div
      role="separator"
      className={cn(
        'shrink-0 bg-border',
        orientation === 'horizontal' ? 'h-px w-full' : 'h-full w-px',
        className,
      )}
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
