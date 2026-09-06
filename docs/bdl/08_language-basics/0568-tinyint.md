---
title: "TINYINT"
source: "fgl-topics/c_fgl_datatypes_TINYINT.html"
breadcrumb: "Language basics > Primitive Data types > TINYINT"
type: "concept"
---

# TINYINT

> The TINYINT data type is used for storing very small whole numbers.

## Syntax

```
TINYINT
```

## Usage

The storage of `TINYINT` variables
is based on 1 byte of signed data ( = 8 bits ).

The value range
is from -128 to +127.

`TINYINT` variables can be initialized with [integer
literals](0586-integer-literals.md "Integer literals define a whole number in an expression."):

```
MAIN
  DEFINE i TINYINT
  LET i = 101
  DISPLAY i 
END MAIN
```

When assigning a whole number that exceeds
the `TINYINT` range, the overflow error [-8097](../15_library-reference/4483-genero-bdl-errors.md) will
be raised.

`TINYINT` variables are initialized
to zero in functions, modules and globals.

The `TINYINT` variables
cannot be `NULL`.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[SMALLINT](0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.")

[INTEGER](0562-integer.md "The INTEGER data type is used for storing large whole numbers.")

[BIGINT](0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.")
