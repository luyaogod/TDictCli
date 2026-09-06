---
title: "Enabling and disabling actions"
source: "fgl-topics/c_fgl_prog_dialogs_action_activation.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Enabling and disabling actions"
type: "concept"
---

# Enabling and disabling actions

> By default, dialog actions are enabled. However, it is recommended that an action be disabled when not allowed in the current context.

Dialog actions are enabled to let the user invoke the action handler (`ON
ACTION`/`COMMAND`) by clicking on the corresponding action
view (button) or by pressing its accelerator key. In most situations, actions remain
active during the whole dialog execution. However, to follow GUI standards, actions must
be disabled when not allowed in the current context. For example, a print action is
disabled if no record is currently shown in the form. After a database query, when the
form is filled with a given record, the print action can be activated.

During a dialog instruction, enable or disable an action with the
`setActionActive()` method of the `ui.Dialog`
built-in class. This method takes the name of the action (in lowercase
letters) and a boolean expression (`0` or `FALSE`,
`1` or `TRUE`) as arguments.

```
  BEFORE INPUT
    CALL DIALOG.setActionActive( "zoom", FALSE )
```

Consider centralizing action activation / deactivation in a setup
function specific to the dialog, passing the `DIALOG` object
as the parameter. Centralizing the action activation defines the rules in
a single location:

```
FUNCTION cust_dialog_setup(d)
  DEFINE d ui.Dialog
  DEFINE can_modify BOOLEAN
  LET can_modify = (cust_rec.is_new OR user_info.is_admin)
  CALL d.setActionActive("update", can_modify)
  CALL d.setActionActive("delete", can_modify)
  ...
END FUNCTION
```

Some [predefined dialog actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.") such as
insert / append / delete of `INPUT ARRAY` are automatically enabled/disabled based on
the context. For example, if the maximum number of rows (`MAXCOUNT`) is reached in an
`INPUT ARRAY`, insert and append actions are disabled.

When the action activation depends on the focus being in a specific field, consider using the
[`INFIELD`](2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.")
clause of `ON ACTION` to automatically disable an action if the focus
leaves the specified field.

Inside a `DIALOG` block, actions can be defined at different levels, and may need
to be identified with the sub-dialog prefix, when you invoke the
`ui.Dialog.setActionActive()` method outside of the context of the sub-dialog. In the
following example, the `check_row` action must be prefixed by the
`s_ord` sub-dialog name, because `setActionActive()` is called from
the `INPUT BY NAME` sub-dialog context, to disable an action from the `DISPLAY
ARRAY`
sub-dialog:

```
DIALOG ATTRIBUTES(UNBUFFERED)
   DISPLAY ARRAY a_ord TO s_ord.*
      -- sub-dialog-level action
      ON ACTION check_row
         ...
   END DISPLAY
   ...
   INPUT BY NAME rec.* ...
      ON CHANGE consolidation
          -- Must use sub-dialog name to identify the check_row action:
          CALL DIALOG.setActionActive( "s_ord.check_row", FALSE )
      ...
   END INPUT
END DIALOG
```

## Related links

**Related concepts**  

[Identifying actions in ui.Dialog methods](../15_library-reference/3238-identifying-actions-in-ui-dialog-methods.md "Identifying actions in ui.Dialog methods")
