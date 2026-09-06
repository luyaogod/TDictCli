---
title: "fgl_dialog_setbuffer()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_SETBUFFER.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_setbuffer()"
type: "concept"
---

# fgl_dialog_setbuffer()

> Sets the input buffer of the current field.

## Syntax

```
FUNCTION fgl_dialog_setbuffer(
   value STRING )
```

1. value is the text to set in the current input buffer.

## Usage

In the default buffered input mode, this function modifies the input
buffer of the current field; the corresponding input variable is not
assigned. It makes no sense to call this function in `BEFORE
FIELD` blocks of `INPUT` and `INPUT
ARRAY`. However, if the statement is using the `UNBUFFERED`
mode, the function will set both the field buffer and the program variable.
If the string set by the function does not represent a valid value
that can be stored by the program variable, the buffer and the variable
will be set to [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.").

The `fgl_dialog_setbuffer()` function must be used in
`INPUT`, `INPUT ARRAY` and
`CONSTRUCT` blocks.

This function sets the modification flag for both [`FIELD_TOUCHED()`](../08_language-basics/0673-field-touched-function.md "The FIELD_TOUCHED() operator checks if fields were modified during the dialog execution.") and [`fgl_buffertouched()`](2736-fgl-buffertouched.md "Returns TRUE if the input buffer was modified in the current field.")
functions. There is a slight difference between both functions. The modification flag for
`fgl_buffertouched()` is reset to `FALSE` when entering the
field.

The function is especially useful in a [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") instruction, because there is no variable associated with
fields in this case.

## Related links

**Related concepts**  

[fgl\_dialog\_getbuffer()](2743-fgl-dialog-getbuffer.md "Returns the text of the input buffer of the current field.")

[The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
