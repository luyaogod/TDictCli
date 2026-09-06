---
title: "DISPLAY TO / BY NAME in dialogs"
source: "fgl-topics/c_fgl_DIALOG_instr_DISPLAY_TO_BY_NAME.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control instructions > DISPLAY TO / BY NAME in dialogs"
type: "concept"
description: "The DISPLAY variable-list TO field-list or DISPLAY BY NAME variable-list instruction fills the value buffers of specified form fields with the values contained in the specified program variables. The ..."
---

# DISPLAY TO / BY NAME in dialogs

The `DISPLAY variable-list TO field-list` or `DISPLAY BY NAME
variable-list` instruction fills the value buffers of specified form fields with the
values contained in the specified program variables. The [`DISPLAY TO / BY NAME`](1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.")
instruction changes the buffers directly in the current form, not the program variables bound to the
dialog. `DISPLAY` can be used outside any dialog instruction, in the same way as the
[`CLEAR`](1884-clear-field-list.md "The CLEAR field-list instruction clears specific fields in the current form.") instruction.
`DISPLAY` also sets the modification flag of fields.

As `DIALOG` is typically used with the [`UNBUFFERED` mode](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields."), there is no
reason to set field buffers in a `DIALOG` block since any variable assignment will
synchronize field buffers. Actually, changing the field buffers with the `DISPLAY` or
`CLEAR` instruction will have no visual effect if the fields are used by a dialog
working in `UNBUFFERED` mode, because the variables bound to the dialog will be used
to reset the field buffer just before giving control back to the user. So if you want to set field
values, just assign the variables and the fields will be synchronized. However, when using a [`CONSTRUCT`](2084-the-construct-sub-dialog.md "The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.") binding, you may want
to set field buffers with this `DISPLAY` instruction, as there are no program
variables bound to fields (with `CONSTRUCT`, only one string variable is bound to
hold the SQL condition).

Instead of using a `DISPLAY` instruction to set the modification flag of fields to
simulate user input, use the [`DIALOG.setFieldTouched()`](../15_library-reference/3227-ui-dialog-setfieldtouched.md "Sets the modification flag of the specified field.") method instead.

## Related links

**Related concepts**  

[Input field modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.")
