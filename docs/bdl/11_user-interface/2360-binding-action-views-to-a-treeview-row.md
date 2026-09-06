---
title: "Binding action views to a treeview row"
source: "fgl-topics/c_fgl_treeviews_rowbound_actions.html"
breadcrumb: "User interface > User interface programming > Tree views > Binding action views to a treeview row"
type: "concept"
---

# Binding action views to a treeview row

> A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering.

Rowbound actions are action defined with the [`ROWBOUND`](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.") action attribute, when using a `DISPLAY ARRAY` or
`INPUT ARRAY` dialog to control a table or tree view:

```
DISPLAY ARRAY arr TO sr.*
   ...
   ON ACTION copy_row ATTRIBUTES(ROWBOUND, TEXT="Copy row")
      ...
```

Rowbound actions are available to the end user as options of a three-dots button, on the right of
`TABLE`, `TREE` and `SCROLLGRID` rows:

- On devices with mouse pointer (desktop), the three-dots button appears when the mouse hovers a
  row.
- On touch devices (mobile), the three-dots button is displayed when selecting the current
  row.

If the action is disabled, the corresponding rowbound action view
of the three-dots button will not be displayed.

![Treeview with rowbound actions](../_images/TreeviewWithRowBoundActions2.jpg)

*Treeview with rowbound actions*

> **Important:**
>
> The main purpose of the [`ROWBOUND`](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.") attribute is to have the corresponding action enabled and disabled
> automatically by the runtime system, depending on the existence of a row. The default decoration of
> such action is front-end platform driven and is the consequence of using the
> `ROWBOUND` action attribute. A `DISPLAY ARRAY` or `INPUT
> ARRAY` can be used with a plain form that show a single row at a time. The rowbound actions
> can also be used in such case, using form buttons as action views for example. These buttons will be
> automatically enabled/disabled according to the list content.

## Related links

**Related concepts**  

[Actions bound to the current row](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.")
