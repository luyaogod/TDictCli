---
title: "tbreak"
source: "fgl-topics/c_fgl_Debugger_tbreak.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > tbreak"
type: "concept"
---

# tbreak

> The tbreak command sets a temporary breakpoint.

## Syntax

```
tbreak [ location ] [ if condition ]
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
4. condition is an [expression](2583-expressions-in-debugger-commands.md "A limited expression syntax can be used in debugger commands.")
   evaluated dynamically.

## Usage

The `tbreak` (temporary break) command sets a breakpoint for one stop only.

The breakpoint is set in the same way as with the [`break`](2588-break.md "The break command defines a break point to stop the program execution at a given line or function.") command, but the breakpoint is
automatically deleted after the first time your program stops there.

If a condition is specified, the program stops at the breakpoint only if
the condition evaluates to true.

If you do not specify any location (function or line number), the breakpoint is created for
the current line. For example, if you write "`tbreak if var = 1`", the debugger
adds a conditional breakpoint for the current line, and the program will only stop if the
variable is equal to 1 when reaching the current line again.

## Example

```
MAIN
    DEFINE x INTEGER

    FOR x = 1 TO 3
        DISPLAY x
    END FOR

END MAIN
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) tbreak 5
Breakpoint 1 at 0x00000000: file prog.4gl, line 5.

(fgldb) run
main() at prog.4gl:5
   2         DEFINE x INTEGER
   3
   4         FOR x = 1 TO 3
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN

(fgldb) continue
Continuing.
          1
          2
          3
Program exited normally.
```
