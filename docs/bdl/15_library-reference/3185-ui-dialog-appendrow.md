---
title: "ui.Dialog.appendRow"
source: "fgl-topics/c_fgl_ClassDialog_appendRow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.appendRow"
type: "concept"
---

# ui.Dialog.appendRow

> Appends a new row in the specified list.

## Syntax

```
appendRow(
   name STRING )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

The `appendRow()` method appends a row to the end of the array controlled by the
dialog.

> **Important:**
>
> This method is designed to be used in an
> `ON ACTION` block. It must not be called in control blocks such as `BEFORE
> ROW`, `AFTER ROW`, `BEFORE INSERT`, `AFTER
> INSERT`, `BEFORE DELETE`, or `AFTER DELETE`.

The method is similar to appending a new element to the program array, except the internal
dialog registers are automatically updated (like the total number of rows returned by [`getArrayLength()`](3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.")). If the list
is decorated with [cell attributes](3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs)."), the
program array defining the attributes will also be synchronized. If [multi-row selection](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") is enabled, selection
flags of existing rows are kept. The new row is inserted at the end of the list with the selection
flag set to zero.

> **Note:**
>
> The purpose of this method is to implement business logic required to modify the record list in
> the current dialog. It is typically used in a `DISPLAY ARRAY` dialog. Avoid using
> this method in `INPUT ARRAY`. To allow the end user to append, modify or delete rows
> in a `DISPLAY ARRAY`, use [list
> modification interaction blocks](../11_user-interface/2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.").

After the method is called, a new row is created in the program array. You can assign values to
the variables before the control goes back to the user. The `getArrayLength()` method
will return the new row count.

The method does not set the current row and does not give the focus to the list; you need to call
`setCurrentRow()` and execute [`NEXT
FIELD`](../11_user-interface/1955-next-field-instruction.md) to give the focus.

This method does not execute any `BEFORE ROW`, `BEFORE INSERT`,
`AFTER INSERT` or `AFTER ROW` control blocks.

The `appendRow()` method does not create a [temporary row](../11_user-interface/2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") as the "append" action of `INPUT
ARRAY`; the row is considered permanent once it is added.

## Example

This example implements a user-defined action to append ten rows at the end of the
list.

```
ON ACTION append_some_rows 
  FOR i = 1 TO 10
    CALL DIALOG.appendRow("sa")
    LET r = DIALOG.getArrayLength("sa")
    LET p_items[r].item_quantity = 1.00
  END FOR
```
