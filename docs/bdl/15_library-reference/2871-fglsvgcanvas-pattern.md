---
title: "fglsvgcanvas.pattern()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_pattern.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.pattern()"
type: "concept"
---

# fglsvgcanvas.pattern()

> Produces an SVG "pattern" element.

## Syntax

```
FUNCTION pattern(
   id STRING,
   x STRING,
   y STRING,
   width STRING,
   height STRING,
   patternUnits STRING,
   patternContentUnits STRING,
   patternTransform STRING,
   preserveAspectRatio STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. x defines the X coordinate where the patterm starts.
3. y defines the Y coordinate where the patterm starts.
4. width defines the width of the pattern.
5. height defines the height of the pattern.
6. patternUnits defines the patternUnits attribute.
7. patternContentUnits defines the patternContentUnits attribute.
8. patternTransform defines the patternTransform attribute.
9. preserveAspectRatio is the aspect ratio to preserve.

## Usage

This function creates a `"pattern"` SVG DOM element from the parameters.

The resulting DOM can be used in a `"defs"` element created with the [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") function, to
produce pattern definitions.

## Example

```
DEFINE pattern, n om.DomNode
...
LET pattern = fglsvgcanvas.pattern( "pattern1", 0,0,20,20, "userSpaceOnUse",
                                    NULL, "rotate(45)", NULL )
CALL pattern.appendChild( n:=fglsvgcanvas.rect(0,0,1000,10,NULL,NULL) )
CALL n.setAttribute(SVGATT_STYLE, 'stroke:none; fill:blue;' )
CALL pattern.appendChild( n:=fglsvgcanvas.rect(0,10,1000,10,NULL,NULL) )
CALL n.setAttribute(SVGATT_STYLE, 'stroke:none; fill:navy;' )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
