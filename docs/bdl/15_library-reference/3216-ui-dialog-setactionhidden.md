---
title: "ui.Dialog.setActionHidden"
source: "fgl-topics/c_fgl_ClassDialog_setActionHidden.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setActionHidden"
type: "concept"
---

# ui.Dialog.setActionHidden

> Showing or hiding a default action view.

## Syntax

```
setActionHidden(
   name STRING,
   hidden BOOLEAN )
```

1. name is the name of the action, see [Identifying actions in dialog methods](3238-identifying-actions-in-ui-dialog-methods.md).
2. hidden is a boolean value.

## Usage

Use the `setActionHidden()` method to hide the [default view](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") (and [context menu](../11_user-interface/1603-action-default-attributes-reference-4ad.md "This topic contains all attributes you can define in a .4ad action defaults file.") option) of an
action.

```
CALL DIALOG.setActionHidden( "confirm", TRUE )
```

The first parameter identifies the action object of the dialog.

The second parameter indicates if the implicit action views corresponding to this action must be
visible (`TRUE`) or hidden (`FALSE`).

This method will only affect the rendering of
implicit action views such as [default action views](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action
views](../11_user-interface/2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](../11_user-interface/2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row."):
The explicit action views (such as buttons in the form layout) will not get the new
attribute value. Note that the `comment` attribute will not be rendered on context
menu and rowbound action views.
