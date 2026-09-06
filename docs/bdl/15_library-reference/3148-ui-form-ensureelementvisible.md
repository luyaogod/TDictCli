---
title: "ui.Form.ensureElementVisible"
source: "fgl-topics/c_fgl_ClassForm_ensureElementVisible.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.ensureElementVisible"
type: "concept"
---

# ui.Form.ensureElementVisible

> Ensure the visibility of a form element.

## Syntax

```
ensureElementVisible(
   name STRING )
```

1. name defines the name of the form element, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).

## Usage

Use the `ensureElementVisible()` method, to show the given form element (not form
field) to the user. This method is typically used to show a folder page containing a static
image.

The `ensureElementVisible()` method must only be used for static form elements. To
make form fields temporarly visible without giving the focus to the field, use the
`ensureFieldVisible()` method instead.

The form element is identified by its name. If several form elements have the same name, the
first form element found is selected.

The `ensureElementVisible()` method can only show the specified element, if the
focus handling in the current active dialog allows it. For more details, see the [`ensureFieldVisible()`](3149-ui-form-ensurefieldvisible.md "Ensure visibility of a form field.")
instead.

## Related links

**Related concepts**  

[Giving the focus to a form element](../11_user-interface/2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code.")

[PAGE item type](../11_user-interface/1697-page-item-type.md "Defines the content of a folder page.")
