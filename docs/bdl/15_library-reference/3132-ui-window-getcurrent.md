---
title: "ui.Window.getCurrent"
source: "fgl-topics/c_fgl_ClassWindow_getCurrent.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.getCurrent"
type: "concept"
---

# ui.Window.getCurrent

> Get the current window object.

## Syntax

```
ui.Window.getCurrent()
  RETURNS ui.Window
```

## Usage

The `ui.Window.getCurrent()` class
method returns the `ui.Window` object corresponding
to the current window.

Declare a variable of type `ui.Window` to
hold the window object reference.

When invoking the `ui.Window.getCurrent()` method to get the initial default
"screen" window, an `OPEN FORM`/`DISPLAY FORM` must have been
executed. Otherwise, the method will return `NULL`.

When no form is yet displayed, and you want to get the default "screen" window object before any
dialog instruction, use the `CURRENT WINDOW IS screen` instruction, then use
`ui.Window.getCurrent()`. When the first dialog is a `MENU`
instruction and no form is yet displayed, the `ui.Window.getCurrent()` method can be
called in the `BEFORE MENU` block. See [`ui.Window.setText()`](3139-ui-window-settext.md "Set the window title.") for an code example. See also [`ui.Window.createForm()`](3131-ui-window-createform.md "Create a new empty form in a window."), to create a
form dynamically in the initial "screen" window.

## Example

```
DEFINE w ui.Window
OPEN WINDOW w1 WITH FORM "custform"
LET w = ui.Window.getCurrent()
...
```

## Related links

**Related concepts**  

[Example 2: Get a the current form and hide a groupbox](3142-example-2-get-a-the-current-form-and-hide-a-groupbox.md "Example 2: Get a the current form and hide a groupbox")

[ui.Window.forName](3130-ui-window-forname.md "Get a window object by name.")
