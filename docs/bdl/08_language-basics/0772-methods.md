---
title: "Methods"
source: "fgl-topics/c_fgl_Functions_methods.html"
breadcrumb: "Language basics > Functions > Methods"
type: "concept"
---

# Methods

> A function declared with a receiver type defines a method for this type.

## Purpose of methods

Methods are functions that perform on a variable with a specific [user-defined type](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."), and cannot be used for another type.

Use methods to implement the [interface](0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") (the access
methods) for a type.

The concept of methods allows you to write robust code like in Object-Oriented Programming
languages, without the complexity and traps of OOP.

## Defining a method

A method is a `FUNCTION` defined with a receiver argument specified in parentheses
before the function name.

The receiver consists of an identifier followed by a user-defined type. The receiver identifier
is then referenced in the function body as the target of this
method:

```
TYPE Rectangle RECORD
    height, width FLOAT
END RECORD

FUNCTION (r Rectangle) area() RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION
```

## Method definition rules

The receiver type must be a `TYPE` defined in the same module as the method, and
it must define a structured `RECORD` type (it cannot be a flat type defined with a
primitive type such as `INTEGER`).

> **Important:**
>
> Receiver types and methods for this type must be defined in the same module.
> If the receiver type is not defined in the same module as the method, the compiler produces the
> error -8426.

Method names and receiver type field names must be different: It is not legal to define a method
with the same name as a field of the receiver
type:

```
PUBLIC TYPE Rectangle RECORD
    height, width FLOAT
END RECORD
...
PUBLIC FUNCTION (r Rectangle) width () RETURNS INTEGER
| Method and field names must be different. Type rectangle.Rectangle has field width.
| See error number -8427.
    RETURN r.width
END FUNCTION
```

The compiler is more strict regarding methods definitions, compared to regular functions. For
example, when a method returns values, it must be defined with the `RETURNS`
clause:

```
PUBLIC TYPE Rectangle RECORD
    height, width FLOAT
END RECORD
...
PUBLIC FUNCTION (r Rectangle) getWidth ()
    RETURN r.width
| A method without return type (RETURNS) can not return values.
| See error number -8431.
END FUNCTION
```

## Invoking a method

A method for type needs a variable defined to reference an instance of that type. To invoke a
method, prefix the method name with the variable of the type used as receiver for the
method:

```
DEFINE r1, r2 Rectangle
DISPLAY r1.area()
DISPLAY r2.area()
```

## Polymorphism

The same function name can be reused for a different receiver type. Using a different receiver
type for the same function name implies another
method:

```
FUNCTION (r Rectangle) area() RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION

FUNCTION (r Triangle) area() RETURNS FLOAT
    RETURN ( r.height * r.base ) / 2
END FUNCTION
```

## Receiver variable is passed by reference

A method for a type can implicitly modify the receiver (the receiver variable is passed by
reference to the
method):

```
FUNCTION (r Rectangle) setDimensions(w FLOAT, h FLOAT)
    CALL myAssert( (w>0 AND h>0), "Invalid dimensions!" ) -- stops program
    LET r.width = w
    LET r.height = h
END FUNCTION
```

## Methods with record parameters

When defining a method taking a structured record type as parameter, you must call the method
with a record variable name. The record will be passed as a copy: The original record cannot be
modified, unless the `INOUT` clause is used.

The [`.*` dot star notation](../09_advanced-features/0835-passing-records-as-parameter.md) (expending the record fields on the stack) is not
allowed when calling a method with a record as parameter. This technique is supported with regular
functions for backward compatibility.

```
TYPE t_rec1 RECORD
         f1 INT
     END RECORD

TYPE t_rec2 RECORD
         f1 INT
     END RECORD

FUNCTION (r t_rec1) method1(p t_rec2)
    LET p.f1 = 999
END FUNCTION

FUNCTION (r t_rec1) method2(p t_rec2 INOUT)
    LET p.f1 = 999
END FUNCTION

FUNCTION main()
    DEFINE r1 t_rec1, r2 t_rec2
    CALL r1.method1( r2 )
    DISPLAY r2.f1 -- not modified
    CALL r1.method2( r2 )
    DISPLAY r2.f1 -- 999
    -- CALL r1.method1( r2.* ) -- invalid!
END FUNCTION
```

## Methods returning a record structure

Methods can return a complete `RECORD`: The method must be defined with
`RETURNS type-name` and use the `RETURN` clause by
specifying the record variable without the `.*` notation (this would expand all
members of the record on the
stack):

```
TYPE t_cust RECORD
    id INT,
    name VARCHAR(50)
END RECORD

TYPE t_status RECORD
    errcode INT,
    message STRING
END RECORD

FUNCTION (r t_cust) checkData() RETURNS t_status
    DEFINE s t_status
    CASE
        WHEN r.id IS NULL OR r.id < 1
            LET s.errcode = -9
            LET s.message = "Invalid identifier"
        WHEN LENGTH(r.name) = 0
            LET s.errcode = -8
            LET s.message = "Name is empty"
    END CASE
    RETURN s
END FUNCTION

FUNCTION main()
    DEFINE r t_cust
    DEFINE s t_status

    LET s = r.checkData()
    DISPLAY s.*

    LET r.id = 101
    LET s = r.checkData()
    DISPLAY s.*

    LET r.name = "John Lambert"
    LET s = r.checkData()
    DISPLAY s.*

END FUNCTION
```

## Related links

**Related concepts**  

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")

[Commenting a method](../13_programming-tools/2557-commenting-a-method.md "Commenting a method")
