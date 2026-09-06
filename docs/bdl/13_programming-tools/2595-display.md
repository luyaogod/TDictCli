---
title: "display"
source: "fgl-topics/c_fgl_Debugger_display.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > display"
type: "concept"
---

# display

> The display command displays the specified expression's value each time program execution stops.

## Syntax

```
display expression
```

1. expression is a combination of variables, constants and operators. Read [Expressions in debugger commands](2583-expressions-in-debugger-commands.md "A limited expression syntax can be used in debugger commands.").

## Usage

The `display` command allows you to add an expression to an automatic
display list. The values of the expressions in the list are printed each time program
execution stops. Each expression in the list is assigned a number to identify it.

This command is useful in tracking how the values of expressions change during the
program's execution.

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

(fgldb) run
Breakpoint 1, main() at prog.4gl:5
   2         DEFINE x INTEGER
   3
   4         FOR x = 1 TO 10
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN

(fgldb) display x
1: x = 1

(fgldb) continue
Continuing.
          1
Breakpoint 1, main() at prog.4gl:5
1: x = 2
   2         DEFINE x INTEGER
   3
   4         FOR x = 1 TO 10
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN

(fgldb) continue
Continuing.
          2
Breakpoint 1, main() at prog.4gl:5
1: x = 3
   2         DEFINE x INTEGER
   3
   4         FOR x = 1 TO 10
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN
```
