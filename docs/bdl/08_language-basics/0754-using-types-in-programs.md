---
title: "Using types in programs"
source: "fgl-topics/c_fgl_user_types_004.html"
breadcrumb: "Language basics > Types > Using types in programs"
type: "concept"
---

# Using types in programs

> Define a type as a synonym for an existing data type, or as a shortcut for records and array structures.

## Defining and using a type

After declaring a type, it can be used as a normal data type to define [variables](0686-variables.md "Explains how to define program variables.").

```
TYPE t_customer RECORD
         cust_num INTEGER,
         cust_name VARCHAR(50),
         cust_addr VARCHAR(200)
    END RECORD

DEFINE c1 t_customer

DEFINE custlist DYNAMIC ARRAY OF t_customer
```

The scope of a type is the same as for variables and constants. Types can be global,
module-specific, or local to a function.

## Using a type defined in another module

A good practice is to define types that belong to the same domain in a single
.4gl module, and import that module in the modules where the types are
needed.

By default, module-specific types are private; They cannot be used by an other module of the
program. To make a module type public, add the `PUBLIC` keyword before
`TYPE`. When a module type is declared as public, it can be referenced by another
module by using the [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
instruction:

Source file customers.4gl:

```
PUBLIC TYPE t_ord RECORD
           ord_id INTEGER,
           ord_date DATE,
           ord_total DECIMAL(10,2)
       END RECORD
PUBLIC TYPE t_cust RECORD
           cust_id INTEGER,
           cust_name VARCHAR(50),
           orders DYNAMIC ARRAY OF t_ord
       END RECORD
```

Source file main.4gl:

```
IMPORT FGL customers
MAIN
    DEFINE custlist DYNAMIC ARRAY OF t_cust
    DISPLAY custlist.getLength()
END MAIN
```

## Types for function references

Types can also be used to declare a function signature, in order to define program variables that
[reference functions](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") with that
signature:

```
TYPE callback_function FUNCTION(p1 INT, p2 INT) RETURNS INT
DEFINE v callback_function
    ...
    LET v = FUNCTION add
    ...
```

## Completing types with methods and interfaces

Types define data structures that can be manipulated from [methods](0772-methods.md "A function declared with a receiver type defines a method for this type."), to use the concept of encapsulation and make
your code more robust:

```
IMPORT util
PUBLIC TYPE Circle RECORD
    pos_x FLOAT,
    pos_y FLOAT,
    diameter FLOAT
END RECORD

PUBLIC FUNCTION (c Circle) area() RETURNS FLOAT
    RETURN util.Math.pi() * (c.diameter / 2) ** 2
END FUNCTION
```

Furthermore, to manipulate similar but different data structures with the same set of methods,
you can define interfaces to indirectly manipulate different [instances](0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") of types:

```
PUBLIC TYPE Circle RECORD ... END RECORD
PUBLIC FUNCTION (c Circle) area() RETURNS FLOAT ... END FUNCTION

PUBLIC TYPE Rectangle RECORD ... END RECORD
PUBLIC FUNCTION (r Rectangle) area() RETURNS FLOAT ... END FUNCTION

PUBLIC TYPE Triangle RECORD ... END RECORD
PUBLIC FUNCTION (r Triangle) area() RETURNS FLOAT ... END FUNCTION

PUBLIC TYPE Shape INTERFACE
    area() RETURNS FLOAT
END INTERFACE
```

## Related links

**Related concepts**  

[Records](0715-records.md "Records allow structured program variables definitions.")

[Arrays](0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
