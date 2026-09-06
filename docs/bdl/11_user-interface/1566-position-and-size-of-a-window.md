---
title: "Position and size of a window"
source: "fgl-topics/c_fgl_windows_and_forms_008.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Position and size of a window"
type: "concept"
---

# Position and size of a window

> Window objects can be created with a position and size for the TUI mode.

## Defining the window position and size in TUI mode

A typical [`OPEN
WINDOW`](1572-open-window.md "Creates and displays a new window.") instruction for TUI mode specifies the position on the screen:

```
OPEN WINDOW w1 AT 10,5 WITH FORM "custlist"
```

When using the [TUI mode](1517-text-mode-rendering-tui-mode.md) (or [traditional GUI mode](1519-graphical-mode-with-traditional-display.md)), the `AT line,
column` clause defines the position of the top-left corner of the window on the
terminal screen. The terminal character-cell coordinates start at line 1 and column 1, referencing
the top-left corner of the terminal.

The `WITH lines ROWS, characters COLUMNS` clause specifies explicit
vertical and horizontal dimensions for the window. The expression at the left of the
`ROWS` keyword specifies the height of the window, in character unit lines. This must
be an integer between 1 and max, where max is the maximum number of lines that the screen can
display. The integer expression after the comma at the left of the `COLUMNS` keyword
specifies the width of the window, in character unit columns. This must return a whole number
between 1 and length, where length is the number of characters that your monitor can display on one
line.

## Reserved window lines in TUI mode

In addition to the lines needed for a form, consider allowing room for the
`COMMENT` line, the `MENU` line, the `MENU` comment
line, and the `ERROR` line for the TUI mode.

The runtime system issues a runtime error if the window does not include sufficient lines to
display both the form and these additional reserved lines.

The minimum number of lines required to display a form in a window is the number of lines in the
form, plus an additional line below the form for prompts, messages, and comments.

## Window geometry in GUI mode

When using the full [GUI mode](1518-graphical-mode-rendering-gui-mode.md) (without the
traditional mode), the `AT line, column` clause is optional and if
used, the `WITH lines ROWS, characters COLUMNS` clause is ignored,
because the size of the window is automatically calculated, and the position is controlled by the
window manager.

The size of a [regular window](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window.") is defined by
the [window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), and the content of
the current form adapts to the size of the container. When using a [modal window](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window."), the size of the window is computed from
its content. The size of a modal window that shows a form defining a `TABLE`, is
computed from the size of the `TABLE` element (see [Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.") for more details)

## Related links

**Related concepts**  

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")

[Defining the position of reserved lines](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")
