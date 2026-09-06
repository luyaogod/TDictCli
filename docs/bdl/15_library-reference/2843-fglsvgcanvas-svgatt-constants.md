---
title: "fglsvgcanvas.SVGATT_ constants"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_SVGATT.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.SVGATT_ constants"
type: "concept"
---

# fglsvgcanvas.SVGATT_ constants

> List of predefined SVG attributes.

## Syntax

```
CONSTANT SVGATT_TRANSFORM          = "transform"
CONSTANT SVGATT_CLASS              = "class"
CONSTANT SVGATT_STYLE              = "style"

CONSTANT SVGATT_ONCLICK            = "onclick"
CONSTANT SVGATT_ONMOUSEOVER        = "onmouseover"
CONSTANT SVGATT_ONMOUSEOUT         = "onmouseout"

CONSTANT SVGVAL_ELEM_CLICKED       = "elem_clicked(this)"
CONSTANT SVGVAL_ELEM_MOUSE_OVER    = "elem_mouse_over(this)"
CONSTANT SVGVAL_ELEM_MOUSE_OUT     = "elem_mouse_out(this)"

CONSTANT SVGATT_FILL               = "fill"
CONSTANT SVGATT_FILL_OPACITY       = "fill-opacity"
CONSTANT SVGATT_FILL_RULE          = "fill-rule"
CONSTANT SVGATT_STROKE             = "stroke"
CONSTANT SVGATT_STROKE_WIDTH       = "stroke-width"
CONSTANT SVGATT_STROKE_OPACITY     = "stroke-opacity"
CONSTANT SVGATT_STROKE_LINECAP     = "stroke-linecap"
CONSTANT SVGATT_STROKE_LINEJOIN    = "stroke-linejoin"
CONSTANT SVGATT_STROKE_MITERLIMIT  = "stroke-miterlimit"
CONSTANT SVGATT_STROKE_DASHARRAY   = "stroke-dasharray"
CONSTANT SVGATT_STROKE_DASHOFFSET  = "stroke-dashoffset"
CONSTANT SVGATT_FONT_FAMILY        = "font-family"
CONSTANT SVGATT_FONT_SIZE          = "font-size"
CONSTANT SVGATT_FONT_SIZE_ADJUST   = "font-size-adjust"
CONSTANT SVGATT_FONT_STRETCH       = "font-stretch"
CONSTANT SVGATT_FONT_STYLE         = "font-style"
CONSTANT SVGATT_FONT_VARIANT       = "font-variant"
CONSTANT SVGATT_FONT_WEIGHT        = "font-weight"
CONSTANT SVGATT_TEXT_ANCHOR        = "text-anchor"
CONSTANT SVGATT_MARKER_START       = "marker-start"
CONSTANT SVGATT_MARKER_MID         = "marker-mid"
CONSTANT SVGATT_MARKER_END         = "marker-end"
```

## Usage

The `SVGATT_` predefined constants are available from the fglsvgcanvas module.

Such constants are typically used to build an SVG attribute list for [`setAttributes()`](2877-fglsvgcanvas-setattributes.md "Sets the SVG attributes from an attribute set.") or [`styleAttributeList()`](2881-fglsvgcanvas-styleattributelist.md "Builds a string with a list of attributes to be used in a style attribute.")
