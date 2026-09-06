---
title: "FGLTRACE_EXCLUDE"
source: "fgl-topics/c_fgl_EnvVariables_FGLTRACE_EXCLUDE.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLTRACE_EXCLUDE"
type: "concept"
---

# FGLTRACE_EXCLUDE

> Defines the list of functions to be excluded from the program execution trace.

The FGLTRACE\_EXCLUDE environment variable defines the list of patterns to identify functions (and
class methods) that must be excluded from the call stack trace.

By default, all functions are traced, starting from the `MAIN` function, or from
the functions listed in [FGLTRACE\_FUNCTIONS](0535-fgltrace-functions.md "Defines the list of functions to be followed by the program execution trace.").

In order to exclude functions you don't want to trace, define the FGLTRACE\_EXCLUDE environment
variable with space-separated list of patterns:

```
exclude-pattern [...]
```

1. exclude-pattern is a string with wildcards like in a `MATCHES`
   expression:
   - The `*` wildcard represents 0 to n characters.
   - The `?` wildcard represent a single character.
   - The `[ ]` wildcards can be used to define a single character in the specified
     range (use a `^` starting caret for negation).

> **Important:**
>
> Unlike FGLTRACE\_FUNCTIONS, user functions in FGLTRACE\_EXCLUDE must
> be specified with their module prefix. For example, if you want to exclude the
> `check_order()` function of the orders.4gl module, use the
> qualified name "orders.check\_order" in FGLTRACE\_EXCLUDE.

The FGLTRACE\_EXCLUDE environment variable is typically used to exclude build-in functions and
classes such as `base.Array.*`. Tracing build-in functions can produce a huge log and
is not always relevant. For example, if the program uses the base.Array.getLength() method, the
trace will report each call to the method, as in `FOR i=1 TO arr.getLength()`.

To exclude all global built-in functions such as `fgl_getenv()`, use the
`<builtin>.*` exclusion pattern.

For example, to exclude all methods from the build-in dynamic array class, from the
`util.JSON` class, all om.\* package calls and all functions of
`mymodule` with the format `debug_[0-9]???`, define the variable as
follows:

On UNIX™:

```
$ FGLTRACE_EXCLUDE="<builtin>.* base.Array.* util.JSON.* om.* mymodule.debug_[0-9]???"
$ export FGLTRACE_EXCLUDE
```

On Windows®:

```
C:\> set FGLTRACE_EXCLUDE=<builtin>.* base.Array.* util.JSON.* om.* mymodule.debug_[0-9]???
```

## Related links

**Related concepts**  

[Execution trace](../13_programming-tools/2632-execution-trace.md "Print a function call stack of your program.")
