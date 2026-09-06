---
title: "GROUP container"
source: "fgl-topics/c_fgl_FormSpecFiles_GROUP_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > GROUP container"
type: "concept"
---

# GROUP container

> Defines a layout area to group other layout elements together.

## Syntax

```
GROUP [identifier] [ ( attribute [,...] ) ]
  layout-container
  [...] 
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. layout-container is another child container.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`collapsible`](1639-group-style-attributes.md), [`collapserPosition`](1639-group-style-attributes.md), [`initiallyCollapsed`](1639-group-style-attributes.md).

## Can hold

`VBOX`, `HBOX`, `GROUP`, `FOLDER`, `GRID`, `SCROLLGRID`, `TABLE`, `TREE`.

## Usage

In a `LAYOUT` tree definition, use a `GROUP` container
to hold other containers such as a `VBOX` with children, or a
`GRID` container.

> **Note:**
>
> When defining a `GROUP` container,
> you cannot set the [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container.")
> attribute. This attribute makes sense only for a group item defined with a layout tag contained in a
> `GRID` area.

For more details about this item type, see [GROUP item type](1693-group-item-type.md "Defines a layout area to group other layout elements together.").

## Example

```
GROUP ( TEXT = "Customer" )
  VBOX
    GRID
    {
     ...
    }
    END
    TABLE
    {
     ...
    }
    END
  END
END
```

## Related links

**Related concepts**  

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")
