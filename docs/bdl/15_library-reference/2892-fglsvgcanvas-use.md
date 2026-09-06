---
title: "fglsvgcanvas.use()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_use.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.use()"
type: "concept"
---

# fglsvgcanvas.use()

> Produces an SVG "use" element.

## Syntax

```
FUNCTION use(
   name STRING,
   x STRING,
   y STRING )
  RETURNS om.DomNode
```

1. name is the xlink:href reference (without #).
2. x defines the X coordinate to place the shape.
3. y defines the Y coordinate to place the shape.

## Usage

This function creates a `"use"` SVG DOM element from the parameters.

The `"use"` element can reuse an SVG shape from elsewhere in the SVG document, for
example to make several copies of shapes defined in the [`"defs"`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") element.

## Example

```
DEFINE n1, n2 om.DomNode
LET n1 = fglsvgcanvas.use( "myshape1", 100, 100 )
LET n2 = fglsvgcanvas.use( "myshape1", 110, 110 )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")

[fglsvgcanvas.symbol()](2885-fglsvgcanvas-symbol.md "Produces an SVG \"symbol\" element.")
