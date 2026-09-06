# TDict — ERP 数据字典 CLI / T100 作业调试工具

面向 **T100 / Genero(4GL) ERP** 的本地命令行工具，两条主线：

1. **ERP 数据字典查询**（给"读代码、做配置"提供上下文）：查 ERP 数据字典与业务元数据——数据表结构与字段中文含义（`tdict rt`）、字段校验规则（`rv`）、下拉选项/系统分类码（`scc`）、字段画面规格（`desc`）、可复用开窗（`rq`）、报错消息文本（`msg`）、系统与单据参数说明（`sysp`/`docp`）。数据随工具本地保存（`tdict db sync` 从 ERP 刷新），也可 `--conn` 直查远程库。
2. **AI 人机协同调试 T100 作业**（`tdict debug`）：通过 SSH 在 T100 服务器上驱动 `fglrun -d` 的 (fgldb) 文本调试协议，提供 Web 调试界面与命令行控制端，让 AI/人协同排查作业逻辑错误、跟踪变量、验证接口报文场景。

所有输出默认使用**简体中文 (zh_CN)**。

## 能查到的数据（规模）

查询命令覆盖的数据与当前体量（以恒烁正式区数据为例，同步后合计约 85 万行）：

| 内容 | 数据是什么 | 体量 | 查询命令 |
|---|---|---|---|
| 数据表字典 | 系统 **3,882 张表**的登记与字段定义：每张表做什么、字段中文含义/类型/长度/主键、键值与索引 | 约 50 万行 | `tdict rt <表名>`（表名示例：`dzea_t` 表登记档、`dzeb_t` 字段档、`oobd_t` 业务表） |
| 校验带值 (r.v) | 字段校验规则模板：校验 SQL、外部参数、判断条件、错误讯息 | 1,442 条 | `tdict rv [识别码]` |
| 系统分类码 (SCC) | 下拉/选项字典：分类码 →「值 → 说明」 | 2,066 个分类码、1.1 万分类值 | `tdict scc [分类码]` |
| 字段画面规格 | 字段在画面的控件/下拉来源/格式/必填等设计参考 | 15.6 万条 | `tdict desc <表名> [字段]` |
| 可复用开窗 (r.q) | 开窗选单定义：SQL、参数、显现/回传列 | 5,141 个 | `tdict rq [开窗码]` |
| 系统消息 | 所有提示/报错编号的文本、建议处理、技术细节（多语言） | 约 6 万条 | `tdict msg <编号>` |
| 参数定义 | 系统/企业/据点/单据参数的用途说明、型态、值域、预设值 | 1,034 个参数（含 2,578 条单据性质绑定） | `tdict sysp <编号>` / `tdict docp <编号>` |

查询读的是本地数据；`tdict db sync` 从 ERP 刷新后即包含以上全部内容（详见各命令节与「数据维护」）。

## 安装

### 方式一：全局安装（推荐）

安装到系统 PATH 后，可在任意目录直接调用 `tdict`：

```bash
# 1. 编译
cd TDictCli
GOPROXY=https://goproxy.cn,direct go build -o tdict.exe .

# 2. 复制到 GOPATH/bin（已在 PATH 中）
cp tdict.exe $GOPATH/bin/

# 3. 设置数据库路径环境变量（永久生效，需重启终端）
setx TDICT_DB "D:\我的项目\TDictCli\erp_data.db"

# 或者为当前会话设置
export TDICT_DB="D:\我的项目\TDictCli\erp_data.db"
```

数据库查找优先级：`$TDICT_DB` > `-d` 参数 > 可执行文件同目录 > 当前工作目录。

### 方式二：手动指定数据库路径

将 `tdict.exe` 和 `erp_data.db` 放到同一目录：
```bash
./tdict rt dzea_t -d ./erp_data.db
```

### 方式三：从源码编译

```bash
# 要求 Go 1.21+
cd TDictCli
GOPROXY=https://goproxy.cn,direct go build -o tdict.exe .
```

### 安装 Skill（推荐）

安装 Claude Code Skill 后，AI Agent 可自动理解并使用本工具。技能内容内嵌在 `tdict.exe` 中，安装的就是当前版本：

```bash
# 安装到当前项目
tdict install

# 安装到指定项目
tdict install /path/to/project
```

## 快速上手

