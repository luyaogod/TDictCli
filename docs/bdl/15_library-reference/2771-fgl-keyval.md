---
title: "fgl_keyval()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_KEYVAL.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_keyval()"
type: "concept"
---

# fgl_keyval()

> Returns the key code of a logical or physical key.

## Syntax

```
FUNCTION fgl_keyval(
   keyName STRING )
  RETURNS INTEGER
```

1. keyName can be a single character, a digit,
   a printable symbol like @, #, $ or a special keyword such as ACCEPT.

## Usage

`fgl_keyval()` can be used in form-related statements
to examine a value returned by the `fgl_lastkey()` and
`fgl_getkey()`
functions.

Key names recognized by `fgl_keyval()` are: `ACCEPT`,
`DELETE`, `DOWN`, `END`,
`ESC`/`ESCAPE`, `HELP`, `HOME`,
`INSERT`, `INTERRUPT`, `LEFT`,
`NEXT`/`NEXTPAGE`,
`PREVIOUS`/`PREVPAGE`, `RETURN`,
`RIGHT`, `SPACE`, `TAB`, `UP`,
`F1` through `F64`,
`CONTROL-character` (where character can be any
letter except A, D, H, I, J, L, M, R, or X).

The function returns NULL if the parameter does not correspond to a valid key.

If you specify a single character, `fgl_keyval()` considers the case and returns
the ASCII value of the character. In all other instances, the function ignores the case of its
argument, which can be uppercase or lowercase letters.

To determine whether the user has performed an action, such as inserting a row, specify the
logical name of the action (such as INSERT) rather than the name of the physical key (such as F1).
For example, the logical name of the Accept action is ACCEPT, while the default physical key is
ESCAPE. To test if the key most recently pressed by the user corresponds to the Accept action,
specify `fgl_keyval("ACCEPT")` rather than `fgl_keyval("ESCAPE")` or
`fgl_keyval("ESC")`. Otherwise, if a key other than ESCAPE is set as the Accept key
and the user presses that key, `fgl_lastkey()` does not return a code equal to
`fgl_keyval("ESCAPE")`.

This
function is provided for backward compatibility especially for TUI
mode applications. `fgl_keyval()` is well supported
in text mode, but this function can only be emulated in GUI mode,
because the front-ends communicate with the runtime system with other
events as keystrokes.

## Related links

**Related concepts**  

[fgl\_lastkey()](2772-fgl-lastkey.md "Returns the key code corresponding to the logical key that the user most recently typed in the form.")

[fgl\_getkey()](2763-fgl-getkey.md "Waits for a keystroke and returns the key number.")
