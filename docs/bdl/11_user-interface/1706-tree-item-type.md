---
title: "TREE item type"
source: "fgl-topics/c_fgl_FormSpecFiles_TREE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > TREE item type"
type: "concept"
---

# TREE item type

> Defines a tree view widget.

> **Note:**
>
> This topic covers the `TREE` item type definition in form files. See also [the chapter dedicated to tree view programming](2352-tree-views.md "Describes how to implement tree views.").

## TREE item basics

A `TREE` form item type defines a tree view to show a structured tree of data records
with an optional set of columns.

![TREE rendering](../_images/FormItemType_TREE_1.jpg)

*TREE form item type*

The `TREE` basically acts like a `TABLE` container, in the sense of
its function as a record list container. A `TREE` container renders as a tree view
widget, with regular table columns on the right of the tree view.

A `TREE` is controlled by a `DISPLAY ARRAY` dialog, see [List dialogs](2299-list-dialogs.md "Describes how to program dialogs controlling list containers.") for more
details about this type of dialog.

## Defining a TREE

The `TREE` form item defines a tree view widget to show a structured
set of data records. It is bound to a [screen
array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") grouping form fields which define the columns of the tree view.

The screen array definition must have exactly the same number of columns as the
`TREE` form item. Use [`PHANTOM` fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field)."), if the number of record members in the program array
exceeds the number of columns to be displayed in the `TREE` container.

`TREE` container definitions are very similar to regular
`TABLE` containers; before reading further about tree views, you may need to familiarize
yourself with [`TABLE`](1703-table-item-type.md "Defines a list view widget.") containers.

The first column in the `TREE` must be the field defining the text of
the tree view nodes.

Column titles can be defined with static labels in the `TREE` layout. However,
this prevents you from localizing the column text. Consider using the [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item.") attribute in the form field
definition corresponding to the table columns.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [Tree style attributes](1649-tree-style-attributes.md "Tree presentation style attributes apply to the TREE container.").

## Where to use a TREE

A `TREE` form item can be defined in different ways:

1. As a [`TREE` container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")
   in a `LAYOUT` tree.
2. As a [`<TREE >` layout
   tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") with a [`TREE` item
   definition](1749-tree-item-definition.md "Defines attributes for a tree layout tag.") in the `ATTRIBUTES` section.

## Defining the TREE layout and size

The number of field columns composing the `TREE` container in the form layout
defines the initial width of the tree view.

The width of a `TREE` container is by default defined by the number of visible
columns in its layout, and the initial height is defined by the number of rows. To specify
explicitly the width and height of a tree view, use the [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.") and [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.") attributes.

Tree view columns can be moved around, their width can be adapted, they can be hidden/shown, and
can be selected to sort the record list automatically. To turn off these features, use respectively
the `UNMOVABLECOLUMNS`,
`UNSIZABLECOLUMNS`,
`UNHIDABLECOLUMNS` and
`UNSORTABLECOLUMNS`
attributes.

`TREE` colums can be made stretchable, by combining the [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable."), [`STRETCH={X|NONE}`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") and [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.") attributes.

For more details, see also [Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.").
