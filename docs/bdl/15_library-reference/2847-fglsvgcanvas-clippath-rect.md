---
title: "fglsvgcanvas.clipPath_rect()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_clippath_rect.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.clipPath_rect()"
type: "concept"
---

# fglsvgcanvas.clipPath_rect()

> Produces an SVG "clipPath" element with a "rect" element.

## Syntax

```
FUNCTION clipPath_rect(
   id STRING,
   x STRING,
   y STRING,
   width STRING,
   height STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. x is the X coordinate of the top-left corner of the clip rectangle.
3. x is the Y coordinate of the top-left corner of the clip rectangle.
4. width is the width of the clip rectangle.
5. height is the height of the clip rectangle.

## Usage

This function creates a `"clipPath"` SVG DOM element using a rectangle as clipping
shape.

A clip path can be used in a text or image element, to define the clipping borders.

The clip path element must be added to the `defs` node of the `svg`
element.

To use a clip path in an SVG DOM element, set the `"clip-path"` attribute with the
string generate by the [`url()`](2891-fglsvgcanvas-url.md "Produces a \"url(#name)\" reference for SVG elements.") function.

## Example

```
DEFINE defs om.DomNode,
       clip_id STRING
...
LET clip_id = "clip_gh_1"
CALL defs.appendChild(
             fglsvgcanvas.clipPath_rect( clip_id, tx, ty, -1, sy )
     )
...
LET n = fglsvgcanvas.text( tx + 3, (ty + (sy/2) + 1), "Some text", NULL )
CALL n.setAttribute("clip-path", fglsvgcanvas.url("clip_gh_1") )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
