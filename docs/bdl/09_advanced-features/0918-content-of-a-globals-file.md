---
title: "Content of a globals file"
source: "fgl-topics/c_fgl_Globals_006.html"
breadcrumb: "Advanced features > Globals > Content of a globals file"
type: "concept"
---

# Content of a globals file

> A globals file contains a GLOBALS ... END GLOBALS block.

As a `GLOBALS` block can also be defined in regular modules, it is
possible to include a source containing more than a `GLOBALS` block. When
including such module, the sections before and after the `GLOBALS` block
are ignored by the compiler. The source defining the global elements can be compiled
individually.

For example, it is allowed to define a module A with a `GLOBALS ... END GLOBALS`
block, followed by [function](../08_language-basics/0761-functions.md "Describes user defined functions.") definitions. This module
can be compiled and functions will be taken into account. Module A can then be included in module B
with a `GLOBALS "filename"` instruction, and when compiling module
B the function definitions of the included module A will be ignored. [`IMPORT`](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") instructions before the a
`GLOBALS ... END GLOBALS` block will also be ignored in such case.
