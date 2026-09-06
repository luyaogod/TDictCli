---
title: "reflect.Value.set"
source: "fgl-topics/c_fgl_ext_reflect_Value_set.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.set"
type: "concept"
---

# reflect.Value.set

> Assigns the specified value to this value object.

## Syntax

```
set( x reflect.Value )
```

1. x is a `reflect.Value` object to be used to set the
   value.

## Usage

The `set()` method sets the value of the current `reflect.Value`
object, from another `reflect.Value` object passed as parameter.

The method assigns the current object with the original value passed as parameter.

In order to set values from expressions, use the [`reflect.Value.copyOf()`](4356-reflect-value-copyof.md "Creates a new reflect.Value object from a copy of an expression.") class method, to create a
`reflect.Value` object from an original [expression](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language.") or [variable](../08_language-basics/0686-variables.md "Explains how to define program variables."):

```
CALL val.set( reflect.Value.copyOf( 999.99 ) )
```

The parameter of the `set()` method can also be a `reflect.Value`
object created from [reflect.Value.valueOf()](4357-reflect-value-valueof.md "Creates a new reflect.Value object as a reference to the original variable."), referencing another variable.

## Example

```
IMPORT reflect
MAIN
    DEFINE src, dst reflect.Value
    DEFINE rec1, rec2 RECORD
        f1, f2 DECIMAL(10, 2)
    END RECORD
    LET rec1.f1 = -123.45
    LET rec1.f2 = 999.99
    DISPLAY "rec1.f1 = ", rec1.f1
    DISPLAY "rec1.f2 = ", rec1.f2
    LET src = reflect.Value.valueOf(rec1)
    LET dst = reflect.Value.valueOf(rec2)
    CALL dst.set( src )
    DISPLAY "rec2.f1 = ", rec2.f1
    DISPLAY "rec2.f2 = ", rec2.f2
END MAIN
```

Shows:

```
rec1.f1 =      -123.45
rec1.f2 =       999.99
rec2.f1 =      -123.45
rec2.f2 =       999.99
```
