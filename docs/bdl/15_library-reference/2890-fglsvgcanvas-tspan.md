---
title: "fglsvgcanvas.tspan()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_tspan.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.tspan()"
type: "concept"
---

# fglsvgcanvas.tspan()

> Produces an SVG "tspan" element.

## Syntax

```
FUNCTION tspan(
   x STRING,
   y STRING,
   dx STRING,
   dy STRING,
   style STRING,
   content STRING )
  RETURNS om.DomNode
```

1. x defines the X position of the tspan element.
2. y defines the Y position of the tspan element .
3. dx defines the delta-X for horizontal positioning of the tspan element.
4. dy defines the delta-Y for vertical positioning of the tspan element.
5. style defines the in-line style to be applied on the tspan element.
6. content is the actual text.

## Usage

This function creates a `"tspan"` SVG DOM element from the parameters.

A tspan element can be used to draw multiple lines of text in SVG.

The DOM nodes created by this function must be added to a `"text"` element
produced by the [`text()`](2886-fglsvgcanvas-text.md "Produces an SVG \"text\" element.") function.

## Example

```
DEFINE root_svg, t om.DomNode
...
LET t = fglsvgcanvas.text(NULL,200,NULL,"style_2")
CALL root_svg.appendChild( t )
CALL t.appendChild( fglsvgcanvas.tspan(120,NULL,NULL,30,NULL,"First line" ) ) 
CALL t.appendChild( fglsvgcanvas.tspan(120,NULL,NULL,30,NULL,"Second line" ) )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
