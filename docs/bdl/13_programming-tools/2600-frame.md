---
title: "frame"
source: "fgl-topics/c_fgl_Debugger_frame.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > frame"
type: "concept"
---

# frame

> The frame command selects and prints a stack frame.

## Syntax

```
frame [ frame-num ]
```

1. frame-num is the stack frame number of the frame that you wish to
   select.

## Usage

The `frame` command allows you to move from one stack frame to another,
and to print the stack frame that you select. Each stack frame is associated with one call
to one function within the currently executing program. Without an argument, the current
stack frame is printed.

See [stack frames](2581-stack-frames-in-the-debugger.md "The stack frame contains information about function call stack.") for a brief discussion of
frames.

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

(fgldb) break func2
Breakpoint 1 at 0x00000000: file prog.4gl, line 11.

(fgldb) run
Breakpoint 1, func2() at prog.4gl:11
   8     END FUNCTION
   9
   10    FUNCTION func2() RETURNS (INT,STRING)
-> 11        RETURN 999, "xxxxxx"
   12    END FUNCTION
   13

(fgldb) backtrace
#0 func2() at prog.4gl:11
#1 func1() at prog.4gl:7
#2 main() at prog.4gl:2

(fgldb) frame 1
#1 func1() at prog.4gl:7
   4
   5     FUNCTION func1() RETURNS ()
   6         DEFINE x INT, s STRING
-> 7         CALL func2() RETURNING x, s
   8     END FUNCTION
   9
   10    FUNCTION func2() RETURNS (INT,STRING)

(fgldb) frame 2
#2 main() at prog.4gl:2
   1     MAIN
-> 2         CALL func1()
   3     END MAIN
   4
   5     FUNCTION func1() RETURNS ()
```

## Related links

**Related concepts**  

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")
