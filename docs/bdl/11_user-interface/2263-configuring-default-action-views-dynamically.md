---
title: "Configuring default action views dynamically"
source: "fgl-topics/c_fgl_prog_dialogs_defactview_methods.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Configuring default action views dynamically"
type: "concept"
description: "Attributes of default action views can be changed dynamically during a dialog execution with ui.Dialog.setAction*() methods. For example, to change the text and icon of the default action view bound ..."
---

# Configuring default action views dynamically

Attributes of [default action
views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") can be changed dynamically during a dialog execution with
`ui.Dialog.setAction*()`
methods.

For example, to change the text and icon of the default action view bound to the "print" action:

```
CALL DIALOG.setActionText("print", "Print order" )
CALL DIALOG.setActionImage("print", "printer" )
```

The corresponding default action view will be decorated with the new text and icon.

The attributes set with `ui.DIALOG.setAction*()` methods are volatile and last
only for the duration of the current dialog.

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
