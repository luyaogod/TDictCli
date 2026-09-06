---
title: "Understanding the program execution trace"
source: "fgl-topics/c_fgl_prog_trace_basics.html"
breadcrumb: "Programming tools > Execution trace > Understanding the program execution trace"
type: "concept"
---

# Understanding the program execution trace

> This is an introduction to the program execution trace.

The program execution trace prints all function calls of your program, with the values of
parameters passed to and values returned from functions.

> **Important:**
>
> 1. Sensitive and personal data may be written to the output. Make sure that the log output produced
>    by the runtime trace is written to files that can only be read by application administrators.
> 2. If you plan to use the program execution trace on a production site, pay attention to the amount
>    of data produced by the log: Using the trace can fill up the disk and potentially crash the system
>    in a very short time.

This tool can help to understand the code structure of complex programs, and identify bugs in
your code.

To enable the program execution trace, start fglrun with the `-trace` option, for
example:

```
fglrun --trace myprog
```

The function call stack trace will be printed to the stderr
stream:

```
Enter prog.main()
  Enter prog.func1(12499,"abcdef") from prog.main at line 8
    Enter prog.func2(12499,1) from prog.func1 at line 16
    Return 12500 from prog.func2 at line 23
    ...
```

The output format of FGL trace is for debug purpose only and can change in future product
releases.

Additional control is available with FGLTRACE\* environment variables, as described in the [usage topic](2635-using-the-program-execution-trace.md "The program execution trace is typically used in a development environment, to understand the program flow.").
