---
title: "Defining a breakpoint in the code"
source: "fgl-topics/c_fgl_Debugger_009.html"
breadcrumb: "Programming tools > Integrated debugger > Defining a breakpoint in the code"
type: "concept"
---

# Defining a breakpoint in the code

> Set a breakpoint in the program source code with the BREAKPOINT instruction.

If the program flow encounters this instruction, the program stops as if the break
point was set by the [break
command](2588-break.md "The break command defines a break point to stop the program execution at a given line or function."):

```
MAIN
  DEFINE i INTEGER
  LET i=123
  BREAKPOINT
  DISPLAY i
END MAIN
```

The [BREAKPOINT](../09_advanced-features/0832-breakpoint.md "The BREAKPOINT instruction sets a program breakpoint when running in debug mode.") instruction is
simply ignored when running in normal mode.
