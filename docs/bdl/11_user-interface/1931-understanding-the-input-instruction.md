---
title: "Understanding the INPUT instruction"
source: "fgl-topics/c_fgl_record_input_002.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Understanding the INPUT instruction"
type: "concept"
---

# Understanding the INPUT instruction

> The INPUT instruction controls a single record input from form fields.

The `INPUT` statement binds [program
variables](../08_language-basics/0686-variables.md "Explains how to define program variables.") to [screen-records](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") for
data entry in form fields. The `INPUT` statement uses the current form in the current
window. Before executing the `INPUT` statement, record data must be fetched from the
database table into the program variables using the input statement.

An `INPUT` will start with empty fields, unless the [`WITHOUT DEFAULTS`](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.") option is used.
Using `WITHOUT DEFAULTS` clause instructs the `INPUT` dialog to use
the current data of the program variables/record bound to the dialog.

During the `INPUT` statement execution, the user can edit the record fields, while
the program controls the behavior of the instruction with control blocks.

When the user moves from field to field or changes values, [dialog control blocks](1938-input-control-blocks.md) such as `BEFORE
FIELD` are executed.

When the user clicks on an action view (button), or when an asynchronous event occurs, [dialog interaction blocks](1945-input-interaction-blocks.md) like `ON ACTION`
are executed.

The code inside an `INPUT` dialog can use [control instructions](1950-input-control-instructions.md), [dialog control
functions](2223-dialog-control-functions.md "The language provides several built-in functions and operators to be used in a dialog instruction."), and the [`ui.Dialog`](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.") class, to implement the dialog behavior.

To terminate the `INPUT` execution, the user can validate (or cancel) the dialog
to commit (or invalidate) the modifications made in the record.

When the statement completes execution, the form is deactivated.
After the user terminates the input (for example, with the "accept" key), the program must test the
`int_flag` variable to check if the
dialog was validated (or canceled), and then can use the `INSERT` or
`UPDATE` SQL statements to modify the appropriate database tables.

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
