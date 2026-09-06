---
title: "INTERFACE"
source: "fgl-topics/c_fgl_Interface_syntax.html"
breadcrumb: "Language basics > Interfaces > INTERFACE"
type: "concept"
---

# INTERFACE

> An interface is defined by a list of methods for a type.

## Syntax

```
INTERFACE
    method-name ( 
        parameter-name type-specification
        [,...]
      )
     [ RETURNS { type-specification
               | ( [ type-specification [,...] ] )
               } ]
   [,...]
END INTERFACE
```

1. method-name defines the name of a method.
2. parameter-name is the name of a formal argument of the method.
3. type-specification can be one of:
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

## Usage

An `INTERFACE` structure defines a list of methods that apply to types defined as
records with [methods](0772-methods.md "A function declared with a receiver type defines a method for this type."). The interface defines the
how, the type defines the what.

An interface is associated to a record type through the list of methods defined for that type.
For example, a `Rectangle` record type defining an `area()` method can
be associated to an interface defining that method:

```
TYPE Rectangle RECORD
    height FLOAT,
    width FLOAT
END RECORD
PUBLIC FUNCTION (r Rectangle) area() RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION
```

All elements inside an `INTERFACE` must be and can only be [methods for a user-defined type](0772-methods.md "A function declared with a receiver type defines a method for this type."), and must be specified in
the interface by using the same parameter names, parameter types and return types of the methods it
refers to. An interface is typically defined as a `TYPE`, to simplify its
reusage:

```
TYPE Shape INTERFACE
    area() RETURNS FLOAT,
    kind() RETURNS STRING
END INTERFACE
```

Methods of multiple individual types associated to an interface can be invoked indirectly by
declaring a variable with the interface structure:

```
DEFINE s Shape, x FLOAT
LET x = s.area()  -- Can be the area() method for types Circle, Rectangle, etc.
```

Before it can be used, an interface variable must be set by referencing an instance of a another
variable defined with a corresponding record
type:

```
DEFINE r Rectangle, s Shape, x FLOAT
LET s = r
LET x = s.area()
```

Several interfaces can be defined for a given type. This provides a high level of
flexibility:

```
TYPE Rectangle RECORD
    height FLOAT,
    width FLOAT
END RECORD
FUNCTION (r Rectangle) area() RETURNS FLOAT
    RETURN (r.height * r.width)
END FUNCTION
FUNCTION (r Rectangle) domainName() RETURNS STRING
    RETURN "Geometry"
END FUNCTION

TYPE Shape INTERFACE
    area() RETURNS FLOAT
END INTERFACE

TYPE Domain INTERFACE
    domainName() RETURNS STRING
END INTERFACE

FUNCTION main()
    DEFINE r Rectangle = ( height:10, width:20 )
    DEFINE v1 Shape
    DEFINE v2 Domain
    LET v1 = r
    DISPLAY v1.area()
    LET v2 = r
    DISPLAY v2.domainName()
END FUNCTION
```

## Related links

**Related concepts**  

[Variables](0686-variables.md "Explains how to define program variables.")