```bash
# 查询单张表的完整字典（表名、字段、键值、索引）
tdict rt dzea_t

# 查询多个表（逗号分隔）
tdict rt "dzea_t,dzeb_t,dzed_t"

# 输出 JSON / CSV
tdict rt dzea_t --json
tdict rt dzea_t --csv

# 从 ERP 实时刷新数据
tdict db sync

# 列出全部校验带值定义，并查看指定 dzcd001 的详情
tdict rv
tdict rv v_ooba002_07

# 列出全部系统分类码 (SCC)，并查看指定 gzca001 的详情
tdict scc
tdict scc 4

# 查询指定表的字段规格（控件/SCC码/格式），以及单字段完整规格
tdict desc oobd_t
tdict desc oobd_t oobd002

# 列出全部可复用开窗，并查看指定 dzca001 的详情
tdict rq
tdict rq q_apca001

# 报错/参数编号 → 消息文本与处理建议、参数用途说明
tdict msg std-00006
tdict sysp A-SYS-0040
tdict docp D-MFG-0076

# 数据较旧先刷新;或临时直查某环境(不需要本地数据)
tdict db sync
tdict msg aoo-00120 --conn 恒烁正式区
```

## 命令

> 命名说明：`rt`（数据表，对应表格设计器 r.t）、`desc`（字段规格）、`rv`（校验带值 r.v）、`rq`（开窗 r.q）采用 T100 体系术语；旧名称 `table`/`spec`/`check`/`win` 仍可作为别名使用。

### `tdict rt <表名>`

查询一张或多张表的完整字典：这张表在系统里做什么（表说明/所属模块/表类型）、每个字段的中文含义、类型与主键必填、键值与索引。读代码、看 SQL、查界面字段含义时用它。

返回内容包含：

- **表名/表说明**：表名与中文说明、所属模块、表类型（取值含义见文末「输出中的类型码」）
- **字段**：序号、字段名、字段说明、数据类型、长度、主键、必填、备注
- **键值**：键名、类型（主键/外键/唯一）、键值字段、外键表、外键字段
- **索引**：索引名、类型、索引字段

```
$ tdict rt dzea_t
=== dzea_t ===
表说明: 数据表主档
模块: ADZ    类型: X

字段 (31):
序号  字段名        字段说明      数据类型       长度  主键  必填  备注
--  ---------  ----------  ---------  ---  --  --  --
1   dzeastus   状态码       varchar2   10
2   dzea001    table编号   varchar2   15   PK  Y
...

键值 (1):
键名       类型  键值字段   外键表  外键字段
-------  --  -------  ---  ----
dzea_pk  主键  dzea001

索引 (1):
索引名      类型  索引字段
-------  --  -------
dzea_pk  U   dzea001
```

`--json` 输出结构化对象（`表名/表说明/模块/类型/字段/键值/索引`）；`--csv` 输出字段明细（含表名列）。

### `tdict rv [dzcd001]`

查询 **字段校验规则 (r.v)**——系统保存数据前做的各种检查是可复用的"校验模板"，每条校验一个识别码：含要执行的校验 SQL（SQL 里用 `<field>`、`arg1~9`、`:TODAY` 等占位符，运行时代入实际字段与参数）、外部参数定义与判断条件。想知道某字段会被怎样校验、校验不过显示什么错误讯息，或要给字段绑定校验时用它。

校验识别码示例：`v_ooba002_07`（形如 `v_表_用途`）；也可 `tdict rv --kw 料号` 按识别码/说明搜索。

```bash
# 列出全部校验定义（识别码/客制/说明/型态/错误讯息/备注）
tdict rv

# 按关键字过滤（识别码或说明）
tdict rv --kw 料号

# 指定 dzcd001 的完整详情
tdict rv v_ooba002_07

# 指定说明语言别（默认 zh_CN）
tdict rv v_ooba002_07 --lang zh_TW

# JSON 输出（列表/详情）；--csv 仅列表模式
tdict rv v_ooba002_07 --json
tdict rv --csv
```

详情包含（标准/客制变体分别显示）：

