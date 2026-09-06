---
title: "BUTTONEDIT and ON ACTION INFIELD"
source: "fgl-topics/c_fgl_Migrate_to_310_buttonedit_action.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > BUTTONEDIT and ON ACTION INFIELD"
type: "concept"
---

# BUTTONEDIT and ON ACTION INFIELD

> With ON ACTION INFIELD, a BUTTONEDIT action is now always considered as a field-qualified action.

Before version 3.10, when implementing the action handler for a `BUTTONEDIT`
action with `ON ACTION action-name INFIELD
field-name`, it was mandatory to add the
`"field-name."` prefix in the `ACTION` attribute,
to make the button always active, even when the focus was not in the
field:

```
-- Form file:
BUTTONEDIT f1 = customer.cust_city, ACTION = cust_city.zoom;

-- Program file:
ON ACTION zoom INFIELD cust_city
```

When specifying only the action name in the `ACTION` attribute, the
`BUTTONEDIT` button was only enabled when the focus was in the field. However, this
is not the expected behavior of a `BUTTONEDIT` button.

Starting with Genero 3.10, the expected behavior is now implicitly achieved, even when the field
name is not specified in the `ACTION`
attribute:

```
BUTTONEDIT f1 = customer.cust_city, ACTION = zoom;
```

For more details see [Field-specific actions (INFIELD clause)](../11_user-interface/2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus."), [BUTTONEDIT item type](../11_user-interface/1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.").
