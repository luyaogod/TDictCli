---
title: "print"
source: "fgl-topics/c_fgl_Debugger_print.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > print"
type: "concept"
---

# print

> The print command displays the current value of the specified expression.

## Syntax

```
print expression
```

1. expression is a combination of variables, constants and operators. Read [Expressions in debugger commands](2583-expressions-in-debugger-commands.md "A limited expression syntax can be used in debugger commands.").

## Usage

The `print` command allows you to examine the data in your program.

It evaluates and prints the value of the specified expression from your program,
in a format appropriate to its data type.

`p` is an alias for the
`print` command.

## Example

```
MAIN
    DEFINE x INTEGER

    FOR x = 100 TO 200
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
   4         FOR x = 100 TO 200
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN

(fgldb) print x
$1 = 100
```
