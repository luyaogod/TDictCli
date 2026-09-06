---
name: erp-modify
description: 修改框架生成的 TIPTOP ERP Genero BDL (.4gl) 源文件。只有 #add-point:... / #end add-point 区块内的代码可编辑——绝不修改区块外的框架代码。当用户要求在 ERP 4GL 程序中新增或修改业务逻辑时触发本技能。
---

# ERP 修改 — 框架生成程序修改指南

## 概述

本 Skill 用于**修改** TIPTOP ERP 系统中由程序生成器（a00/a02/a03/a26 等样板）自动生成的 Genero BDL 源代码（`.4gl` 文件）。

这类文件在文件头有两个标记：

```
#该程式未解开Section, 採用最新樣板產出!
#该程式非freestyle程式!
```

**核心规则：只有 `#add-point:` 到 `#end add-point` 之间的可编辑区块才能新增/修改代码，区块之外的框架代码一律不得修改。**

框架代码由样板生成，修改后在下一次样板重新生成时会被覆盖丢失；add-point 区块则是生成器刻意保留给客制代码的插入点。

## 识别框架生成文件

- 文件顶部有 `{<section id="<程序名>.<section>" type="s" >}` ... `{</section>}` 的 section 区块
- 各区块内有 `#應用 aXX 樣板自動產生(Version:Y)` 标记（如 `a00`、`a02`、`a03`、`a26`），标示该区块由哪个样板/版本生成
- 文件内有大量 `#add-point:` 标记
- 文件头有 `#+ Description:` 描述程序用途（如 `IC下單平台維護作業`）

## add-point 标记语法

开始标记（单行）：

```
#add-point:<说明文字> name="<region_id>"
```

结束标记（两种写法，照抄原文件大小写）：

```
#end add-point        -- 小写 end（多数区域）
#END add-point        -- 大写 END（ON ACTION / DIALOG handler 内）
```

- `name="<region_id>"` 是分层式识别码，格式为 `<section>.<context>.<detail>`（如 `input.a.sfajwfdocno`、`menu.insert`）
- 缩进继承所在代码区块：文件顶层为第 0 栏，MAIN/函数内 3 个空格，DIALOG handler 内 12 个空格

## 可编辑区域分层

add-point 区块分三种，修改前先判断属于哪一类：

### A. 客制用区域（安全，优先使用）

描述文字带「(客製用)」，或名称属于客制/动作 hook 区域：

- 各 section 的 `*.define_customerization`（如 `main.define_customerization`、`init.define_customerization`）
- `global.memo`、`global.memo_customerization`、`global.variable_customerization`
- `menu.*` 工具栏动作（`menu.insert`、`menu.modify`、`menu.delete`、`menu.query`、`menu.output`...）
- `input.*` DIALOG 字段 handler（`input.b.<字段>` BEFORE FIELD、`input.a.<字段>` AFTER FIELD、`input.c.<字段>` ON ACTION controlp INFIELD）
- `construct.*` CONSTRUCT 字段 handler（`construct.b.<字段>`、`construct.a.<字段>`、`construct.c.<字段>`）
- `ui_dialog.page<n>.before_display`、`.before_row`、`.action` 等画面 hook
- `browser_fill.*` 列表填值 hook

### B. 框架自有区块（警告，勿修改）

描述文字含警告 `(請盡量不要在客製環境修改此段落內容, 否則將後續patch的調整需人工處理)`。

例如 `global.variable`（模块变量声明区）、各 section 的 `*.define`。这些区块虽然在 add-point 标记之间，但内容是框架生成的，修改后后续 patch 调整需要人工处理。

### C. readonly 区段（生成器永不覆写，手工代码）

section 标记为 `readonly="Y"` 的区段（如 `other_dialog`、`other_function`）。这是手工编写的客制代码区，生成器不会重新生成；大量自定义代码可写在这里。

## 修改规则

1. **绝不删除或移动** `#add-point:` / `#end add-point` / `#END add-point` 标记本身
2. **只把代码写在标记之间**，并保持与所在区块一致的缩进
3. **不修改标记外的任何框架代码** — 样板重新生成时会覆盖丢失
4. 优先将新增代码放在 **A 类客制用区域**；若需在现有 add-point 区块内改动，先确认不是 **B 类警告区块**
5. 结束标记大小写照抄原文件：`#end add-point` vs `#END add-point`
6. 在 add-point 区块内可加自己的注释，可用工单编号标注（如 `#250715-00042#1 add ---s`）
7. 修改后确认每个 `#add-point:` 都有对应的结束标记

## 重要规则

1. **只有 add-point 区块可编辑** — 这是最重要的规则
2. 修改前先判断区块分层（A 客制用 / B 框架自有警告 / C readonly 手工）
3. 不修改含 `(請盡量不要在客製環境修改此段落內容...)` 警告的区块
4. 框架代码由样板生成，乱改会在重新生成时被覆盖
5. 若需大量自定义代码，可写在 `readonly="Y"` 的 `other_function` 区段
6. 用 `tdict` 查询表/字段中文含义，不要猜测编号式命名

## 工具参考

- 解读 ERP 源代码（表名/字段命名规范）：参见 [erp-read.md](erp-read.md)
- 查询数据字典（表/字段中文含义）：参见 [tdict.md](tdict.md)
- BDL 语言语法/内置函数参考：项目内 `docs/bdl`(路径用 `tdict bdldoc dir` 查/设)
