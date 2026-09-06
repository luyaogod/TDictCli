---
title: "ui.Dialog.setSelectionRange"
source: "fgl-topics/c_fgl_ClassDialog_setSelectionRange.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setSelectionRange"
type: "concept"
---

# ui.Dialog.setSelectionRange

> Sets the row selection flags for a range of rows.

## Syntax

```
setSelectionRange(
   name STRING,
   start INTEGER,
   end INTEGER,
   value BOOLEAN )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. start is the starting row index.
3. end is the ending row index.
4. value is the selection flag to set.

## Usage

If multi-row selection is enabled
with [setSelectionMode()](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list."),
you can set the selection flags for a range of rows with the `setSelectionRange()`
method.

```
ON ACTION select_all 
  CALL DIALOG.setSelectionRange( "sr", 1, -1, TRUE)
```

The
start and end index must be in the range of possible row indexes (from
1 to `DIALOG.getArrayLength()`).

If
you specify an end index of -1, it will set the flags from start index
to the end of the list.
