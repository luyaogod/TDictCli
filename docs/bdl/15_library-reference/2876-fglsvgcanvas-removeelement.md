---
title: "fglsvgcanvas.removeElement()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_removeelement.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.removeElement()"
type: "concept"
---

# fglsvgcanvas.removeElement()

> Deletes an SVG element from the SVG canvas.

## Syntax

```
FUNCTION removeElement(
    node om.DomNode )
```

1. node is the `om.DomNode` to be removed.

## Usage

This function deletes an SVG element from the current SVG canvas.

The DOM node to remove is passed as second parameter.

The SVG DOM node object is de-referenced and destroyed.

## Example

```
DEFINE root_svg, g om.DomNode
...
LET g = fglsvgcanvas.g( "group1" )
CALL root_svg.appendChild( g )
...
CALL fglsvgcanvas.removeElement( g )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
