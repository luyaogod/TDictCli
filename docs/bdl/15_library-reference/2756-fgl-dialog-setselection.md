---
title: "fgl_dialog_setselection()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_SETSELECTION.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_setselection()"
type: "concept"
---

# fgl_dialog_setselection()

> Selects the text in the current field.

## Syntax

```
FUNCTION fgl_dialog_setselection(
   start INTEGER,
   end INTEGER )
```

1. start defines the edit cursor position.
2. end defines the selection end position.

## Usage

A call to `fgl_dialog_setselection(cursor, end)` sets the text
selection in the current form field. The start parameter defines the character
position of the edit cursor (equivalent to `fgl_dialog_getcursor()` position),
while end defines the character position of the end of the text selection
(equivalent to `fgl_dialog_getselectionend()` position).

> **Important:**
>
> When using byte length semantics, the positions are expressed in bytes.
> When using [char length semantics](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs."),
> the unit is characters. This matters when using a multibyte locale such as UTF-8.

start can be lower, greater or equal to end.

This function has only an effect when staying in the current field, it is not recommended to
call it in an `AFTER FIELD` or `AFTER ROW` event for example.

## Related links

**Related concepts**  

[fgl\_dialog\_getcursor() / fgl\_getcursor()](2752-fgl-dialog-getcursor-fgl-getcursor.md "Returns the position of the edit cursor in the current field.")

[fgl\_dialog\_getselectionend()](2754-fgl-dialog-getselectionend.md "Returns the position of the last selected character in the current field.")
