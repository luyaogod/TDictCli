---
title: "INPUT ATTRIBUTES clause"
source: "fgl-topics/c_fgl_DIALOG_block_INPUT_ATTRIBUTES.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Procedural DIALOG block configuration > INPUT ATTRIBUTES clause"
type: "concept"
---

# INPUT ATTRIBUTES clause

> INPUT specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.

## HELP option

The `HELP` attribute defines the number of the
[help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to be
displayed when invoked and focus is in the list controlled by the
`INPUT` sub-dialog. The predefined 'help' action
is automatically created by the runtime system.
You can bind [action
views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the 'help' action.
The `HELP` clause overrides the `HELP`
attribute.

## NAME option

The `NAME` attribute can be used to identify the `INPUT`
sub-dialog, especially useful to qualify [sub-dialog actions](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.")

## WITHOUT DEFAULTS option

The `WITHOUT DEFAULTS` option indicates if the
fields controlled by `INPUT` must be filled (`FALSE`) or not
(`TRUE`) with the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") values defined in the form specification file. The runtime system
assumes that the field satisfies the [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") attribute, when `WITHOUT DEFAULTS` is used. If the
`WITHOUT DEFAULTS` option is not used, all fields defined with the
`REQUIRED` attribute must be visited and modified.

For more details see [Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.").
