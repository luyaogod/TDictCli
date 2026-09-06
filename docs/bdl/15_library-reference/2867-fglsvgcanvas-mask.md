---
title: "fglsvgcanvas.mask()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_mask.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.mask()"
type: "concept"
---

# fglsvgcanvas.mask()

> Produces an SVG "mask" element.

## Syntax

```
FUNCTION mask(
   id STRING,
   x STRING,
   y STRING,
   width STRING,
   height STRING,
    node om.DomNode,
    name STRING )
```

1. id is the SVG object identifier.
2. x is the X coordinate of the mask.
3. x is the Y coordinate of the mask.
4. width is the width of the mask.
5. height is the height of the mask.
6. maskUnits defines the coordinate system for x,y,width,height.
7. maskContentUnits defines the coordinate system for the content of the
   mask.

## Usage

This function creates a `"mask"` SVG DOM element from the parameters.

To use a mask in an SVG DOM element, set the `"mask"` attribute with the string
generate by the [`url()`](2891-fglsvgcanvas-url.md "Produces a \"url(#name)\" reference for SVG elements.") function.

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
