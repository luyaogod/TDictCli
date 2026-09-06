---
title: "fgl_drawbox()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DRAWBOX.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_drawbox()"
type: "concept"
---

# fgl_drawbox()

> Draws a rectangle in the current window.

## Syntax

```
FUNCTION fgl_drawbox(
   height INTEGER,
   width INTEGER,
   posY INTEGER,
   posX INTEGER,
   color INTEGER )
```

1. height is the height of the rectangle.
2. width is the width of the rectangle.
3. posY is the vertical coordinate (line) of the upper side of the rectangle.
4. posX is the horizontal coordinate (column) of the left side of the rectangle.
5. color is the color number (ignored).

## Usage

The `fgl_drawbox()` function draws a rectangle based on the character
terminal coordinates in the current open window.

Dimensions and coordinates are specified in grid cell units (characters).

This function is provided for backward compatibility. A call to this function will
be ignored if the current window is not SCREEN based. The function is supported
to draw rectangles in text mode applications.

## Related links

**Related concepts**  

[fgl\_drawline()](2758-fgl-drawline.md "Draws a line in the current window (TUI and traditional mode).")
