---
title: "fglsvgcanvas.text()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_text.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.text()"
type: "concept"
---

# fglsvgcanvas.text()

> Produces an SVG "text" element.

## Syntax

```
FUNCTION text(
   x STRING,
   y STRING,
   content STRING,
   class STRING )
  RETURNS om.DomNode
```

1. x and y define the position of the text.
2. content is the actual text.
3. class defines a reference to a CSS style.

## Usage

This function creates a `"text"` SVG DOM element from the parameters.

To specify the text font attributes, define a CSS style in a [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") element with
the [styleList()](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.") function,
and reference the text style in the class parameter of this function.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.text(10,10,"Hello!","mystyle_1")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
