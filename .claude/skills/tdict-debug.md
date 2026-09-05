---
name: tdict-debug
description: 通过 TDict 命令行调试 T100 ERP 作业(4GL/Genero)——对指定作业启动 fgldb 调试会话、透传全部标准调试命令(break/continue/print/where/info/watch 等)并得到原生输出、查询接口报文日志、按指定日志重放调试。用于排查 T100 作业逻辑错误、跟踪变量取值、验证接口报文场景。所有输出使用简体中文 (zh_CN)。
---

# TDict Debug — T100 作业命令行调试

## 前提

- 先运行 `tdict debug serve` 启动调试服务(默认**后台常驻、单实例**,打印地址后立即返回,当前会话可继续输入命令):
  ```bash
  tdict debug serve                # 后台启动;Web 界面 http://127.0.0.1:28670(端口占用自动顺延)
  tdict debug serve --stop         # 停止后台实例
  tdict debug serve --foreground   # 前台运行,日志直出终端(调试用)
  ```
  重复执行 serve 不会起第二个实例,只会提示已在运行;控制端命令(start/exec/…)自动发现后台实例真实地址。
- `config.json` 的 `debug` 节配置了 SSH 与区域
- 作业的 GUI 界面会弹到用户的 GDC 上;断点命中时程序挂起,用户可在 GDC 操作后再继续

## 工作流(AI 调试范式)

```bash
# 1. 启动调试:连 SSH → 启动作业 → 等入口停站(返回 JSON 快照)
tdict debug start bsft001_wf -m asf

# 2. 透传任意 fgldb 标准命令,原样返回输出
tdict debug exec "break 4450"          # 按行下断点
tdict debug exec "break bsft001_wf.b_fill"   # 按函数下断点
tdict debug exec "info breakpoints"    # 查看断点
tdict debug exec "continue"            # 继续运行(阻塞到下次停站)
tdict debug exec "where"               # 调用栈
tdict debug exec "print ls_sql"        # 求值变量/表达式(原生格式)
tdict debug exec "print g_qryparam.cond"  # record 字段
tdict debug exec "next"                # 步过
tdict debug exec "ptype g_qryparam"    # 变量类型/结构
tdict debug exec "info locals"         # 局部变量

# 3. 结束会话(作业窗口随之关闭)
tdict debug quit
```

`exec` 透传全部 fgldb 标准命令并返回**原生文本**,命令清单见 Genero 文档
Debugger commands(backtrace/where、break、call、clear、continue、delete、disable、
display、down/up、enable、finish、frame、ignore、info、list、next、output、print、
ptype、run、set、signal、step、tbreak、until、watch、whatis)。
注意:`continue`/`run`/`until` 会阻塞到程序再次停站才返回输出,耗时长时加
`--timeout 300`。

## 环境(SSH)切换与 TOPENT

调试前可先查看/切换当前生效的服务器环境(SSH 配置),并设置本轮的 TOPENT(企业编号):

```bash
# 查看当前生效环境与全部已配置环境(含各自 SSH/区域/数据库与 TOPENT 默认)
tdict debug env

# 切换到另一环境:持久化 activeEnv 并把当前会话重连到该环境(空闲 idle)
tdict debug env 恒烁测试区

# 查看当前会话的 TOPENT override 与配置级默认
tdict debug topent

# 设置会话级 TOPENT(仅会话空闲时;下一轮调试启动采用;值不限数字/文本)
tdict debug topent 99

# 清除会话级 override(回退到该环境 db.ent 的配置默认)
tdict debug topent --clear
```

- `env`/`topent` 都依赖运行中的 `tdict debug serve`,默认自动发现其地址。
- `env <名称>` 会把目标环境设为默认并结束当前调试重连过去,切换前请确认无重要调试进行。
- TOPENT 也可以在会话外由配置决定:每个环境的 `db.ent` 就是它的 TOPENT 默认值;
  会话级 override 优先于配置默认,`--clear` 后回到配置默认。

## 接口报文日志调试

排查接口(wssp/awsp)报文问题时:

```bash
# 列出日志(可 --service wssp900 过滤、--fail 只看失败)
tdict debug wslogs --service wssp900 --fail

# 按指定日志重放调试:解析该次调用的作业与报文参数,自动重放并停在入口
tdict debug wsdebug <rowid>
# 之后同样用 exec 透传调试命令
```

