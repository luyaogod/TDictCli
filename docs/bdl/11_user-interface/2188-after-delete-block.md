---
title: "AFTER DELETE block"
source: "fgl-topics/c_fgl_dialog_AFTER_DELETE_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control blocks > AFTER DELETE block"
type: "concept"
description: "Syntax AFTER DELETE instruction [...] Usage The AFTER DELETE block is executed each time the user deletes a row of an INPUT ARRAY list, after the row has been deleted from the list. The AFTER DELETE ..."
---

# AFTER DELETE block

## Syntax

```
AFTER DELETE
   instruction [...]
```

## Usage

The `AFTER DELETE` block is executed each time the user deletes a row of an
`INPUT ARRAY` list, after the row has been deleted from the list.

The `AFTER DELETE` block is executed after the `BEFORE
DELETE` block and  before the `AFTER ROW` block for the deleted row
and the `BEFORE ROW` block of the new current row.

When an `AFTER DELETE` block executes, the program array has already been
modified; the deleted row no longer exists in the array (except in the special case when deleting
the last row). The [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") function or the [`ui.Dialog.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") method
returns the same index as in `BEFORE ROW`, but it is the index of the new current
row. The `AFTER ROW` block is also executed just after the `AFTER
DELETE` block.

When deleting the last row of the list, `AFTER DELETE` is executed for the delete
row, and `DIALOG.getCurrentRow()` / `arr_curr()` will be one greater
than `DIALOG.getArrayLength()` / `ARR_COUNT()`. Ensure you avoid
accessing a dynamic array with a row index that is greater than the total number of rows, otherwise
the runtime system will adapt the total number of rows to the actual number of rows in the program
array. When using a static array, you must ignore the values in the rows after
`ARR_COUNT()`.

Here the `AFTER DELETE` block is used to re-number the rows with a new item line
number (note that [`DIALOG.getArrayLength()`](../15_library-reference/3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.") / [`ARR_COUNT()`](../15_library-reference/2727-arr-count.md "Returns the number of rows entered during an INPUT ARRAY statement.") may return
zero):

```
INPUT ARRAY p_items FROM s_items.*
   AFTER DELETE
      LET r = DIALOG.getCurrentRow("s_items")
      FOR i=r TO DIALOG.getArrayLength("s_items")
          LET p_items[i].item_lineno = i
      END FOR
...
```

It is not possible to use the `CANCEL DELETE` instruction in an `AFTER
DELETE` block. At this time it is too late to cancel row deletion, as the data row no longer
exists in the program array.

## Related links

**Related concepts**  

[BEFORE INSERT block](2020-before-insert-block.md "BEFORE INSERT block")

[AFTER ROW block](1976-after-row-block.md "AFTER ROW block")
