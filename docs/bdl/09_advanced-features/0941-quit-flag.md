---
title: "quit_flag"
source: "fgl-topics/c_fgl_programs_QUIT_FLAG.html"
breadcrumb: "Advanced features > Program registers > quit_flag"
type: "concept"
---

# quit_flag

> quit_flag is a predefined variable set to TRUE when a quit signal is detected.

## Syntax

```
quit_flag
```

## Usage

`quit_flag` is set to `TRUE` when a quit signal is detected by the
runtime system. The quit signal is raised when the user presses the quit signal key
([Ctrl]+[Backslash]) in TUI mode, or when another process sends the SIGQUIT signal to the
fglrun process.

`quit_flag` must be used with the [`DEFER QUIT`](0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received.") configuration instruction.
If the `DEFER QUIT` instruction is not specified, and quit signal will stop the
program execution.

When the quit signal arrives during a procedural instruction (`FOR` loop) and
`DEFER QUIT` is used, the runtime system sets `quit_flag` to
`TRUE` and continues the program execution. It is up to the program to check the
`quit_flag` variable.

When the quit signal arrives during an interactive instruction (`INPUT`,
`CONSTRUCT`) and `DEFER QUIT` is used, the runtime system sets
`quit_flag` to `TRUE` and continues with the execution of the
interactive instruction.

Once `quit_flag` is set to `TRUE`, it must be reset to
`FALSE` to detect a new quit event.

## Example

```
MAIN
  DEFINE n INTEGER
  DEFER QUIT
  LET quit_flag = FALSE
  FOR n = 1 TO 1000
     IF quit_flag THEN EXIT FOR END IF
     ...
  END FOR
END MAIN
```

## Related links

**Related concepts**  

[User interruption handling](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure.")

[int\_flag](0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.")
