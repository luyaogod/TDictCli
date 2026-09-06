---
title: "Defining the action for a row choice"
source: "fgl-topics/c_fgl_ui_tables_curr_row_select.html"
breadcrumb: "User interface > User interface programming > Table views > Defining the action for a row choice"
type: "concept"
---

# Defining the action for a row choice

> The row choice in a TABLE can be associated with a dedicated action.

When using a [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
dialog to control a list view like a `TABLE`, the physical event to choose a row has
the following results, depending on the type of front-end:

- With the desktop / web browser front-end, by default, a mouse double-click changes the current
  row, and fires the "accept" action, if available. If the default accept action is fired, the dialog
  will end, except if the accept action has been disabled or was overwritten by an `ON ACTION
  accept` handler. This default behavior applies to most record lists of desktop applications,
  where the main purpose is to let the user choose a row from the list. The
  `DOUBLECLICK` attribute can be used to define a different action to be fired for a
  double-click with the mouse. If needed, the physical event (single or double click) to trigger this
  action can be defined with the `rowActionTrigger` style attribute.
- With a mobile device, a single tap and double tap on a row only changes the current row. If a
  single or double tap on a row must fire an action, define the action to be triggered with the
  `DOUBLECLICK` attribute, and configure the physical event to fire this action with
  the `rowActionTrigger` style attribute.

The action to be fired on double click (desktop) or double tap (mobile) can be defined with the
`DOUBLECLICK` option of `DISPLAY ARRAY` or with the [`DOUBLECLICK`](1775-doubleclick-attribute.md "The DOUBLECLICK attribute defines the action for row choice on TABLE/TREE/SCROLLGRID rows.") attribute of
`TABLE`, `TREE` or `SCROLLGRID` list container in the
form definition file. See [next
section](2304-defining-the-action-for-a-row-choice.md) for more details.

When using an [`INPUT ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
dialog, no row choice action ("accept") is possible by default: A mouse double-click is typically
used to select text in the current cell. However, when using the `DOUBLECLICK`
attribute in the `TABLE`, `TREE` or `SCROLLGRID`, a
double-click can fire the specified action during an `INPUT ARRAY`. See [Using the DOUBLECLICK attribute with INPUT ARRAY](2304-defining-the-action-for-a-row-choice.md).

For more details see [Defining the action for a row choice](2304-defining-the-action-for-a-row-choice.md "The row choice in the DISPLAY ARRAY/INPUT ARRAY dialog can be associated with a dedicated action.").
