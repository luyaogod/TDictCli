---
title: "ui.Window.createForm"
source: "fgl-topics/c_fgl_ClassWindow_createForm.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.createForm"
type: "concept"
---

# ui.Window.createForm

> Create a new empty form in a window.

## Syntax

```
createForm(
   name STRING )
  RETURNS ui.Form
```

1. name is the name for the form.

## Usage

The `createForm()` method can be used to create a new
empty form in the window object. This is typically used to build forms dynamically to be controlled
with dynamic dialogs.

After creating the form element with createForm(), the program can build
the form elements with the [om API](3280-the-om-package.md "These topics cover the built-in classes of the om package").

It is mandatory to
create a form in a window with the `createForm()` method. A form element created
directly with the om API is not usable.

If the form passed as parameter does not exist, the
`createForm()` method returns a new `ui.Form` instance. Otherwise,
when the form name passed as the parameter identifies an existing form used by the window, the
method returns `NULL`.

The method must be invoked by using the current window,
typically created by an `OPEN WINDOW` instruction
like:

```
OPEN WINDOW w1 WITH 10 ROWS, 30 COLUMNS
```

To create a form
dynamically as the very initial form of the program, and no user interface instruction what
performed before, execute the [`CURRENT WINDOW IS SCREEN`](../11_user-interface/1576-current-window.md "Makes a specified window the current window.") instruction, to get the current window node, and
invoke the `createForm()` method:

```
CURRENT WINDOW IS SCREEN
LET f = ui.Window.getCurrent().createForm("myform")
```

## Example

```
DEFINE w ui.Window,
       f ui.Form,
       n, g om.DomNode
OPEN WINDOW w1 WITH 10 ROWS, 20 COLUMNS
LET w = ui.Window.getCurrent()
LET f = w.createForm("myform")
LET n = f.getNode()
LET g = n.createChild("Grid")
...
```

## Related links

**Related concepts**  

[Create forms dynamically](../11_user-interface/2427-create-forms-dynamically.md "Dynamic dialogs are typically used with forms that are generated at runtime.")

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[ui.Window.forName](3130-ui-window-forname.md "Get a window object by name.")

[ui.Window.getCurrent](3132-ui-window-getcurrent.md "Get the current window object.")
