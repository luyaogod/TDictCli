---
name: tdict
description: 通过 TDict CLI 查询 ERP 数据字典——从 erp_data.db 查询数据表的完整字典（表名、字段、键值、索引）；查询可复用的字段校验带值定义（r.v, dzcd001，含校验 SQL、参数、判断条件）；查询系统分类码（SCC, gzca001，含分类值列表）；查询字段规格（dzep_t：控件类型、SCC 码、画面设计器使用的格式）；查询可复用开窗定义（r.q, dzca001，含 SQL、参数、显现设定）；或通过 tdict db sync 从 ERP 数据库刷新数据。所有输出使用简体中文 (zh_CN)。
---

# TDict — ERP 数据字典 CLI

## 工具说明

TDict 是一个 Go 编写的 CLI 工具，查询本地 SQLite 数据库（`erp_data.db`），内含完整的 T100 ERP 数据字典。覆盖 **3,882 张表**、**148,539 个字段定义**及键值/索引定义。数据通过 `tdict db sync` 从 ERP 数据库实时刷新。

所有输出默认使用**简体中文 (zh_CN)** 翻译。

## 何时使用本工具

在以下场景使用 TDict：

- 查询数据表的完整字典 — **表名、字段、键值、索引**（`tdict rt <表名>`）
- 按 dzcd001 查询**校验带值 (r.v) 定义** — 存储在 `dzcd_t`/`dzce_t`/`dzch_t` 中的可复用字段校验 SQL，包括 SQL 文本、外部参数和判断条件（`tdict rv <dzcd001>`）
- 按 gzca001 查询**系统分类码 (SCC)** — 驱动所有下拉/组合框组件的分类编码，包括单头（gzca_t/gzcal_t）及其分类值列表（gzcb_t/gzcbl_t）（`tdict scc <gzca001>`）
- 查询数据表的**字段规格**（`dzep_t`）— 画面设计器生成画面时的参考配置：每个字段用什么控件、下拉用哪个 SCC 码、格式、必填标记等（`tdict desc <表名> [字段]`）
- 按 dzca001 查询**可复用开窗定义 (开窗 r.q)** — SQL（dzca003）被 adzp210 编译成实际 `q_*.4gl` 程序的查寻窗口，包括其参数（dzcb_t）和显现/回传列（dzcc_t）（`tdict rq <dzca001>`）
- 从 ERP 数据库刷新字典数据（`tdict db sync`）

**始终优先使用 TDict 而非直接读取 CSV 文件** — SQLite 数据库有索引、查询更快，且已包含 zh_CN 翻译关联。

## 命令

### `tdict rt <表名>`

查询一张或多张表的完整字典。多个表名用逗号分隔。

```bash
# 单张表
tdict rt dzea_t

# 多张表
tdict rt "dzea_t,dzeb_t,dzed_t"

# JSON 输出
tdict rt stga_t --json

# CSV 输出（字段明细平铺，含表名列）
tdict rt dzea_t --csv
```

输出包含四部分：

1. **表名/表说明** — 表名及 zh_CN 表说明（来自 `dzea_t` + `dzeal_t`）
2. **字段** — 字段定义（来自 `dzeb_t` + `dzebl_t`）：序号、字段名、字段说明(zh_CN)、数据类型、长度、主键、必填、备注
3. **键值** — 键值定义（来自 `dzed_t`）：键名、类型（主键/外键/唯一）、键值字段、外键表、外键字段
4. **索引** — 索引定义（来自 `dzec_t`）：索引名、类型、索引字段

注意：精简版 CLI 未暴露跨表字段搜索和任意 SQL 命令。对于已知所属表的字段，直接查询整张表即可。

### `tdict rv [dzcd001]`

查询**校验带值 (r.v) 定义** — 存储在 `dzcd_t` 中的可复用字段校验 SQL（SQL 文本在 `dzcd003`），运行时由 `cl_chk_validate.4gl` 在标签替换后执行（PREPARE/EXECUTE）。需要了解某字段执行什么校验、运行什么 SQL、参数/条件是什么时使用本命令。

```bash
# 列出全部校验定义（识别码/客制/说明/型态/错误讯息）
tdict rv

# 按关键字过滤（匹配识别码和 zh_CN 说明）
tdict rv --kw 料号

# 指定 dzcd001 的完整详情
tdict rv v_ooba002_07

# 繁体中文说明
tdict rv v_ooba002_07 --lang zh_TW

# JSON 输出（列表或详情）
tdict rv --json
tdict rv v_ooba002_07 --json

# CSV 输出（仅列表模式）
tdict rv --csv
```

