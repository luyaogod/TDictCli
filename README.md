# TDict — ERP 数据字典 CLI

面向 **T100 / Genero(4GL) ERP** 的本地命令行工具：

**ERP 数据字典查询**（给"读代码、做配置"提供上下文）：查 ERP 数据字典与业务元数据——数据表结构与字段中文含义（`tdict r.t`）、字段校验规则（`tdict r.v`）、下拉选项/系统分类码（`scc`）、字段画面规格（`desc`）、可复用开窗（`tdict r.q`）、报错消息文本（`msg`）、系统与单据参数说明（`sysp`/`docp`）。数据随工具本地保存（`tdict db sync` 从 ERP 刷新），也可 `--conn` 直查远程库。

此外保留 SSH 环境与数据库连接管理（`tdict serve` 可视化配置页 / `tdict env` / `tdict db`）与服务器源码镜像（`tdict mirror`）。

T100 作业调试（原 `tdict debug`）已拆分到独立项目，本仓库不再包含调试代码。

所有输出默认使用**简体中文 (zh_CN)**。

## 能查到的数据（规模）

查询命令覆盖的数据与当前体量（以正式区数据为例，同步后合计约 85 万行）：

| 内容 | 数据是什么 | 体量 | 查询命令 |
|---|---|---|---|
| 数据表字典 | 系统 **3,882 张表**的登记与字段定义：每张表做什么、字段中文含义/类型/长度/主键、键值与索引 | 约 50 万行 | `tdict r.t <表名>`（表名示例：`dzea_t` 表登记档、`dzeb_t` 字段档、`oobd_t` 业务表） |
| 校验带值 (r.v) | 字段校验规则模板：校验 SQL、外部参数、判断条件、错误讯息 | 1,442 条 | `tdict r.v [识别码]` |
| 系统分类码 (SCC) | 下拉/选项字典：分类码 →「值 → 说明」 | 2,066 个分类码、1.1 万分类值 | `tdict scc [分类码]` |
| 字段画面规格 | 字段在画面的控件/下拉来源/格式/必填等设计参考 | 15.6 万条 | `tdict desc <表名> [字段]` |
| 可复用开窗 (r.q) | 开窗选单定义：SQL、参数、显现/回传列 | 5,141 个 | `tdict r.q [开窗码]` |
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
./tdict r.t dzea_t -d ./erp_data.db
```

### 方式三：从源码编译

```bash
# 要求 Go 1.21+
cd TDictCli
GOPROXY=https://goproxy.cn,direct go build -o tdict.exe .
```

### 安装 Skill（推荐）

AI 技能文件（`tdict`、`tdict-debug`、`erp-read`）以普通目录 **`skills/`** 与 `tdict.exe` 放在一起（便携版/仓库里都是），**不内嵌进二进制**，可直接编辑：

```bash
# 把 exe 同目录的 skills/ 复制到当前目录(生成 ./skills/)
tdict install skills

# 之后按所用 AI 工具自己摆放,例如给 Claude Code 用:
#   mv skills .claude/skills
```

## 快速上手

```bash
# 查询单张表的完整字典（表名、字段、键值、索引）
tdict r.t dzea_t

# 查询多个表（逗号分隔）
tdict r.t "dzea_t,dzeb_t,dzed_t"

# 输出 JSON / CSV
tdict r.t dzea_t --json
tdict r.t dzea_t --csv

# 从 ERP 实时刷新数据
tdict db sync

# 列出全部校验带值定义，并查看指定 dzcd001 的详情
tdict r.v
tdict r.v v_ooba002_07

# 列出全部系统分类码 (SCC)，并查看指定 gzca001 的详情
tdict scc
tdict scc 4

# 查询指定表的字段规格（控件/SCC码/格式），以及单字段完整规格
tdict desc oobd_t
tdict desc oobd_t oobd002

# 列出全部可复用开窗，并查看指定 dzca001 的详情
tdict r.q
tdict r.q q_apca001

# 报错/参数编号 → 消息文本与处理建议、参数用途说明
tdict msg std-00006
tdict sysp A-SYS-0040
tdict docp D-MFG-0076

