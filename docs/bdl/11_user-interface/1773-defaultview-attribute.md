---
title: "DEFAULTVIEW attribute"
source: "fgl-topics/c_fgl_FSFAttributes_DEFAULTVIEW.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > DEFAULTVIEW attribute"
type: "concept"
---

# DEFAULTVIEW attribute

> The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.

## Syntax

```
DEFAULTVIEW = { AUTO | YES | NO }
```

## Usage

This attribute is an action attribute that can be specified in form `ACTION
DEFAULTS`.

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

For more details, see [DEFAULTVIEW action attribute](2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.").

## Related links

**Related concepts**  

[DEFAULT attribute](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.")
