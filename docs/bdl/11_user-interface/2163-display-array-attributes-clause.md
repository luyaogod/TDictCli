---
title: "DISPLAY ARRAY ATTRIBUTES clause"
source: "fgl-topics/c_fgl_DIALOG_block_DISPLAY_ARRAY_ATTRIBUTES_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Declarative DIALOG block configuration > DISPLAY ARRAY ATTRIBUTES clause"
type: "concept"
---

# DISPLAY ARRAY ATTRIBUTES clause

> DISPLAY ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.

## HELP option

The `HELP` attribute defines the number of the [help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to be displayed when invoked and
focus is in the list controlled by the `DISPLAY ARRAY` sub-dialog. The
predefined 'help' action is automatically created by the runtime system. You can bind
[action views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the 'help' action.

The `HELP` clause overrides the `HELP` attribute.

## COUNT option

The `COUNT` attribute defines the number of valid rows in the [static array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") to be displayed as default rows. If you
do not use the `COUNT` attribute, the runtime system cannot determine how
much data to display, so the screen array remains empty. The `COUNT` option
is ignored when using a dynamic array, unless page mode is used. In this case, the
`COUNT` attribute must be used to define the total number of rows, because
the dynamic array will only hold a page of the entire row set. If the value of
`COUNT` is negative or zero, it defines an empty list.

See also [Controlling the number of rows](2302-controlling-the-number-of-rows.md "Methods are provided to set and get the total number of rows in a read-only or editable list of records.").

## DOUBLECLICK option

The `DOUBLECLICK` option can be used to define the action that will
be fired when the user chooses a row from the list. Different configuration options are
available to control the row selection action of desktop and mobile devices. For more
details, see [Defining the action for a row choice](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.").

## FOCUSONFIELD option

When the `FOCUSONFIELD` option is used, the `DISPLAY ARRAY` allows
focus at the field (or cell) level. It is then possible to implement `BEFORE FIELD`
and `AFTER FIELD` blocks, as well as using `NEXT FIELD` instructions.
However, the dialog still manages a read-only list. For more details, see [Field-level focus in DISPLAY ARRAY](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.").
