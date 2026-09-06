---
title: "TABINDEX attribute"
source: "fgl-topics/c_fgl_FSFAttributes_TABINDEX.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > TABINDEX attribute"
type: "concept"
---

# TABINDEX attribute

> The TABINDEX attribute defines the tab order for a form item.

## Syntax

```
TABINDEX = integer
```

1. integer defines the order of the item in the tab sequence.

## Usage

This attribute can be used to define the order in which the form items are selected as the user
"tabs" from field to field.

To take `TABINDEX` attributes into account in dialogs, the program must defined
the form tabbing order with the `OPTIONS FIELD ORDER FORM` instruction.
Alternatively, a dialog can use the `FIELD ORDER FORM` option as well.

`TABINDEX` can be set to zero in order to exclude the form item from the tabbing
list. The item can still get the focus with the mouse.

> **Important:**
>
> `TABINDEX` values must be unique in a form. Except when using zero.

The `TABINDEX` attribute can also be used to define which field must get the focus
when a `FOLDER` page is selected.

By default, form items get a tab index depending on the order in which they appear in the
`LAYOUT` section.

For more details, see [Defining field tabbing order method](../09_advanced-features/0928-defining-field-tabbing-order-method.md)

## Example

```
EDIT f001 = customer.fname, TABINDEX = 1;
EDIT f002 = customer.lname, TABINDEX = 2;
EDIT f003 = customer.comment,
     TABINDEX = 0; -- Excluded from tabbing list
```

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")

[FOLDER container](1720-folder-container.md "Defines the parent container for folder pages.")

[Defining the tabbing order](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.")
