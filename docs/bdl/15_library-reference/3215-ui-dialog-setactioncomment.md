---
title: "ui.Dialog.setActionComment"
source: "fgl-topics/c_fgl_ClassDialog_setActionComment.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setActionComment"
type: "concept"
---

# ui.Dialog.setActionComment

> Set the comment/hint of a default action view.

## Syntax

```
setActionComment(
   name STRING,
   comment STRING )
```

1. name is the name of the action, see [Identifying actions in dialog methods](3238-identifying-actions-in-ui-dialog-methods.md).
2. comment is the comment to be set.

## Usage

Use the `setActionComment()` method to define the comment/hint for the [default view](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") of an
action.

```
CALL DIALOG.setActionComment( "confirm", "Confirm current order" )
```

The first parameter identifies the action object of the dialog.

This method will only affect the rendering of
implicit action views such as [default action views](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action
views](../11_user-interface/2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](../11_user-interface/2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row."):
The explicit action views (such as buttons in the form layout) will not get the new
attribute value. Note that the `comment` attribute will not be rendered on context
menu and rowbound action views.

## Related links

**Related concepts**  

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
