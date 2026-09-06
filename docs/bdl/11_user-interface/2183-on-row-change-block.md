---
title: "ON ROW CHANGE block"
source: "fgl-topics/c_fgl_dialog_ON_ROW_CHANGE_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control blocks > ON ROW CHANGE block"
type: "concept"
description: "Syntax ON ROW CHANGE instruction [...] Usage The ON ROW CHANGE block is executed in a list controlled by an INPUT ARRAY , when leaving the current row and when the row has been modified since it got ..."
---

# ON ROW CHANGE block

## Syntax

```
ON ROW CHANGE
   instruction [...]
```

## Usage

The `ON ROW CHANGE` block is executed in a list controlled by an `INPUT
ARRAY`, when leaving the current row and when the row has been modified since it got the
focus. This is typically used to detect row modification.

The code in `ON ROW CHANGE` will not be executed when leaving new rows created by
the user with the default append or insert action. To detect row creation, you must use the
`BEFORE INSERT` or `AFTER INSERT` control blocks.

The `ON ROW CHANGE` block is only executed if at least one field value in the
current row has changed since the row was entered, and the modification flag of the field is set.
The modified field(s) may not be the current field, and several field values can be changed. Values
may have been changed by the user or by the program. The modification flag is reset for all fields
when entering another row, when going to another sub-dialog, or when leaving the dialog
instruction.

`ON ROW CHANGE` is executed after the `AFTER FIELD` block
and before the `AFTER ROW` block.

When called in this block, [DIALOG.getCurrentRow()](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") return the index of the current row that has been changed.

You can, for example, code database modifications (`UPDATE`) in the `ON ROW
CHANGE` block:

```
 INPUT ARRAY p_items FROM s_items.*
    ...
    ON ROW CHANGE
      LET r = DIALOG.getCurrentRow("s_items")
      UPDATE items SET
              items.item_code        = p_items[r].item_code,
              items.item_description = p_items[r].item_description,
              items.item_price       = p_items[r].item_price,
              items.item_updatedate  = TODAY
             WHERE items.item_num = p_items[r].item_num
```

## Related links

**Related concepts**  

[Input field modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.")

[AFTER ROW block](1976-after-row-block.md "AFTER ROW block")
