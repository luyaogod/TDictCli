---
title: "reflect.Method.getSignature"
source: "fgl-topics/c_fgl_ext_reflect_Method_getSignature.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Method class > reflect.Method methods > reflect.Method.getSignature"
type: "concept"
---

# reflect.Method.getSignature

> Returns the signature of a method, as a string.

## Syntax

```
getSignature()
  RETURNS STRING
```

## Usage

The `getSignature()` method returns a string the defines the signature of the
method represented by this `reflect.Method` object.

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
    DEFINE cus Customer
    LET typ = reflect.Type.typeOf(cus)
    DISPLAY "signature 1 = ", typ.getMethod(1).getSignature()
    DISPLAY "signature 2 = ", typ.getMethod(2).getSignature()
END FUNCTION
```

Shows:

```
signature 1 = (id:INTEGER, name:VARCHAR(30)) RETURNS SMALLINT
signature 2 = (id:INTEGER) RETURNS SMALLINT
```
