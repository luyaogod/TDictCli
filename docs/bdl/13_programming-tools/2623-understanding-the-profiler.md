---
title: "Understanding the profiler"
source: "fgl-topics/c_fgl_profiler_002.html"
breadcrumb: "Programming tools > Program profiler > Understanding the profiler"
type: "concept"
---

# Understanding the profiler

> The profiler is a tool built into the runtime system that allows you to know where your program spends processing time, and which function calls which function.

The profiler can help you to identify parts of your program that are executing slower than
expected.

> **Important:**
>
> The program profiler shows percentages of actual processing times: If the program is waiting for
> a user-interaction, or when the program sleeps, this has no impact on the results. However, the
> numbers include the time spent waiting for disk I/O, to connect to the front-end, or to perform SQL
> statements in the database server.

In order to enable the profiler during the execution of a program, you must start
fglrun with the `-p` option, for example:

```
fglrun -p myprog
```

When the program ends, the profiler dumps profiling information to the standard error stream.

The times reported by the profiler can change from one execution to the other, depending on the
available system resources. Consider executing your program several times, to get an average
time.
