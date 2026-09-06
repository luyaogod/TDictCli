---
title: "Syntax of the program execution trace"
source: "fgl-topics/c_fgl_prog_trace_syntax.html"
breadcrumb: "Programming tools > Execution trace > Syntax of the program execution trace"
type: "concept"
---

# Syntax of the program execution trace

> The execution trace is enable by using the --trace option of fglrun.

In order to enable program execution trace, start the [fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") command with the `--trace` option:

```
fglrun --trace program[.42r] [argument [...]]
```

1. program is the name of the BDL program.
2. argument is a command line argument passed to the program.

Function calls are printed to stderr during program execution.

By default, the trace starts in `MAIN`.

To define a list of functions to be traced, use the [FGLTRACE\_FUNCTIONS](../07_configuration/0535-fgltrace-functions.md "Defines the list of functions to be followed by the program execution trace.") environment
variable.

To exclude a set of functions from the trace, use the [FGLTRACE\_EXCLUDE](../07_configuration/0536-fgltrace-exclude.md "Defines the list of functions to be excluded from the program execution trace.") environment variable.
