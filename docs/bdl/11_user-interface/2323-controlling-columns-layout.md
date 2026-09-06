---
title: "Controlling columns layout"
source: "fgl-topics/c_fgl_prog_dialogs_list_cols.html"
breadcrumb: "User interface > User interface programming > Table views > Controlling columns layout"
type: "concept"
---

# Controlling columns layout

> By default, a user can position, hide, show, and resize columns in TABLE and TREE containers.

> **Important:**
>
> This feature is not supported on mobile platforms.

## Resizing columns

By default, columns can be resized. On desktop front-ends, the user can drag the right edge of a
column header to increase or decrease the width of the column.

To prevent column resizing for all columns in a table, add the [`UNSIZABLECOLUMNS`](1835-unsizablecolumns-attribute.md "The UNSIZABLECOLUMNS attribute indicates that the columns of the table cannot be resized by the user.") attribute
to the `TABLE` or `TREE` container.

To prevent column resizing for an individual column, add the [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user.") attribute to the form
field definition for that column.

## Hiding/showing columns

By default, the user can control the visibility of columns. On desktop front-ends, scrolling up
shows a configuration button on the top-right of the table, that allows the show/hide columns.

To disable the column visibility option for all columns in a table, add the [`UNHIDABLECOLUMNS`](1831-unhidablecolumns-attribute.md "The UNHIDABLECOLUMNS attribute indicates that the columns of the table cannot be hidden or shown by the user with the context menu.") attribute
to the `TABLE` or `TREE` container.

To disable the column visibility option for an individual column, add the [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu.") attribute to the form
field definition for that column.

To hide a column initially but allow column visibility, set the [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.") attribute with the value
`USER` in the form field definition for that column. This hides the column by
default, and allows the user to show the column if needed.

## Changing column positions

By default, columns can be moved around. On desktop front-ends, a user can rearrange columns by
dragging the column header to a the left or to the right.

To disable this option, add the [`UNMOVABLECOLUMNS`](1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table.") attribute to the `TABLE` or
`TREE` container.

To disable this option for an individual column, add the [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table.") attribute to the form
field definition for that column.
