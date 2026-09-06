---
title: "fglsvgcanvas.defs()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_defs.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.defs()"
type: "concept"
---

# fglsvgcanvas.defs()

> Produces an SVG "defs" element.

## Syntax

```
FUNCTION defs(
   id STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.

## Usage

This function creates a `"defs"` SVG DOM element from the parameters.

The `defs` element can be used to centralize SVG drawing elements such as [SVG patterns](2871-fglsvgcanvas-pattern.md "Produces an SVG \"pattern\" element.") or [CSS styles](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list."), that can be reused in
other SVG elements.

Content of a `"defs"` element can be generated with the [`styleList()`](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.")
function.

## Example

Creating a pattern:

```
DEFINE root_svg, defs, pattern, n om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL root_svg.appendChild( defs )

LET pattern = fglsvgcanvas.pattern( "pattern1", 0, 0, 20, 20, "userSpaceOnUse",
                                    NULL, "rotate(45)", NULL )
CALL pattern.appendChild( n:=fglsvgcanvas.rect(0,0,1000,10,NULL,NULL) )
CALL n.setAttribute(SVGATT_STYLE, 'stroke:none; fill:blue;' )
CALL pattern.appendChild( n:=fglsvgcanvas.rect(0,10,1000,10,NULL,NULL) )
CALL n.setAttribute(SVGATT_STYLE, 'stroke:none; fill:navy;' )
CALL defs.appendChild( pattern )
```

Creating CSS styles:

```
CONSTANT STYLE_1 = 1
CONSTANT STYLE_2 = 2
CONSTANT STYLE_3 = 3
DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes
DEFINE root_svg, defs om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL defs.appendChild( fglsvgcanvas.styleList(
                              fglsvgcanvas.styleDefinition(".style_1",attr[STYLE_1])
                           || fglsvgcanvas.styleDefinition(".style_2",attr[STYLE_2])
                           || fglsvgcanvas.styleDefinition(".style_3",attr[STYLE_3])
                           )
                         )
CALL root_svg.appendChild( defs )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
