---
title: "Defining the tabbing order"
source: "fgl-topics/c_fgl_prog_dialogs_tabbing_order.html"
breadcrumb: "User interface > User interface programming > Input fields > Defining the tabbing order"
type: "concept"
---

# Defining the tabbing order

> Control the order of tabbing through the fields with the TABINDEX attribute.

When a dialog is executing, the end-user can jump from field to field with the keyboard by using
the Tab and Shift-Tab keys. One can tab out of an `INPUT ARRAY` sub-dialog with Ctrl-Tab
and Shift-Ctrl-Tab accelerators (in `INPUT ARRAY`, Tab and Shift-Tab loop in the
fields of the current row).

Consider using the [`OPTIONS INPUT
WRAP`](../09_advanced-features/0927-defining-the-field-input-loop.md "The OPTIONS INPUT [NO] WRAP instructions defines field wrapping in dialogs.") instruction to force the focus to stay in the dialog, when tabbing out of the
last form element controlled by the dialog.

The order in which the fields can be visited with the Tab key can be controlled with a program
option and the `TABINDEX` form field attribute.

The [`FIELD
ORDER`](2089-dialog-attributes-clause.md) dialog attribute defines the way tabbing order works (it can also be defined
globally with `OPTIONS FIELD ORDER mode`). Tabbing order can be
based on the dialog binding list (`FIELD ORDER CONSTRAINED`, the default) or it can
be based on the form tabbing order (`FIELD ORDER FORM`). It is recommended that you
use the `FIELD ORDER FORM` option, to use the tabbing order specified in the form
file.

The [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.") field
attribute allows tabbing order in the form to be defined for each form item. By default,
the form compiler assigns a tabbing index for each form item based on the position of
the item in the layout.

> **Important:**
>
> `TABINDEX` values must be unique in a form.

Form elements that can get the focus are:

- Simple form fields controlled by `INPUT`
  or `CONSTRUCT`,
- Read-only lists or treeviews controlled by `DISPLAY ARRAY`,
- Editable list cells controlled by `INPUT ARRAY`,
- Simple buttons controlled by a `COMMAND`
  interaction block.

If you use the keyboard to tab into a form element, the focus will go to the next (or previous)
element that is visible and activated. In other words, if a form item is hidden or
disabled, it is removed from the tabbing list.

The tabbing position of a read-only list driven by a `DISPLAY ARRAY` binding is
defined by the `TABINDEX` of the first field.

When `TABINDEX` is set to zero, the form item is excluded from the tabbing list.
However, the item with `TABINDEX=0` can still get the focus with the mouse (or when
you tap on it on a mobile device).

The [`NEXT FIELD`](1955-next-field-instruction.md) instruction can
also use the tabbing order, when executing `NEXT FIELD NEXT` and `NEXT FIELD
PREVIOUS`.

When moving columns around in a [`TABLE`](1703-table-item-type.md "Defines a list view widget.") or [`TREE`](1706-tree-item-type.md "Defines a tree view widget.") container, new tab indexes are assigned to columns, so that input
tabbing order corresponds to the visual column order. If the order of the columns in an editable
list shouldn't be changed, freeze the table columns with the [`UNMOVABLECOLUMNS`](1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table.") attribute
of the `TABLE`/`TREE`.

## Related links

**Related concepts**  

[Dialog configuration with FGLPROFILE](2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior.")
