---
title: "fglsvgcanvas.getItemid()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_getitemid.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.getItemid()"
type: "concept"
---

# fglsvgcanvas.getItemid()

> Returns SVG element id after a user action.

## Syntax

```
FUNCTION getItemId( cid SMALLINT )
     RETURNS STRING
```

1. cid is the SVG canvas id, as returned by
   `fglsvgcanvas.create()`.

## Usage

This function returns the `id` attribute of the SVG element involved in a
user-defined action.

When the fglsvgcanvas web component does not have the focus, it is not possible to return the SVG
element id in the field value. If user-defined actions are bound to SVG events that can be fired
when the field does not have the focus, the only way to identify the SVG element is the
`getItemId()` function.

## Example

```
DEFINE id STRING
...
ON ACTION mouse_over
   LET id = fglsvgcanvas.getItemId(cid)
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
