---
title: "BOOLEAN"
source: "fgl-topics/c_fgl_datatypes_BOOLEAN.html"
breadcrumb: "Language basics > Primitive Data types > BOOLEAN"
type: "concept"
---

# BOOLEAN

> The BOOLEAN data type stores a logical value, TRUE or FALSE.

## Syntax

```
BOOLEAN
```

## Usage

`BOOLEAN` values can be `NULL`, `TRUE` (integer 1)
or `FALSE` (integer 0).

The default value of a `BOOLEAN` variable is `FALSE`.

Boolean variables are typically used to store the result of a [boolean
expression](0594-boolean-expressions.md "This section covers boolean expression evaluation rules."):

```
FUNCTION checkOrderStatus( cid INTEGER )
  DEFINE b BOOLEAN
  LET b = ( isValid(cid) AND isStored(cid) )
  IF NOT b THEN
    ERROR "The order is not ready."
  END IF
END FUNCTION
```

Data type conversion can be controlled by catching the runtime exceptions. For more details, see
[Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

SQL Database vendor specific implementation of the boolean SQL type may not correspond exactly to
the Genero `BOOLEAN` values. However, `BOOLEAN` variables can be used
in SQL statements: The conversion from the BDL `BOOLEAN` values
`TRUE/1` and `FALSE/0` to/from the native SQL boolean value will be
done by the database client or by the ODI driver. For more details, see [The BOOLEAN data type](../10_sql-support/1010-the-boolean-data-type.md "SQL implementation of the BOOLEAN data type varies on the database type.").
