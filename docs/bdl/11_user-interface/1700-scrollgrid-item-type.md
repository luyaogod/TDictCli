---
title: "SCROLLGRID item type"
source: "fgl-topics/c_fgl_FormSpecFiles_SCROLLGRID.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > SCROLLGRID item type"
type: "concept"
---

# SCROLLGRID item type

> Defines a scrollable grid view widget.

> **Note:**
>
> This topic covers the `SCROLLGRID` item type definition in form files. See also
> [the chapter dedicated to scrollgrid view
> programming](2335-scrollgrid-views.md "Describes how to implement scrollgrid views.").

## SCROLLGRID item basics

A `SCROLLGRID` form item type defines a grid to show a scrolling list of data
records in a set of positioned form fields.

![SCROLLGRID rendering](../_images/FormItemType_SCROLLGRID_1.jpg)

*SCROLLGRID form item type*

The `SCROLLGRID` basically acts like a `TABLE` container, in the
sense of its function as a record list container.

By default, scrollgrids display a fixed number of rows, and can be configured to be resizable in
height.

A `SCROLLGRID` is controlled by a `DISPLAY ARRAY` or `INPUT
ARRAY` dialog, see [List dialogs](2299-list-dialogs.md "Describes how to program dialogs controlling list containers.") for more details about
this type of dialogs.

## Defining a SCROLLGRID

The `SCROLLGRID` form item defines a formatted list view to show a structured set
of data records. It is bound to a [screen
array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") grouping form fields which define the list fields of the scrollgrid.

The screen array definition must have exactly the same number of fields as the
`SCROLLGRID` form item. Use [`PHANTOM` fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field)."), if the number of record members in the program array
exceeds the number of fields to be displayed in the `SCROLLGRID` container.

The `SCROLLGRID` form item declares a formatted text block, defining the
dimensions and the position of the logical elements of a screen for a multi-record presentation.

A `SCROLLGRID` is similar to the `GRID`, that repeats on several
"row-templates", in order to design a view for multiple-records that display with a vertical
scrollbar.

Inside a `SCROLLGRID`, the same layout rules apply as in a `GRID`
container.

Static labels can be used as in a regular `GRID` container. However, this prevents
you from localizing the label text. Consider using `LABEL` fields with a [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") attribute, instead of static text
labels.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [ScrollGrid style attributes](1646-scrollgrid-style-attributes.md "ScrollGrid presentation style attributes apply to SCROLLGRID container.").

## Where to use a SCROLLGRID?

A `SCROLLGRID` form item must be defined in different ways:

1. As a [`SCROLLGRID`
   container](1723-scrollgrid-container.md "Defines a scrollable grid view widget.") in a `LAYOUT` tree.
2. As a [`<SCROLLGRID >` layout
   tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") with a [`SCROLLGRID` item definition](1743-scrollgrid-item-definition.md "Defines attributes for a scrollgrid layout tag.") in the `ATTRIBUTES` section.

## Defining the scrollgrid size and layout

The global size of a `SCROLLGRID` element is defined by its content, and the
number of record rows displayed.

When using a `SCROLLGRID` container, you cannot set the [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container.")
attribute. This attribute makes sense only for a scrollgrid defined with a layout tag, within a
`GRID` area.

By default, a scrollgrid is not resizable in height: The number of visible rows is defined by the
number of row templates in the form layout. Use the [`WANTFIXEDPAGESIZE=NO`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element.")
attribute, to allow the scrollgrid to resize.

When using `WANTFIXEDPAGESIZE=NO`, a resizable scrollgrid is displayed with one
single row, if no other element in the form layout makes the scrollgrid stretch implicitly. To force
a default number of rows in stretchable scrollgrids, use the [`INITIALIPAGESIZE`](1792-initialpagesize-attribute.md "The INITIALPAGESIZE attribute defines the initial page size of a list element.") attribute.

A usual pattern on the web is to render information as a responsive tile list, using tiles
displayed in pages. Each tile will flow, depending on the container size.

To render stretchable scrollgrids as a paged responsive tile list, define the [`customWidget`](1646-scrollgrid-style-attributes.md "ScrollGrid presentation style attributes apply to SCROLLGRID container.")
presentation style attribute to `pagedScrollGrid`. With this attribute, each scrollgrid
row will be rendered as a tile (the page size of the scrollgrid defines the number of tiles in each
page):

```
<Style name="ScrollGrid.paged" >
    <StyleAttribute name= "customWidget" value= "pagedScrollGrid" />
</Style>
```

For more details, see also [Controlling scrollgrid rendering](2339-controlling-scrollgrid-rendering.md "Scrollgrid rendering can be controlled by the use of presentation styles and scrollgrid attributes.").

## Related links

**Related concepts**  

[TABLE item type](1703-table-item-type.md "Defines a list view widget.")
