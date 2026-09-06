---
title: "ui.Dialog.insertRow"
source: "fgl-topics/c_fgl_ClassDialog_insertRow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.insertRow"
type: "concept"
---

# ui.Dialog.insertRow

> Inserts a new row in the specified list.

## Syntax

```
insertRow(
   name STRING,
   index INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. index is the index where the row must be inserted
   (starts at 1).

## Usage

The `insertRow()` method inserts a row in the list, at a given position.

> **Important:**
>
> This method is designed to be used in an
> `ON ACTION` block. It must not be called in control blocks such as `BEFORE
> ROW`, `AFTER ROW`, `BEFORE INSERT`, `AFTER
> INSERT`, `BEFORE DELETE`, or `AFTER DELETE`.

The method is similar to inserting a new element in the program array, except the internal
dialog registers are automatically updated (like the total number of rows returned by [`getArrayLength()`](3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.")). If the list
is decorated with [cell attributes](3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs)."), the
program array defining the attributes will also be synchronized. If [multi-row selection](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") is enabled, selection
flags of existing rows are kept. Selection information is synchronized (i.e., flags are shifted
down) for all rows after the new inserted row.

> **Note:**
>
> The purpose of this method is to implement business logic required to modify the record list in
> the current dialog. It is typically used in a `DISPLAY ARRAY` dialog. Avoid using
> this method in `INPUT ARRAY`. To allow the end user to append, modify or delete rows
> in a `DISPLAY ARRAY`, use [list
> modification interaction blocks](../11_user-interface/2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.").

After the method is called, a new row is created in the program array, so you can assign values
to the variables before the control goes back to the user. The `getArrayLength()`
method will return the new row count.

The method does not set the current row and does not give the focus to the list; you need to call
`setCurrentRow()` and execute [`NEXT
FIELD`](../11_user-interface/1955-next-field-instruction.md) to give the focus.

The `insertRow()` method must not be used when controlling a tree view. Use the
[`insertNode()`](3204-ui-dialog-insertnode.md "Inserts a new node in the specified tree.") method
instead.

This method does not execute any `BEFORE ROW`, `BEFORE INSERT`,
`AFTER INSERT` or `AFTER ROW` control blocks.

If the index is greater than the number of rows, a new row is appended at the end of the list.
This is the equivalent of calling the `appendRow()` method.

If the list is empty, `getCurrentRow()` returns zero. If zero is returned, use 1
to reference the first row, otherwise you can get a [-1326](4483-genero-bdl-errors.md) error when using the program
array.

## Example

This example shows a user-defined action to insert ten rows in the list at the current
position:

```
ON ACTION insert_some_rows 
  LET r = DIALOG.getCurrentRow("sa")
  IF r == 0 THEN LET r = 1 END IF
  FOR i = 10 TO 1 STEP -1
    CALL DIALOG.insertRow("sa", r)
    LET p_items[r].item_quantity = 1.00
  END FOR
```
