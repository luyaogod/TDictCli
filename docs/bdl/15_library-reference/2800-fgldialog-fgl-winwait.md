---
title: "fgldialog.fgl_winwait()"
source: "fgl-topics/c_fgl_utility_functions_FGL_WINWAIT.html"
breadcrumb: "Library reference > Utility modules > fgldialog: Common dialog functions > fgldialog.fgl_winwait()"
type: "concept"
---

# fgldialog.fgl_winwait()

> Displays an interactive message box and waits for user validation.

## Syntax

```
FUNCTION fgl_winwait(
   text STRING )
```

1. text is the message displayed in the message box.
   Use '`\n`' to separate lines (not working on ASCII client).

## Usage

The `fgl_winwait()` function displays a message to the end user and waits until
the user presses the OK button.

> **Important:**
>
> With front-ends implementing this function with a system dialog box API
> creating a modal window, the end user will have to close the modal window first, before continuing
> within the window of another program. Consider using a [menu with
> "dialog" style](../11_user-interface/1906-syntax-of-the-menu-instruction.md "The MENU instruction defines a set of options the end user can select to trigger actions in a program.") instead so as not to block other programs.
