---
title: "Importing modules"
source: "fgl-topics/c_fgl_programs_IMPORT.html"
breadcrumb: "Advanced features > Importing modules"
type: "concept"
---

# Importing modules

> Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.

The `IMPORT {JAVA|FGL}` instruction can be used to declare
the usage of an external module. All (public) symbols of the external module can be referenced
in the current module.

The `IMPORT {JAVA|FGL}` instruction must be the first
instruction in the current module. If you specify this instruction after `DEFINE`,
`CONSTANT` or `GLOBALS`, fglcomp will report a
syntax error.

The `IMPORT {JAVA|FGL}` instruction can import Genero
modules, Java classs or C extension libraries:

- `IMPORT FGL [package-path.]module-name`: Imports a
  specific Genero module.
- `IMPORT FGL package-path.*`: Imports all modules that belong to this
  Genero package.
- `IMPORT JAVA class-name`: Imports a Java class or class element.
- `IMPORT lib-name`: Imports a C extension implementing functions and
  variables.

The name specified after the `IMPORT FGL` or `IMPORT JAVA`
instruction is case-sensitive; program module (.4gl) or Java class must exactly
match the file name.

> **Tip:**
>
> The `IMPORT FGL` instruction is case sensitive, but some files
> systems such as Microsoft™ Windows® NTFS are not case sensitive. Consider using lower-case
> directory/package names and module/filenames, to avoid issues, especially when the development
> platform is different from the deployment platform: `IMPORT FGL mymodule` of
> MyModule.4gl works on Windows, but fails on UNIX/Linux.

For backward compatibility, C extension library names specified after `IMPORT` are
converted to lowercase by the compiler (therefore, we recommend you to use lowercase file names for
C extensions). A character case mismatch will be detected on UNIX™ platforms, but not on Windows where the file
system is not case-sensitive. Regarding the usage of imported symbols in the rest of the code (not
in the `IMPORT` instruction): C extensions and Genero symbols are case-insensitive,
while Java symbols are case-sensitive.

## Child topics

- [Importing FGL modules](0814-importing-fgl-modules.md): FGL modules can be organized by using the IMPORT FGL and PACKAGE instructions.
- [IMPORT JAVA](0826-import-java.md): The IMPORT JAVA instruction imports Java module elements.
- [IMPORT (C-Extension)](0827-import-c-extension.md): The IMPORT instruction imports c extension module elements to be used by the current module.
