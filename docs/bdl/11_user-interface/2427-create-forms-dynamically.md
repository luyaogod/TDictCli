---
title: "Create forms dynamically"
source: "fgl-topics/c_fgl_dynamic_dialogs_forms.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Create forms dynamically"
type: "concept"
---

# Create forms dynamically

> Dynamic dialogs are typically used with forms that are generated at runtime.

## Purpose of dynamic forms

Before you instantiate a new `ui.Dialog` object, you must load an existing
compiled .42f form, or create a new form dynamically in your program.

Since dynamic dialogs are build at runtime (because form fields are not known at compiled time),
the corresponding form is also created (or completed) at runtime.

The current form (in the current window) is automatically attached to the new created dialog.

## Dynamic form creation with `createForm()`

Forms built at runtime must be created with the [`ui.Window.createForm()`](../15_library-reference/3131-ui-window-createform.md "Create a new empty form in a window.") method, and must contain a valid definition with
layout containers, form fields, and screen records.

The `createForm()` method must be invoked by using the current window. The window
is typically created with `OPEN WINDOW` such as:

```
OPEN WINDOW w1 WITH 10 ROWS, 30 COLUMNS
```

Then create the `ui.Form` object, and get the `om.DomNode` object
to build your form:

```
DEFINE f ui.Form, n om.DomNode
LET f = ui.Window.getCurrent().createForm("myform")
LET n = f.getNode()
... create the AUI elements of the form
```

To create the very initial application form dynamically, when no other form was displayed in the
default screen window, execute the `CURRENT WINDOW IS SCREEN` instruction before
using the `ui.Window.getCurrent()` method and create the form. Otherwise,
`getCurrent()` would return
`NULL`:

```
CURRENT WINDOW IS SCREEN
LET f = ui.Window.getCurrent().createForm("myform")
```

Use `om` classes, to build your form dynamically. A good practice in creating
dynamic forms is to write first a .per file, that implements a static version
of one of the forms you want to build at runtime. Compile the .per to a
.42f and inspect the generated XML file, to understand the structure of the
form file.

## COMBOBOX initializers

[Combobox initializers](2250-filling-a-combobox-item-list.md "The item list of COMBOBOX fields can be initialized at runtime.") are called when
executing the `DISPLAY FORM` or `OPEN WINDOW WITH FORM`
instructions.

When creating a form dynamically with `COMBOBOX` fields, the initialization
functions are not called. The combobox items must be created as part of the combobox node.

## Related links

**Related concepts**  

[The om package](../15_library-reference/3280-the-om-package.md "These topics cover the built-in classes of the om package")
