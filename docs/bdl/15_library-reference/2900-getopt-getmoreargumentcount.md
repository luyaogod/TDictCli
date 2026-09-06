---
title: "Getopt.getMoreArgumentCount()"
source: "fgl-topics/c_fgl_utility_functions_getopt_getmoreargumentcount.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.getMoreArgumentCount()"
type: "concept"
---

# Getopt.getMoreArgumentCount()

> Returns the number of command line arguments left to be processed after the known options.

## Syntax

```
FUNCTION (r Getopt) getMoreArgumentCount( )
  RETURNS INTEGER
```

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, that returns the number of command line arguments left to
process, after all possible options defined in the `Getopt` object have been
detected.

For more details see [`getMoreArgument()`](2901-getopt-getmoreargument.md "Returns the additional argument at the specified index.").
