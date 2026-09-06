---
title: "ui.Form.ensureFieldVisible"
source: "fgl-topics/c_fgl_ClassForm_ensureFieldVisible.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.ensureFieldVisible"
type: "concept"
---

# ui.Form.ensureFieldVisible

> Ensure visibility of a form field.

## Syntax

```
ensureFieldVisible(
   name STRING )
```

1. name defines the name of the form field, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).

## Usage

The `ensureFieldVisible()` method makes the given form field
visible to the user. This method can for example be used to show a folder
page by passing a field that is located in the folder page, even if the field
is not used in a dialog.

The form field is identified by name, with an optional prefix
(`table.column` or `column`).

This method does not give the focus to the field passed as parameter: The folder page or screen
area shown by this method call is temporarily visible and can disappear at the next user
interaction, depending on focus management.

For example, consider a folder having two pages. The focus is in a field on
the first page. A call to the `ensureFieldVisible()` method
makes the second folder page visible, passing a field located in the second page.
When the user presses the TAB key, the focus goes to the next field on the
first page, bringing the first page to the top. If you want to show a folder
page and give the focus to a specific field in that page, you must explicitly
give the focus to a field of the page, with [`NEXT FIELD`](../11_user-interface/1955-next-field-instruction.md).

The `ensureFieldVisible()` method is used for form fields. To show static form
elements such as labels or images, use the [`ensureElementVisible()`](3148-ui-form-ensureelementvisible.md "Ensure the visibility of a form element.") method instead.
