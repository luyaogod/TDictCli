---
title: "GRID item type"
source: "fgl-topics/c_fgl_FormSpecFiles_GRID.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > GRID item type"
type: "concept"
---

# GRID item type

> Defines a layout area based on a grid of cells.

## GRID item basics

A `GRID` form item defines an area in the layout section to place
children form items by X,Y position in layout cells.

![GRID rendering](../_images/FormItemType_GRID_1.jpg)

*GRID form item type*

## Defining an GRID

The `GRID` container declares a formatted text block defining the
dimensions and the positions of the form items contained in the grid.

You can specify the position of labels, form fields for data entry or additional
interactive objects such as buttons.

A `GRID` container can hold static text, item tags, field tags, hbox
tags, and layout tags to define other containers such as `TABLE`,
`TREE` and `SCROLLGRID`.

A `GRID` can hold form items such as labels, fields, or buttons at a
specific position. Form items are located with item tags in the grid layout area.
You can use layout tags to place some type of containers inside a grid.

Front-ends support different presentation and behavior options, which can be
controlled by a [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.")
attribute. For more details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.").

## Where to use a GRID

A `GRID` form item must be defined as a [`GRID` container](1722-grid-container.md "Defines a layout area based on a grid of cells.") in a
`LAYOUT` tree.

See also [Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.").

## Related links

**Related concepts**  

[Layout items](1668-layout-items.md "Layout items are containers with a body that can hold other form items.")
