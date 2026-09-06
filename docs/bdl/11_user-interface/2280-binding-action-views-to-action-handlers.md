---
title: "Binding action views to action handlers"
source: "fgl-topics/c_fgl_prog_dialogs_action_binding.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Binding action views to action handlers"
type: "concept"
---

# Binding action views to action handlers

> How are action views of the forms bound to action handlers in the program code?

Form [action views](2278-defining-action-views-in-forms.md "How to define action views that will fire action events.") (such as
buttons) are bound to [action handlers](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.") by
the `name` attribute. Action handlers are defined in interactive instructions with an
`ON ACTION` clause or `COMMAND` / `ON KEY` clauses.

For example, in the `ATTRIBUTES` section of the form, a button may be
defined as follows:

```
BUTTON b1: show_help, TEXT="Show Help";
```

The corresponding action handler (code) in the program will use the
"`show_help`" action
name:

```
ON ACTION show_help 
   CALL ShowHelp()
```

Other type of action views can for example be toolbar
items:

```
TOOLBAR tb
  ITEM show_help ( TEXT="Show Help" )
  ...
```

Or `BUTTONEDIT` buttons (using the `ACTION` attribute to define the
action name):

```
BUTTONEDIT f1 = customer.cust_city, ACTION = open_city_list;
```

The `COMMAND` / `ON KEY` clauses are typically used to
write text mode programs. Such clauses define the name of the action and the decoration
label. It is recommended that you use `ON ACTION` clauses instead,
because they identify user actions with an abstract name. However, if required, you can
use a `COMMAND` clause in a non-menu dialog to include the corresponding
action view in the focusable form items.

In the `ON ACTION action-name` clause, the name of the action must
be a valid identifier, preferably written in lowercase letters. In the abstract user
interface tree (where the action views are defined), action names are case sensitive (as
they are standard DOM attribute values). However, identifiers are not case sensitive in
the language. The fglcomp compiler always converts the action
identifiers of `ON ACTION` clauses to
lowercase:

```
ON ACTION PrintRecord   -- will be compiled as "printrecord"
```

To avoid confusion, always use lowercase names for action names (for example,
`print_record` instead of `PrintRecord`).

In the context of a multiple dialog, action views can be bound to action handlers by using a
simple action name, or by using a qualified action name with the sub-dialog name as
prefix:

```
BUTTON b1 = cust_list.append, TEXT="Add new customer";   -- cust_list is the sub-dialog id
```

With
a simple name, the action will be automatically enabled, if the action exists for the sub-dialog
having the focus, and the action is enabled. With a sub-dialog name as prefix, the action view will
be active even if the focus is not in the corresponding sub-dialog. For more details, read [Sub-dialog actions in procedural DIALOG blocks](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.")

## Related links

**Related concepts**  

[Field-specific actions (INFIELD clause)](2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.")
