---
title: "delete"
source: "fgl-topics/c_fgl_Debugger_delete.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > delete"
type: "concept"
---

# delete

> The delete command allows you to remove breakpoints that you have specified in your debugger session.

## Syntax

```
delete breakpoint
```

1. breakpoint is the number assigned to the breakpoint by the debugger.

## Usage

The `delete` command allows you to remove breakpoints set by [break](2588-break.md "The break command defines a break point to stop the program execution at a given line or function.") when they are no longer needed in your debugger session.

If you prefer you may disable the breakpoint instead, see the
[`disable`](2594-disable.md "The disable command disables the specified breakpoint.") command.

`d` is an alias for the `delete` command.

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

(fgldb) delete 1

(fgldb) info breakpoints
No breakpoints or watchpoints.
```
