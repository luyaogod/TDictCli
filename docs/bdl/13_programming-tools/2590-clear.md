---
title: "clear"
source: "fgl-topics/c_fgl_Debugger_clear.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > clear"
type: "concept"
---

# clear

> The clear command clears the breakpoint at a specified line or function.

## Syntax

```
clear [ location ]
```

where location
is:

```
{ [module.]function
| [module:]line
}
```

1. module is the name of a source file, without extension.
2. function is a function name.
3. line is a source code line.

## Usage

With the `clear` command, you can delete breakpoints set by [break](2588-break.md "The break command defines a break point to stop the program execution at a given line or function.") wherever they are in your program.

Use the `clear` command with no arguments to delete any
breakpoints at the next instruction to be executed in the selected
stack frame.

Use the [delete](2592-delete.md "The delete command allows you to remove breakpoints that you have specified in your debugger session.") command
to delete individual breakpoints by specifying their breakpoint numbers.

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

(fgldb) info breakpoints
Num Type           Disp Enb Address    What
1   breakpoint     keep y   0x00000000 in func2 at prog.4gl:11

(fgldb) clear func2
Deleted breakpoint 1

(fgldb) info breakpoints
No breakpoints or watchpoints.
```
