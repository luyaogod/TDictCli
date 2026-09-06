---
title: "Actions bound to the current row"
source: "fgl-topics/c_fgl_prog_dialogs_rowbound_actions.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Actions bound to the current row"
type: "concept"
---

# Actions bound to the current row

> Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.

When using a `DISPLAY ARRAY` or `INPUT ARRAY` dialog to control a
table view, actions can get the [`ROWBOUND`](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.") attribute, in order to make the action only available when there
is a current row in the list.

The `ROWBOUND` attribute is typically used with `TABLE` and
`TREE` containers, but can also by used with `SCROLLGRID` and static
lists in `GRID` containers (however in last case there will be no row-specific
decoration of the action views)

> **Important:**
>
> The main purpose of the [`ROWBOUND`](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.") attribute is to have the corresponding action enabled and disabled
> automatically by the runtime system, depending on the existence of a row. The default decoration of
> such action is front-end platform driven and is the consequence of using the
> `ROWBOUND` action attribute. A `DISPLAY ARRAY` or `INPUT
> ARRAY` can be used with a plain form that show a single row at a time. The rowbound actions
> can also be used in such case, using form buttons as action views for example. These buttons will be
> automatically enabled/disabled according to the list content.

In `DISPLAY ARRAY`, the actions created from [modification triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.") `ON UPDATE`
and `ON DELETE` are automatically defined as rowbound actions. Actions for `ON
INSERT` and `ON APPEND` are not considered as specific to a row.

The actions defined with the `ROWBOUND` attribute will be available by selecting
the three-dots button on the right of a table row.
If the action is disabled, the corresponding rowbound action view
of the three-dots button will not be displayed.

In the following example, the `DISPLAY ARRAY` dialog implements three actions:

- The "refresh" action is not "rowbound", and will always be available (active/visible), even if
  the list is empty.
- The "check" action is rowbound, and will only be available if there is a (current) row in the
  list.
- The "delete" action created by the `ON DELETE` modification trigger is implicitly
  "rowbounded".

```
DISPLAY ARRAY a_orders TO sr.* ATTRIBUTES(UNBUFFERED)
   ...
   ON ACTION refresh -- not rowbound
      CALL fetch_orders()
   ON ACTION check ATTRIBUTES(ROWBOUND,TEXT="Check")
      CALL check_order(arr_curr())
   ON DELETE -- implicitly rowbound (text defined in default.4ad)
      CALL delete_order(arr_curr())
   ...
END DISPLAY
```

## Related links

**Related concepts**  

[Default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.")

[Binding action views to a table row](2322-binding-action-views-to-a-table-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering.")

[Binding action views to a scrollgrid row](2341-binding-action-views-to-a-scrollgrid-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering.")

[Binding action views to a treeview row](2360-binding-action-views-to-a-treeview-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering.")
