---
title: "ui.Dialog.setActionText"
source: "fgl-topics/c_fgl_ClassDialog_setActionText.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setActionText"
type: "concept"
---

# ui.Dialog.setActionText

> Defining the text of a default action view.

## Syntax

```
setActionText(
   name STRING,
   text STRING )
```

1. name is the name of the action, see [Identifying actions in dialog methods](3238-identifying-actions-in-ui-dialog-methods.md).
2. text is the text to be set.

## Usage

Use the `setActionText()` method to define the label of the [default view](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") of an
action.

```
CALL DIALOG.setActionText( "confirm", "Confirm" )
```

The first parameter identifies the action object of the dialog.

The second parameter defines the label to be used in the implicit action view corresponding to
the specified action.

This method will only affect the rendering of
implicit action views such as [default action views](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action
views](../11_user-interface/2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](../11_user-interface/2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row."):
The explicit action views (such as buttons in the form layout) will not get the new
attribute value. Note that the `comment` attribute will not be rendered on context
menu and rowbound action views.

## Related links

**Related concepts**  

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
