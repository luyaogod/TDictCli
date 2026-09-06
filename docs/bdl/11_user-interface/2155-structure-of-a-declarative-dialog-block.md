---
title: "Structure of a declarative DIALOG block"
source: "fgl-topics/c_fgl_declarative_dialogs_structure.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Structure of a declarative DIALOG block"
type: "concept"
---

# Structure of a declarative DIALOG block

> A declarative DIALOG instruction is made of a single sub-dialog block, with an optional DEFINE clause to declare local variables.

Unlike procedural `DIALOG` blocks, declarative `DIALOG` blocks can
only define one sub-dialog block.

The dialog instruction in the declarative `DIALOG` block binds program variables
to form fields and defines the type of interaction that will take place for the data model (simple
input, list input or query).

The program variables used by the declarative dialog can be local to the current module or passed
as parameter in the `DIALOG` block definition.

The dialog implements individual [control
blocks](2099-dialog-control-blocks.md "Dialog control blocks are predefined dialog triggers where you can implement specific code to control the interactive instruction.") that allows you to control the behavior of the interactive instruction. The dialog can
also hold [action handlers](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.").

The declarative `DIALOG` block can define the following dialog
types:

- Simple record input with the [`INPUT` sub-dialog block](2083-the-input-sub-dialog.md "The INPUT sub-dialog implements single record input in fields of the current form.").
- Query by example input with the [`CONSTRUCT` sub-dialog block](2084-the-construct-sub-dialog.md "The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.").
- Read-only record list navigation with the [`DISPLAY ARRAY` sub-dialog block](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.").
- Editable record list handling with the [`INPUT ARRAY` sub-dialog block](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.").

## Related links

**Related concepts**  

[Structure of a procedural DIALOG block](2082-structure-of-a-procedural-dialog-block.md "Structure of a procedural DIALOG block")

## Child topics

- [The DEFINE clause](2156-the-define-clause.md): The DEFINE clause can be used to define program variables with a scope that is local to the declarative dialog block.
- [The INPUT sub-dialog](2157-the-input-sub-dialog.md): The INPUT sub-dialog implements single record input in fields of the current form.
- [The CONSTRUCT sub-dialog](2158-the-construct-sub-dialog.md): The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.
- [The DISPLAY ARRAY sub-dialog](2159-the-display-array-sub-dialog.md): The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.
- [The INPUT ARRAY sub-dialog](2160-the-input-array-sub-dialog.md): The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.
