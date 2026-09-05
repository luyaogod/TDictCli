# TDict — ERP 数据字典 CLI / T100 作业调试工具

面向 **T100 / Genero(4GL) ERP** 的本地命令行工具，两条主线：

1. **数据字典查询**（给"读代码、做配置"提供上下文）：基于 T100 数据字典，查询命令（`tdict rt` 表字典 / `rv` 校验带值 / `scc` 系统分类码 / `desc` 字段规格 / `rq` 可复用开窗）统一走同一数据源接口——既可用**本地 SQLite 镜像**（`tdict db sync` 提前拉取 24 张字典表），也可经 `config.json` 的 `query.source` 或 `--conn` 切换为**远程 ERP 库直查**（金仓/Oracle，同一批表、同一输出）。
2. **AI 人机协同调试 T100 作业**（`tdict debug`）：通过 SSH 在 T100 服务器上驱动 `fglrun -d` 的 (fgldb) 文本调试协议，提供 Web 调试界面与命令行控制端，让 AI/人协同排查作业逻辑错误、跟踪变量、验证接口报文场景。

所有输出默认使用**简体中文 (zh_CN)**。

## 数据规模

| 数据表 | 说明 | 行数 |
|---|---|---|
| `dzea_t` | 资料表主档（表注册信息） | 3,882 |
| `dzeal_t` | 资料表多语言档 | 8,840 |
| `dzeb_t` | 资料表字段档（字段定义） | 148,539 |
| `dzebl_t` | 资料表字段多语言档 | 308,773 |
| `dzec_t` | 资料表索引档 | 4,430 |
| `dzed_t` | 资料表键值档（PK/FK 定义） | 5,413 |
| `dzee_t` | 表格更新包纪录档 | 1 |
| `dzef_t` | 字段参考设置档 | 18,895 |
| `dzeg_t` | 数据表多语言档 | 60 |
| `dzcd_t` 等 5 张 | 校验带值 (r.v) 定义 | 需 `tdict db sync` 同步 |
| `gzca_t` 等 4 张 | 系统分类码 (SCC) | 需 `tdict db sync` 同步 |
| `dzep_t` | 字段规格（画面设计器参考配置） | 需 `tdict db sync` 同步 |
| `dzca_t` 等 5 张 | 可复用开窗 (r.q) | 需 `tdict db sync` 同步 |

