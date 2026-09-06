---
title: "fgl_getwin_width()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETWIN_WIDTH.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getwin_width()"
type: "concept"
---

# fgl_getwin_width()

> Returns the width of the current window as a number of columns.

## Syntax

```
FUNCTION fgl_getwin_width()
  RETURNS INTEGER
```

## Usage

The `fgl_getwin_width()` function returns the width of the current
window, in character units.

This function is provided for text mode applications, in GUI mode, windows are resizable and thus
their width is variable.

## Related links

**Related concepts**  

[fgl\_getwin\_height()](2767-fgl-getwin-height.md "Returns the number of rows of the current window.")
