---
title: "SMALLINT"
source: "fgl-topics/c_fgl_datatypes_SMALLINT.html"
breadcrumb: "Language basics > Primitive Data types > SMALLINT"
type: "concept"
---

# SMALLINT

> The SMALLINT data type is used for storing small whole numbers.

## Syntax

```
SMALLINT
```

## Usage

The storage of `SMALLINT` variables
is based on 2 bytes of signed data ( = 16 bits ).

The value
range is from -32,767 to +32,767.

`SMALLINT` variables can be initialized with [integer
literals](0586-integer-literals.md "Integer literals define a whole number in an expression."):

```
MAIN
  DEFINE i SMALLINT
  LET i = 1234
  DISPLAY i 
END MAIN
```

When assigning a whole number
that exceeds the `SMALLINT` range, the overflow error [-1214](../15_library-reference/4483-genero-bdl-errors.md) will
be raised.

`SMALLINT` variables are initialized
to zero in functions, modules and globals.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[INTEGER](0562-integer.md "The INTEGER data type is used for storing large whole numbers.")

[BIGINT](0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.")

[TINYINT](0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.")