- **单头**：型态（1=检查存在 / 2=带值 / 3=检查存在并带值）、错误讯息代号、行业别、状态码、说明、备注
- **SQL 指令**：校验 SQL 原文，附标签图例（`<field>` `<table>` `<wc>` `<count>`、`arg1~9`、`:TODAY`、`:DEPT` 等全局变量）
- **参数**：顺序、参数名称（对应 arg1~9）、日期型态（1=年月日 / 2=年月日时分秒 / 3=毫秒 / 其他=字符串）、说明、备注
- **判断条件**：顺序、条件 SQL、条件成立时显示的错误讯息

### `tdict scc [gzca001]`

查询 **系统分类码 (SCC)**——系统里各种下拉/选项的"选项字典"：币别、单据类型、料件属性等分类统一按分类码组织，一个分类码 = 一组「值 → 说明」，画面上的下拉选项就来自这里。查字段可选项、下拉数据来源、按分类码取说明时用它。

分类码示例：`4`、`241`（多为数字）；`--kw 币别` 按分类码/名称搜索。

```bash
# 列出全部分类码（分类码/群组/状态/名称/值数）
tdict scc

# 按关键字过滤（分类码或名称）
tdict scc --kw 币别

# 指定 gzca001 的完整详情
tdict scc 4

# 指定说明语言别（默认 zh_CN）
tdict scc 4 --lang zh_TW

# JSON 输出（列表/详情）；--csv 仅列表模式
tdict scc 4 --json
tdict scc --csv
```

详情包含：

- **单头**：群组、状态（Y=启用/N=停用）、名称、说明、说明2
- **分类值**：值、说明（即下拉选项显示文字）、排序（下拉按此排序）、标准/客制
- **扩展数据列**（`gzcb003`~`gzcb015`，仅显示非空值）：含义由各分类码自己定义（如某分类码用它们存群组编号区间、另一个存格式），结合该分类码的语境理解

### `tdict desc <表名> [字段名]`

查询**字段的画面规格**——每个字段在画面上用什么控件（输入框/下拉/日期…）、下拉取哪个系统分类码、显示格式与宽度、必填与否、默认值、开窗程序、最大/最小值等，画面设计器按它生成画面。要复刻/理解画面字段行为、查字段编辑控件或下拉来源时用它。

示例：`tdict desc oobd_t`（整表全部字段规格）；`tdict desc oobd_t oobd002`（单字段完整规格）。

```bash
# 列出该表全部字段规格（序号/字段/字段名/控件/SCC码/必填/宽度/格式/默认值/校验带值）
tdict desc oobd_t

# 单字段完整规格（含 SCC 名称、取值范围、开窗、串查、报表栏宽等全部列）
tdict desc oobd_t oobd002

# 指定说明语言别（默认 zh_CN）
tdict desc oobd_t oobd002 --lang zh_TW

# JSON 输出（列表/详情）；--csv 输出全列明细
tdict desc oobd_t oobd002 --json
tdict desc oobd_t --csv
```

输出语义（每条规格的字段含义）：

| 输出项 | 含义 |
|---|---|
| 控件 | 字段在画面的控件类型：03=ComboBox（下拉）05=Edit（输入）01=ButtonEdit（带开窗按钮）04=DateEdit 09=RadioGroup 11=SpinEdit 12=TextEdit 13=TimeEdit 34=DateTimeEdit 等 |
| SCC码 | 下拉/单选控件的选项来源（哪个系统分类码；用 `tdict scc <码>` 查可选项） |
| 必填 | Y/N |
| 显示宽度 | 控件宽度（"整数,小数"） |
| 格式 | 显示格式 format |
| 大小写 | U/L |
| 默认值 | 默认值 |
| 最大值/最小值 | 取值范围（配比较符号） |
| 编辑开窗 / 查询开窗 | 字段上开窗按钮弹出的程序 |
| 校验带值 | 该字段绑定的校验识别码（用 `tdict rv <码>` 查校验内容） |
| 串查型态 / 串查程序 | 串查/参考程序 |
| 报表栏宽 / 小数位 | 报表输出参考 |
| 客制识别 | 标准/客制 |

### `tdict rq [dzca001]`

查询**可复用开窗 (r.q)**——代码里 `CALL q_xxx()` 弹出的查寻选单定义：每条开窗 = 一段带占位符的 SQL（选取字段/表/条件区用 `<field>/<table>/<wc>` 标记）+ 外部参数（对应 arg1~9）+ 显现与回传列。读代码遇到开窗调用、想知道它查什么表、回传哪些值时用它。

