---
name: erp-code-reader
description: 阅读并分析 TIPTOP ERP Genero BDL 源代码。使用 TDict CLI 将表/字段编号解码为中文含义。当用户打开包含 Genero BDL 代码的 .4gl/.per/.txt 文件，或询问 ERP 程序逻辑、代码中引用的数据库表结构、ERP 源文件中的字段含义时触发本技能。
---

# ERP 代码阅读 — Genero BDL 源代码阅读指南

## 获取源码:本地镜像优先

阅读服务器源码前,先确认本地源码镜像是否存在并新鲜:

```bash
tdict mirror path [<环境名>]    # 拿到镜像目录绝对路径(如 D:\dev\erp-src\主机正式区)
tdict mirror pull  [<环境名>]   # 下载/更新该环境源码镜像(增量;--full 全量重建)
```

- 镜像只含各模块 **4gl**(源码)与 **4fd**(前端字段描述)两棵树,目录与服务器同构(`<镜像根>/<环境名>/erp/<模块>/{4gl,4fd}/...`、`com/{lib,sub,qry,wss}/...`);per/编译产物(42m/42r)/多语言等不在镜像内。
- 拿到路径后用本地文件工具直接读/搜索(ls/rg/按行读),不要 ssh/scp 去服务器;镜像未覆盖的文件(如 per)才走 `tdict debug source` 白名单通道。
- 镜像更新时间滞后于服务器时(他人改过代码),先 `tdict mirror pull` 再读;调试停站的行号来自服务器编译表,若与镜像不同代,以服务器文件为准。

## 命令规范说明

### 模块命名原则

模块名称由三个英文字母组成（适用于 ERP 模块及共用 COM 模块等）。

**第一码**为标准识别码：A=标准 ERP，B=行业专用 ERP，C=客制 ERP，D=客制行业专用 ERP，E=客制自创 ERP。**第二、三码**采用系统英文缩写，若已被使用则第二码改以 X 或其他相关缩写替代。

格式：**AXX**（XX 为 ERP 模块英文缩写代码）

**范例**：AIM（料件主档及库存管理，IM 取自 Item Master / Inventory Management）、AXM（销售管理，因 ASM 已被系统基本资料管理使用而改之）。

**共用模块（$COM）代码**：

| 代码 | 说明 | 英文名称 |
| :---: | :--- | :--- |
| lib | 共用程序 | Common Library |
| lng | 语言资料与提速档副程序 | Lingual sub function |
| sub | 一般副程序 | Common Sub function |
| qry | 查询副程序 | Common Query Function |
| wss | 整合与 Web Service 程序 | Web Service Subsystem |


### 程序命名原则

#### 主程序编号

格式：**SSSQ999_MM**（七码：三码英文 + 四码数字，均小写）

- SSS：模块代码
- Q：程序类别 — i=基本资料维护，m=主档维护，t=交易处理，s=参数设定，p=批次处理，q=查询，r=报表
- 999：三码流水号
- _MM：行业别专用代码（如 _ic=IC 设计业，_sc=鞋服饰，_ph=制药业；省略则通用）

**范例**：aimi100（料件基本资料维护）、cimi001（客制料件维护）、bphi100_ph（食品添加物登记证维护）

#### 子程序与子画面

- **子程序**：SSSQ999_88（一般）/ SSSQ999_MM_88（行业专用）
- **子画面**：SSSQ999_s88 或 SSSQ999_88_s88

**范例**：aimi100_01、cimi100_01；aimi100_s01、aimi100_01_s01

#### Library 文件

- **标准**：cl_xxxxxxx（共用程序，如 cl_err、cl_about），长度 1–17 字符
- **客制**：ccl_xxxxxxxxx（如 ccl_trim），长度 1–17 字符

> 客制共用程序暂不开放，lib 功能已完备。

#### 应用元件

- **标准**：s_XXXXXXX 或 s_SSSQ999_XXXXXX（如 s_transaction）
- **客制**：CS_XXXXXXXXX 或 cs_SSSQ999_XXXXXX（如 cs_date）

#### 报表相关文件

报表主程序命名同 2.1。报表元件格式：

| 类型 | 标准 | 行业专用 |
| :---: | :--- | :--- |
| 凭证报表 | SSSQ999_g88 | SSSQ999_MM_g88 |
| 查询报表 | SSSQ999_x88 | SSSQ999_MM_x88 |

- **报表结构档（.rdd）**：与元件同名，如 axmr500_g01
- **报表样板（.4rp）**：主报表 SSSQ999_g88（多样板加 _77），子报表加 _subrep66

**范例汇总**：

| 场景 | 主程序 | 报表元件 | 样板 |
| :--- | :--- | :--- | :--- |
| 标准报表 | axmr500 | axmr500_g01 | axmr500_g01 / axmr500_g01_subrep01 |
| 多样件 | axmr500 | axmr500_g02 | axmr500_g02 / axmr500_g02_subrep01 |
| 多样板 | axmr500 | axmr500_g01 | axmr500_g01_02 / axmr500_g01_02_subrep01 |
| 行业别 | axmr500_ph | bxmr500_ph_g01 | bxmr500_ph_g01 / bxmr500_ph_g01_subrep01 |
| 客制标准 | axmr500 | axmr500_g01 | axmr500_g01 / axmr500_g01_subrep01 |
| 新增客制 | cxmr501 | cxmr501_g01 | cxmr501_g01 / cxmr501_g01_subrep01 |

