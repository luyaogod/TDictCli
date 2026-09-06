---
title: "CANVAS item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_CANVAS_Item_Type_2.html"
breadcrumb: "User interface > User interface programming > Canvases > CANVAS item definition"
type: "concept"
---

# CANVAS item definition

> Defines attributes for a CANVAS drawing area.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
CANVAS item-tag: item-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item
   tag in the layout section.
2. item-name identifies the form item, it is mandatory to identify the canvas in
   programs.
3. attribute-list defines the aspect and behavior of the form
   item.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: none.

## Usage

Define the rendering and behavior of a canvas drawing area [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."), with a `CANVAS` element in the `ATTRIBUTES` section.

> **Important:**
>
> The `CANVAS` feature is deprecated, consider using the [built-in fglsvgcanvas
> `WEBCOMPONENT`](2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.").

## Example

```
LAYOUT
GRID
{
[cvs1             ]
[                 ]
[                 ]
 ...

}
END
END

ATTRIBUTES
CANVAS cvs1: canvas1;
...
```

## Related links

**Related concepts**  

[Canvases](2438-canvases.md "Canvases are form drawing areas.")
