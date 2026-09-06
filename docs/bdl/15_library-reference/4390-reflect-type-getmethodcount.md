---
title: "reflect.Type.getMethodCount"
source: "fgl-topics/c_fgl_ext_reflect_Type_getMethodCount.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getMethodCount"
type: "concept"
---

# reflect.Type.getMethodCount

> Returns the number of methods in a record with methods or in an interface.

## Syntax

```
getMethodCount()
  RETURNS INTEGER
```

## Usage

The `getMethodCount()` method returns the number of methods defined in an record
with methods, or in an interface, that is represented by this `reflect.Type`
object.

The `reflect.Type` object used to call this method must have been
created with a variable of [record-type with
methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), or an [`INTERFACE`](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") variable,
or is a `reflect.Type` object returned from a method like [`reflect.Value.getType()`](4371-reflect-value-gettype.md "Returns the reflect.Type object of a reflect.Value object.") or [`getElementType()`](4383-reflect-type-getelementtype.md "Returns the type of the elements in an array or in a dictionary."), and
references a record-type with methods or an interface type.

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
    DEFINE shp Shape
    DEFINE cir Circle
    LET typ = reflect.Type.typeOf(shp)
    DISPLAY "shp method count = ", typ.getMethodCount()
    LET typ = reflect.Type.typeOf(cir)
    DISPLAY "cir method count = ", typ.getMethodCount()
END FUNCTION
```

Shows:

```
shp method count =           2
cir method count =           2
```
