---
title: "Profiler output: Call graph"
source: "fgl-topics/c_fgl_profiler_006.html"
breadcrumb: "Programming tools > Program profiler > Profiler output: Call graph"
type: "concept"
---

# Profiler output: Call graph

> The profiler call graph provides detailed function call information.

The section "Call graph" provides the following details for each function:

1. The functions that called it, the number of calls, and an estimation
   of the percentage of time spent in these functions.
2. The functions called, the number of calls, and an estimation of
   the time that was spent in the subroutines called from this function.

> **Tip:**
>
> Runtime system internal function names start with the `rts_` prefix. For example,
> the `rts_display()` function implements the `DISPLAY` instruction.
> Showing these internal functions in the profiler output is useful to analyse the code. Usually
> `rts_*` functions appear at the top of the flat profile list, because most of the
> processing time is spend in these runtime functions (for example to execute SQL statements). To
> improve your code, spot the first user function that appears in the flat profile list, and check the
> `%self` number: This is where user code can be improved.

| Column name | Description |
| --- | --- |
| `index` | Each function has an index which appears at the beginning of its primary line. |
| `%total` | Percentage of execution time spent in this function. Includes time spent in subroutines called from this function. |
| `%self` | Percentage of execution time spent in this function excluding the time spent in subroutines called from this function. |
| `%child` | Percentage of execution time spent in the functions called from this function. |
| `calls/of` | Number of calls / Total number of calls |
| `name` | Function name |

Output example:

```
Call graph (order by self)

index     %total    %self   %child     calls/of           name
...
-------------------
            9.67     1.27     8.40         1/2        <-- myprog.main
           23.16     3.31    19.85         1/2        <-- myprog.fB
[5]        32.82     4.58    28.24           2        *** myprog.fC
           28.24     0.00    28.24         7/8        --> mymod.fA
-------------------
...
```

Description:

- The index `[5]` corresponds to the line in the top ["Flat profile"](2625-profiler-output-flat-profile.md "The flat profile shows a summary of the functions called during the program execution.") output.
- The three stars `***` indicate the function that is analyzed:
  `myprog.fC`.
- The `%total` column indicates that the `myprog.fC` function has
  consumed 32.82% of the execution time:
  - 4.58% of the execution time is spent in the current function code (`%self`
    column)
  - 28.24% of the execution time is spent in the called functions (`%child`
    column)
- The `calls/of` column shows that the `myprog.fC` function has been
  called 2 times:
  - One time by `myprog.main`.
  - One time by `myprog.fB`.
- The `myprog.fC` function has called the `mymod.fA` function 7
  times, on a total of 8 calls in the whole program.
