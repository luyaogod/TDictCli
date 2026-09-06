---
title: "CLOSE WINDOW"
source: "fgl-topics/c_fgl_windows_and_forms_CLOSE_WINDOW.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > CLOSE WINDOW"
type: "concept"
---

# CLOSE WINDOW

> Closes and destroys a window.

## Syntax

```
CLOSE WINDOW { identifier | SCREEN }
```

1. identifier is the name of the window.

## Usage

The `CLOSE WINDOW` closes the specified window. If the [`OPEN WINDOW`](1572-open-window.md "Creates and displays a new window.") statement
includes the `WITH FORM` clause,
it closes both the form and the window.

Closing a window has no effect on
variables that were set while the window was open.

Closing
the current window makes the next window on the stack the new current
window. If you close any other window, the runtime system deletes
it from the stack, leaving the current window unchanged.

If the window is currently being used for input, `CLOSE
WINDOW` generates a runtime error.

You can close
the default screen window with the `CLOSE WINDOW SCREEN` instruction.

## Example

```
MAIN
  OPEN WINDOW w1 WITH FORM "customer"
  MENU "Test"
    COMMAND KEY(INTERRUPT) "exit" EXIT MENU
  END MENU
  CLOSE WINDOW w1
END MAIN
```

## Related links

**Related concepts**  

[OPEN WINDOW](1572-open-window.md "Creates and displays a new window.")
