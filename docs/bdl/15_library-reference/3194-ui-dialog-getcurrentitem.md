---
title: "ui.Dialog.getCurrentItem"
source: "fgl-topics/c_fgl_ClassDialog_getCurrentItem.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getCurrentItem"
type: "concept"
---

# ui.Dialog.getCurrentItem

> Returns the current item having focus.

## Syntax

```
getCurrentItem()
  RETURNS STRING
```

## Usage

The `getCurrentItem()` method returns the name of the current form item having the
focus, or the last triggered `MENU` action.

When the dialog is a [`MENU`](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from."),
`getCurrentItem()` returns the name of the last triggered action, no matter the type
of action handler (`COMMAND`, `COMMAND KEY` or `ON
ACTION`).

For all kind of dialogs except [`MENU`](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from."),
when the focus is on an action view (typically, a [`BUTTON`](../11_user-interface/1684-button-item-type.md "Defines a push-button that can trigger an action.") in the form layout), `getCurrentItem()` returns the
name of the corresponding action. If several action views are bound to the same action handler with
a unique name, there is no way to distinguish which action view has the focus.

If the focus is in a simple field controlled by an [`INPUT`](../11_user-interface/2083-the-input-sub-dialog.md "The INPUT sub-dialog implements single record input in fields of the current form.") or [`CONSTRUCT`](../11_user-interface/2084-the-construct-sub-dialog.md "The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.") singular- or sub-dialog, `getCurrentItem()`
returns the [table-name.]field-name of that current field. The
table-name prefix is added if a `FROM` clause is used with an
explicit list of fields. No prefix is added if `FROM screen-record.*` is used or if
`BY NAME` clause is used. In the context of a `BEFORE INPUT` or
`BEFORE CONSTRUCT` block, the `getCurrentItem()` method always returns
`NULL`.

If the focus is in a list controlled by a [`DISPLAY ARRAY`](../11_user-interface/2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.") singular- or sub-dialog, `getCurrentItem()`
returns the screen-array name identifying the list for a regular `DISPLAY
ARRAY`. If the `DISPLAY ARRAY` is defined with the
`FOCUSONFIELD` attribute, the method returns
screen-array.field-name. In some context, the current field is undefined. For
example when entering the `DISPLAY ARRAY`, `getCurrentItem()` will
only return the screen-array, in the context of the `BEFORE ROW`
control block.

If the focus is in a field of a list controlled by an [`INPUT ARRAY`](../11_user-interface/2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.") singular- or
sub-dialog, `getCurrentItem()` returns screen-array.field-name,
identifying both the list and the current field. In some context, the current field is undefined.
For example when entering the `INPUT ARRAY`, `getCurrentItem()` will
only return the screen-array, in the context of the `BEFORE ROW`
control block. In the context of a `BEFORE INPUT` or `BEFORE
CONSTRUCT` block, the `getCurrentItem()` method always returns
`NULL`.
