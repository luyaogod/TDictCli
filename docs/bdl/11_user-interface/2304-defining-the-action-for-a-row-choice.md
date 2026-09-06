---
title: "Defining the action for a row choice"
source: "fgl-topics/c_fgl_prog_dialogs_row_selection_action.html"
breadcrumb: "User interface > User interface programming > List dialogs > Defining the action for a row choice"
type: "concept"
---

# Defining the action for a row choice

> The row choice in the DISPLAY ARRAY/INPUT ARRAY dialog can be associated with a dedicated action.

## Action fired by default by current row selection

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

## Defining the name of the row choice action

The action to be fired when a row choice occurs can be defined in the form file with the [`DOUBLECLICK`](1775-doubleclick-attribute.md "The DOUBLECLICK attribute defines the action for row choice on TABLE/TREE/SCROLLGRID rows.") attribute of the
`TABLE`, `TREE` or `SCROLLGRID` containers, or with the
`DOUBLECLICK` attribute of the [`DISPLAY ARRAY`](1965-display-array-instruction-configuration.md) dialog:

```
-- In .per form file definition:
TABLE (DOUBLECLICK=select, ...)
   ...

-- In .4gl program code:
DISPLAY ARRAY arr TO sr.*
        ATTRIBUTES(UNBUFFERED, DOUBLECLICK=select)
   ...
```

The `DOUBLECLICK` attribute in `DISPLAY ARRAY` takes precedence
over the `DOUBLECLICK` attribute in the list container of a form file.

When defining a `DOUBLECLICK` action in the list container of the form file, or in
a dialog attribute, you declare an explicit action view, and no [default action view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") will be displayed for
this action (you can force it with `DEFAULTVIEW=YES`).

## Using the DOUBLECLICK attribute with INPUT ARRAY

With an [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."), field editing
is implicit: A double click with the mouse is typically used to select text in an editable field.
Therefore, unlike `DISPLAY ARRAY`, no `accept` action is fired by
default when the user double-clicks on a cell during an `INPUT ARRAY`.

However, if the `TABLE`, `TREE` or `SCROLLGRID` list
container defines the [`DOUBLECLICK`](1775-doubleclick-attribute.md "The DOUBLECLICK attribute defines the action for row choice on TABLE/TREE/SCROLLGRID rows.") attribute, the corresponding action is fired when a double
click occurs, and a row choice action can be implemented in an `INPUT ARRAY`. This
configuration is typically used when the list container has columns defined with non-editable fields
like `LABEL` or `EDIT` using `NOENTRY`.

## Physical event triggering the row choice action

For desktop and web UI applications, the physical event that fires the row choice action is the
mouse double-click, which can be changed with the [`rowActionTrigger`](1648-table-style-attributes.md) presentation style attribute for the
`Table`, `ScrollGrid` and `Tree` classes.

For example, in web applications, the row choice is typically done by a single mouse click.
Define the simple-click physical event to fire the row choice action in your
.4st style file as follows:

```
...
  <Style name="Table">
     <StyleAttribute name="rowActionTrigger" value="singleClick" />
  </Style>
...
```

On touchscreen / mobile devices, `singleClick` corresponds to a single tap and
`doubleClick` corresponds to a double tap.

## Row choice action handler in programs

To handle row choice actions in the program code, define the `DOUBLECLICK`
attribute for the `DISPLAY ARRAY` dialog, and the corresponding action handler block
`ON ACTION action-name`:

```
DISPLAY ARRAY arr TO sr.*
        ATTRIBUTES(UNBUFFERED, DOUBLECLICK=select)
   ON ACTION select
      MESSAGE "Current selected row index:", DIALOG.getCurrentRow("sr")
END DISPLAY
```

If the `DOUBLECLICK` attribute is defined, it will only configure the action for
the corresponding physical event. By default, the "accept" action is still available, and the [Ok]
button or the [Return] key will still fire the accept action and leave the dialog. To avoid the
default accept action, add `ACCEPT=FALSE` to the `DISPLAY ARRAY`
attribute list.

By default, the row selection action can be fired (it is active), even when the array is
empty.

To automatically disable the row selection action when no rows exist in the list, use the [`ROWBOUND`](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row.") action with the
action defined by the `DOUBLECLICK` attribute. The action will then only be active,
if there is a current row in the list:

```
DISPLAY ARRAY arr TO sr.*
        ATTRIBUTES(UNBUFFERED, DOUBLECLICK=select)
    ON ACTION select ATTRIBUTES(ROWBOUND)
```

Otherwise, the program code needs to check if there is a current row, before executing the code
for the row selection:

```
DEFINE x INTEGER
DISPLAY ARRAY arr TO sr.*
        ATTRIBUTES(UNBUFFERED, DOUBLECLICK=select)
    ON ACTION select
       LET x = DIALOG.getCurrentRow("sr")
       IF x > 0 THEN
           DISPLAY "Selected row: ", arr[x].*
       END IF
...
```

## Execution order of row change control blocks

During a `DISPLAY ARRAY`, if the selected row is not the current row, the
`AFTER ROW` and `BEFORE ROW` control blocks execute before the
`ON ACTION` block, in the following order:

1. `AFTER ROW` (for the previous current row)
2. `BEFORE ROW` (for the new current row)
3. `ON ACTION double-click-action`

## Related links

**Related concepts**  

[DISPLAY ARRAY instruction configuration](1965-display-array-instruction-configuration.md "DISPLAY ARRAY instruction configuration")
