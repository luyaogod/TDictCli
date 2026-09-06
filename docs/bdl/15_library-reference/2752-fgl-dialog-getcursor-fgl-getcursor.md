---
title: "fgl_dialog_getcursor() / fgl_getcursor()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETCURSOR.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getcursor() / fgl_getcursor()"
type: "concept"
---

# fgl_dialog_getcursor() / fgl_getcursor()

> Returns the position of the edit cursor in the current field.

## Syntax

```
FUNCTION fgl_dialog_getcursor()
  RETURNS INTEGER
```

## Usage

The `fgl_dialog_getcursor()` function can be used in conjunction with
`fgl_dialog_getselectionend()` to get the position of the edit cursor and the piece
of text that is selected in the current field.

> **Important:**
>
> When using byte length semantics, the position is expressed in bytes. When
> using [char length semantics](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs."), the
> unit is characters. This matters when using a multibyte locale such as UTF-8.

## Related links

**Related concepts**  

[fgl\_dialog\_getselectionend()](2754-fgl-dialog-getselectionend.md "Returns the position of the last selected character in the current field.")

[fgl\_dialog\_setcursor()](2747-fgl-dialog-setcursor.md "This function sets the position of the edit cursor in the current field.")

[fgl\_dialog\_setselection()](2756-fgl-dialog-setselection.md "Selects the text in the current field.")
