---
title: "Sorting rows in a list"
source: "fgl-topics/c_fgl_prog_dialogs_list_sort.html"
breadcrumb: "User interface > User interface programming > Table views > Sorting rows in a list"
type: "concept"
---

# Sorting rows in a list

> List controllers implement a built-in sort. This feature can be disabled if not required.

When a [`DISPLAY
ARRAY`](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.") or [`INPUT
ARRAY`](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.") block is combined with a [`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") container, the
row sorting feature is implicitly available. Row sorting is supported on [`TREE`](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.") containers with
`DISPLAY ARRAY` dialogs only.

To sort rows in a list, the user must select a column header of the table, using a mouse click on
desktop. Multiple sort columns can be selected, by using alt-click.

The program dialog code can also define primary sort columns (also known as "group by" columns),
with the [`setGroupBy()`](../15_library-reference/3229-ui-dialog-setgroupby.md "Defines the grouping column of a list dialog.") and
[`addGroupBy()`](../15_library-reference/3181-ui-dialog-addgroupby.md "Appends a grouping column of a list dialog.") dialog methods.
Primary sort columns are part of the sorting columns list. Primary sort column force rows to be
always grouped by the values of the given columns. To specify the initial sort order of a primary
column, you can use the [`setGroupByDesc()`](../15_library-reference/3230-ui-dialog-setgroupbydesc.md "Defines sort order for a grouping column of a list dialog.") method.

When a given column is already selected for ascending sort, the next [alt-]click will select the
column for descending order, and when already selected for descending order, a next [alt-]click will
de-select that sort column, except if the column is defined as primary sort column by program. In
this case, the primary sort column must remain and the end user can only switch from ascending to
descending order.

Selecting a table column header triggers a GUI event, that instructs the runtime system to
reorder the rows displayed in the list container.

In fact, the rows are only sorted from a visual point of view; the data rows in the program array
(the model) are left untouched. Therefore, when sorting applies, the visual position of the current
row might be different from the current row index in the program array. To convert a visual index
to/from the program array index, use the [`visualToArrayIndex()`](../15_library-reference/3234-ui-dialog-visualtoarrayindex.md "Converts the visual index to the program array index for a given screen array.")/[`arrayToVisualIndex()`](../15_library-reference/3187-ui-dialog-arraytovisualindex.md "Converts the program array index to the visual index for a given screen array.") dialog methods.

The program array used by a `DISPLAY ARRAY` or `INPUT ARRAY` can be
modified in different ways:

- Interactively:
  - With [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")
    insert/append/delete actions, current row edition.
  - With `DISPLAY ARRAY` [modification triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.") such as `ON INSERT`.
- By program:
  - With `DIALOG` methods [`deleteRow()`](../15_library-reference/3192-ui-dialog-deleterow.md "Deletes a row from the specified list."), [`insertRow()`](../15_library-reference/3205-ui-dialog-insertrow.md "Inserts a new row in the specified list."), [`appendRow()`](../15_library-reference/3185-ui-dialog-appendrow.md "Appends a new row in the specified list."), [`deleteAllRows()`](../15_library-reference/3190-ui-dialog-deleteallrows.md "Deletes all rows from the specified list."), [`deleteNode()`](../15_library-reference/3191-ui-dialog-deletenode.md "Deletes a node from the specified tree-view."), [`insertNode()`](../15_library-reference/3204-ui-dialog-insertnode.md "Inserts a new node in the specified tree."), [`appendNode()`](../15_library-reference/3186-ui-dialog-appendnode.md "Appends a new node in the specified tree-view.").
  - With `DYNAMIC ARRAY` methods [`deleteElement()`](../15_library-reference/2948-dynamic-array-deleteelement.md "Removes an element from the array."), [`insertElement()`](../15_library-reference/2950-dynamic-array-insertelement.md "Inserts a new element at the given index."), [`appendElement()`](../15_library-reference/2945-dynamic-array-appendelement.md "Adds a new element to the end of the array."), [`clear()`](../15_library-reference/2946-dynamic-array-clear.md "Removes all elements of the array."), [`copyTo()`](../15_library-reference/2947-dynamic-array-copyto.md "Copies a complete array to the destination array passed as parameter.").

When the program array is modified interactively, or by program with `DIALOG`
methods, the runtime does not perform row sorting, to keep new created rows visible. However,
the runtime will perform row sorting, when the length of program array changes, after using
`DYNAMIC ARRAY` methods like `appendElement()`: Since the exact array
operation is not known by the runtime when manipulating directly the program array, the runtime must
re-build the entire internal list of visual indexes. Note that if you delete and then append/insert
a new row with array methods, the number of rows does not change and not automatic new sort will
apply. Consequently, it is recommended using only `DIALOG` methods or `DISPLAY
ARRAY` modification triggers, to manipulate the array rows by program during the dialog
execution.

To sort rows on a character string column, the runtime system uses the standard collation order
of the system, following the current [locale](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
settings. As a result, the rows might be ordered a bit differently than when using the database
server sort (with an `ORDER BY` clause of the `SELECT` statement),
since database servers can define their own collation sequences to sort character data.

The data used to sort rows is the raw data in the program array. When using a
`COMBOBOX` with [`ITEMS`](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.") defining key/label pairs, the runtime system uses the key values to
sort the table rows.

The built-in sort is enabled by default. To prevent sorting in `TABLE` or
`TREE` containers, define the [`UNSORTABLECOLUNMS`](1837-unsortablecolumns-attribute.md "The UNSORTABLECOLUMNS attribute indicates that the columns of the table cannot be selected by the user for sorting.") attribute at the list container level, or set the
`UNSORTABLE` attribute at the column/field level. As rows can be created and modified
during an `INPUT ARRAY` instruction, you may want to use the
`UNSORTABLECOLUMNS` attribute for tables controlled by `INPUT
ARRAY`.

To execute code after a sort was performed, use the [`ON SORT`](1989-on-sort-block.md) interaction block in the dialog, for example to display the current
row position with [`DIALOG.arrayToVisualIndex()`](../15_library-reference/3187-ui-dialog-arraytovisualindex.md "Converts the program array index to the visual index for a given screen array."). With a page-mode `DISPLAY
ARRAY` using `ON FILL BUFFER`, the build-in sort is disabled. Use the
`ON SORT` trigger to re-sort the result providing rows in `ON FILL
BUFFER`. The dialog methods [`getSortKeyAt()`](../15_library-reference/3203-ui-dialog-getsortkeyat.md "Returns the name of field used as grouping or sorting column, for a given sort column position.") and [`isSortReverseAt()`](../15_library-reference/3208-ui-dialog-issortreverseat.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending), for a given sort column position.") methods must be used to get the sort columns and sort
order. For more details, see [Paged mode with sorting feature](2309-paged-mode-of-display-array.md).

When an application window is closed, the selected sort column and order is saved. The sort will
be automatically re-applied the next time the window is created. This way, the rows will appear
sorted when the program restarts. The saved sort column and order is specific to each list
container.

In order to cleanup the sort columns (primary sort columns and user-defined sort columns), call
the [`resetSort()`](../15_library-reference/3211-ui-dialog-resetsort.md "Resets the sort columns of a list dialog.") dialog method.
This method can be used in `BEFORE DISPLAY` or `BEFORE DIALOG` blocks,
and during the dialog execution, in an `ON ACTION` block.

Custom sorting rules can be implemented with a comparison function, that must be associated to a
column field name with the [`ui.Dialog.setColumnComparisonFunction()`](../15_library-reference/3222-ui-dialog-setcolumncomparisonfunction.md "Associates a comparison function to a form field of a list dialog.") method. Do not implement complex
code in such function: It will have an impact on performances with a large set or rows, compared to
the build-in sorting algorithm.

## Related links

**Related concepts**  

[Record list (DISPLAY ARRAY)](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")

[Editable record list (INPUT ARRAY)](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")

[Paged mode of DISPLAY ARRAY](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.")

[Full list mode of DISPLAY ARRAY](2308-full-list-mode-of-display-array.md "In order to handle short/medium result sets, use the full list mode of DISPLAY ARRAY.")
