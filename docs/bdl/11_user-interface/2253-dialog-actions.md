---
title: "Dialog actions"
source: "fgl-topics/c_fgl_prog_dialogs_actions.html"
breadcrumb: "User interface > User interface programming > Dialog actions"
type: "concept"
---

# Dialog actions

> Describes how to program action handling when the end user triggers an action on the front-end.


## Child topics

- [Action handling basics](2254-action-handling-basics.md): This topic describes the basic concepts of dialog actions.
- [Predefined actions](2255-predefined-actions.md): Genero predefines some action names for common operations of interactive instructions.
- [Default action views](2258-default-action-views.md): A default action view is created to render an action handler, when no explicit action view exists for it.
- [Configuring actions](2259-configuring-actions.md): Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.
- [Defining action views in forms](2278-defining-action-views-in-forms.md): How to define action views that will fire action events.
- [Implementing dialog action handlers](2279-implementing-dialog-action-handlers.md): How to execute user code in ON ACTION blocks when an action is fired.
- [Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md): How are action views of the forms bound to action handlers in the program code?
- [Data validation at action invocation](2281-data-validation-at-action-invocation.md): The validate action attribute controls field validation when an action is fired.
- [Enabling and disabling actions](2282-enabling-and-disabling-actions.md): By default, dialog actions are enabled. However, it is recommended that an action be disabled when not allowed in the current context.
- [Hiding and showing default action views](2283-hiding-and-showing-default-action-views.md): If needed, default action views can be hidden or shown.
- [Sub-dialog actions in procedural DIALOG blocks](2284-sub-dialog-actions-in-procedural-dialog-blocks.md): This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.
- [Field-specific actions (INFIELD clause)](2285-field-specific-actions-infield-clause.md): Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.
- [Multilevel action conflicts](2286-multilevel-action-conflicts.md)
- [Action views in chromebar](2287-action-views-in-chromebar.md): Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.
- [Action display in the context menu](2288-action-display-in-the-context-menu.md): The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.
- [Implementing the close action](2289-implementing-the-close-action.md): The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button).
- [Implementing the back action](2290-implementing-the-back-action.md): The back action is a predefined action dedicated to move back in the stack of windows/forms.
- [Actions bound to the current row](2291-actions-bound-to-the-current-row.md): Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.
- [Keyboard accelerator names](2292-keyboard-accelerator-names.md): Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.
- [Setting action key labels](2296-setting-action-key-labels.md): Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.
- [Action views on mobile devices](2297-action-views-on-mobile-devices.md): Action views are rendered following mobile specific standards.
- [Automatic action views](2298-automatic-action-views.md): Action views can be rendered automatically in some form elements.