## 约定与注意事项

- 同一时间只允许一个调试会话;`start`/`wsdebug` 前无需手动 quit,后端会自动结束旧会话
- 只调试**测试区**(config zone),禁止对生产区随意下断点
- `exec "print"` 大数组输出可能很长,优先 print 具体字段
- 用户在 GDC 上的操作(点按钮/单据流)会驱动程序走到断点;等待用户操作时
  `exec "continue"` 可能长时间阻塞,属正常现象
- 会话状态随时可查:`tdict debug status`

## 原生 fgldb 参考(Genero BDL User Guide 6.00 节选)

> 来源:D:\T100\4gl文档\BDL-Markdown\13_programming-tools\
> `2520-fgldb.md`(fgldb 工具页)与 `2586-debugger-commands.md`(Debugger commands 清单)。
> `tdict debug exec "<命令>"` 等价于在原生 `(fgldb)` 提示符下逐条输入,输出为原生文本。

### fgldb — interface program for remote debugging

The fgldb command line tool is an interface tool to attach to the FGL integrated
debugger remotely.

**Syntax 1: Debugging an application running on the computer**

```
fgldb -p process-id
```

- `process-id` is the process identifier of the `fglrun` process.

**Syntax 2: Debugging an app running on a mobile device**

```
fgldb -m host[:port]
```

- `host` is the hostname or IP address of the mobile device where the program executes.
- `port` is the TCP debug port number to connect to, default is 6400.

**Options**

| Option | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |
| `-p process-id` | Attach to an `fglrun` process running on the same computer, by using its process id. |
| `-m host[:port]` | Attach to a running mobile app to debug, using the mobile device hostname/IP and debug port. |

The fgldb tool can be used to:

- Debug an `fglrun` process currently running on the same computer, by using its process id.
- Debug an application on a mobile device, by using the mobile device IP address and its TCP debug port.

### Debugger commands — full list

All commands below are passed through with `tdict debug exec "..."`.
As noted in the workflow above, `continue`/`run`/`until` (and similar commands that
resume program execution) block until the program stops again — add `--timeout 300`
(or larger) when running them.

| Command | Description |
| --- | --- |
| `backtrace` / `where` | Prints a summary of how your program reached the current state. |
| `break` | Defines a breakpoint to stop the program execution at a given line or function. |
| `call` | Calls a function in the program. |
| `clear` | Clears the breakpoint at a specified line or function. |
| `continue` | Continues the execution of the program after a breakpoint. |
| `delete` | Removes breakpoints that you have specified in your debugger session. |
| `detach` | Closes the TCP connection of a remote debug session. |
| `disable` | Disables the specified breakpoint. |
| `display` | Displays the specified expression's value each time program execution stops. |
| `down` | Moves down in the call stack. |
| `echo` | Prints the specified text as prompt. |
| `enable` | Enables breakpoints that have previously been disabled. |
| `finish` | Continues the execution of a program until the current function returns normally. |
| `frame` | Selects and prints a stack frame. |
| `help` | Provides information about debugger commands. |
| `ignore` | Defines the number of times a breakpoint must be ignored. |
| `info` | Describes the current state of your program (e.g. `info breakpoints`, `info locals`, `info sources`). |
| `list` | Prints source code lines of the program being executed. |
| `next` | Continues running the program by executing the next source line in the current stack frame, and then stops. |
| `output` | Prints only the value of the specified expression, suppressing any other output. |
| `print` | Displays the current value of the specified expression. |
| `ptype` | Prints the data type or structure of a variable. |
| `quit` | Terminates the debugger session. |
| `run` | Starts the program. |
| `set` | Configures your debugger session and changes program variable values (e.g. `set verbose on`). |
| `source` | Executes a file of debugger commands. |
| `signal` | Sends an interruption signal to the program. |
| `step` | Continues running the program by executing the next line of source code, and then stops. |
| `tbreak` | Sets a temporary breakpoint. |
| `tty` | Resets the default program input and output for future `run` commands. |
| `undisplay` | Cancels expressions to be displayed when the program execution stops. |
| `until` | Continues running the program until the specified location is reached. |
| `up` | Selects and prints the function that called this one, or the function specified by the frame number in the call stack. |
| `watch` | Sets a watchpoint for an expression (optionally with a condition, e.g. `watch i if i >= 3`). |
| `whatis` | Prints the data type of a variable. |
