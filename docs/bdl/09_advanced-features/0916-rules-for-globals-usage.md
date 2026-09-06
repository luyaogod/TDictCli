---
title: "Rules for globals usage"
source: "fgl-topics/c_fgl_Globals_004.html"
breadcrumb: "Advanced features > Globals > Rules for globals usage"
type: "concept"
---

# Rules for globals usage

> Follow the rules described in this topic in order to use globals properly.

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

Rules for `GLOBALS` usage:

1. If you modify the globals file, you must recompile all the modules that include the file.
2. If a local variable has the same name as another variable that you declare in the
   `GLOBALS` statement, only the local variable is visible within its scope of
   reference.
3. You can declare several `GLOBALS .. END GLOBALS` blocks in the same module.
4. A source file with `GLOBALS .. END GLOBALS` block must not contain any executable
   statement: These will be ignored when including the global file into another module.
5. Do not write a declaration statement outside a `GLOBALS ... END GLOBALS` block in
   the globals file.
6. You do not need to compile the source file containing the `GLOBALS` block.
   However, it is recommended to compile the globals file to detect errors.
7. You can declare several `GLOBALS "filename"` instructions in
   the same globals using module.
8. Although you can include multiple `GLOBALS ... END GLOBALS` statements in the
   same application, do not declare the same identifier within more than one `GLOBALS`
   declaration. Even if several declarations of a global elements defined in multiple places are
   identical, declaring any global element more than once can result in compilation errors or
   unpredictable runtime behavior.
9. A `GLOBALS .. END GLOBALS` block can hold `GLOBALS
   "filename"` instructions. In such case, the specified files will be
   included recursively.
