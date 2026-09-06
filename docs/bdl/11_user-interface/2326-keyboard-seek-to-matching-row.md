---
title: "Keyboard seek to matching row"
source: "fgl-topics/c_fgl_prog_dialogs_list_seek.html"
breadcrumb: "User interface > User interface programming > Table views > Keyboard seek to matching row"
type: "concept"
---

# Keyboard seek to matching row

> The keyboard seek feature allows a user to find a row in a read-only list, by typing characters.

When a [`DISPLAY ARRAY`](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.")
is used with a list, the keyboard seek feature is automatically implemented. A user may type alphabetic
characters on the keyboard to have the runtime system automatically seek the next row having a character
field that contains a value starting with the typed characters. The seek search restarts from the
current row when the user types new characters on the keyboard.

> **Important:**
>
> This feature is not supported on mobile platforms.

This feature works with any list container ([`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records."), [`TREE`](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget."), [`SCROLLGRID`](1723-scrollgrid-container.md "Defines a scrollable grid view widget.")
).

Numeric, date/time and large data (`TEXT`/`BYTE`) columns are
ignored. Only character columns are searched, fields using widgets like image, radio-group or
checkbox are ignored. Furthermore, the seek function ignores `PHANTOM` fields,
hidden fields and fields defined with the `INVISIBLE` attribute.

The user can rapidly type several characters on the keyboard, to search for a value that starts
with the typed characters. After a given timeout (less than a second), the seek buffer is cleared
and a new search filter can be applied.

The seek search is case-insensitive.

If no row is found from the typed characters, the [Not found] error [-8105](../15_library-reference/4483-genero-bdl-errors.md) is displayed
automatically.

If an alphabetic character is used as action accelerator, the built-in seek feature is disabled,
because the accelerator must fire the corresponding action.

Only rows in memory can be searched. When using [page-mode](2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.") (`ON FILL BUFFER`),
the built-in seek is disabled. When implementing [dynamic tree
views](2352-tree-views.md "Describes how to implement tree views."), the built-in seek will only search the tree nodes available in the program array.

By default, any character column of the list is scanned. But if the list gets [sorted](2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required."), the runtime system considers that the sort
column is the most important and searches only in that column.

## Related links

**Related concepts**  

[Finding rows matching a pattern](2325-finding-rows-matching-a-pattern.md "List controllers implement a built-in find. This feature can be disabled if not required.")
