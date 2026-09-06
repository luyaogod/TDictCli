---
title: "Understanding tables views"
source: "fgl-topics/c_fgl_ui_tables_intro.html"
breadcrumb: "User interface > User interface programming > Table views > Understanding tables views"
type: "concept"
---

# Understanding tables views

> Table views define the graphical element to display a list of records.

The end user can navigate in the list to select a row or edit rows, depending on the
dialog controlling the table.

If the front-end platform standards allow it, the user can resize the table, sort rows,
move/resize/hide columns, make multiple-row selections, search rows by criterion, and more.

![Table view screenshot](../_images/SimpleList2_gbc.jpg)

*Table View*

Tables views are controlled by a `DISPLAY ARRAY` or `INPUT ARRAY`
instruction, using a form screen-array bound to a `TABLE` container.

When controlled by a `DISPLAY ARRAY`, the table view is by default read-only.
However, you can implement [modification
triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list."), to let the end user append, modify and delete rows.

When controlled by an `INPUT ARRAY`, the table view allows immediate data
modification: The rows are editable.

Note that a table view can also be used with an `INPUT` or `CONSTRUCT`
dialog.

You can customize the rendering and the behavior of table views with form attributes in
the [`TABLE` container](1724-table-container.md "Defines a re-sizable table designed to display a list of records."),
and in the program using the [dialog implementation](2315-table-views.md "Describes how to implement table/list views.").

## Related links

**Related concepts**  

[List dialogs](2299-list-dialogs.md "Describes how to program dialogs controlling list containers.")
