---
title: "DIALOG ATTRIBUTES clause"
source: "fgl-topics/c_fgl_multiple_dialogs_DIALOG_ATTRIBUTES.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Procedural DIALOG block configuration > DIALOG ATTRIBUTES clause"
type: "concept"
description: "FIELD ORDER FORM option By default, the form tabbing order is defined by the variable list in the binding specification . You can control the tabbing order by using the FIELD ORDER FORM attribute; ..."
---

# DIALOG ATTRIBUTES clause

## FIELD ORDER FORM option

By
default, the form tabbing order is defined by the variable list in
the [binding specification](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.").
You can control the tabbing order by using the `FIELD ORDER
FORM` attribute; when this attribute is used, the tabbing
order is defined by the
[`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.")
attribute of the form items.

The field order mode can also be specified
globally with the [`OPTIONS
FIELD ORDER`](../09_advanced-features/0928-defining-field-tabbing-order-method.md) instruction.

With `FIELD
ORDER FORM`, if the user changes the focus from field A to
a distant field B with the mouse, the dialog does not execute
the `BEFORE FIELD` / `AFTER FIELD` triggers
of intermediate fields which appear in the binding specification
between field A and field B. Unlike singular dialogs, if the default
`FIELD ORDER CONSTRAINT` mode is used in a multiple dialog
instruction, intermediate triggers are never executed (i.e. the [`Dialog.fieldOrder`](2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior.") FGLPROFILE
entry is ignored by `DIALOG`.)

See also [Defining the tabbing order](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.").

## UNBUFFERED option

Indicates that the dialog must be sensitive to program variable changes. When using this
option, you bypass the traditional `BUFFERED` mode.

When using the traditional "buffered" mode, program variable changes are not automatically
displayed to form fields; You need to execute a `DISPLAY TO` or `DISPLAY BY
NAME`. Additionally, if an action is triggered, the value of the current field is not
validated and is not copied into the corresponding program variable. The only way to get the text of
the current field is to use [`GET_FLDBUF()`](../08_language-basics/0671-get-fldbuf-function.md "The GET_FLDBUF() operator returns as character strings the current values of the specified fields.").

If the "unbuffered" mode is used, program variables and form fields are automatically synchronized.
You don't need to display explicitly values with a `DISPLAY TO` or `DISPLAY
BY NAME`. When an action is triggered, the value of the current field is validated and is
copied into the corresponding program variable.

See also [The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.").

The unbuffered mode can be set globally for all `DIALOG` instructions with the
[`ui.Dialog.setDefaultUnbuffered()`](../15_library-reference/3177-ui-dialog-setdefaultunbuffered.md "Set the default unbuffered mode for all dialogs.")
class method:

```
CALL ui.Dialog.setDefaultUnbuffered(TRUE)
DIALOG -- Will work in UNBUFFERED mode     ...
END DIALOG
```