# 数据较旧先刷新;或临时直查某环境(不需要本地数据)
tdict db sync
tdict msg aoo-00120 --conn 正式区
```

## 命令

> 命名说明：`r.t`（数据表，对应表格设计器 r.t）、`r.v`（校验带值 r.v）、`r.q`（开窗 r.q）、`desc`（字段规格）、`scc`（分类码）均对齐 T100 原生工具/术语；短名 `rt`/`rv`/`rq` 与旧名 `table`/`check`/`win`/`spec` 仍作为别名可用。

### `tdict r.t <表名>`

查询一张或多张表的完整字典：这张表在系统里做什么（表说明/所属模块/表类型）、每个字段的中文含义、类型与主键必填、键值与索引。读代码、看 SQL、查界面字段含义时用它。

返回内容包含：

- **表名/表说明**：表名与中文说明、所属模块、表类型（取值含义见文末「输出中的类型码」）
- **字段**：序号、字段名、字段说明、数据类型、长度、主键、必填、备注
- **键值**：键名、类型（主键/外键/唯一）、键值字段、外键表、外键字段
- **索引**：索引名、类型、索引字段

```
$ tdict r.t dzea_t
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

### `tdict r.v [dzcd001]`

查询 **字段校验规则 (r.v)**——系统保存数据前做的各种检查是可复用的"校验模板"，每条校验一个识别码：含要执行的校验 SQL（SQL 里用 `<field>`、`arg1~9`、`:TODAY` 等占位符，运行时代入实际字段与参数）、外部参数定义与判断条件。想知道某字段会被怎样校验、校验不过显示什么错误讯息，或要给字段绑定校验时用它。

校验识别码示例：`v_ooba002_07`（形如 `v_表_用途`）；也可 `tdict r.v --kw 料号` 按识别码/说明搜索。

```bash
# 列出全部校验定义（识别码/客制/说明/型态/错误讯息/备注）
tdict r.v

# 按关键字过滤（识别码或说明）
tdict r.v --kw 料号

# 指定 dzcd001 的完整详情
tdict r.v v_ooba002_07

# 指定说明语言别（默认 zh_CN）
tdict r.v v_ooba002_07 --lang zh_TW

# JSON 输出（列表/详情）；--csv 仅列表模式
tdict r.v v_ooba002_07 --json
tdict r.v --csv
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
| 校验带值 | 该字段绑定的校验识别码（用 `tdict r.v <码>` 查校验内容） |
| 串查型态 / 串查程序 | 串查/参考程序 |
| 报表栏宽 / 小数位 | 报表输出参考 |
| 客制识别 | 标准/客制 |

### `tdict r.q [dzca001]`

查询**可复用开窗 (r.q)**——代码里 `CALL q_xxx()` 弹出的查寻选单定义：每条开窗 = 一段带占位符的 SQL（选取字段/表/条件区用 `<field>/<table>/<wc>` 标记）+ 外部参数（对应 arg1~9）+ 显现与回传列。读代码遇到开窗调用、想知道它查什么表、回传哪些值时用它。

开窗码示例：`q_apca001`（`q_` 开头）；`--kw 料号` 按说明搜索。

```bash
# 列出全部开窗（识别码/客制/说明/状态/每页笔数/HardCode/行业别）
tdict r.q

# 按关键字过滤（识别码或说明）
tdict r.q --kw 料号

# 指定 dzca001 的完整详情
tdict r.q q_apca001

# 指定说明语言别（默认 zh_CN）
tdict r.q q_apca001 --lang zh_TW

# JSON 输出（列表/详情）；--csv 仅列表模式
tdict r.q q_apca001 --json
tdict r.q --csv
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
tdict msg aoo-00120 --conn 正式区
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

> 通常不必用 CLI：`tdict serve` 的「数据同步」视图可选环境并显示逐表进度（见「可视化配置页」）。下列命令等价，便于脚本/自动化。

上文各查询命令读的是本地数据；`tdict db sync` 从 ERP 把它刷新到最新（覆盖表字典、校验、分类码、画面规格、开窗、消息、参数等全部内容）。

