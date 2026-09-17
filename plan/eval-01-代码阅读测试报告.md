# TDict 用于「读懂 T100/Genero 4GL 源码」的实测评价

- 评测人角色：第一次拿到 `tdict` 的开发者
- 评测日期：2026-09-17
- 工具：`D:\我的项目\TDictCli\tdict.exe`（38,590,464 字节，Sep 17 09:26）
- 数据：`D:\我的项目\TDictCli\erp_data.db`（59,195,392 字节，Aug 26 13:18）
- 文档：`README.md`、`skills/tdict/SKILL.md`、`skills/erp-code-reader/SKILL.md`
- 源码树：`D:\t100_source_code`（20,218 个 `.4gl`）

**来源标注约定**：本文每条结论后标注 `[工具]`（来自 tdict 输出）、`[源码]`（来自我读 .4gl）、`[数据]`（来自我直接看 erp_data.db）、`[推测]`（我的推断，未证实）。所有命令都在 Git Bash 下执行。

---

## 0. 摘要（先给结论）

`tdict` 的 **`r.t` / `r.q` / `desc` / `msg` / `scc` / `r.v` / `sysp` / `docp` 一旦有数据，在「读代码」这件事上是**降维打击**——尤其 `r.q`（直接吐出开窗 SQL 全文）和 `desc`（吐出字段绑定的校验码与开窗程序），裸读源码花几百行才能得到的信息，一条命令就有。

但是：

1. **随包给的 `erp_data.db` 只装了 8 个命令里的 1 个**（只有 `dzea_t/dzeb_t` 这一族表字典）。`r.v`/`scc`/`desc`/`r.q`/`msg`/`sysp`/`docp` **全部报错**，而 README 与 SKILL.md 都写「本地已随工具带好数据（约 85 万行）」。这是**上手阶段唯一的、也是最大的摩擦点**。[数据][工具]
2. 唯一稳定可用的本地命令 `r.t` **没有 `--kw` 搜索**，无法「按中文说明找表」。[工具]
3. 工具**完全没有「程序/作业字典」**：输入 `aapi011`、`aapq110`、`q_adzi052`、`cl_abi` 得不到任何「这个程序做什么」。这次四个文件的「业务用途」**全部靠读 .4gl 文件头**。[工具][源码]
4. `erp-code-reader/SKILL.md` 的程序命名规则里 **`_88` 后缀是错的**：全树 `*_88.4gl` **只有 1 个**，而 `*_01.4gl` 有 484 个、`*_x01.4gl` 有 542 个、`*_g01.4gl` 有 302 个。按文档去找子程序会**找不到**。[源码]

---

## 任务 1：按说明上手

### 1.1 我做了什么、报了什么错

我**只**读了 README.md 和两份 SKILL.md，然后按 README「快速上手」的顺序试。

**第 1 步：确认文件在**

```
$ cd "D:/我的项目/TDictCli" && ls -la tdict.exe erp_data.db README.md
-rw-r--r-- 1 18526 197609    38914 Sep 17 13:34 README.md
-rw-r--r-- 1 18526 197609 59195392 Aug 26 13:18 erp_data.db
-rwxr-xr-x 1 18526 197609 38590464 Sep 17 09:26 tdict.exe
```

**第 2 步：`tdict --help`** —— 成功，命令清单很清楚。`[工具]`

**第 3 步：README 第一个例子 `tdict r.t dzea_t`** —— 成功。

```
$ ./tdict.exe r.t dzea_t
=== dzea_t ===
表说明: 数据表主档
模块: ADZ    类型: X

字段 (31):
序号  字段名        字段说明       数据类型       长度   主键  必填  备注
--  ---------  ---------  ---------  ---  --  --  --
1   dzeastus   状态码        varchar2   10
2   dzea001    table编号    varchar2   15   PK  Y
...
```

**到这里我以为通了。然后按 README「快速上手」继续试后面每一条，全部失败：**

```
$ ./tdict.exe r.v v_ooba002_07
Error: 本地库尚未包含校验定义 (dzcd_t 等表)数据。请先在有数据库的环境执行: tdict db sync

$ ./tdict.exe scc 4
Error: 本地库尚未包含系统分类码 (gzca_t 等表)数据。请先在有数据库的环境执行: tdict db sync

$ ./tdict.exe desc oobd_t
Error: 本地库尚未包含字段规格 (dzep_t)数据。请先在有数据库的环境执行: tdict db sync

$ ./tdict.exe r.q q_apca001
Error: 本地库尚未包含开窗定义 (dzca_t 等表)数据。请先在有数据库的环境执行: tdict db sync

$ ./tdict.exe msg std-00006
Error: 本地库尚未包含消息档 (gzze_t/gzzal_t)。请先执行 tdict db sync,或用 --conn <环境名> 远程直查

$ ./tdict.exe sysp A-SYS-0040
Error: 本地库尚未包含参数档 (gzsz_t/gzszl_t)。请先执行 tdict db sync,或用 --conn <环境名> 远程直查

$ ./tdict.exe docp D-MFG-0076
Error: 本地库尚未包含参数档 (gzsz_t/gzszl_t)。请先执行 tdict db sync,或用 --conn <环境名> 远程直查
```

**8 个查询命令，7 个不可用。** 而 README 的「能查到的数据（规模）」表明确写着这些内容"查询读的是本地数据"，SKILL.md 更写着「本地已随工具带好数据（约 3,882 张表的字段定义、148,539 个字段，以及校验/分类码/开窗/消息/参数等元数据，合计约 85 万行）」。

**第 4 步：直接查 SQLite，确认到底装了什么** ——

```
$ python -c "... select name from sqlite_master ..."
tables: ['dzea_t','dzeal_t','dzeb_t','dzebl_t','dzec_t','dzed_t','dzee_t','dzef_t','dzeg_t']
  dzea_t: 3882     # 表登记
  dzeal_t: 8840
  dzeb_t: 148539   # 字段定义  <- 与 SKILL.md 的 148,539 完全一致
  dzebl_t: 308773
  dzec_t: 4430
  dzed_t: 5413
  dzee_t: 1
  dzef_t: 18895
  dzeg_t: 60
```

`[数据]` 结论：随包的 `erp_data.db` 只含 **r.t 一族**（表登记 + 字段 + 键值/索引），合计约 49.5 万行；**dzcd_t(r.v) / gzca_t(scc) / dzep_t(desc) / dzca_t(r.q) / gzze_t(msg) / gzsz_t(sysp,docp) 一张都没有**。README 的「约 85 万行」是把还没同步的内容也算进去了。

**第 5 步：找出路。** SKILL.md 提过一句 `--conn <环境名>` 可以远程直查。先看有哪些环境：

```
$ ./tdict.exe env
当前默认(activeEnv): 主机正式区
环境 (3):
   * 主机正式区        172.16.1.109:22      zone=36   TOPENT=DSCNJ  db=oracle 172.16.1.109:1521/t35prd
     恒烁正式区        172.19.93.55:22      zone=1    TOPENT=9999   db=kingbase 172.19.93.55:54321/topprd
     恒烁测试区        172.19.93.58:22      zone=2    TOPENT=9999   db=kingbase 172.19.93.58:54321/topprd
```

**第 6 步：用 `--conn 主机正式区` 重试 7 个失败的命令 —— 全部成功。**

```
$ ./tdict.exe msg std-00006 --conn 主机正式区
=== std-00006 (zh_CN) ===
类型:     错误 (1)
状态:     启用
文本:     输入的资料已存在
建议处理: 请重新审核您所需要输入的资料
```

`恒烁测试区/恒烁正式区` 连不通（`dial error: timeout`），`主机正式区`（Oracle）可以。`[工具]`

**第 7 步：验证 README 推荐的修复路径 `db sync` 是否真的有用**（写到临时库，不动随包的 `erp_data.db`）：

```
$ ./tdict.exe db sync --conn 主机正式区 -d D:/t100_source_code/.tmp_eval/eval.db --table dzca_t,dzcal_t
正在从 ERP 拉取字典数据 (环境: 主机正式区) -> D:/t100_source_code/.tmp_eval/eval.db
同步完成: 2 张表, 共 16201 行
数据库已更新: D:/t100_source_code/.tmp_eval/eval.db

$ ./tdict.exe r.q q_adzi052 -d .../eval.db     # 立刻就能查了
=== q_adzi052 ===
...
```

`db sync` 是有效的，报错信息给的指引也是对的。（临时库已删除，源码树保持只读。）

### 1.2 上手过程的评价

- 工具**本身**装好了：exe 和 db 同目录，双击即用，无需配置，`--help` 清楚。**这一步零摩擦**。
- **真正的摩擦在上手第 5 分钟**：照着 README「快速上手」抄，第 4 条命令开始全红。文档承诺的「本地离线可用」在这份数据上不成立，而文档里没有任何地方说「先确认本地库是否完整」。
- 文档对「本地 vs 远程」的写法会**把人带偏**：README 把 `--conn` 描述成一个「可选优化」（"不想维护本地数据时…"），而不是「本地库不全时的救命手段」。
- 有一个隐性的坑：`-d/--db` 的默认值是**相对路径** `erp_data.db`，查找顺序是 `$TDICT_DB > -d > exe 同目录 > 当前工作目录`。我在 `D:/我的项目/TDictCli` 下用 `./tdict.exe` 恰好命中 exe 同目录；如果我在别处用绝对路径调 exe，`-v` 显示的解析结果就未必如预期。`-v` 能打印解析到的路径，这点是好的。
- 另一处小坑（非工具问题）：`which tdict` 指向 `D:\APPS\tdict-portable\tdict`（另一份旧便携版），**不是**本次要评的这份。文档没有提醒"PATH 上可能已有一个旧 tdict"。

