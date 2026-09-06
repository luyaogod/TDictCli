---
title: "ui.Dialog.setActionImage"
source: "fgl-topics/c_fgl_ClassDialog_setActionImage.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setActionImage"
type: "concept"
---

# ui.Dialog.setActionImage

> Set the image of a default action view.

## Syntax

```
setActionImage(
   name STRING,
   image STRING )
```

1. name is the name of the action, see [Identifying actions in dialog methods](3238-identifying-actions-in-ui-dialog-methods.md).
2. image is the image name to be set.

## Usage

Use the `setActionImage()` method to define the image of the [default view](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") of an
action.

```
CALL DIALOG.setActionImage( "confirm", "smiley" )
```

The first parameter identifies the action object of the dialog.

The second parameter defines the image resource to be used in the implicit action view
corresponding to the specified action.

This method will only affect the rendering of
implicit action views such as [default action views](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action
views](../11_user-interface/2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](../11_user-interface/2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row."):
The explicit action views (such as buttons in the form layout) will not get the new
attribute value. Note that the `comment` attribute will not be rendered on context
menu and rowbound action views.

## Related links

**Related concepts**  

[ui.Dialog.setActionText](3218-ui-dialog-setactiontext.md "Defining the text of a default action view.")

[ui.Dialog.setActionComment](3215-ui-dialog-setactioncomment.md "Set the comment/hint of a default action view.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
