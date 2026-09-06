---
title: "CANCEL DIALOG instruction"
source: "fgl-topics/c_fgl_DIALOG_instr_CANCEL_DIALOG.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control instructions > CANCEL DIALOG instruction"
type: "concept"
description: "Syntax CANCEL DIALOG Usage The CANCEL DIALOG statement terminates a procedural DIALOG block, after executing the AFTER INPUT , AFTER DISPLAY , AFTER CONSTRUCT control block of the current sub-dialog, ..."
---

# CANCEL DIALOG instruction

## Syntax

```
CANCEL DIALOG
```

## Usage

The `CANCEL DIALOG` statement terminates a procedural `DIALOG`
block, after executing the `AFTER INPUT`, `AFTER DISPLAY`,
`AFTER CONSTRUCT` control block of the current sub-dialog, and the `AFTER
DIALOG` control block.

When used in a [declarative
`DIALOG` block](2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs."), the `CANCEL DIALOG` instruction does only make
sense when the declarative dialog block is included in a procedural dialog block with the
`SUBDIALOG` clause.

The `CANCEL DIALOG` instruction can be used from multiple dialogs to mimic the
`cancel` default action of single dialogs.

The `CANCEL DIALOG` instruction makes the following:

1. Set the [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") register to
   `TRUE`.
2. If defined, execute the code in the [`AFTER
   INPUT`](1941-after-input-block.md), [`AFTER
   DISPLAY`](1974-after-display-block.md) or [`AFTER
   CONSTRUCT`](2058-after-construct-block.md) block of the current sub-dialog.
3. If defined, execute the code in the [`AFTER DIALOG`](2102-after-dialog-block.md) block.

The statements appearing after the `CANCEL DIALOG` instruction will be
skipped.

You typically code an `CANCEL DIALOG` in an `ON ACTION cancel`
block:

```
ON ACTION cancel CANCEL DIALOG
```

The default settings regarding action attributes for the `cancel` action define
the `validate` attribute to `"no"`, in order to avoid current field
validation for this action. This is important when using the `UNBUFFERED` mode. For
more details, see [Actions configuration for field validation](2235-the-buffered-and-unbuffered-modes.md).

## Related links

**Related concepts**  

[EXIT DIALOG instruction](2141-exit-dialog-instruction.md "EXIT DIALOG instruction")

[ACCEPT DIALOG instruction](2142-accept-dialog-instruction.md "ACCEPT DIALOG instruction")

[ui.Dialog.cancel](../15_library-reference/3188-ui-dialog-cancel.md "Cancels a parent dialog from a sub-dialog.")
