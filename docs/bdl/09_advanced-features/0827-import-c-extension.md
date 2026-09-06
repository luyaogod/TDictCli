---
title: "IMPORT (C-Extension)"
source: "fgl-topics/c_fgl_programs_IMPORT_CEXT.html"
breadcrumb: "Advanced features > Importing modules > IMPORT (C-Extension)"
type: "concept"
---

# IMPORT (C-Extension)

> The IMPORT instruction imports c extension module elements to be used by the current module.

## Syntax

```
IMPORT cextname
```

1. cextname is an [identifier](../08_language-basics/0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.") defining the C extension module to be imported (without the file extension).

## Usage

Using `IMPORT cextname` instructs the compiler and runtime
system to use the cextname C extension for the current module.

At runtime, all imported C extension modules are loaded when the program starts.

The name of the module
specified after the `IMPORT` keyword is converted to
lowercase by the compiler. Therefore it is recommended to
use lowercase file names only.

The C extension must exist as a shared library (.DLL or
.so) and be loadable (environment
variables must be set properly). C extension modules used with the
`IMPORT` instruction do not have to be linked
to fglrun: the runtime system loads dependent C
extension modules dynamically.

The [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable
specifies the directories to search for the C extension modules. You may also have to set up the
system environment properly (such as PATH on Windows® and
LD\_LIBRARY\_PATH on UNIX™) if the C extension library is
dependent on other libraries.

By default, the runtime system tries to load a C extension module with the name
`userextension`, if it exists. This
simplifies the migration of existing C extensions; you just need to
create a shared library named userextension.so
(or userextension.dll on Windows), and copy the file to one
of the directories defined in FGLLDPATH.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
