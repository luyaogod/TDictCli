---
title: "Getopt.getMoreArgument()"
source: "fgl-topics/c_fgl_utility_functions_getopt_getmoreargument.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.getMoreArgument()"
type: "concept"
---

# Getopt.getMoreArgument()

> Returns the additional argument at the specified index.

## Syntax

```
FUNCTION (r Getopt) getMoreArgument( ind INTEGER )
  RETURNS STRING
```

1. ind is the index of the additional argument, from 1 to [`getMoreArgumentCount()`](2900-getopt-getmoreargumentcount.md "Returns the number of command line arguments left to be processed after the known options.").

## Usage

After processing all command line options, it is possible to scan for additional command line
arguments, that are not part of the options defined in the [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.")
array of the `Getopt` object.

`getMoreArgument()` is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, that
returns the additional argument at the specified index, after all possible options defined in the
`Getopt` object have been detected.

For example, with a command using a set of filenames as argument, you want to
allow:

```
$ fglrun reduce --verbose --ratio=5 file1 file2 file3
```

In such case, the `--verbose` and `--ratio` options must be defined
by the `Getopt` object and processed with the [getopt()](2902-getopt-getopt.md "Process the next command line option.") method, while
file1, file2 and file3 are additional
arguments to be processed after the `getopt()` loop, using the
`getMoreArgumentCount()` method to get the total number of additional arguments, and
the `getMoreArgument()` method to retrieve the additional argument.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    DEFINE cnt, ind INTEGER
    ...
    WHILE g.getopt() == getopt.SUCCESS
        ...
    END WHILE

    IF g.invalidOptionSeen() THEN
        CALL g.displayUsage(NULL)
        EXIT PROGRAM 1
    ELSE
        LET cnt = g.getMoreArgumentCount()
        IF cnt == 0 THEN
            DISPLAY "ERROR: No files were provided..."
            EXIT PROGRAM 1
        ELSE
            FOR ind = 1 TO cnt
                DISPLAY SFMT("File to process: %1", g.getMoreArgument(ind))
            END FOR
        END IF
    END IF

END MAIN
```
