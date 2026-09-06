---
title: "Understanding list dialogs"
source: "fgl-topics/c_fgl_ui_list_dialogs_intro.html"
breadcrumb: "User interface > User interface programming > List dialogs > Understanding list dialogs"
type: "concept"
---

# Understanding list dialogs

> List dialogs are dialogs controlling a list of records rendered in a list container such as TABLE, TREE or SCROLLGRID.

Genero provides the `DISPLAY ARRAY` and `INPUT ARRAY` list dialogs
to control a list of records that are defined in a [program
array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.").

The `DISPLAY ARRAY` and `INPUT ARRAY` dialogs can use dynamic or
static arrays. Static arrays are supported for backward compatibility. Consider using dynamic arrays
in new implementations.

A `DISPLAY ARRAY` handles by default a read-only list. However, you can implement
[modification triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list."), to let the end user
append, modify and delete rows.

An `INPUT ARRAY` dialog allows immediate data modification: The rows are editable
by default.

The topics in this chapter are common to all list dialogs.

The form must define a list container to display the records, it can be one of the following:

- [`TABLE`](1703-table-item-type.md "Defines a list view widget.")
- [`TREE`](1706-tree-item-type.md "Defines a tree view widget.")
- [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.")

Note that the list dialog can control a plain grid with basic form fields. In this case, only
one record will be displayed at a time, but list navigation is still available.

## Related links

**Related concepts**  

[Record list (DISPLAY ARRAY)](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")

[Editable record list (INPUT ARRAY)](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")
