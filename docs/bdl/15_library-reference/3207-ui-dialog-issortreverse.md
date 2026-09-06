---
title: "ui.Dialog.isSortReverse"
source: "fgl-topics/c_fgl_ClassDialog_isSortReverse.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.isSortReverse"
type: "concept"
---

# ui.Dialog.isSortReverse

> Indicates the sort order direction (FALSE=ascending, TRUE=descending)

## Syntax

```
isSortReverse(
   name STRING )
  RETURNS BOOLEAN
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

The `isSortReverse()` method returns `FALSE` if the
sort order is ascending, and `TRUE` if the sort is in descending order.

This method is used in the context of the `ON SORT` trigger.

To identify multiple sort columns, use [`isSortKeyReverseAt()`](3208-ui-dialog-issortreverseat.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending), for a given sort column position.").

## Related links

**Related concepts**  

[ON SORT block](../11_user-interface/1989-on-sort-block.md "ON SORT block")

[Populating a DISPLAY ARRAY](../11_user-interface/2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.")

[ui.Dialog.getSortKey](3202-ui-dialog-getsortkey.md "Returns the name of the sort field selected by the user.")
