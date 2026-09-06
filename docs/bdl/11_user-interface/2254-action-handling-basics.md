---
title: "Action handling basics"
source: "fgl-topics/c_fgl_prog_dialogs_actions_basics.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Action handling basics"
type: "concept"
---

# Action handling basics

> This topic describes the basic concepts of dialog actions.

A dialog action is a general term, used to identify an application function that can
be triggered by the end user. For example, the "`accept`" action validates a dialog
input, the "`close`" action closes the current window.

We distinguish user-defined actions from predefined actions:

- User-defined actions are created with an `ON ACTION`,
  `COMMAND` or `ON KEY` block, in a dialog instruction.
- Predefined actions are actions with a reserved name such as
  `accept`, `cancel`, `interrupt`. There are several
  types of predefined actions: Dialogs create automatic actions like
  `accept` and `cancel`. Applications can also use special
  actions and particular actions such as `interrupt`, to let the
  user cancel a running application procedure. See [Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.") for more details.

In the user interface of the application, action views can produce action
events, that will execute user code in the corresponding action handler, defined
in the current interactive instruction of the program.

We distinguish explicit action views from implicit action views and
default action views:

- Explicit actions views are for example `BUTTON` form items,
  `TOOLBAR` items or `TOPMENU` commands. For more details, see [Defining action views in forms](2278-defining-action-views-in-forms.md "How to define action views that will fire action events.").
- Implicit action views action views include the concept of [default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action views](2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.").
- Default action views are created, when no explicit action views are
  defined for that action in the current form. A default action view is typically a
  button that appears in a specific area, located and decorated following the front-end platform
  standards. For more details, see [Default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.").

The action views are bound to action handlers by the name of the
action. For more details, see [Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?").

An action event is an AUI event produced by a click on an action view,
or by using an action [accelerator key](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.").

An action handler is an `ON ACTION action-name` or
`COMMAND "action-name"` dialog block, implementing the user code
in the current interactive dialog. See [Implementing dialog action handlers](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.") for
more details.

Actions can belong to a specific context:

- dialog-level actions: The `ON ACTION` block is defined in a [singular dialog](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling."), or is defined in a [`DIALOG`](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.") block, at the same level
  as subdialogs.
- subdialog-level actions: The `ON ACTION` block is defined inside a
  [subdialog block](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.") of a
  `DIALOG` block.
- row-level actions: The `ON ACTION` block is defined with [`ROWBOUND`](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.") attribute.
- field-level actions: The `ON ACTION` block is defined with [`INFIELD`](2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.") clause.

Actions can be configured with action attributes. These attributes can be defined
explicitly at the action view level (button in form) or with action
defaults. Dialog-level action configuration is possible with `ON ACTION
name ATTRIBUTES(...)`, to define functional attributes and decorate
default action views. For more details, see [Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.").