详情按 标准/客制 变体分别输出：

1. **单头**（`dzcd_t` + `dzcdl_t`）：型态（`dzcd005`：1=检查存在 / 2=带值 / 3=检查存在并带值）、错误讯息（`dzcd004`）、行业别（`dzcd006`）、状态码、说明（`dzcdl003`）、备注（`dzcdl004`）
2. **SQL 指令**（`dzcd003`）— 带标签的原始校验 SQL（`<field>` `<table>` `<wc>` `<count>`、`arg1~arg9`、`:TODAY`、`:DEPT`/`:SITE`/`:USER` 等）；末尾打印标签图例
3. **参数**（`dzce_t` + `dzcel_t`）：顺序、参数名称（`dzce004`）、日期型态（`dzcestus`：1=年月日 / 2=年月日时分秒 / 3=毫秒 / 其他=字符串）、说明（`dzcel004`）、备注
4. **判断条件**（`dzch_t`）：顺序、条件 SQL（`dzch003`）、条件成立时的错误讯息（`dzch004`）

数据来源 `dzcd_t dzcdl_t dzce_t dzcel_t dzch_t`；本地库未同步时命令会提示先执行 `tdict db sync`。

### `tdict scc [gzca001]`

查询**系统分类码 (SCC)** — ERP 中用于数据/业务逻辑分类的编码，也是所有下拉/组合框的数据源。运行期用法：`cl_set_combo_scc`（cl_set_combo.4gl）按 `SELECT gzcb002, gzcbl004 FROM gzcb_t LEFT JOIN gzcbl_t ... WHERE gzcb001=? AND gzcastus='Y' ORDER BY gzcb012` 填充下拉；`cl_get_scc_desc`（cl_get.4gl）从 `gzcbl_t` 返回描述。

```bash
# 列出全部 SCC（分类码/群组/状态/名称/值数）
tdict scc

# 按关键字过滤（分类码或名称）
tdict scc --kw 币别

# 指定 gzca001 的完整详情
tdict scc 4

# 繁体中文说明
tdict scc 4 --lang zh_TW

# JSON 输出（列表或详情）
tdict scc --json
tdict scc 4 --json

# CSV 输出（仅列表模式）
tdict scc --csv
```

详情包含：

1. **单头**（`gzca_t` + `gzcal_t`）：群组（`gzca004`）、状态（`gzcastus` Y=启用/N=停用）、名称（`gzcal003`）、说明（`gzcal004`）、说明2（`gzcal005`）
2. **分类值**（`gzcb_t` + `gzcbl_t`，按 `gzcb012` 排序）：值（`gzcb002`）、说明（`gzcbl004`，即下拉显示文字）、排序（`gzcb012`）、标准/客制（`gzcb013`）
3. **扩展数据列**（`gzcb003`~`gzcb011`/`gzcb014`/`gzcb015`）：按 SCC 语义复用的通用数据列（如 SCC 4 的 gzcb004/gzcb005 存群组编号区间、SCC 241 的 gzcb003 存格式）；文本模式仅列非空值，JSON 输出全列

数据来源 `gzca_t gzcal_t gzcb_t gzcbl_t`；本地库未同步时命令会提示先执行 `tdict db sync`。

### `tdict desc <表名> [字段名]`

查询**字段规格**（`dzep_t`）— 画面设计器（表格设计器 adzi140、画面设计器 adzp168）根据表生成画面时的参考配置：每个字段用什么控件、下拉用哪个 SCC 码、显示格式、必填标记、默认值、开窗程序、校验范围等。运行期 `q_*.4gl` 查询程序读取 `dzep_t.dzep011` 动态给组合框灌入 SCC 值。

```bash
# 列出该表全部字段规格（序号/字段/字段名/控件/SCC码/必填/宽度/格式/默认值/校验带值）
tdict desc oobd_t

# 单字段完整详情（全部列，含 SCC 名称、取值范围、开窗）
tdict desc oobd_t oobd002

# 繁体中文说明
tdict desc oobd_t oobd002 --lang zh_TW

# JSON 输出（列表或详情）；--csv 输出全部列
tdict desc oobd_t --json
tdict desc oobd_t oobd002 --json
tdict desc oobd_t --csv
```

`dzep_t` 关键列：

