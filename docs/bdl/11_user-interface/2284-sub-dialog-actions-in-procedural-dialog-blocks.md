---
title: "Sub-dialog actions in procedural DIALOG blocks"
source: "fgl-topics/c_fgl_prog_dialogs_subdialog_action.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Sub-dialog actions in procedural DIALOG blocks"
type: "concept"
---

# Sub-dialog actions in procedural DIALOG blocks

> This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.

We distinguish dialog actions from sub-dialog actions. When the
`ON ACTION` handler is defined at the same level as a `BEFORE DIALOG`
control block, it is a dialog action, and the action name is a simple identifier as in singular
interactive instructions:

```
action-name
```

When the `ON ACTION` handler is defined inside a sub-dialog, or if the action
is an automatic action such as `insert` in `INPUT ARRAY`, it is a
sub-dialog action. A sub-dialog action is qualified with the name of the
sub-dialog:

```
sub-dialog-name.action-name
```

The `INPUT ARRAY` and `DISPLAY ARRAY` sub-dialogs are
implicitly identified with the screen-record name defined in the form. For
`INPUT` and `CONSTRUCT` sub-dialogs, the sub-dialog identifier
can be specified with the `NAME` attribute.

This example defines two 'check' actions in different sub-dialog contexts, and a 'close' action
at the dialog level:

```
DIALOG
  INPUT BY NAME ... ATTRIBUTES (NAME = "cust")
    ON ACTION check                   -- sub-dialog action "cust.check"
    ...
  END INPUT
  DISPLAY ARRAY arr_orders TO sr_ord.*
    ...
    ON ACTION check                   -- sub-dialog action "sr_ord.check"
    ...
  END DISPLAY
  BEFORE DIALOG
    ...
  ON ACTION close                      -- dialog action "close"
    ...
END DIALOG
```

By using the sub-dialog identifier in form definition files, you can bind action views to
specific sub-dialog actions. Action views bound to sub-dialog actions with qualified sub-dialog
action names will always be active, even if the focus is not in the sub-dialog of the action. You
typically use fully-qualified sub-dialog action names for buttons in the form body or in topmenu
options. However, it does not make much sense to use this technique for toolbar buttons, where
buttons must be enabled/disabled based on the
context.

```
TOOLBAR
   ...
   ITEM append
   ...
END

TOPMENU
   ...
   GROUP orders (TEXT="Orders")
     COMMAND sr_ord.append
   ...
END

LAYOUT
GRID
{
   ...
   [b002  ]
}
END
END

ATTRIBUTES
BUTTON b002: sr_ord.append;
END
```

If you bind an action view with a simple action name (without the sub-dialog prefix), the action
view will be attached to any sub-dialog action with the matching name. This is especially useful for
common actions such as the insert / append / delete actions created by `INPUT ARRAY`,
when the dialog handles multiple editable lists. Bind toolbar buttons to these actions without the
sub-dialog prefix; the buttons will apply to the current list that has the focus. The action views
bound to sub-dialog actions without the sub-dialog qualifier will automatically be enabled or
disabled when entering or leaving the group of fields controlled by the sub-dialog (typical
navigation buttons in the toolbar will be disabled if the focus is not in a list).

If a sub-dialog action is invoked when the focus is not in the
sub-dialog of the action, the focus will automatically be given to
the first field of the sub-dialog, before executing the user code
defined in the `ON ACTION` clause. This will trigger
the same validation rules and control blocks as if the user had selected
the first field of the sub-dialog by hand.

When using `DIALOG.setActionActive()` (or any method that takes an action name as parameter),
you can specify the action name with or without a sub-dialog identifier. If you qualify the action
with the sub-dialog identifier, the sub-dialog action is clearly identified. If you don't specify a
sub-dialog prefix, the action will be identified based on the focus context - when the focus is in
the sub-dialog of the action, non-qualified action names identify the local sub-dialog action;
otherwise, they identify a dialog action if one exists with the same name. Disabling an action by
the program with `setActionActive()`, will take precedence over the built-in
activation rules (this means that if the action is disabled by the program, the action will not be
activated when entering the sub-dialog).

For action views bound to sub-dialog actions with qualifiers, the
action defaults defined with the corresponding action name will
be used to set the attributes with the default values. In other
words, the prefix will be ignored. For example, if an action view
is defined with the name "`custlist.append`",
it will get the action defaults defined for the "`append`"
action.

## Related links

**Related concepts**  

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Topmenus](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")

[Toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")
