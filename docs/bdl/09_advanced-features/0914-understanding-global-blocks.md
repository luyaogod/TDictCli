---
title: "Understanding global blocks"
source: "fgl-topics/c_fgl_Globals_002.html"
breadcrumb: "Advanced features > Globals > Understanding global blocks"
type: "concept"
---

# Understanding global blocks

> Global symbols can be defined with the GLOBALS instruction

The `GLOBALS` instruction can be used to declare variables, constants and types
for the whole program.

> **Important:**
>
> Using `GLOBALS` is strongly discouraged. This feature is supported for backward
> compatibility, to compile legacy source code. Instead of `GLOBALS`, use
> `PUBLIC` symbols in modules to be imported with [`IMPORT FGL`](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").
>
> Why not globals? Global variables visible and modifiable across the entire program code is an old
> concept that reduces modularity and flexibility. `GLOBALS` implements specific
> semantics that restrain source code organization; for example, it is allowed to specify a file path
> such as `GLOBALS "../../common/myglobals.4gl"`. The globals file can define functions
> after the `GLOBALS` block that are ignored when including such file.
>
> Conversely, the `IMPORT FGL` instruction brings more flexibility and allows a
> better organization of your source code, by grouping public variables in modules where other symbols
> (functions, types, contants) belong to the same domain.

## Related links

**Related concepts**  

[Importing modules](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
