---
title: "Linking programs"
source: "fgl-topics/c_fgl_CompilingPrograms_005.html"
breadcrumb: "Programming tools > Compiling source files > Linking programs"
type: "concept"
---

# Linking programs

> Describes how to link .42m modules together to build a .42r program file.

## Purpose of .42r program files

Traditional Genero programming allows you to call a function without seeing the definition of
that function. The goal of linking programs is to resolve function symbols and build a
.42r program file.

When writing new applications, consider using [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") instead of traditional linking. When using `IMPORT
FGL`, the link stage is no longer required, and the 42m module containing the
`MAIN` block can be directly executed.

Genero .42r program files are created by linking several
.42m modules and/or .42x [libraries](2536-linking-libraries.md "Describes how to link .42m modules together to build a .42x library file.") together, where one of the modules defines
a `MAIN` block.

By convention, the resulting program file gets the .42r extension.

Program linking is done with the [fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") tool by using
the `-l` option.

The following lines show a link procedure to create a program in a UNIX™ shell session:

```
$ fglcomp main.4gl
$ fglcomp store.4gl
$ fgllink -o stores.42r main.42m store.42m
```

> **Important:**
>
> If you omit the `-o` option in the fgllink
> command, the default output file will get the .42x extension (used for
> libraries), with the name of the module containing the `MAIN` block. The
> .42r file extension is used by convention, to distinguish a program dictionary
> file from a library dictionary file.

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
main.42m
module1.42m
module2.42m
lib1.42x
$ fgllink -o prog.42x @myfiles.txt
```

## Symbol resolution with the linker

The purpose of the linking phase is to check for missing function symbols, and reference all the
symbols in the resulting .42r program file.

The `IMPORT FGL` method is the preferred method for linking. However, you can mix
both methods. The job of the linker is to resolve symbols that are not already resolved by the
compiler from `IMPORT FGL` usage.

Any function used in the .42m modules specified in the link line must be
provided. Missing symbols will result in a [-1338](../15_library-reference/4483-genero-bdl-errors.md) linker error:

```
$ cat main.4gl
MAIN
  CALL myfunc()
END MAIN

$ fglcomp main.4gl
$ fgllink -o prog.42r main.42m
ERROR(-1338):The function 'myfunc' has not been defined in any module in the program.
```

Symbol resolution in only be done when linking programs. When linking a 42x
library, there can be references to undefined functions.

When linking a 42r program, global symbols must be unique; otherwise, error [-6203](../15_library-reference/4483-genero-bdl-errors.md) will be returned by the
linker. The same error will be returned when linking a 42x library by using modules defining the
same functions.

The link process searches recursively for the functions used by the program. For example, if the
`MAIN` block calls function FA in module MA, and FA calls FB in module MB, all
functions from module MA and MB will be included in the 42r program definition.

## The linking process steps in detail

When linking a .42r program (with modules possibly using [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.")), the linker works as
follows:

1. Loads all modules specified on the command-line. Builds a symbol-table of all public functions.
   Raises error [-6203](../15_library-reference/4483-genero-bdl-errors.md), if a
   public function is defined more than once.
2. Loads all imported modules (if not specified on the command line). Adds public functions to the
   symbol-table.
   1. If an imported function is already defined by a linked module, the imported function will be
      ignored.
   2. If an imported function is not defined by a linked module, and is defined more than once, then
      this function is marked as "ambiguous" (this is legal, no warning or error is produced).
3. All unresolved symbols on any module (linked modules and imported modules) will be resolved.
   - If a function is not defined, error [-1338](../15_library-reference/4483-genero-bdl-errors.md) is thrown.
   - If a function is an ambiguous function (2.b), error [-8401](../15_library-reference/4483-genero-bdl-errors.md) is thrown.

Understand the difference in case of error [-6203](../15_library-reference/4483-genero-bdl-errors.md) and [-8401](../15_library-reference/4483-genero-bdl-errors.md):

- Error -6203: A function in a linked module is defined more then once. That's illegal.
- Error -8401: A function in an imported module is define more then once, and used by the
  linker to resolve a symbol. It's legal to define an imported function more than once, but this
  function can not be used by the linker.

## Content of a .42r program file

The generated 42r program files do not contain the 42m p-code, it is basically a dictionary of
global symbols used by the program.

When deploying an application, both 42m modules and 42r
program files must be provided.

Since 42x library files are only used to build programs, you do not have to deploy 42x library
files.

When linking a 42r program by using 42x libraries, the modules defined in a library are included
only if one of the symbols in the module is used by the program. However, all symbols of 42m modules
specified in the command line will always be referenced in the resulting 42r program file.

During the link, when the same function symbols are defined in distinct libraries, the linker
will select the function of the first library that was specified in the command line.

All symbols referenced in a module must exist in the final 42r program dictionary file. If a
symbol is not found, the runtime system stops with error [-1338](../15_library-reference/4483-genero-bdl-errors.md). This error is fatal and
cannot be trapped with an exception handler.

## P-Code module find path FGLLDPATH

If you do not specify an absolute path for a file, the linker searches by default for 42m modules
and 42x libraries in the current directory.

You can specify the 42m search path with the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment
variable:

```
$ FGLLDPATH=/usr/dev/lib/maths:/usr/dev/lib/utils
$ export FGLLDPATH
$ ls /usr/dev/lib/maths
mathlib1.42x
mathlib2.42x
mathmodule11.42m
mathmodule12.42m
mathmodule22.42m
$ ls /usr/dev/lib/utils
fileutils.42m
userutils.42m
dbutils.42m
$ fgllink -o myprog.42r mymodule.42m mathlib1.42x fileutils.42m
```

In this example the linker will find the specified files in the
/usr/dev/lib/maths and /usr/dev/lib/utils directories
defined in FGLLDPATH.

## Library versus module function precedence

When creating a .42r program by linking .42m modules
with .42x libraries, if the same function is defined in a 42m and in a module
of a 42x library, the function of the specified 42m module will be selected by the linker, and the
function of the library will be ignored.

However, the linker will raise error [-6203](../15_library-reference/4483-genero-bdl-errors.md), if two 42m modules specified
in the link command define the same function.

## Exclusion of unused library module

When linking a .42r program by using a .42x library, if
none of the functions of a module in the .42x library are used in the program,
the complete module is excluded by the linker.

Unused module exclusion may cause undefined function errors at runtime, such as when a function
is only used in a dynamic call (an initialization function, for example.)

The following case illustrates this behavior:

```
$ cat module1.4gl
FUNCTION func11()
END FUNCTION

$ cat module2.4gl
FUNCTION func21()
END FUNCTION

$ cat main.4gl
MAIN
  CALL func11()  -- from module1.4gl
END MAIN

$ fglcomp module1.4gl
$ fglcomp module2.4gl
$ fglcomp main.4gl

$ fgllink -o lib.42x module1.42m module2.42m

$ fgllink -o prog.42r main.42m lib.42x  -- Only module1.42m is included in .42r
```

Here, `module1.42m` with functions `func11` will be referenced in
the prog.42r program file, but module `module2.42m` will not. At
runtime, any dynamic call to function `func21` will fail with an untrappable error
[-1338](../15_library-reference/4483-genero-bdl-errors.md).

## Symbol conflicts with IMPORT FGL and linking

The job of the linker is to resolve symbols that are not already solved by the compiler from
[`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") usage.

If the same function is defined by a linked module and an imported module, and the function is
called without the module prefix, the imported function takes precedence over the linked module
function.

The following example illustrates this case:

- `module1` defines `func11`, so does `module3`.
- The link command includes `main`, `module2` and
  `module3`.
- `module2` imports `module1`, and calls `func11`
  with no prefix, resolving to `module1.func11`.
- `main` does not import any module, and calls `func11` with no
  prefix. The linker will resolve `func11` to `module3.func11`, because
  `module3` is used in the link command.

```
$ cat module1.4gl
FUNCTION func11()
    DISPLAY "module1.func11()"
END FUNCTION

$ cat module2.4gl
IMPORT FGL module1
FUNCTION func21()
    DISPLAY "module2.func21()"
    CALL func11() -- from module1, because it is imported
END FUNCTION

$ cat module3.4gl
FUNCTION func11() -- Same name as in module1.4gl
    DISPLAY "module3.func11()"
END FUNCTION

$ cat main.4gl
MAIN
    CALL func11() -- from module3, because it is linked
    CALL func21()
END MAIN

$ fglcomp module1.4gl
$ fglcomp module2.4gl
$ fglcomp module3.4gl
$ fglcomp main.4gl
$ fgllink -o prog.42r main.42m module2.42m module3.42m
$ fglrun prog.42r
module3.func11()
module2.func21()
module1.func11()
```

To make your code more readable, consider prefixing imported symbols with the module name and
thus avoid any ambiguity.

## Imported module using linked functions

When linking a program with modules using the [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") instruction, the imported modules do not have to be specified in
the link line.

However, if the imported module uses functions that come from other modules, which are not
imported by this module, these non-imported modules must be specified in the link command line.

For example, if the main module imports module `module1` to call the
`func11()`, which in turn calls `func21()` from
`module2`, but `module1` does not import `module2`,
then module2 must be linked to the program:

```
$ cat module1.4gl
FUNCTION func11()
    DISPLAY "module1.func11()"
    CALL func21()
END FUNCTION

$ cat module2.4gl
FUNCTION func21()
    DISPLAY "module2.func21()"
END FUNCTION

$ cat main.4gl
IMPORT FGL module1
MAIN
    CALL module1.func11()
END MAIN

$ fglcomp module1.42m
$ fglcomp module2.42m
$ fglcomp main.4gl

$ fglrun main.42m
module1.func11()
Program stopped at 'module1.4gl', line number 3.
FORMS statement error number -1338.
The function 'func21' has not been defined in any module in the program.

$ fgllink -o prog.42r main.42m module2.42m

$ fglrun prog.42r
module1.func11()
module2.func21()
```

## Linking programs with C Extensions

If you are using [C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code."), you may need to use
the `-e` option to specify the list of extension modules if the
`IMPORT` keyword is not used:

```
$ fgllink -e extlib,extlib2,extlib3 -o stores.42r main.42m store.42m
```

## Related links

**Related concepts**  

[Linking programs using C-Extensions](../14_extending-the-language/2709-linking-programs-using-c-extensions.md "When creating a 42r program or 42x library, the linker needs to resolve all function names, including C-Extension functions.")
