---
title: "fglsvgcanvas.polyline()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_polyline.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.polyline()"
type: "concept"
---

# fglsvgcanvas.polyline()

> Produces an SVG "polyline" element.

## Syntax

```
FUNCTION polyline(
   points STRING )
  RETURNS om.DomNode
```

1. points is a `STRING` defining the points of the polygon.

## Usage

This function creates a `"polygon"` SVG DOM element from the parameters.

The points is a string containing a list of X,Y coordinates to draw the shape,
in the form `"x1,y1 x2,y2 ..."`.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.polyline("10,10 10,20 20,20")
CALL n.setAttribute(SVGATT_STYLE, "stroke:#660000;")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
