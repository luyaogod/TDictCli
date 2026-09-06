---
title: "AFTER INSERT block"
source: "fgl-topics/c_fgl_dialog_AFTER_INSERT_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > AFTER INSERT block"
type: "concept"
description: "Syntax AFTER INSERT instruction [...] Usage The AFTER INSERT block of INPUT ARRAY is executed when the creation of a new row is validated. In this block, you can for example implement SQL to insert a ..."
---

# AFTER INSERT block

## Syntax

```
AFTER INSERT
   instruction [...]
```

## Usage

The `AFTER INSERT` block of `INPUT ARRAY` is executed when the
creation of a new row is validated. In this block, you can for example implement SQL to insert a new
row in the database table.

The `AFTER INSERT` block is executed after the `AFTER FIELD`
block and before the `AFTER ROW` block.

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") returns the index of the newly-created row.

When the user appends a new row at the end of the list, then moves UP to another row or validates
the dialog, the `AFTER INSERT` block is only executed if at least one field was
edited. If no data entry is detected, the dialog automatically removes the new appended row and thus
does not trigger the `AFTER INSERT` block.

When executing a `NEXT FIELD` in the `AFTER INSERT` block, the
dialog will keep the focus in the list and stay in the current row. Use this behavior to implement
row input validation and prevent the user from leaving the list or moving to another row. However,
this will not cancel the row insertion and will not invoke the `BEFORE INSERT` /
`AFTER INSERT` triggers again. The only way to keep the focus in the current row
after the row was inserted is to execute a `NEXT FIELD` in the `AFTER
ROW` block.

In this example, the `AFTER INSERT` block inserts a new row in the database and
cancels the operation if the SQL command fails:

```
INPUT ARRAY p_items FROM s_items.*
    ...
    AFTER INSERT
      WHENEVER ERROR CONTINUE
      INSERT INTO items VALUES ( p_items[DIALOG.getCurrentRow("s_items")].* )
      WHENEVER ERROR STOP
      IF sqlca.sqlcode<>0 THEN
         ERROR SQLERRMESSAGE
         CANCEL INSERT
      END IF
```

## Related links

**Related concepts**  

[NEXT FIELD instruction](1955-next-field-instruction.md "NEXT FIELD instruction")

[AFTER ROW block](1976-after-row-block.md "AFTER ROW block")
