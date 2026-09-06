---
title: "Action display in the context menu"
source: "fgl-topics/c_fgl_prog_dialogs_contextmenu.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Action display in the context menu"
type: "concept"
---

# Action display in the context menu

> The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.

A context menu can be shown with a mouse right-click, with all the active actions
that are possible in the current form.

> **Tip:**
>
> By default, when using the GAS/browser, a right-click with the mouse opens the browser context
> menu. See the Genero Browser Client User Guide, to understand how to enable the BDL context menu.

Displaying all actions might not be adapted to your needs. To control if an action must be
displayed in the context menu, set the `CONTEXTMENU` attribute in action defaults.
Values for `CONTEXTMENU` can be `YES`, `NO` and
`AUTO`.

```
ACTION DEFAULTS
   ...
   ACTION insert ( ... CONTEXTMENU = YES ... )
   ACTION append ( ... CONTEXTMENU = YES ... )
   ACTION delete ( ... CONTEXTMENU = YES ... )
   ...
   ACTION validate_order ( ... CONTEXTMENU = NO ... )
   ...
END
```

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
