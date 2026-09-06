---
title: "fglsvgcanvas.circle()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_circle.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.circle()"
type: "concept"
---

# fglsvgcanvas.circle()

> Produces an SVG "circle" element.

## Syntax

```
FUNCTION circle(
   cx STRING,
   cy STRING,
   r STRING )
  RETURNS om.DomNode
```

1. cx defines the X coordinate of the center point of the circle.
2. cy defines the Y coordinate of the center point of the circle.
3. r defines the radius of the circle.

## Usage

This function creates a `"circle"` SVG DOM element from the parameters.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.circle(100,100,50)
CALL n.setAttribute(SVGATT_STYLE,"stroke:#006600;")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
