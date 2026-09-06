---
title: "fgl_dialog_getselectionend()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETSELECTIONEND.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getselectionend()"
type: "concept"
---

# fgl_dialog_getselectionend()

> Returns the position of the last selected character in the current field.

## Syntax

```
FUNCTION fgl_dialog_getselectionend()
  RETURNS INTEGER
```

## Usage

The `fgl_dialog_getselectionend()` function returns the edit cursor position of
the last selected character in the text of the current field.

> **Important:**
>
> When using byte length semantics, the position is expressed in bytes. When
> using [char length semantics](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs."), the
> unit is characters. This matters when using a multibyte locale such as UTF-8.

The function returns zero if the complete text is selected.

The edit cursor position returned by `fgl_dialog_getcursor()` will be lower than
the position returned by `fgl_dialog_getselectionend()` if the text has been selected
backwards.

## Related links

**Related concepts**  

[fgl\_dialog\_getcursor() / fgl\_getcursor()](2752-fgl-dialog-getcursor-fgl-getcursor.md "Returns the position of the edit cursor in the current field.")

[fgl\_dialog\_setselection()](2756-fgl-dialog-setselection.md "Selects the text in the current field.")