**总计 3,882 张表，148,539 个字段，498,833 行数据。**（数据可通过 `tdict db sync` 从 ERP 实时刷新；校验带值相关 5 张表 `dzcd_t dzcdl_t dzce_t dzcel_t dzch_t` 存储可复用的字段校验 SQL，同步后可用 `tdict rv` 查询；系统分类码 4 张表 `gzca_t gzcal_t gzcb_t gzcbl_t` 为各种下拉组件的数据来源，同步后可用 `tdict scc` 查询；字段规格表 `dzep_t` 记录画面设计器生成画面时的参考配置，同步后可用 `tdict desc` 查询；开窗 5 张表 `dzca_t dzcal_t dzcb_t dzcbl_t dzcc_t` 存储可复用的开窗选取 SQL，同步后可用 `tdict rq` 查询）

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
```

## 命令

> 命名说明：`rt`（数据表，对应表格设计器 r.t）、`desc`（字段规格）、`rv`（校验带值 r.v）、`rq`（开窗 r.q）采用 T100 体系术语；旧名称 `table`/`spec`/`check`/`win` 仍可作为别名使用。

### `tdict rt <表名>`

查询一张或多张表的完整字典。返回内容包含：

- **表名/表说明**（`dzea_t` + `dzeal_t` zh_CN）
- **字段**（`dzeb_t` + `dzebl_t`）：序号、字段名、字段说明、数据类型、长度、主键、必填、备注
- **键值**（`dzed_t`）：键名、类型（主键/外键/唯一）、键值字段、外键表、外键字段
- **索引**（`dzec_t`）：索引名、类型、索引字段

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

查询 **校验带值 (r.v) 定义**——ERP 中可复用的、用于校验字段的 SQL。运行时由公共函数库 `cl_chk_validate.4gl` 取出 `dzcd_t.dzcd003` 的 SQL 文本，替换占位符后动态执行（`PREPARE/EXECUTE`）。数据来源 `dzcd_t dzcdl_t dzce_t dzcel_t dzch_t`，需先执行 `tdict db sync` 同步。

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

- **单头**（`dzcd_t` + `dzcdl_t`）：型态（`dzcd005`：1=检查存在 / 2=带值 / 3=检查存在并带值）、错误讯息代号（`dzcd004`）、行业别（`dzcd006`）、状态码、说明、备注
- **SQL 指令**（`dzcd003`）：校验 SQL 原文，附标签图例（`<field>` `<table>` `<wc>` `<count>`、`arg1~arg9`、`:TODAY`、`:DEPT` 等全局变量）
- **参数**（`dzce_t` + `dzcel_t`）：顺序、参数名称、日期型态（`dzcestus`：1=年月日 / 2=年月日时分秒 / 3=毫秒 / 其他=字符串）、说明、备注
- **判断条件**（`dzch_t`）：顺序、条件 SQL（`dzch003`）、条件成立时显示的错误讯息（`dzch004`）

### `tdict scc [gzca001]`

查询 **系统分类码 (SCC)** 定义——T100 中维护各种数据/业务逻辑分类的编码，系统中各种下拉组件（`cl_set_combo_scc`，cl_set_combo.4gl:90）与取描述函数（`cl_get_scc_desc`，cl_get.4gl:1368）均从这套表取值。数据来源 `gzca_t gzcal_t gzcb_t gzcbl_t`，需先执行 `tdict db sync` 同步。

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

- **单头**（`gzca_t` + `gzcal_t`）：群组（`gzca004`）、状态（`gzcastus` Y=启用/N=停用）、名称（`gzcal003`）、说明（`gzcal004`）、说明2（`gzcal005`）
- **分类值**（`gzcb_t` + `gzcbl_t`）：值（`gzcb002`）、说明（`gzcbl004`，下拉选项显示文字）、排序（`gzcb012`，下拉按此排序）、标准/客制（`gzcb013`）
- **扩展数据列**（`gzcb003`~`gzcb011`/`gzcb014`/`gzcb015`）：按分类码语义复用的通用数据列（如 SCC 4 的 gzcb004/gzcb005 存群组编号区间、SCC 241 的 gzcb003 存格式），仅显示非空值

### `tdict desc <表名> [字段名]`

查询**字段规格表 (dzep_t)**——数据库表所有字段的规格，即画面设计器（表格设计器 adzi140、画面设计器 adzp168）生成画面时的参考配置：哪个字段用什么控件、下拉用哪个 SCC 码、格式化方式、是否必填等。运行期 `q_*.4gl` 查询程序也用 `dzep_t.dzep011` 给下拉框动态灌 SCC 值。数据来源 `dzep_t` + `dzeb_t/dzebl_t`（字段名/说明）+ `dzea_t/dzeal_t`（表说明），需先执行 `tdict db sync` 同步。

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

dzep_t 关键列语义：

| 列 | 含义 |
|---|---|
| `dzep001` / `dzep002` | 表代码 / 字段代码（复合键，关联 `dzea_t.dzea001`、`dzeb_t.dzeb002`） |
| `dzep010` | 控件类型（dzej_t GENERO_WIDGETS：01=ButtonEdit 02=CheckBox **03=ComboBox** 04=DateEdit 05=Edit 09=RadioGroup 11=SpinEdit 12=TextEdit 13=TimeEdit 34=DateTimeEdit 等） |
| `dzep011` | **SCC 码**（系统分类码，下拉/单选控件的选项数据源，仅 03/09 控件有效） |
| `dzep005` | 必填标记 Y/N |
| `dzep009` | 控件显示宽度（"整数,小数"） |
| `dzep012` | 默认值 |
| `dzep013` / `dzep014` | 最大值 / 最小值（配 `dzep025`/`dzep026` 比较符号） |
| `dzep017` / `dzep018` | 编辑时开窗 / 查询时开窗程序代码 |
| `dzep019` | r.v 校验带值代码（关联 dzcd_t） |
| `dzep020` / `dzep022` | 串查/参考程序代码 / 串查型态 |
| `dzep021` | 显示格式 format |
| `dzep023` | 字段大小写 U/L |
| `dzep027` / `dzep028` | 报表栏宽 / 报表小数位数 |
| `dzepstus` | 环境/客制识别（s=标准 c=客制） |

### `tdict rq [dzca001]`

查询**可复用开窗 (r.q) 定义**——封装一段 SQL、供需要开窗选取的字段复用的查寻窗口（设计上与校验带值同族）。由设计器 `adzi210.4gl` 维护、`adzp210.4gl` 依据 `dzca003` 的 SQL（含 `<field>/<table>/<wc>/<inwc>` 标签和 `arg1~arg9` 占位符）生成实体开窗程序 `com/qry/4gl/q_*.4gl`；调用方设置全局 `g_qryparam`（arg1~9、where、state 等）后 `CALL q_xxx()`，选取值回传 `g_qryparam.return1~9`。数据来源 `dzca_t dzcal_t dzcb_t dzcbl_t dzcc_t`，需先执行 `tdict db sync` 同步。

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

- **单头**（`dzca_t` + `dzcal_t`）：状态、SQL 指令（`dzca003` 原文，附标签图例）、每页笔数（`dzca004`）、作业串查编号（`dzca005`）、HardCode（`dzca006`，Y=跳过自动产生）、行业别（`dzca007`）、说明（`dzcal003`）、助记码（`dzcal004`）
- **参数**（`dzcb_t` + `dzcbl_t`）：顺序、参数名称（`dzcb003`，对应 argN）、日期型态（`dzcbstus`：1=年月日 / 2=年月日时分秒 / 3=毫秒 / 其他=字符串）、说明
- **显现设定**（`dzcc_t`）：显现顺序、字段编号（`dzcc003`）、表格别名（`dzcc008`）、显示控件（`dzcc004`）、是否回传（`dzcc005`，Y 的字段按顺序构成 return1~9）、大小写、显示格式（`dzcc010`）、标签缀字（`dzcc007`）

### `tdict db sync`

从 ERP 数据库拉取 24 张字典表（`dzea_t dzeal_t dzeb_t dzebl_t dzec_t dzed_t dzee_t dzef_t dzeg_t` + 校验带值 `dzcd_t dzcdl_t dzce_t dzcel_t dzch_t` + 系统分类码 `gzca_t gzcal_t gzcb_t gzcbl_t` + 字段规格 `dzep_t` + 可复用开窗 `dzca_t dzcal_t dzcb_t dzcbl_t dzcc_t`）的最新数据，写入本地 SQLite（默认 `-d/--db` 或 `TDICT_DB` 指向的 `erp_data.db`）。

同步采用**临时库 + 原子替换**：写入临时文件成功后整体替换目标库，中途失败不影响原库；替换前自动备份原库为 `<数据库>.bak`。所有列以 TEXT 存储（与现有 schema 一致），并根据 `dzed_t` 的 PK 定义重建唯一索引。

```bash
# 全量同步 24 张字典表
tdict db sync

