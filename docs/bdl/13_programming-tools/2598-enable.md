---
title: "enable"
source: "fgl-topics/c_fgl_Debugger_enable.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > enable"
type: "concept"
---

# enable

> The enable command enables breakpoints that have previously been disabled.

## Syntax

```
enable breakpoint
```

1. breakpoint is the number assigned to the breakpoint by the debugger.

## Usage

The `enable` command allows you to reactivate a breakpoint in the current
debugger session.

The breakpoint must have been disabled using the [`disable`](2594-disable.md "The disable command disables the specified breakpoint.") command.

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

(fgldb) disable 1

(fgldb) info breakpoints
Num Type           Disp Enb Address    What
1   breakpoint     keep n   0x00000000 in func1 at prog.4gl:7

(fgldb) enable 1

(fgldb) run
Breakpoint 1, func1() at prog.4gl:7
   4
   5     FUNCTION func1() RETURNS ()
   6         DEFINE x INT, s STRING
-> 7         CALL func2() RETURNING x, s
   8     END FUNCTION
   9
   10    FUNCTION func2() RETURNS (INT,STRING)
```
