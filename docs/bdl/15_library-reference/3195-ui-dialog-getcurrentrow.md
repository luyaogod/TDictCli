---
title: "ui.Dialog.getCurrentRow"
source: "fgl-topics/c_fgl_ClassDialog_getCurrentRow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getCurrentRow"
type: "concept"
---

# ui.Dialog.getCurrentRow

> Returns the current row of the specified list.

## Syntax

```
getCurrentRow(
   name STRING )
  RETURNS INTEGER
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

Use the `getCurrentRow()` method to retrieve the current row of
an `INPUT ARRAY` or `DISPLAY ARRAY` list.

You must pass the name of the screen array to identify the list.

```
DIALOG
  DISPLAY ARRAY custlist TO sa_custlist.*
    BEFORE ROW
      MESSAGE "Current row: " || DIALOG.getCurrentRow("sa_custlist")
    ...
  END DISPLAY
  INPUT ARRAY ordlist TO sa_ordlist.*
    BEFORE ROW
      MESSAGE "Current row: " || DIALOG.getCurrentRow("sa_ordlist")
    ...
  END INPUT
  ...
```
