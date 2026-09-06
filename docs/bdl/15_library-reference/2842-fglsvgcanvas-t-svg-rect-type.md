---
title: "fglsvgcanvas.t_svg_rect type"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_t_svg_rect.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.t_svg_rect type"
type: "concept"
---

# fglsvgcanvas.t_svg_rect type

> The t_svg_rect type defines the position and dimensions of a rectangle.

## Syntax

```
TYPE t_svg_rect RECORD
    x DECIMAL,
    y DECIMAL,
    width DECIMAL,
    height DECIMAL
  END RECORD
```

## Usage

This user-defined type defines a record structure with the position (`x`,
`y`) and size (`width`, `height`) of a rectangle.

The structure is for example used with the [fglsvgcanvas.getBBox()](2860-fglsvgcanvas-getbbox.md "Returns the bounding box of an SVG element.") function.