开窗码示例：`q_apca001`（`q_` 开头）；`--kw 料号` 按说明搜索。

```bash
# 列出全部开窗（识别码/客制/说明/状态/每页笔数/HardCode/行业别）
tdict rq

# 按关键字过滤（识别码或说明）
tdict rq --kw 料号

# 指定 dzca001 的完整详情
tdict rq q_apca001

# 指定说明语言别（默认 zh_CN）
tdict rq q_apca001 --lang zh_TW

# JSON 输出（列表/详情）；--csv 仅列表模式
tdict rq q_apca001 --json
tdict rq --csv
```

详情包含（标准/客制变体分别显示）：

- **单头**：状态、SQL 指令（原文，附标签图例）、每页笔数、作业串查编号、HardCode（Y=跳过自动产生）、行业别、说明、助记码
- **参数**：顺序、参数名称（对应 argN）、日期型态（1=年月日 / 2=年月日时分秒 / 3=毫秒 / 其他=字符串）、说明
- **显现设定**：显现顺序、字段编号、表格别名、显示控件、是否回传（Y 的字段按显现顺序构成 return1~9）、大小写、显示格式、标签缀字

### `tdict msg <编号>`

查询**系统消息**——所有提示/报错消息的记录，每个消息一个编号。程序报错、日志里出现的编号，都能查到它的完整文本、建议处理方式与技术细节。消息编号格式：`std-00001`、`azz-00041`、`lib-xxxxx`（类型-流水），负整数（SQLCODE 如 `-100`）需用 `--` 分隔：`tdict msg -- -100`。

```bash
# 默认语言(zh_CN)行
tdict msg std-00006

# 指定语言
tdict msg azz-00041 --lang zh_TW

# 多编号、JSON
tdict msg "std-00006,azz-00041" --json

# 远程直查
tdict msg aoo-00120 --conn 主机正式区
```

输出长这样：

```
=== std-00006 (zh_CN) ===
类型:     错误 (1)
状态:     启用
文本:     输入的资料已存在
建议处理: 请重新审核您所需要输入的资料
```

返回：**文本**（报错原文）、**建议处理**、**建议作业**（附作业名称）、**技术细节**（给程式人员的详细讯息）、**类型**（0=警告 / 1=错误 / 2=资讯）、**状态**（启用/停用）。消息按语言各一条：默认只显示 `--lang`（zh_CN）行，该编号无此语言时命令会列出可用语言（与源系统一致：精确匹配、无自动回退）。

### `tdict sysp <编号>` / `tdict docp <编号>`

查询**系统/单据参数说明**——决定系统行为与单据流程的可配置项。每个参数有编号、名称、中文说明与"怎么填"的定义（型态/值域/预设值）。

- `tdict sysp` 查**系统/企业/据点级参数**，编号格式 = 型态码 + 领域 + 4 位流水：`A-SYS-0040`（A=系统级）、`E-CIR-0001`（E=企业级）、`S-BAS-0028`（S=据点级）；
- `tdict docp` 查**单据别参数**（编号 `D-` 开头，如 `D-MFG-0076`），并附该参数适用的单据性质（模块 + 单据性质清单）。

```bash
tdict sysp A-SYS-0040              # 系统参数(默认 zh_CN 说明)
tdict sysp S-FIN-3014 --lang zh_TW
tdict docp D-MFG-0101              # 单据别参数 + 单据性质绑定
tdict docp "D-MFG-0076,D-BAS-0058" --json
tdict docp A-SYS-0040              # 群不对会提示改用 sysp
```

输出长这样：

```
=== A-SYS-0040 (zh_CN) ===
名称:     系统是否允许LIB相关资源被签出
说明:     系统是否允许LIB相关资源被签出
群:       gzsa_t  系统级参数(系统全域参数)
型态:     Y/N (1)
领域:     SYS
预设值:   N
异常处理: 改抓设定时的预设值 (1)
状态:     启用
```

返回：名称/说明（多语言）、参数群与级别、**型态**（1=Y/N 2=整数选项 3=范围设定 4=字符或SCC 5=日期）、领域、预设值、值域、SCC 选项、校核/开窗引用、取参异常处理、修改频度、即时抓取、状态与长备注；`tdict docp` 另附**单据性质绑定**表。注意：这里查的是参数的定义与说明，各环境下参数**实际设的值**在参数值维护作业/画面里，不在本命令范围。

