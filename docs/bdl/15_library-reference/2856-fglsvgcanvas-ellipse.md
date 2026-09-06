---
title: "fglsvgcanvas.ellipse()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_ellipse.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.ellipse()"
type: "concept"
---

# fglsvgcanvas.ellipse()

> Produces an SVG "ellipse" element.

## Syntax

```
FUNCTION ellipse(
   cx STRING,
   cy STRING,
   rx STRING,
   ry STRING )
  RETURNS om.DomNode
```

1. cx defines the X coordinate of the center of the ellipse.
2. cy defines the Y coordinate of the center of the ellipse.
3. `rx` defines the X radius of the ellipse.
4. `ry` defines the Y radius of the ellipse.

## Usage

This function creates an `"ellipse"` SVG DOM element from the parameters.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.ellipse(100,100,50,60)
CALL n.setAttribute(SVGATT_STYLE,"stroke:#006600;")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
