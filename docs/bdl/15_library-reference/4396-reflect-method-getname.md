---
title: "reflect.Method.getName"
source: "fgl-topics/c_fgl_ext_reflect_Method_getName.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Method class > reflect.Method methods > reflect.Method.getName"
type: "concept"
---

# reflect.Method.getName

> Returns the name of the method.

## Syntax

```
getName()
  RETURNS STRING
```

## Usage

The `getName()` method returns the name of the method represented by this
`reflect.Method` object.

The `reflect.Method` object used to call this method must have
been created with the [`reflect.Type.getMethod()`](4389-reflect-type-getmethod.md "Returns a reflect.Method object of record with methods or an interface.") method, from a `reflect.Type`
object created with a [`RECORD` with
methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), or and [`INTERFACE`](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") type.

When the source type of the method is a record-type, the method name is prefixed with the name of
the module where the record-type is defined. When the source type is an interface-type, the method
name is not prefixed.

## Example

```
IMPORT reflect
  
TYPE Shape INTERFACE
        kind() RETURNS STRING,
        area() RETURNS FLOAT
    END INTERFACE

TYPE Circle RECORD
        x, y FLOAT,
        r FLOAT
    END RECORD
FUNCTION (c Circle) kind() RETURNS STRING
    RETURN "circle"
END FUNCTION
FUNCTION (c Circle) area() RETURNS FLOAT
    RETURN (c.r * c.r * 3.1416)
END FUNCTION

FUNCTION main()
    DEFINE typ reflect.Type
    DEFINE met reflect.Method
    DEFINE shp Shape
    DEFINE cir Circle
    LET typ = reflect.Type.typeOf(shp)
    LET met = typ.getMethod(1)
    DISPLAY "shp method name 1 = ", met.getName()
    LET typ = reflect.Type.typeOf(cir)
    LET met = typ.getMethod(1)
    DISPLAY "cir method name 1 = ", met.getName()
END FUNCTION
```

Shows:

```
shp method name 1 = kind
cir method name 1 = Circle.kind
```
