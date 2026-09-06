---
title: "fglsvgcanvas.path()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_path.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.path()"
type: "concept"
---

# fglsvgcanvas.path()

> Produces an SVG "path" element.

## Syntax

```
FUNCTION path(
   d STRING )
  RETURNS om.DomNode
```

1. d is the SVG path specification.

## Usage

This function creates a `"path"` SVG DOM element from the parameters.

The path element draws complex SVG shapes combined from lines, arcs, curves, etc.

For more details about path elements, see SVG specification.

## Example

Displaying an image from a URL:

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.path("M50,50 A30,30 0 0,1 35,20 ...")
CALL n.setAttribute(SVGATT_STYLE, "stroke:#006600; stroke-width:0.4")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
