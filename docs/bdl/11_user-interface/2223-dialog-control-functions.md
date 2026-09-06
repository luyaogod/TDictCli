---
title: "Dialog control functions"
source: "fgl-topics/c_fgl_prog_dialogs_control_functions.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Dialog control functions"
type: "concept"
---

# Dialog control functions

> The language provides several built-in functions and operators to be used in a dialog instruction.

Use the dialog functions and operators to keep track of the relative states of the current row,
the program array, and the screen array, or to access the field buffers and keystroke
buffers.

Typical control functions used in dialogs are: [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY."),
[`arr_count()`](../15_library-reference/2727-arr-count.md "Returns the number of rows entered during an INPUT ARRAY statement."), [`fgl_set_arr_curr()`](../15_library-reference/2782-fgl-set-arr-curr.md "Moves to a specific row in a record list."), [`set_count()`](../15_library-reference/2788-set-count.md "Defines the number of rows containing explicit data in a static array used by the next dialog."),
[`field_touched()`](../08_language-basics/0673-field-touched-function.md "The FIELD_TOUCHED() operator checks if fields were modified during the dialog execution."), [`GET_FLDBUF()`](../08_language-basics/0671-get-fldbuf-function.md "The GET_FLDBUF() operator returns as character strings the current values of the specified fields."), [`INFIELD()`](../08_language-basics/0672-infield-function.md "The INFIELD() operator checks for the current screen field."), [`fgl_dialog_getfieldname()`](../15_library-reference/2745-fgl-dialog-getfieldname.md "Returns the name of the current input field."), [`fgl_dialog_getbuffer()`](../15_library-reference/2743-fgl-dialog-getbuffer.md "Returns the text of the input buffer of the current field.").

As an alternative to functions and operators (especially for those
taking hard-coded parameters such as `INFIELD()`, use the
methods provided in the [ui.Dialog](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.")
class.
