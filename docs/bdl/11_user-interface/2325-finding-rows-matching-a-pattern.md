---
title: "Finding rows matching a pattern"
source: "fgl-topics/c_fgl_prog_dialogs_list_find.html"
breadcrumb: "User interface > User interface programming > Table views > Finding rows matching a pattern"
type: "concept"
---

# Finding rows matching a pattern

> List controllers implement a built-in find. This feature can be disabled if not required.

The [`DISPLAY ARRAY`](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.")
and [`INPUT ARRAY`](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.") blocks
support the built-in find feature by default.

This feature works with any list container ([`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records."), [`TREE`](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget."), [`SCROLLGRID`](1723-scrollgrid-container.md "Defines a scrollable grid view widget.") ).

The built-in find creates automatically the "find" and "findnext" actions. These actions can be
decorated, enabled and disabled as regular actions.

If the dialog defines an explicit `ON ACTION find` or `ON ACTION
findnext`, the default built-in find is disabled.

When the user triggers the "find" action (default accelerator is Ctrl-F), the dialog opens a
pop-up window to enter a search value and search options:

![Table displaying the dialog to find rows.](../_images/table_find_dialog_1.jpg)

*Table showing the find dialog*

On validation with the OK button, the dialog starts to search a row where a field value matches
the value entered in the find dialog. The "find" action starts the search from the current row (and
in the field after the current field, if the dialog is an `INPUT ARRAY`).

On mobile devices (without a physical keyboard), since the Ctrl-F accelerator cannot be
triggered, there is by default no way to trigger the "find" action. If needed, add a dedicated
action view with the name "find" to the toolbar, topmenu or set the `DEFAULTVIEW=YES`
action default attribute for this action.

After a "find" action, the user can trigger the "findnext" action (default accelerator is
Ctrl-G), in order to continue the search, without opening the find dialog again (the current search
value will be reused).

By default, any table column is scanned, but the user can select a specific column in the find
dialog box, as long as a column title is available. Case-sensitive or insensitive search as well as
wraparound options are also available.

Only rows in memory can be searched. When using the [paged-mode](2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.") (`ON FILL
BUFFER`), the built-in search is disabled. When implementing [dynamic tree views](2352-tree-views.md "Describes how to implement tree views."), the built-in find will only search the
tree nodes available in the program array.

The value entered in the find dialog is compared to column fields defined with basic primitive
types. Columns defined with type `TEXT` or `BYTE` are excluded.

Comparisons are based on the formatted display value. For example, to match a value in a
`MONEY` column, the user must type the exact formatted string—including the currency
symbol, thousands and decimal separator. The same applies to `DECIMAL` values. For
example, if the column displays "$1,234", searching for "1234" will not yield a match.

Only text widgets displaying values are searched. Columns using widgets such as images,
radio-groups, checkboxes are not searched. Furthermore, the find function ignores [`PHANTOM` fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field)."), hidden fields
and fields defined with the `INVISIBLE` attribute.

With [`COMBOBOX` fields](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values."), the
find searches in the visible values of combobox items.

When the dialog is an `INPUT ARRAY` and no specific search column is selected in
the find dialog, the search scans each cell. The search starts in the current row, after the current
field. If no cell value matches in the current row, the search continues on the first field of the
next row.

The built-in find feature is also available in [text
mode](1517-text-mode-rendering-tui-mode.md). In graphical mode, the default keyboard accelerator is Ctrl-F. When using text mode,
the accelerator is the `/` slash key.
