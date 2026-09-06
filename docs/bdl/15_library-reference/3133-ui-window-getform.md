---
title: "ui.Window.getForm"
source: "fgl-topics/c_fgl_ClassWindow_getForm.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.getForm"
type: "concept"
---

# ui.Window.getForm

> Get the current form of a window.

## Syntax

```
getForm()
  RETURNS ui.Form
```

## Usage

The `getForm()` method
returns the `ui.Form` object corresponding to the
current form used by the window object.

Declare a variable
of type `ui.Form` to hold the form object reference.

Consider
using the `ui.Dialog.getForm()` method to get the form
used by the current dialog.

## Example

```
DEFINE f ui.Form
OPEN WINDOW w1 WITH FORM "custform"
LET w = ui.Window.getCurrent()
LET f = w.getForm()
...
```

## Related links

**Related concepts**  

[Example 2: Get a the current form and hide a groupbox](3142-example-2-get-a-the-current-form-and-hide-a-groupbox.md "Example 2: Get a the current form and hide a groupbox")

[The Form class](3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.")

[ui.Dialog.getForm](3200-ui-dialog-getform.md "Returns the current form used by the dialog.")
