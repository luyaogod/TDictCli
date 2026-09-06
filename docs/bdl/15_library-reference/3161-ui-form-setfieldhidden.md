---
title: "ui.Form.setFieldHidden"
source: "fgl-topics/c_fgl_ClassForm_setFieldHidden.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setFieldHidden"
type: "concept"
---

# ui.Form.setFieldHidden

> Show or hide a form field.

## Syntax

```
setFieldHidden(
   name STRING,
   hidden INTEGER )
```

1. name defines the name of the form field, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).
2. hidden the integer value to show or hide the element.

## Usage

Change the visibility of a form field with the `setFieldHidden()` method.

> **Important:**
>
> Make sure that the form layout is prepared to adapt, when elements are hidden dynamically.
> Containers such as `TABLE` allow hiding form fields / columns with little impact on
> the layout. Make also sure that the dialog controlling the form is prepared for fields hiding: Any
> hidden field (directly or because a parent element is hidden), will be considered as a disabled
> field.

Pass the identifier of the form field, as defined in the form definition. The
form field is identified by column name, with an optional prefix (`table.column` or
`column`).

The value passed to hide/show the element can be 0, 1 or 2:

| Hidden value | Description |
| --- | --- |
| `0` | Makes the field visible. |
| `1` | The field is hidden and the user cannot make it visible. Typically used to hide information the user is not allowed to see. |
| `2` | The element is hidden and the user can make it visible. |

Hiding form elements by program can be used in conjunction with [`HIDDEN@screen-size`](../11_user-interface/1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.")
attributes, to get responsive layout:

- When hiding an element by program (value=1 or 2), it will always be hidden to the user, no
  matter the `HIDDEN@screen-size` definitions.
- When showing an element by program (value=0), it will be visible or hidden, depending on the
  `HIDDEN@screen-size` definitions.

Do not disable or hide all fields of a dialog, otherwise the dialog
execution stops: At least one field must be able to get the focus during a dialog execution.

Do not disable/hide the current field having the focus. When
disabling or when hiding the current field, the `AFTER FIELD` block of the current
field and the `BEFORE FIELD` block on the next field in the tabbing order will be
executed. In a multiple dialog, if the next field in the tabbing order is in another sub-dialog, the
`AFTER INPUT` block of the current sub-dialog and the `BEFORE
INPUT/CONSTRUCT/DISPLAY` or the next sub-dialog are executed. In all these control blocks, a
`NEXT FIELD` to the field that is disabled or hidden will result in an endless
loop.

## Related links

**Related concepts**  

[Example 2: Hide form elements dynamically](3167-example-2-hide-form-elements-dynamically.md "Example 2: Hide form elements dynamically")

[ui.Form.setFieldStyle](3162-ui-form-setfieldstyle.md "Change the style of a form field.")

[ui.Form.setElementStyle](3158-ui-form-setelementstyle.md "Change the style of form elements.")
