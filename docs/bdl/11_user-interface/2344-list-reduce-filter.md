---
title: "List reduce filter"
source: "fgl-topics/c_fgl_prog_dialogs_list_filter_2.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > List reduce filter"
type: "concept"
---

# List reduce filter

> The reduce filter allows a user to limit the row set in the list by using a filter.

## Understanding the reduce filter

When using a [`DISPLAY
ARRAY`](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.") with a [`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") or [`SCROLLGRID`](1723-scrollgrid-container.md "Defines a scrollable grid view widget.") container, the user can enter a criterion in a search field, to
show only the rows matching the content of the filter.

The reduce filter is supported on mobile and desktop front-end platforms.

To open the reduce filter field, click/tap on the filtering icon in the [chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

![Android list view with filter field](../_images/android_listview_filter_gbc.jpg)

*Android™ list view with filter field*

## Reduce filter usage details

The filter search is case-insensitive.

The value entered in the filter field is compared to all fields of visible columns, except
columns of the type `TEXT` or `BYTE`. The comparison is based on the
formatted value. For example, a `MONEY` column will display values formatted with the
currency symbol. To match values in that column, the user must enter exactly the same value (for
example, with the currency symbol and the correct decimal separator). When using
`COMBOBOX` fields, the find searches in the visible values of combobox items.

Only text widgets displaying values are searched. Columns using widgets such as images,
radio-groups or checkboxes are not searched. The filter function ignores `PHANTOM`
fields, hidden fields and fields defined with the `INVISIBLE` attribute.

Only rows in memory can be searched. When using [page-mode](2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.") (`ON FILL
BUFFER`), the built-in filter is disabled. When implementing [dynamic tree views](2352-tree-views.md "Describes how to implement tree views."), the built-in filter will only search the
tree nodes available in the program array.

Row sorting features of list dialogs can be used in conjonction with the reduce filter: When a
reduce filter applies, the rows remain sorted.

If the current row is filtered out, all [rowbound actions](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.") are disabled automatically.

When adding new rows with [`ui.Dialog.appendRow()`](../15_library-reference/3185-ui-dialog-appendrow.md "Appends a new row in the specified list.") or [`ui.Dialog.insertRow()`](../15_library-reference/3205-ui-dialog-insertrow.md "Inserts a new row in the specified list."), or with the [`ON APPEND` / `ON INSERT`
triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list."), the new rows are visible, even if they do not match the reduce filter.

[Builtin search actions](2325-finding-rows-matching-a-pattern.md "List controllers implement a built-in find. This feature can be disabled if not required.") disable the reduce
filter.

The reduce filter is disabled when the dialog is an [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.").

## Controlling the reduce filter

The reduce filter can be enabled ("yes") or disabled ("no") with the
`reduceFilter` style attribute:

```
<Style name="Table">
  <StyleAttribute name="reduceFilter" value= "no" />
</Style>
```

The reduce filter is enabled by default, when rendering on a mobile device.

See [`Table.reduceFilter`](1648-table-style-attributes.md) and [`ScrollGrid.reduceFilter`](1646-scrollgrid-style-attributes.md).
