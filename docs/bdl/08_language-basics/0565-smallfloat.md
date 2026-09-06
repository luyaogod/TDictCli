---
title: "SMALLFLOAT"
source: "fgl-topics/c_fgl_datatypes_SMALLFLOAT.html"
breadcrumb: "Language basics > Primitive Data types > SMALLFLOAT"
type: "concept"
---

# SMALLFLOAT

> The SMALLFLOAT data type stores values as single-precision floating-point binary numbers with up to 8 significant digits.

## Syntax

```
SMALLFLOAT
```

1. `SMALLFLOAT` and `REAL` are synonyms.

## Usage

The storage of `SMALLFLOAT` variables is based on 4 bytes of signed
data ( =32 bits), this type is equivalent to the `float` data type in C.

The `SMALLFLOAT` data type is not recommended for exact decimal storage; use the
`DECIMAL` data type instead.

`SMALLFLOAT` variables are initialized to zero in functions, modules
and globals.

`SMALLFLOAT` values can be converted to strings based on the [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") (or DBMONEY) environment variable
setting.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[DECIMAL(p,s)](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.")
