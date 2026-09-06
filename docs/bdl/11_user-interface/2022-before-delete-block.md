---
title: "BEFORE DELETE block"
source: "fgl-topics/c_fgl_dialog_BEFORE_DELETE.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY control blocks > BEFORE DELETE block"
type: "concept"
description: "Syntax BEFORE DELETE instruction [...] Usage The BEFORE DELETE block is executed each time the user deletes a row of an INPUT ARRAY list, before the row is removed from the list. You typically code ..."
---

# BEFORE DELETE block

## Syntax

```
BEFORE DELETE
   instruction [...]
```

## Usage

The `BEFORE DELETE` block is executed each time the user deletes a row of an
`INPUT ARRAY` list, before the row is removed from the list.

You typically code the database table synchronization in the `BEFORE DELETE`
block, by executing a DELETE SQL statement using the primary key of the current row. In the
`BEFORE DELETE` block, the row to be deleted still exists in the program array, so
you can access its data to identify what record needs to be removed.

The `BEFORE DELETE` block is executed before the `AFTER
DELETE` block.

If needed, the deletion can be canceled with the [`CANCEL DELETE`](2036-cancel-delete-instruction.md) instruction.

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") returns the index of the row that will be deleted.

This example uses the `BEFORE DELETE` block to remove the row from the database
table and cancels the deletion operation if an SQL error occurs:

```
INPUT ARRAY p_items FROM s_items.*
  BEFORE DELETE
    LET r = DIALOG.getCurrentRow("s_items")
    WHENEVER ERROR CONTINUE
    DELETE FROM items
       WHERE item_num = p_items[r].item_num
    WHENEVER ERROR STOP
    IF sqlca.sqlcode <> 0 THEN
      ERROR SQLERRMESSAGE
      CANCEL DELETE
    END IF
...
```

## Related links

**Related concepts**  

[AFTER DELETE block](2023-after-delete-block.md "AFTER DELETE block")
