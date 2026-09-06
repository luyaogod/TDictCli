---
title: "FGLTRACE_FUNCTIONS"
source: "fgl-topics/c_fgl_EnvVariables_FGLTRACE_FUNCTIONS.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLTRACE_FUNCTIONS"
type: "concept"
---

# FGLTRACE_FUNCTIONS

> Defines the list of functions to be followed by the program execution trace.

The FGLTRACE\_FUNCTIONS environment variable defines the list of functions to
be traced with the program execution trace option.

By default, the trace starts with the `MAIN` function and all functions of the
program are traced.

In order to limit the trace to a given set of functions, define the FGLTRACE\_FUNCTIONS
environment variable with space-separated list of function names. The functions can be prefixed by a
module name:

```
 {
 | function-name
 | module-name.function-name
 }
 [...]
```

1. module-name is the name of a .42m module.
2. function-name is the name of a function.

> **Important:**
>
> Unlike [FGLTRACE\_EXCLUDE](0536-fgltrace-exclude.md "Defines the list of functions to be excluded from the program execution trace."), functions in FGLTRACE\_FUNCTIONS can be specified with or without the
> module prefix. For example, if you want to include the `check_order()` function of
> the orders.4gl module, you can specify "`orders.check_order`" or
> "`check_order`" in FGLTRACE\_FUNCTIONS.

For example, to enable the trace log in the "invoice\_report\_1" function (without specifying a
module), and to enable the trace in the "add\_customer" function, defined in the "custmod"
module:

On UNIX™:

```
$ FGLTRACE_FUNCTIONS="invoice_report_1 custmod.add_customer"
$ export FGLTRACE_FUNCTIONS
```

On Windows™:

```
C:\> set FGLTRACE_FUNCTIONS=invoice_report_1 custmod.add_customer
```

FGLTRACE\_FUNCTIONS has a higher priority than FGLTRACE\_EXCLUDE. The trace is enabled for a
function listed in FGLTRACE\_FUNCTIONS, when is it called from a function excluded by
FGLTRACE\_EXCLUDE.

## Related links

**Related concepts**  

[Execution trace](../13_programming-tools/2632-execution-trace.md "Print a function call stack of your program.")