```bash
# 全量刷新（约 85 万行，视网络约 2~4 分钟）
tdict db sync

# 只刷部分（一般不必）
tdict db sync --table dzea_t

# 指定数据源环境与库文件
tdict db sync --conn 正式区 -d D:/path/to/erp_data.db
```

- 同步需能访问 ERP 服务器（内网/VPN）；采用临时库整体替换，失败不影响原库，完成后原库自动备份为 `<库>.bak`。
- 缺省数据源 = 默认环境（activeEnv，`tdict env` 查看/切换）的库；环境与库配置见「数据库连接与查询数据源」。
- 不想维护本地数据时，查询命令加 `--conn <环境名>` 直接查远程（与本地同一批数据、同一输出）。

### `tdict install skills`

把**与 `tdict.exe` 同目录的 `skills/`** 复制到**当前工作目录**（命令在哪运行就装到哪）。`skills/` 是普通的 markdown 文件、可直接编辑（不再内嵌二进制）；装完由你按所用 AI 工具改名/移动，如 Claude Code 的 `.claude/skills/`。

```bash
# 把 exe 同目录的 skills/ 复制到当前目录(生成 ./skills/)
tdict install skills

# 之后自己摆放,例如:
#   mv skills .claude/skills
```

### `tdict env [<环境名>]`

查看/设置默认 SSH 环境（写入 `config.json` 顶层 `hosts.activeEnv`）：

```bash
tdict env                     # 列出全部环境与当前默认
tdict env 正式区           # 把默认环境写入 config.json
tdict env --json
```

缺省数据源（`db sync`、`mirror`、`--conn` 直查）都取该默认环境；`--conn <环境名>` 可对单次调用覆盖。

## 可视化配置页（`tdict serve`）

在浏览器里维护 `config.json` 的 SSH 环境与数据库连接，并可直接跑源码镜像与数据同步（通常不必走 CLI）：

```bash
tdict serve                              # 默认 127.0.0.1:28670(占用自动顺延),Ctrl+C 停止
tdict serve --listen 127.0.0.1:9123
```

左侧活动栏切换四个视图：

- **环境配置**（左侧环境列表 + 右侧「SSH 服务器 / 数据库」Tab 编辑同一环境）：

| 操作 | 说明 |
|---|---|
| SSH 服务器 | 环境名、主机、端口、账号、密码、登录区域(zone)、TOPENT；可增删环境、设为默认环境 |
| 数据库 | 每个环境一对一挂载：类型(oracle/kingbase)、主机、端口、服务名(service)/库名(database)、账号列表 |
| 从服务器获取数据库配置 | 用该环境 SSH 登录服务器**只读探测**并回填连接要素：oracle 按区域加载环境后读 `ORACLE_HOME`/`sqlplus`/`TWO_TASK` 并解析 `tnsnames.ora` 得到 host/port/service；kingbase 发现运行中实例的数据目录、端口与库名。解析出内部主机名(客户端不可达)时回填 SSH 主机 |
| 账号行 ⚡ 验证 | 服务器侧以该「账号/密码」连显式目标库执行 `select 1`(只读) |
| 数据库「测试连接」 | 客户端直连测试（与 `tdict db ping` 同链路：账号取列表首项） |
| 保存 | 整块写回顶层 `hosts` 键（`query`/`mirror`/`bdldoc` 原样保留），并兼容迁移旧键 `debug` |

- **源码镜像**（等价 `tdict mirror dir` / `pull`，带实时进度）：

| 操作 | 说明 |
|---|---|
| 镜像根目录 | 查看/设置 `config.json` 顶层 `mirror.dir` |
| 环境下拉框 | 选择要拉取的环境，**默认选中默认环境**（`hosts.activeEnv`，标「（默认）」）；下方显示该环境的镜像目录与是否已拉取（有无完整基线 `.tdict-mirror.ok`，增量前提） |
| 增量更新 / 全量重建 | 后台执行 `host.MirrorPullProgress`；全量重建需二次确认（整目录替换）。同一时间只允许一个拉取任务 |
| 拉取进度 | 分阶段显示（连接 → 探测 T100 目录 → 服务器打包 → 下载并解压 → 完成）：打包阶段轮询服务器归档大小，下载阶段按压缩字节显示 `已传输/总量 (百分比)`、文件数与用时 |

