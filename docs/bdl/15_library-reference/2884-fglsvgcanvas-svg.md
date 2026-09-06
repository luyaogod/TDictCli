---
title: "fglsvgcanvas.svg()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_svg.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.svg()"
type: "concept"
---

# fglsvgcanvas.svg()

> Produces an SVG "svg" element.

## Syntax

```
FUNCTION svg(
   id STRING,
   x STRING,
   y STRING,
   width STRING,
   height STRING,
   viewBox STRING,
   preserveAspectRatio STRING
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. x, `y`, `width`, height
   define the SVG viewport.
3. viewBox defines the SVG viewbox.
4. `preserveAspectRatio` is the aspect ratio to preserve.

## Usage

This function creates a "svg" SVG DOM element definition from the parameters.

SVG allows you to create nested svg sub-elements with their own coordinate system.

Build a tree of svg elements as you need.

The root svg element must be created with the [`setRootSVGAttributes()`](2879-fglsvgcanvas-setrootsvgattributes.md "Produces the root SVG element.") function.

The viewport is defined by the `x`, `y`, `width`
and `height` attributes.

The viewbox is defined by the `viewBox` string for example as `"0 0 100
100"`.

See SVG reference documentation for more details about viewport, viewbox and aspect ratio
concepts.

## Example

```
DEFINE root_svg, n om.DomNode
...
LET n = fglsvgcanvas.svg("day_1",
           NULL, NULL, NULL, NULL,
           "0 0 500 500",
           "xMidYMid meet"
        )
...
CALL root_svg.appendChild( n )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
