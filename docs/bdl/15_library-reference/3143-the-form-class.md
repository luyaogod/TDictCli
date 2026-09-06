---
title: "The Form class"
source: "fgl-topics/c_fgl_ClassForm.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class"
type: "concept"
---

# The Form class

> The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.

A form object allows you to manipulate form elements by program.
For example, you can hide parts of a form with the `setElementHidden()` method.
The runtime system is able to handle hidden fields during a
dialog instruction. You can, for example, hide a `GROUP` containing
fields and labels.

Outside dialogs, get a [`ui.Form`](3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.")
instance of the current form with the [`ui.Window.getForm()`](3133-ui-window-getform.md "Get the current form of a window.") method.
When executing a dialog, use the [`ui.Dialog.getForm()`](3200-ui-dialog-getform.md "Returns the current form used by the dialog.")
method.

Note that the [`OPEN
FORM`](../11_user-interface/1578-open-form.md "Declares a compiled form in the program.") instruction does not load a form; it
simply declares a handle. The form will be created in the AUI
tree when executing the [`DISPLAY FORM`](../11_user-interface/1579-display-form.md "Displays and associates a form with the current window.")
instruction. Therefore, the corresponding `ui.Form` object
is only available after `DISPLAY FORM` is executed.

## Child topics

- [ui.Form methods](3144-ui-form-methods.md): Methods of the ui.Form class.
- [Usage](3163-usage.md)
- [Examples](3165-examples.md): ui.Form usage examples.
