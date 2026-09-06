---
title: "fgl_dialog_setcursor()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_SETCURSOR.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_setcursor()"
type: "concept"
---

# fgl_dialog_setcursor()

> This function sets the position of the edit cursor in the current field.

## Syntax

```
FUNCTION fgl_dialog_setcursor(
   x INTEGER )
```

1. x is the edit cursor position in the text.

## Usage

The `fgl_dialog_setcursor()` moves the edit cursor to the specified position in the
current field. The function must be called in interactive instructions control blocks, when staying
in the current field.

This function has only an effect when staying in the current field, it is not recommended to call
it in an `AFTER FIELD` or `AFTER ROW` event for example.

Note that you can use `FGL_DIALOG_SETSELECTION()` to select a piece of text in a
field.

> **Important:**
>
> When using byte length semantics, the position is expressed in bytes. When
> using [char length semantics](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs."), the unit
> is characters. This matters when using a multibyte locale such as UTF-8.

## Related links

**Related concepts**  

[fgl\_dialog\_getcursor() / fgl\_getcursor()](2752-fgl-dialog-getcursor-fgl-getcursor.md "Returns the position of the edit cursor in the current field.")

[fgl\_dialog\_setselection()](2756-fgl-dialog-setselection.md "Selects the text in the current field.")
