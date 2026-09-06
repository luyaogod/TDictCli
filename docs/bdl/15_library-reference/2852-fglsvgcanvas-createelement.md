---
title: "fglsvgcanvas.createElement()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_createelement.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.createElement()"
type: "concept"
---

# fglsvgcanvas.createElement()

> Produces an SVG DOM element with the tag name specified as parameter.

## Syntax

```
FUNCTION createElement(
   tagName STRING,
   id STRING )
  RETURNS om.DomNode
```

1. tagName is the DOM tag name.
2. id is the SVG object identifier.

## Usage

This function creates an SVG DOM element with the specified tag name.

Use the `createElement()` function when you need to create an SVG element
that cannot be created with one of the fglsvgcanvas functions.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.createElement( "feBlend", "myblend" )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
