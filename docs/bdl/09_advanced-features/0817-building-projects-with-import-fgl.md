---
title: "Building projects with IMPORT FGL"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_projects.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Building projects with IMPORT FGL"
type: "concept"
---

# Building projects with IMPORT FGL

> Use IMPORT FGL instead of linking.

`IMPORT FGL` instructs the [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")
compiler and [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") runtime system to load/check the
specified modules.

When using only `IMPORT FGL` to define module dependency, there is no longer a
need to [link](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.") programs or use libraries.

With `IMPORT FGL`, the compiler can check the number of parameters and returning
values in functions calls, and the autocompletion in source code editors is improved as it can
suggest all imported symbols.

`IMPORT FGL` can be combined with the `PACKAGE` instruction, to
organize source modules as a tree of module groups. See [Organizing modules in packages](0818-organizing-modules-in-packages.md "Modules to be imported can be grouped in packages.") for more details.

By default, the fglcomp compiler creates the .42m files
in the current working directory, or beside the .4gl source files of modules
imported via packages. Consider using the `--output-dir` option of
fglcomp, if you want to create the .42m files in a different
(runtime) directory as the current development directory. For more details, see [Output directory for .42m pcode files](../13_programming-tools/2534-compiling-program-code-files-4gl.md).
