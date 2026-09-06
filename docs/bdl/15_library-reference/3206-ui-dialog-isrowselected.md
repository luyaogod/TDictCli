---
title: "ui.Dialog.isRowSelected"
source: "fgl-topics/c_fgl_ClassDialog_isRowSelected.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.isRowSelected"
type: "concept"
---

# ui.Dialog.isRowSelected

> Queries row selection for a given list and row.

## Syntax

```
isRowSelected(
   name STRING,
   row INTEGER )
  RETURNS BOOLEAN
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. row is a row index.

## Usage

If multi-row selection is enabled with [`setSelectionMode()`](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list."),
you can check whether a row is selected with the `isRowSelected()`
method:

```
ON ACTION check_current_row_selected 
  IF DIALOG.isRowSelected( "sr", DIALOG.getCurrentRow("sr") ) THEN
    MESSAGE "Current row is selected."
  END IF
```

If multi-row selection is off, the method
returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") for
the current row and [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") for
other rows.