- **数据同步**（等价 `tdict db sync`，带实时进度）：

| 操作 | 说明 |
|---|---|
| 目标数据库 | **可在页面上编辑并保存**（写入 `config.json` 顶层 `sync.target`）；默认是 **exe 同目录** 的 `erp_data.db`——便携版分发到任何机器都成立，**不再受宿主机 `TDICT_DB` 影响**。文件（含父目录）不存在时同步会自动创建；「恢复默认」清空自定义值。原库自动备份为 `.bak` |
| 环境下拉框 | 只列**已挂数据库**的环境，**默认选中默认环境**（`hosts.activeEnv`） |
| 开始同步 | 后台执行 `dbsync.Run`：远程逐表拉取 → 写临时库 → 建主键索引 → 原子替换本地库。需二次确认；同一时间只允许一个同步任务 |
| 同步进度 | 数据表 `第 x/y 张`、当前表与行数、累计行数、阶段（连接 → 拉取 → 建索引 → 替换 → 完成）、用时、原库备份路径与索引警告 |

- **设置**（命令行安装 / BDL 文档目录 / 运行信息）：

| 操作 | 说明 |
|---|---|
| 命令行安装 | 显示当前 `tdict` 可执行文件与所在目录，一键**加入用户 PATH**（Windows 写注册表 `HKCU\Environment\Path`，用户级、无需管理员），之后任意位置可直接运行 `tdict`；也可一键移除。已加入/未加入有明确状态 |
| 生效时机 | 新开的终端立即可用；**已打开的终端需重开**（写入后广播 `WM_SETTINGCHANGE`） |
| 非 Windows | 不自动改 PATH，页面给出等价的 `export PATH="$PATH:<目录>"` 手动命令 |
| BDL 语言文档目录 | 查看/设置 `config.json` 顶层 `bdldoc.dir`（等价 `tdict bdldoc dir <目录>`），并提示该目录在本机是否存在；接口 `GET/PUT /api/bdldoc` |
| 运行信息 | 配置文件路径、当前服务地址 |

- 配置文件取 `--config` / `TDICT_CONFIG`（缺省 `config.json`）；文件不存在时首次保存自动创建。
- **便携版发布为空配置、且不含业务数据**：打包脚本用 `config.empty.json` 生成空的 `config.json`（`hosts.sshs` 为空），首次运行用本页添加自己的环境；同时也**不打包 `erp_data.db`**（含客户表字典/schema/企业码等数据），配好环境后自行用「数据同步」或 `tdict db sync` 拉取。技能以普通目录 `skills/` 随包提供（**不内嵌二进制**，可直接编辑）。仓库中不提交 `config.json` 与 `erp_data.db`（见 `.gitignore`）。
- 服务器执行工具（sqlplus/ksql）路径自动探测，无需配置；SSH/DB 探测逻辑与 `tdict db discover` 同源（`host` 包）；镜像逻辑与 `tdict mirror pull` 同源，同步逻辑与 `tdict db sync` 同源（`dbsync` 包）。
- 该服务只做配置读写、只读探测与同步拉取，不启动任何调试会话。

## 本地源码镜像（tdict mirror）

> 通常不必用 CLI：`tdict serve` 的「源码镜像」视图可直接设置镜像根、执行增量/全量拉取并显示实时进度（见「可视化配置页」）。下列命令等价，便于脚本/自动化。

把某环境(T100 服务器)的源代码镜像到本地目录,供 **AI 用本地文件工具读码**(ls/rg/带行号读),避免每次读取都经服务器往返、也避免给 AI 服务器端查询权限。**只镜像**:各模块 `4gl`(源码)、`4fd`(前端字段描述)、`42s`(编译字符串,**只取 `zh_CN` 语言目录**——文件所在目录的最后一个目录必须是 `zh_CN`),以及 **`*.inc`(4GL include,erp/com 下任意位置)**。per(界面源)、编译产物(42m/42r/42f)、其它语言、设计器辅助目录一律不拉;目录与服务器同构(`<镜像根>/<环境名>/erp/<模块>/{4gl,4fd,42s/zh_CN,inc}/...`、`com/{lib,sub,qry,wss}/...`),客制模块 `c<mod>` 与标准模块并列,同路径下客制优先。

