---
title: "ui.Dialog.getSortKeyAt"
source: "fgl-topics/c_fgl_ClassDialog_getSortKeyAt.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getSortKeyAt"
type: "concept"
---

# ui.Dialog.getSortKeyAt

> Returns the name of field used as grouping or sorting column, for a given sort column position.

## Syntax

```
getSortKeyAt(
   name STRING,
   index INTEGER
)
  RETURNS STRING
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. index is the sort column position (1 = first sort column)

## Usage

The `getSortKeyAt()` method returns the form field name (table column) used as
sorting column, for the record list identified by the screen record name, and the
sort column position specified by index parameter.

This method is used in the context of the `ON SORT` trigger.

When the `ON SORT` trigger is fired because the sort order was reset, the
`getSortKeyAt()` method returns `NULL`.

## Related links

**Related concepts**  

[ui.Dialog.getSortKey](3202-ui-dialog-getsortkey.md "Returns the name of the sort field selected by the user.")

[ON SORT block](../11_user-interface/1989-on-sort-block.md "ON SORT block")

[Populating a DISPLAY ARRAY](../11_user-interface/2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.")

[ui.Dialog.isSortReverse](3207-ui-dialog-issortreverse.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending)")
