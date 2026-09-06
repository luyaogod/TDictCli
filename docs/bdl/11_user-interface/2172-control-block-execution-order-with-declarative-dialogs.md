---
title: "Control block execution order with declarative dialogs"
source: "fgl-topics/c_fgl_declarative_dialogs_ctrlblock_exec_ord.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control blocks > Control block execution order with declarative dialogs"
type: "concept"
description: "Context-dependent behavior Depending on the usage context, dialog-starting and ending control blocks such as BEFORE INPUT / AFTER INPUT are executed each time the focus goes to an element controlled ..."
---

# Control block execution order with declarative dialogs

## Context-dependent behavior

Depending on the usage context, dialog-starting and ending control blocks such as `BEFORE
INPUT` / `AFTER INPUT` are executed each time the focus goes to an element
controlled by the dialog, or only once, when the dialog is started/ended. See control block specific
topics for more details.

## Control block execution order when in multiple dialogs

When the declarative dialog is used with a `SUBDIALOG` keywoard in the context of
a procedural `DIALOG` block, the control blocks are executed in the same order as if
the declarative dialog was included in the procedural `DIALOG` block.

For more details, see [Control block execution order in multiple dialogs](2100-control-block-execution-order-in-multiple-dialogs.md).
