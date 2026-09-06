---
title: "TREE item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_TREE_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > TREE item definition"
type: "concept"
---

# TREE item definition

> Defines attributes for a tree layout tag.

## Syntax

```
TREE layout-tag: [ item-name , ] [ attribute-list ] ;
```

1. layout-tag is an identifier that defines the name of the layout tag.
2. item-name identifies the form item, it is optional but recommended.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), `DOUBLECLICK`, [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`PARENTIDCOLUMN`](1806-parentidcolumn-attribute.md "The PARENTIDCOLUMN attribute specifies the form field that contains the identifier of the parent node of a tree node."), [`IDCOLUMN`](1784-idcolumn-attribute.md "The IDCOLUMN attribute specifies the form field that contains the identifier of a tree node."), [`EXPANDEDCOLUMN`](1777-expandedcolumn-attribute.md "The EXPANDEDCOLUMN attribute specifies the form field that indicates whether a tree node is expanded."), [`ISNODECOLUMN`](1794-isnodecolumn-attribute.md "The ISNODECOLUMN attribute specifies the form field that indicates whether a tree node has children."), [`IMAGEEXPANDED`](1788-imageexpanded-attribute.md "The IMAGEEXPANDED attribute sets the global icon to be used when a tree node is expanded."), [`IMAGECOLLAPSED`](1787-imagecollapsed-attribute.md "The IMAGECOLLAPSED attribute sets the global icon to be used when a tree node is collapsed."), [`IMAGELEAF`](1789-imageleaf-attribute.md "The IMAGELEAF attribute defines the global icon for leaf nodes of a TREE container."), [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`UNHIDABLECOLUMNS`](1831-unhidablecolumns-attribute.md "The UNHIDABLECOLUMNS attribute indicates that the columns of the table cannot be hidden or shown by the user with the context menu."), [`UNMOVABLECOLUMNS`](1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table."), [`UNSIZABLECOLUMNS`](1835-unsizablecolumns-attribute.md "The UNSIZABLECOLUMNS attribute indicates that the columns of the table cannot be resized by the user."), [`UNSORTABLECOLUMNS`](1837-unsortablecolumns-attribute.md "The UNSORTABLECOLUMNS attribute indicates that the columns of the table cannot be selected by the user for sorting."), [`WANTFIXEDPAGESIZE`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element."), [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`allowWebSelection`](1649-tree-style-attributes.md), [`alternateRows`](1649-tree-style-attributes.md), [`forceDefaultSettings`](1649-tree-style-attributes.md), [`headerAlignment`](1649-tree-style-attributes.md), [`headerHidden`](1649-tree-style-attributes.md), [`highlightColor`](1649-tree-style-attributes.md), [`highlightCurrentRow`](1649-tree-style-attributes.md), [`highlightTextColor`](1649-tree-style-attributes.md), [`leftFrozenColumns`](1649-tree-style-attributes.md), [`resizeFillsEmptySpace`](1649-tree-style-attributes.md), [`rightFrozenColumns`](1649-tree-style-attributes.md), [`rowActionTrigger`](1649-tree-style-attributes.md), [`rowHover`](1649-tree-style-attributes.md), [`showGrid`](1649-tree-style-attributes.md), [`tableType`](1649-tree-style-attributes.md).

## Usage

The `TREE` form item type can be used to specify the attributes of a tree
container defined with a layout tag.

For more details about this item type, see [TREE item type](1706-tree-item-type.md "Defines a tree view widget.").
