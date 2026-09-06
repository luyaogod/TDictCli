---
title: "INTERFACE usage"
source: "fgl-topics/c_fgl_Interface_usage.html"
breadcrumb: "Language basics > Interfaces > INTERFACE usage"
type: "concept"
description: "Defining interfaces An interface is defined by a group of methods that apply on user-defined types, to define a common usage interface for several individual types. A method declared in an interface ..."
---

# INTERFACE usage

## Defining interfaces

An interface is defined by a group of methods that apply on user-defined types, to define a
common usage interface for several individual types.

A method declared in an interface must use the same parameter names, parameter types and return
types as the method implementation it refers to.

An interface must be defined as a type with the [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
declaration:

```
TYPE Shape INTERFACE
    kind () RETURNS STRING,
    area () RETURNS FLOAT
END INTERFACE
```

A variable defined with an interface type can receive any type related to the interface.

## Associating types to an interface

A user-defined type for which methods are defined is implicitly associated to any interface that
defines a set of methods for this type.

For example, a type `Rectangle` gets associated to a method named
`area()`:

```
TYPE Rectangle RECORD
    height FLOAT,
    width FLOAT
END RECORD

FUNCTION (r Rectangle) area () RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION
```

An interface defined with the `area()` method becomes implicitly an interface for
the type
`Rectangle`:

```
TYPE Shape INTERFACE
    area() RETURNS FLOAT
END INTERFACE
```

If you define a variable `v` with an `INTERFACE` structure listing
the `area()` method, you can assign a variable defined as `Rectangle`
to `v`, and invoke the method with
`v.area()`:

```
TYPE Rectangle RECORD
    height FLOAT,
    width FLOAT
END RECORD
FUNCTION (r Rectangle) area () RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION

TYPE Shape INTERFACE
    area() RETURNS FLOAT
END INTERFACE

FUNCTION main()
    DEFINE r Rectangle = ( height:10, width:20 )
    DEFINE v Shape
    LET v = r
    DISPLAY v.area()
END FUNCTION
```

## Implementation tips

Write generic code in functions taking as parameter interfaces or collection or interfaces ([dynamic arrays](0734-dynamic-arrays.md), [dictionaries](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.")).

For maximum flexibility, consider implementing the types and corresponding methods in individual
modules, and implement the interface in another individual module. In the parent module using the
types/methods and interfaces, import each module with the [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") instruction.

## Interface referencing an unexisting record

An `INTERFACE` variable references a `RECORD` variable using a
`TYPE`. If the life time of the referenced `RECORD` is over (for
example, when using the `RECORD` variable defined in a function returned as an
`INTERFACE` from that function), the runtime system will raise the error -8083, when
calling a method of this `INTERFACE` variable, because the variable is no more
existing.

For example:

```
TYPE my_interface INTERFACE
         method_1() RETURNS ()
     END INTERFACE

TYPE my_record_type RECORD
         num INTEGER
     END RECORD

FUNCTION (r my_record_type) method_1() RETURNS ()
    DISPLAY r.*
END FUNCTION

FUNCTION foo() RETURNS my_interface
    DEFINE rec my_record_type   -- implements my_interface
    RETURN rec -- this returns an INTERFACE to a local variable
END FUNCTION

FUNCTION main()
    -- when calling method_1, the INTERFACE tries to access the rec RECORD
    -- which is no longer available.
    CALL foo().method_1()  -- expected runtime error -8083
END FUNCTION
```

The above code compiles but produces the following runtime
error:

```
Program stopped at 'm.4gl', line number 21.
FORMS statement error number -8083.
Null pointer exception.
```

## Related links

**Related concepts**  

[Methods](0772-methods.md "A function declared with a receiver type defines a method for this type.")
