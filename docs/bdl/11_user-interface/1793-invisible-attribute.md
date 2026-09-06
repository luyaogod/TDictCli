---
title: "INVISIBLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_INVISIBLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > INVISIBLE attribute"
type: "concept"
---

# INVISIBLE attribute

> The INVISIBLE attribute prevents field data being readable on the screen.

## Syntax

```
INVISIBLE
```

## Usage

The `INVISIBLE` attribute can be used for `EDIT` and
`BUTTONEDIT` fields, to obfuscate the value of the field. A typical usage of
`INVISIBLE` is for password entry fields.

Characters that the user enters in a field with the `INVISIBLE` attribute are
obfuscated during data entry. When displaying data by program to a field (with [`DISPLAY TO / BY NAME`](1876-static-display-display-error-message-clear.md "This section explains the instructions displaying static information to application forms, such as DISPLAY, ERROR, MESSAGE, CLEAR.") or when the [interactive dialog](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.") synchronizes the form field content with
the program variable), the `INVISIBLE` attribute will also make the field content
unreadable.

## Related links

**Related concepts**  

[HIDDEN attribute](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.")

[KEYBOARDHINT attribute](1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly.")
