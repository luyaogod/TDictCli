---
title: "GROUP item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_GROUP_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > GROUP item definition"
type: "concept"
---

# GROUP item definition

> Defines attributes for a group-box layout tag.

## Syntax

```
GROUP layout-tag: [ item-name , ] [ attribute-list ] ;
```

1. layout-tag is an identifier that defines the name of the layout tag.
2. item-name identifies the form item, it is optional but recommended.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container."),
[`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`collapsible`](1639-group-style-attributes.md), [`collapserPosition`](1639-group-style-attributes.md), [`initiallyCollapsed`](1639-group-style-attributes.md).

## Usage

Define the rendering and behavior of a group [layout tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container."),
with a `GROUP` element in the `ATTRIBUTES` section.

For more details about this item type, see [GROUP item type](1693-group-item-type.md "Defines a layout area to group other layout elements together.").

## Example

```
LAYOUT
GRID
{
<GROUP g1             >
 Num: [f001         ]
 ...

}
END
END

ATTRIBUTES
GROUP g1: group1,
   TEXT="Description",
   GRIDCHILDRENINPARENT;
...
```

## Related links

**Related concepts**  

[GROUP container](1719-group-container.md "Defines a layout area to group other layout elements together.")

[Examples](1853-examples.md "Form definition (.per) examples.")
