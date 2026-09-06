---
title: "ROWBOUND action attribute"
source: "fgl-topics/c_fgl_action_attribute_ROWBOUND.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > ROWBOUND action attribute"
type: "concept"
---

# ROWBOUND action attribute

> The ROWBOUND attribute defines if the action is related to the row context of a record list.

## Syntax

(only in action handlers of record list dialog)

```
ROWBOUND
```

## Usage

The `ROWBOUND` attribute is to be used in a `DISPLAY ARRAY` or
`INPUT ARRAY` dialog action handler, when the action depends on the row context.
Actions marked with this attribute will be automatically enabled/disabled, based on current row
existence.

Rowbound action views can be accessed with the three-dot button on the right of
`TABLE`, `TREE` and `SCROLLGRID` rows.
If the action is disabled, the corresponding rowbound action view
of the three-dots button will not be displayed.

The `ROWBOUND` attribute can only be specified in a list handling dialog
(`INPUT ARRAY` or `DISPLAY ARRAY`), as action attribute in the
`ATTRIBUTES()` clause of `ON ACTION` handlers, and applies to the
actions defined by the current dialog in the current window.

Automatic `update` and `delete` actions of `INPUT
ARRAY`, as well as actions created by `ON UPDATE` and `ON
DELETE` modification trigger in `DISPLAY ARRAY`, are implicitly rowbound
actions. Such actions must be available only when at least one row exists in the list. Therefore,
the `ROWBOUND` attribute cannot be specified in `DISPLAY ARRAY`
modification triggers.

> **Important:**
>
> The main purpose of the [`ROWBOUND`](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.") attribute is to have the corresponding action enabled and disabled
> automatically by the runtime system, depending on the existence of a row. The default decoration of
> such action is front-end platform driven and is the consequence of using the
> `ROWBOUND` action attribute. A `DISPLAY ARRAY` or `INPUT
> ARRAY` can be used with a plain form that show a single row at a time. The rowbound actions
> can also be used in such case, using form buttons as action views for example. These buttons will be
> automatically enabled/disabled according to the list content.

## Example

```
DISPLAY ARRAY ...
   ...
   ON ACTION print ATTRIBUTES(ROWBOUND,TEXT="Print")
      CALL print_customer_info(arr_curr())
   ...
```

## Related links

**Related concepts**  

[Actions bound to the current row](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.")

[Default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.")