### `tdict db sync` — 数据维护（刷新本地查询数据）

上文各查询命令读的是本地数据；`tdict db sync` 从 ERP 把它刷新到最新（覆盖表字典、校验、分类码、画面规格、开窗、消息、参数等全部内容）。

```bash
# 全量刷新（约 85 万行，视网络约 2~4 分钟）
tdict db sync

# 只刷部分（一般不必）
tdict db sync --table dzea_t

# 指定数据源环境与库文件
tdict db sync --conn 恒烁正式区 -d D:/path/to/erp_data.db
```

- 同步需能访问 ERP 服务器（内网/VPN）；采用临时库整体替换，失败不影响原库，完成后原库自动备份为 `<库>.bak`。
- 缺省数据源 = 默认环境（activeEnv，`tdict env` 查看/切换）的库；环境与库配置见「数据库连接与查询数据源」。
- 不想维护本地数据时，查询命令加 `--conn <环境名>` 直接查远程（与本地同一批数据、同一输出）。

### `tdict install [目录]`

将 TDict 的 Claude Code 技能文件（`tdict`、`tdict-debug`、`erp-read`、`erp-modify`）安装到目标项目的 `.claude/skills/` 目录，使 Claude Code Agent 能够自动理解和使用本工具，并能安全修改框架生成的 ERP 4GL 源码。技能内容由二进制内嵌（`go:embed`），安装的就是当前版本。

```bash
# 安装到当前工作目录
tdict install

# 安装到指定项目
tdict install D:/path/to/project
```

### T100 作业调试（`tdict debug`）

通过 SSH 在 T100 服务器上驱动 `fglrun -d` 的 (fgldb) 文本调试协议，实现"人操作 GDC 界面 + AI 借助命令行检查分析"的人机协同调试。依赖 `config.json` 的 `debug` 节（SSH/区域/多环境）。

```bash
# 1. 启动调试服务(默认后台常驻、单实例;打印地址后返回,会话不被占用)
tdict debug serve                 # 前端+API: http://127.0.0.1:28670(端口占用自动顺延)
tdict debug serve --stop          # 停止后台实例
tdict debug serve --foreground    # 前台运行(日志直出)

# 2. 连接并启动作业调试(等入口停站,返回 JSON 快照)
tdict debug start bsft001_wf -m asf

# 3. 透传任意 fgldb 标准命令,原样返回输出(原生文本)
tdict debug exec "break 4450"        # 下断点
tdict debug exec "info breakpoints"
tdict debug exec "print ls_sql"      # 求值变量
tdict debug exec "continue"          # 继续(阻塞到下次停站;--timeout 300 调大)
tdict debug exec "where"             # 调用栈

# 4. AI 信息通道(原生命令行给不了的"看"):
tdict debug stop                      # 可复取的停站现场快照
tdict debug source bsft001_wf.4gl -m asf --from 4400 --to 4600  # 读服务器源码(白名单只读)
tdict debug logs --tail 50            # 会话事件日志
tdict debug locate b_fill             # 函数定义定位
tdict debug resolve bsft001_wf -m asf # 作业→实体程序/模块(gzzz_t)

# 5. 环境/会话管理
tdict debug env                       # 查看/切换当前生效环境(SSH 配置)
tdict debug topent 99                 # 设置会话 TOPENT(企业编号)
tdict debug quit                      # 结束本轮调试(会话保留可复用)
```

- **离线设置默认环境(无需 serve)**:`tdict env` 列出全部环境与当前默认;
  `tdict env <环境名>` 直接把默认环境写入 config.json(与设置页「设为默认」等效,
  不影响运行中会话;要连同调试会话一起切换才需要 `tdict debug env`)。
- 控制端命令默认自动发现后台实例真实地址(端口顺延后仍可达),也可 `--url` 显式指定。
- 断点支持持久化(下次调试自动恢复);接口日志 `wslogs`/`wsdebug` 支持报文重放调试。
- 完整用法见内嵌技能 `.claude/skills/tdict-debug.md`(含原生 fgldb 命令参考与人机交接范式)。

## 本地源码镜像（tdict mirror）

