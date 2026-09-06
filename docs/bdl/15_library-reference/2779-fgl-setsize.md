---
title: "fgl_setsize()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SETSIZE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_setsize()"
type: "concept"
---

# fgl_setsize()

> Sets the size of the main application window.

## Syntax

```
FUNCTION fgl_setsize(
   height INTEGER,
   width INTEGER )
```

1. height is the number of lines of the window.
2. width is the number of columns of the window.

## Usage

This function defines the size of the main window when using the
[traditional GUI mode](../11_user-interface/1519-graphical-mode-with-traditional-display.md).

Consider using the [`ui.Interface.setSize()`](3121-ui-interface-setsize.md "Specify the initial size of the window container.") method instead.

## Related links

**Related concepts**  

[fgl\_settitle()](2780-fgl-settitle.md "Sets the title of the current application window.")
