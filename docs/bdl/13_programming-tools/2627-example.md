---
title: "Example"
source: "fgl-topics/c_fgl_profiler_007.html"
breadcrumb: "Programming tools > Program profiler > Example"
type: "concept"
---

# Example

> Program profiler example.

## Source code

File: mymod.4gl

```
PUBLIC FUNCTION fA(s,n_a)
  DEFINE s STRING
  DEFINE n_a,i INTEGER
  FOR i=1 TO n_a
    DISPLAY "fA "||s||" n:"||i
  END FOR
END FUNCTION
```

File: myprog.4gl

```
IMPORT FGL mymod

MAIN
  DISPLAY "Profiler sample"
  CALL fB()
  CALL fC(2)
END MAIN

PRIVATE FUNCTION fB()
  CALL mymod.fA("fB",10)
  CALL fC(5)
END FUNCTION

PRIVATE FUNCTION fC(n_c)
  DEFINE n_c INTEGER
  WHILE n_c > 0
    CALL mymod.fA("fC",2)
    LET n_c=n_c-1
  END WHILE
END FUNCTION
```

## Profiler output (stderr)

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

Call graph (order by self)

index     %total    %self   %child     calls/of           name
           10.69    10.69     0.00         1/25       <-- myprog.main
           34.61    34.61     0.00        24/25       <-- mymod.fA
[1]        45.29    45.29     0.00           25       *** <builtin>.rts_display
-------------------
           76.59    19.34    57.25         1/1        <-- myprog.main
[2]        76.59    19.34    57.25           1        *** myprog.fB
           34.10     8.40    25.70         1/8        --> mymod.fA
           23.16     3.31    19.85         1/2        --> myprog.fC
-------------------
           17.30    17.30     0.00        72/72       <-- mymod.fA
[3]        17.30    17.30     0.00           72       *** <builtin>.rts_Concat
-------------------
           28.24     0.00    28.24         7/8        <-- myprog.fC
           34.10     8.40    25.70         1/8        <-- myprog.fB
[4]        62.34     8.40    53.94           8        *** mymod.fA
           34.61    34.61     0.00        24/25       --> <builtin>.rts_display
           17.30    17.30     0.00        72/72       --> <builtin>.rts_Concat
            2.04     2.04     0.00         8/8        --> <builtin>.rts_forInit
-------------------
            9.67     1.27     8.40         1/2        <-- myprog.main
           23.16     3.31    19.85         1/2        <-- myprog.fB
[5]        32.82     4.58    28.24           2        *** myprog.fC
           28.24     0.00    28.24         7/8        --> mymod.fA
-------------------
          100.00     3.05    96.95         1/1        <-- <top>
[6]       100.00     3.05    96.95           1        *** myprog.main
           76.59    19.34    57.25         1/1        --> myprog.fB
           10.69    10.69     0.00         1/25       --> <builtin>.rts_display
            9.67     1.27     8.40         1/2        --> myprog.fC
-------------------
            2.04     2.04     0.00         8/8        <-- mymod.fA
[7]         2.04     2.04     0.00           8        *** <builtin>.rts_forInit
-------------------
```

## Interpretation

In the top flat profile, you can see that most of the time is spent in built-in
`rts_*` functions:

```
Flat profile (order by self)

index     %total    %self   %child     calls    name
[1]         45.3     45.3      0.0        25    <builtin>.rts_display
[2]         76.6     19.3     57.3         1    myprog.fB
[3]         17.3     17.3      0.0        72    <builtin>.rts_Concat
...
```

This is interesting, but it's not possible for you to improve the built-in functions.

The user function that takes most of the time is `mymod.fB` at the second
position, and most of the time is spend in the child calls
(57.3%):

```
Flat profile (order by self)

index     %total    %self   %child     calls    name
[1]         45.3     45.3      0.0        25    <builtin>.rts_display
[2]         76.6     19.3     57.3         1    myprog.fB
[3]         17.3     17.3      0.0        72    <builtin>.rts_Concat
...
```

The second block in the call graph shows more details about the `mymod.fB`
function:

```
Call graph (order by self)

index     %total    %self   %child     calls/of           name
...
           76.59    19.34    57.25         1/1        <-- myprog.main
[2]        76.59    19.34    57.25           1        *** myprog.fB
           34.10     8.40    25.70         1/8        --> mymod.fA
           23.16     3.31    19.85         1/2        --> myprog.fC
```

The `mymod.fB`, `mymod.fA`, and `mymod.fC` functions
is where you can improve the code. For example, in `mymod.fA`, you can reduce the
number of `DISPLAY` calls (`rts_display`), or `||`
concatenation operator use (`rts_Concat`).
