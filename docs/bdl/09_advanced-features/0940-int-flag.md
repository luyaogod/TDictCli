---
title: "int_flag"
source: "fgl-topics/c_fgl_programs_INT_FLAG.html"
breadcrumb: "Advanced features > Program registers > int_flag"
type: "concept"
---

# int_flag

> int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.

## Syntax

```
int_flag
```

## Usage

The `int_flag` register is set to `TRUE` by the runtime system,
when an interruption signal is sent to the fglrun process, when an asynchronous
interruption event is detected in GUI mode, or when the automatic "cancel" (or "close") action is
fired in a [singular dialog](../11_user-interface/2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.")
(`INPUT`, `CONSTRUCT`).

`int_flag` must be used with the [`DEFER INTERRUPT`](0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received.") configuration instruction: If the `DEFER
INTERRUPT` instruction is not specified, and interruption signal will stop the program
execution.

With singular dialogs, `int_flag` is typically tested after the dialog block (or
in the `AFTER INPUT` block for example), to detect if the user has validated or
canceled the dialog. Note that `int_flag` is not reset to `FALSE` when
the automatic "accept" action is fired. Therefore, it is good practice to initialize
`int_flag` to `FALSE`, before starting a singular dialog.

With a [multiple dialog](../11_user-interface/2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling."),
`int_flag` is not used/set: No automatic "accept" or "cancel" action is created. To
leave the dialog, implement your own (user-defined) action handler, such as `ON ACTION
close`.

The `int_flag` register is also used by the runtime system as diagnostic flag for
predefined action block execution such as [`ON
INSERT`](../11_user-interface/1985-on-insert-block.md) in `DISPLAY ARRAY`.

In TUI mode, the interrupt signal occurs when the user presses the interrupt signal key
(typically, CTRL-C).

In GUI mode, the front-end may send an asynchronous interruption event, while the program is
running in a procedure or SQL query.

When an interruption signal or event occurs during a procedural instruction (`FOR`
loop), the runtime system sets `int_flag` to `TRUE`. It is up to the
program to check the `int_flag` variable and leave the loop.

Once `int_flag` is set to `TRUE`, it must be reset to
`FALSE` in order to detect a new interruption event.

## Example

```
MAIN
  DEFER INTERRUPT
  LET int_flag = FALSE
  INPUT BY NAME ...
     AFTER INPUT
        IF int_flag THEN
           MESSAGE "The input is canceled."
        END IF
     ...
  END INPUT
  ...
END MAIN
```

## Related links

**Related concepts**  

[User interruption handling](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure.")

[quit\_flag](0941-quit-flag.md "quit_flag is a predefined variable set to TRUE when a quit signal is detected.")

[Implementing the close action](../11_user-interface/2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button).")
