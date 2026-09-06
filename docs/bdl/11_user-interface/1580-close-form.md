---
title: "CLOSE FORM"
source: "fgl-topics/c_fgl_windows_and_forms_CLOSE_FORM.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > CLOSE FORM"
type: "concept"
---

# CLOSE FORM

> Closes the resources allocated by OPEN FORM.

## Syntax

```
CLOSE FORM identifier
```

1. identifier is the name of the form.

## Usage

The `CLOSE FORM` instruction releases the memory allocated to the form.

A form associated with a window by the `OPEN WINDOW WITH FORM` instruction is automatically closed when the program closes the
window with a `CLOSE WINDOW` instruction.

## Related links

**Related concepts**  

[CLOSE WINDOW](1575-close-window.md "Closes and destroys a window.")

[OPEN FORM](1578-open-form.md "Declares a compiled form in the program.")
