---
title: "disable"
source: "fgl-topics/c_fgl_Debugger_disable.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > disable"
type: "concept"
---

# disable

> The disable command disables the specified breakpoint.

## Syntax

```
disable breakpoint
```

1. breakpoint is the number assigned to the breakpoint by the debugger.

## Usage

The `disable` command instructs the debugger to ignore the specified
breakpoint when running the program.

Use the [`enable`](2598-enable.md "The enable command enables breakpoints that have previously been disabled.") command
to reactivate the breakpoint for the current debugger session.

## Example

```
MAIN
    CALL func1()
END MAIN

FUNCTION func1() RETURNS ()
    DEFINE x INT, s STRING
    CALL func2() RETURNING x, s
END FUNCTION

FUNCTION func2() RETURNS (INT,STRING)
    RETURN 999, "xxxxxx"
END FUNCTION
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) break func1
Breakpoint 1 at 0x00000000: file prog.4gl, line 7.

(fgldb) info breakpoints
Num Type           Disp Enb Address    What
1   breakpoint     keep y   0x00000000 in func1 at prog.4gl:7

(fgldb) disable 1

(fgldb) info breakpoints
Num Type           Disp Enb Address    What
1   breakpoint     keep n   0x00000000 in func1 at prog.4gl:7

(fgldb) run
Program exited normally.
```
