---
title: "CLEAR WINDOW"
source: "fgl-topics/c_fgl_windows_and_forms_CLEAR_WINDOW.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > CLEAR WINDOW"
type: "concept"
---

# CLEAR WINDOW

> Clears the contents of a window.

## Syntax

```
CLEAR WINDOW { identifier | SCREEN }
```

1. identifier is the name of the window.

## Usage

The `CLEAR WINDOW` instruction clears the content of the specified window that
was declared in an [`OPEN
WINDOW`](1572-open-window.md "Creates and displays a new window.").

In TUI mode, if the window was created with borders, these are left untouched (only the content
of the window is cleared).

If you specify `CLEAR WINDOW SCREEN`, the root screen will be cleared,
except areas occupied by an existing window. `CLEAR WINDOW SCREEN`
will not change the current window setting.

The `CLEAR WINDOW` instruction is typically used in text user interface (TUI)
based applications, as it clears the whole content of the window, including static labels and
messages.

## Related links

**Related concepts**  

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
