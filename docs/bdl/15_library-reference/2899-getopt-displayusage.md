---
title: "Getopt.displayUsage()"
source: "fgl-topics/c_fgl_utility_functions_getopt_displayusage.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.displayUsage()"
type: "concept"
---

# Getopt.displayUsage()

> Display the usage and command line option description to the standard output stream.

## Syntax

```
FUNCTION (r Getopt) displayUsage(
   more_args STRING
   )
```

1. more\_args is a character string to be displayed at the end of the "Usage..." line.

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, used to display the description of the command line options of
the current program.

The more\_args passed as parameter to the `displayUsage()`
method will be shown in the first line, after the option
list:

```
Usage: program-name [options] more_args
```

Before using the `displayUsage()` method, a variable of the type
`Getopt` must be defined and initialized with the [`initDefault()`](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.") or
[`initialize()`](2904-getopt-initialize.md "Initializes a variable defined with the Getopt type for command line argument processing.")
method.

The `displayUsage()` method is typically called when an invalid command line
option is detected (when [`invalidOptionSeen()`](2905-getopt-invalidoptionseen.md "Checks if the command line options are misused.") returns `TRUE`) or when the
`--help/-h` option is used.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    ...
    IF g.opt_char=="h" THEN
        CALL g.displayUsage(NULL)
        EXIT PROGRAM 0
    END IF
    ...
END MAIN
```
