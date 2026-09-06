---
title: "fglsvgcanvas.destroy()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_destroy.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.destroy()"
type: "concept"
---

# fglsvgcanvas.destroy()

> Releases resources allocated for the SVG canvas.

## Syntax

```
FUNCTION destroy( cid SMALLINT )
```

1. cid is the SVG canvas id, as returned by
   `fglsvgcanvas.create()`.

## Usage

This function frees resources allocated for the SVG canvas handler that was created with the
[`create()`](2850-fglsvgcanvas-create.md "Creates a new SVG canvas handler.")
function.

## Example

```
CALL fglsvgcanvas.destroy( cid )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
