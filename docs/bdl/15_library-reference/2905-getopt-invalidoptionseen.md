---
title: "Getopt.invalidOptionSeen()"
source: "fgl-topics/c_fgl_utility_functions_getopt_invalidoptionseen.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.invalidOptionSeen()"
type: "concept"
---

# Getopt.invalidOptionSeen()

> Checks if the command line options are misused.

## Syntax

```
FUNCTION (r Getopt) invalidOptionSeen( )
  RETURNS BOOLEAN
```

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, that returns `TRUE` if the command line options
are not specified correctly as defined by the [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.")
definition of the `Getopt` object.

The command line is considered as invalid in the following cases:

- When an unknown option is used (not defined in the `GetoptOptions` array).
- When a required option argument is missing (`arg_type == getopt.REQUIRED`).

The `invalidOptionSeen()` method can be used after calling the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method processing
command line options in a loop, to detect invalid options usage.

When an invalid command line option usage is detected, display the usage with the [`displayUsage()`](2899-getopt-displayusage.md "Display the usage and command line option description to the standard output stream.")
method and exit the program with an error status with [`EXIT PROGRAM 1`](../09_advanced-features/0831-exit-program.md "The EXIT PROGRAM instruction terminates the execution of the program.").

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    ...
    WHILE g.getopt() == getopt.SUCCESS
        ...
    END WHILE

    IF g.invalidOptionSeen() THEN
        CALL g.displayUsage()
        EXIT PROGRAM 1
    END IF

END MAIN
```