把某环境(T100 服务器)的源代码镜像到本地目录,供 **AI 用本地文件工具读码**(ls/rg/带行号读),避免每次读取都经服务器往返、也避免给 AI 服务器端查询权限。**只镜像各模块 `4gl`(源码)与 `4fd`(前端字段描述)两棵树**——per(界面源)、编译产物(42m/42r/42f)、多语言、设计器辅助目录一律不拉;目录与服务器同构(`<镜像根>/<环境名>/erp/<模块>/{4gl,4fd}/...`、`com/{lib,sub,qry,wss}/...`),客制模块 `c<mod>` 与标准模块并列,同路径下客制优先。

镜像根目录配置在 config.json 顶层 `mirror` 键(**必须显式设置**,不设默认避免隐式落盘):

```json
{ "debug": { ... }, "query": { "source": "local" }, "mirror": { "dir": "D:\\dev\\erp-src" } }
```

```bash
tdict mirror dir                     # 查看当前镜像根
tdict mirror dir D:\dev\erp-src      # 设置/更换镜像根(写入 config.json;相对路径存为绝对)
tdict mirror pull                    # 默认环境(activeEnv/首条)增量更新
tdict mirror pull 主机正式区          # 指定环境增量更新
tdict mirror pull 主机正式区 --full   # 全量重建(整目录替换,删除服务器已不存在的残留)
tdict mirror path                    # 打印镜像目录绝对路径,第一行即路径(AI 前往)
tdict mirror path 主机正式区
```

- **增量机制**:服务器 TOP 下保留 `.tdict-mirror-<环境>.mark` 基线,默认只打包 `find -newer` 的变更文件(秒级);首次拉取与 `--full` 走全量(数百 MB~数 GB,视站点规模,服务器 gzip 打包 + SFTP 流式下载,完成后清理归档、保留 marker)。
- **打包根(topDir)解析**:环境显式 `topDir` > 区域静态表 > SSH 登录探测(如恒烁 zone "1"/"2" 未显式配置时必须探测;失败会报错提示补 topDir)。
- AI 工作流:`tdict mirror path` 拿目录 → 本地工具读码;未镜像的文件(如 per)仍走 `tdict debug source` 白名单通道;调试停站行号来自服务器编译表,与镜像不同代时以服务器文件为准(先 pull)。

## 数据库连接与查询数据源（config.json）

连接配置存于项目根目录 `config.json`（JSON，开发阶段明文，勿用于生产）。数据库连接按 SSH 环境**一对一挂载**：每个环境在 `debug.sshs[]` 内嵌一个 `db`（显式 host/port + `service`(oracle) 或 `database`(kingbase) + 账号列表）；服务器侧 sqlplus/ksql 工具路径自动探测，无需配置。顶层 `query` 键记录查询命令的默认数据源（与 debug 平级，设置页保存环境不会覆盖它）：

```json
{
  "debug": {
    "sshs": [
      { "name": "主机正式区", "host": "172.16.1.109", "port": 22, "user": "tiptop",
        "password": "tiptop", "zone": "36", "topent": "99", "topDir": "/u1/topprd",
        "db": {
          "type": "oracle",
          "host": "172.16.1.109", "port": 1521,
          "service": "t35prd",
          "accounts": [
            { "account": "dsdemo", "password": "dsdemo" },
            { "account": "ds", "password": "ds" }
          ]
        } }
    ],
    "activeEnv": "主机正式区"
  },
  "query": { "source": "local" }
}
```

- **账号规则**：无"主账号"，所有账号都在 `accounts` 列表，不区分默认。客户端直连（查询数据源 / `db ping` / `db sync` / 连接测试）取列表**首项**；服务器侧调试连库由会话 TOPENT 经服务器 `gzou_t` 解析出账号名后回本表查密码（未收录时按"账号=密码"惯例）。
- `viaSsh`（可选）：客户端不可达 DB、但 DB 对 SSH 服务器可达时，经 SSH 端口转发再直连（本地起转发端口 → 驱动连 `127.0.0.1:本地端口`）；缺省远端取连接自身 host/port。
- 支持类型：`kingbase`（人大金仓，PostgreSQL 协议，pgx）、`oracle`（go-ora 纯 Go 驱动）。
- 配置文件查找优先级：`$TDICT_CONFIG` > `--config` 参数 > 可执行文件同目录 > 当前工作目录。

