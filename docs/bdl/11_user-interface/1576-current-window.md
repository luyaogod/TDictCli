---
title: "CURRENT WINDOW"
source: "fgl-topics/c_fgl_windows_and_forms_CURRENT_WINDOW.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > CURRENT WINDOW"
type: "concept"
---

# CURRENT WINDOW

> Makes a specified window the current window.

## Syntax

```
CURRENT WINDOW IS { identifier | SCREEN }
```

1. identifier is the name of the window.

## Usage

Programs with multiple windows might need to switch to a different
open window so that input and output occur in the appropriate window. To make a window the
current window, use the `CURRENT WINDOW` statement.

When a program starts, the screen is the current window. Its name is `SCREEN`.
To make this the current window, specify the keyword `SCREEN` instead of a window
identifier.

If
the window contains a form, that form becomes the current form when a `CURRENT
WINDOW` statement specifies the name of that window. All interactive instruction
such as `CONSTRUCT`, `INPUT` use only the current window for
input and output. If the user displays another form (for example, through an `ON
KEY` clause) in one of these statements, the window containing the new form
becomes the current window. When an interactive instruction resumes, its original window
becomes the current window.

The `CURRENT WINDOW` instruction is typically used in text user interface (TUI)
based applications, when distinct areas of the screen are reserved for different usage. In a GUI
application, windows are typically opened and closed sequentially or on a stack of windows.

## Example

```
MAIN
  OPEN WINDOW w1 WITH FORM "customer"
  ...
  OPEN WINDOW w2 WITH FORM "custlist"
  ...
  CURRENT WINDOW IS w1
  ...
  CURRENT WINDOW IS w2
  ...
  CLOSE WINDOW w1
  CLOSE WINDOW w2
END MAIN
```

## Related links

**Related concepts**  

[OPEN WINDOW](1572-open-window.md "Creates and displays a new window.")

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
