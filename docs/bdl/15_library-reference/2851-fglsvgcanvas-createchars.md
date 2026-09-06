---
title: "fglsvgcanvas.createChars()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_createchars.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.createChars()"
type: "concept"
---

# fglsvgcanvas.createChars()

> Produces an SVG DOM text node.

## Syntax

```
FUNCTION createChars( value STRING )
  RETURNS om.DomNode
```

1. value is the value to be set in the text node.

## Usage

This function creates an SVG DOM text node from the value passed as parameter.

Use the `createChars()` function when you need to create an SVG text node
that cannot be created with one of the fglsvgcanvas functions (such as [`title()`](2889-fglsvgcanvas-title.md "Produces an SVG \"title\" element.")).

## Example

```
DEFINE t, c om.DomNode
LET t = fglsvgcanvas.createElement( "title", "id1" )
LET c = fglsvgcanvas.createChars( "my value" )
CALL t.appendChild( c )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
