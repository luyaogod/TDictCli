---
title: "The Dialog class"
source: "fgl-topics/c_fgl_ClassDialog.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class"
type: "concept"
---

# The Dialog class

> The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.

A `ui.Dialog` object can for example be used to enable or disable actions and form
fields dynamically during the dialog execution.

A dialog object is typically available inside a dialog block, with the predefined
`DIALOG` keyword, and can only be referenced during the execution of that
interactive instruction. After the interactive instruction, the dialog object is
destroyed and its reference becomes invalid.

When using methods of the `ui.Dialog`
class that alter the data model or change the current field or row, control blocks like
`BEFORE FIELD`, `AFTER ROW`, `BEFORE DELETE` are not
executed: These are only fired to detect and control end user activity. Program code is considered
as part of the dialog implementation. For example, methods such as
`ui.Dialog.deleteRow()` must not execute `BEFORE DELETE` /
`AFTER DELETE` control blocks. These control blocks are only fired by an end-user
"delete" action.

Dialog objects can also be created dynamically to handle forms created at runtime. This
feature is only provided for specific needs.

## Related links

**Related concepts**  

[Dialog instructions](../11_user-interface/1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation.")

[The DIALOG control class](../11_user-interface/2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")

## Child topics

- [ui.Dialog methods](3170-ui-dialog-methods.md): Methods of the ui.Dialog class.
- [Usage](3235-usage.md)
- [Examples](3242-examples.md): ui.Dialog usage examples.
