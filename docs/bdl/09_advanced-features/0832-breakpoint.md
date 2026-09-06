---
title: "BREAKPOINT"
source: "fgl-topics/c_fgl_programs_BREAKPOINT.html"
breadcrumb: "Advanced features > Program execution > BREAKPOINT"
type: "concept"
---

# BREAKPOINT

> The BREAKPOINT instruction sets a program breakpoint when running in debug mode.

## Syntax

```
BREAKPOINT
```

## Usage

Normally, to set a breakpoint when you [debug a
program](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs."), you must use the break command of the debugger. But
in some situations, you might need to set the breakpoint in program sources. Therefore,
the `BREAKPOINT` instruction has been added to the language.

When you start [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.")
in debug mode with the `-d` option, if the program flow encounters a
`BREAKPOINT` instruction, the program execution stops and the debug prompt
is displayed, to let you enter a debugger command. The `BREAKPOINT`
instruction is ignored when not running in debug mode.

## Example

```
MAIN
  DEFINE i INTEGER
  LET i=123
  BREAKPOINT
  DISPLAY i 
END MAIN
```

## Related links

**Related concepts**  

[Defining a breakpoint in the code](../13_programming-tools/2582-defining-a-breakpoint-in-the-code.md "Set a breakpoint in the program source code with the BREAKPOINT instruction.")
