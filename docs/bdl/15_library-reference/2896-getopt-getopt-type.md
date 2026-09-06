---
title: "Getopt.Getopt type"
source: "fgl-topics/c_fgl_utility_functions_getopt_t_getopt.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.Getopt type"
type: "concept"
---

# Getopt.Getopt type

> The Getopt structured type is used to process command line options.

## Syntax

```
TYPE Getopt RECORD
    ... private members not documented here ...
    opt_ind INTEGER,
    opt_char CHAR,
    opt_arg STRING
  END RECORD
```

1. `opt_ind` current command line argument index that is processed.
2. `opt_char` is the single-character short name of the current processed
   option.
3. `opt_arg` if present, holds the value parameter of the current processed option
   (`--option=value`). Otherwise, this member is
   `NULL`.

## Usage

This type defines the `Getopt` record that is used with getopt methods to parse
and validate command line options.

A variable of the type `Getopt` must be defined and initialized with the [`initDefault()`](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.") or
[`initialize()`](2904-getopt-initialize.md "Initializes a variable defined with the Getopt type for command line argument processing.")
method, before using the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method in a `WHILE` loop, to process command line
options.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    DEFINE _options getopt.GetoptOptions = [ ... ]

    CALL g.initDefault(_options)
    ...
END MAIN
```
