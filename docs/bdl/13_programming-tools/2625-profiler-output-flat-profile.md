---
title: "Profiler output: Flat profile"
source: "fgl-topics/c_fgl_profiler_005.html"
breadcrumb: "Programming tools > Program profiler > Profiler output: Flat profile"
type: "concept"
---

# Profiler output: Flat profile

> The flat profile shows a summary of the functions called during the program execution.

The flat profile contains a list of the functions called while the programs was running.

> **Tip:**
>
> Runtime system internal function names start with the `rts_` prefix. For example,
> the `rts_display()` function implements the `DISPLAY` instruction.
> Showing these internal functions in the profiler output is useful to analyse the code. Usually
> `rts_*` functions appear at the top of the flat profile list, because most of the
> processing time is spend in these runtime functions (for example to execute SQL statements). To
> improve your code, spot the first user function that appears in the flat profile list, and check the
> `%self` number: This is where user code can be improved.

| Column Name | Description |
| --- | --- |
| `index` | Index of the function in the list. |
| `%total` | Percentage of execution time spent in this function. Includes time spent in subroutines called from this function. |
| `%self` | Percentage of execution time spent in this function excluding the time spent in subroutines called from this function. |
| `%child` | Percentage of execution time spent in functions called from this function. |
| `calls` | Number of times this function was called. |
| `name` | Function name |

100% represents the total processing time. For example, the time spent waiting for user
interaction is ignored.

Output example:

```
Flat profile (order by self)

index     %total    %self   %child     calls    name
[1]         45.3     45.3      0.0        25    <builtin>.rts_display
[2]         76.6     19.3     57.3         1    myprog.fB
[3]         17.3     17.3      0.0        72    <builtin>.rts_Concat
[4]         62.3      8.4     53.9         8    mymod.fA
[5]         32.8      4.6     28.2         2    myprog.fC
[6]        100.0      3.1     96.9         1    myprog.main
[7]          2.0      2.0      0.0         8    <builtin>.rts_forInit
```

Description:

- The lines are ordered by the percentage of time spent in the actual function code
  (`%self` column), in descending order, to show most time consuming functions
  first.
- 45.3% of the time was spent in the `rts_display` function to ouput text to the
  terminal. Since this is a built-in function, there is no user code to improve here.
- 17.3% of the time is spent in the `rts_Concat` built-in function. No improvement
  possible here.
- 19.3% + 8.4% = 1/4 of the program execution time is spent in the `mymod.fB` and
  `myprog.fA` user defined functions. This is where improvements might be
  possible.
