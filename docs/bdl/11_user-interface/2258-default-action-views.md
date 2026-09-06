---
title: "Default action views"
source: "fgl-topics/c_fgl_prog_dialogs_default_action_views.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Default action views"
type: "concept"
---

# Default action views

> A default action view is created to render an action handler, when no explicit action view exists for it.

If no explicit action view is defined, such as a toolbar button, a topmenu item or a simple
button in the form layout, the front-end creates a default action view for each
`COMMAND` or `ON ACTION` action handler, or automatic actions such as
insert/delete in `INPUT ARRAY`, in the current interactive instruction.

The rendering of default action views depends on the platform:

- On a desktop front-end, the default action views appear as buttons in the action panel in the
  right-hand side of the current window.
- On a mobile device, the default action views and toolbar buttons will appear in the [front-end chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

When creating action handlers with `ON KEY`
(or `COMMAND KEY` without a command name in a
`MENU`), the default
action view is invisible. If you define a `text` attribute in the action defaults,
the default action view is made visible.

Control the default action view visibility by using the `DEFAULTVIEW` action attribute.

If one or more action views are defined explicitly for a given action, the front-end considers
that the default view is not needed. Typically, if you define in the form a
`BUTTONEDIT` field, a `BUTTON`, or a `TOOLBAR` item
that triggers the action, you do not need an additional button in the action panel.

The presentation of default action views can be controlled with presentations style attributes
for the `Window` AUI tree nodes. See [Window style attributes: Action Panel](1657-window-style-attributes-action-panel.md "Presentation style attributes that apply to the window action panel.") and [Window style attributes: Ring Menu](1658-window-style-attributes-ring-menu.md "Presentation style attributes that apply to a window ring menu.").

Action views for the actions that are rendered as default action views can also be show in
topmenus and toolbars with `AUTOITEMS` and `AUTOCOMMANDS`
placeholders. See [Defining action views in forms](2278-defining-action-views-in-forms.md "How to define action views that will fire action events.") for more details.

## Related links

**Related concepts**  

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

**Related reference**  

[Window style attributes](1654-window-style-attributes.md "Window presentation style attributes apply to a window element.")
