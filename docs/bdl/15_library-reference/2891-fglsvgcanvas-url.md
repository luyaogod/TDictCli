---
title: "fglsvgcanvas.url()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_url.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.url()"
type: "concept"
---

# fglsvgcanvas.url()

> Produces a "url(#name)" reference for SVG elements.

## Syntax

```
FUNCTION url( name STRING )
 RETURNS STRING
```

1. name is the URL name.

## Usage

This function builds a `"url(#name)"` SVG attribute with the
name passed as parameter.

This URL attribute is typically used in SVG elements to reference a common reusable element
defined in the `"defs"` element, such as SVG patterns.

## Example

```
CALL n.setAttribute(SVGATT_STYLE,
   SFMT('stroke:gray; fill:%1;', fglsvgcanvas.url("pattern_1") )
 )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
