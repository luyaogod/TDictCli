---
title: "fglsvgcanvas.display()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_display.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.display()"
type: "concept"
---

# fglsvgcanvas.display()

> Displays the SVG canvas.

## Syntax

```
FUNCTION display( cid SMALLINT )
```

1. cid is the SVG canvas id, as returned by
   `fglsvgcanvas.create()`.

## Usage

The function sends the SVG content to the front-end for display.

## Example

```
CALL fglsvgcanvas.display( cid )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
