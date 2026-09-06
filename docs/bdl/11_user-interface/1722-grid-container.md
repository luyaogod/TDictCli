---
title: "GRID container"
source: "fgl-topics/c_fgl_FormSpecFiles_GRID_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > GRID container"
type: "concept"
---

# GRID container

> Defines a layout area based on a grid of cells.

## Syntax

```
GRID [identifier] [ ( attribute [,...] ) ]
{
  { text
  | item-tag
  | hbox-tag
  | layout-tag
  | horizontal-line }
  [...]
}
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. [text](1667-static-items.md "A static item defines a simple form item as a final grid element that does not change.") is literal
   text that will appear in the form as a static label.
4. [item-tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") defines the
   position and length of a form item.
5. [hbox-tag](1680-hbox-tags.md "Hbox tags group several item tags within the same horizontal layout box, inside a grid-based container (GRID).") defines the
   position and length of several form items inside an horizontal box.
6. [layout-tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") defines
   the position and length of a layout tag.
7. horizontal-line is a set of (-) hyphen characters defining a horizontal line.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: none.

## Usage

The `GRID` container declares a formatted text block, defining the dimensions and
the positions of children form items.

> **Tip:**
>
> Avoid Tab characters (ASCII 9) inside the curly-brace delimited area. If used, Tab characters
> will be replaced by 8 blanks by fglform.

For more details about this item type, see [GRID item type](1692-grid-item-type.md "Defines a layout area based on a grid of cells.").

## Example

```
GRID
{
<GROUP g1                                   >
 Id:   [f1] Name: [f2                      ]
 Addr: [f3                                 ]
<                                           >
}
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")
