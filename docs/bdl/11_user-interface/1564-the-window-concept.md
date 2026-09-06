---
title: "The window concept"
source: "fgl-topics/c_fgl_windows_and_forms_windows.html"
breadcrumb: "User interface > Form definitions > Windows and forms > The window concept"
type: "concept"
---

# The window concept

> Windows are containers for .42f forms.

## Creating windows

The windows are created from programs; they define a display context for sub-elements like forms,
menus, message and error lines.

A window can contain only one form at a time, but you can display different forms successively in
the same window.

A program creates a new window with the [`OPEN WINDOW`](1572-open-window.md "Creates and displays a new window.") instruction, which also defines the window
identifier:

```
OPEN WINDOW mywindow WITH FORM "myform"
```

## Destroying windows

A window is destroyed with the `CLOSE WINDOW`
instruction:

```
CLOSE WINDOW mywindow
```

## API for window objects

The [`ui.Window`](../15_library-reference/3128-the-window-class.md "The ui.Window class provides an interface to the window objects created with the OPEN WINDOW instruction.") built-in class can be
used to manipulate windows as objects.

The common practice is to get the current window with [`ui.Window.getCurrent()`](../15_library-reference/3132-ui-window-getcurrent.md "Get the current window object."), get the
current displayed form with the [`getForm()`](../15_library-reference/3133-ui-window-getform.md "Get the current form of a window.") window object method, and use it as [`ui.Form`](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.") object to manipulate its content.

## Windows rendering context

When using the [text mode](1517-text-mode-rendering-tui-mode.md) (FGLGUI=0),
`WINDOW` objects are displayed in the character terminal as fixed-size boxes, at a
given line/column position, width and height.

In [graphical mode](1518-graphical-mode-rendering-gui-mode.md), all `WINDOW`
objects are displayed in a single resizable desktop window container or web browser frame (the
window container). The rendering of the window container can be customized. For more
details, see [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").

A GUI application can run in [traditional mode](1519-graphical-mode-with-traditional-display.md)
(`gui.uiMode="traditional"` FGLPROFILE setting), displaying windows as fixed static
frames inside the real graphical parent window.

## Default SCREEN window

When a program starts, the runtime system creates a default window named
`SCREEN`.

This default window can be used as a regular window: it can hold a menu and a form.

Programs typically display their main form in the `SCREEN` window, by using [`OPEN FORM`](1578-open-form.md "Declares a compiled form in the program.") + [`DISPLAY
FORM`](1579-display-form.md "Displays and associates a form with the current window."):

```
MAIN
    -- The SCREEN window exists by default
    ...
    OPEN FORM f_main FROM "customers"
    DISPLAY FORM f_main -- displays in SCREEN
    ...
END MAIN
```

If needed, the default `SCREEN` window can be closed with [`CLOSE WINDOW SCREEN`](1575-close-window.md "Closes and destroys a window.").
However, in most cases, you want to keep this default window and display the main form of the
program with `OPEN FORM` + `DISPLAY FORM`.

Before invoking methods such as `ui.Window.getCurrent()` to get the default screen
window object, and no other user interface instruction was used for this window, execute
`CURRENT WINDOW IS SCREEN` before calling the `ui.Window` methods.

## The current window

A program with user interface must always have a current window.

Several windows can be created, but there can be only one current window when using modal dialogs
(only one dialog is active at the time, thus only the current window can be active).

There is always a current window. The last created window becomes the current window. When the
last created window is closed, the previous window in the window stack becomes the current
window.

Use the [`CURRENT
WINDOW`](1576-current-window.md "Makes a specified window the current window.") instruction to make a specific window current, before executing the
corresponding dialog that is controlling the window
content:

```
OPEN WINDOW w_customers ...
OPEN WINDOW w_orders ...
...
CURRRENT WINDOW IS w_customers
...
CLOSE WINDOW w_customers
CURRRENT WINDOW IS w_orders
...
```

However, this practice is not commonly used: A regular Genero program starts with a main
window/form and opens new windows in cascade that results in a tree of windows. The last created
window is closed to go back to the previous window/form.

## Displaying multiple forms in the same window

When there is a current window, it is possible to display several forms successively in that same
window.

The previous form is removed automatically by the runtime system when displaying a new form to
the
window:

```
OPEN WINDOW mywindow WITH FORM "form1"
INPUT BY NAME ... -- uses form1 elements
...
OPEN FORM f1 FROM "form2"
DISPLAY FORM f1  -- removes "form1" from the window
INPUT BY NAME ... -- uses form2 elements
...
```

## Window presentation styles

Window decoration and behavior options can be defined with a [presentation style](1654-window-style-attributes.md "Window presentation style attributes apply to a window element.").

The window style is identified with the `STYLE` attribute of the [`ATTRIBUTES`
section](1573-open-window-attributes.md "List of attributes for the OPEN WINDOW instruction.") of `OPEN WINDOW`, or it can also be specified at form level, with the
[`WINDOWSTYLE`](1851-windowstyle-attribute.md "The WINDOWSTYLE attribute defines the style to be used by the parent window of a form.") form attribute
in the `LAYOUT` of the form
definition:

```
OPEN WINDOW w_cust WITH FORM "f_cust" ATTRIBUTES(STYLE="common")
```

## Normal and modal windows

If a window uses the style attribute `windowType=modal`, it will appear in a modal
frame on top of the last displayed frame of a window using the default style attribute
`windowType=normal`.

To create a normal window, use the default window style defined with the
`windowType=normal` style attribute (here we do not specify any
`STYLE` attribute in `OPEN
WINDOW`):

```
OPEN WINDOW w_cust WITH FORM "f_cust"
```

To create a modal window, use one of the styles defined with the
`windowType=modal` style attribute, such as the `"dialog"`
style:

```
OPEN WINDOW w_print_options WITH FORM "f_print_options" ATTRIBUTES(STYLE="dialog")
```

For more details about [Configuring windows with styles](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window.").
