---
title: "TABLE item type"
source: "fgl-topics/c_fgl_FormSpecFiles_TABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > TABLE item type"
type: "concept"
---

# TABLE item type

> Defines a list view widget.

> **Note:**
>
> This topic covers the `TABLE` item type definition in form files. See also
> [the chapter dedicated to table view programming](2315-table-views.md "Describes how to implement table/list views.").

## TABLE item basics

A `TABLE` form item type defines a list view to show a scrolling list
of data records in a set of columns.

![TABLE rendering](../_images/FormItemType_TABLE_1.jpg)

*TABLE form item type*

A `TABLE` is controlled by a `DISPLAY ARRAY` or `INPUT ARRAY` dialog, see [List dialogs](2299-list-dialogs.md "Describes how to program dialogs controlling list containers.") for more details about this type of dialogs.

## Defining a TABLE

The `TABLE` form item defines a list view widget to show a set of data
records. It is bound to a [screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
grouping form fields which define the columns of the table.

The screen array definition must have exactly the same number of columns as the
`TABLE` form item. Use [`PHANTOM` fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field)."), if the number of record members in the program array
exceeds the number of columns to be displayed in the `TABLE` container.

Column titles can be defined with static labels in the `TABLE` layout. However,
this prevents you from localizing the column text. Consider using the [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item.") attribute in
the form field definition corresponding to the table columns.

A `TABLE` container can define [`AGGREGATE` fields](1674-aggregate-fields.md "An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.") to display
summary information for columns.

Front-ends support different presentation and behavior options, which can be controlled by
a `STYLE` attribute. For more details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [Table style attributes](1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container.").

## Where to use a TABLE

A `TABLE` form item can be defined in different ways:

1. As a [`TABLE`
   container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") in a `LAYOUT` tree.
2. As a [`<TABLE >` layout
   tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") with a [`TABLE` item
   definition](1746-table-item-definition.md "Defines attributes for a table layout tag.") in the `ATTRIBUTES` section.

## Defining the TABLE layout and size

The number of field columns composing
the `TABLE` container in the form layout defines the initial width of the table
view.

The width of a `TABLE` container is by default defined by the number of
visible columns in its layout, and the initial height is defined by the number of rows. To specify
explicitly the width and height of a table, use the [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.") and [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.") attributes.

Table columns can be moved around, their width can be adapted, they can be
hidden/shown, and can be selected to sort the record list automatically. To turn off these features,
use respectively the `UNMOVABLECOLUMNS`, `UNSIZABLECOLUMNS`, `UNHIDABLECOLUMNS` and `UNSORTABLECOLUMNS`
attributes.

`TABLE` colums can be made stretchable, by combining the [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable."), [`STRETCH={X|NONE}`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") and [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.") attributes.

A `TABLE` can use the [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") attribute, to pack the cells of rows vertically, for small
screens.

For more details, see also [Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.").

## Related links

**Related concepts**  

[SCROLLGRID item type](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.")

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
