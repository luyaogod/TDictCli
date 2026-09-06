---
title: "HIDDEN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_HIDDEN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > HIDDEN attribute"
type: "concept"
---

# HIDDEN attribute

> The HIDDEN attribute indicates if an element must be visible to the user.

## Syntax

```
{ HIDDEN[@screen-size]
| HIDDEN = USER
}
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. `HIDDEN` sets the underlying item attribute to 1, to hide the element.
2. screen-size is a screen size selector that indicates if the element must be
   hidden, depending on the size of the screen. Several
   `HIDDEN@screen-size` attributes can be used for the same
   element.
3. With `HIDDEN=USER`, the underlying item attribute gets a value of 2, to hide the
   element by default, and let end-users show the element (typically for table columns). The
   `USER` option cannot be mixed with screen-size selectors.

## Usage

By default, all form elements are visible. Specify the `HIDDEN` attribute, to hide
a form element, such as a form field or a groupbox.

If the `HIDDEN` keyword is specified alone, the underlying item attribute is set
to 1. The value 1 indicates that the element is definitively hidden to the end user, which cannot
show the element, for example with the context menu of [`TABLE`](1703-table-item-type.md "Defines a list view widget.") headers. In this hidden mode, the [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu.") attribute is ignored by
the front-end.

With `HIDDEN=USER`, the underlying item attribute is set to 2. The value 2
indicates that the element is hidden by default, but the end user can show/hide the element as
needed, when the widget and its container provide a way to do so. For example, in a
`TABLE` or `TREE` container, the end user can use a context menu
option to show a column defined with `HIDDEN=USER` attribute. When
`HIDDEN` is used without the `USER` value, the column is definitively
invisible.

Form fields hidden with `HIDDEN=USER` (value 2) might be shown anyway, if the
field is needed by a dialog for input. In such case, the program dialog takes precedence over the
`HIDDEN` attribute.

The runtime system detects hidden form fields: If you write an [`INPUT`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") block using a hidden field, the
field is ignored, as if it was declared as [`NOENTRY`](1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog."), or disabled with [`DIALOG.setFieldActive("fieldname", FALSE)`](../15_library-reference/3226-ui-dialog-setfieldactive.md "Enable and disable form fields.").

The visibility of form elements can be changed by program with the [`ui.Form.setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.") or
[`ui.Form.setFieldHidden()`](../15_library-reference/3161-ui-form-setfieldhidden.md "Show or hide a form field.")
methods.

When using one of the `@screen-size` modifiers, the
`HIDDEN` attribute is only applied when the size of the screen or web browser window
matches. This can be used to implement responsive layout.

When using `HIDDEN@screen-size`, make sure that the element does not contain a
field (or is not a form field) that will be used for user input: The current input field having the
focus must always be visible.

## Example

```
EDIT f01 = FORMONLY.cust_id, HIDDEN;
EDIT f02 = FORMONLY.cust_name;
EDIT f03 = FORMONLY.cust_address, HIDDEN=USER;
EDIT f04 = FORMONLY.cust_comment, HIDDEN@SMALL, HIDDEN@MEDIUM;
```

## Related links

**Related concepts**  

[Responsive Layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.")

[Form field deactivation](2241-form-field-deactivation.md "Form field deactivation")