---

## 任务 2：用工具读懂四个文件

> 说明：因为本地库只有 `r.t` 可用，以下凡 `[工具]` 标注者，除 `r.t` 外都是 `--conn 主机正式区` 远程查询。

### 2.1 `erp/aap/4gl/aapq110_01.4gl`（224 行）

**a. 用途** —— `[源码]` 文件头第 10 行直接写明：

```
#+ Description: 供應商對帳單明細查詢報表列印
```

它是报表作业 `aapq110`（`供應商對帳單明細查詢`，见 `erp/aap/4gl/aapq110.4gl` 头部）的**查询条件输入子程序**：开一个输入窗让用户选"类型"，把值通过模块变量 `g_apbb_m.l_type` 带回主程序。整个文件就是一个 `INPUT BY NAME` 对话，没有业务逻辑。

**工具在这里帮不上忙**：`tdict` 没有「输入程序名 → 得到用途」的命令。`[工具]`（`tdict --help` 列出的 14 个命令里没有程序/作业维度）。

**b. 访问的表** —— **无**。`[源码]` 全文没有任何 `_t` 表 token、没有任何 SQL 语句（`grep -oE '\b[a-z][a-z0-9]{1,6}_t\b'` 只匹配到 `type_t`，那是 4GL 的类型记录，不是表）。它只 `DISPLAY BY NAME g_apbb_m.l_type`。

> 注意：`tdict` 的能力是「按表名查表」，**没有反向的「按程序查表」**。所以「这个文件访问哪些表」这个问题，工具**结构上就回答不了**，必须 grep 源码。

它的业务表在父程序里（供参考）：

```
$ ./tdict.exe r.t "apbb_t,apba_t,apeb_t,apea_t,pmaa_t" --conn 主机正式区
```

`[源码]`+`[工具]`：`aapq110.4gl` 引用 `apbb_t`(36 次)、`apba_t`(11)、`pmaa_t`(9)、`apea_t/apeb_t`、`isai_t`、`ooef_t` 等。`apbb_t` 的中文名是 **进项发票对账主档** `[工具]`，与文件名 "供應商對帳單"（供应商对账单）语义吻合。

**c. 校验 / 分类码 / 消息 / 开窗 / 参数** ——

- `[源码]` 消息机制：`WHENEVER ERROR CALL cl_err_msg_log`（第 94 行）——本文件内没有出现任何具体消息编号。
- `[源码]` 开窗：无。第 156 行注释 `#此ctrlp無內容` 说明 `l_type` 字段没有绑定校验带值（ctrlp 空）。
- `[源码]` 参数：使用全局 `g_qryparam.state = "i"`。
- `[源码]` 画面：`OPEN WINDOW w_aapq110_01 WITH FORM cl_ap_formpath("aap","aapq110_01")` —— 画面档路径由 `cl_ap_formpath` 动态解析（per/4fd 档），这是 T100 的"画面档"机制。

**d. 关键机制** ——

- `[源码]` **Genero section 样板代码**：整个文件由 `{<section id="aapq110_01.xxx" type="s">}` 切块，每块带 `#應用 a00/a01/a02/a03/a04 樣板自動產生(Version:N)` 与 `#add-point:...` 客制插槽。**读这类文件时 80% 的行是样板噪音**，有效信息只在 add-point 里和 description 里。
- `[源码]` `DIALOG ATTRIBUTES(UNBUFFERED, FIELD ORDER FORM)` + `INPUT BY NAME ... ATTRIBUTE(WITHOUT DEFAULTS)`
- `[源码]` `&include "common_action.4gl"` —— 宏包含，提供 accept/cancel/close/exit 共用 action（**这个被 include 的文件不在本目录，读代码时得自己去找**）
- `[源码]` `GLOBALS "../../cfg/top_global.inc"`（全局变量档）、`IMPORT FGL lib_cl_dlg`、`SCHEMA ds`
- `[源码]` 无事务、无加锁、无报表（本文件内）

---

### 2.2 `com/qry/4gl/q_adzi052.4gl`（1,274 行）

**a. 用途** —— **这是本次评测最有说服力的一个案例。**

`[源码]` 里这个文件**读不出业务用途**：文件头 `# Descriptions...:` 是空的，函数名全是 `q_adzi052_init / _prep_result_set / _pagedata_fill / _set_pagebutton / _get_multiret / _rsfilter` 这类**样板化的开窗引擎函数**（头部注释 `#應用 qry 樣板自動產生(Version:22)`）。1,274 行里绝大部分是通用开窗逻辑：分页、多选、CONSTRUCT 条件、结果过滤、Excel 导出。

`[工具]` 一条命令就给出了答案：

```
$ ./tdict.exe r.q q_adzi052 --conn 主机正式区
=== q_adzi052 ===

[标准]
状态:
每页笔数:
串查编号:
HardCode: N
行业别:   sd
说明:     未解开Section程序清单
助记码:

SQL 指令 (dzca003):
  SELECT DISTINCT <field>dzaf001,dzaf002,dzaf010,dzaf005</field>
    FROM <table>(SELECT dzaf001, MAX(dzaf002) dzaf002, dzaf010,dzaf005
            FROM ds.dzaf_t af0
           WHERE af0.dzaf010 =
                 (SELECT MIN(dzaf010)
                    FROM ds.dzaf_t af1
                   WHERE af1.dzaf001 = af0.dzaf001)
             AND af0.dzaf005 NOT IN ('T', 'MT', 'MG', 'MV')
             AND NOT EXISTS (SELECT 1 FROM ds.dzsf_t WHERE dzsf001 = af0.dzaf001)
             AND dzaf005 NOT IN ('X','G','F')
           GROUP By dzaf001, dzaf005,dzaf010)</table>
   WHERE <wc>1=1</wc>
  ORDER BY dzaf001

显现设定 (dzcc_t, 4):
显现顺序  字段编号     表格别名  显示控件           是否回传  大小写  显示格式  标签缀字
----  -------  ----  -------------  ----  ---  ----  ----
1     dzaf001        Edit (05)      Y
2     dzaf002        Edit (05)      Y
3     dzaf010        CheckBox (02)  N
4     dzaf005        Edit (05)      N

标签说明 (开窗 SQL 中的占位符):
  <field>...</field>   选取/显示字段区 ...
```

**结论**：`q_adzi052` 是「**未解开Section程序清单**」开窗——列出"建构编号"供选择，供"规格/版本管理"使用。

**这是工具对裸读的压倒性胜利**：`说明` + `SQL 全文` + `显现/回传设定` + `占位符图例`，一次调用全给。裸读 1,274 行只能看出"这是个开窗引擎"。

**b. 访问的表** ——

`[源码]` 我一开始用正则从 `FROM/INTO/UPDATE/JOIN` 里抽表名，得到 `dzca_t dzcal_t dzcc_t dzco_t dzeb_t dzep_t gzcc_t` —— **漏了真正的主角** `dzaf_t` 和 `dzsf_t`，因为它们是 schema 限定写法 `ds.dzaf_t`，且出现在**拼接的多行字符串**里。补一次全量 token 扫描才拿到完整清单：

```
$ grep -oE "\b[a-z][a-z0-9]{1,5}_t\b" com/qry/4gl/q_adzi052.4gl | sort | uniq -c | sort -rn
     32 type_t       <- 4GL 类型，非表
     14 dzaf_t
     12 dzco_t
      8 dzca_t
      6 dzeb_t
      5 dzsf_t
      4 gzcc_t
      3 ooaa_t
      2 dzep_t
      2 dzcal_t
      1 dzcc_t
```

`[工具]` 一张命令给出全部中文含义：

```
$ ./tdict.exe r.t "dzaf_t,dzsf_t,dzca_t,dzcal_t,dzcc_t,dzco_t,dzeb_t,dzep_t,gzcc_t,ooaa_t" --conn 主机正式区
dzaf_t   | 规格与程序版本历程对应表        (模块 ADZ)
dzsf_t   | Section异动授权数据档
dzca_t   | 开窗数据表                       <- r.q 的定义档本人
dzcal_t  | 开窗数据表多语言档
dzcc_t   | 开窗显现设置表格
dzco_t   | 开窗模糊查询字段设置
dzeb_t   | 数据表字段档
dzep_t   | 字段规格设计数据档
gzcc_t   | 状态码分类
ooaa_t   | 企业层级参数表
```

关键字段 `[工具]`（`r.t dzaf_t` / `r.t dzsf_t`）：

- `dzaf_t`（规格与程序版本历程对应表）：`dzaf001` 建构编号(PK)、`dzaf002` 建构版号(PK)、`dzaf005` 建构类型(PK)、`dzaf010` 识别标示(PK)、`dzaf003` 规格版号、`dzaf006` 模块、`dzaf009` 客户编号
- `dzsf_t`（Section异动授权数据档）：`dzsf001` 程序编号(PK)、`dzsf002` 异动原因、`dzsf003` 申请解开者

**读出来的业务逻辑**：列出「每个程序(建构编号)的最新版号」，但**排除** `dzaf005 IN ('T','MT','MG','MV','X','G','F')`（这些建构类型）以及**已经送过 Section 异动授权申请的程序**（`NOT EXISTS dzsf_t`）——即"还没被解开 Section 的程序"。

**c. 校验 / 分类码 / 消息 / 开窗 / 参数** ——

`[源码]` 消息编号：
```
adz-00926   adz-01361
```

