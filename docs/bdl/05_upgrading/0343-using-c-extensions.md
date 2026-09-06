---
title: "Using C extensions"
source: "fgl-topics/c_fgl_MigI4GL_006.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > Installation and setup topics > Using C extensions"
type: "concept"
---

# Using C extensions

> When migrating to Genero BDL, you must provide your C-Extensions as shared libraries.

With IBM® Informix® 4GL, you can extend the fglgo runtime executable or link your
binary programs with c4gl by adding your own C functions.

When migrating to Genero Business Development Language, the C-Extensions must be reviewed in
order to provide them as shared libraries. Normally, C extensions modules must be specified in .4gl
modules with the IMPORT instruction. To simplify migration, the runtime system loads the
userextension shared library (or DLL) automatically, so you can group all your
existing C functions in a unique shared library and use it without changing the source code of your
programs.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
