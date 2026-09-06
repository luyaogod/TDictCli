---
title: "Getopt.copyArguments()"
source: "fgl-topics/c_fgl_utility_functions_getopt_copyarguments.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.copyArguments()"
type: "concept"
---

# Getopt.copyArguments()

> Returns a dynamic array of string with all command line arguments starting from the provided index.

## Syntax

```
FUNCTION copyArguments( ind INTEGER )
  RETURNS DYNAMIC ARRAY OF STRING
```

1. ind is the command line argument index to start with, to produce the dynamic
   array of string. Use 1 for the first argument (0 is the program name).

## Usage

The `copyArguments()` function builds a dynamic array of strings with command line
arguments, starting from the index passed as parameter.

Use this function as second parameter for the `initialize()` method, to provide
the list of command line arguments to be processed.

For more details, see [`initialize()`](2904-getopt-initialize.md "Initializes a variable defined with the Getopt type for command line argument processing.").
