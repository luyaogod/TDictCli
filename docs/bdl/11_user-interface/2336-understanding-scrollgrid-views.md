---
title: "Understanding scrollgrid views"
source: "fgl-topics/c_fgl_ui_scrollgrid_intro.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > Understanding scrollgrid views"
type: "concept"
---

# Understanding scrollgrid views

> A scrollgrid view defines a graphical element to display a scrolling list of data records in a set of form fields positioned in a grid.

The scrollgrid is defined as a template for form fields that make up the scroll grid elements.
The front-end clones this template for each record that fits in the enclosing scrollgrid container.
The end user sees a multiple-record view like a list with a vertical scrollbar that can be
navigated to select or edit rows, depending on the dialog controlling the scrollgrid.

By default, scrollgrids display a fixed number of visible records, as defined by the number of
templates in the grid layout. The scrollgrid can be configured to be resizable in height, so that if
screen or page size grows, the scrollgrid resizes and more records are visible.

![Screenshot of form with resizable scrollgrid view](../_images/scrollgrid_list_gbc.jpg)

*Form with resizable scrollgrid view*

A usual pattern on the Web is to render information as a responsive tile list, using tiles
displayed in pages. A resizable scrollgrid can easily be configured as a responsive tile list:

![Screenshot of form with paged scrollgrid view](../_images/scrollgrid_paged_gbc.jpg)

*Form with scrollgrid as tilelist view*

## Scrollgrid view controllers

Scrollgrid views are similar to table views in that they are controlled by a
`DISPLAY ARRAY` or `INPUT ARRAY` instruction, using a form
screen-array bound to a `SCROLLGRID` container.

When controlled by a `DISPLAY ARRAY`, the scrollgrid view is by default
read-only. However, you can implement [modification triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list."), to let the end
user append, modify and delete rows.

When controlled by an `INPUT ARRAY`, the scrollgrid view allows immediate
data modification: The rows are editable.

A scrollgrid can also be used with an `INPUT` or `CONSTRUCT`
dialog: In such case, the user can input field values in the first scrollgrid row only.

You can customize the rendering and the behavior of scrollgrid views with form attributes
in the [`SCROLLGRID`
container](1723-scrollgrid-container.md "Defines a scrollable grid view widget."), and in the program using the [dialog implementation](2335-scrollgrid-views.md "Describes how to implement scrollgrid views.").

## Related links

**Related concepts**  

[List dialogs](2299-list-dialogs.md "Describes how to program dialogs controlling list containers.")
