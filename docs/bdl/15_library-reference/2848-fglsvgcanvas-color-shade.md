---
title: "fglsvgcanvas.color_shade()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_color_shade.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.color_shade()"
type: "concept"
---

# fglsvgcanvas.color_shade()

> Applies a shade factor to an RGB color.

## Syntax

```
FUNCTION color_shade(
   source STRING,
   factor DECIMAL )
 RETURNS STRING
```

1. source is the source RGB color.
2. factor is the shade factor.

## Usage

This function modifies the RGB color specification by applying the shade factor passed as
parameter.

## Example

```
LET m = fglsvgcanvas.color_shade( "#FFAA34", 0.42 )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