# 仅同步部分表
tdict db sync --table dzea_t,dzeal_t

# 使用指定连接与数据库
tdict db sync --conn 恒烁dsdemo -d D:/path/to/erp_data.db
```

> 提示：同步需能访问 ERP 服务器（内网/VPN），数据量约 50 万行，受网络带宽限制通常耗时约 2 分钟。

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

`rt/rv/desc/scc/rq` 五个查询命令统一走同一查询接口（`db.Source`）：本地 SQLite 镜像与远程 ERP 库查的是**同一批 24 张字典表**，输出完全一致；命令层不感知数据源。解析优先级：

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

## 数据库 Schema

### 核心 JOIN 关系

```sql
-- 表名 → zh_CN 翻译
SELECT a.dzea001, COALESCE(al.dzeal003, a.dzea002) AS 表说明
FROM dzea_t a
LEFT JOIN dzeal_t al ON al.dzeal001 = a.dzea001 AND al.dzeal002 = 'zh_CN'

-- 字段名 → zh_CN 翻译
SELECT b.dzeb001, b.dzeb002, COALESCE(bl.dzebl003, b.dzeb003) AS 字段说明
FROM dzeb_t b
LEFT JOIN dzebl_t bl ON bl.dzebl001 = b.dzeb002 AND bl.dzebl002 = 'zh_CN'

-- 查询表主键
SELECT dzed004 FROM dzed_t WHERE dzed001 = '表名' AND dzed003 = 'P'
```

### 表类型码 (dzea004)

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

### 字段类型码 (dzeb006)

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
