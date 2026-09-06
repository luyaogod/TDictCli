---
title: "fgl_dialog_getbuffer()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETBUFFER.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getbuffer()"
type: "concept"
---

# fgl_dialog_getbuffer()

> Returns the text of the input buffer of the current field.

## Syntax

```
FUNCTION fgl_dialog_getbuffer()
  RETURNS STRING
```

## Usage

The `fgl_dialog_getbuffer()` function returns the content of the
input buffer of the current field. It must be used in `INPUT`,
`INPUT ARRAY` and `CONSTRUCT` blocks.

The function is especially useful in a [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") instruction, because there is no variable associated to fields in
this case.

Consider using the [`ui.Dialog.getFieldBuffer()`](3197-ui-dialog-getfieldbuffer.md "Returns the input buffer of the specified field.") method instead.

## Related links

**Related concepts**  

[fgl\_dialog\_setbuffer()](2744-fgl-dialog-setbuffer.md "Sets the input buffer of the current field.")

[GET\_FLDBUF() [function]](../08_language-basics/0671-get-fldbuf-function.md "The GET_FLDBUF() operator returns as character strings the current values of the specified fields.")
