---
title: "EXIT DIALOG instruction"
source: "fgl-topics/c_fgl_DIALOG_instr_EXIT_DIALOG_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control instructions > EXIT DIALOG instruction"
type: "concept"
description: "Syntax EXIT DIALOG Usage The EXIT DIALOG statement terminates a procedural DIALOG block without any further control block execution. When used in a declarative DIALOG block, the EXIT DIALOG ..."
---

# EXIT DIALOG instruction

## Syntax

```
EXIT DIALOG
```

## Usage

The `EXIT DIALOG` statement terminates a procedural `DIALOG` block
without any further control block execution.

When used in a declarative `DIALOG` block, the `EXIT DIALOG`
instruction does only make sense when the declarative dialog block is included in a procedural
dialog block with the `SUBDIALOG` clause.

Program flow resumes at the Instruction following the `END DIALOG` keywords.
Blocks such as `AFTER DIALOG` will not be executed.

```
   ON ACTION quit 
       EXIT DIALOG
```

When leaving the `DIALOG` instruction, all form items used by the dialog will be
disabled until another interactive statement takes control.

The `EXIT DIALOG` instruction leaves the `DIALOG` block
immediately, while [`CANCEL
DIALOG`](2143-cancel-dialog-instruction.md) makes some additional tasks.

## Related links

**Related concepts**  

[CANCEL DIALOG instruction](2143-cancel-dialog-instruction.md "CANCEL DIALOG instruction")

[ACCEPT DIALOG instruction](2142-accept-dialog-instruction.md "ACCEPT DIALOG instruction")
