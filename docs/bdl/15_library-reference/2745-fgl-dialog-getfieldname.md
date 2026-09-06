---
title: "fgl_dialog_getfieldname()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETFIELDNAME.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getfieldname()"
type: "concept"
---

# fgl_dialog_getfieldname()

> Returns the name of the current input field.

## Syntax

```
FUNCTION fgl_dialog_getfieldname()
  RETURNS STRING
```

## Usage

This function returns the name of the current input field during a dialog
execution. It must be use in `INPUT`, `INPUT ARRAY`
or `CONSTRUCT` blocks.

Only the column part of the field name is returned (screen record name is
omitted).

The `fgl_dialog_getfieldname()` is similar to the
[`INFIELD()`](../08_language-basics/0672-infield-function.md "The INFIELD() operator checks for the current screen field.")
operator and [`fgl_dialog_infield()`](2746-fgl-dialog-infield.md "This function checks for the current input field.")
function.
