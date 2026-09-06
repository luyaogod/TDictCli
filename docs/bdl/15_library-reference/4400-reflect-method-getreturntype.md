---
title: "reflect.Method.getReturnType"
source: "fgl-topics/c_fgl_ext_reflect_Method_getReturnType.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Method class > reflect.Method methods > reflect.Method.getReturnType"
type: "concept"
---

# reflect.Method.getReturnType

> Returns the type of a return value of a method.

## Syntax

```
getReturnType(
     index INTEGER )
  RETURNS reflect.Type
```

1. index is the ordinal position of the return value.

## Usage

The `getReturnType()` method returns a [`reflect.Type`](4379-the-reflect-type-class.md "The reflect.Type class is a generic API to inspect types.") object representing the type of the return value at the
specified index, for the method represented by this `reflect.Method` object.

The `reflect.Method` object used to call this method must have
been created with the [`reflect.Type.getMethod()`](4389-reflect-type-getmethod.md "Returns a reflect.Method object of record with methods or an interface.") method, from a `reflect.Type`
object created with a [`RECORD` with
methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), or and [`INTERFACE`](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") type.

## Example

```
IMPORT reflect
TYPE Customer INTERFACE
        create(id INTEGER, name VARCHAR(30)) RETURNS SMALLINT,
        delete(id INTEGER) RETURNS SMALLINT
    END INTERFACE
FUNCTION main()
    DEFINE typ reflect.Type
    DEFINE met reflect.Method
    DEFINE cus Customer
    LET typ = reflect.Type.typeOf(cus)
    LET met = typ.getMethod(1)
    DISPLAY "return type 1 = ", met.getReturnType(1).toString()
END FUNCTION
```

Shows:

```
return type 1 = SMALLINT
```
