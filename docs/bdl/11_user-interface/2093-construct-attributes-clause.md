---
title: "CONSTRUCT ATTRIBUTES clause"
source: "fgl-topics/c_fgl_DIALOG_block_CONSTRUCT_ATTRIBUTES.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Procedural DIALOG block configuration > CONSTRUCT ATTRIBUTES clause"
type: "concept"
---

# CONSTRUCT ATTRIBUTES clause

> CONSTRUCT specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.

## HELP option

The `HELP` attribute defines the number of the
[help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to
be displayed when invoked and focus is in the list controlled by
the `CONSTRUCT` sub-dialog.
The predefined 'help' action is automatically created by the runtime
system. You can bind [action
views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the 'help' action.

The `HELP` clause overrides the `HELP` attribute.

## NAME option

The `NAME` attribute can be used to identify the `CONSTRUCT`
sub-dialog; this is especially useful to qualify [sub-dialog actions](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.").