### 查询数据源切换（本地 SQLite ⇄ 远程库）

`rt/rv/desc/scc/rq` 五个查询命令统一走同一查询接口（`db.Source`）：本地 SQLite 镜像与远程 ERP 库查的是**同一批表(24 张字典 + 消息档)**，输出完全一致；命令层不感知数据源。解析优先级：

1. `--conn <环境名>`：本次调用远程直查该环境的库（`--conn local` 回本地）；
2. `config.json` 顶层 `query.source`：设为环境名（如 `"恒烁正式区"`）即默认远程直查该环境；
3. 缺省 `local`：本地 SQLite（`-d`/`TDICT_DB` 定位的 `erp_data.db`），与旧版行为一致。

```bash
tdict rt dzea_t                                # 默认本地 SQLite
tdict rt dzea_t --conn 恒烁正式区               # 远程直查恒烁正式区(金仓 dsdata)
tdict desc oobd_t oobd002 --conn 主机正式区     # 远程直查主机正式区(oracle dsdemo)
tdict rv v_ooba002_07 --conn local             # 显式切回本地
```

远程直查使用客户端驱动直连 `db.host:port`（账号取列表首项，`--conn`/`query.source` 与此无关），不要求本地已 sync；金仓与 Oracle 均为完整支持（列表 `--kw` 过滤大小写不敏感，与本地一致）。查询是只读单条 SELECT，值经白名单/转义内联。

`db` 子命令（连接管理，与查询数据源独立；`--conn` 语义同为环境名）：

```bash
tdict db list                                  # 按环境列出挂载的库(类型/地址/账号数)
tdict db ping [--conn <环境名>]                # 验证连接可达(只读;缺省 activeEnv 环境的库)
tdict db sync [--conn <环境名>]                # 拉取该环境字典数据写入本地 SQLite(见前节)
tdict db discover --env 主机正式区 --type oracle [--save]   # SSH 自动发现连接要素 → 预览或写入该环境 db
```

- `db discover` 给出"服务器视角"连接要素；客户端不可达时改 `db.host` 或补 `viaSsh`。
- 注意：root 的 `--conn`（查询数据源切换）与 `db sync/ping` 组内自带的 `-c/--conn` 是两个独立 flag，值都取 SSH 环境名。

## 输出中的类型码

`tdict rt` 输出的取值含义。

### 表类型(rt 的「类型」列:每张表在系统里的角色)

| 码 | 含义 |
|---|---|
| B | 基础数据 (Basic) |
| M | 主档 (Master) |
| T | 交易单头 (Transaction header) |
| D | 交易单身/明细 (Detail) |
| L | 多语言 (Language) |
| V | 提速档/视图 (View) |
| X | 系统/交叉 (Cross-reference) |
| H | 历史/暂存 (History/Temp) |

### 字段数据类型(rt 字段表的「数据类型」列常见值)

| 码 | 含义 |
|---|---|
| `N004`, `N101` | 数字 (number) |
| `C003`, `C004`, `C105` | 字符 (varchar2) |
| `D001` | 日期 (date) |
| `Z012` | 时间戳 (timestamp) |
| `Z501` | 一般 flag (Y/N) |
| `Z509` | 短说明 |
| `Z510` | 长说明 |
| `Z521` | 文件编号 (Table 引用) |
| `Z522` | 字段编号 (Field 引用) |
| `Z504` | 人员编号 |
| `Z505` | 组织编号 |

## 项目结构

