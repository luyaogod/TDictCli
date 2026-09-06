---
title: "Hiding and showing default action views"
source: "fgl-topics/c_fgl_prog_dialogs_action_visibility.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Hiding and showing default action views"
type: "concept"
---

# Hiding and showing default action views

> If needed, default action views can be hidden or shown.

When an action is rendered with a default action view (for example, by a button on the action
frame of a desktop front-end, or in the top action panel on a mobile front-end), it is sometimes
required to hide the action button when the operation is not possible and there is not much space on
the screen.

> **Important:**
>
> Hiding an action will only make the default action view invisible, if
> there is a keyboard accelerator associated to the action, it can still fire the action.
> Consider disabling the action completely with `setActionActive()`.

During a dialog instruction, show or hide an action with the `setActionHidden()` method of the
`ui.Dialog` built-in class. This method takes the name of the action (in lowercase
letters) and a boolean expression (`FALSE` or `TRUE`) as
arguments:

```
  BEFORE INPUT
    CALL DIALOG.setActionHidden( "zoom", TRUE )
```

Consider centralizing action visibility control in a setup function specific to the dialog,
passing the `DIALOG` object as the parameter. Centralizing the action
activation defines the rules in a single
location:

```
FUNCTION cust_dialog_setup(d)
  DEFINE d ui.Dialog
  DEFINE can_modify BOOLEAN
  LET can_modify = (cust_rec.is_new OR user_info.is_admin)
  CALL d.setActionActive("update", can_modify)
  CALL d.setActionHidden("update", IIF(can_modify,0,1))
  CALL d.setActionActive("delete", can_modify)
  CALL d.setActionHidden("delete", IIF(can_modify,0,1))
  ...
END FUNCTION
```

Pay attention to multi-level action definitions inside a `DIALOG` block: Inside a
`DIALOG` block, actions must be hidden/shown with the
`ui.Dialog.setActionHidden()` method by specifying a simple action
name:

```
DIALOG ATTRIBUTES(UNBUFFERED)
     ...
   BEFORE DIALOG
     CALL DIALOG.setActionHidden( "print", TRUE )
     ...
   ON ACTION query 
     -- query the database and fill the record
     ...
     CALL DIALOG.setActionHidden( "print", (cust_id IS NULL) )
     ...
END DIALOG
```
