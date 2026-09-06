---
title: "Getopt.isSuccess()"
source: "fgl-topics/c_fgl_utility_functions_getopt_issuccess.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.isSuccess()"
type: "concept"
---

# Getopt.isSuccess()

> Checks if a command line option parsing succeeded.

## Syntax

```
FUNCTION (r Getopt) isSuccess( )
  RETURNS BOOLEAN
```

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, that returns `TRUE` if last command line
argument was processed successfully.

The `isSuccess()` method can be used after calling the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method processing
command line options in a loop, to detect correct options usage.

The processing status of the current option is typically checked with the return code of the
`getopt()` method.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    ...
    WHILE g.getopt() == getopt.SUCCESS
        ...
    END WHILE

    IF g.isSuccess() THEN
        ...
    END IF

END MAIN
```
