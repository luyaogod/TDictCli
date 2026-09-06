---
title: "fgl_lastkey()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_LASTKEY.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_lastkey()"
type: "concept"
---

# fgl_lastkey()

> Returns the key code corresponding to the logical key that the user most recently typed in the form.

## Syntax

```
FUNCTION fgl_lastkey()
  RETURNS INTEGER
```

## Usage

The `fgl_lastkey()` function returns a numeric code corresponding to the user's
last keystroke before the function was called. For example, if the last key that the user pressed
was a lowercase `a`, the function returns the code `61`.

> **Note:**
>
> The value of `fgl_lastkey()` is undefined in a [`MENU`](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.") statement.

It is not required to know the specific key codes returned by `fgl_lastkey()`: The
`fgl_keyval()` function can be used to compare the key code of the last key pressed.
The  `fgl_keyval()` function allows you to compare the last key pressed with a
logical of physical key. For example, you do not need to know the physical key defined to validate a
dialog, you can use the logical name "accept" instead. For a complete list of key codes and logical
key names, see the [Key code
table](2793-the-key-code-table.md).

Pay attention to the fact that this function is provided for backward compatibility: The abstract
user interface protocol is based on logical events, not only key events. For example, in GUI mode,
when selecting a new row with the mouse in a table, there is no key press as when moving in a screen
array in TUI mode. However, the runtime system tries to emulate as much as possible keystrokes from
non-keystroke events like mouse clicks. Consider reviewing the code logic, in order to use control
events of the dialog instruction such as `BEFORE FIELD`, `AFTER
FIELD`, `BEFORE ROW`, `AFTER ROW`, to detect field and record
list navigation events.

## Related links

**Related concepts**  

[fgl\_keyval()](2771-fgl-keyval.md "Returns the key code of a logical or physical key.")
