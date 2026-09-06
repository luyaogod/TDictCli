---
title: "fgl_scr_size()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SCR_SIZE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_scr_size()"
type: "concept"
---

# fgl_scr_size()

> Returns the size of the specified screen array in the current form.

## Syntax

```
FUNCTION fgl_scr_size(
   name STRING )
  RETURNS INTEGER
```

1. name is the name of a screen-array in the current displayed form.

## Usage

The `fgl_scr_size()` function takes the name of a screen array as parameter
identifying an array in the currently-opened form and returns an integer that corresponds to the
number of screen records in that screen array.

This function is typically used with traditional text mode forms
having [screen
arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") with a constant size, to display data in screen array
rows with the `DISPLAY TO` instruction.

For modern GUI applications, consider using the
[UNBUFFERED](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
mode in dialogs, to get automatic form field synchronization with
program variables.

Error [-1108](4483-genero-bdl-errors.md)
will be raised if the passed screen-array does not exits in the current
form, and error [-1114](4483-genero-bdl-errors.md)
is returned if no form is currently displayed.
