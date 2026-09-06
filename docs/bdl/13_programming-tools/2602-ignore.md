---
title: "ignore"
source: "fgl-topics/c_fgl_Debugger_ignore.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > ignore"
type: "concept"
---

# ignore

> The ignore command defines the number of times a breakpoint must be ignored.

## Syntax

```
ignore breakpoint count
```

1. breakpoint is the breakpoint number.
2. count is the number of times the breakpoint will be ignored.

## Usage

The `ignore` command defines the number of times a breakpoint
is ignored when the program flow reaches that breakpoint.

The next count times the breakpoint is reached, the program
execution will continue, and no [breakpoint
condition](2588-break.md "The break command defines a break point to stop the program execution at a given line or function.") is checked.

You can specify a count  of zero to make the breakpoint stop
the next time it is reached.

When using the [continue](2591-continue.md "The continue command continues the execution of the program after a breakpoint.") command
to resume the execution of the program from a breakpoint, you can specify a an
ignore count directly as an argument.

## Example

```
MAIN
    DEFINE x INTEGER

    FOR x = 1 TO 10
        DISPLAY x
    END FOR

END MAIN
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) break 5
Breakpoint 1 at 0x00000000: file prog.4gl, line 5.

(fgldb) ignore 1 3
Will ignore next 3 crossings of breakpoint 1.
(fgldb) run
          1
          2
          3
Breakpoint 1, main() at prog.4gl:5
   2         DEFINE x INTEGER
   3
   4         FOR x = 1 TO 10
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN

(fgldb) print x
$1 = 4
```
