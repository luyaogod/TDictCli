---
title: "Form items"
source: "fgl-topics/c_fgl_FormSpecFiles_Form_Items.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items"
type: "concept"
---

# Form items

> The concept of form item includes all elements used in the definition of a form.

## What is a form item?

A form item can be an input field such as an `EDIT` field, a push
`BUTTON`, a `GROUPBOX`, or a `TABLE`
container. A form item can also be an element of a `TOOLBAR`,
`TOPMENU`, and `ACTION DEFAULTS` definition.

A form item can be:

- [A satellite item](1666-satellite-items.md "Satellite items are display elements defined outside the LAYOUT section.")
- [A static item](1667-static-items.md "A static item defines a simple form item as a final grid element that does not change.")
- [A layout item](1668-layout-items.md "Layout items are containers with a body that can hold other form items.")
- [An action view](1669-action-views.md "An action view defines a form item that can trigger an action in the program.")
- [A form field](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

## Form item types

A form item is defined by its type, called a form item type. For
example, a form field can be an `EDIT`, or a `COMBOBOX`.
A form layout container can be a `GROUP`, or a `GRID`.
A toolbar item can be an `ITEM` or a `SEPARATOR`.

For a detailed description, see [Form item types](1683-form-item-types.md "The form item types defines the purpose of form elements.").

## Defining a form item

In a `GRID` container, form items (typically, form fields) must be
defined with a form tag in the [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section, bound by the tag name to a definition in the
[`ATTRIBUTES`](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
section.

The form tag defines the position and length of the form item, while the appearance
and the behavior of the form item is defined by a set of attributes in the
`ATTRIBUTES`
section:

```
LAYOUT
GRID
{
   [f1            ]
   ...
}
END
END
...
ATTRIBUTES
EDIT f1 = customer.cust_name, ... ;
END
```

## Satellite form items

Other kinds of form items are defined in the section it belongs to (for example, an
`ITEM` element of a `TOOLBAR` definition).

## Related links

**Related concepts**  

[Examples](1853-examples.md "Form definition (.per) examples.")

## Child topics

- [Satellite items](1666-satellite-items.md): Satellite items are display elements defined outside the LAYOUT section.
- [Static items](1667-static-items.md): A static item defines a simple form item as a final grid element that does not change.
- [Layout items](1668-layout-items.md): Layout items are containers with a body that can hold other form items.
- [Action views](1669-action-views.md): An action view defines a form item that can trigger an action in the program.
- [Form fields](1670-form-fields.md): Form fields are form elements designed for data input and/or data display.
- [Identifying form items](1675-identifying-form-items.md): Elements defined in a form file can be identified with a name, to be used in programs.
- [Screen records / arrays](1676-screen-records-arrays.md): Form fields can be grouped in a screen record or screen array definition.
- [Form tags](1677-form-tags.md): Form tags define layout elements inside a grid-based container.
