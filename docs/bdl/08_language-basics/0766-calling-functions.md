---
title: "Calling functions"
source: "fgl-topics/c_fgl_Functions_calls.html"
breadcrumb: "Language basics > Functions > Calling functions"
type: "concept"
---

# Calling functions

> Functions can be invoked, to execute the code they define.

## How can functions be invoked?

A function can be invoked in different ways:

1. with the `CALL` instruction,
2. in an expression,
3. in a callback mechanism.

The symbol used to identify the function to be called can be a static function name, or a
variable referencing a function.

## Typical function invocation with `CALL`

In most cases, functions are invoked with the [`CALL`](0675-call.md "The CALL instruction invokes a specified function or method.")
instruction:

```
FUNCTION open_database(dbname STRING)
   ...
END FUNCTION

...
    CALL open_database("stock")
```

## Function invoked as part of an expression

When a function returns a single value, it can be invoked in an [expressions](0592-expressions.md "Shows the possible expressions supported in the language."):

```
MAIN
    DEFINE r INTEGER
    LET r = 500 + add(50,5)
END MAIN

FUNCTION add(x,y) RETURNS INTEGER
    DEFINE x, y INTEGER
    RETURN (x + y)
END FUNCTION
```

## Callback mechanisms

Genero BDL provides some callback mechanisms where functions are invoked when needed, for
example:

- the [`WHENEVER ERROR CALL`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.")
  instruction,
- the [`INITIALIZER`](../11_user-interface/1791-initializer-attribute.md "The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item.") form
  field attribute,
- the [`ui.Form.setDefaultInitializerFunction()`](../15_library-reference/3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.") method,
- ...

## Recursion

A function can invoke itself recursively:

```
MAIN
    CALL recursive(1)
END MAIN

FUNCTION recursive(x)
    DEFINE x INTEGER
    DISPLAY "x = ", x
    IF x<10 THEN
       CALL recursive(x+1)
    END IF
END FUNCTION
```

> **Important:**
>
> Each time a function calls itself, parameters are pushed on the [stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions."). A deep level of recursion can result in out of memory
> errors.

## Calling a function from its reference variable

A function can be [referenced](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") by a variable,
and be invoked through this variable in a `CALL` instruction, or in an
expression:

```
TYPE t_func_ref FUNCTION (p1 INT, p2 INT) RETURNS INT
DEFINE fr t_func_ref
LET fr = FUNCTION add   -- Function with the same signature as t_func_ref
DISPLAY fr(100,200)
```

## Naming parameters in a function call

To improve code readability, function parameters can be qualified with the name specified in the
[function
definition](0763-function-definitions.md "A FUNCTION definition defines a named procedure with a set of statements."):

```
MAIN
    CALL cleanup( mode: "full", verbose: TRUE )
END MAIN

FUNCTION cleanup( mode STRING, verbose BOOLEAN )
    IF verbose THEN
       DISPLAY "Cleanup mode: ", mode
    END IF
END FUNCTION
```

Built-in functions and class/object methods can also be invoked with named
parameters:

```
DEFINE s STRING
DISPLAY s.subString(startIndex: 1, endIndex: 3)
...
DEFINE io base.Channel
LET io = base.Channel.create()
CALL io.openClientSocket(host: "localhost", port: 4711, mode: "u", timeout: 0)
```

If the parameter name in the function call does not match the parameter in the function
definition, the compiler will produce error [-8420](../15_library-reference/4483-genero-bdl-errors.md).

Named function call parameters does not imply free ordering, nor does it allow you to omit
parameters. All parameters must be specified, in the same order as in the function declaration.

In order to check the parameter names, fglcomp needs to know the function
definition. When the called function is defined in another module, use [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") to let the compiler know
the function definition.

> **Tip:**
>
> Named function parameters are provided by the [source code completer](../13_programming-tools/2544-configure-vim-for-genero-bdl.md).

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")
