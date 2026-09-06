---
title: "ui.Dialog.deleteRow"
source: "fgl-topics/c_fgl_ClassDialog_deleteRow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.deleteRow"
type: "concept"
---

# ui.Dialog.deleteRow

> Deletes a row from the specified list.

## Syntax

```
deleteRow(
   name STRING,
   index INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. index is the index of the row to be deleted
   (starts a 1).

## Usage

The `deleteRow()` method deletes the row in the array controlled by the
dialog.

> **Important:**
>
> This method is designed to be used in an
> `ON ACTION` block. It must not be called in control blocks such as `BEFORE
> ROW`, `AFTER ROW`, `BEFORE INSERT`, `AFTER
> INSERT`, `BEFORE DELETE`, or `AFTER DELETE`.

The method is similar to deleting an element to the program array, except that internal dialog
registers are automatically updated (like the total number of rows returned by [`getArrayLength()`](3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.")). If the list
is decorated with [cell attributes](3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs)."), the
program array defining the attributes will also be synchronized. If [multi-row selection](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") is enabled, selection
information is synchronized (selection flags are shifted up) for all rows after the deleted row.

> **Note:**
>
> The purpose of this method is to implement business logic required to modify the record list in
> the current dialog. It is typically used in a `DISPLAY ARRAY` dialog. Avoid using
> this method in `INPUT ARRAY`. To allow the end user to append, modify or delete rows
> in a `DISPLAY ARRAY`, use [list
> modification interaction blocks](../11_user-interface/2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.").

After the method is called, the row no longer exists in the program array, and the
`getArrayLength()` method will return the new row count.

If the `deleteRow()` method is called during an `INPUT ARRAY` that
has the focus, control blocks such as `BEFORE ROW` and `BEFORE
FIELD` will be executed, if you delete the current row. This is required to reset the
internal state of the dialog. However, the method does not execute any `BEFORE ROW`
or `AFTER ROW` control blocks in a `DISPLAY ARRAY` dialog.

If the `deleteRow()` method is called during an `INPUT ARRAY`, and
if no more rows are in the list after the call, the dialog will automatically append a new [temporary row](../11_user-interface/2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") if the focus is in the list, to let the
user enter new data. When using `AUTO APPEND = FALSE` attribute, no temporary row
will be created and the current row register will be automatically changed to make sure that it will
not be greater than the total number of rows.

If you pass zero as row index, the method does nothing (if no rows are in the list,
`getCurrentRow()` returns zero).

## Example

This example implements a user-defined action to remove rows that have a specific
property:

```
ON ACTION delete_invalid_rows 
  FOR r = 1 TO DIALOG.getArrayLength("sa")
    IF NOT s_orders[r].is_valid THEN
      CALL DIALOG.deleteRow("sa",r)
      LET r = r - 1
    END IF
  END FOR
```
