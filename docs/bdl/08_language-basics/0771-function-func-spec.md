---
title: "FUNCTION func-spec"
source: "fgl-topics/c_fgl_Functions_FUNCTION_selector.html"
breadcrumb: "Language basics > Functions > FUNCTION func-spec"
type: "concept"
---

# FUNCTION func-spec

> The FUNCTION keyword provides the reference to the specified function.

## Syntax

```
FUNCTION [module-name.]function-name
```

1. module-name is the name of an imported module.
2. function-name is the name of a function defining in the current module or in
   an imported module.

## Usage

Inside a `MAIN` or `FUNCTION` block, the `FUNCTION`
keyword instructs the compiler to use the next symbol as the name of a function, rather than the
name of a variable (the language allows the declaration of variables and functions with the same
name in the same module).

The function specification following the `FUNCTION` keyword can be a single
function name or a function name prefixed by a [module
name](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").

The return-types of the referenced functions must be known: If the referenced function returns
one or more values, the function type and the function itself must be defined with a [`RETURNS` clause](0763-function-definitions.md). Otherwise, the compiler will produce the error [-8419](../15_library-reference/4483-genero-bdl-errors.md).

A `FUNCTION func-spec` expression is typically used to assign a
variable defined with a [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
referencing a function. It can also be used as parameter in a function call, but it cannot be
combined with other expressions.

> **Important:**
>
> In order to assign a function reference to a variable, the variable must have
> been defined with a function type that matches the referenced function: The function parameter
> names and types, as well as the return types must be the same. If the
> signatures of the function type and function reference do not match, the compiler produces the error
> [-6631](../15_library-reference/4483-genero-bdl-errors.md).

## Example

```
IMPORT FGL mymodule

TYPE callback_function FUNCTION(p1 INT, p2 INT) RETURNS INT

FUNCTION add(p1 INT, p2 INT) RETURNS INT
    RETURN p1 + p2
END FUNCTION

...
    DEFINE v callback_function
    LET v = FUNCTION add  -- Assign function reference to the variable
...
    CALL process( FUNCTION add, ... ) -- Function reference passed as parameter
...
    LET v = FUNCTION mymodule.sub  -- Using a module prefix
```

## Related links

**Related concepts**  

[Function references](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
