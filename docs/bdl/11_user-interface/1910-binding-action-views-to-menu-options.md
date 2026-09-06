---
title: "Binding action views to menu options"
source: "fgl-topics/c_fgl_menus_021.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > Binding action views to menu options"
type: "concept"
description: "A MENU statement is a controller for user actions, defining action handlers triggered by action views. Actions views in the form file (such as TOOLBAR buttons, TOPMENU items or push BUTTON ) are bound ..."
---

# Binding action views to menu options

A `MENU` statement is a controller for user actions, defining action handlers
triggered by action views. Actions views in the form file (such as `TOOLBAR` buttons, `TOPMENU` items or push `BUTTON`) are bound to menu options by name.
For example, if a `MENU` instruction defines `ON ACTION sendmail`, a
form button with the name "sendmail" will be attached to that action handler.

When binding action views to menu option clauses, the action name is case sensitive. The
compiler converts `COMMAND` labels and `ON ACTION` identifiers to lowercase to create the
action name. It is recommended that you use all lowercase letters when defining the action name for
action views and menu options.

Menu options can also be defined with the [`COMMAND`](1916-command-key-option-block.md) clause. Unlike `ON ACTION`, the
`COMMAND` clause takes a string literal as argument, that defines both the action
name and the default text to be displayed in the default action view. For example, `COMMAND
"Help"` will define the action name help and the default button text "Help". Action views
must be bound with the action name in lowercase (`help`).

When the menu is rendered as a pop-up of dialog box, no explicit action views need to be
defined, default action views will be created and will get the decoration specified in action
defaults.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