`[工具]`（注意：本地库没有消息档，必须带 `--conn`）：
```
$ ./tdict.exe msg "adz-00926,adz-01361" --conn 主机正式区
=== adz-00926 (zh_CN) ===
类型:     错误 (1)
文本:     开窗%1设定有异常。
建议处理: 请联络系统管理员检查r.q SQL指令设定
技术细节 (gzze006):
  验证sql错误:%2
=== adz-01361 (zh_CN) ===
类型:     错误 (1)
文本:     警告:找不到开窗%1模糊查询字段，请确认开窗模糊查询设置数据。
```

`[源码]` 系统参数：
```
$ grep -ohE "\"[A-Z]-[A-Z]{3}-[0-9]{4}\"" com/qry/4gl/q_adzi052.4gl
"A-SYS-0073" "E-SYS-0002" "E-SYS-0003"
```

`[工具]`：
```
$ ./tdict.exe sysp A-SYS-0073 --conn 主机正式区
名称:     启用依录入值过滤开窗结果功能
说明:     开窗结果会依据原字段录入的值进一步筛选
型态:     Y/N (1)   预设值: N

$ ./tdict.exe sysp E-SYS-0002 --conn 主机正式区
名称:     开窗每页显现数据笔数
说明:     若r.q没有设置开窗每页显现数据笔数,则依参数的设置显示
型态:     整数选项 (2)   预设值: 200   值域: >=10

$ ./tdict.exe sysp E-SYS-0003 --conn 主机正式区
名称:     开窗选择笔数数据上限
预设值:   1000   值域: >=200
```

**这三条参数正好解释了源码里 `gi_page_count` / `g_sel_limit` 两个变量的取值来源**——工具在这里做的是"跨文件语义补齐"，裸读源码只能看到"有个变量叫 page_count"，看不到"默认 200、来自企业级参数 E-SYS-0002"。

`[源码]` 系统分类码：本文件没有硬编码 SCC，它是**动态**从 `dzeb_t → gzcc_t` 取（第 1037-1041 行 `SELECT dzeb001 FROM dzeb_t WHERE dzeb002 = l_col_id` → `SELECT DISTINCT gzcc003 FROM gzcc_t WHERE gzcc001 = l_dzeb001`），或用 `cl_set_combo_scc_part` / `cl_set_combo_scc` 设置（第 1051/1060 行）。

`[工具]` 本文件用到 `ooaa_t`（企业层级参数表）；`dzca_t` 的 `dzca004` 即"每页笔数"，取值优先级源码注释写着 `ds.dzca_t > dsdemo.ooaa_t`（第 177 行）`[源码]`——这个"覆盖链"工具没有说明，只能靠读注释。

**d. 关键机制** ——

`[源码]`：

- **元数据驱动的动态 SQL**：程序自身不含业务 SQL，业务 SQL 存在 `dzca_t.dzca003`，运行时拼装（`q_adzi052_prep_result_set`）
- **分页**：`ROWNUM AS RANK` 子查询 + `g_page_idx/g_page_last`（第 622-638 行）
- **多选**：`INPUT ARRAY gr_qry FROM s_qry.*` + `gr_qry_sel` 暂存已勾选行（第 344 行、`_qry_check`）
- **CONSTRUCT 查询条件**：`gi_cons_where`（第 806 行 `q_adzi052_sqlwhere`）
- **结果过滤 rsfilter**：`q_adzi052_rsfilter`（第 1111 行）
- **Excel 导出**：`cl_export_to_excel`
- **SCC 下拉**：`cl_set_combo_scc_part` / `cl_set_combo_scc`
- **画面档**：`OPEN WINDOW` + `ui.Dialog`
- 无事务、无加锁

---

### 2.3 `com/lib/4gl/cl_abi.4gl`（2,073 行，共用函数库）

**a. 用途** —— `[源码]` 文件头：

```
#+ Description: ABI Library
#+ Modifier...: No.211101-00032#3 22/03/08 14910 1.Copy From cl_fr
```

即 **T100 的 ABI 报表函数库**（从报表库 `cl_fr` 复制改造而来）。`[源码]` 公开函数：

| 函数 | 我的理解 |
|---|---|
| `cl_abi_view()` | 选择 ABI 报表样板（`gziu_t`）并调用报表主机出报表 |
| `cl_abi_view_logfile()` | 看报表日志 |
| `cl_abi_direct_print()` | 直接送印（读 `gzgn_t` 报表直接送印设置档） |
| `cl_abi_maillist()` | 取邮件名单 |
| `cl_abi_check_reptable()` / `cl_abi_trun_reptable()` | 报表资料临时表（`dsrept` schema）的检查/清理 |
| `cl_abi_get_cptdir()` | 报表元件路径 |
| `cl_abi_get_gzgh011()` / `cl_abi_update_gzpe009()` | XtraGrid 中继档 / 排程信件清单维护 |
| `cl_abi_data_rowcnt()` | 表行数 |
| `cl_abi_get_sid_timestamp()` / `cl_abi_chk_env_exist()` / `cl_abi_get_env_global()` | 环境/时间戳工具 |

**工具帮不上忙**：`tdict` 没有任何"函数库/程序说明"命令。这个问题必须读源码。

**b. 访问的表** ——

`[源码]` 全量 token 扫描（同样有明显的 `type_t` 噪音）：
```
36 gziu_t   13 gzpe_t   11 gzgn_t    7 oojn_t     6 gziul_t    6 gzgu_t
 6 gzgh_t    5 gzze_t    5 gzgj_t    4 ooef_t     4 gzzy_t     4 gzgv_t
 3 gztd_t    3 gzou_t    3 gzgdl_t   3 gzde_t     2 ooek_t     2 ooag_t
 2 gzzj_t    2 gzyal_t   2 gzgm_t    2 glap_t     1 ooefl_t    1 gzxa_t
 1 gzgq_t    1 gzgf_t    1 gzgd_t    1 gzda_t     1 gzcbl_t    1 glaa_t
```

`[工具]` 一条命令拿到全部中文名：
```
$ ./tdict.exe r.t "gziu_t,gziul_t,gzpe_t,gzgn_t,gzgh_t,gzgu_t,oojn_t,ooek_t,gzgm_t,gzgj_t,gzgdl_t,gztd_t,gzde_t,glap_t" --conn 主机正式区
gziu_t  | ABI样板数据主档          (AZZ)
gziul_t | ABI样板数据主档多语言档
gzpe_t  | 排程细项信件清单          (B)
gzgn_t  | 报表直接送印设置档
gzgh_t  | XtraGrid中继档
gzgu_t  | 3.0首页可视化报表设置档(abi)   (M)
oojn_t  | 报表邮件默认收件人单身档
ooek_t  | 报表Logo纪录档
gzgm_t  | 报表元件表头设置档
gzgj_t  | 报表表头设置单头档
gzgdl_t | 报表样板说明多语言档(GR+XtraGrid)
gztd_t  | 字段属性定义档
gzde_t  | 子程序及应用元件基本数据表
glap_t  | 凭证单头档
```

`[源码]` 字段级注释（源码里就在 `LIKE gziu_t.gziu000 #報表樣板ID` 这种位置）：
`gziu000` 报表样板ID、`gziu002` 样板代号、`gziu003` 客制否、`gziu004` 角色、`gziu005` 用户、`gziu007` 列印签核。`[工具]` `r.t gziu_t` 能给出同样的信息（中文名 + 类型），但**中文名更粗**（"ABI样板数据主档"），源码里的行内注释更细。

**c. 校验 / 分类码 / 消息 / 开窗 / 参数** ——

`[源码]` 参数：`"A-SYS-0071" "A-SYS-0081" "A-SYS-0094" "A-SYS-0095" "A-SYS-0096" "A-SYS-0097"`

`[工具]`：
```
$ ./tdict.exe sysp A-SYS-0095 --conn 主机正式区
名称:     可视化报表2主机网址，例「http://报表主机IP:端口号」
型态:     字符或SCC (4)   领域: SYS
```

`[源码]` 消息编号（本文件引用最多，20+ 条）：`azz-00665 azz-00666 azz-11183 lib-00073 lib-00082 lib-00112 lib-00113 lib-00115 lib-00117 lib-00119 lib-00142 lib-00154 lib-00172 lib-00177 lib-00185 lib-00186 lib-00188 lib-00211 lib-00236 lib-00315 ...`

`[工具]`（抽样）：
```
$ ./tdict.exe msg "azz-00665,azz-00666,lib-00073,lib-00082,lib-00112" --conn 主机正式区
azz-00665 | 错误 | 所有用户
azz-00666 | 错误 | 所有角色
lib-00073 | 错误 | 无报表资料生成...
lib-00082 | 错误 | 程式闲置时间过久,请重新执行程式.
lib-00112 | 资讯 | 系统通知
```

**这套消息在一眼扫过时基本没用**——`所有用户` / `所有角色` 光看文本不知道它在什么场景报。如果只看 `msg` 文本，反而可能误判。要真理解得回源码看它在哪个分支被 CALL。**这是工具的"信息孤岛"问题：它给你文本，但不给你上下文。**

`[源码]` SCC：本文件没有硬编码 SCC 码。

**d. 关键机制** ——

`[源码]`：
- **报表资料临时库**：`IF cl_null(g_report) THEN LET g_report = "dsrept"` / `g_rep_db = "dsrept."`（`cl_abi_check_reptable`，第 1677 行）——报表执行时把结果写进 `dsrept` schema 的临时表
- **HTTP 直连报表主机**：`DEFINE req com.HTTPRequest / resp com.HTTPResponse`（`cl_abi_view`）
- **排程整合**：`GLOBALS "../../cfg/top_schedule.inc"`
- **邮件**：`cl_abi_maillist` / `cl_rpt_get_default_maillist` / `cl_gr_nodata_sendmail`
- **直接送印**：按 Client IP 优先序匹配 `gzgn_t` 设置（`l_sort`，第 1207 行）
- **日志**：`cl_xg_view_logfile` / `cl_abi_view_logfile`
- **多语言**：`gziul_t` / `gzgdl_t`
- 无事务、无加锁（这是报表库）