1. `dzep001`/`dzep002` — 表代码 / 字段代码（复合键；关联 `dzea_t.dzea001`、`dzeb_t.dzeb002`）
2. `dzep010` — 控件类型（dzej_t GENERO_WIDGETS：01=ButtonEdit 02=CheckBox **03=ComboBox** 04=DateEdit 05=Edit 09=RadioGroup 11=SpinEdit 12=TextEdit 13=TimeEdit 34=DateTimeEdit 等）
3. **`dzep011` — SCC 码**（ComboBox/RadioGroup 的选项来源；仅控件 03/09 有效）
4. `dzep005` 必填 Y/N；`dzep009` 控件宽度（"整数,小数"）；`dzep012` 默认值；`dzep021` 显示格式；`dzep023` 大小写 U/L
5. `dzep013`/`dzep014` 最大/最小值（配 `dzep025`/`dzep026` 比较符号）；`dzep017`/`dzep018` 编辑/查询开窗程序；`dzep019` r.v 校验带值代码；`dzep020`/`dzep022` 参考程序/类型；`dzep027`/`dzep028` 报表栏宽/小数位
6. `dzepstus` — 环境/客制标记（s=标准 c=客制）

数据来源 `dzep_t` 关联 `dzeb_t`/`dzebl_t`（字段名，按 `dzeb021` 排序）和 `dzea_t`/`dzeal_t`（表说明）；本地库未同步时命令会提示先执行 `tdict db sync`。

### `tdict rq [dzca001]`

查询**可复用开窗定义 (开窗 r.q)** — 封装参数化 SQL 的查寻窗口，供字段做选取列表复用，设计与校验带值同族。由 `adzi210.4gl` 维护；生成器 `adzp210.4gl` 把 `dzca003`（标签 `<field>/<table>/<wc>/<inwc>`、占位符 `arg1~arg9`、伪变量 `:SITE/:TODAY/...`）编译成实际开窗程序 `com/qry/4gl/q_*.4gl`。运行期调用方设置全局 `g_qryparam`（arg1~9、where、state）后 `CALL q_xxx()`，选取结果回传到 `g_qryparam.return1~9`。

```bash
# 列出全部开窗（识别码/客制/说明/状态/每页笔数/HardCode/行业别）
tdict rq

# 按关键字过滤（识别码或说明）
tdict rq --kw 料号

# 指定 dzca001 的完整详情
tdict rq q_apca001

# 繁体中文说明
tdict rq q_apca001 --lang zh_TW

# JSON 输出（列表或详情）；--csv 仅列表
tdict rq q_apca001 --json
tdict rq --csv
```

详情按 标准/客制 变体分别输出：

1. **单头**（`dzca_t` + `dzcal_t`）：状态、SQL 指令（`dzca003` 原文及标签图例）、每页笔数（`dzca004`）、作业串查编号（`dzca005`）、HardCode（`dzca006`，Y=跳过自动产生）、行业别（`dzca007`）、说明（`dzcal003`）、助记码（`dzcal004`）
2. **参数**（`dzcb_t` + `dzcbl_t`）：顺序、参数名称（`dzcb003`，即 argN 序号）、日期型态（`dzcbstus`：1=年月日 / 2=年月日时分秒 / 3=毫秒 / 其他=字符串）、说明（`dzcbl004`）
3. **显现设定**（`dzcc_t`）：显现顺序、字段编号（`dzcc003`）、表格别名（`dzcc008`）、显示控件（`dzcc004`）、是否回传（`dzcc005` — Y 的字段按显现顺序构成 return1~9）、大小写、显示格式（`dzcc010`）、标签缀字（`dzcc007`）

数据来源 `dzca_t dzcal_t dzcb_t dzcbl_t dzcc_t`；本地库未同步时命令会提示先执行 `tdict db sync`。

### `tdict msg <编号>`

查询**系统消息档**（所有提示/报错消息的记录，由作业 `azzi920` 维护，运行时 `cl_err`/`cl_getmsg` 按编号取用）。当代码/日志里出现消息编号（形如 `std-00006`、`azz-00041`、`lib-xxxxx` 或负整数 SQLCODE `-100`）时，用本命令取文本与建议处理方式。

```bash
# 查默认语言(zh_CN)的消息
tdict msg std-00006

# 繁体语言行
tdict msg azz-00041 --lang zh_TW

# 多编号 / JSON
tdict msg "std-00006,azz-00041" --json

# 负整数编号(SQLCODE)需用 -- 分隔,避免被当作 flag
tdict msg -- -263

# 未同步/跨环境时远程直查
tdict msg aoo-00120 --conn 恒烁正式区
```

