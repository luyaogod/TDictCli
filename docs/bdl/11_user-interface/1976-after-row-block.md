---
title: "AFTER ROW block"
source: "fgl-topics/c_fgl_dialog_AFTER_ROW.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY control blocks > AFTER ROW block"
type: "concept"
description: "Syntax AFTER ROW instruction [...] AFTER ROW block in singular DISPLAY ARRAY, INPUT ARRAY dialogs In a singular DISPLAY ARRAY or INPUT ARRAY instruction, the AFTER ROW block is executed each time the ..."
---

# AFTER ROW block

## Syntax

```
AFTER ROW
   instruction [...]
```

## AFTER ROW block in singular DISPLAY ARRAY, INPUT ARRAY dialogs

In a singular `DISPLAY ARRAY` or `INPUT ARRAY` instruction, the
`AFTER ROW` block is executed each time the user moves to another row, before the
current row is left. This trigger can also be executed in other situations, such as when you delete
a row, or when the user inserts a new row.

A `NEXT FIELD` instruction executed in the `AFTER ROW`
control block will keep the user entry in the current row. Use this behavior to implement
row validation and prevent the user from leaving the list or moving to another row.

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") returns the index of the row that you are leaving.

## AFTER ROW block in DISPLAY ARRAY and INPUT ARRAY of procedural DIALOG

In an `INPUT` or `INPUT ARRAY` sub-dialog of a procedural
`DIALOG` instruction, the `AFTER ROW` block is executed when a
`DISPLAY ARRAY` or `INPUT ARRAY` list loses the focus, or when the
user moves to another row in a list. This trigger can also be executed in other situations, for
example when you delete a row, or when the user inserts a new row.

`AFTER ROW` is executed after the `AFTER
FIELD`, `AFTER INSERT` and before
`AFTER DISPLAY` or `AFTER INPUT` blocks.

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") returns the index of the of the row that you are leaving.

For both `INPUT ARRAY` and `DISPLAY ARRAY` sub-dialogs,
a `NEXT FIELD` executed in the `AFTER
ROW` control block will keep the focus in the list and
stay in the current row. Use this feature to implement row validation
and prevent the user from leaving the list or moving to another row.

## AFTER ROW and temporary rows in INPUT ARRAY

After creating a [temporary row](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") at the end of
a list driven by INPUT ARRAY, if you leave that row to go to a previous row without data input
(setting the touched flag), or when the cancel action is invoked, the temporary row will be
automatically removed. The `AFTER ROW` block will be executed for the temporary row,
but `ui.Dialog.getCurrentRow()`/`arr_curr()` will be one row greater
than `ui.Dialog.getArrayLength()`/`ARR_COUNT()`. In this case, it is
recommended that you ignore the `AFTER ROW` event. For example, it is recommended
that you avoid executing a `NEXT FIELD` or `CONTINUE INPUT`
instruction, and trying to access the dynamic array with a row index that is greater than the total
number of rows, otherwise the runtime system will adapt the total number of rows to the actual
number of rows in the program array.

In this example, the `AFTER ROW` block checks the current row index and verifies a
variable value to force the focus to stay in the current row if the value is
wrong:

```
INPUT ARRAY p_items FROM s_items.*
   ...
   AFTER ROW
      LET r = DIALOG.getCurrentRow("s_items")
      IF r <= DIALOG.getArrayLength("s_items") THEN
        IF NOT item_is_valid_quantity(p_item[r].item_quantity) THEN
           ERROR "Item quantity is not valid"
           NEXT FIELD item_quantity
         END IF
      END IF
```

Another way to handle the case of temporary rows in `AFTER ROW` is to use a flag
to know if the `AFTER INSERT` block was executed. The `AFTER INSERT`
block is not executed if the temporary row is automatically removed. By setting a first value in
`BEFORE INSERT` and changing the flag in `AFTER INSERT`, you can
detect if the row was permanently added to the
list:

```
INPUT ARRAY p_items FROM s_items.*
  ...
  BEFORE INSERT
    LET op = "T"
    ...
  AFTER INSERT
    LET op = "I"
    ...
  AFTER ROW
    IF op == "I" THEN
      IF NOT item_is_valid_quantity(p_item[arr_curr()].item_quantity) THEN
         ERROR "Item quantity is not valid"
         NEXT FIELD item_quantity
        END IF
      WHENEVER ERROR CONTINUE
      INSERT INTO items (item_num, item_name, item_quantity) 
                         VALUES ( p_item[arr_curr()].* )
      WHENEVER ERROR STOP
      IF sqlca.sqlcode<0 THEN
         ERROR "Could not insert the record into database!"
         NEXT FIELD CURRENT
      ELSE
         MESSAGE "Record has been inserted successfully"
      END IF
   END IF
...
```

## Related links

**Related concepts**  

[NEXT FIELD instruction](1955-next-field-instruction.md "NEXT FIELD instruction")

[BEFORE ROW block](1975-before-row-block.md "BEFORE ROW block")

[ON ROW CHANGE block](2018-on-row-change-block.md "ON ROW CHANGE block")
