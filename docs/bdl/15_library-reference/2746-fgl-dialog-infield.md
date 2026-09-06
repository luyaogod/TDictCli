---
title: "fgl_dialog_infield()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_INFIELD.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_infield()"
type: "concept"
---

# fgl_dialog_infield()

> This function checks for the current input field.

## Syntax

```
FUNCTION fgl_dialog_infield(
   name STRING )
  RETURNS INTEGER
```

1. name is the name if the form field.

## Usage

The `fgl_dialog_infield()` function returns
`TRUE` if the field name passed as the parameter
is the current input field.

The function must be called in `INPUT`, `INPUT
ARRAY` or `CONSTRUCT` blocks.

This function is the equivalent of the [`INFIELD()`](../08_language-basics/0672-infield-function.md "The INFIELD() operator checks for the current screen field.")
operator, except that the function takes a string expression as parameter,
while the `INFIELD()` operator expects a hard-coded form
field name.

## Related links

**Related concepts**  

[Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
