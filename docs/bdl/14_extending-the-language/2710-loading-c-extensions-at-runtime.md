---
title: "Loading C-Extensions at runtime"
source: "fgl-topics/c_fgl_CExtensions_008.html"
breadcrumb: "Extending the language > C-Extensions > Loading C-Extensions at runtime"
type: "concept"
---

# Loading C-Extensions at runtime

> The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files.

> **Important:**
>
> When running iOS platforms, the C-Extensions are linked statically to the GMI application.

Directories are searched for the C-Extensions libraries based on the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable rules.

If the C-Extension library depends on other shared libraries, make sure that the library loader
of the operating system can find theses shared objects. You may need to set the LD\_LIBRARY\_PATH
environment variable on UNIX™ or the PATH environment variable
on Windows® to point to the directory where these other
libraries are located.

There are three ways to bind a C-Extension with the runtime system:

1. Using the `IMPORT` instruction in sources.
2. Using the default C-Extension name: userextension (.so
   or .DLL)
3. Using the `-e` option of fglrun.

## Using the IMPORT instruction

The `IMPORT` instruction allows you to declare an external module in a
.4gl source file. It must appear at the beginning of the source file.

The name of the module specified after the `IMPORT` keyword is converted to
lowercase by the compiler. Therefore it is recommended to use lowercase file names only.

The compiler and the runtime system automatically know which C-Extensions must be loaded, based
on the `IMPORT`
instruction:

```
IMPORT mylib1
MAIN
  CALL myfunc1("Hello World")  -- C function defined in mylib1
END MAIN
```

When the `IMPORT` instruction is used, no other action has to be taken at runtime.
The module name is stored in the 42m p-code and is automatically loaded when
needed.

## Using the default C-Extension name

It is recommended that all modules using a function from a C-Extension now use the [`IMPORT`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") instruction. However, this can mean
a major change to existing sources.

To simplify migration of existing C-Extensions, the runtime system loads by default a module with
the name `userextension`. Create this shared library with your existing C-Extensions,
and the runtime system will load it automatically if found in the directories specified by
FGLLDPATH.

## Using the -e fglrun option

In some cases you need several C-Extension libraries, which are used by different groups of
programs, so you cannot use the default `userextension` solution. However, you don't
want to review all your sources in order to use the `IMPORT` instruction.

You can specify the C-Extensions to be loaded by using the `-e` option of [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs."). The `-e` option takes a comma-separated
list of module names, and can be specified multiple times in the command line. The following example
loads five extension
modules:

```
fglrun -e myext1,myext2,myext3 -e myext4,myext5 myprog.42r
```

By using the `-e` option, the runtime system loads the modules specified in the
command line instead of loading the default `userextension` module.

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")
