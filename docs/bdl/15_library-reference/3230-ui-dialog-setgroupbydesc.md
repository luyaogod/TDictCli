---
title: "ui.Dialog.setGroupByDesc"
source: "fgl-topics/c_fgl_ClassDialog_setGroupByDesc.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setGroupByDesc"
type: "concept"
---

# ui.Dialog.setGroupByDesc

> Defines sort order for a grouping column of a list dialog.

## Syntax

```
setGroupByDesc(
  name STRING,
  val BOOLEAN
)
```

1. name is the name of form field of a [list dialog](../11_user-interface/2300-understanding-list-dialogs.md "List dialogs are dialogs controlling a list of records rendered in a list container such as TABLE, TREE or SCROLLGRID."), defined to group rows. See [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).
2. val is a boolean value: `TRUE` for descending order,
   `FALSE` for ascending order (default when defining a grouping column).

## Usage

After defining a grouping column with the [`setGroupBy()`](3229-ui-dialog-setgroupby.md "Defines the grouping column of a list dialog.") or [`addGroupBy()`](3181-ui-dialog-addgroupby.md "Appends a grouping column of a list dialog.") method, use `setGroupByDesc()` to define the
sort order for the grouping column passed as first argument.

By default, the sort order is ascending. If the second parameter is `TRUE`, the
sort order is descending.

For more details and code example, read the [ui.Dialog.addGroupBy](3181-ui-dialog-addgroupby.md "Appends a grouping column of a list dialog.")
reference page.