---

### 2.4 `erp/aap/4gl/aapi011.4gl`（6,592 行，大作业）

**a. 用途** —— `[源码]` 文件头：

```
#+ Description: 應付帳款類別依帳套設定科目作業
#+ Standard Version.: SD版次:0031(2025-09-03) PR版次:0031(2025-09-03)
```

即**应付账款（AP）模块的「依账套设定会计科目」基本资料维护作业**（`i` 类）。功能：按「账套(glabld) + 账务类型分类码(glab002)」维护对应会计科目。

`[工具]` 完全答不了这个——没有程序字典。`[工具]` 一个**间接**线索：查 `aap-00001` 的"建议作业"字段指出同族作业 `aapi010 (应付账务类型维护作业)`，暗示了模块语境，但不给 aapi011 本身。

**b. 访问的表** ——

`[源码]` 全量 token 扫描：
```
154 glab_t   94 type_t    7 glaa_t    6 gzcb_t    6 glad_t    5 ooag_t
  4 oocq_t    3 oocql_t   3 glac_t    2 ooefl_t   2 isao_t    2 gzcbl_t
  2 glaal_t   1 faacl_t   1 dzzc_t    1 apcb_t    1 apca_t
```
`[源码]` 字符串 SQL 里出现的：`FROM glaa_t / glab_t / glad_t / gzcb_t / isao_t / oocq_t`、`INTO glab_t`、`UPDATE glab_t`。

`[工具]`：
```
$ ./tdict.exe r.t "glab_t,glaa_t,glac_t,glad_t,gzcb_t,gzcbl_t,ooag_t,oocq_t,oocql_t,isao_t,glaal_t,faacl_t,apca_t,apcb_t,dzzc_t" --conn 主机正式区
glab_t  | 账套应用会计科目设置档     (AGL, 类型 D)
glaa_t  | 账别数据档
glac_t  | 会计科目数据档
glad_t  | 账别科目管理设置档
gzcb_t  | 系统分类值档
gzcbl_t | 系统分类值多语言档
ooag_t  | 员工数据档
oocq_t  | 应用分类码档(ACC)
oocql_t | 应用分类码多语言档
isao_t  | 营运据点开立发票数据设置档
glaal_t | 账别数据档多语言档
faacl_t | 固定资产主要类型档多语言档
apca_t  | 应付凭单单头
apcb_t  | 应付凭单单身
dzzc_t  | 简易画面数据设置
```

主档关键字段 `[工具]`（`r.t glab_t`，58 个字段）——**这个输出含金量极高，连 SCC 引用都写在"备注"列里**：

```
=== glab_t 账套应用会计科目设置档  模块 AGL  类型 D
1  glabent   企业编号        number    5    PK Y
2  glabld    账套            varchar2  5    PK Y
3  glab001   设置类型        varchar2  10   PK Y  SCC.8301科目應用設定分類
4  glab002   分类码          varchar2  10   PK Y  SCC系統分類碼；ACC應用分類碼
5  glab003   分类码值        varchar2  10   PK Y  SCC系統分類碼值；參考scc_8301設定分類之備註說明。
6  glab004   科目参照表编号   varchar2  10
7  glab005   会计科目编号一    varchar2  24       主要的預設會計科目。
8  glab006   会计科目编号二    varchar2  24       依部門性質區分科目者...
12 glab010   其他设置值      varchar2  100      非會計科目型態之設定值者，例如 比率、金額、對應碼別等。
13 glab011   科目汇总方式    varchar2  10       SCC_8315科目彙整方式
14 glabud001 自定义字段(文本)001 ... (ud001~ud030)
```

**"备注"列直接告诉你 SCC 码编号 8301 / 8315**——这是工具最有价值的一处设计。

**c. 校验 / 分类码 / 消息 / 开窗 / 参数** ——

`[源码]` SCC 硬编码：
```
396:      CALL cl_set_combo_scc('glab011','8315')
409:   CALL cl_set_combo_scc('glab011_2','8315')
410:   CALL cl_set_combo_scc('glab003','8504')
411:   CALL cl_set_combo_scc('glab003_2','8504')
```

`[工具]`：
```
$ ./tdict.exe scc 8315 --conn 主机正式区
名称: 科目汇整方式
分类值 (3):
值  说明        排序    标准/客制
1  依明细项成立科目  0010  标准
2  同会计科目合并   0040  标准

$ ./tdict.exe scc 8504 --conn 主机正式区
名称:   应付账务类型科目设置项目
说明2:  帳務類型科目設定，應設定項目。
        :代碼意義: 0* 借方科目  2* 貸方科目  1*中間性科目
        應用欄位一: 區別aapt400作業下的區別
        應用欄位二: 預設科目借貸方( D借 C貸 )
分类值 (29):
...
```

**工具在这里的表现非常好**：把 `'8315'` 这个魔法数字变成"科目汇整方式"，并且把 SCC 的自定义扩展列语义也讲清楚了。

`[源码]` 消息：`aap-00001 aap-00002 adz-01032 lib-00419 std-00003 std-00004 std-00006 std-00009`

`[工具]`：
```
$ ./tdict.exe msg "aap-00001,std-00006,adz-01032" --conn 主机正式区
=== aap-00001 ===
类型:     错误 (1)
文本:     输入的数据不存在于应用分类码档中！
建议处理: 请至[应付账务类型维护作业aapi010]维护后，重新输入！
建议作业: aapi010 (应付账务类型维护作业)      <- 直接给出"该去哪个作业处理"
=== std-00006 === 输入的资料已存在
=== adz-01032 === 无上一笔资料可供复制
```

**`建议作业` 这一列是宝藏**——它是工具里唯一能把"程序编号 ↔ 作业中文名"连起来的入口。

`[工具]` 字段级校验与开窗（`desc`）：
```
$ ./tdict.exe desc glab_t glab002 --conn 主机正式区
=== glab_t (账套应用会计科目设置档) ===
字段:     glab002 (分类码)
控件:     Edit (05)     必填: Y    大小写: upper
开窗:     编辑 q_faac001 / 查询 q_faac001
校验带值: v_faac001
客制识别: Y

注: dzep011 (SCC码) 为下拉/单选控件的选项数据源, 仅控件为 ComboBox (03)/RadioGroup (09) 时有效;
    各列含义可对照维护作业 adzi150 (字段规格) 与画面设计器 adzp168。
```

**这是我给 `desc` 打高分的实证**：源码里 6,592 行，你看到 `glab002` 只能猜；工具直接告诉你它绑了开窗 `q_faac001`、校验 `v_faac001`。而且输出末尾还主动给了"去哪看全部含义"。

`[源码]` 开窗调用：`q_oocq002`、`q_glac002_6`、`q_glac002`、`q_glaald_03`、`q_dzeb002_10`、`q_authorised_ld`
`[源码]` 参数：文件内无 `X-XXX-0000` 字面量（参数走运行时查询）。

**d. 关键机制** ——

`[源码]` 关键字计数：
```
53 cl_err                 47 s_transaction_end       13 s_transaction_begin
13 cl_set_act_visible     10 s_hint_show_set_comments 10 s_desc_get_account_desc
 7 cl_navigator_setting    4 s_axrt                   4 cl_export_to_excel
 4 cl_set_comp_entry       4 cl_set_combo_scc         3 FOR UPDATE
 3 q_glac*                 3 cl_show_fld_cont         3 cl_log_act
 1 cl_msgcentre_notify     1 ui.Dialog                1 OPEN WINDOW
```

- **事务**：`s_transaction_begin` / `s_transaction_end`（47 次 end）——T100 标准的交易包
- **加锁**：`DECLARE aapi011_cl CURSOR FROM g_forupd_sql  # LOCK CURSOR`（第 294、3027 行）、`FOR UPDATE` 3 处 → 乐观锁/行锁
- **账套权限控管**：`q_authorised_ld`、`s_desc_get_ld_desc`（与头部注释 `#160812-00027#5 全面盤點應付程式帳套權限控管` 对应）
- **复制**：`aapi011_reproduce()` / `_detail_reproduce()` + `adz-01032 无上一笔资料可供复制`
- **画面**：`OPEN WINDOW w_aapi011` + `cl_ap_formpath('aap',g_code)`（动态画面档）
- **浏览页**：`DISPLAY ARRAY`(5) / `INPUT ARRAY`(2) + `aapi011_browser_fill`
- **单头/单身**：`glab_t` 为主档，另有 `aapi011_insert_b/_update_b/_delete_b/_lock_b/_unlock_b` 一套 `_b` 后缀的批量操作函数
- **Excel 导出**：`cl_export_to_excel`(4)
- **消息中心**：`cl_msgcentre_notify`
- **自定义字段**：`glabud001~glabud030`（文本/数字/日期三段，符合 erp-code-reader 文档的 ud 规则）
- **报表元件**：`CALL s_axrt`(4)
- 头部有 **31 个版次的修改记录**（`#160318-00005#2 ... 240130-00004#1`），是理解历史包袱的最快入口——**但工具完全没有触及这类"变更历史"信息**

---

## 任务 3：对照实验 —— 工具 vs 裸读

**题目**：`com/qry/4gl/q_adzi052.4gl` 访问哪些表？这些表分别是干什么的？

### 路线 A：用工具

