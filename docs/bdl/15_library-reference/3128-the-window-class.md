---
title: "The Window class"
source: "fgl-topics/c_fgl_ClassWindow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class"
type: "concept"
---

# The Window class

> The ui.Window class provides an interface to the window objects created with the OPEN WINDOW instruction.

A window object is created with a form with the `OPEN WINDOW WITH FORM`
instruction, and can be manipulated with the `ui.Window` methods, after getting the
window object instance with [`ui.Window.getCurrent()`](3132-ui-window-getcurrent.md "Get the current window object.") or [`ui.Window.forName()`](3130-ui-window-forname.md "Get a window object by name.").

When the window contains a form, you can use the [`ui.Form`](3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.") class to manipule form element.

## Child topics

- [ui.Window methods](3129-ui-window-methods.md): Methods of the ui.Window class.
- [Examples](3140-examples.md): ui.Window usage examples.
