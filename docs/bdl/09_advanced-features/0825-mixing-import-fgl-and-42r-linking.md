---
title: "Mixing IMPORT FGL and .42r linking"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_import_and_link.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Mixing IMPORT FGL and .42r linking"
type: "concept"
---

# Mixing IMPORT FGL and .42r linking

> Traditional linking is still supported for backward compatibility, and can be mixed with IMPORT FGL method.

To ease migration from traditional linking to imported modules, you can mix `IMPORT
FGL` usage with fgllink/fglrun -l.

By default, even when `IMPORT FGL` is used, fglcomp does not
raise an error, if a referenced function is not found in the imported modules. This is mandatory to
compile the 42m file to be linked later with the module defining the missing
function.

Use the fglcomp `-W implicit` or the
`--resolve-calls` options, to check that all symbols are resolved with a
corresponding `IMPORT FGL` instruction.

When the `-W implicit` option is used , fglcomp will print
warning [-8406](../15_library-reference/4483-genero-bdl-errors.md) for any
referenced function that cannot be found in an imported module.

The `-W implicit` option is to be used when migrating linked modules to a solution
where module dependency is only based on `IMPORT FGL`.

To enable strict symbol resolution by the compiler, use the `--resolve-calls`
option. This option will force the compiler to check all function symbols referenced in a module,
and raise error [-8406](../15_library-reference/4483-genero-bdl-errors.md), if a
symbol is not found in the imported modules.

The `--resolve-calls` option should be used to compile programs that are only
based on `IMPORT FGL` and no longer use the link phase.

For more details about the linker, see [Linking programs](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.").
