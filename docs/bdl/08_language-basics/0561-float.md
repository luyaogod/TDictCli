---
title: "FLOAT"
source: "fgl-topics/c_fgl_datatypes_FLOAT.html"
breadcrumb: "Language basics > Primitive Data types > FLOAT"
type: "concept"
---

# FLOAT

> The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.

## Syntax

```
FLOAT [(precision)]
```

1. `FLOAT` and `DOUBLE PRECISION` are synonyms.
2. The precision can be specified, but it has no effect in programs.

## Usage

The storage of `FLOAT` variables is based on 8 bytes of signed data (64 bits),
this type is equivalent to the `double` data type in C.

The `FLOAT` data type is not recommended for exact decimal storage; use the
`DECIMAL` type instead.

`FLOAT` variables are initialized to zero in functions, modules and
globals.

`FLOAT` values can be converted to strings based on the [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") (or DBMONEY) environment variable
setting.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[DECIMAL(p,s)](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.")
