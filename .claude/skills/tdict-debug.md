---
name: tdict-debug
description: 通过 TDict 命令行调试 T100 ERP 作业(4GL/Genero)——对指定作业启动 fgldb 调试会话、透传全部标准调试命令(break/continue/print/where/info/watch 等)并得到原生输出、查询接口报文日志、按指定日志重放调试。用于排查 T100 作业逻辑错误、跟踪变量取值、验证接口报文场景。所有输出使用简体中文 (zh_CN)。
---

# TDict Debug — T100 作业命令行调试

## 前提

- 用户本机已运行 `tdict debug serve`(Web 界面同时可用;未运行时先启动它)
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
