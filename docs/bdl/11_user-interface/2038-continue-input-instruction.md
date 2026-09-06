---
title: "CONTINUE INPUT instruction"
source: "fgl-topics/c_fgl_InputArray_CONTINUE_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY control instructions > CONTINUE INPUT instruction"
type: "concept"
description: "Syntax CONTINUE INPUT Usage CONTINUE INPUT skips all subsequent statements in the current control block and gives the control back to the dialog. This instruction is useful when program control is ..."
---

# CONTINUE INPUT instruction

## Syntax

```
CONTINUE INPUT
```

## Usage

`CONTINUE INPUT` skips all subsequent statements in the current control block and
gives the control back to the dialog. This instruction is useful when program control is nested
within multiple conditional statements, and you want to return the control to the dialog.

If this instruction is called in a control block that is not [`AFTER INPUT`](1941-after-input-block.md), further control blocks
might be executed depending on the context.

`CONTINUE INPUT` instructs the dialog to continue as if the code in the control
block was terminated (it's a kind of `GOTO end_of_control_block`). However, when
executed in `AFTER INPUT`, the focus returns to the current row and current field in
the list, giving the user another chance to enter data in that field. In this case the
`BEFORE ROW` and `BEFORE FIELD` triggers will be invoked.

In this example, an `ON ACTION` block gives control back to the dialog, skipping
all subsequence instructions of the `ON ACTION` block:

```
ON ACTION zoom 
   IF p_cust.cust_id IS NULL OR p_cust.cust_name IS NULL THEN
      ERROR "Zoom window cannot be opened."
      CONTINUE INPUT
   END IF
   IF p_cust.cust_address IS NULL THEN
      ...
```

Use the `NEXT FIELD` control
instruction to give the focus to a specific field and force the dialog to continue. However, unlike
`CONTINUE INPUT`, the `NEXT FIELD` instruction will also skip the
further control blocks that are normally executed.

The `CONTINUE INPUT` instruction can only be used in a singular `INPUT
ARRAY` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog
block.

## Related links

**Related concepts**  

[AFTER FIELD block](1944-after-field-block.md "AFTER FIELD block")

[The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")
