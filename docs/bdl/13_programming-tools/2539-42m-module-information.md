---
title: "42m module information"
source: "fgl-topics/c_fgl_CompilingPrograms_007.html"
breadcrumb: "Programming tools > Compiling source files > 42m module information"
type: "concept"
---

# 42m module information

> Describes how to handle module information in .42m p-code files.

## Compiler and runtime compatibility

The runtime system ([fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.")) used
to execute programs must be compatible with the [fglcomp](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler version used to build the .42m
modules.

To control this compatibility, the compiler writes Genero and p-code version information in the
generated 42m files, in the p-code header.

The compiler also writes additional information in the p-code header, such as the source filename
used at compilation, a timestamp, or a user-defined string/tag.

The p-code header information of a 42m file can be read with the fglrun -b
module.42m command. This is useful on production sites, to check the version of the
compiler that was used to build the 42m nodules.

It is also possible to get additional information about the p-code module such as the size used
in memory with fglrun --module-size.

## Identifying the version of the compiler and runtime system

To check if the version of the compiler or runtime system, use the `-V`
option:

```
$ fglcomp -V
fglcomp 4.00.03 internal-build-number
Genero 4gl compiler
Target l64xl217
...

$ fglrun -V
fglrun 4.00.03 internal-build-number
Genero virtual machine
Target l64xl217
...
```

## Extracting p-code header information

To extract build information from a .42m file, run fglrun
with the `-b` option:

```
$ fglrun -b mymodule.42m
3.10.12 /home/devel/stores/mymodule.4gl 24
```

The p-code header contains the following fields:

1. An optional timestamp, if the `--timestamp` option was used with
   fglcomp.
2. The Genero product version.
3. The full path of the source file, or the module name, if the `--omit-source-name`
   option was used with fglcomp.
4. The internal identifier of the p-code version.
5. An optional user-defined tag, if the `--tag=string` option was
   used with fglcomp.

## Avoiding the full source file path in the p-code header

By default, fglcomp writes the source filename (full path) in the resulting
.42m
module:

```
$ fglcomp mymodule.4gl

$ fglrun -b mymodule.42m
3.10.12 /home/devel/stores/mymodule.4gl 24
```

To avoid the source name in the .42m module, use the
`--omit-source-name` option of fglcomp, to write only the module
name in the p-code header:

```
$ fglcomp --omit-source-name mymodule.4gl

$ fglrun -b mymodule.42m
3.10.12 mymodule 24
```

## Writing a custom string to the p-code header

Use the `--tag="custom-string"` option of fglcomp, to
add a user-defined string to the p-code header. When reading p-code header information with
fglrun -b, the user-defined string is printed as
`tag="custom-string"`.

The tag field can for example be used to stamp the 42m file with the custom product
version:

```
$ fglcomp --tag="OXOGEN 5.23" mymodule.4gl

$ fglrun -b mymodule.42m
3.10.12 /home/devel/stores/mymodule.4gl 24 tag="OXOGEN 5.23"
```

## Writing a compilation timestamp to the p-code header

To write timestamp information in the p-code header, use the `--timestamp` option
of fglcomp:

```
$ fglcomp --timestamp mymodule.4gl

$ fglrun -b mymodule.42m
2008-12-24 11:22:33 3.10.12 /home/devel/stores/mymodule.4gl 24
```

> **Important:**
>
> When using the `--timestamp` compiler option to write
> build timestamp information in p-code modules, you will not be able to easily compare 42m files
> (based on a checksum, for example). Without the timestamp, fglcomp generates
> exactly the same p-code module as if the source file was not modified.

## Reading p-code header of older versions

fglrun -b can read the header of p-code modules compiled with older versions
of fglcomp and display version information for such old modules.

If fglrun cannot recognize a p-code module, it returns an execution status
that is different from zero.

## Reading p-code header of 42x and 42r files

When reading build information of a 42x or 42r file, fglrun scans all modules
used to build the library or program. You will see different versions in the first column if the
modules were compiled with different versions of fglcomp. However, it's not
recommended that you mix versions on a production site:

```
$ fglrun -b myprogram.42r
3.10.11 /home/devel/stores/mymodule1.4gl 24
3.10.02 /home/devel/stores/mymodule2.4gl 24
3.10.12 /home/devel/stores/mymodule3.4gl 24
```

## Computing the p-code size of a module or program

The fglrun command provides the `--module-size` option to
compute the p-code size of a module:

```
$ fglrun --module-size mymodule.42m
 12.34K mymodule
```

This size is defined by the amount of p-code instructions of the module.

P-code modules are shared by all fglrun instances executing on the same
computer. For more details, see [Elements shared by multiple programs](../09_advanced-features/0969-elements-shared-by-multiple-programs.md).

The `--program-size` option reports the total p-code size of all modules of a
program. The argument must be a .42m module containing the
`MAIN` (with main module importing other modules by `IMPORT
FGL`):

```
$ fglrun --module-size myprogram.42m
  0.34K module1
  2.59K module2
  ...
12.76K total
```

The `--program-size` option can also use a .42r program file
as argument, for linked programs:

```
$ fglrun --module-size myprogram.42r
  0.34K module1
  2.59K module2
  ...
12.76K total
```
