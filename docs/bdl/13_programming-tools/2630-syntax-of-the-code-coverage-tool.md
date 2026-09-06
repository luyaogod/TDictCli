---
title: "Syntax of the code coverage tool"
source: "fgl-topics/c_fgl_coverage_syntax.html"
breadcrumb: "Programming tools > Source code coverage > Syntax of the code coverage tool"
type: "concept"
---

# Syntax of the code coverage tool

> The code coverage tool is enabled by setting the FGLCOV environment variable.

## Enabling coverage data generation

To collect coverage data while executing programs, you must set the [FGLCOV](../07_configuration/0521-fglcov.md "Enables code coverage data collection.") environment variable to a value different from
zero:

```
$ export FGLCOV=1
```

## Merging coverage data files and source files

After producing coverage data files, these must be merged with source files by using the [fglrun --merge-cov](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.")
command:

```
fglrun --merge-cov name.4gl [name.42m.cov]
```

Specify the path to the .42m.cov file, if this file and the corresponding
.42m file are not located in the same directory as the
.4gl source file.

## Interpreting coverage indicators and execution counts

In the .4gl.cov result file, each source line gets a counting indicator,
that must be interpreted as follows:

- `num` (a number): The line was executed num
  times.
- `-` (hyphen-minus): The line is not reachable.
- `=====` : The line is reachable, but has never been executed.