> **新增文件类型会自动补齐**:白名单带版本号(记录在本地基线标记 `.tdict-mirror.ok` 里)。白名单升级后,下次「增量更新」检测到版本不一致会**自动转全量**一次,把新增类型的历史文件补齐(不必手动 `--full`);进度/日志会提示本次实际为全量。

> **备份/临时文件也排除**:`*.bak`、`*.bck`、`*.bck1`、`*.old`、`*.orig`、`*.tmp`、`*.swp`、`*~` 等(服务器侧 `find` 与本地清理共用同一份后缀列表)。历史上已拉取到本地的备份文件会在下次拉取时自动清理(无需 `--full`)。

镜像根目录配置在 config.json 顶层 `mirror` 键(**必须显式设置**,不设默认避免隐式落盘):

```json
{ "hosts": { "sshs": [ ... ], "activeEnv": "正式区" },
  "query": { "source": "local" }, "mirror": { "dir": "D:\\dev\\erp-src" },
  "bdldoc": { "dir": "D:\\path\\to\\docs\\bdl" } }
```

```bash
tdict mirror dir                     # 查看当前镜像根
tdict mirror dir D:\dev\erp-src      # 设置/更换镜像根(写入 config.json;相对路径存为绝对)
tdict mirror pull                    # 默认环境(activeEnv/首条)增量更新
tdict mirror pull 正式区          # 指定环境增量更新
tdict mirror pull 正式区 --full   # 全量重建(整目录替换,删除服务器已不存在的残留)
tdict mirror path                    # 打印镜像目录绝对路径,第一行即路径(AI 前往)
tdict mirror path 正式区
```

- **增量机制**:服务器 TOP 下保留 `.tdict-mirror-<环境>.mark` 基线,默认只打包 `find -newer` 的变更文件(秒级);首次拉取与 `--full` 走全量(数百 MB~数 GB,视站点规模,服务器 gzip 打包 + SFTP 流式下载,完成后清理归档、保留 marker)。
- **打包根(TOP)解析**:T100 路径**不允许静态配置**——打包根按登录区域在服务器上动态探测获取(登录 zone → 环境脚本回读 TOP/ERP/COM),探测失败即报错(检查该环境 SSH 登录与 zone 设置)。
- AI 工作流:`tdict mirror path` 拿目录 → 本地工具读码。

## BDL 语言参考文档(tdict bdldoc)

项目内随带 **Genero BDL(4GL)语言参考文档**(markdown,已排除图片):`docs/bdl/`(约 5,200 篇,含 01 用户指南 ~ 19 版权、`llms.txt` 全文索引)。供 AI/开发者查询 BDL 语法、内置函数、fgldb 调试器命令等语言级问题——写 4GL 代码、读调试器输出遇到语言问题先查它,而不是猜。

文档存放路径记录在 config.json 顶层 `bdldoc.dir`,用命令查看/设置(**只改配置,不移动文件**):

```bash
tdict bdldoc dir                    # 输出文档目录(第一行即绝对路径;未设置时给出提示)
tdict bdldoc dir D:\path\to\docs    # 设置目录(写入 config.json;相对路径存为绝对)
```

- 默认值指向仓库内 `docs/bdl`(克隆/解压后按需用 `tdict bdldoc dir` 校正)。
- 相关章节速查:BDL 语言基础/进阶在 `08_language-basics`/`09_advanced-features`;SQL 支持在 `10_sql-support`;画面/报表在 `11_user-interface`/`12_reports`;fgldb 调试器与工具在 `13_programming-tools`。

## 数据库连接与查询数据源（config.json）

