---
title: "watch"
source: "fgl-topics/c_fgl_Debugger_watch.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > watch"
type: "concept"
---

# watch

> The watch command sets a watchpoint for an expression.

## Syntax

```
watch expression [if boolean-expression]
```

1. expression is a combination of variables, constants and operators. Read [Expressions in debugger commands](2583-expressions-in-debugger-commands.md "A limited expression syntax can be used in debugger commands.").
2. boolean-expression is an optional boolean expression.

## Usage

The `watch` command
stops the program execution when the value of the expression changes.

If boolean-expression is
provided, the `watch` command stops the execution
of the program if the expression value has changed and the boolean-expression evaluates
to true.

The watchpoint cannot be set if the program is not
in the context where expression can be evaluated.
Before using a watchpoint, you typically set a breakpoint in the function
where the expression makes sense, then you run
the program, and then you set the watchpoint. This example illustrates
this procedure.

## Example

```
MAIN
   DEFINE i INTEGER

   LET i = 1 
   DISPLAY i
   LET i = 2 
   DISPLAY i
   LET i = 3 
   DISPLAY i

 END MAIN
```

```
(fgldb) break main
breakpoint 1 at 0x00000000: file test.4gl, line 4
(fgldb) run
Breakpoint 1, main() at test.4gl:4
4      LET i = 1
(fgldb) watch i if i >= 3
Watchpoint 1:  i
(fgldb) continue
1
2
Watchpoint 1:  i

Old value = 2
New value = 3
main() at t.4gl:9
9         DISPLAY i
(fgldb)
```
