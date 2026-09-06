---
title: "Dialog action handler attributes"
source: "fgl-topics/c_fgl_prog_dialogs_on_action_attributes.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Dialog action handler attributes"
type: "concept"
description: "Action attributes can be specified at the dialog instruction level for implicit action views such as default action views , context menu action views and rowbound action views . The action attributes ..."
---

# Dialog action handler attributes

Action attributes can be specified at the dialog instruction level for implicit action views such
as [default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action views](2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.").

The action attributes specified at the dialog-level will overwrite the attributes defined in
action defaults.

To define dialog-level action attributes for an action, add the
`ATTRIBUTES()` clause to `ON ACTION`, with a
comma-separated list of action default attributes:

```
ON ACTION print
   ATTRIBUTES (TEXT = "Print",
               COMMENT = "Print the current record",
               IMAGE = "printer",
               VALIDATE = NO)
```

It is possible to use localized strings in action attributes such as `TEXT` and
`COMMENT`:

```
ON ACTION print
   ATTRIBUTES (TEXT = %"common.print.label",
               COMMENT = %"common.print.comment",
               ... )
```

Dialog-level action attributes are typically used when the dialog is not related to a
specific form, for example with independent `MENU` dialogs.

If the current form defines explicit action views (buttons in layout, toolbar buttons,
topmenu items) with the same name as the `ON ACTION` handler defining
action attributes with the `ATTRIBUTES()` clause, the explicit action
views will not get the action attributes defined by the `ON ACTION`.

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
