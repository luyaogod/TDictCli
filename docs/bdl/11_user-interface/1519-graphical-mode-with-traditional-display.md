---
title: "Graphical mode with Traditional Display"
source: "fgl-topics/c_fgl_DynamicUI_005.html"
breadcrumb: "User interface > User interface basics > Genero user interface modes > Graphical mode rendering (GUI mode) > Graphical mode with Traditional Display"
type: "concept"
description: "Important: Traditional mode is under construction and not fully supported. What is the Traditional GUI mode? The Traditional GUI mode can be used to ease migration from TUI based applications to ..."
---

# Graphical mode with Traditional Display

> **Important:**
>
> Traditional mode is under construction and not fully supported.

## What is the Traditional GUI mode?

The Traditional GUI mode can be used to ease migration from [TUI based applications](1517-text-mode-rendering-tui-mode.md) to [graphical mode](1518-graphical-mode-rendering-gui-mode.md).

With the Traditional GUI mode, application windows bound to forms using a [`SCREEN`](1714-screen-section.md "The SCREEN section defines the form layout for TUI mode forms.") section will be
displayed as simple boxes in a main front-end window. Other windows bound to forms defined with the
[`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section will be
displayed as new GUI windows.

![Traditional GUI mode rendering screenshot](../_images/TradiMod1_gbc.jpg)

*Traditional GUI mode example*

Here the same form displayed with default GUI rendering (no traditional mode):

![Default GUI mode rendering screenshot](../_images/TradiMod1_gbc_2.jpg)

*Default GUI mode example*

## Enabling the Traditional GUI mode

The Traditional GUI mode can be enabled with the following [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") entry:

```
gui.uiMode = "traditional"
```

By default, the Traditional GUI mode is off.

## Window rendering rules

If the Traditional GUI mode is enabled, the [`OPEN WINDOW`](1572-open-window.md "Creates and displays a new window.") statement works differently depending on the layout type of
bound forms.

On the front-end side, there is one unique main graphical window (a top-level widget called
"compatibility window container") created to host all the windows created by a program.
Traditional forms are form files which have a `SCREEN` section instead of the
`LAYOUT` section. When migrating from a TUI mode project, all forms initially contain
a `SCREEN` section; hence all windows opened in traditional mode will appear in the
compatibility window container.

To rebuild a form file with graphical items such as group boxes, buttons
and tables, use a `LAYOUT` section. If the rebuilt form file
is loaded via `OPEN WINDOW ... WITH FORM form-file` then,
even in traditional mode, the newly created window will appear as a new
top-level widget on the front-end side. This opens a smooth migration path
using the traditional mode; as a first step, it is possible to migrate and
enhance some application forms like typical search lists, while keeping the
rest of the application forms running in the traditional rendering.

Note, however, that following instructions do not work in Traditional GUI mode:

1. `OPEN WINDOW window_id AT line, column WITH height ROWS, width COLUMNS`
2. `OPEN FORM form_id FROM "form_file"`

   (where *form\_file* is defined with a `LAYOUT` section)
3. `DISPLAY FORM form_id`

A runtime error results, because you cannot display a form with dynamic
geometry in a fixed geometry container.
Only forms with a `SCREEN` section can be displayed at a
later stage in a window that was initially opened inside the compatibility
window container.

## Function key shifting

When the Traditional GUI mode is enabled, you can map Shift-Fx and Ctrl-Fx key strokes to
F(x+offset) actions. The offset is defined with the `gui.key.add_function` entry:

```
gui.key.add_function = 12
```

This entry defines the number of function keys of the keyboard (default
is 12). When defined as 12, a Shift-F1 will be received as an F13 (12+1)
action event by the program, and a Control-F1 will be F25 (12\*2+1).

## Related links

**Related concepts**  

[Graphical mode rendering (GUI mode)](1518-graphical-mode-rendering-gui-mode.md "Graphical mode rendering (GUI mode)")
