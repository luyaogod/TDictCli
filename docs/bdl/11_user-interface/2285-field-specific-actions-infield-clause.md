---
title: "Field-specific actions (INFIELD clause)"
source: "fgl-topics/c_fgl_prog_dialogs_infield_action.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Field-specific actions (INFIELD clause)"
type: "concept"
---

# Field-specific actions (INFIELD clause)

> Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.

The `ON ACTION` interaction block of `INPUT`,
`CONSTRUCT` and `INPUT ARRAY` (as singular dialogs or sub-dialogs in
`DIALOG` instruction), can be specified with the `INFIELD
field-name` clause. With this clause, the action will only be active when the focus
is on one of the fields. The same action name can be used for several fields.

```
INPUT ARRAY custarr WITHOUT DEFAULTS FROM sr_cust.*
  ON ACTION zoom INFIELD cust_city
    LET custarr[arr_curr()].cust_city = zoom_city() 
  ON ACTION zoom INFIELD cust_state
    LET custarr[arr_curr()].cust_state = zoom_state() 
END INPUT
```

Actions defined with the `INFIELD field-name` clause can
be identified with the field name as
prefix:

```
field-name.action-name
```

You can bind form action views with or without the field name prefix:

```
-- Form file:
EDIT f1 = customer.cust_name;
BUTTON b1 : cust_name.clear, ... ;    -- field-qualified action
BUTTON b2 : clear, ... ;              -- unqualified action

-- Program code:
ON ACTION clear INFIELD cust_name
   ...
```

- Without the field name prefix, the action view is enabled and disabled automatically depending
  on the current field: The action view is enabled when the corresponding `INFIELD`
  field has the focus.
- When binding the action view with the fully-qualified name including the field name prefix, the
  action view will always be active, and the focus will jump to the first corresponding field, if the
  action is fired.

When using `ON ACTION action-name INFIELD
field-name` for one or several [`BUTTONEDIT`](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.") fields, the runtime
system implicitly handles the `BUTTONEDIT` action as a field-qualified action, even
if the `ACTION` attribute is defined without the field
name:

```
BUTTONEDIT f1 = customer.cust_city, ACTION = zoom;
```

Is equivalent
to:

```
BUTTONEDIT f1 = customer.cust_city, ACTION = cust_city.zoom;
```

Actions defined in sub-dialogs of the `DIALOG` instruction get the name of the
sub-dialog as prefix. If `ON ACTION action-name INFIELD
field-name` is used in a sub-dialog, the action object name is prefixed with
the name of the sub-dialog, followed by the name of the field. The fully-qualified action name
will
be:

```
sub-dialog-name.field-name.action-name
```

When the field-specific action is invoked (for example by a button
of the toolbar bound with the fully-qualified action name) and if
the field does not have the focus, the runtime system first selects
that field before executing the code of the `ON ACTION INFIELD` block.
The field selection forces data validation and `AFTER FIELD` of
the current field, followed by `BEFORE FIELD` of the
target field associated to the action.

It's still possible to enable and disable field-specific action
objects by the program using the `DIALOG.setActionActive()` method.
When specifying a fully-qualified action name with the field
name prefix, that field-specific action will be enabled or disabled.
When disabled by the `setActionActive()` method,
the corresponding action views will always be disabled, even
if the field has the focus. If you do not specify a fully-qualified
name in the method call, and if several actions are defined
with the same action name in different sub-dialogs and/or using
the `INFIELD` clause, the method will identify
the action according to the current focus context. For example, if
you define `ON ACTION zoom INFIELD cust_city` and `ON
ACTION zoom INFIELD cust_addr`, when the focus is
in `cust_city`, a call to `DIALOG.setActionActive("zoom",
FALSE)` will disable the action specific to the `cust_city` field.

Fields can be enabled or disabled dynamically with the `DIALOG.setFieldActive()`
method. If an `ON ACTION INFIELD` is declared
on a field and if you enable/disable the field dynamically,
then the field-specific action (and corresponding action views
in the form) will be enabled or disabled accordingly.

For action views bound to field actions with qualifiers, the action
defaults defined with the corresponding action name will
be used to set the attributes with the default values. In other
words, the prefix will be ignored. For example, if an action view
is defined with the name "`cust_addr.check`",
it will get the action defaults defined for the "`check`"
action.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Topmenus](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")

[Toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")
