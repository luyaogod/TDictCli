---
title: "CLEAR instruction in dialogs"
source: "fgl-topics/c_fgl_dialog_CLEAR_instruction_5.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control instructions > CLEAR instruction in dialogs"
type: "concept"
description: "The CLEAR field-list and CLEAR SCREEN ARRAY screen-array .* instructions clear the value buffer of specified form fields. The buffers are directly changed in the current form, and the program ..."
---

# CLEAR instruction in dialogs

The [`CLEAR
field-list`](1884-clear-field-list.md "The CLEAR field-list instruction clears specific fields in the current form.") and [`CLEAR SCREEN ARRAY
screen-array.*`](1883-clear-screen-array.md "The CLEAR SCREEN ARRAY instruction clears the values of all rows of the form list identified by the specified screen array.") instructions clear the value buffer of specified
form fields. The buffers are directly changed in the current form, and the program variables bound
to the dialog are left unchanged. `CLEAR` can be used outside any dialog instruction,
such as the [`DISPLAY BY NAME /
TO`](1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.") instructions.

When a dialog is configured with the `UNBUFFERED` mode, there is no reason to
clear field buffers since any variable assignment will synchronize field buffers. Actually, changing
the field buffers with `DISPLAY` or `CLEAR` instruction in an
`UNBUFFERED` dialog will have no visual effect, because the variables bound to the
dialog will be used to reset the field buffer just before giving control back to the user. To clear
fields of an `UNBUFEFERED` dialog, just set to `NULL` the variables
bound to the dialog. However, when using a `CONSTRUCT`, no program variables are
associated with the dialog and no `UNBUFFERED` concept exits, and the
`CLEAR` or `DISPLAY TO / BY NAME` instructions are the only way to
modify the `CONSTRUCT` fields.

A screen array with a screen-line specification doesn't make much
sense in a GUI application using `TABLE` containers,
you can therefore use the `CLEAR SCREEN ARRAY` instruction
to clear all rows of a list.

## Related links

**Related concepts**  

[Static display (DISPLAY/ERROR/MESSAGE/CLEAR)](1876-static-display-display-error-message-clear.md "This section explains the instructions displaying static information to application forms, such as DISPLAY, ERROR, MESSAGE, CLEAR.")
