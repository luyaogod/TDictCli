---
title: "CONTINUE DIALOG instruction"
source: "fgl-topics/c_fgl_DIALOG_instr_CONTINUE_DIALOG.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control instructions > CONTINUE DIALOG instruction"
type: "concept"
description: "Syntax CONTINUE DIALOG Usage The CONTINUE DIALOG statement continues the execution of a DIALOG instruction, skipping all statements appearing after this instruction. Control returns to the dialog ..."
---

# CONTINUE DIALOG instruction

## Syntax

```
CONTINUE DIALOG
```

## Usage

The `CONTINUE DIALOG` statement continues the execution of a
`DIALOG` instruction, skipping all statements appearing after this instruction.

Control returns to the dialog instruction, which executes remaining control blocks as if the
program reached the end of the current control block. Then the control goes back to the user and the
dialog waits for a new event.

The `CONTINUE DIALOG` statement is useful when program control is nested within
multiple conditional statements, and you want to return control to the user by skipping the rest of
the statements.

In the following code example, an `ON ACTION` block gives control back to the
dialog, skipping all instructions below line 04:

```
ON ACTION zoom
  IF p_cust.cust_id IS NULL OR p_cust.cust_name IS NULL THEN
    ERROR "Zoom window cannot be opened if no info to identify customer"
    CONTINUE DIALOG
  END IF
  IF p_cust.cust_address IS NULL THEN
     ...
```

If `CONTINUE DIALOG` is called in a control block that is not `AFTER
DIALOG`, further control blocks might be executed depending on the context. Actually,
`CONTINUE DIALOG` just instructs the dialog to continue as if the code in the control
block was terminated (it is a kind of `GOTO end_of_control_block`). However, when
executed in `AFTER DIALOG`, the focus returns to the current field or read-only list.
In this case the `BEFORE ROW` and `BEFORE FIELD` triggers will be
invoked.

A `CONTINUE DIALOG` in `AFTER FIELD`, `AFTER INPUT`,
`AFTER DISPLAY` or `AFTER CONSTRUCT` will only stop the program flow
of the current block of statements; instructions after `CONTINUE DIALOG` will not be
executed. If the user has selected a field in a different sub-dialog, this new field will get the
focus and all necessary AFTER / BEFORE control blocks will be executed.

In case of input error in a field, the best practice is to use a `NEXT FIELD`
instruction to stay in the dialog and set the focus to the field that the user has to correct.

## Related links

**Related concepts**  

[NEXT FIELD instruction](1955-next-field-instruction.md "NEXT FIELD instruction")
