---
title: "Multilevel action conflicts"
source: "fgl-topics/c_fgl_prog_dialogs_multilevel_actions.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Multilevel action conflicts"
type: "concept"
description: "Actions can be defined at two levels in a singular dialog, and three levels in the context of a DIALOG block: Dialog level Sub-dialog level (procedural DIALOG only) Field level ( ON ACTION with ..."
---

# Multilevel action conflicts

Actions can be defined at two levels in a singular dialog, and three
levels in the context of a `DIALOG` block:

1. Dialog level
2. Sub-dialog level (procedural `DIALOG` only)
3. Field level (`ON ACTION` with `INFIELD` clause)

It is not good practice to use the same action name at different levels of a dialog: This makes
action view bindings and action handling (enabling / disabling) very complex, because there are many
possible combinations. Therefore, when using the same action name at different dialog levels, the
fglcomp compiler will raise a warning [-8409](../15_library-reference/4483-genero-bdl-errors.md). However, it is legal to use
the same action name for a given level of action handlers in sub-dialogs or for field-actions. For
example, using the "zoom" action name for multiple `ON ACTION INFIELD` handlers is a
common practice.

When binding action views with full qualified names, the `ON
ACTION` handler is clearly identified, and the corresponding
user code will be executed. However, when you do not specify the complete
prefix of a sub-dialog or field action, the runtime system searches
for the best `ON ACTION` handler to be executed, according
to the current focus context.

Take for example a `DIALOG` instruction defining
three `ON ACTION print` handlers at the dialog,
sub-dialog and field level:

```
DIALOG
  INPUT BY NAME ... ATTRIBUTES (NAME = "cust")
    ...
    ON ACTION print INFIELD cust_name -- field-level action (1)
      ...
    ON ACTION print                   -- sub-dialog-level action (2)
      ...
  END INPUT
  ...
  ON ACTION print                     -- dialog-level action (3)
    ...
END DIALOG
```

The action views of the form will behave as follows:

- Action views bound with the name "`print`" will
  always be active, and invoke the `ON ACTION print` handler
  corresponding to the current focus context:
  - (1) is invoked if the focus is in the `cust_name` field.
  - (2) is invoked if the focus is in the `cust` sub-dialog,
    but not in `cust_name` field.
  - (3) is invoked if the focus is in another sub-dialog as `cust`
    sub-dialog.
- Action views bound with the name "`cust.print`" will always be
  active, even if the focus is not the `cust` sub-dialog, and invoke
  the `ON ACTION print` handler depending on the focus context:
  - (1) is invoked if the focus is in the `cust_name` field.
  - (2) is invoked if the focus is in the `cust` sub-dialog,
    but not in `cust_name` field.
- Action views bound with the name "`cust.cust_name.print`"
  will always be active, and invoke the `ON ACTION print
  INFIELD cust_name` handler after giving the focus
  to the `cust_name` field.

If the first field of a sub-dialog defines an `ON ACTION
INFIELD` with the same action name as a sub-dialog action,
and the focus is not in that sub-dialog when the user selects an action
view bound with the name *sub-dialog-name.action-name*, the runtime
system gives the focus to the first field of the sub-dialog. This
field becomes the current field, and the runtime system executes the
field-specific action handler instead of the sub-dialog action handler.

To avoid mistakes and complex combinations, it is recommended that you use specific
action names for each dialog level.
