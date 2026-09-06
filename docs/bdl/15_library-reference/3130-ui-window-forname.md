---
title: "ui.Window.forName"
source: "fgl-topics/c_fgl_ClassWindow_forName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.forName"
type: "concept"
---

# ui.Window.forName

> Get a window object by name.

## Syntax

```
ui.Window.forName(
   name STRING )
  RETURNS ui.Window
```

1. name defines the name of the window.

## Usage

The `ui.Window.forName()` class method returns the `ui.Window`
object corresponding to an identifier used to create the window with the [`OPEN WINDOW`](../11_user-interface/1572-open-window.md "Creates and displays a new window.") instruction, or
the predefined "screen" window.

Declare a variable of type `ui.Window` to hold the window object reference.

The name of the window passed as parameter can use the same letter case as in the `OPEN
WINDOW` instruction: The lookup is case-insensitive.

When invoking the `ui.Window.forName()` method to get the initial default "screen"
window, an `OPEN FORM`/`DISPLAY FORM` must have been executed.
Otherwise, the method will return `NULL`. To create a form dynamically in the initial
"screen" window, see [ui.Window.createForm](3131-ui-window-createform.md "Create a new empty form in a window.").

## Example

```
DEFINE w ui.Window
OPEN WINDOW w1 WITH FORM "custform"
LET w = ui.Window.forName("w1")
...
```

## Related links

**Related concepts**  

[Example 1: Get a window by name and change the title](3141-example-1-get-a-window-by-name-and-change-the-title.md "Example 1: Get a window by name and change the title")

[ui.Window.getCurrent](3132-ui-window-getcurrent.md "Get the current window object.")
