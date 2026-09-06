---
title: "fglsvgcanvas.text_path()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_text_path.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.text_path()"
type: "concept"
---

# fglsvgcanvas.text_path()

> Produces the SVG "text" element with a "textPath" sub-element.

## Syntax

```
FUNCTION text_path(
   x STRING,
   y STRING,
   content STRING,
   path STRING
   class STRING )
  RETURNS om.DomNode
```

1. x and y define the position of the text.
2. content is the actual text.
3. path is the xlink:href reference (without #).
4. class defines a reference to a CSS style.

## Usage

This function creates a `"text"` SVG DOM element from the parameters, including a
`"textPath"` sub-element that references a `"path"` element defined in
a `"defs"` element.

The path parameter is used to build an
`"xlink:href=#path"` reference.

The actual path can be created with the [`path()`](2870-fglsvgcanvas-path.md "Produces an SVG \"path\" element.") function, and
included in a `"defs"` element created with the [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") function.

To specify the text font attributes, define a CSS style in a [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") element with
the [styleList()](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.") function,
and reference the text style in the class parameter of this function.

## Example

```
DEFINE root_svg, defs, p, n om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL root_svg.appendChild( defs )
...
CALL defs.appendChild( p:=fglsvgcanvas.path("M150,400 C175,380 225,320 450,450") )
CALL p.setAttribute("id", "path_1")
LET n = fglsvgcanvas.text_path( NULL,NULL,
           "This text follows a path...","path_1","style_4")
...
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