```
TDictCli/
├── main.go              # 入口(内嵌 web/dist 与 .claude/skills)
├── go.mod / go.sum      # Go module
├── config.json          # debug + query(查询数据源) + mirror(源码镜像根) 配置(JSON,开发阶段明文)
├── cli/
│   ├── root.go          # 根命令 + 全局 --json/--csv/-d/--config/--conn
│   ├── table.go         # tdict rt(表名/字段/键值/索引)
│   ├── check.go         # tdict rv(校验带值 dzcd001 查询)
│   ├── scc.go           # tdict scc(系统分类码 gzca001 查询)
│   ├── spec.go          # tdict desc(字段规格 dzep_t 查询)
│   ├── win.go           # tdict rq(可复用开窗 dzca001 查询)
│   ├── source.go        # 查询数据源解析(本地 SQLite ⇄ --conn/query.source 远程库)
│   ├── mirror.go        # tdict mirror(本地源码镜像 dir/pull/path)
│   ├── install.go       # tdict install(安装内嵌 Claude Code 技能)
│   ├── db.go            # tdict db 命令组(连接配置加载)
│   ├── db_sync.go       # tdict db sync(拉取 ERP 字典数据 → SQLite)
│   ├── dbops.go         # tdict db list / ping / discover(连接管理 + SSH 自动发现)
│   ├── debug*.go        # tdict debug 命令族(serve/start/exec/...)
│   ├── debugctl.go      # debug 控制端(REST 薄封装)
│   ├── debugctx.go      # debug 信息通道(stop/source/logs/locate/resolve/interrupt)
│   ├── debugenv.go      # debug env/topent 切换
│   └── servebg*.go      # debug serve 后台常驻(单实例/端口顺延/stop)
├── dbconfig/
│   └── dbconfig.go      # db 连接配置模型(dbconfig.Connection;每环境挂一个)
├── erpdb/
│   ├── erpdb.go         # ERP 连接器(只读 Query);kingbase 经 pgx
│   ├── oracle.go        # oracle 经 go-ora
│   └── ident.go         # SQL 标识符/字面量安全(查询内联)
├── db/
│   ├── db.go            # SQLite 连接与查询封装 + db.Source 接口定义
│   ├── check.go         # 校验带值查询(dzcd_t/dzce_t/dzch_t 等)
│   ├── scc.go           # 系统分类码查询(gzca_t/gzcb_t 等)
│   ├── spec.go          # 字段规格查询(dzep_t)
│   ├── win.go           # 可复用开窗查询(dzca_t/dzcb_t/dzcc_t 等)
│   └── source.go        # 查询数据源统一接口(db.Source)/缺表判定
├── debug/
│   ├── api.go           # 调试服务 REST + WS + 静态前端
│   ├── session.go       # fgldb 会话/PTY 协议驱动
│   ├── manager.go       # 会话管理器 + 事件流/源码读取/作业解析
│   ├── config.go        # debug 配置(SSH/envs/TOPENT 等)
│   ├── db.go            # 数据库探查(服务器端 sqlplus/ksql)
│   ├── mirror.go        # 源码镜像引擎(服务器 4gl/4fd 打包 → SFTP 下载 → 本地解压)
│   ├── sshx.go          # SSH/SFTP/PTY 封装
│   └── ...              # parser/tenv/wslog/bpsstore 等
├── live/
│   ├── live.go          # 远程数据源(live.Open;viaSsh 隧道支持)
│   └── source.go        # 远程查询实现(db.Source 的金仓/Oracle 方言)
├── sshtun/
│   └── sshtun.go        # SSH 端口转发隧道组件
├── output/
│   └── output.go        # 表格/JSON/CSV 输出格式化
├── web/                 # 调试 Web 前端(React/Vite;产物嵌入 dist)
├── .claude/
│   └── skills/
│       ├── tdict.md              # 数据字典查询 Skill
│       ├── tdict-debug.md        # T100 作业调试 Skill
│       ├── erp-read.md           # 阅读/分析 ERP 4GL 源码 Skill
│       └── erp-modify.md         # 修改框架生成 4GL 源码(add-point)Skill
├── erp_data.db          # SQLite 数据库(tdict db sync 刷新)
├── dist/                # 便携版打包产物(build_portable.bat)
└── README.md
```

## 技术栈

| 项目 | 选择 |
|---|---|
| 语言 | Go 1.26 |
| SQLite 驱动 | [modernc.org/sqlite](https://modernc.org/sqlite)（纯 Go，无 CGO） |
| Kingbase 驱动 | [jackc/pgx](https://github.com/jackc/pgx)（纯 Go，PostgreSQL 协议） |
| Oracle 驱动 | [sijms/go-ora/v2](https://github.com/sijms/go-ora)（纯 Go） |
| SSH/SFTP | [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto/ssh) + [pkg/sftp](https://github.com/pkg/sftp) |
| WebSocket | [coder/websocket](https://github.com/coder/websocket) |
| CLI 框架 | [cobra](https://github.com/spf13/cobra) |
| Web 前端 | React + Vite + Monaco（调试界面） |
| 输出 | `text/tabwriter` + `encoding/json` |
