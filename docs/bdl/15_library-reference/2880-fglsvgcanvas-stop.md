---
title: "fglsvgcanvas.stop()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_stop.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.stop()"
type: "concept"
---

# fglsvgcanvas.stop()

> Produces an SVG "stop" element for gradients.

## Syntax

```
FUNCTION stop(
   offset STRING,
   color STRING,
   opacity STRING )
  RETURNS om.DomNode
```

1. offset defines the offset SVG attribute.
2. color defines the stop-color SVG attribute.
3. opacity defines the stop-opacity SVG attribte.

## Usage

This function creates a `"stop"` SVG DOM element from the parameters.

An SVG `"stop"` element defines the ramp of colors to use on a gradient in a
`"linearGradient"` or `"radialGradient"` element.

For a usage example, see [fglsvgcanvas.linearGradient()](2865-fglsvgcanvas-lineargradient.md "Produces an SVG \"linearGradient\" element.").

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
