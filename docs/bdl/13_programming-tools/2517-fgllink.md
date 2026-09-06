---
title: "fgllink"
source: "fgl-topics/c_fgl_tools_fgllink.html"
breadcrumb: "Programming tools > Command reference > fgllink"
type: "concept"
---

# fgllink

> The fgllink tool assembles p-code modules produced with fglcomp into a .42r program or a .42x library.

## Syntax 1: Create a 42x library

```
fgllink [options] -o outfile.42x file-list
```

1. In this form, fgllink links a .42x library.
2. options are described in [fglrun
   -l linking options](2513-fglrun.md).
3. outfile.42x is the name of the library to be
   created.

where file-list is:

```
{ module.42m
| pattern
| @argfile
} [...]
```

1. module.42m is a p-code module
   compiled with fglcomp.
2. pattern is a `MATCHES`-style pattern to
   find files, like `'[a-z]*.42m'`.
3. argfile defines a file that contains a list of .42m
   files. Each line must specify a filename or a pattern.

## Syntax 2: Create a 42r program

```
fgllink [options] -o outfile.42r file-list
```

1. In this form, fgllink links a .42r program.
2. options are described in [fglrun
   -l linking options](2513-fglrun.md).
3. outfile.42r is the name of the program to be
   created.

where file-list
is:

```
{ { module.42m | library.42x }
| pattern
| @argfile
} [...]
```

1. module.42m is a p-code module
   compiled with fglcomp.
2. library.42x is a name of a library to be used for
   linking.
3. pattern is a `MATCHES`-style pattern to
   find files, like `'[a-z]*.42m'`.
4. argfile defines a file that contains a list of
   .42m or .42x files. Each line must specify a filename or a
   pattern.

## Syntax 3: Information options

```
fgllink info-option
```

1. info-option can be any of the informational options.

## Options

| Option | Description |
| --- | --- |
| `-V` or `--version` | Displays version information. |
| `-h` or `--help` | Displays options for the tool. |

## Usage

The fgllink command line tool links .42m p-code modules together to
create a .42x library or a .42r program file.

```
fgllink -o myprog.42x module1.42m module2.42m lib1.42x
```

Note that fgllink is a wrapper calling [fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") with the `-l` option.

If case of error, the fgllink command execution status is different from zero.
Consequently, it is possible to detect compilation errors in scripts and makefiles.

## Related links

**Related concepts**  

[Linking libraries](2536-linking-libraries.md "Describes how to link .42m modules together to build a .42x library file.")

[Linking programs](2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.")
