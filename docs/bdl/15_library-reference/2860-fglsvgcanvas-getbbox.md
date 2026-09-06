---
title: "fglsvgcanvas.getBBox()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_getbbox.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.getBBox()"
type: "concept"
---

# fglsvgcanvas.getBBox()

> Returns the bounding box of an SVG element.

## Syntax

```
FUNCTION getBBox( cid SMALLINT, id STRING )
     RETURNS t_svg_rect
```

1. cid is the SVG canvas id, as returned by
   `fglsvgcanvas.create()`.
2. id identifies the SVG element from which the bounding box is got.

## Usage

This function returns the bounding box of the SVG element identified by the id
parameter.

The values returned are `x`, `y`, `width` and
`height` (in the current user space).

The values returned by `getBBox()` can be held in a record variable defined with
the [`t_svg_rect`](2842-fglsvgcanvas-t-svg-rect-type.md "The t_svg_rect type defines the position and dimensions of a rectangle.") type.

If no SVG element is found, the values returned will be `NULL`.

## Example

```
DEFINE rect fglsvgcanvas.t_svg_rect
...
ON ACTION get_bbox
   CALL fglsvgcanvas.getBBox(cid, "label_23") RETURNING rect.*
   DISPLAY rect.x, rect.y, rect.width, rect.height
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
