---
title: "FGLLDPATH"
source: "fgl-topics/c_fgl_EnvVariables_FGLLDPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLLDPATH"
type: "concept"
---

# FGLLDPATH

> Defines a list of paths to find program modules.

The FGLLDPATH environment variable defines the search paths to find FGL modules and C
extensions.

When compiling (`IMPORT FGL`), linking and when executing the program, the
compiler and the runtime system must known where to search for
.4gl/.42m (and C extension) modules. Use the FGLLDPATH
environment variable, to define the search paths to find program modules.

When using packages, FGLLDPATH can define the top-dir to find modules from the
package paths. See [PACKAGE](../09_advanced-features/0816-package.md "Defines the package the module belongs to.") for more details.

FGLLDPATH must contain a list of paths, separated by the operating system specific path separator.
The path separator is **":"** on UNIX™ platforms and
**";"** on Windows® platforms.

The directories are searched in the following order:

1. The current working directory.
2. The directory where the program file resides (the .42m module containing
   `MAIN` or the .42r program file).
3. A path defined in the FGLLDPATH environment variable.
4. The $FGLDIR/lib directory.

FGLLDPATH is also used by the debugger to find program sources. For more details, see [FGLSOURCEPATH](0533-fglsourcepath.md "Defines a list of paths to program source files.").

## Related links

**Related concepts**  

[FGLDIR](0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")

[Importing modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