返回：**文本**（`gzze003`）、**建议处理**（`gzze004`）、**建议作业**（`gzze005`，名称取 `gzzal_t`）、**技术细节**（`gzze006`，程式人员用）、**类型**（`gzze007`：0警告/1错误/2资讯）、状态（`gzzestus` Y=启用）。语言策略与源系统一致：精确 `(编号, 语言)` 匹配、无自动回退 —— 默认只显示 `--lang`（zh_CN）行，该编号无此语言时命令会列出可用语言。数据来源 `gzze_t gzzal_t`（已并入 `tdict db sync` 全量表清单）。

### `tdict sysp <编号>` / `tdict docp <编号>`

查询**参数定义**（`gzsz_t` + 多语言 `gzszl_t`）。代码里看到参数编号（形如 `A-SYS-0040`、`E-CIR-0001`、`S-BAS-0028`、`D-MFG-0076`）或想了解某系统/单据参数时使用：

- `tdict sysp`：azzi990 视域 —— 系统级 (A, gzsa_t)/ 企业级 (E, ooaa_t)/ 据点级 (S, ooab_t) 参数定义与说明；
- `tdict docp`：azzi991 视域 —— 单据别参数 (D, ooac_t),附绑定单据性质清单（`gzsy_t`）。

```bash
# 查系统参数(默认 zh_CN 说明)
tdict sysp A-SYS-0040

# 繁体语言
tdict sysp S-FIN-3014 --lang zh_TW

# 单据别参数 + 单据性质绑定
tdict docp D-MFG-0101

# 多编号 / JSON
tdict docp "D-MFG-0076,D-BAS-0058" --json
```

返回定义与说明（名称/说明/型态/领域/预设值/值域/校核开窗引用等），不查运行时当前值（值在客户化值表）。数据来源 `gzsz_t gzszl_t gzsy_t`（已并入 `tdict db sync` 全量表清单）。

### `tdict install [目录]`### `tdict install [目录]`

将 TDict 的 Claude Code 技能文件（`tdict`、`erp-code-reader`、`erp-modify`）安装到目标项目的 `.claude/skills/` 目录，使 AI Agent 能自动理解和使用本工具。技能内容内嵌于二进制中。

```bash
# 安装到当前工作目录
tdict install

# 安装到指定项目
tdict install /path/to/project
```

### `tdict db sync`

从 ERP 数据库拉取 29 张表（9 张基础字典 + 校验带值 5 张 `dzcd_t dzcdl_t dzce_t dzcel_t dzch_t` + 系统分类码 4 张 `gzca_t gzcal_t gzcb_t gzcbl_t` + 字段规格表 `dzep_t` + 开窗 5 张 `dzca_t dzcal_t dzcb_t dzcbl_t dzcc_t` + 系统消息档 2 张 `gzze_t gzzal_t` + 参数定义档 3 张 `gzsz_t gzszl_t gzsy_t`）写入本地 SQLite。采用临时库 + 原子替换；原库自动备份为 `<db>.bak`。

```bash
# 全量同步 29 张表
tdict db sync

# 仅同步部分表
tdict db sync --table dzea_t,dzeal_t

# 指定连接与数据库
tdict db sync --conn 恒烁dsdemo -d D:/path/to/erp_data.db
```

需要能访问 ERP 服务器；约 50 万行数据，视网络约 2 分钟。

## 数据库 Schema

