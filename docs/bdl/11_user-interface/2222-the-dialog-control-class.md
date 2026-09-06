---
title: "The DIALOG control class"
source: "fgl-topics/c_fgl_prog_dialogs_class_dialog.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > The DIALOG control class"
type: "concept"
---

# The DIALOG control class

> This topic explains the purpose of the ui.DIALOG class.

Inside a dialog instruction, the `DIALOG` predefined keyword represents the
current dialog object. This dialog object can be used to execute methods provided by the
`ui.DIALOG` built-in class.

For example, you can enable or disable an action with the `setActionActive()`
dialog method, or you can hide or show the default action view with the
`setActionHidden()`
method:

```
BEFORE INPUT
  CALL DIALOG.setActionActive("zoom",FALSE)
AFTER FIELD field1
  CALL DIALOG.setActionHidden("zoom",TRUE)
```

The `setFieldActive()` method can be used to enable or disable a field
during the dialog:

```
ON CHANGE custname
   CALL DIALOG.setFieldActive( "custaddr",
               (rec.custname IS NOT NULL) )
```

The `ui.Dialog` class also provides methods to configure the dialog, for
example to enable multiple row selection:

```
BEFORE DIALOG
   CALL DIALOG.setSelectionMode( "sr1", 1 )
```

When using methods of the `ui.Dialog`
class that alter the data model or change the current field or row, control blocks like
`BEFORE FIELD`, `AFTER ROW`, `BEFORE DELETE` are not
executed: These are only fired to detect and control end user activity. Program code is considered
as part of the dialog implementation. For example, methods such as
`ui.Dialog.deleteRow()` must not execute `BEFORE DELETE` /
`AFTER DELETE` control blocks. These control blocks are only fired by an end-user
"delete" action.

## Related links

**Related concepts**  

[Dynamic Dialogs](2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
