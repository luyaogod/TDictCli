---
title: "GLOBALS"
source: "fgl-topics/c_fgl_Globals_003.html"
breadcrumb: "Advanced features > Globals > GLOBALS"
type: "concept"
---

# GLOBALS

> The GLOBALS / END GLOBALS block and the GLOBALS instruction.

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

## Syntax 1: Global block declaration

```
GLOBALS
  declaration-statement
  [,...]
END GLOBALS
```

1. declaration-statement is a variable, constant or type
   declaration.

## Syntax 2: Importing definitions from a globals file

```
GLOBALS "filename"
```

1. filename is the file containing the definition of globals.
   filename can be a relative or an absolute path. To specify a path, the slash
   (`/`) directory separator can be used for UNIX™
   and Windows® platforms.
2. Use this syntax to include global declarations in the current module.

## Usage

To extend the scope of [variables](../08_language-basics/0686-variables.md "Explains how to define program variables."), [constants](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values.") or [user
types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") to the whole program, define a module containing a `GLOBALS ... END
GLOBALS` block and include this global module with the `GLOBALS
"filename"` statement in other modules.

Initialization of global variables (`DEFINE x INT = 100`) is not supported.

The filename should contain the `.4gl` suffix. However, the
compiler uses the filename as it is. Therefore it can accept other file extensions such as
`GLOBALS "stock.glb"`.

The filename can be a relative or an absolute path. To specify a path, the
slash (`/`) directory separator can be used for UNIX and Windows platforms.

When the filename specified by `GLOBALS` defines a relative
file path, the globals file will be searched relatively to the location of the source provided to
fglcomp.

For compatibility, fglcomp searches globals files in pwd first, then the
search is relative to the source file path. This can simplify build scripts: the compiler can be
called from any directory.

## Related links

**Related concepts**  

[Variables](../08_language-basics/0686-variables.md "Explains how to define program variables.")

[Constants](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values.")

[Types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
