---
title: "DEFER INTERRUPT / QUIT"
source: "fgl-topics/c_fgl_programs_DEFER.html"
breadcrumb: "Advanced features > Configuration options > DEFER INTERRUPT / QUIT"
type: "concept"
---

# DEFER INTERRUPT / QUIT

> The DEFER instruction defines the program behavior when interrupt or quit signals are received.

## Syntax

```
DEFER { INTERRUPT | QUIT }
```

## Usage

The `DEFER` instruction controls the behavior of the program when an interrupt
signal (`SIGINT`) or quit signal (`SIGQUIT`) has been received.

`DEFER INTERRUPT` and `DEFER QUIT` instructions should
only be used in the [`MAIN`](0789-the-main-block-function.md "The MAIN block is the starting point of the program.")
block, to be executed at the beginning of the program.

`DEFER INTERRUPT` indicates that the program must continue when it receives an
interrupt signal (`SIGINT`). By default, the program stops when receiving an
interrupt signal.

Once deferred, you cannot reset to the default behavior.

When an interrupt signal is caught by the runtime system and `DEFER INTERRUPT` is
used, the [`int_flag`](0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") predefined
global variable is set to `TRUE` by the runtime system.

Interrupt signals are raised on terminal consoles when the user presses a key like CTRL-C,
depending on the stty configuration. When a program is displayed through a front-end, no terminal
console is used; therefore, users cannot send interrupt signals by pressing a key like CTRL-C. To
send an interruption request from the front-end, you must define an action view with the name
'`interrupt`'.

`DEFER QUIT` indicates that the program must continue when it receives a quit
signal (`SIGQUIT`). By default, the program stops when receiving a quit signal.

When a quit signal is caught by the runtime system and `DEFER QUIT` is used, the
[`quit_flag`](0941-quit-flag.md "quit_flag is a predefined variable set to TRUE when a quit signal is detected.") predefined global
variable is set to `TRUE` by the runtime system.

## Related links

**Related concepts**  

[User interruption handling](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure.")
