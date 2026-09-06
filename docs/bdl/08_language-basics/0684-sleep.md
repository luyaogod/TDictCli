---
title: "SLEEP"
source: "fgl-topics/c_fgl_FlowControl_SLEEP.html"
breadcrumb: "Language basics > Flow control > SLEEP"
type: "concept"
---

# SLEEP

> The SLEEP instruction causes the program to pause for the specified number of seconds.

## Syntax

```
SLEEP seconds
```

1. seconds must be an integer expression.

## Usage

The `SLEEP` instruction is typically invoked to let the end user read a message
displayed on a character terminal.

For example:

```
MAIN
  DISPLAY "Please wait 5 seconds..."
  SLEEP 5
  DISPLAY "Thank you."
END MAIN
```

With graphical applications, the `SLEEP` command is seldom used and should be
avoided.

When seconds is lower than zero or is null, the program continues immediately
with the next statement.

The `SLEEP` instruction can be interrupted by a process signal like SIGINT. If
your program uses [`DEFER INTERRUPT`](../09_advanced-features/0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received.") to
prevent program termination on SIGINT, the `SLEEP` instruction will return
immediately when the SIGINT signal is caught by the process. The program will then continue after
`SLEEP`, without waiting for the number of seconds specified.

To prevent interruption of the `SLEEP` command, put a `SLEEP 1`
(second) instruction in a `WHILE` loop and continue until the requested time has
expired, as shown in the following
example:

```
MAIN
    DEFER INTERRUPT
    CALL mysleep(3)
END MAIN

FUNCTION mysleep(secs SMALLINT)
    DEFINE te DATETIME YEAR TO FRACTION
    IF secs <= 0 OR secs IS NULL THEN
        RETURN
    END IF
    LET te = CURRENT + secs UNITS SECOND
    WHILE secs > 0
        SLEEP secs
        IF CURRENT >= te THEN
            EXIT WHILE
        END IF
        LET secs = ((te - CURRENT) / 1 UNITS SECOND) + 0.5
    END WHILE
END FUNCTION
```

## Related links

**Related concepts**  

[Integer expressions](0595-integer-expressions.md "This section covers integer expression evaluation rules.")
