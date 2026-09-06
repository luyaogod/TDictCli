---
title: "CONTEXTMENU attribute"
source: "fgl-topics/c_fgl_FSFAttributes_CONTEXTMENU.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > CONTEXTMENU attribute"
type: "concept"
---

# CONTEXTMENU attribute

> The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action.

## Syntax

```
CONTEXTMENU = { AUTO | YES | NO }
```

## Usage

This attribute is an action attribute that can be specified in form `ACTION DEFAULTS`

`CONTEXTMENU` is an action attribute defining whether the context menu option must
be displayed for an action.

Actions to be displayed in a context menu must have a `TEXT` attribute. If the
`TEXT` attribute is not defined or is empty, the corresponding action view will not
be shown in the context menu.

If the action is disabled, the corresponding action view will not be displayed in the context
menu.

Possible values for `CONTEXTMENU` are:

1. `NO` indicates that no context menu option must be displayed for this
   action.
2. `YES` indicates that a context menu option must always be displayed for this
   action, if the action is visible.
3. `AUTO` means that the context menu option is displayed if no explicit action view
   is used for that action and the action is visible.

The default is `YES`.

This attribute applies to the actions defined by the current dialog
in the current window. It can be specified as action default attribute in a global
.4ad file, in the `ACTION DEFAULTS` section of form files, or as
dialog action attribute (using `ON ACTION name
ATTRIBUTES(...)`).

For more details, see [CONTEXTMENU action attribute](2272-contextmenu-action-attribute.md "The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action.").

## Related links

**Related concepts**  

[ui.Dialog.setActionHidden](../15_library-reference/3216-ui-dialog-setactionhidden.md "Showing or hiding a default action view.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Action display in the context menu](2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.")
