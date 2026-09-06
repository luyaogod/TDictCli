---
title: "fglsvgcanvas.linearGradient()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_lineargradient.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.linearGradient()"
type: "concept"
---

# fglsvgcanvas.linearGradient()

> Produces an SVG "linearGradient" element.

## Syntax

```
FUNCTION linearGradient(
   id STRING,
   x1 STRING,
   y1 STRING,
   x2 STRING,
   y2 STRING,
   spreadMethod STRING,
   gradientTransform STRING,
   gradientUnits STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. x1 defines the X coordinate of the linear gradient start point.
3. y1 defines the Y coordinate of the linear gradient start point.
4. x2 defines the X coordinate of the linear gradient end point.
5. y2 defines the Y coordinate of the linear gradient end point.
6. spreadMethod defines how the gradient is spread out through the shape.
7. gradientTransform defines the gradient transformation.
8. gradientUnits defines the coordinate system to be used.

## Usage

This function creates a `"linearGradient"` SVG DOM element from the
parameters.

An SVG `"linearGradient"` element can be used to define fill colors of shapes.

The element must contain `"stop"` elements that can be created with the [`stop`](2880-fglsvgcanvas-stop.md "Produces an SVG \"stop\" element for gradients.") function.

The resulting gradient definition can be added to a `"defs"` SVG DOM element.

## Example

```
DEFINE root_svg, defs, lg om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL root_svg.appendChild( defs )

LET lg = fglsvgcanvas.linearGradient( "gradient_1",
                                      "0%", "0%", "0%", "100%",
                                      "pad", NULL, NULL )
CALL lg.appendChild( fglsvgcanvas.stop(   "0%", "gray", 0.8 ) )
CALL lg.appendChild( fglsvgcanvas.stop( "100%", "navy", 0.6 ) )
CALL defs.appendChild( lg )
...
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
