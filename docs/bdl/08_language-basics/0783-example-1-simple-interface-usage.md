---
title: "Example 1: Simple INTERFACE usage"
source: "fgl-topics/c_fgl_Interface_example_1.html"
breadcrumb: "Language basics > Interfaces > Examples > Example 1: Simple INTERFACE usage"
type: "concept"
---

# Example 1: Simple INTERFACE usage

> Defines an INTERFACE to handle shape objects.

This first module called rectangle.4gl defines the
`Rectangle` type and the methods that apply to this
type:

```
PUBLIC TYPE Rectangle RECORD
    height, width FLOAT
END RECORD

PUBLIC FUNCTION (r Rectangle) area() RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION
    
PUBLIC FUNCTION (r Rectangle) kind() RETURNS STRING
    RETURN "Rectangle"
END FUNCTION
    
PUBLIC FUNCTION (r Rectangle) setDimensions(w FLOAT, h FLOAT) RETURNS()
    LET r.width = w
    LET r.height = h
END FUNCTION
```

The circle.4gl module defines the `Circle` type and methods
that apply to this
type:

```
IMPORT util

PUBLIC TYPE Circle RECORD
    diameter FLOAT
END RECORD

PUBLIC FUNCTION (c Circle) area() RETURNS FLOAT
    RETURN util.Math.pi() * (c.diameter / 2) ** 2
END FUNCTION

PUBLIC FUNCTION (c Circle) kind() RETURNS STRING
    RETURN "Circle"
END FUNCTION

PUBLIC FUNCTION (c Circle) setDiameter(d FLOAT) RETURNS()
    LET c.diameter = d
END FUNCTION
```

The "shapes.4gl" module implements the `Shape` interface, to
group the `area()` and `kind()` methods, which both apply to the
`Rectangle` and `Circle` types. The `totalArea()`
functions takes a dynamic array of `Shape` objects, knowing only about the behavior
defined by the `Shape` interface.

In this example, the module implementing the `Shape` interface does not need to
import the `rectangle` and `circle` modules: The interface definition
is abstract and is independent from the types definitions: Only the method definitions matters.
Other shape types could be implemented, such as `Triangle`,
`Trapeziod`, without touching the `Shapes` interface.

Note also that the `Rectangle.setDimensions()` and
`Circle.setDiameter()` methods are not part of the `Shape` interface:
These methods are specific to each shape types and therefore cannot be part of a common
interface.

```
PUBLIC TYPE Shape INTERFACE
    area() RETURNS FLOAT,
    kind() RETURNS STRING
END INTERFACE 
    
PUBLIC TYPE ShapesArray DYNAMIC ARRAY OF Shape

PUBLIC FUNCTION totalArea(shapes ShapesArray) RETURNS FLOAT
    DEFINE i INT
    DEFINE area FLOAT
    FOR i = 1 TO shapes.getLength()
        LET area = area + shapes[i].area()
    END FOR 
    RETURN area
END FUNCTION
```

The main module imports the `rectangle`, `circle` and
`shapes` modules, defines `rectangle` and `circle`
variables, that can be assigned to an dynamic array of `shapes`, that can in turn be
used with the abstract interface:

```
IMPORT FGL rectangle
IMPORT FGL circle
IMPORT FGL shapes

FUNCTION main()
    DEFINE r1 rectangle.Rectangle
    DEFINE c1 circle.Circle
    DEFINE sa shapes.ShapesArray

    CALL r1.setDimensions(10, 20)
    CALL c1.setDiameter(20)

    LET sa[1] = r1
    LET sa[2] = c1

    DISPLAY sa[1].kind(), sa[1].area()
    DISPLAY sa[2].kind(), sa[2].area()
    DISPLAY "Total area:", shapes.totalArea(sa)

END FUNCTION
```
