---
title: "Identifying actions in ui.Dialog methods"
source: "fgl-topics/c_fgl_ClassDialog_identify_action.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Usage > Identifying actions in ui.Dialog methods"
type: "concept"
description: "In ui.Dialog methods such as setActionActive() , the first parameter identifies the action object to be modified. This parameter can be fully-qualified or partly-qualified. If you don't specify a ..."
---

# Identifying actions in ui.Dialog methods

In `ui.Dialog` methods such as [`setActionActive()`](3213-ui-dialog-setactionactive.md "Enabling and disabling dialog actions."), the first parameter identifies the action object to be
modified. This parameter can be fully-qualified or partly-qualified. If you don't specify a
fully-qualified name, the action object will be identified based on the focus context.

The action name specification can be any of the following:

- action-name
- dialog-name.action-name
- dialog-name.field-name.action-name
- field-name.action-name (singular dialogs only)

Here action-name identifies the name of the action specified in
`ON ACTION action-name` or `COMMAND "action-name"`
handlers, while dialog-name identifies the [singular dialog](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") or [sub-dialog](../11_user-interface/2082-structure-of-a-procedural-dialog-block.md) and field-name
defines the field bound to the action [`INFIELD`](../11_user-interface/2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.") clause
of `ON ACTION`.

The runtime system will raise the error [**-8089**](4483-genero-bdl-errors.md) if the action
specified by [dialog-name.][field-name.]action-name can not be found
within the current dialog.

As a general rule, assign unique action names for each specific dialog action, to avoid the usage
of dialog and/or field identifiers.

The name of the action passed as parameter can use the same letter case as the action definition:
The lookup is case-insensitive:

```
ON ACTION PrintReport
...
CALL DIALOG.setActionActive("PrintReport", FALSE)
```

In the [`DIALOG`](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
instruction, actions can be prefixed with the [sub-dialog identifier](../11_user-interface/2081-identifying-sub-dialogs-in-dialog.md "Sub-dialogs need to be identified by a name to distinguish the different contexts."). However, if methods like `setActionActive()` are
called in the context of the sub-dialog, the prefix can be omitted. When using a field-specific
action defined with the `INFIELD` clause of `ON ACTION`, you can
identify the action with the fully-qualified name
dialog-name.field-name.action-name. Like sub-dialog actions, if you specify only
action-name, the runtime system will search for the action object based on the
focus context.

Note that an [`INPUT`](../11_user-interface/2083-the-input-sub-dialog.md "The INPUT sub-dialog implements single record input in fields of the current form.") or
[`CONSTRUCT`](../11_user-interface/2084-the-construct-sub-dialog.md "The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.") sub-dialogs
have no identifier by default. The dialog name can be defined with the `NAME`
attribute. For more details, see [Identifying sub-dialogs in DIALOG](../11_user-interface/2081-identifying-sub-dialogs-in-dialog.md "Sub-dialogs need to be identified by a name to distinguish the different contexts.").

When using a singular dialog like [`INPUT`](../11_user-interface/1932-syntax-of-the-input-instruction.md "The INPUT statement supports data entry in fields of the current form."), you can identify field-specific
actions by field-name.action-name, if the dialog was defined without a
`NAME` attribute.

## Example

```
MAIN
  DEFINE cust_rec RECORD
            num INT,
            name VARCHAR(50),
            city INT
         END RECORD
  DEFINE orders DYNAMIC ARRAY OF RECORD
            ord_num INT,
            cust_num INT,
            ord_date DATE
         END RECORD
  OPEN FORM f1 FROM "cust_ord"
  DISPLAY FORM f1
  DIALOG ATTRIBUTES(UNBUFFERED)
   INPUT BY NAME cust_rec.* ATTRIBUTES(NAME="cust")
      ON ACTION compare
         CALL compare()
      ON ACTION check INFIELD cust_city
         CALL check_city(cust_rec.city)
   END INPUT
   DISPLAY ARRAY orders TO sr_ord.*
      ON ACTION archive
         CALL archive()
   END DISPLAY
   ON ACTION print
      CALL print()
   ON ACTION disable_all
      CALL DIALOG.setActionActive("cust.compare", FALSE)
      CALL DIALOG.setActionActive("cust.cust_city.check", FALSE)
      CALL DIALOG.setActionActive("sr_ord.archive", FALSE)
      CALL DIALOG.setActionActive("print", FALSE)
  END DIALOG
END MAIN
```

## Related links

**Related concepts**  

[Enabling and disabling actions](../11_user-interface/2282-enabling-and-disabling-actions.md "By default, dialog actions are enabled. However, it is recommended that an action be disabled when not allowed in the current context.")

[Binding action views to action handlers](../11_user-interface/2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")
