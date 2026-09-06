---
title: "DEFAULTVIEW action attribute"
source: "fgl-topics/c_fgl_action_attribute_DEFAULTVIEW.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > DEFAULTVIEW action attribute"
type: "concept"
---

# DEFAULTVIEW action attribute

> The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.

## Syntax 1 (Dialog action handlers and form action defaults):

```
DEFAULTVIEW = { AUTO | YES | NO }
```

## Syntax 2 (Global .4ad action defaults file):

```
defaultView = { "yes" | "no" | "auto" }
```

## Usage

`DEFAULTVIEW` is an
action attribute defining whether the default action view (a button) must be displayed for an
action.

Note that default action views can be hidden with [`DIALOG.setActionHidden()`](../15_library-reference/3216-ui-dialog-setactionhidden.md "Showing or hiding a default action view.").

Possible values for `DEFAULTVIEW` are:

- `NO` indicates that no default action view must be displayed for this
  action.
- `YES` indicates that a default action view must always be displayed for this
  action, except if it is hidden by `DIALOG.setActionHidden()`.
- `AUTO` means that a default action view is displayed, if no explicit action view
  is defined for that action, and not hidden by `DIALOG.setActionHidden()`.

The default is `AUTO`.

This attribute applies to the actions defined by the current dialog
in the current window. It can be specified as action default attribute in a global
.4ad file, in the `ACTION DEFAULTS` section of form files, or as
dialog action attribute (using `ON ACTION name
ATTRIBUTES(...)`).

The `DEFAULTVIEW` attribute does not apply to default action views displayed by a
`MENU` instruction using the `STYLE="dialog"` or
`"popup"` attribute: `DEFAULTVIEW` is related to the action panel of a
window, [pop-up and dialog menus](1909-rendering-modes-of-a-menu.md) have no relation to the
action panel.

## Example

```
-- As action handler attribute
ON ACTION zoom ATTRIBUTES(DEFAULTVIEW=YES)

-- As action default
ACTION DEFAULTS
  ACTION zoom (DEFAULTVIEW=YES)
END

-- In a global action defaults file
<ActionDefault name="zoom" defaultView="yes" ... />
```

## Related links

**Related concepts**  

[DEFAULTVIEW attribute](1773-defaultview-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.")

[Default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.")

[ui.Dialog.setActionHidden](../15_library-reference/3216-ui-dialog-setactionhidden.md "Showing or hiding a default action view.")
