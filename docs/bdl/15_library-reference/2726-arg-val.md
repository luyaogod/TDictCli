---
title: "arg_val()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ARG_VAL.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > arg_val()"
type: "concept"
---

# arg_val()

> Returns a command line argument by position.

## Syntax

```
FUNCTION arg_val(
   index INTEGER )
  RETURNS STRING
```

1. index is an integer defining the argument position.

## Usage

This function provides a mechanism for passing values to the program through
the command line that invokes the program. You can design a program to expect
or allow arguments after the name of the program in the command line.

The index parameter defines the argument to be returned.
0 returns the name of the program, 1 returns the first argument.

Like all built-in functions, `arg_val()` can be invoked from
any program block. You can use it to pass values to `MAIN`,
which cannot have formal arguments, but you are not restricted to calling
`arg_val()` from the `MAIN` statement.

Use the `arg_val()` function to retrieve individual arguments during program
execution. Use the [`num_args()`](2731-num-args.md "Returns the number of program arguments.") function to determine how many arguments follow the program name
on the command line.

If index is greater than 0, `arg_val(index)`
returns the command-line argument used at a given position. The value of index
must be between 0 and the value returned by `num_args()`, the number of command-line
arguments. The expression `arg_val(0)` returns the name of the application
program.

If the argument is negative or greater than `num_args()`,
the method returns `NULL`.

## Related links

**Related concepts**  

[base.Application.getArgument](2970-base-application-getargument.md "Returns the command line argument by position.")
