---
title: "GLOBALS variables usage"
source: "fgl-topics/c_fgl_MigI4GL_025.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > GLOBALS variables usage"
type: "concept"
---

# GLOBALS variables usage

> This topic describes differences between I4GL and FGL regading global variables definitions and usage.

## Avoid GLOBALS

> **Important:**
>
> Using `GLOBALS` is strongly discouraged. This feature is supported for backward
> compatibility, to compile legacy source code. Instead of `GLOBALS`, use
> `PUBLIC` symbols in modules to be imported with [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").
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

## Mismatching global variable definitions

The c4gl C-code compiler of IBM® Informix® 4GL has a weakness that allows global variable declarations of
the same variable with different data types. Each different declaration found by the c4gl compiler
defines a distinct global variable, which can be used separately. This can actually be very
confusing (the same global variable name can, for example, reference a `DATE` value
in module A and an `INTEGER` value in module B).

IBM Informix 4GL
RDS (fglpc / fglgo) does not allow multiple global variable declaration with different types. The
fglgo runner raises error -1337 if this happens.

The next code example shows two .4gl modules defining the same global variable with different
data types:

Main.4gl:

```
GLOBALS
  DEFINE v INTEGER
END GLOBALS
MAIN
  LET v = 123
END MAIN
```

Module.4gl:

```
GLOBALS
  DEFINE v DATE
END GLOBALS
FUNCTION test()
  LET v = TODAY
END FUNCTION
```

The fglcomp tool compiles both modules separately without problem, but when linking with
fgllink or fglrun -l, the linker raises error [-1337](../15_library-reference/4483-genero-bdl-errors.md).

You must review your code and use the same data type for all global variables having the same
name.

## Duplicated global variable definition

IBM Informix 4GL
produces a compilation error when same variable and type is defined in a `GLOBALS / END
GLOBALS` section in the current module, and the globals file included with a `GLOBALS
"filename"` instruction that defines the same variable.

Genero BDL accepts the duplicated global variable definition, as long as the types match. This is
supported for backward compatiblity with older Genero and I4GL
versions.

```
GLOBALS
    DEFINE var1 INTEGER
END GLOBALS
```

```
GLOBALS "myglobals.4gl"
GLOBALS
    DEFINE var1 INTEGER
END GLOBALS
MAIN
    DISPLAY "var1 = ", var1
END MAIN
```

As a general advice, avoid the use of global variables and use modules variables with
`PUBLIC` scope.

## Related links

**Related concepts**  

[Declaration context](../08_language-basics/0693-declaration-context.md "A variable can be declared in different contexts, which defines its visibility.")
