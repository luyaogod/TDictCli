---
title: "continue"
source: "fgl-topics/c_fgl_Debugger_continue.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > continue"
type: "concept"
---

# continue

> The continue command continues the execution of the program after a breakpoint.

## Syntax

```
continue [ ignore-count ]
```

1. ignore-count defines the number of times to ignore a
   breakpoint at this location.

## Usage

The `continue` command continues the execution of the program until
the program completes normally, another breakpoint is reached, or a signal is received.

`c` is an alias for the `continue` command.

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

(fgldb) run
Breakpoint 1, func1() at prog.4gl:7
   4
   5     FUNCTION func1() RETURNS ()
   6         DEFINE x INT, s STRING
-> 7         CALL func2() RETURNING x, s
   8     END FUNCTION
   9
   10    FUNCTION func2() RETURNS (INT,STRING)

(fgldb) continue
Continuing.
Program exited normally.
```