连接配置存于项目根目录 `config.json`（JSON，开发阶段明文，勿用于生产）。数据库连接按 SSH 环境**一对一挂载**：每个环境在 `hosts.sshs[]` 内嵌一个 `db`（显式 host/port + `service`(oracle) 或 `database`(kingbase) + 账号列表）；服务器侧 sqlplus/ksql 工具路径自动探测，无需配置。顶层 `query` 键记录查询命令的默认数据源（与 hosts 平级）：

```json
{
  "hosts": {
    "sshs": [
      { "name": "正式区", "host": "10.0.0.1", "port": 22, "user": "youruser",
        "password": "yourpassword", "zone": "36", "topent": "99",
        "db": {
          "type": "oracle",
          "host": "10.0.0.1", "port": 1521,
          "service": "YOUR_SERVICE",
          "accounts": [
            { "account": "your_schema", "password": "your_password" },
            { "account": "your_schema2", "password": "your_password2" }
          ]
        } }
    ],
    "activeEnv": "正式区"
  },
  "query": { "source": "local" }
}
```

- **T100 路径不静态配置**:没有 `topDir`/`moduleRoots` 配置项——源码查找根/打包根等路径一律登录该环境后按其 zone 动态探测获取(登录 zone → 环境脚本回读 TOP/ERP/COM),探测失败即报错,请检查 SSH 登录与 zone 设置。
- **账号规则**：无"主账号"，所有账号都在 `accounts` 列表，不区分默认。客户端直连（查询数据源 / `db ping` / `db sync` / 连接测试）取列表**首项**；未收录的账号按"账号=密码"惯例兜底。
- `viaSsh`（可选）：客户端不可达 DB、但 DB 对 SSH 服务器可达时，经 SSH 端口转发再直连（本地起转发端口 → 驱动连 `127.0.0.1:本地端口`）；缺省远端取连接自身 host/port。
- 支持类型：`kingbase`（人大金仓，PostgreSQL 协议，pgx）、`oracle`（go-ora 纯 Go 驱动）。
- 配置文件查找优先级：`$TDICT_CONFIG` > `--config` 参数 > 可执行文件同目录 > 当前工作目录。

### 查询数据源切换（本地 SQLite ⇄ 远程库）

`r.t/r.v/desc/scc/r.q` 五个查询命令统一走同一查询接口（`db.Source`）：本地 SQLite 镜像与远程 ERP 库查的是**同一批表(24 张字典 + 消息档)**，输出完全一致；命令层不感知数据源。解析优先级：

1. `--conn <环境名>`：本次调用远程直查该环境的库（`--conn local` 回本地）；
2. `config.json` 顶层 `query.source`：设为环境名（如 `"正式区"`）即默认远程直查该环境；
3. 缺省 `local`：本地 SQLite（`-d`/`TDICT_DB` 定位的 `erp_data.db`），与旧版行为一致。

```bash
tdict r.t dzea_t                                # 默认本地 SQLite
tdict r.t dzea_t --conn 正式区                   # 远程直查正式区(金仓 your_schema)
tdict desc oobd_t oobd002 --conn 正式区         # 远程直查正式区(oracle your_schema)
tdict r.v v_ooba002_07 --conn local             # 显式切回本地
```

远程直查使用客户端驱动直连 `db.host:port`（账号取列表首项，`--conn`/`query.source` 与此无关），不要求本地已 sync；金仓与 Oracle 均为完整支持（列表 `--kw` 过滤大小写不敏感，与本地一致）。查询是只读单条 SELECT，值经白名单/转义内联。

`db` 子命令（连接管理，与查询数据源独立；`--conn` 语义同为环境名）：

```bash
tdict db list                                  # 按环境列出挂载的库(类型/地址/账号数)
tdict db ping [--conn <环境名>]                # 验证连接可达(只读;缺省 activeEnv 环境的库)
tdict db sync [--conn <环境名>]                # 拉取该环境字典数据写入本地 SQLite(见前节)
tdict db discover --env 正式区 --type oracle [--save]   # SSH 自动发现连接要素 → 预览或写入该环境 db
```

