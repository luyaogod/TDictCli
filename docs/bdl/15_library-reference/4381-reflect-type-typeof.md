---
title: "reflect.Type.typeOf"
source: "fgl-topics/c_fgl_ext_reflect_Type_typeOf.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.typeOf"
type: "concept"
---

# reflect.Type.typeOf

> Creates a new reflect.Type object representing a type.

## Syntax

```
reflect.Type.typeOf(
     val any-type )
  RETURNS reflect.Type
```

1. val is an [expression](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language."). This is
   typically a [variable](../08_language-basics/0686-variables.md "Explains how to define program variables."), but it can be any expression
   supported by the language.
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `reflect.Type.typeOf()` class method returns a `reflect.Type`
object created from the expression passed a parameter.

The `reflect.Type` object can then be used to call object methods to describe the
type.

The expression passed as parameter is typically a program variable. However, the typeOf() method
can also directly be called with any expression.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
                   pkey INTEGER,
                   name VARCHAR(30)
               END RECORD
    DEFINE typ reflect.Type
    LET typ = reflect.Type.typeOf( arr )
    DISPLAY "type name 1 = ", typ.toString()
    DISPLAY "type name 2 = ", reflect.Type.typeOf( 4 / 2 ).toString()
END MAIN
```

Shows:

```
type name 1 = DYNAMIC ARRAY OF RECORD
type name 2 = DECIMAL
```
