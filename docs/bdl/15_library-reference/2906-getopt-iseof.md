---
title: "Getopt.isEof()"
source: "fgl-topics/c_fgl_utility_functions_getopt_iseof.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.isEof()"
type: "concept"
---

# Getopt.isEof()

> Checks if there are more command line options to be read.

## Syntax

```
FUNCTION (r Getopt) isEof( )
  RETURNS BOOLEAN
```

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, used to check if there are no more command line options to
process.

The `isEof()` method can be used after calling the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method processing
command line options, to stop the processing when no more known command line options are to be
processed.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    DEFINE s INTEGER
    ...

    WHILE TRUE
        LET s = g.getOpt()
        IF g.isEof() THEN EXIT WHILE END IF
        DISPLAY g.opt_char
    END WHILE

END MAIN
```
