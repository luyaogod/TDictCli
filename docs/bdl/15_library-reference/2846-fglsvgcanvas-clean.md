---
title: "fglsvgcanvas.clean()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_clean.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.clean()"
type: "concept"
---

# fglsvgcanvas.clean()

> Deletes all SVG elements inside the SVG canvas.

## Syntax

```
FUNCTION clean( cid SMALLINT )
```

1. cid is the SVG canvas id, as returned by
   `fglsvgcanvas.create()`.

## Usage

This function cleans the SVG canvas.

The `clean()` function will not automatically display the cleaned SVG canvas. To
see a visual result, re-display the SVG content with the [`display()`](2855-fglsvgcanvas-display.md "Displays the SVG canvas.")
function.

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
