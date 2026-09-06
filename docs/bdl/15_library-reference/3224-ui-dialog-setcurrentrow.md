---
title: "ui.Dialog.setCurrentRow"
source: "fgl-topics/c_fgl_ClassDialog_setCurrentRow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setCurrentRow"
type: "concept"
---

# ui.Dialog.setCurrentRow

> Sets the current row in the specified list.

## Syntax

```
setCurrentRow(
   name STRING,
   row INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. row is the new row in the array.

## Usage

Use the `setCurrentRow()` method to change the current row in an `INPUT
ARRAY` or `DISPLAY ARRAY` list. You must pass the name of the screen array to
identify the list, and the new row
number.

```
DEFINE x INTEGER
DIALOG
  DISPLAY ARRAY custlist TO sa_custlist.*
    ...
  END DISPLAY
  ON ACTION goto_x 
    CALL DIALOG.setCurrentRow("sa_custlist", x)
    ...
```

Moving to a different row with `setCurrentRow()` will not trigger [control blocks](../11_user-interface/2099-dialog-control-blocks.md "Dialog control blocks are predefined dialog triggers where you can implement specific code to control the interactive instruction.") such as `BEFORE ROW` /
`AFTER ROW`, as the [`fgl_set_arr_curr()`](2782-fgl-set-arr-curr.md "Moves to a specific row in a record list.") built-in function does.

The `setCurrentRow()` method will not set the focus; you need to use `NEXT
FIELD` to set the focus to a list. (This works with `DISPLAY ARRAY` as well
as with `INPUT ARRAY`.)

If the passed row index is lower than 1, the first row will be selected. If the row index is
greater than the total number of rows, the last row will be selected.

If the new current row is not in the current view, the dialog will adapt the list offset to make
the new current row visible.

If [multi-row selection](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") is enabled,
all selection flags of rows are cleared, and the new current row gets automatically selected.

A `setCurrentRow(screen-array,row-index)` in
a [paged mode `DISPLAY ARRAY`](../11_user-interface/2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.")
will subsequently trigger the `ON FILL BUFFER` code, if the requested row is not in
the current page of visible rows.

If `DISPLAY ARRAY` using `ON FILL BUFFER` was started with
`COUNT=-1`, and the row-index provided by
`setCurrentRow()` is greater than the actual number of rows already fetched through
`ON FILL BUFFER`, `setCurrentRow()` will have no effect. The actual
number of rows must be provided with [`DIALOG.setArrayLength(screen-array,count)`](3220-ui-dialog-setarraylength.md "Sets the number of rows in a DISPLAY ARRAY using paged mode.")
where count >= row-index, before calling
`setCurrentRow()`:

```
DISPLAY ARRAY ... ATTRIBUTES(COUNT=-1)
ON FILL BUFFER
   ...
ON ACTION goto_last_row
   SELECT COUNT(*) INTO cnt FROM ... -- Count total rows
   CALL DIALOG.setArrayLength("sr", cnt)
   CALL DIALOG.setCurrentRow("sr", cnt)
...
```