| 步骤 | 命令 | 结果 |
|---|---|---|
| 1 | `grep -oE '\b[a-z][a-z0-9]{1,5}_t\b' q_adzi052.4gl \| sort \| uniq -c` | 11 个 token（含 1 个 `type_t` 噪音） |
| 2 | `tdict r.t "dzaf_t,dzsf_t,dzca_t,dzcal_t,dzcc_t,dzco_t,dzeb_t,dzep_t,gzcc_t,ooaa_t" --conn 主机正式区` | 10 张表的中文名，一次到手 |
| 3 | `tdict r.t "dzaf_t,dzsf_t" --conn 主机正式区` | 关键字段含义（建构编号/版号/类型、程序编号/异动原因/申请解开者） |
| 4 | `tdict r.q q_adzi052 --conn 主机正式区` | SQL 全文 + 显现/回传设定（**这一步裸读完全做不到**） |

**合计 4 次工具调用**（1 次 grep + 3 次 tdict），耗时约 1 分钟（其中 3 次是远程往返）。

### 路线 B：纯裸读源码

不许用工具。我只能从源码里能看到的证据推断：

| 表 | 源码里的证据（行号） | 我**能**推出什么 | 我**推不出**什么 | 真值（工具给的） |
|---|---|---|---|---|
| `dzaf_t` | 573/576 `FROM ds.dzaf_t af0/af1`，用 `dzaf001/dzaf002/dzaf010/dzaf005`，且 `dzaf010 = MIN(...) GROUP BY dzaf001` | 「按 dzaf001 分组取版本」→ 是一张**带版次的主档** | 具体业务对象是"程序建构" | 规格与程序版本历程对应表 |
| `dzsf_t` | 579 `NOT EXISTS (SELECT 1 FROM ds.dzsf_t WHERE dzsf001 = af0.dzaf001)` | 一张**排除清单**（按程序编号） | **排除的原因/业务含义** | Section异动授权数据档 |
| `dzca_t` | 142 `SELECT dzca001,dzcal003 FROM dzca_t`；180 `SELECT dzca004` | 一张"id + 说明 + 一个数字"的**配置表** | 它其实就是**开窗定义档本人**（程序正在读自己的定义） | 开窗数据表 |
| `dzco_t` | 1160/1178 `FROM dzco_t` | 只出现表名，**推不出任何东西** | 全部 | 开窗模糊查询字段设置 |
| `dzeb_t` | 1037 `SELECT dzeb001 FROM dzeb_t WHERE dzeb002 = l_col_id` | "列 → 字段"对照表 | 它是**数据表字段档**（整个数据字典的字段层） | 数据表字段档 |
| `dzep_t` | 1054 `SELECT dzep011 FROM dzep_t WHERE dzep002 = l_col_id` | "列 → 规格"对照表 | 字段规格设计数据档 | 字段规格设计数据档 |
| `gzcc_t` | 1041 `SELECT DISTINCT gzcc003 FROM gzcc_t WHERE gzcc001=...`；1083 `SELECT gzcc004` | 一张分类/代码表 | 是"状态码分类" | 状态码分类 |
| `dzcal_t` | 143 `LEFT JOIN dzcal_t ON dzcal002 = g_lang` | 多语言档 ✅ | —— | 开窗数据表多语言档 ✅ |
| `dzcc_t` | 1182 `SELECT 1 FROM dzcc_t` | 只出现表名，**推不出任何东西** | 全部 | 开窗显现设置表格 |
| `ooaa_t` | 177 注释 `取值優先權: ds.dzca_t>dsdemo.ooaa_t` | 一张**被覆盖的企业级配置表** | 企业层级参数表 | 企业层级参数表 |

**准确度**：10 张表，我能**说对**的只有 2 张（`dzcal_t`、部分 `ooaa_t`），**方向对但含义错**的 2 张（`dzsf_t` 我推成"排除清单"，实际是"异动授权申请档"；`dzeb_t` 推成"列对照表"，实际是"字段档"），**完全推不出**的 6 张。**准确率 20%**。

而且最关键的一层——**这个开窗业务上是干什么的**——裸读**给不出**。因为文件里只有引擎，业务在元数据里。**路线 B 在"用途"这个问题上是 0 分。**

### 对比

| 维度 | 路线 A（工具） | 路线 B（裸读） |
|---|---|---|
| 步数 | 4 次调用 | 至少 10 次 grep（每表 1 次），实际要读 1,274 行才敢下结论 |
| 耗时 | ~1 分钟 | ~5-10 分钟 |
| 表清单准确度 | 100% | 漏 `dzaf_t`/`dzsf_t`（schema 限定 + 多行拼接字符串，朴素正则会漏） |
| 表**语义**准确度 | 100% | 20% |
| 能回答"业务用途"吗 | 能（`r.q` 的 `说明`） | **不能** |
| 上下文消耗 | ~1.5k token（10 行输出） | 至少 15k token（1,274 行 + 自己的推理） |
| 依赖 | 需要数据（本地不全 → 需网络） | 无依赖 |

**结论**：在"表是什么"这件事上工具**省 5~10 倍步数与 10 倍上下文**，且准确率高一个量级。**但前提是数据齐**——这次因为本地库缺 `dzca_t`，路线 A 的 4 步里有 3 步必须走网络。

---

## 任务 4：自选探索（非 aap 模块）

**选定**：`erp/axm/4gl/axmi172.4gl`（销售管理模块，此前没读过）

**第 1 步（源码，1 次调用）**：
```
$ sed -n '1,20p' erp/axm/4gl/axmi172.4gl
#+ Description: 集團銷售預測分配比率設定作業

$ grep -oE "\b[a-z][a-z0-9]{1,6}_t\b" axmi172.4gl | sort | uniq -c | sort -rn | head -15
     77 type_t    74 xmie_t    33 xmif_t    7 ooefl_t    3 xmial_t
      3 rtaxl_t    3 pmaal_t    3 oojdl_t   3 ooag_t     3 imaal_t
      2 imda_t     2 gzze_t     1 oocql_t   1 dzzc_t

$ grep -ohE "\b(xm|azz|lib|std|aim|axm)-[0-9]{5}\b" axmi172.4gl
aim-00070 axm-00105 axm-00167 axm-00175 azz-00079 azz-00087 std-00003 std-00004 std-00006 std-00009
```

**第 2 步（工具，1 次调用）**：
```
$ ./tdict.exe r.t "xmie_t,xmif_t,xmial_t,rtaxl_t,pmaal_t,oojdl_t,imda_t,imaal_t" --conn 主机正式区
xmie_t  | 集团销售预测分配比率设置单头档
xmif_t  | 集团销售预测分配比率设置单身档
xmial_t | 销售预测编号设置档多语言档
rtaxl_t | 品类基本数据档多语言档
pmaal_t | 交易对象主档多语言档
oojdl_t | 渠道基本数据档多语言档
imda_t  | 料件引入营运据点产品分类档
imaal_t | 料件多语言档
```

**第 3 步（工具，1 次调用）**：
```
$ ./tdict.exe msg "axm-00105,axm-00167,axm-00175,aim-00070" --conn 主机正式区
axm-00105 | 输入的资料无效！
axm-00167 | 单身分配比率总和不等于100！
axm-00175 | 单身已有资料，是否删除单身资料重新产生？
aim-00070 | 输入的资料无效！
```

**结论（3 步，约 40 秒）**：`axmi172` 是「**集团销售预测分配比率设定作业**」——维护「集团 → 各销售预测编号」的分配比率，单头 `xmie_t` / 单身 `xmif_t`；校验"单身分配比率总和必须 = 100%"（`axm-00167`）；辅助数据来自料件 `imda_t`、渠道 `oojdl_t`、品类 `rtaxl_t`、交易对象 `pmaal_t`。

**哪些命令真正有用**：
- `sed/grep` 拿 Description + 表 token + 消息编号（**必做，工具替代不了**）
- `tdict r.t <多表>`（**最有用**，1 次调用 8 张表中文名）
- `tdict msg <多编号>`（有用，尤其带"建议作业"时）

**哪些白费**：
- `tdict r.t` 里对 `type_t` / `imda_t` 这类噪音没过滤（是我 grep 的锅，不是工具的）
- 我试的 `tdict r.t "gc_t"`（不存在的表）→ 一次无效网络往返，浪费约 20 秒。工具正确报 `未找到表 'gc_t' 或该表无字段定义`，但**这个错误是在连完数据库之后才给的**——对远程查询来说，本该能更早判断。
- `tdict desc <表>` 在本任务**没用上**（因为目的是"用途和数据依赖"，不是字段细节）。这是**任务决定的**，不是命令没用。

**没能回答的**：`axmi172` 这个"作业"在 T100 菜单里的正式作业名/所属菜单——工具没有作业目录，只能说"Description 里写了"。

---

## 任务 5：评价

### 5.1 摩擦点清单

---

**F1. 随包的本地库只装了 8 个命令里的 1 个，而文档说"数据都在本地"** —— **严重度：挡路**

