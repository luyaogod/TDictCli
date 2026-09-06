---
title: "fgl_getkey()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETKEY.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getkey()"
type: "concept"
---

# fgl_getkey()

> Waits for a keystroke and returns the key number.

## Syntax

```
FUNCTION fgl_getkey()
  RETURNS INTEGER
```

## Usage

`fgl_getkey()` waits for a keystroke and returns the
[key code](2793-the-key-code-table.md)
corresponding to the pressed physical key.

It is recommended to only use this function in text mode.

Unlike `fgl_lastkey()`, which can return a value indicating
the logical effect of whatever key the user pressed,`fgl_getkey()`
returns an integer representing the key code of the physical key that the user pressed.
The `fgl_getkey()` function recognizes the same codes for keys that the
[`fgl_keyval()`](2771-fgl-keyval.md "Returns the key code of a logical or physical key.")
function returns. Unlike `fgl_keyval()`, which can only return keystrokes
that are entered during dialogs, `fgl_getkey()` can be called outside
a dialog context.

## Related links

**Related concepts**  

[fgl\_lastkey()](2772-fgl-lastkey.md "Returns the key code corresponding to the logical key that the user most recently typed in the form.")
