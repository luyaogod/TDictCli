---
title: "reflect.Type.getMethod"
source: "fgl-topics/c_fgl_ext_reflect_Type_getMethod.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getMethod"
type: "concept"
---

# reflect.Type.getMethod

> Returns a reflect.Method object of record with methods or an interface.

## Syntax

```
getMethod(
     index INTEGER )
  RETURNS reflect.Method
```

1. index is the ordinal position of the method in the `INTERFACE`
   definition.

## Usage

The `getMethod()` method returns a `reflect.Method` object from a
`reflect.Type` object that represents a record with methods or an interface. The
returned object can then be used to describe the methods of the record type of interface type.

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
    DISPLAY "shp method 1 = ", typ.getMethod(1).getName()
    DISPLAY "shp method 2 = ", typ.getMethod(2).getName()
    LET typ = reflect.Type.typeOf(cir)
    DISPLAY "cir method 1 = ", typ.getMethod(1).getName()
    DISPLAY "cir method 2 = ", typ.getMethod(2).getName()
END FUNCTION
```

Shows:

```
shp method 1 = kind
shp method 2 = area
cir method 1 = Circle.kind
cir method 2 = Circle.area
```

## Related links

**Related concepts**  

[reflect.Type.getMethodCount](4390-reflect-type-getmethodcount.md "Returns the number of methods in a record with methods or in an interface.")
