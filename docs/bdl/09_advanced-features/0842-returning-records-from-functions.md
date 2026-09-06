---
title: "Returning records from functions"
source: "fgl-topics/c_fgl_runtime_stack_return_record.html"
breadcrumb: "Advanced features > Runtime stack > Returning records from functions"
type: "concept"
description: "Returning records as single structured value Records can be returned from functions and methods as a single structured value, by using only the record name in the RETURNS instruction. Important: To ..."
---

# Returning records from functions

## Returning records as single structured value

Records can be returned from functions and methods as a single structured value, by using only
the record name in the `RETURNS` instruction.

> **Important:**
>
> To return a record as a single structured value, functions must have been
> defined with the [fully typed
> syntax](../08_language-basics/0763-function-definitions.md). Methods for types can only be defined with a [fully typed
> syntax](../08_language-basics/0763-function-definitions.md).

For example:

```
TYPE t_cust RECORD
    pkey INTEGER,
    name VARCHAR(50)
END RECORD

FUNCTION newCustomer() RETURNS t_cust
    DEFINE c t_cust
    LET c.pkey = 0
    LET c.name = "<undefined>"
    RETURN c
END FUNCTION

FUNCTION main()
    DEFINE r t_cust
    LET r = newCustomer()
    DISPLAY r.*
END FUNCTION
```

This syntax improves code readability and robustness: With the legacy `.*`
notation, the compiler is more permissive and allows to pass and return record members with
different types (this is usually done by mistake). When assigning a record without the
`.*` notation from a function return, the compiler produces error [-4325](../15_library-reference/4483-genero-bdl-errors.md), if the types of the target
record and the type used in the function `RETURNS` clause do not match.

> **Tip:**
>
> Another way to return records is is to [pass records as reference](0835-passing-records-as-parameter.md) with the `INOUT` keyword.

## Returning records by expansion (`.*` notation)

Records can be returned from functions in the `RETURN` instruction, by expanding
the members on the stack, using `.*` (dot-star) after the record
name:

```
TYPE t_cust RECORD
    pkey INTEGER,
    name VARCHAR(50)
END RECORD

FUNCTION newCustomer()
    DEFINE c t_cust
    LET c.pkey = 0
    LET c.name = "<undefined>"
    RETURN c.*
END FUNCTION

FUNCTION main()
    DEFINE r t_cust
    CALL newCustomer() RETURNING r.*
    DISPLAY r.*
END FUNCTION
```

Returning records by expansion is supported for backward compatiblity. Consider returning records
as a single structured value, without the `.*` notation, and fully typed function
definitions.

## Related links

**Related concepts**  

[RETURN](../08_language-basics/0676-return.md "The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.")

[Methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type.")
