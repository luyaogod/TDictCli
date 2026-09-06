---
title: "fgl_drawline()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DRAWLINE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_drawline()"
type: "concept"
---

# fgl_drawline()

> Draws a line in the current window (TUI and traditional mode).

## Syntax

```
FUNCTION fgl_drawline(
   posY INTEGER,
   posX INTEGER,
   width INTEGER,
   type CHAR(1),
   color INTEGER)
```

1. posY is the vertical coordinate (line) of the start of the line.
2. posX is the horizontal coordinate (column) of the start of the line.
3. width is the width of the line.
4. type (ignored).
5. color is the color number (ignored).

## Usage

The `fgl_drawline()` function draws a line based on the character
terminal coordinates in the current open window.

Dimensions and coordinates are specified in grid cell units (characters).

This function is provided for backward compatibility. A call to this function
will be ignored if the current window is not SCREEN based. The function is supported
to draw lines in text mode applications.

## Related links

**Related concepts**  

[fgl\_drawbox()](2757-fgl-drawbox.md "Draws a rectangle in the current window.")
