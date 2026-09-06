---
title: "ui.Dialog.getForm"
source: "fgl-topics/c_fgl_ClassDialog_getForm.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getForm"
type: "concept"
---

# ui.Dialog.getForm

> Returns the current form used by the dialog.

## Syntax

```
getForm()
  RETURNS ui.Form
```

## Usage

The `getForm()` method returns a `ui.Form` object as a handle to
the current form used by the dialog.

Use
this form object to modify elements of the current form. For example,
you can hide some parts of the form with the [`ui.Form.setElementHidden()`](3156-ui-form-setelementhidden.md "Show or hide form elements.") method.
