---
title: "fgl_settitle()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SETTITLE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_settitle()"
type: "concept"
---

# fgl_settitle()

> Sets the title of the current application window.

## Syntax

```
FUNCTION fgl_settitle(
   newTitle STRING )
```

1. newTitle is the text of the title.

## Usage

The `fgl_settitle()` function defines the title of the current window, as well
as the default title for new created windows.

This function is provided for backward compatibility, the title of a window can be defined
with the [`TEXT`](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
attribute of a `LAYOUT` section.

## Related links

**Related concepts**  

[ui.Window.setText](3139-ui-window-settext.md "Set the window title.")
