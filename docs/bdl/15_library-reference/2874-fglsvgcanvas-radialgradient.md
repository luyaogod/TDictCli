---
title: "fglsvgcanvas.radialGradient()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_radialgradient.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.radialGradient()"
type: "concept"
---

# fglsvgcanvas.radialGradient()

> Produces an SVG "radialGradient" element.

## Syntax

```
FUNCTION radialgradient(
   id STRING,
   cx STRING,
   cy STRING,
   fx STRING,
   fy STRING,
   r STRING,
   spreadMethod STRING,
   gradientTransform STRING,
   gradientUnits STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. cx defines the X coordinate of the radial gradient center.
3. cy defines the Y coordinate of the radial gradient center.
4. fx defines the X coordinate of the radial gradient focal point.
5. fy defines the Y coordinate of the radial gradient focal point.
6. spreadMethod defines how the gradient is spread out through the shape.
7. gradientTransform defines the gradient transformation.
8. gradientUnits defines the coordinate system to be used.

## Usage

This function creates a `"radialGradient"` SVG DOM element from the
parameters.

An SVG `"radialGradient"` element can be used to define fill colors of shapes.

The element must contain `"stop"` elements that can be created with the [`stop`](2880-fglsvgcanvas-stop.md "Produces an SVG \"stop\" element for gradients.") function.

The resulting gradient definition can be added to a `"defs"` SVG element.

## Example

```
DEFINE root_svg, defs, rg om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL root_svg.appendChild( defs )

LET rg = fglsvgcanvas.radialGradient( "gradient_1",
                                      NULL, NULL, "5%", "5%", "65%",
                                      "pad", NULL, NULL )
CALL rg.appendChild( fglsvgcanvas.stop(   "0%", "gray", 0.4 ) )
CALL rg.appendChild( fglsvgcanvas.stop( "100%", "navy", 0.7 ) )
CALL defs.appendChild( rg )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
