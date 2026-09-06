---
title: "The form concept"
source: "fgl-topics/c_fgl_windows_and_forms_forms.html"
breadcrumb: "User interface > Form definitions > Windows and forms > The form concept"
type: "concept"
---

# The form concept

> Forms define the layout and presentation of areas used by the dialogs, to display or input data.

## Loading forms in programs

Forms are loaded by programs from external files with the .42f extension,
the compiled version of .per [form
specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.").

Forms are loaded with the [`OPEN
FORM`](1578-open-form.md "Declares a compiled form in the program.") instruction followed by a `DISPLAY FORM`, to display the form
into the current window, or forms can be used directly as window creation argument with the [`OPEN WINDOW ... WITH FORM`](1574-with-form-clause.md "Creating a window object with a form.")
instruction:

```
OPEN FORM f_cust FROM "f_cust"
DISPLAY FORM f_cust -- into current window
...
OPEN WINDOW w_cust WITH FORM "f_cust"
```

## The version of a form

Forms can be stamped with the [`VERSION`](1846-version-attribute.md "The VERSION attribute is used to specify a user version string for an element.") attribute: This attribute is used to indicate that the form content
has changed.

The front-end is then able to distinguish different form versions and avoid saved settings being
applied for new form versions.

```
LAYOUT (VERSION="1.45")
...
```

## Controlling forms with dialogs

The form that is used by interactive instructions like `INPUT` is defined by the
[current window](1564-the-window-concept.md) containing the form.

Switching between existing windows (and thus, between forms associated with each window) is done
with the [`CURRENT
WINDOW`](1576-current-window.md "Makes a specified window the current window.") instruction.

Several forms can be successively displayed in the same (current) window. The last displayed form
will be used by the next dialog, while the form displayed before will be automatically removed from
the
window:

```
OPEN WINDOW w_common WITH 20 ROWS, 60 COLUMNS
...
OPEN FORM f1 FROM "f_cust"
DISPLAY FORM f1 -- f_cust is shown
INPUT BY NAME rec_cust.* ...
...
OPEN FORM f2 FROM "f_ord"
DISPLAY FORM f2 -- f_ord is shown (f_cust is removed)
INPUT BY NAME rec_ord.* ...
```

## API for form objects

The [`ui.Form`](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.") built-in class is provided
to manipulate form elements.

You can, for example, hide some parts of a form with the [`setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.") method.

A `ui.Form` object for the current form can be obtained with the [`ui.Window.getForm()`](../15_library-reference/3133-ui-window-getform.md "Get the current form of a window.") method or [`ui.Dialog.getForm()`](../15_library-reference/3200-ui-dialog-getform.md "Returns the current form used by the dialog."). Forms can also by
created dynamically with the [`ui.Window.createForm()`](../15_library-reference/3131-ui-window-createform.md "Create a new empty form in a window.") method, to generate forms at runtime.
