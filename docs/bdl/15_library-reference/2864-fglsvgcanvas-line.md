---
title: "fglsvgcanvas.line()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_line.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.line()"
type: "concept"
---

# fglsvgcanvas.line()

> Produces an SVG "line" element.

## Syntax

```
FUNCTION line(
   x1 STRING,
   y1 STRING,
   x2 STRING,
   y2 STRING )
  RETURNS om.DomNode
```

1. x1 defines the X coordinate of the start point.
2. y1 defines the Y coordinate of the start point.
3. x2 defines the X coordinate of the end point.
4. y2 defines the Y coordinate of the end point.

## Usage

This function creates a `"line"` SVG DOM element from the parameters.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.line(10,10,50,30)
CALL n.setAttribute(SVGATT_STYLE,"stroke:#006600;")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
