---
title: "BIGINT"
source: "fgl-topics/c_fgl_datatypes_BIGINT.html"
breadcrumb: "Language basics > Primitive Data types > BIGINT"
type: "concept"
---

# BIGINT

> The BIGINT data type is used for storing very large whole numbers.

## Syntax

```
BIGINT
```

## Usage

The storage of `BIGINT` variables is based on 8 bytes of signed data (
= 64 bits ).

The value range is from -9,223,372,036,854,775,807 to +9,223,372,036,854,775,807.

`BIGINT` variables can be initialized with [integer
literals](0586-integer-literals.md "Integer literals define a whole number in an expression."):

```
MAIN
  DEFINE i BIGINT
  LET i = 9223372036854775600
  DISPLAY i
END MAIN
```

When assigning a whole number that exceeds the `BIGINT` range, the
overflow error [-1284](../15_library-reference/4483-genero-bdl-errors.md) will be raised.

`BIGINT` variables are initialized to zero in functions, modules and
globals.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[INTEGER](0562-integer.md "The INTEGER data type is used for storing large whole numbers.")

[SMALLINT](0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.")

[TINYINT](0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.")
