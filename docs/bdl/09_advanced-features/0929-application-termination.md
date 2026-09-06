---
title: "Application termination"
source: "fgl-topics/c_fgl_programs_009.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Application termination"
type: "concept"
---

# Application termination

> The OPTIONS TERMINATE SIGNAL defines a callback function in case of SIGTERM signal.

## Syntax

> **Important:**
>
> This feature is only supported on UNIX-like operating systems.

```
OPTIONS ON TERMINATE SIGNAL CALL function
```

1. function is the name of a function with no parameters or return value.

## Usage

The `OPTIONS ON TERMINATE SIGNAL CALL function` defines
the function that must be called when the application receives the SIGTERM signal.

With this option, you can control program termination. If this statement is not called, the
program is stopped with an exit value of SIGTERM (15).

When the SIGTERM signal occurs and `OPTIONS ON TERMINATE CALL` is used, the
callback function is executed, and the program continues. To stop the program, perform an
`EXIT PROGRAM` in the callback function, or set the int\_flag and test this flag in
the processing code, to stop properly.

Use the `OPTIONS ON TERMINATE SIGNAL CALL function` instruction with
care, and do not execute complex code in the callback function. The code is expected to contain only simple and short cleanup operations; any interactive
instruction is avoided.

## Example

```
MAIN
    OPTIONS ON TERMINATE SIGNAL CALL end_app
    MENU "test"
        COMMAND "quit" EXIT MENU
    END MENU
END MAIN

FUNCTION end_app()
    DISPLAY "Program terminated by SIGTERM signal..."
    EXIT PROGRAM 1
END FUNCTION
```
