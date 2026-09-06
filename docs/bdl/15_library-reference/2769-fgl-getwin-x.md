---
title: "fgl_getwin_x()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETWIN_X.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getwin_x()"
type: "concept"
---

# fgl_getwin_x()

> Returns the horizontal position of the current window.

## Syntax

```
FUNCTION fgl_getwin_x()
  RETURNS INTEGER
```

## Usage

The `fgl_getwin_x()` function returns the horizontal coordinate
of the top/left corner of the current window.

This function is provided for text mode applications, in GUI mode, windows are
movable and thus their position is variable.

## Related links

**Related concepts**  

[fgl\_getwin\_y()](2770-fgl-getwin-y.md "Returns the vertical position of the current window.")
