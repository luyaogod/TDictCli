---
title: "Canvas drawing functions"
source: "fgl-topics/r_fgl_Canvas_functions.html"
breadcrumb: "User interface > User interface programming > Canvases > Canvas drawing functions"
type: "reference"
description: "Important: This feature is deprecated, its use is discouraged although not prohibited. This table describes the helper functions provided to ease canvas usage. Use these functions or use the DOM API ..."
---

# Canvas drawing functions

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

This table describes the helper functions provided to ease canvas usage. Use these
functions or use the DOM API to directly create canvas elements in the form. The helper
functions are implemented in $FGLDIR/src/fgldraw.4gl. See the
source file for more details.

| Name | Description |
| --- | --- |
| `drawInit()` | Initializes the drawing API. It is mandatory to call this function at the beginning of your program, before the first display instruction. |
| `drawSelect()` | Selects a canvas area for drawing. |
| `drawDisableColorLines()` | By default, simple lines drawn with `drawLine()` are colored by `drawFillColor()`. Pass `TRUE` to the function to get black lines. |
| `drawFillColor()` | Defines the fill color for shapes and lines. Color value are named colors like "red", "green", "blue", etc. |
| `drawLineWidth()` | Defines the width of lines. |
| `drawAnchor()` | Defines the anchor hint for texts. |
| `drawLine()` | Draws a line in the selected canvas. |
| `drawCircle()` | Draws a circle in the selected canvas. |
| `drawArc()` | Draws an arc in the selected canvas. |
| `drawRectangle()` | Draws a rectangle in the selected canvas. |
| `drawOval()` | Draws an oval in the selected canvas. |
| `drawText()` | Draws a text in the selected canvas. |
| `drawPolygon()` | Draws a polygon in the selected canvas. |
| `drawClear()` | Clears the selected canvas. |
| `drawButtonLeft()` | Enables left mouse click on a canvas element. |
| `drawButtonRight()` | Enables right mouse click on a canvas element. |
| `drawClearButton()` | Disables all mouse clicks on a canvas element. |
| `drawGetClickedItemId()` | Returns the id of the last clicked canvas element |
