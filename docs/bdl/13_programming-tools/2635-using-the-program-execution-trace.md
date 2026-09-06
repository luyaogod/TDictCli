---
title: "Using the program execution trace"
source: "fgl-topics/c_fgl_prog_trace_usage.html"
breadcrumb: "Programming tools > Execution trace > Using the program execution trace"
type: "concept"
---

# Using the program execution trace

> The program execution trace is typically used in a development environment, to understand the program flow.

## Data produced by the execution trace can be huge

> **Important:**
>
> The amount of data produced can fill up the disk and potentially crash the
> system in a very short time. Consider using this feature carefully, especially in production
> environments.

## Starting a program with trace

Take the following code
example:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF INTEGER,
           c INTEGER
    LET arr[1] = 123
    LET arr[2] = 456
    LET c = incr_elems(arr,5)
END MAIN

FUNCTION incr_elems(arr,val)
    DEFINE arr DYNAMIC ARRAY OF INTEGER,
           val INTEGER
    DEFINE i INTEGER
    FOR i=1 TO arr.getLength()
        LET arr[i] = arr[i]+val
    END FOR
    RETURN arr.getLength()
END FUNCTION
```

In order to get a function call stack trace, start your program with the `--trace`
option of fglrun:

```
fglrun --trace prog.42m
```

The resulting trace printed to the
stderr stream will look as
follows:

```
Enter prog.main()
  Enter prog.incr_elems({123,456},5) from prog.main at line 6
    Enter base.Array.getLength({123,456}) from prog.incr_elems at line 13
    Return 2 from base.Array.getLength
    Enter base.Array.getLength({128,456}) from prog.incr_elems at line 13
    Return 2 from base.Array.getLength
    Enter base.Array.getLength({128,461}) from prog.incr_elems at line 13
    Return 2 from base.Array.getLength
    Enter base.Array.getLength({128,461}) from prog.incr_elems at line 16
    Return 2 from base.Array.getLength
  Return 2 from prog.incr_elems at line 16
Return from prog.main at line 7
```

## Defining the functions to trace

By default, the trace starts with the `MAIN` function and prints all subsequent
functions calls of the program.

To define the functions where the trace should start, define the [FGLTRACE\_FUNCTIONS](../07_configuration/0535-fgltrace-functions.md "Defines the list of functions to be followed by the program execution trace.") environment
variable.

This variable takes a space-separated list of function names.

In FGLTRACE\_FUNCTIONS, the functions can be specified with or without their module prefix.

For example, with the following definition:

```
$ export FGLTRACE_FUNCTIONS="prog.incr_elems otherfunc1 otherfunc2"
```

The new resulting trace will now start at
`prog.incr_elems()`:

```
Enter prog.incr_elems({123,456},5) from prog.main at line 6
  Enter base.Array.getLength({123,456}) from prog.incr_elems at line 13
  Return 2 from base.Array.getLength
  Enter base.Array.getLength({128,456}) from prog.incr_elems at line 13
  Return 2 from base.Array.getLength
  Enter base.Array.getLength({128,461}) from prog.incr_elems at line 13
  Return 2 from base.Array.getLength
  Enter base.Array.getLength({128,461}) from prog.incr_elems at line 16
  Return 2 from base.Array.getLength
Return 2 from prog.incr_elems at line 16
```

## Excluding functions from the trace

By default, the trace includes also Genero BDL library APIs, such as
`base.Array.getLength()`.

Usually you are not interested in built-in function calls: You want to trace your own
functions.

In order to exclude a complete set of methods, class or even a complete package, use the [FGLTRACE\_EXCLUDE](../07_configuration/0536-fgltrace-exclude.md "Defines the list of functions to be excluded from the program execution trace.") environment variable, to
define a space-separated list of exclude-patterns, using `*` and `?`
wildcards and `[a-z]` character ranges, as with the `MATCHES`
operator.

In FGLTRACE\_EXCLUDE, user functions must be specified with their module prefix.

In our example, we want to exclude all methods from the build-in dynamic array class, and
additionally we can also want to exclude all the calls to the classes of the om.\* package (even if
these are not used in this code
example):

```
$ export FGLTRACE_EXCLUDE="<builtin>.* base.Array.* om.*"
```

> **Tip:**
>
> To exclude all global built-in functions such as `fgl_getenv()`, use
> the `<builtin>.*` exclusion pattern.

The new resulting trace will now look as follows, focusing on the user function calls
only:

```
Enter prog.incr_elems({123,456},5) from prog.main at line 6
Return 2 from prog.incr_elems at line 16
```

## Mixing FGLTRACE\_FUNCTIONS and FGLTRACE\_EXCLUDE

If both FGLTRACE\_FUNCTIONS and FGLTRACE\_EXCLUDE identify the same function, then
FGLTRACE\_FUNCTIONS has a higher priority than FGLTRACE\_EXCLUDE.

For example, with the following settings, the function `mod2.delete_all` will be
traced:

```
$ export FGLTRACE_FUNCTIONS="mod1.incr_elems mod2.delete_all"
$ export FGLTRACE_EXCLUDE="mod2.*"
```

The function call chain does not matter: A function listed in FGLTRACE\_FUNCTIONS is always
traced. Consequently, the trace is enabled for a function listed in FGLTRACE\_FUNCTIONS, even when is
it called from a function excluded by FGLTRACE\_EXCLUDE. This might cause gaps in the trace
output.

## Limited trace output

In order to avoid huge logs when the program creates large string variable or arrays,
the output of values is voluntarily limited to a given size. In such case, the trace prints
ellipsis (...), to indicate that the actual value is larger.

Try for example to change the MAIN block to fill the array with more
elements:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF INTEGER,
           c INTEGER
    DEFINE i INTEGER
    FOR i=1 TO 100
        LET arr[i] = i
    END FOR
    LET c = incr_elems(arr,5)
END MAIN
...
```

Compile and run the program again. The new trace output will show only the first elements of the
array passed as parameter to the `incr_elems()`
function:

```
Enter prog.incr_elems({1,2,3,4,5,6,7,8,9,10,11,...more...},..) from prog.main at line 8
Return 100 from prog.incr_elems at line 18
```
