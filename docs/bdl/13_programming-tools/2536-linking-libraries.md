---
title: "Linking libraries"
source: "fgl-topics/c_fgl_CompilingPrograms_004.html"
breadcrumb: "Programming tools > Compiling source files > Linking libraries"
type: "concept"
---

# Linking libraries

> Describes how to link .42m modules together to build a .42x library file.

## Grouping .42m modules in .42x libraries

Compiled .42m modules can be grouped in libraries using the [fgllink](2517-fgllink.md "The fgllink tool assembles p-code modules produced with fglcomp into a .42r program or a .42x library.") linker. The library file gets the .42x
extension.

The linker can be used to create .42x libraries or .42r
program files. If none of the modules provided to the linker defines the `MAIN`
block, the linker creates a library file; if a `MAIN` block is present, the linker
creates a program file. Make sure to use the correct file extension.

Linking is supported for backward compatibility, it is recommended that you use [`IMPORT FGL`](2535-importing-modules.md "Describes how to define module interdependence with IMPORT FGL.") instead.

Library linking is done with the fgllink tool, or with the [fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") tool by using the `-l` option. The [fgllink](2517-fgllink.md "The fgllink tool assembles p-code modules produced with fglcomp into a .42r program or a .42x library.") tool is a wrapper script for
fglrun -l .

The following lines show a link procedure to create a library in a UNIX™ shell session:

```
$ fglcomp fileutils.4gl
$ fglcomp userutils.4gl
$ fgllink -o libutils.42x fileutils.42m userutils.42m
```

When you create a library, all functions of the 42m modules used in the link command are
registered in the 42x file.

> **Important:**
>
> The 42x library file does not contain the 42m p-code.
> When deploying your application, you must provide all compiled 42m modules.

When creating a 42x library, all functions must be uniquely defined; otherwise, error [-6203](../15_library-reference/4483-genero-bdl-errors.md) will be returned by the
linker.

## Providing the files to link in an arguments file

The fgllink linker supports the `@argfile`
argument, to provide a file that contains the list of .42m modules and
.42x libraries to be used for the link. This can be used when it is not
possible to pass all files in the command line.

Only link files must be specified in the arguments file. Linker options must be provided in the
command line.

The argument file must contain one file per line:

```
$ cat myfiles.txt
module1.42m
module2.42m
$ fgllink -o lib1.42x @myfiles.txt
```

## Using libraries when linking programs

The 42x libraries are typically used to link the final 42r programs:

```
$ fglcomp mymain.4gl
$ fgllink -o myprog.42r mymain.42m libutils.42x
```

The 42r programs must be re-linked, if the content of the 42x libraries changes.

In this example, if a function of the userutils.4gl source file was removed,
you must recompile userutils.4gl, re-link the libutils.42x
library and re-link the myprog.42r program.

## Linking libraries with other libraries

It is possible to create a library by referencing other 42x library files in the link command, as
long as 42m modules can be found:

```
$ fglcomp module_1.4gl
$ fglcomp module_2.4gl
$ fgllink -o lib_A.42x module_1.42m
$ fgllink -o lib_B.42x module_2.42m lib_A.42x
$ fgllink -o myprog.42r lib_B.42x 
  -- will hold functions of module_1 and module_2.
```

## P-Code module find path FGLLDPATH

If you do not specify an absolute path for a file, the linker searches by default for 42m modules
and 42x libraries in the current directory.

If the 42m modules are not in the current directory, you can specify the 42m module search path
with the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable.

## Linking libraries when using C Extensions

If you are using [C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code."), you may need to use
the `-e` option to specify the list of extension modules, if the
`IMPORT` keyword is not used:

```
$ fgllink -e extlib,extlib2,extlib3 -o libutils.42x fileutils.42m userutils.42m
```

## Related links

**Related concepts**  

[Compiling program code files (.4gl)](2534-compiling-program-code-files-4gl.md "The .4gl source files must be compiled to .42m p-code files, in order to be loaded by the runtime system.")
