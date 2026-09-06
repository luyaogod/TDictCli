---
title: "Passing records as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_record.html"
breadcrumb: "Advanced features > Runtime stack > Passing records as parameter"
type: "concept"
description: "Passing records by reference (with INOUT ) To pass records by reference, define a TYPE with the RECORD structure, and use this type in the parameter definition of the function, followed by the INOUT ..."
---

# Passing records as parameter

## Passing records by reference (with `INOUT`)

To pass records by reference, define a `TYPE` with the `RECORD`
structure, and use this type in the parameter definition of the function, followed by the
`INOUT` keyword.

The function must be known by the fglcomp compiler: It is not possible to use
this feature for linked functions. The function defining the `INOUT` parameter must
be declared in the same module, or it must be defined in a module imported with [`IMPORT FGL`](0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.").

The function can then be invoked by specifying the record name as parameter.

The record structure defined in the caller can then be modified inside the function
body:

```
TYPE t_cust RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
     END RECORD

MAIN
    DEFINE r t_cust
    CALL initialize_customer(r)
    DISPLAY r.*
END MAIN

FUNCTION initialize_customer(r t_cust INOUT) RETURNS ()
    LET r.cust_id = 0
    LET r.cust_name = "<undefined>"
END FUNCTION
```

See also the optimization topic [Passing records by reference](0977-passing-records-by-reference.md).

## Passing records as single structured value

When only specifying the name of the record as parameter to a function, it is passed as a single
structured value.

Functions should be been defined with the [fully typed
syntax](../08_language-basics/0763-function-definitions.md).

The record can be used in the function body, but when modifying members of the record, the record
of the caller will not be modified.

```
TYPE t_cust RECORD
    cust_id INTEGER,
    cust_name VARCHAR(50)
END RECORD

MAIN
    DEFINE r t_cust
    CALL display_customer(r)
END MAIN

FUNCTION display_customer(r t_cust) RETURNS ()
    DISPLAY r.*
    LET r.cust_id = 999 -- has no effect
END FUNCTION
```

This syntax improves code readability and robustness: With the legacy `.*`
notation, the compiler is more permissive and allows to pass and return record members with
different types (this is usually done by mistake). When using simple record names without
`.*` as parameter, the compiler produces error [-4325](../15_library-reference/4483-genero-bdl-errors.md), if the types of the passed
record and the type of the function parameter do not match:

```
TYPE t_cust RECORD
    cust_id INTEGER,
    cust_name VARCHAR(50)
END RECORD

MAIN
    DEFINE r t_cust
    CALL display_customer(r)
| The source and destination records in this record assignment statement are not compatible...
| See error number -4325.
END MAIN

FUNCTION display_customer(r)
    DEFINE r RECORD
        cust_id SMALLINT, -- Different type as t_cust.cust_id!
        cust_name VARCHAR(50)
    END RECORD
END FUNCTION
```

If the function is not know by the compiler (linker is used), the `CALL`
instruction with a simple record name paramerter will produce the compiler error [-4340](../15_library-reference/4483-genero-bdl-errors.md). To make the compiler check
with a valid call, the function must be declared in the same module, or it must be defined in a
module imported with [`IMPORT
FGL`](0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.").

Records can also be passed as single structured value to [methods for
types](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."):

```
FUNCTION (c t_cust) setLastOrderId(o t_order) RETURNS ()
    LET c.last_order_id = o.order_id
END FUNCTION
```

Methods for types can only be defined with a [fully typed
syntax](../08_language-basics/0763-function-definitions.md).

## Passing records as single structured value to APIs

API methods such as [`util.JSON.stringify()`](../15_library-reference/3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.") use records passed as single structured
value:

```
IMPORT util
...
DEFINE rec RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
       END RECORD
...
    CALL util.JSON.stringify( rec )
...
```

## Passing records implicitly by reference APIs

With API methods such as [`util.JSON.parse()`](../15_library-reference/3582-util-json-parse.md "Parses a JSON string and fills program variables with the values."), the record is implicitly passed by reference, and it can
be modified by the
method:

```
IMPORT util
...
DEFINE rec RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
       END RECORD
...
    CALL util.JSON.parse( '{"cust_id":2735, "cust_name":"McCarlson"}', rec )
...
```

## Passing records by expansion (`.*` notation)

All elements of a `RECORD` structure can be passed by value to a function with the
dot star (`.*`) notation.

The `record.*` parameter syntax is supported for backward compatibility. Consider
passing records by reference (`INOUT`), or by value, by specifying the record name
only.

When using the `.*` notation as function parameter, the record is expanded: Each
member of the record structure is pushed on the stack. The receiving local variables in the function
must be defined with the same record structure as the caller.

For backward compatibility, the runtime system allows to call functions defined with the
individual parameters matching the record members (or the other way around - calling a function with
individual values that match the record parameter definition of the function). However, this is bad
practice.

> **Important:**
>
> When using [methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), the
> parameter type checking is enforced. It is not possible to pass records with the `.*`
> notation like with legacy functions: The compiler will reject the code, if the complex parameters
> types (records, arrays) do not match.

In the next example, the `func_r()` function is defined with a record parameter,
while the `func_ab()` function is defined with individual integer and string
parameters:

```
MAIN
    DEFINE rec RECORD
        a INT,
        b VARCHAR(50)
    END RECORD
    CALL func_r(rec.*)
    CALL func_ab(rec.*)     -- Bad practice!
    CALL func_r(101, 'aaa') -- Bad practice!
END MAIN

-- Function defining a record like that in the caller
FUNCTION func_r( r )
    DEFINE r RECORD
        a INT,
        b VARCHAR(50)
    END RECORD
    DISPLAY "func_r : ", r.*
END FUNCTION

-- Function defining two individual variables
FUNCTION func_ab(a, b)
    DEFINE
        a INT,
        b VARCHAR(50)
    DISPLAY "func_ab: ", a, b
END FUNCTION
```

## Related links

**Related concepts**  

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")