- **我想做什么** → 按 README「快速上手」把每条示例都跑一遍。
- **我期望什么** → 文档说了"查询读的是本地数据"、"本地已随工具带好数据…合计约 85 万行"，我期望 8 条示例全通。
- **实际发生什么** →
  ```
  $ ./tdict.exe r.v v_ooba002_07
  Error: 本地库尚未包含校验定义 (dzcd_t 等表)数据。请先在有数据库的环境执行: tdict db sync
  $ ./tdict.exe scc 4
  Error: 本地库尚未包含系统分类码 (gzca_t 等表)数据。请先在有数据库的环境执行: tdict db sync
  $ ./tdict.exe desc oobd_t
  Error: 本地库尚未包含字段规格 (dzep_t)数据。请先在有数据库的环境执行: tdict db sync
  $ ./tdict.exe r.q q_apca001
  Error: 本地库尚未包含开窗定义 (dzca_t 等表)数据。请先在有数据库的环境执行: tdict db sync
  $ ./tdict.exe msg std-00006
  Error: 本地库尚未包含消息档 (gzze_t/gzzal_t)。请先执行 tdict db sync,或用 --conn <环境名> 远程直查
  $ ./tdict.exe sysp A-SYS-0040
  Error: 本地库尚未包含参数档 (gzsz_t/gzszl_t)。请先执行 tdict db sync,或用 --conn <环境名> 远程直查
  ```
  直接查 SQLite 证实：库内只有 `dzea_t/dzeal_t/dzeb_t/dzebl_t/dzec_t/dzed_t/dzee_t/dzef_t/dzeg_t`（r.t 一族），约 49.5 万行。[数据]
- **我怎么绕过去的** → 试了 `--conn 主机正式区`，成功。`恒烁正式区/恒烁测试区` 连不通。
- **绕不过去的部分** → 没有网络/没有账号的人，**8 个命令只能用到 1 个**，工具价值砍掉 7/8。而且文档从头到尾没说"你可能需要先检查本地库完整性"。
- **附注**：`db sync` 确实有效（我写临时库验证过：`同步完成: 2 张表, 共 16201 行`），错误信息给的指引也是对的。**问题在文档的承诺，不在工具的能力。**

---

**F2. 工具没有"程序/作业字典"，而这是读代码的第一个问题** —— **严重度：挡路**

- **我想做什么** → 看到 `aapi011.4gl`，想知道"这个程序做什么"。
- **我期望什么** → 至少一个 `tdict prog aapi011` 之类的命令给中文作业名 + 所属模块 + 相关表。
- **实际发生什么** → `tdict --help` 里 14 个命令，没有一个接受程序编号。`tdict r.t` 只认表名。
  ```
  $ ./tdict.exe r.t aapi011
  未找到表 'aapi011' 或该表无字段定义。
  ```
- **我怎么绕过去的** → 只能读源码文件头 `#+ Description:`。四个文件的用途**全部**是这么得到的 `[源码]`。
- **绕不过去** → 读一个**没有写 Description 的文件**（比如 `q_adzi052.4gl`，Descriptions 为空）就完全卡死，只能靠 `tdict r.q q_adzi052` 这条**歪打正着**的路径。而 `cl_abi.4gl` 这种函数库，既没有作业码也没有开窗码，**只能一行行读**。
- **严重度理由**：读代码的第一步就是"这是什么"，工具在这里缺席，等于每次都要先花 10-30 秒读文件头。

---

**F3. `r.t` 没有 `--kw`，无法按中文说明找表** —— **严重度：烦人**

- **我想做什么** → 看代码时想到一个业务概念（"供应商对账单"、"账套科目"），想知道对应哪张表。
- **我期望什么** → `tdict r.t --kw 对账` 列出匹配的表；跟 `r.v` / `scc` / `r.q` 一样有 `--kw`。
- **实际发生什么** →
  ```
  $ ./tdict.exe r.t --help
  Flags:
    -h, --help   help for r.t      <- 只有 help，没有任何 --kw
  $ ./tdict.exe r.t --kw 对账
  Error: accepts 1 arg(s), received 0
  ```
  对比 `r.v`/`scc`/`r.q` 都有 `--kw 按识别码/说明过滤`（`--help` 里可查）。**同一套工具，唯独最基础的表字典没有搜索。**
- **我怎么绕过去的** → 猜表名前缀（`ap*_t` 是应付、`gl*_t` 是总账），或者先 grep 源码拿到表名再反查。
- **绕不过去** → 3,882 张表，靠猜前缀不可能。这是**读代码时最自然的检索方向**（从业务词找表），工具正好缺。

---

**F4. 输出里的 `type_t` 之类噪音没有区分，且工具不告诉你"这个 token 不是表"** —— **严重度：小瑕疵**

- **我想做什么** → 从源码里快速提取"这个文件用到的表"。
- **我期望什么** → 工具能帮我做这一步（哪怕是个 `tdict scan <file>` 之类的辅助）。
- **实际发生什么** → 工具完全不做源码扫描（可以理解，它是数据字典不是 linter）。我自己写的正则把 `type_t`（4GL 类型记录，不是表）也算进去了；`q_adzi052` 里 32 次 `type_t` 是最高的 token。**更危险的是漏**：`ds.dzaf_t` 这种 schema 限定写法、以及拼接字符串里的表名，朴素正则会**整个漏掉**——我第一遍就漏了主角表 `dzaf_t`/`dzsf_t`。
- **我怎么绕过去的** → 改用"全量 `_t` token 扫描 + 人工筛"。
- **绕不过去** → 每次都要人工判断哪些是表。如果工具有 `--kw` 反查（F3），至少能批量确认。

---

**F5. `erp-code-reader/SKILL.md` 的程序命名规则与实际代码不符（`_88` 后缀不存在）** —— **严重度：挡路（会直接找不到文件）**

- **我想做什么** → 按 SKILL.md 的规则找 `aapq110` 的子程序。
- **我期望什么** → SKILL.md 说「**子程序**：SSSQ999_88（一般）/ SSSQ999_MM_88（行业专用）」，那我应该去找 `aapq110_88`。
- **实际发生什么** → 全树扫了一遍：
  ```
  $ find . -name "*_88.4gl" | wc -l
  1                              <- 整个 20,218 个 .4gl 里只有 1 个
  $ find erp -name "*_88.4gl" | wc -l
  0
  $ find erp -name "*_[0-9][0-9].4gl" | sed 's/.*_\([0-9][0-9]\)\.4gl/\1/' | sort | uniq -c | sort -rn
   484 01      222 02      112 03       64 04       38 05
    22 06       19 07       14 09
  $ find erp -name "*_x[0-9][0-9].4gl" | sed 's/.*\(_x[0-9][0-9]\)\.4gl/\1/' | sort | uniq -c | sort -rn
   542 _x01     73 _x02     23 _x03     11 _x04      9 _x05
  $ find erp com -name "*_g01.4gl" | wc -l
  302
  ```
  真实约定是 **`_01`/`_02`…（子程序）**、**`_x01`/`_x02`…（报表查询元件）**、**`_g01`/`_g02`…（报表元件）**，文档里写的 `_88` / `_x88` / `_g88` **一个都不存在**（`_x88`=0）。
- **我怎么绕过去的** → 自己 ls 目录看实际文件名，忽略文档这一节。
- **绕不过去** → 一个**完全信任文档**的人，会浪费大量时间去 find 一个不存在的文件，然后开始怀疑"是不是镜像没拉全"。
- **附带**：`_wf` 后缀（487 个文件，如 `axmi125_wf.4gl`、`aapp132_wf.4gl`）文档**完全没提**。

---

**F6. `desc` / `r.t` 的"备注"列有时极有价值，有时是空的，文档没说这个列会说什么** —— **严重度：小瑕疵**

- **我想做什么** → 判断一个字段的取值来源。
- **我期望什么** → 文档说 `r.t` 输出的"备注"是"备注"（README 表格里只写"备注"两字）。
- **实际发生什么** → 查 `glab_t`，备注列直接给出了 `SCC.8301科目應用設定分類`、`SCC_8315科目彙整方式`、`主要的預設會計科目。`——这是**整份输出里最有价值的一列**，但 README/SKILL.md 对它只有"备注"这个字面描述，完全没提它可能包含 SCC 引用和业务说明。
  ```
  3  glab001  设置类型  varchar2  10  PK Y  SCC.8301科目應用設定分類
  13 glab011  科目汇总方式 varchar2 10       SCC_8315科目彙整方式
  ```
- **我怎么绕过去的** → 偶然看到的，纯运气。
- **绕不过去** → 不会造成失败，但**埋没了工具最好的功能之一**。

---

**F7. `msg` 只给文本，不给"这条消息在代码里的语境"** —— **严重度：烦人**

- **我想做什么** → 理解 `cl_abi.4gl` 里为什么满屏都是 `CALL cl_err(...)`。
- **我期望什么** → 至少 `msg` 的输出能让我判断这条消息是"正常提示"还是"异常分支"。
- **实际发生什么** →
  ```
  $ ./tdict.exe msg "azz-00665,azz-00666" --conn 主机正式区
  azz-00665 | 错误 (1) | 所有用户
  azz-00666 | 错误 (1) | 所有角色
  ```
  两条都标"错误"，但文本是"所有用户"/"所有角色"——这**看起来是下拉选项标签，被登记成了错误消息**。光看 `msg` 输出，我会把它当成"报错文案"，回源码上下文才知道不是。
- **我怎么绕过去的** → 回源码看用法。
- **绕不过去** → `msg` 在"批量扫消息编号"时会产生**误导性上下文**：它给了文本，没给"谁在什么时候调它"。对读代码来说，**消息编号本身没意义，调用点才有意义**。

---

**F8. 两个 `--conn` 语义不同（根命令 vs `db sync` 组内）** —— **严重度：烦人（但文档已警告）**

- **我想做什么** → 用 `db sync --conn 恒烁测试区` 指定环境刷数据。
- **我期望什么** → `--conn` 在哪儿都是"环境名"。
- **实际发生什么** → `db sync` 里 `-c/--conn` 是它自己的 flag（`连接名称`），根命令的 `--conn` 是"查询数据源"。README 第 588 行确实写了「注意：root 的 `--conn`（查询数据源切换）与 `db sync/ping` 组内自带的 `-c/--conn` 是两个独立 flag」。**文档说了，但两个同名 flag 在同一工具里语义不同，本身就是设计味道。**
- **我怎么绕过去的** → 读了文档这一句。
- **绕不过去** → 不至于。

