---
title: "fgl_dialog_setfieldorder()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_SETFIELDORDER.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_setfieldorder()"
type: "concept"
---

# fgl_dialog_setfieldorder()

> This function enables or disables field order constraint.

## Syntax

```
FUNCTION fgl_dialog_setfieldorder(
   constrainded INTEGER )
```

1. When constrainded is `TRUE`, the field
   order is constrained. When constrainded is `FALSE`,
   the field order is not constrained.

## Usage

Typical applications control user input with `BEFORE FIELD`
and `AFTER FIELD` blocks.
In many cases the field order and the sequential execution of `AFTER
FIELD` blocks is important in order to validate the data
entered by the user. But with graphical front-ends you can use the
mouse to move to a field. By default the runtime system executes
all `BEFORE FIELD` and `AFTER FIELD` blocks
of the fields used by the interactive instruction, from the origin
field to the target field selected by mouse click. If needed, you
can force the runtime system to ignore all intermediate field triggers,
by calling this function with a `FALSE` attribute.

This function must be called outside interactive dialog blocks, typically
at the beginning of the program.

Consider using the [`Dialog.fieldOrder`](../11_user-interface/2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior.")
parameter when all programs are affected. The FGLPROFILE profile entry is
the default when the `fgl_dialog_setfieldorder()`
function is not used.

Consider using [`OPTIONS FIELD ORDER FORM`](../09_advanced-features/0928-defining-field-tabbing-order-method.md)
for new developments with graphical rendering.
