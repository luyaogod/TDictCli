---
title: "Text attribute shows default action view"
source: "fgl-topics/c_fgl_prog_dialogs_action_text_default_view.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Text attribute shows default action view"
type: "concept"
description: "When creating actions with ON KEY (or COMMAND KEY without a command name in a MENU ), the default action view is invisible. However, when a text action attribute is defined for the corresponding key ..."
---

# Text attribute shows default action view

When creating actions with `ON KEY` (or `COMMAND KEY` without a
command name in a `MENU`), the [default action view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") is invisible.
However, when a `text` action attribute is defined for the corresponding key action,
the default action view is made visible.

You can also control the visibility of the default action view with the
`DEFAULTVIEW` action attribute.

Note that it is also possible to set key labels with form attributes (`KEY`) or
with function calls (`FGL_SETKEYLABEL()`), this feature is supported for backward
compatibility. Use action default text attributes in new developments.

## Related links

**Related concepts**  

[Setting action key labels](2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.")
