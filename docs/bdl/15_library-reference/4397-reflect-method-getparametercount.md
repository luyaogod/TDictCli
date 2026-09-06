---
title: "reflect.Method.getParameterCount"
source: "fgl-topics/c_fgl_ext_reflect_Method_getParameterCount.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Method class > reflect.Method methods > reflect.Method.getParameterCount"
type: "concept"
---

# reflect.Method.getParameterCount

> Returns the number of parameters of a method.

## Syntax

```
getParameterCount()
  RETURNS INTEGER
```

## Usage

The `getParameterCount()` method returns the number of parameters of the method
object represented by this `reflect.Method` object.

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
    DISPLAY "param count   = ", met.getParameterCount()
END FUNCTION
```

Shows:

```
param count  =           2
```
