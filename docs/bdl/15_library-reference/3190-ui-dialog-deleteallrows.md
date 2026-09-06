---
title: "ui.Dialog.deleteAllRows"
source: "fgl-topics/c_fgl_ClassDialog_deleteAllRows.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.deleteAllRows"
type: "concept"
---

# ui.Dialog.deleteAllRows

> Deletes all rows from the specified list.

## Syntax

```
deleteAllRows(
   name STRING )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

The `deleteAllRows()` method
removes all the rows of a list driven by a `DISPLAY ARRAY` or `INPUT
ARRAY`. This is equivalent to a `deleteRow()` call,
but instead of deleting one particular row, it removes all rows of
the specified list.

This method must not be called in control
blocks such as `BEFORE ROW`, `AFTER ROW`, `BEFORE
INSERT`, `AFTER INSERT`, `BEFORE
DELETE`, `AFTER DELETE`, it is designed to
be used in an `ON ACTION` block.

After
the method is called, all rows are deleted from the program array,
and the `getArrayLength()` method will return zero.

The
method takes the name of the screen-array as parameter.

If
the `deleteAllRows()` method is called during an `INPUT
ARRAY`, the dialog will automatically append a new [temporary row](../11_user-interface/2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") if the
focus is in the list, to let the user enter new data. When using `AUTO
APPEND = FALSE` attribute, no temporary row will be created
and the current row register will be automatically changed to make
sure that it will not be greater than the total number of rows.

If the `deleteAllRows()` method is called during an `INPUT ARRAY`
or `DISPLAY ARRAY` that has the focus, the `BEFORE ROW` control
block will be executed if you delete the current row. This is required to reset the internal state
of the dialog.

If the list was
decorated with [cell
attributes](3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs)."), the program array defining the attributes will
be cleared. If [multi-row
selection](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") is enabled, selection flags are cleared.