- `db discover` 给出"服务器视角"连接要素；客户端不可达时改 `db.host` 或补 `viaSsh`。
- 注意：root 的 `--conn`（查询数据源切换）与 `db sync/ping` 组内自带的 `-c/--conn` 是两个独立 flag，值都取 SSH 环境名。

## 输出中的类型码

`tdict r.t` 输出的取值含义。

### 表类型(r.t 的「类型」列:每张表在系统里的角色)

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

### 字段数据类型(r.t 字段表的「数据类型」列常见值)

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
├── main.go              # 入口(内嵌 web/dist)
├── go.mod / go.sum      # Go module
├── config.json          # hosts(SSH 环境 + 数据库连接) + query(数据源) + mirror(镜像根) + bdldoc(BDL 文档目录)(明文,勿提交真实凭据)
├── cfgfile/             # config.json 统一读写入口(读取/校验/原子更新)
├── cli/
│   ├── root.go          # 根命令 + 全局 --json/--csv/-d/--config/--conn
│   ├── table.go         # tdict r.t(数据表字典)
│   ├── check.go / scc.go / spec.go / win.go   # tdict r.v / scc / desc / rq 查询
│   ├── msg.go / param.go                       # tdict msg / sysp / docp
│   ├── source.go        # 查询数据源解析(本地 SQLite ⇄ --conn/query.source 远程)
│   ├── env.go           # tdict env(离线查看/设置默认环境)
│   ├── mirror.go        # tdict mirror(镜像根 dir / pull / path)
│   ├── bdldoc.go        # tdict bdldoc(BDL 语言文档目录 dir)
│   ├── install.go       # tdict install skills(把 exe 同目录的 skills/ 复制到当前目录)
│   ├── db.go / db_sync.go / dbops.go   # tdict db(sync/list/ping/discover)
│   ├── serve.go         # tdict serve(可视化配置页:SSH 环境 + 数据库)
├── server/              # 配置服务实现(静态前端 + /api/config、/api/dbprobe、/api/conntest)
├── host/                # 远程能力共享层(CLI 命令共用的"服务器问路"层)
│   ├── env.go           # 环境模型 NamedSsh + LoadHosts(读 config 环境清单)
│   ├── ssh.go           # SSH 连接(Dial/PTY/SFTP/exec)
│   ├── term.go          # 终端行解析/提示符判定
│   ├── tenv.go          # 登录动态路径探测(登录 zone → TOP/ERP/COM)
│   ├── mirror.go        # 源码镜像引擎(服务器 4gl/4fd/42s-zh_CN/*.inc 打包 → SFTP → 解压)
│   └── dbprobe.go       # 服务器侧 DB 探测/账号验证
├── dbconfig/            # db 连接配置模型(每环境挂一个,host/port/service|库名+账号列表)
├── erpdb/               # ERP 连接器(只读 Query;kingbase 经 pgx / oracle 经 go-ora)
├── db/                  # SQLite 查询实现 + db.Source 统一查询接口
├── live/                # 远程数据源(live.Open;viaSsh 隧道支持)
├── sshtun/              # SSH 端口转发隧道组件
├── output/              # 表格/JSON/CSV 输出格式化
├── server/              # 本地配置服务(tdict serve):静态前端 + 配置读写/连接探测 REST
├── web/                 # 前端(React/Vite/Tailwind;SSH 与数据库配置页,产物嵌入)
├── skills/              # AI 技能文件(普通 markdown,不内嵌;tdict install skills 复制到当前目录)
│   ├── tdict.md              # 数据字典查询 Skill
│   ├── tdict-debug.md        # T100 作业调试 Skill(调试功能已迁出,供独立调试项目使用)
│   └── erp-read.md           # 阅读/分析 ERP 4GL 源码 Skill
├── docs/
│   └── bdl/             # BDL(4GL)语言参考文档(markdown,无图片;路径记于 bdldoc.dir)
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
| WebSocket | 已移除（随调试功能迁出） |
| CLI 框架 | [cobra](https://github.com/spf13/cobra) |
| Web 前端 | React + Vite + Tailwind（SSH/数据库配置页，产物 `go:embed` 嵌入） |
| 输出 | `text/tabwriter` + `encoding/json` |
