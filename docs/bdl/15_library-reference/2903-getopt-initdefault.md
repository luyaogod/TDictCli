---
title: "Getopt.initDefault()"
source: "fgl-topics/c_fgl_utility_functions_getopt_initdefault.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.initDefault()"
type: "concept"
---

# Getopt.initDefault()

> Initializes a variable defined with the Getopt type.

## Syntax

```
FUNCTION (r Getopt) initDefault(
   options GetoptOptions
   )
```

1. options is a [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.") array that holds the definition of the command line
   options.

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, used to initialize the variable of this type, using the
current program name (`arg_val(0)`) and where the command line argument starting
index is 1 ([`copyArguments(1)`](2898-getopt-copyarguments.md "Returns a dynamic array of string with all command line arguments starting from the provided index.")).

A variable of the type `Getopt` must be defined, as well as a [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.")
dynamic array containing the definitions of the command line options for the current program.

After initializing the `Getopt` record, use the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method to parse
and validate command line options.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    DEFINE _options getopt.GetoptOptions = [ ... ]

    CALL g.initDefault(_options)
    WHILE g.getopt() == getopt.SUCCESS
        ...
    END WHILE

END MAIN
```
