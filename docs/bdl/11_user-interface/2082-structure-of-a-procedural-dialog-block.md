---
title: "Structure of a procedural DIALOG block"
source: "fgl-topics/c_fgl_multiple_dialogs_structure.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Structure of a procedural DIALOG block"
type: "concept"
description: "A procedural DIALOG instruction is made up of several sub-dialogs , plus global control blocks such as BEFORE DIALOG and action handlers such as ON ACTION or COMMAND . Sub-dialogs can be defined ..."
---

# Structure of a procedural DIALOG block

A procedural `DIALOG` instruction is made up of several sub-dialogs,
plus global control blocks such as `BEFORE DIALOG` and action handlers such as
`ON ACTION` or `COMMAND`.

Sub-dialogs can be defined inside the `DIALOG` instruction, or can be declared
externally in another module and attached to the current `DIALOG` block with the
`SUBDIALOG` clause. A dialog defined in the scope of a function is know as a
procedural dialog block, while a dialog declared in the scope of a module is named
a declarative dialog block.

The sub-dialogs bind program variables to form fields and define
the type of interaction that will take place for the data model
(simple input, list input or query). The sub-dialogs implement
individual [control blocks](2099-dialog-control-blocks.md "Dialog control blocks are predefined dialog triggers where you can implement specific code to control the interactive instruction.") which
let you control the behavior of the interactive instruction. Sub-dialogs
can also hold action handlers, which will define local [sub-dialog actions](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?").

The `DIALOG` procedural instruction can hold the following type of
sub-dialogs:

1. Simple record input with the `INPUT` sub-dialog
   block.
2. Query by example input with the `CONSTRUCT` sub-dialog
   block.
3. Read-only record list navigation with the `DISPLAY ARRAY` sub-dialog
   block.
4. Editable record list handling with the `INPUT ARRAY` sub-dialog
   block.
5. A `SUBDIALOG` clause referencing a declarative
   sub-dialog by name.

## Related links

**Related concepts**  

[Declarative dialogs (DIALOG - at module level)](2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs.")

## Child topics

- [The INPUT sub-dialog](2083-the-input-sub-dialog.md): The INPUT sub-dialog implements single record input in fields of the current form.
- [The CONSTRUCT sub-dialog](2084-the-construct-sub-dialog.md): The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.
- [The DISPLAY ARRAY sub-dialog](2085-the-display-array-sub-dialog.md): The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.
- [The INPUT ARRAY sub-dialog](2086-the-input-array-sub-dialog.md): The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.
- [The SUBDIALOG clause](2087-the-subdialog-clause.md)
