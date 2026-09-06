---
title: "fglsvgcanvas.rect()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_rect.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.rect()"
type: "concept"
---

# fglsvgcanvas.rect()

> Produces an SVG "rect" element.

## Syntax

```
FUNCTION rect(
   x STRING,
   y STRING,
   width STRING,
   height STRING,
   rx STRING,
   ry STRING )
  RETURNS om.DomNode
```

1. x defines the X coordinate of the top/left corner.
2. y defines the Y coordinate of the top/left corner.
3. width defines the width of the rectangle.
4. height defines the height of the rectangle.
5. rx defines the X radius for rounded corners.
6. ry defines the Y radius for rounded corners.

## Usage

This function creates a `"rect"` SVG DOM element from the parameters.

Rounded borders are got by setting the rx, ry
parameters.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.rect(10,10,50,30,1,1)
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