---

**F9. 远程查询的失败反馈慢** —— **严重度：小瑕疵**

- **我想做什么** → `tdict db sync --conn 恒烁测试区 --table gc_t`。
- **实际发生什么** → 先花了几十秒等 TCP 超时（`dial error: timeout: context deadline exceeded`），才知道网络不通。
- **我怎么绕过去的** → 换环境。
- **绕过不去** → 每个新环境都要付一次"不通"的代价。缺一个 `tdict db ping` 之外的快速连通性预检提示。

---

### 5.2 打分（1–5 分）

| 项目 | 分数 | 理由 |
|---|---|---|
| **上手难度** | **2 / 5** | 装好即用、`--help` 清楚（这部分是 5 分）；但照着 README 抄，**第 4 条命令开始全红**，8 个命令 7 个不能用，而这个"文档承诺 vs 实际数据"的落差要自己花 6-7 步才能诊断出来。上手成本主要不在学工具，在**排掉文档造成的误判**。 |
| **输出可读性** | **4.5 / 5** | `r.t` 的表格对齐、`r.q` 的 SQL 分段 + 占位符图例、`desc` 末尾主动给"去哪看更多"、`msg` 有个"建议作业"列、`docp` 给单据性质绑定——**输出设计是成熟产品级的**。唯一扣分：`msg` 的孤立文本（F7）、`r.t` 的"备注"列含义未在文档解释（F6）。 |
| **信息完整度** | **5 / 5（有数据时）/ 1 / 5（随包数据）** | 有数据时：表 3,882 张、字段 14.8 万、开窗 SQL 全文、SCC 含自定义扩展列、参数含型态/值域/预设值/维护作业——**这是我见过最完整的 T100 元数据出口**。随包数据：**8 个命令家族只装了 1 个**。综合给 **3 / 5**。 |
| **相对裸读源码的增益** | **5 / 5** | 任务 3 的实测：表语义准确率 100% vs 20%，步数 4 vs 10+，上下文 1.5k vs 15k token。**`r.q` 和 `desc` 是裸读在原理上不可能企及的**——业务语义和字段绑定根本不在 `.4gl` 里。 |
| **SKILL.md 可用性** | **2.5 / 5** | 好的一面：命令速查表、输出示例、"始终优先用 TDict 而不是去翻 CSV/源码"这句导向很对；`erp-code-reader` 的表/字段命名规范、通用字段表（`crtid/stus/docno/ent/site`）、索引规范**非常有用而且是裸读得不到的**。差的一面：程序后缀 `_88` **整节错误**（F5）、完全没提 `_x01`/`_wf`、没提"本地库可能不全"、没有"程序名怎么映射到用途"的指导（而这恰是第一个问题）。 |
| **README 可用性** | **3 / 5** | 结构清晰、示例详尽、类型码表齐全（表类型 B/M/T/D/L/V/X/H、字段数据类型 N004/C003...）**这些是我实际用上的**。但**「能查到的数据（规模）」表和「快速上手」都在承诺一个不存在的本地库**（F1）；`r.t` 没 `--kw` 这么大的缺口，文档的全局参数表也没露面（只在 `r.v/scc/r.q` 各自小节里提）。 |

---

### 5.3 工具改进建议（按对"读懂代码"的价值排序，最多 8 条）

1. **加「程序/作业字典」命令** —— `tdict prog <程序编号>`：给出中文作业名、所属模块、类型（i/m/t/s/p/q/r）、主档表、被引用的开窗。**这是读代码要回答的第一个问题，现在完全空白**。数据源就在 `gzde_t`（子程序及应用元件基本数据表）、`gzzj_t`（模块编号明细表）里，工具已经连得上。
   - 最低成本的临时版本：把 `msg` 的"建议作业"字段全表导出成一个 `程序编号 → 作业名` 的索引。

2. **给 `r.t` 加 `--kw`（按表名/表说明搜索）** —— 与 `r.v/scc/r.q` 对齐。读代码时"从业务词找表"是最自然的动作，现在只能猜前缀。顺带把 `desc` 也加上。

3. **修 F1：让"本地库不全"在第一时间、以可操作的方式暴露** ——
   - 在 `tdict --help` / 每个数据命令的 `--help` 里显示一行：`本地数据: 表字典 ✓ / 校验 ✗ / 分类码 ✗ ...（tdict db sync 补齐）`
   - 新增 `tdict db status`：列出本地库各数据族的行数与覆盖状态。
   - 把 `-v` 从"只打印路径"升级为"路径 + 各数据族行数"。

4. **新增 `tdict file <路径>`（源码扫描辅助）** —— 输入一个 `.4gl`，输出：文件头 Description、引用到的表（**识别 `ds.xxx_t` schema 限定写法、`LIKE xxx_t.field` 字段定义、拼接字符串里的表名**，并剔除 `type_t` 这类 4GL 类型）、消息编号、SCC 码、`X-XXX-0000` 参数码、开窗调用。**再自动带上每个结果的 `r.t`/`msg`/`scc` 解码。** 这一条能把"读一个文件"从 4-10 步压到 1 步。
   - 若嫌重，退一步：`tdict r.t --from-file <路径>`。

5. **`msg` 输出加"引用此消息的程序"** —— 消息表（`gzze_t`）里若有引用方信息就直接输出；没有就至少标注"此编号为通用消息，需回源码看调用点"。避免 F7 的误导。

6. **`r.t` 输出把"备注"列单独做语义说明，并在有 SCC 引用时把它结构化** —— 现在 `SCC.8301科目應用設定分類` 是纯文本，如果解析成 `SCC 8301（科目應用設定分類）` 并允许直接 `tdict scc 8301`，价值翻倍。README 表格里给"备注"补一行说明。

7. **加 `tdict prog --list-tables <程序编号>` / 或在 `r.t` 里支持反查** —— 输入表名，列出「哪些开窗 (`r.q`) 查了这张表」「哪些校验 (`r.v`) 用了这张表」。现在 `dzca_t` 这种"程序正在读自己的定义"的关系，只有读了源码才知道。

8. **给 `--conn` 加连通性预检与超时提示** —— `tdict db ping` 已有，但查询命令在环境不通时要等 TCP 超时（F9）。检测到指定环境的 SSH/DB 从未连通时，直接给出"该环境不可达，是否改用 local？"的提示，而不是等几十秒后报 `dial error`。

---

### 5.4 提示词（SKILL.md）改进建议（最多 8 条，附建议原文）

**针对 `skills/tdict/SKILL.md`：**

1. **开头加"先验数据完整性"** —— 现在第二节写「本地已随工具带好数据（约 3,882 张表的字段定义、148,539 个字段，以及校验/分类码/开窗/消息/参数等元数据，合计约 85 万行）」。这句话在本机上**是假的**，会让人把"数据没同步"误判成"工具坏了"。改为：

   > **先确认本地数据是否完整**：`tdict` 的各命令需要不同的数据族。若某命令报「本地库尚未包含 XXX 数据」，说明本地 `erp_data.db` 只同步了部分内容——**这不代表命令不可用**，任选其一：(a) 加 `--conn <环境名>` 远程直查（先用 `tdict env` 看有哪些环境）；(b) 执行 `tdict db sync` 补齐本地库。随包/随仓库的 `erp_data.db` **可能只含表字典一族**（`dzea_t/dzeb_t/...`），其余需自行同步。

2. **把「始终优先用 TDict」这句加个前提** —— 现在是「**始终优先用 TDict 而不是去翻 CSV/源码文件**：查到的说明自带中文翻译与字段语义，快且全。」这句在数据不全时会**直接导致失败**。改为：

   > **有数据时始终优先用 TDict，而不是去翻 CSV**：说明自带中文翻译与字段语义，快且全。**但要先确认该命令的数据族在本地可用**（见上）。TDict **不回答**「某个 .4gl 程序做什么」——那要看源码文件头的 `#+ Description:`。

3. **补一条"程序编号 → 用途"的路子** —— 现在完全没有。加：

   > **想知道某个程序（如 `aapi011`）做什么**：本工具没有程序字典。三条路：(1) 读源码文件头 `#+ Description:`（最可靠）；(2) 若该程序是开窗引擎，用 `tdict r.q q_xxxxx` 查其登记说明与 SQL；(3) 用 `tdict msg <该程序相关的消息编号>`，看「建议作业」字段反推同族作业中文名。

**针对 `skills/erp-code-reader/SKILL.md`：**

4. **修正 `_88` 后缀（最关键的一条）** —— 现在是：
   > - **子程序**：SSSQ999_88（一般）/ SSSQ999_MM_88（行业专用）
   > - **子画面**：SSSQ999_s88 或 SSSQ999_88_s88

   实测 `*_88.4gl` 全树只有 1 个（erp 下 0 个）。改为：

   > - **子程序**：SSSQ999_**01**（一般）/ SSSQ999_MM_**01**（行业专用）；多支时依序 `_02`、`_03`…（实测：`_01` 484 个、`_02` 222 个、`_03` 112 个）
   > - **子画面**：SSSQ999_s**01** 或 SSSQ999_**01**_s**01**
   > - **作业流程（Workflow）**：SSSQ999_**wf**（实测 487 个，如 `axmi125_wf`、`aapp132_wf`）

