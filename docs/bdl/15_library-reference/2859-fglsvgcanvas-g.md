---
title: "fglsvgcanvas.g()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_g.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.g()"
type: "concept"
---

# fglsvgcanvas.g()

> Produces an SVG "g" element.

## Syntax

```
FUNCTION g(
   id STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.

## Usage

This function creates a `"g"` SVG DOM element from the parameters.

The `g` element is used to group SVG elements together, apply a transformation,
and use common attributes for the group.

For example, you can draw complex shapes inside the "g" element, and rotate the whole object by
180 degrees around point (50,50), by using `"rotate(180 50 50)"` for the
`"transform"` parameter. For more details about the `"transform"` SVG
attribute, see SVG specification

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.g( "group1" )
CALL n.setAttribute(SVGATT_TRANSFORM, "rotate(45 50 50)")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
