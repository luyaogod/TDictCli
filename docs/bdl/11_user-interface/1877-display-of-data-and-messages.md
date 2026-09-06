---
title: "Display of data and messages"
source: "fgl-topics/c_fgl_record_display_002.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > Display of data and messages"
type: "concept"
---

# Display of data and messages

> The values contained in program variables can be displayed to the current form with the DISPLAY BY NAME or DISPLAY TO instruction.

Forms can be cleared with the [`CLEAR
FORM`](1882-clear-form.md "The CLEAR FORM instruction clears all fields in the current form.") or [`CLEAR
field-list`](1884-clear-field-list.md "The CLEAR field-list instruction clears specific fields in the current form.") instructions. Complete record lists (in
`SCROLLGRID`, `TABLE` or `TREE` containers) can be
cleared with the [`CLEAR SCREEN
ARRAY`](1883-clear-screen-array.md "The CLEAR SCREEN ARRAY instruction clears the values of all rows of the form list identified by the specified screen array.") instruction.

Application messages and warnings can be displayed to the user with the [`MESSAGE`](1885-message.md "The MESSAGE instruction displays a message to the user.") and [`ERROR`](1886-error.md "The ERROR instruction displays an error message to the user.") instructions.

The `DISPLAY BY
NAME`, [`DISPLAY
TO`](1880-display-to.md "The DISPLAY ... TO instruction displays data to specific form fields.") instructions are not interactive, and are usually not needed, if the program is
always in the context of a dialog controlling the form fields:

- When the dialog starts, the data of the program variables will be displayed in form fields if
  the [`WITHOUT
  DEFAULTS`](1936-input-instruction-configuration.md) option is specified.
- During the dialog execution, form fields will be automatically synchronized with the program
  variables when using the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") option. With this option, setting the program variable is
  sufficient to show the value to the user in the form fields.

## Related links

**Related concepts**  

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")
