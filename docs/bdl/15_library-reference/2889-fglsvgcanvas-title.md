---
title: "fglsvgcanvas.title()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_title.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.title()"
type: "concept"
---

# fglsvgcanvas.title()

> Produces an SVG "title" element.

## Syntax

```
FUNCTION title( text STRING )
  RETURNS om.DomNode
```

1. text defines text for the title element.

## Usage

This function creates a `"title"` SVG DOM element from the text provided as
parameter.

In SVG, a title element is usually rendered with a tooltip on desktop platforms.

The title element must be the first child of the parent SVG element. See SVG specification for
more details.

## Example

```
DEFINE g om.DomNode
LET g = fglsvgcanvas.g( "g1" )
CALL g.appendChild( fglsvgcanvas.title("This is my text") )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
