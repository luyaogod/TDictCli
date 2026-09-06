---
title: "fglsvgcanvas.setCurrent()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_setcurrent.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.setCurrent()"
type: "concept"
---

# fglsvgcanvas.setCurrent()

> Selects the SVG canvas handler for subsequent SVG canvas API calls.

## Syntax

```
FUNCTION setCurrent(
   cid SMALLINT )
```

1. cid is the SVG canvas id, as returned by `fglsvgcanvas.create()`.

## Usage

This function selects the SVG canvas identified by the specified id handler,
for other fglsvgcanvas API calls.

Any subsequent calls to an fglsvgcanvas function that requires the SVG canvas handler will be
done with this identifier.

## Example

```
DEFINE cid1, cid2 SMALLINT
LET cid1 = fglsvgcanvas.create("formonly.canvas1") -- current canvas is cid1
LET cid2 = fglsvgcanvas.create("formonly.canvas1") -- current canvas is cid2
CALL fglsvgcanvas.setcurrent( cid1 )               -- current canvas is cid1
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
