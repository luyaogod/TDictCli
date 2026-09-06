---
title: "fgl_getwin_height()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETWIN_HEIGHT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getwin_height()"
type: "concept"
---

# fgl_getwin_height()

> Returns the number of rows of the current window.

## Syntax

```
FUNCTION fgl_getwin_height()
  RETURNS INTEGER
```

## Usage

The `fgl_getwin_height()` function returns the height of the
current window, in character units.

This function is provided for text mode applications, in GUI mode, windows are resizable and thus
their height is variable.

## Related links

**Related concepts**  

[fgl\_getwin\_width()](2768-fgl-getwin-width.md "Returns the width of the current window as a number of columns.")
