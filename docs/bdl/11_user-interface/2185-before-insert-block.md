---
title: "BEFORE INSERT block"
source: "fgl-topics/c_fgl_dialog_BEFORE_INSERT_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control blocks > BEFORE INSERT block"
type: "concept"
description: "Syntax BEFORE INSERT instruction [...] Usage The BEFORE INSERT block is executed when a new row is created in an INPUT ARRAY . You typically use this trigger to set some default values in the ..."
---

# BEFORE INSERT block

## Syntax

```
BEFORE INSERT
   instruction [...]
```

## Usage

The `BEFORE INSERT` block is executed when a new row is created in an
`INPUT ARRAY`. You typically use this trigger to set some default values in the
newly-created row. A new row can be created by moving down after the last row, by executing a insert
action, or by executing an append action.

The `BEFORE INSERT` block is executed after the `BEFORE ROW`
block and before the `BEFORE FIELD` block.

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") returns the index of the newly-created row.

To distinguish row insertion from an appended row, compare the current row
(`DIALOG.getCurrentRow("screen-array")`) with the total number of rows
(`DIALOG.getArrayLength("screen-array")`). If the current row index and the
total number of rows correspond, the `BEFORE INSERT` concerns a temporary row,
otherwise it concerns an inserted row.

Row creation can be stopped by using the [`CANCEL INSERT`](2037-cancel-insert-instruction.md) instruction inside `BEFORE INSERT`. If
possible, it is however better to disable the insert and append actions to prevent the user
executing the actions with [`DIALOG.setActionActive()`](../15_library-reference/3213-ui-dialog-setactionactive.md "Enabling and disabling dialog actions.").

When inserting/appending a new row, the fields are initialized with the default values,
before the `BEFORE INSERT` code is executed. In an `INPUT
ARRAY`, the default values for new rows are taken from the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attributes of the form
fields.

In this example, the `BEFORE INSERT` block checks if the user can create rows and
prevents new row creation if needed; otherwise, it sets some default values:

```
INPUT ARRAY p_items FROM s_items.*
    ...
    BEFORE INSERT
      IF NOT user_can_append THEN
         ERROR "You are not allowed to append rows"
         CANCEL INSERT
      END IF
      LET r = DIALOG.getCurrentRow("s_items")
      LET p_items[r].item_num = get_new_serial("items")
      LET p_items[r].item_name = "undefined"
```

## Related links

**Related concepts**  

[Appending rows in INPUT ARRAY](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.")

[BEFORE ROW block](1975-before-row-block.md "BEFORE ROW block")