| 表 | 行数 | 说明 |
|---|---|---|
| dzea_t | 3,882 | 资料表主档（dzea001=表名, dzea002=说明, dzea003=模块, dzea004=类型） |
| dzeal_t | 8,840 | 资料表多语言档（关联 dzeal001=dzea001, dzeal002='zh_CN'） |
| dzeb_t | 148,539 | 资料表字段档（dzeb001=表, dzeb002=字段, dzeb007=类型, dzeb008=长度） |
| dzebl_t | 308,773 | 字段多语言档（关联 dzebl001=dzeb002, dzebl002='zh_CN'） |
| dzec_t | 4,430 | 索引定义（dzec001=表, dzec002=索引, dzec003=类型, dzec004=列） |
| dzed_t | 5,413 | 键值定义（dzed001=表, dzed002=键, dzed003=P/F/U, dzed004=列） |
| dzee_t | 1 | 更新包运行纪录 |
| dzef_t | 18,895 | 字段参考设置档 |
| dzeg_t | 60 | 数据表多语言档 |
| dzcd_t | — | 校验带值设计表: dzcd001=识别码, dzcd002=客制(s/c), dzcd003=校验SQL, dzcd004=错误讯息, dzcd005=型态(1/2/3) |
| dzcdl_t | — | 校验带值多语言: dzcdl001=识别码, dzcdl002=语言别, dzcdl003=说明, dzcdl004=备注 |
| dzce_t | — | 校验带值参数设计: dzce001=识别码, dzce002=顺序, dzce003=客制, dzce004=参数名称, dzcestus=日期型态 |
| dzcel_t | — | 参数多语言: dzcel001=识别码, dzcel002=顺序, dzcel003=语言别, dzcel004=说明, dzcel005=备注 |
| dzch_t | — | 校验带值判断条件: dzch001=识别码, dzch002=顺序, dzch003=条件SQL, dzch004=错误讯息, dzch005=客制 |
| gzca_t | — | 系统分类码主档: gzca001=分类码, gzca004=群组, gzcastus=状态(Y/N) |
| gzcal_t | — | 系统分类码多语言: gzcal001=分类码, gzcal002=语言别, gzcal003=名称, gzcal004=说明, gzcal005=说明2 |
| gzcb_t | — | 系统分类值档: gzcb001=分类码, gzcb002=值, gzcb012=排序, gzcb013=标准/客制(s/c), gzcb003~011/014/015=通用数据列 |
| gzcbl_t | — | 分类值多语言: gzcbl001=分类码, gzcbl002=值, gzcbl003=语言别, gzcbl004=说明, gzcbl005~007=说明2~4 |
| dzep_t | — | 字段规格表: dzep001=表, dzep002=字段, dzep005=必填, dzep009=宽度, dzep010=控件, dzep011=SCC码, dzep012=默认值, dzep013/014=最大/最小值, dzep017/018=编辑/查询开窗, dzep019=校验带值, dzep020/022=串查, dzep021=格式, dzep023=大小写, dzep027/028=报表宽/小数位, dzepstus=客制识别 |
| dzca_t | — | 开窗数据表: dzca001=识别码(q_xxx), dzca002=客制(s/c), dzca003=SQL指令(带标签), dzcastus=状态, dzca004=每页笔数, dzca005=串查编号, dzca006=HardCode(Y/N), dzca007=行业别 |
| dzcal_t | — | 开窗多语言: dzcal001=识别码, dzcal002=语言别, dzcal003=说明, dzcal004=助记码 |
| dzcb_t | — | 开窗设计参数: dzcb001=识别码, dzcb002=顺序, dzcb003=参数名称(argN), dzcb004=客制, dzcbstus=日期型态(1/2/3) |
| dzcbl_t | — | 开窗参数多语言: dzcbl001=识别码, dzcbl002=顺序, dzcbl003=语言别, dzcbl004=说明, dzcbl005=助记码 |
| dzcc_t | — | 开窗显现设定: dzcc001=识别码, dzcc002=显现顺序, dzcc003=字段编号, dzcc004=显示控件, dzcc005=是否回传(Y构成return1~9), dzcc006=大小写, dzcc007=标签缀字, dzcc008=表格别名, dzcc009=客制, dzcc010=显示格式 |

校验带值 5 张表（dzcd/dzcdl/dzce/dzcel/dzch）、系统分类码 4 张表（gzca/gzcal/gzcb/gzcbl）、dzep_t 和开窗 5 张表（dzca/dzcal/dzcb/dzcbl/dzcc）在未对含该数据的 ERP 执行 `tdict db sync` 前没有数据行。

### 关键 JOIN 模式

```sql
-- 表说明 (zh_CN)
SELECT a.dzea001, COALESCE(al.dzeal003, a.dzea002) FROM dzea_t a
LEFT JOIN dzeal_t al ON al.dzeal001 = a.dzea001 AND al.dzeal002 = 'zh_CN'

-- 字段说明 (zh_CN)
SELECT b.dzeb001, b.dzeb002, COALESCE(bl.dzebl003, b.dzeb003) FROM dzeb_t b
LEFT JOIN dzebl_t bl ON bl.dzebl001 = b.dzeb002 AND bl.dzebl002 = 'zh_CN'
```

## 全局参数

| 参数 | 说明 |
|---|---|
| `--json` | 输出缩进 JSON |
| `--csv` | 输出 CSV |
| `-d <路径>` | SQLite 数据库路径（默认: erp_data.db） |
| `--config <路径>` | ERP 连接配置文件（默认: config.json） |
