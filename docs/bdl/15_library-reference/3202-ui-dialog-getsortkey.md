---
title: "ui.Dialog.getSortKey"
source: "fgl-topics/c_fgl_ClassDialog_getSortKey.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getSortKey"
type: "concept"
---

# ui.Dialog.getSortKey

> Returns the name of the sort field selected by the user.

## Syntax

```
getSortKey(
   name STRING )
  RETURNS STRING
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

The `getSortKey()` method returns the form field name (table column) selected by
the user to sort rows.

This method is used in the context of the `ON SORT` trigger.

When the `ON SORT` trigger is fired because the sort order was reset, the
`getSortKey()` method returns `NULL`.

To identify multiple sort columns, use [`getSortKeyAt()`](3203-ui-dialog-getsortkeyat.md "Returns the name of field used as grouping or sorting column, for a given sort column position.").

## Related links

**Related concepts**  

[ui.Dialog.getSortKeyAt](3203-ui-dialog-getsortkeyat.md "Returns the name of field used as grouping or sorting column, for a given sort column position.")

[ON SORT block](../11_user-interface/1989-on-sort-block.md "ON SORT block")

[Populating a DISPLAY ARRAY](../11_user-interface/2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.")

[ui.Dialog.isSortReverse](3207-ui-dialog-issortreverse.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending)")
