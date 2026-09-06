---
title: "SIZEPOLICY attribute"
source: "fgl-topics/c_fgl_FSFAttributes_SIZEPOLICY.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > SIZEPOLICY attribute"
type: "concept"
---

# SIZEPOLICY attribute

> The SIZEPOLICY attribute is a sizing directive based on the content of a form item.

## Syntax

```
SIZEPOLICY = { INITIAL | FIXED | DYNAMIC }
```

## Usage

The `SIZEPOLICY` attribute defines how front-ends compute the widget size (width
and height), or the widget width only of form items, based on the content, or definition of the form
field or form item.

For most form items types like `LABEL`, `CHECKBOX`, or
`COMBOBOX`, the `SIZEPOLICY` attribute applies only to the width of
the widget. With two-dimensional form items such as `IMAGE`, the
`SIZEPOLICY` attribute applies to the width and height. Note that the height of a
(vertical) `RADIOGROUP` is defined by the number of `ITEMS`, and only
the width can be controlled with the `SIZEPOLICY` attribute.

The `SIZEPOLICY` attribute applies to AUI tree leaf elements of a form (it does
not apply to containers such as `TABLE` or `GRID`):
`SIZEPOLICY` can be used with form items having content with variable size such as
`IMAGE`, `LABEL` (for localization) and with form items where the
definition impacts the widget size, like the width of labels in `ITEMS` of a
`COMBOBOX`.

Elements allowing user input such as `EDIT`, or elements where the size does not
depend on the value of content such as `PROGRESSBAR`, `SLIDER` do not
use the `SIZEPOLICY` attribute.

In `TABLE` or `TREE` containers, `SIZEPOLICY` applies
only to `COMBOBOX`, `CHECKBOX` and `RADIOGROUP`
columns. This concerns the size policy specified by the user, or the default
(`INITIAL`). In such list views, for widgets like `LABEL`, the size
policy is implicitly fixed, and the column width is defined by the form layout.

When the `SIZEPOLICY` is not specified, it defaults to `INITIAL`.
The behavior then depends on the type of form item. For more details, see Table 1.

## SIZEPOLICY = FIXED

When `SIZEPOLICY` is `FIXED`, the form element's size (or its width
only) is defined in the layout of the form specification file.

The size of the element is computed from the width and height in the form grid and the font
used on the front-end side.

The element keeps the size, even if the content is modified. However, for form items
supporting the `STRETCH` attribute, if the `STRETCH` attribute is set
to `X`, `Y` or `BOTH`, the form element can still
stretch when the parent window size changes.

## SIZEPOLICY = DYNAMIC

When `SIZEPOLICY` is `DYNAMIC`, the form element's size (or its
width only) grows and shrinks to adapt to the size of the content (like the width of a
`LABEL` text), or to adapt to the widget definition (like the picture size for an
`IMAGE`), during the lifetime of the application.

This can be used for `COMBOBOX` or `RADIOGROUP` fields, when the
width of the widget must fit exactly to its content, which can vary during the program
execution.

With `SIZEPOLICY=DYNAMIC`, elements such as `BUTTON`,
`LABEL`, `IMAGE` and `RADIOGROUP` can shrink and grow,
while `COMBOBOX` elements can only grow.

With `WEBCOMPONENT` and `TEXTEDIT` fields,
`SIZEPOLICY=DYNAMIC` is ignored.

## SIZEPOLICY = INITIAL

When `SIZEPOLICY` is `INITIAL`, the first time the element appears
on the screen, its size (or its width only) is computed from the initial content (like the text of a
`LABEL`) or the widget definition (like the labels in `ITEMS` of
`COMBOBOX`). Once the widget displays, its size is frozen.

However, for form items supporting the `STRETCH` attribute, if the
`STRETCH` attribute is set to `X`, `Y` or
`BOTH`, the form element can still stretch when the parent window size changes.

`SIZEPOLICY=INITIAL` is ignored for `WEBCOMPONENT` fields and
`TEXTEDIT` fields.

`SIZEPOLICY=INITIAL` is typically used when the size of the element must be fixed,
but is not known at design time. For example, when populating a `COMBOBOX` item list
from a database table, the size of the `COMBOBOX` may need to be increased, based on
the size of items labels in the drop-down list. This size policy mode is also useful when the text
of labels is unknown at design time because of internationalization.

With `SIZEPOLICY=INITIAL`, the behavior differs depending on the form element
type.

| Form item | Behavior with SIZEPOLICY=INITIAL |
| --- | --- |
| [`BUTTON`](1684-button-item-type.md "Defines a push-button that can trigger an action.") | The size defined in the form is a minimum size. If the initial button text is bigger, the size grows (width and height). |
| [`CHECKBOX`](1686-checkbox-item-type.md "Defines a boolean or three-state checkbox field.") | The width defined in the form is a minimum width. If the checkbox label (`TEXT`) is larger as the minimum width, this defines the initial size of the widget. Then, the checkbox size remains fixed for the life time of the form. |
| [`COMBOBOX`](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.") | The width defined in the form is a minimum width. If a label of combobox `ITEMS` is larger than the minimum width, the size grows in order for the combobox to fully display the largest item. Then, the combobox size remains fixed for the life time of the form. |
| [`LABEL`](1696-label-item-type.md "Defines a simple text area to display a read-only value.") | The width defined in the form is a minimum width. If the length of the initial field value (or the length of the `TEXT` attribute for static labels) is larger as the minimum width, this defines the initial size of the widget. Then, the label keeps the same size for the life time of the form. The size defined in the form is ignored. |
| [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.") | Image form items adapt their size to the initial image displayed and keep that size. If no initial image is displayed (i.e. the image field value is `NULL`), the form item does not take space in the layout (and also does not adapt the size if an image is displayed later on). Images can use the `STRETCH` attribute, so that the widget size is dependent on the parent container, overriding the `SIZEPOLICY` attribute. If the `WIDTH` and `HEIGHT` attributes must be used, the `SIZEPOLICY` attribute must be set to `FIXED`. |
| [`RADIOGROUP`](1699-radiogroup-item-type.md "Defines a mutual exclusive set of options field.") | For vertical radiogroups: The height of the widget is defined by the number of `ITEMS`. The width defined in the form is a minimum width. If a label of radiogroup `ITEMS` is larger than the minimum width, the width grows in order for the radiogroup to fully display the largest item. Then, the radiogroup size remains fixed for the life time of the form. |
| [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") | N/A: `SIZEPOLICY=INITIAL` is ignored for `TEXTEDIT` fields. |
| [`WEBCOMPONENT`](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.") | N/A: `SIZEPOLICY=INITIAL` is ignored for `WEBCOMPONENT` fields. |

## Example

```
IMAGE img1 = FORMONLY.main_image,
    SIZEPOLICY=DYNAMIC;

COMBOBOX f001 = customer.city,
    ITEMS=((1,"Paris"),(2,"Madrid"), (3,"London")),
    SIZEPOLICY=INITIAL; -- Default

WEBCOMPONENT wc1 = FORMONLY.chart,
    COMPONENTTYPE="chart",
    SIZEPOLICY=FIXED,
    STRETCH=BOTH;
```

## Related links

**Related concepts**  

[Controlling the image layout](1585-controlling-the-image-layout.md "Explains how image form items can be sized in different front-end layout systems.")

[Web components](2378-web-components.md "This section describes how to use web components in your application.")

[STRETCH attribute](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")
