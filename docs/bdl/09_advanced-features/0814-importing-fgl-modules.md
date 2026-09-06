---
title: "Importing FGL modules"
source: "fgl-topics/c_fgl_programs_fgl_import.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules"
type: "concept"
---

# Importing FGL modules

> FGL modules can be organized by using the IMPORT FGL and PACKAGE instructions.

## Related links

**Related concepts**  

[Program execution](0828-program-execution.md "This section describes program execution and language instructions related to program execution.")

## Child topics

- [IMPORT FGL](0815-import-fgl.md): The IMPORT FGL instruction imports module symbols.
- [PACKAGE](0816-package.md): Defines the package the module belongs to.
- [Building projects with IMPORT FGL](0817-building-projects-with-import-fgl.md): Use IMPORT FGL instead of linking.
- [Organizing modules in packages](0818-organizing-modules-in-packages.md): Modules to be imported can be grouped in packages.
- [Auto-compilation of imported modules](0819-auto-compilation-of-imported-modules.md): Imported local and package modules are compiled automatically if needed.
- [Circular module references](0820-circular-module-references.md): Circular references between imported modules are allowed.
- [Identifying modules to be imported](0821-identifying-modules-to-be-imported.md): Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions.
- [Scope of module symbols (PRIVATE/PUBLIC)](0822-scope-of-module-symbols-private-public.md): The PRIVATE/PUBLIC modifiers can be used to hide / publish symbols to other modules.
- [Using [package.]module prefix](0823-using-package-module-prefix.md): Resolve symbol name conflicts with module prefix.
- [Defining aliases for imported modules](0824-defining-aliases-for-imported-modules.md): Use aliases when package path or module names are too long.
- [Mixing IMPORT FGL and .42r linking](0825-mixing-import-fgl-and-42r-linking.md): Traditional linking is still supported for backward compatibility, and can be mixed with IMPORT FGL method.
