---
title: "ui.Dialog.visualToArrayIndex"
source: "fgl-topics/c_fgl_ClassDialog_visualToArrayIndex.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.visualToArrayIndex"
type: "concept"
---

# ui.Dialog.visualToArrayIndex

> Converts the visual index to the program array index for a given screen array.

## Syntax

```
visualToArrayIndex(
   name STRING,
   visualIndex INTEGER
 ) RETURNS INTEGER
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. visualIndex is the index of the row as seen by the end user.

## Usage

When the end user sorts rows in a table, the visual row index may differ from the
program array index.

Use this method to convert a row index as seen by the end user, to the program array index. For
example, if the application implements a feature that allows the user to enter a row
index to jump to that row, it will be entered as a visual row index. You must
convert this index to the program array index, for example to make a
`setCurrentRow()`.

```
CALL DIALOG.setCurrentRow("sr", DIALOG.visualToArrayIndex("sr", user_index))
```

## Related links

**Related concepts**  

[Handling the current row](../11_user-interface/2303-handling-the-current-row.md "Query and control the current row in a read-only or editable list of records.")

[ui.Dialog.arrayToVisualIndex](3187-ui-dialog-arraytovisualindex.md "Converts the program array index to the visual index for a given screen array.")
