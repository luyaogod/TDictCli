---
title: "Function references"
source: "fgl-topics/c_fgl_Functions_references.html"
breadcrumb: "Language basics > Functions > Function references"
type: "concept"
---

# Function references

> Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.

## Syntax

A function type is defined with the following
syntax:

```
FUNCTION (
     parameter-name type-specification [ attributes-list ]
     [,...]
  )
 [ attributes-list ]
 [ RETURNS { type-specification [ attributes-list ]
           | ( type-specification [ attributes-list ] [,...] )
           | ( )
           } ]
```

1. parameter-name is the name of a formal argument of the function type.
2. type-specification can be one of:
   - A [primitive type](0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
3. attributes-list is a comma-separated list of name = value
   pairs or name attributes, and defines [attributes for the function](0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.").

## Purpose of function references

A function reference points to a function definition, that can be called at runtime. The actual
function is not known at compile time, only the function type (number and type of parameters and
return values) is known.

This feature allows you to manipulate functions dynamically, for example to implement generic
module, that can be configured with callback functions.

Function references are based on function types. Referenced functions must be defined with the
syntax defining parameter types in parentheses (and the `RETURNS` clause, if the
functions return values). For more details, see [`FUNCTION` syntax 2](0763-function-definitions.md "A FUNCTION definition defines a named procedure with a set of statements.").

## Defining function types

A function type identifies the signature of a function from the number, names and types of
parameters and return values of that function:

```
FUNCTION(p1 INT, p2 INT) RETURNS INT
```

The name of the parameters is part of the function signature. Function types using the same
number of parameters/types and return types, but different parameter names are considered as a
different function types by the compiler.

A function type can be used as other types, to declare simple variables, members of a structured
record, or arrays:

```
DEFINE fx FUNCTION(p1 INT, p2 INT) RETURNS INT
```

To simplify function reference usage, define a user-type with the `TYPE`
instruction, with the function type that will match functions to be called by reference:

```
TYPE callback_function FUNCTION(p1 INT, p2 INT) RETURNS INT
```

For more details about user-defined type definitions, see [Types](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.").

## Variable definition for function references

When the user-type for the function reference is available, declare a program variable to hold
such function reference:

```
DEFINE callback callback_function
```

For more details about variable definitions, see [Variables](0686-variables.md "Explains how to define program variables.").

## Get the FUNCTION reference

To get the reference of a function, use the `FUNCTION` keyword followed by the
name of the function to be referenced. The function must be defined in the current module, or in a
module imported with [`IMPORT
FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.").

```
LET callback = FUNCTION add
```

In the above example, the function "`add()`" must be defined with the same
function type as the "`callback`" variable.

For more details, see [FUNCTION func-spec](0771-function-func-spec.md "The FUNCTION keyword provides the reference to the specified function.").

## Invoking a function with the CALL statement

Functions referenced in a variable can be invoked with the `CALL` instruction, by using the variable.
The referenced function will be called as in a regular function
call:

```
CALL callback(100,200) RETURNING result
```

## Using function references in expressions

Variables referencing functions can be used in expressions, like in this example:

```
LET get_count_func = FUNCTION get_total_items()
LET c = get_count_func()
LET get_count_func = FUNCTION get_total_elements()
LET c = c + get_count_func()
```

## Passing function references as function parameters

Like other values, function references can be passed as function
parameters:

```
CALL process( FUNCTION add, FUNCTION sub, callback )
...
FUNCTION process( f1 callback_function,
                  f2 callback_function,
                  f3 callback_function )
   DISPLAY f1(100,200) + f2(200,50) + f3(150,300)
END FUNCTION
```

## Related links

**Related concepts**  

[Flow control](0674-flow-control.md "Definition of language elements and instructions that control the flow of a program.")
