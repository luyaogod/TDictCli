---
title: "Auto-compilation of imported modules"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_autocomp.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Auto-compilation of imported modules"
type: "concept"
---

# Auto-compilation of imported modules

> Imported local and package modules are compiled automatically if needed.

When the imported module is located in the same directory as the compiled module or in a
sub-directory following the package paths, fglcomp will automatically compile the
imported module, if the .42m file does not exist, or is older than the
corresponding .4gl source file.

The fglcomp compiler builds the .42m in the current
directory, in the package path directory, or in in the output directory specified by
`--output-dir` option.

The automatic recompilation of imported modules applies recursively: For example, when a
main.4gl module imports `module1` which in turn imports
`module2`, and module2.4gl is more recent as
module2.42m, fglcomp will automatically compile
module2.4gl.

fglcomp uses [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.")
content to search for the .42m modules used by `IMPORT FGL`.
Modules located in other directories and found by FGLLDPATH should already be compiled.

> **Important:**
>
> Auto-compilation of imported modules is only supported if the imported module is in the current
> directory or in a package path. Regular modules located in other directories and found by FGLLDPATH
> must already be compiled.

When using packages, fglcomp can find modules in sub-directories by following
the package-path specified in `IMPORT FGL
package-path.*`. Therefore, FGLLDPATH should not be set during
compilation with packages. FGLLDPATH can be set at runtime to find .42m modules
from a top-dir that is different from the source top-dir.

To avoid implicit compilation of imported modules, use the `--implicit=none`
option of fglcomp. If the .42m file exists but the
.4gl source file cannot be found, fglcomp imports the
.42m file as is.

The fglcomp --implicit=none option is provided for specific cases. Do not use
this option when not really needed: Auto-compilation of imported modules is the recommended
default.
