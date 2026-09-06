---
title: "Syntax of the program profiler"
source: "fgl-topics/c_fgl_profiler_003.html"
breadcrumb: "Programming tools > Program profiler > Syntax of the program profiler"
type: "concept"
---

# Syntax of the program profiler

> The program profiler is enabled by using the -p option of fglrun.

To activate the program profiler, start the [fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.")
command with the `-p` option:

```
fglrun -p program[.42r] [argument [...]]
```

1. program is the name of the BDL program.
2. argument is a command line argument passed to the program.

Profiling statistics are collected during program execution, and printed when
the program ends.
