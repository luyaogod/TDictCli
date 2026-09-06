---
title: "Getopt.getopt()"
source: "fgl-topics/c_fgl_utility_functions_getopt_getopt.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.getopt()"
type: "concept"
---

# Getopt.getopt()

> Process the next command line option.

## Syntax

```
FUNCTION (r Getopt) getopt( )
  RETURNS INTEGER
```

1. The returned status can be one of:
   - `getopt.SUCCESS`: A command line argument was detected and can be processed.
   - `getopt.EOF`: No more command line arguments to process, exit the
     `WHILE` loop.
   - `getopt.BAD_ARGUMENT`: An invalid command line argument was detected, exit the
     `WHILE` loop and inform the user.

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, that processes the next command line option.

Before using the `getopt()` method, a variable of the type `Getopt`
must be defined and initialized with the [`initDefault()`](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.") or
[`initialize()`](2904-getopt-initialize.md "Initializes a variable defined with the Getopt type for command line argument processing.")
method.

The `getopt()` method is typically used with a `WHILE` loop to
process all command line arguments, by checking that the method returns
`getopt.SUCCESS`:

```
    WHILE g.getopt() == getopt.SUCCESS
        ...
    END WHILE
```

Inside the `WHILE` loop, check the properties of the Getopt record to identify the
current option that has been found. The `opt_char` member identifies the current
option with its single-character short name, and the `opt_arg` member holds the
option argument (`--option=value`), if one is
available:

```
        CASE g.opt_char
            WHEN 'e'
               CALL extract(g.opt_arg)
            ...
        END CASE
```

When the `WHILE` loop ends because `getopt()` has returned a status
different from `getopt.SUCCESS`, handle errors by checking the status with the [`invalidOptionSeen()`](2905-getopt-invalidoptionseen.md "Checks if the command line options are misused.")
method. If all provided options are correct, process additional command line arguments if needed.
For more details, see [`getMoreArgument()`](2901-getopt-getmoreargument.md "Returns the additional argument at the specified index.").

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE outfile STRING
    DEFINE g getopt.Getopt
    DEFINE _options getopt.GetoptOptions
        = [(name: "version",
            description: "Version information",
            opt_char: 'v',
            arg_type: getopt.NONE),
           (name: "help",
            description: "This help page",
            opt_char: 'h',
            arg_type: getopt.NONE),
           (name: "outfile",
            description: "Output filename",
            opt_char: 'o',
            arg_type: getopt.REQUIRED)]

    CALL g.initDefault(_options)
    WHILE g.getOpt() == getopt.SUCCESS
        CASE g.opt_char
            WHEN 'v'
                DISPLAY "Version 1.50"
                EXIT PROGRAM 0
            WHEN 'h'
                CALL g.displayUsage("file ...")
                EXIT PROGRAM 0
            WHEN 'o'
                LET outfile = g.opt_arg
        END CASE
    END WHILE

    IF g.invalidOptionSeen() THEN
        CALL g.displayUsage("file ...")
        EXIT PROGRAM 1
    END IF

END MAIN
```
