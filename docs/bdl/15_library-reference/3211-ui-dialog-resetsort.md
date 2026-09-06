---
title: "ui.Dialog.resetSort"
source: "fgl-topics/c_fgl_ClassDialog_resetSort.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.resetSort"
type: "concept"
---

# ui.Dialog.resetSort

> Resets the sort columns of a list dialog.

## Syntax

```
resetSort(
  name STRING )
```

1. name is the name of form field of a [list dialog](../11_user-interface/2300-understanding-list-dialogs.md "List dialogs are dialogs controlling a list of records rendered in a list container such as TABLE, TREE or SCROLLGRID."). See [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

The `resetSort()` method can be used to reset the grouping columns defined by
program ([`setGroupBy()`](3229-ui-dialog-setgroupby.md "Defines the grouping column of a list dialog.")), and
sort columns selected by the end-user during a `DISPLAY ARRAY` or `INPUT
ARRAY` list dialog.

By default, when a form is loaded and its corresponding dialog controller starts, the front-end
sends stored settings that can containing the last sort specification, which can trigger visual row
sorting.

The `resetSort()` method can for example be called when re-executing the SQL query
to fill data rows with an SQL `ORDER BY` clause, and get a default row sorting.

When using a `DISPLAY ARRAY` in [paged mode](../11_user-interface/2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY."), the `ON FILL BUFFER` block is executed to re-fill the current
page after calling `resetSort()`. Note that the `ON SORT` block is not
fired in this case.

## Example

The next code example resets the sort specification each time the DISPLAY ARRAY dialog
starts:

```
DISPLAY ARRAY files TO sr.*
    BEFORE DISPLAY
        CALL DIALOG.resetSort("sr")
    ...
END DISPLAY
```
