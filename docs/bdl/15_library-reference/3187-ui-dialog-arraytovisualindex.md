---
title: "ui.Dialog.arrayToVisualIndex"
source: "fgl-topics/c_fgl_ClassDialog_arrayToVisualIndex.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.arrayToVisualIndex"
type: "concept"
---

# ui.Dialog.arrayToVisualIndex

> Converts the program array index to the visual index for a given screen array.

## Syntax

```
arrayToVisualIndex(
   name STRING,
   arrayIndex INTEGER
 ) RETURNS INTEGER
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. arrayIndex is the index of the program array row.

## Usage

When the end user sorts rows in a table, the program array index may differ from the
visual row index.

Use this method to convert a program array row index (`arr_curr()`) to
a row index as seen by the end user. For example, if you want to display a typical
message with (current-row / total-rows),
convert the current program array row to a visual row index before displaying the
value:

```
MESSAGE SFMT( "Row: %1/%2",
   DIALOG.arrayToVisualIndex( "sr", DIALOG.getCurrentRow("sr") ),
   DIALOG.getArrayLength( "sr" )
 )
```

## Related links

**Related concepts**  

[Handling the current row](../11_user-interface/2303-handling-the-current-row.md "Query and control the current row in a read-only or editable list of records.")

[ON SORT block](../11_user-interface/1989-on-sort-block.md "ON SORT block")

[ui.Dialog.visualToArrayIndex](3234-ui-dialog-visualtoarrayindex.md "Converts the visual index to the program array index for a given screen array.")

[ui.Dialog.getCurrentRow](3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.")

[ui.Dialog.getArrayLength](3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.")
