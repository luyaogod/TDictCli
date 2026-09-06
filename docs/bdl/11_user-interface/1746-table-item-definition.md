---
title: "TABLE item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_TABLE_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > TABLE item definition"
type: "concept"
---

# TABLE item definition

> Defines attributes for a table layout tag.

## Syntax

```
TABLE layout-tag: [ item-name , ] [ attribute-list ] ;
```

1. layout-tag is an identifier that defines the name of the layout tag.
2. item-name identifies the form item, it is optional but recommended.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`AGGREGATETEXT`](1759-aggregatetext-attribute.md "The AGGREGATETEXT attribute defines a label to be displayed for aggregate fields."),
`DOUBLECLICK`, [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`UNHIDABLECOLUMNS`](1831-unhidablecolumns-attribute.md "The UNHIDABLECOLUMNS attribute indicates that the columns of the table cannot be hidden or shown by the user with the context menu."), [`UNMOVABLECOLUMNS`](1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table."), [`UNSIZABLECOLUMNS`](1835-unsizablecolumns-attribute.md "The UNSIZABLECOLUMNS attribute indicates that the columns of the table cannot be resized by the user."), [`UNSORTABLECOLUMNS`](1837-unsortablecolumns-attribute.md "The UNSORTABLECOLUMNS attribute indicates that the columns of the table cannot be selected by the user for sorting."), [`WANTFIXEDPAGESIZE`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element."), [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`allowWebSelection`](1648-table-style-attributes.md), [`alternateRows`](1648-table-style-attributes.md), [`forceDefaultSettings`](1648-table-style-attributes.md), [`headerAlignment`](1648-table-style-attributes.md), [`headerHidden`](1648-table-style-attributes.md), [`headerPosition`](1648-table-style-attributes.md), [`highlightColor`](1648-table-style-attributes.md), [`highlightCurrentCell`](1648-table-style-attributes.md), [`highlightCurrentRow`](1648-table-style-attributes.md), [`highlightTextColor`](1648-table-style-attributes.md), [`leftFrozenColumns`](1648-table-style-attributes.md), [`reduceFilter`](1648-table-style-attributes.md), [`resizeFillsEmptySpace`](1648-table-style-attributes.md), [`rightFrozenColumns`](1648-table-style-attributes.md), [`rowActionTrigger`](1648-table-style-attributes.md), [`rowAspect`](1648-table-style-attributes.md), [`rowHover`](1648-table-style-attributes.md), [`showGrid`](1648-table-style-attributes.md), [`tableType`](1648-table-style-attributes.md).

## Usage

Define a `TABLE` element in the `ATTRIBUTES` section, to
configure a table layouted with a `<TABLE >`
[layout tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.").

For more details about this item type, see [TABLE item type](1703-table-item-type.md "Defines a list view widget.").

## Example

```
LAYOUT
GRID
{
<TABLE t1                    >
[c1  |c2        |c3          ]
[c1  |c2        |c3          ]
 ...

}
END
END

ATTRIBUTES
TABLE t1: table1, UNSORTABLECOLUMNS;
...
```
