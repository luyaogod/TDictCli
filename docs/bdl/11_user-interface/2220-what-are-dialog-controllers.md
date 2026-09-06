---
title: "What are dialog controllers?"
source: "fgl-topics/c_fgl_prog_dialogs_the_dialogs.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > What are dialog controllers?"
type: "concept"
---

# What are dialog controllers?

> Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.

The interactive instructions allow the program to respond to user actions and data
input.

## Simple display (non-interactive)

Instructions such as `DISPLAY BY NAME` display program variable data in the fields
of a form and continue the program flow without giving control to the end user. This is in fact not
an interactive instruction, as it just displays data to the current form, and returns immediately.
However, it may be used in interactive instructions to display information to the end user.

When using the `UNBUFFERED` mode of a dialog, you do not need to use the
`DISPLAY BY NAME` instruction to synchronize program variables and form fields.

The `MESSAGE` and `ERROR` instructions are also simple display
instructions without user interaction. These instructions are typically used to display a warning
message to the end user.

## The interactive dialog blocks

The singular [`MENU`](1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.") instruction handles a
list of choices to activate a specific function of the program. No field input is possible with this
instruction. The user can only select an action from the list.

The singular [`INPUT`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") instruction
is designed for simple record input. It enables the fields in a form for input, waits while the user
types data into the fields, and proceeds after the user accepts or cancels the dialog.

The singular [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
instruction is used to browse a list of records. It allows the user to view the contents of a
program array of records, scrolling the record list on the screen and choosing a specific record.
`DISPLAY ARRAY` implements by default a read-only list of records, but can be
extended to become a modifiable list with list modification triggers such as `ON
INSERT`.

The singular [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")
instruction supports record list input. It allows the user to alter the contents of records of a
program array, and to insert and delete records.

The singular [`CONSTRUCT`](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") instruction
is designed to let the user enter search criteria for a database query. The user can enter a value
or a range of values for one or several form fields, and your program looks up the database rows
that satisfy the requirements.

The [procedural `DIALOG`](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
instruction (placed in the program flow) allows you to combine several `INPUT`,
`DISPLAY ARRAY`, `INPUT ARRAY` and `CONSTRUCT`
functionality within the same form. This is also known as multiple-dialog block.

With singular dialog blocks such as `INPUT`, the automatic "cancel" action sets
[`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") to `TRUE`.
Note that `int_flag` is not reset to `FALSE` when the automatic
"accept" action is fired.

The [declarative
`DIALOG`](2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs.") block (defined in a module at the same level as a function)
implements individual `MENU`, `INPUT`, `DISPLAY ARRAY`,
`INPUT ARRAY` and `CONSTRUCT` functionality, that can be associated
with a procedural `DIALOG` instruction through the `SUBDIALOG` clause.
See [Using declarative dialogs](2154-using-declarative-dialogs.md "Dialog coding concepts, configuration and code structure.") for more details.

## Dynamic dialogs

When the form structure and field definitions can only be determined at runtime, it is possible
to implement [Dynamic Dialogs](2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class."), using generic code that will control a form
generated on the fly.

Dynamic dialogs can implement all types of dialogs like static dialogs do. However, this feature
should only be used in specific cases where regular static dialog instructions cannot be used.
Static dialog code is more readable than generic code implementing dynamic dialogs. A generic record
selection from a list (zoom) is the perfect candidate for a dynamic dialog.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
