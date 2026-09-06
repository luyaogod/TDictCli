---
title: "backtrace / where"
source: "fgl-topics/c_fgl_Debugger_backtrace_where.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > backtrace / where"
type: "concept"
---

# backtrace / where

> The backtrace commands prints a summary of how your program reached the current state.

## Syntax

```
backtrace
```

## Usage

The `backtrace` command prints a summary of your program's entire stack,
one line per frame.

Each line in the output shows the frame number and function name.

`bt` and `where` are aliases for the `backtrace`
command.

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
```
