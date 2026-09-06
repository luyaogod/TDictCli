---
title: "GetOpt constants"
source: "fgl-topics/c_fgl_utility_functions_getopt_constants.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > GetOpt constants"
type: "concept"
---

# GetOpt constants

> List of predefined constants for the getopt API.

## Syntax

```
CONSTANT NONE      = 1
CONSTANT REQUIRED  = 2
CONSTANT OPTIONAL  = 3

CONSTANT SUCCESS       = 0
CONSTANT EOF           = 1
CONSTANT BAD_ARGUMENT  = 2
```

Constant values are provided in this syntax diagram for information. You should obviously use the
constant names instead of the values.

## Usage

The GetOpt predefined constants are used to implement command line option processing with the
getopt module.

The `getopt.NONE`, `getopt.REQUIRED`,
`getopt.OPTIONAL` constants are used to define the `opt_arg` member of
elements of the [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.") array used to define the command line options.

The `getopt.SUCCESS`, `getopt.EOF`,
`getopt.BAD_ARGUMENT` are return codes of the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method.
