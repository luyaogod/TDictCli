---
title: "fgl_getwin_y()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETWIN_Y.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getwin_y()"
type: "concept"
---

# fgl_getwin_y()

> Returns the vertical position of the current window.

## Syntax

```
FUNCTION fgl_getwin_y()
  RETURNS INTEGER
```

## Usage

The `fgl_getwin_y()` function returns the vertical coordinate
of the top/left corner of the current window.

This function is provided for text mode applications, in GUI mode, windows
are movable and thus their position is variable.

## Related links

**Related concepts**  

[fgl\_getwin\_x()](2769-fgl-getwin-x.md "Returns the horizontal position of the current window.")