#### 查询元件

- **标准**：q_XXXXXXX（如 q_imaa001），长度 1–18 字符，以表格名命名
- **客制**：cq_XXXXXXXXX，长度 1–17 字符
- **行业专用**：q_XXXXXXXXX_mm（如 q_ooea001_ph），前缀长度 1–15 字符

> 避免与动态查询副程序名称冲突。

#### Web Service 程序

- **主程序**：wssp999（标准）/ cwssp999（客制）
- **子程序**：wssp999_88（标准）/ cwssp999_88（客制）

存放于 $COM/WSS，子程序不对外供一般 ERP 调用。

#### 共用参数文件（.inc）

- **一般用途**：SSSQ999_88.inc（放 4g1 目录）
- **跨模块**：top_XXXXXX.inc（放 $ERP/cfg，链接至 $COM/cfg）

#### 扩展名说明

| 类别 | 扩展名 | 说明 |
| :--- | :--- | :--- |
| 原始程序 | 4g1 / inc / 42m / 42r | 源码 / 共用参数 / 编译目标 / 可执行 |
| 画面 | 4fd / per / 42f | 设计档 / 对照档 / 编译档 |
| 报表 | rdd / 4rp | 结构档 / 样板档 |
| 资源 | str / 42s / sch / 4tm/4ad/4tb / 4pw | 翻译原始档 / 翻译编译档 / 参考内容 / Genero 设定 / 专案设定 |


### 函数命名原则

格式：**xxxxxxxxx_yyyy**（前 7 码为程序文件名，yyyy 为功能描述）

**范例**：azzi100_insert

**常用函数**：_insert、_delete、_show、fetch、_modify、_input、_query、set_entry、set_no_entry

> 同一主程序下函数名不可重复。


### 变量命名原则

| 类型 | 格式 | 说明 |
| :--- | :--- | :--- |
| 全局变量 | g_XXXXXXXXX | 定义于 $TOP/cfg/top_global.inc，如 g_gui_type、g_errno |
| 局部变量 | l_XXXXXXXXX | 仅限当前 Function 有效 |
| 传递参数 | p_XXXXXXXXX | 跨函数传递 |
| 屏幕变量 | s_detailN / s_browser | N 为流水号，异动可能影响代码生成 |


### 数据库表格命名原则

#### 表格名称

格式：**xxxx_t**（4 码小写 + _t），前两码为模块名，后两码流水号。如 imaa_t（料件主档）。

- **行业包辅助表**：xxxxmm_t（如 imaaic_t、imaasl_t）
- **客制表**：标准表不可删，新建客制表加 uc（如 imaauc_t）

#### 字段名称

格式：**xxxx999z**（表格编号 + 3 码流水号 + 阶层码）。如 imaa001（料件编号）、imaa001a（子阶）。

- **客制字段**：客制表用 apqquc001，标准表加 ud（如 apqqua001、imaaud001）
- **行业包辅助字段**：xxxxmm099（如 imaaic001）
- **自定义字段**：主/明/交易档加 ud，001–010 文字型（C003），011–020 数值型（N101），021–030 日期型（D002）

**固定字尾规范**：

| 用途 | 字尾 | 说明 |
| :--- | :--- | :--- |
| 建立 | crtid / crtdt / crtdp | 员工 / 日期 / 部门 |
| 拥有 | ownid / owndp | 员工 / 部门 |
| 修改 | modid / moddt | 员工 / 日期 |
| 确认 | cnfid / cnfdt | 员工 / 日期（adp 负责） |
| 过账 | pstd / pstdt | 员工 / 日期（adp 负责） |

**通用字段**：

| 字段 | 字尾 | 类型 | 备注 |
| :--- | :--- | :--- | :--- |
| 状态码 | stus | C001 | |
| 单号 | docno | C203 | |
| 单据日期 | docdt | D001 | T 类表单号后必跟日期 |
| 项次 | seq / seq1 / seq2 | N004 | 需逐层依赖 |
| 企业编号 | ent | N802 | |
| 法人 | comp | C813 | |
| 营运据点 | site | C813 | |
| 账别 | ld | C501 | 小写 |
| 组织对象 | unit | C007 | |
| 账务归属 | orga | C007 | |
| 核算组织 | legl | C007 | |
| 时间戳记 | stamp | D003 | 不可为 KEY |

**必要字段**：主/参数/基础资料档需 crtid、crtdp、crtdt、ownid、owndp、modid、moddt、stus；交易档另加 cnfid、cnfdt。

#### 索引与键值

- **索引**：xxxx_Qy（Q=n/u，y 流水号），如 imaa_n01、imaa_u01
  - 客制新表比照标准（imaauc_n01）；标准表加客制索引前缀 tic_（tic_imaa_n01）
  - 行业包辅助表：imaaic_n01
- **Primary Key**：xxxx_pk（r.t 自动命名，唯一不可改），如 imaa_pk
  - 客制：imaauc_pk；行业包：imaaic_pk
- **Footing Data**：xxxxxx_fdy（r.t 自动命名，仅供 r.a 及产生器参考，不实际建库），如 imaa_fd
  - 客制：imaauc_fd；行业包：imaaic_fd
  
## 工具参考

- 查询数据字典(表/字段中文含义):参见 [tdict.md](tdict.md)
- BDL 语言语法/库函数参考:项目内 `docs/bdl`(Genero BDL 文档;路径用 `tdict bdldoc dir` 查/设)