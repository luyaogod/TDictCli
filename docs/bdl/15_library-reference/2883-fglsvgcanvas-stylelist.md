---
title: "fglsvgcanvas.styleList()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_stylelist.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.styleList()"
type: "concept"
---

# fglsvgcanvas.styleList()

> Produces a CSS style list.

## Syntax

```
FUNCTION styleList(
    content STRING )
  RETURNS om.DomNode
```

1. content is the string that is a concatenation of style definitions produced
   by `styleDefinition()` functions.

## Usage

This function creates a `"style"` SVG DOM element with a text sub-node containing
a list of style definitions, to create a CSS style. The resulting DOM element can then be used in a
`"defs"` element.

A single CSS style definition can be produced with the [`styleDefinition()`](2882-fglsvgcanvas-styledefinition.md "Produces a CSS style definition with a selection and list of attributes.") function.

## Example

```
CONSTANT STYLE_1 = 1
CONSTANT STYLE_2 = 2
CONSTANT STYLE_3 = 3
DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes,
       root_svg, defs om.DomNode
... 
LET attr[STYLE_1] = om.SaxAttributes.create()
CALL attr[STYLE_1].setAttribute(SVGATT_STROKE,"red")
...
LET attr[STYLE_2] = om.SaxAttributes.create()
...
LET attr[STYLE_3] = om.SaxAttributes.create()
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