5. **修正报表元件表** —— 现在是：
   > | 凭证报表 | SSSQ999_g88 | SSSQ999_MM_g88 |
   > | 查询报表 | SSSQ999_x88 | SSSQ999_MM_x88 |

   实测 `_g01` 302 个、`_x01` 542 个，`_x88` 0 个。改为：

   > | 凭证报表 | SSSQ999_**g01**（多支依序 `_g02`…） | SSSQ999_MM_**g01** |
   > | 查询／条件元件 | SSSQ999_**x01**（多支依序 `_x02`…） | SSSQ999_MM_**x01** |

   并把下面的「范例汇总」表同步改掉（现在写的是 `axmr500_g01`，这个恰好是对的——说明表里大部分地方对，只有 `_88` 系统性写错）。

6. **补一句"读 .4gl 的入口顺序"** —— 现在文档大篇幅讲命名规则，但**没讲"打开一个 .4gl 该先看哪**"。加：

   > **读一个 .4gl 的顺序**：① 文件头 `#+ Description:`（业务用途，最可靠）；② 文件头的 `#+ Modifier...` 修改记录（31 条历史往往能一眼看出这个程序的坑）；③ `SCHEMA` / `GLOBALS` / `IMPORT` 段；④ 用 `grep -oE '\b[a-z][a-z0-9]{1,6}_t\b' | sort | uniq -c` 提表名（**注意 `type_t` 是 4GL 类型不是表；注意 `ds.xxx_t` 这种 schema 限定写法会被朴素正则漏掉**），再交给 `tdict r.t` 解码；⑤ 用 `tdict r.q` 查被 `CALL q_xxx()` 的开窗（**业务语义在元数据里，不在 .4gl 里**）。⑥ 注意：`{<section ...>}` 与 `#add-point:` 是样板噪音，有效信息只在 add-point 内和文件头。

7. **补"哪些业务语义不在源码里"的警告** —— 这是本次最大的认知收获。加：

   > **重要：T100 的业务语义经常不在 `.4gl` 里。** 例如 `com/qry/4gl/q_adzi052.4gl`（1,274 行）本身只是**通用开窗引擎**，它的表、条件、显现/回传列全部登记在元数据表 `dzca_t/dzcc_t/dzco_t` 里——**裸读源码在原理上读不出它的用途**，必须用 `tdict r.q q_adzi052`。同理 `desc <表> <字段>` 能给出字段绑定的开窗程序与校验码，这也是源码里看不到的。

8. **补 `msg` 的语境警告 + 建议作业的用法** —— 加：

   > `tdict msg` 给的是**消息文本**，不是**调用语境**。像 `azz-00665 所有用户`、`azz-00666 所有角色` 这种"文本看起来像选项标签"的消息，光看文本会误导；要判断它是不是异常分支，得回源码看 `CALL cl_err(...)` 的位置。反过来，`msg` 的 **「建议作业」** 字段很有用——它是工具里唯一能把"程序编号 ↔ 作业中文名"连起来的入口（例：`aap-00001` → `aapi010 (应付账务类型维护作业)`）。

---

### 5.5 一句话结论

**推荐，但必须配一句提醒。** `tdict` 的 `r.q`（开窗 SQL 全文）、`desc`（字段绑定的开窗与校验码）、`r.t`（3,882 张表的中文名与字段语义）能回答**裸读源码在原理上回答不了**的问题——这次实测里 `q_adzi052` 读 1,274 行读不出用途、一条 `tdict r.q` 就出来了，这是不可替代的价值；但前提是**先把数据补齐**（`--conn` 远程或 `db sync`），并且**不要相信随包 `erp_data.db` 是完整的**，也**别指望 `erp-code-reader/SKILL.md` 的 `_88` 后缀命名规则**——那两处是这次最坑人的地方，修好之后我会毫不犹豫地推给同事。

---

## 附录：本次执行的完整命令记录

```bash
# --- 任务 1：上手 ---
ls -la tdict.exe erp_data.db README.md ; which tdict ; ./tdict.exe --help
./tdict.exe r.t dzea_t ; ./tdict.exe r.t dzea_t -v
./tdict.exe r.t dzea_t                       # OK
./tdict.exe r.v v_ooba002_07                 # Error: 本地库尚未包含校验定义
./tdict.exe scc 4                            # Error: 本地库尚未包含系统分类码
./tdict.exe desc oobd_t                      # Error: 本地库尚未包含字段规格
./tdict.exe r.q q_apca001                    # Error: 本地库尚未包含开窗定义
./tdict.exe msg std-00006                    # Error: 本地库尚未包含消息档
./tdict.exe sysp A-SYS-0040                  # Error: 本地库尚未包含参数档
./tdict.exe docp D-MFG-0076                  # Error: 本地库尚未包含参数档
python -c "sqlite3 ... select name from sqlite_master"   # 只有 r.t 一族
./tdict.exe env ; ./tdict.exe db list
./tdict.exe msg std-00006 --conn 主机正式区   # OK
./tdict.exe r.v v_ooba002_07 --conn 主机正式区 # OK
./tdict.exe scc 4 --conn 主机正式区            # OK
./tdict.exe desc oobd_t --conn 主机正式区      # OK
./tdict.exe r.q q_apca001 --conn 主机正式区    # OK
./tdict.exe sysp A-SYS-0040 --conn 主机正式区  # OK
./tdict.exe docp D-MFG-0076 --conn 主机正式区  # OK
./tdict.exe db sync --conn 恒烁测试区 -d <tmp> --table dzca_t   # dial timeout
./tdict.exe db sync --conn 主机正式区 -d <tmp> --table dzca_t,dzcal_t  # 同步完成 16201 行
./tdict.exe r.q q_adzi052 -d <tmp>            # 从临时库可查
rm -rf /d/t100_source_code/.tmp_eval          # 还原源码树

# --- 任务 2 ---
# aapq110_01.4gl
head -80 erp/aap/4gl/aapq110_01.4gl ; sed -n '80,224p' ...
grep -oE "\b[a-z][a-z0-9]{1,6}_t\b" aapq110_01.4gl        # 无表
sed -n '1,20p' erp/aap/4gl/aapq110.4gl                   # Description: 供應商對帳單明細查詢
# q_adzi052.4gl
head -60 com/qry/4gl/q_adzi052.4gl ; grep -nE "^...FUNCTION " ...
grep -nEi "(SELECT|PREPARE|DECLARE .* CURSOR|FOREACH )" ...
grep -oE "\b[a-z][a-z0-9]{1,5}_t\b" | sort | uniq -c        # 11 个 token
./tdict.exe r.t "dzaf_t,dzsf_t,dzca_t,dzcal_t,dzcc_t,dzco_t,dzeb_t,dzep_t,gzcc_t,ooaa_t" --conn 主机正式区
./tdict.exe r.q q_adzi052 --conn 主机正式区                 # 说明 + SQL 全文 + 显现设定
./tdict.exe msg "adz-00926,adz-01361" --conn 主机正式区
./tdict.exe sysp "A-SYS-0073" / "E-SYS-0002" / "E-SYS-0003" --conn 主机正式区
# cl_abi.4gl
head -50 com/lib/4gl/cl_abi.4gl ; grep -nE "^(PUBLIC|PRIVATE) FUNCTION " ...
grep -oE "\b[a-z][a-z0-9]{1,6}_t\b" | sort | uniq -c
sed -n '121,190p' / '1190,1230p' / '1659,1680p' / '1853,1872p' cl_abi.4gl
./tdict.exe r.t "gziu_t,gziul_t,gzpe_t,gzgn_t,gzgh_t,gzgu_t,oojn_t,ooek_t,..." --conn 主机正式区
./tdict.exe msg "azz-00665,azz-00666,lib-00073,lib-00082,lib-00112,..." --conn 主机正式区
./tdict.exe sysp A-SYS-0095 --conn 主机正式区
# aapi011.4gl
sed -n '1,40p' erp/aap/4gl/aapi011.4gl ; grep -nE "^(PUBLIC|PRIVATE) FUNCTION " ...  # 51 个函数
grep -oE "\b[a-z][a-z0-9]{1,6}_t\b" | sort | uniq -c
grep -nE "SCHEMA |DECLARE .* CURSOR|PREPARE " ...
./tdict.exe r.t "glab_t,glaa_t,glac_t,glad_t,..." --conn 主机正式区
./tdict.exe r.t glab_t --conn 主机正式区                    # 58 字段 + SCC 备注
./tdict.exe scc 8315 / scc 8504 --conn 主机正式区
./tdict.exe msg "aap-00001,aap-00002,std-00006,adz-01032" --conn 主机正式区
./tdict.exe desc glab_t glab002 --conn 主机正式区
grep -nE "set_combo_scc|ON ACTION controlp|FOR UPDATE" aapi011.4gl

# --- 任务 4 ---
ls erp/ ; ls erp/axm/4gl/ | shuf -n 3
sed -n '1,20p' erp/axm/4gl/axmi172.4gl ; grep -oE ... ; grep -ohE ...msg...
./tdict.exe r.t "xmie_t,xmif_t,xmial_t,rtaxl_t,pmaal_t,oojdl_t,imda_t,imaal_t" --conn 主机正式区
./tdict.exe msg "axm-00105,axm-00167,axm-00175,aim-00070" --conn 主机正式区

# --- 命名规则验证（F5） ---
find . -name "*_88.4gl" | wc -l ; find erp -name "*_[0-9][0-9].4gl" | sed ... | sort | uniq -c
find erp -name "*_x[0-9][0-9].4gl" | ... ; find erp com -name "*_g01.4gl" | wc -l
find erp com -name "*_wf.4gl" | wc -l
```

**工具调用总量**：约 26 次 shell 调用（含 3 次 Python/SQLite 直查、1 次 `db sync` 验证、约 34 次 `tdict` 子命令，其中约 20 次为远程 `--conn`）。
**源码修改**：无（临时库目录已删除）。
