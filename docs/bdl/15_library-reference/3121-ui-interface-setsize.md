---
title: "ui.Interface.setSize"
source: "fgl-topics/c_fgl_ClassInterface_setSize.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.setSize"
type: "concept"
---

# ui.Interface.setSize

> Specify the initial size of the window container.

## Syntax

```
ui.Interface.setSize(
   h STRING,
   w STRING )
```

1. h is the height of the window container.
2. w is the width of the window container.
3. Default unit is form cells, can accept graphical units like `px`, or
   `%` for screen ratio.

## Usage

Call the `ui.Interface.setSize()` method at the beginning of the program, before
displaying any window/form and before starting any interactive instruction.

The parameters can be `INTEGER` or `STRING` values. By default, the
unit is the character grid cells, to be used with [traditional
mode](../11_user-interface/1519-graphical-mode-with-traditional-display.md). For regular rendering, use the `px` unit to specify the height and width
in pixels, or use the `%` sign as unit, to represent a ratio of the desktop
screen.

When using the GDC front end with regular rendering mode, the
`ui.Interface.setSize()` method defines the initial size of the [window container](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."). This method has
no effect when using the GAS+browser or with a mobile device, since the web browser tab or device
screen size cannot be controlled. When [stored
settings](../11_user-interface/1560-front-end-stored-settings.md "Front-ends can store some layout properties of windows and form elements, for subsequent program executions.") apply, the size is taken from the stored settings.

When using [traditional mode](../11_user-interface/1519-graphical-mode-with-traditional-display.md) for the rendering, the
`ui.Interface.setSize()` method defines the initial size of the area containing the
fixed windows created by `OPEN WINDOW`, as a replacement for the [`fgl_setsize()`](2779-fgl-setsize.md "Sets the size of the main application window.") built-in
function.

## Example

```
MAIN
    CALL ui.Interface.setSize("30%","80%")
    OPEN FORM f1 FROM "main_form"
    DISPLAY FORM f1
    MENU "test"
    COMMAND "quit" EXIT MENU
    END MENU
END MAIN
```
