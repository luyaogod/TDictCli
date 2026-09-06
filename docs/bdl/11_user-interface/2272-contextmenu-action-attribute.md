---
title: "CONTEXTMENU action attribute"
source: "fgl-topics/c_fgl_action_attribute_CONTEXTMENU.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > CONTEXTMENU action attribute"
type: "concept"
---

# CONTEXTMENU action attribute

> The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action.

## Syntax 1 (Dialog action handlers and form action defaults):

```
CONTEXTMENU = { AUTO | YES | NO }
```

## Syntax 2 (Global .4ad action defaults file):

```
contextMenu = { "yes" | "no" | "auto" }
```

## Usage

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

## Example

```
-- As action handler attribute
ON ACTION zoom ATTRIBUTES(CONTEXTMENU=YES)

-- As action default
ACTION DEFAULTS
  ACTION zoom (CONTEXTMENU=YES)
END

-- In a global action defaults file
<ActionDefault name="zoom" contextMenu="yes" ... />
```

## Related links

**Related concepts**  

[CONTEXTMENU attribute](1768-contextmenu-attribute.md "The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action.")

[ui.Dialog.setActionHidden](../15_library-reference/3216-ui-dialog-setactionhidden.md "Showing or hiding a default action view.")
