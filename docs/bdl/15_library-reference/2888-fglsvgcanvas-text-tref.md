---
title: "fglsvgcanvas.text_tref()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_text_tref.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.text_tref()"
type: "concept"
---

# fglsvgcanvas.text_tref()

> Produces the SVG "text" element with a "tref" sub-element.

## Syntax

```
FUNCTION text_tref(
   x STRING,
   y STRING,
   tref STRING,
   class STRING )
  RETURNS om.DomNode
```

1. x and y define the position of the text.
2. tref is the xlink reference to the tref.
3. class defines a reference to a CSS style.

## Usage

This function creates a `"text"` SVG DOM element from the parameters, including a
`"tref"` sub-element that references a `"text"` element defined in a
`"defs"` element.

The tref parameter is used to build an
`"xlink:href=#path"` reference.

The `"text"` element referenced by the "tref" attribute must be defined in a
`"defs"` element, created with the [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") function.

To specify the text font attributes, define a CSS style in a [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") element with
the [styleList()](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.") function,
and reference the text style in the class parameter of this function.

## Example

```
DEFINE root_svg, defs, t, n om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL root_svg.appendChild( defs )
...
CALL defs.appendChild( t:=fglsvgcanvas.text( NULL, NULL,
                       "The is the referenced text..." ) )
CALL t.setAttribute("id","text_1")
LET n = fglsvgcanvas.text_tref(10,20,"text_1","style_4")
...
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
