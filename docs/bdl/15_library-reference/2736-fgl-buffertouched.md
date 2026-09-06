---
title: "fgl_buffertouched()"
source: "fgl-topics/c_fgl_BuiltInFunctions_BUFFERTOUCHED.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_buffertouched()"
type: "concept"
---

# fgl_buffertouched()

> Returns TRUE if the input buffer was modified in the current field.

## Syntax

```
FUNCTION fgl_buffertouched()
  RETURNS INTEGER
```

## Usage

The function returns `TRUE` if the input buffer has been modified after the
current field has got the focus.

Call this function in `AFTER FIELD`, `AFTER INPUT`,
`AFTER CONSTRUCT`, `ON KEY`, `ON ACTION`
blocks.

This function is not equivalent to [`FIELD_TOUCHED()`](../08_language-basics/0673-field-touched-function.md "The FIELD_TOUCHED() operator checks if fields were modified during the dialog execution.") or [`DIALOG.getFieldTouched()`](3198-ui-dialog-getfieldtouched.md "Returns the modification flag for a field."): The modification status of
`fgl_buffertouched()` is reset when entering a new field, while
`FIELD_TOUCHED()`/`DIALOG.getFieldTouched()` returns
`TRUE`, when a field was modified during the interactive instruction.

## Related links

**Related concepts**  

[Input field modification flag](../11_user-interface/2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.")
