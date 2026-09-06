---
title: "SCROLLGRID item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_SCROLLGRID_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > SCROLLGRID item definition"
type: "concept"
---

# SCROLLGRID item definition

> Defines attributes for a scrollgrid layout tag.

## Syntax

```
SCROLLGRID layout-tag: [ item-name , ] [ attribute-list ] ;
```

1. layout-tag is an identifier that defines the name of the layout tag.
2. item-name identifies the form item, it is optional but recommended.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), `DOUBLECLICK`, [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container."),
[`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`INITIALPAGESIZE`](1792-initialpagesize-attribute.md "The INITIALPAGESIZE attribute defines the initial page size of a list element."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`customWidget`](1646-scrollgrid-style-attributes.md), [`highlightColor`](1646-scrollgrid-style-attributes.md), [`highlightCurrentCell`](1646-scrollgrid-style-attributes.md), [`highlightCurrentRow`](1646-scrollgrid-style-attributes.md), [`highlightTextColor`](1646-scrollgrid-style-attributes.md), [`rowActionTrigger`](1646-scrollgrid-style-attributes.md), [`itemsAlignment`](1646-scrollgrid-style-attributes.md).

## Usage

The `SCROLLGRID` form item type to specify the attributes of a scrollgrid container
defined with a layout tag.

For more details about this item type, see [SCROLLGRID item type](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.").

## Example

```
LAYOUT
GRID
{
<SCROLLGRID sg1             >
 [f001         ]
 ...

}
END
END

ATTRIBUTES
SCROLLGRID sg1: scrollgrid1,
  GRIDCHILDRENINPARENT;
```
