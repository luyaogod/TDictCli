---
title: "Action views on mobile devices"
source: "fgl-topics/c_fgl_prog_dialogs_action_views_mobile.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Action views on mobile devices"
type: "concept"
---

# Action views on mobile devices

> Action views are rendered following mobile specific standards.

## Action views in the front-end chromebar

The default action views of the action panel, ring menu panel and the toolbar action views are by
default displayed in the chromebar, to adapt to mobile device GUI standards.

This can be controlled with presentation style attributes.

For more details, see [Action views in chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

## The Back button

The Back button on both Android™ and iOS devices is considered a default action view for the `close`,
`cancel`, or `accept` action of the current dialog:

- If a `close` action is defined and active, it is assigned to the back
  button.
- If no active `close` action exists in the current dialog, but the
  `cancel` action is defined and active, it is assigned to the back button.
- If neither `close` nor `cancel` actions are defined / active, but
  the `accept` action is defined and active, it is assigned to the back button.

For more details about the close action usage, see [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). Read also [Implementing the back action](2290-implementing-the-back-action.md "The back action is a predefined action dedicated to move back in the stack of windows/forms.").
