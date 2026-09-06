---
title: "MONEY(p,s)"
source: "fgl-topics/c_fgl_datatypes_MONEY.html"
breadcrumb: "Language basics > Primitive Data types > MONEY(p,s)"
type: "concept"
---

# MONEY(p,s)

> The MONEY data type is provided to store currency amounts with exact decimal storage.

## Syntax

```
MONEY [ (precision[,scale]) ]
```

1. precision defines the number of significant digits (limit is 32, default is
   16).
2. scale defines the number of digits to the right of the decimal point.
3. When no scale is specified, it defaults to 2.
4. When no (precision, scale) is specified,
   it defaults to `MONEY(16,2)`.

## Usage

The `MONEY` data type is provided to store currency amounts. Its behavior is
similar to the `DECIMAL` data type, with some important differences:

A `MONEY` variable is displayed with the currency symbol defined in the
[DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") (or DBMONEY) environment variable.

When the the scale is not specified for the `MONEY` type,
the default is 2. A `MONEY` without precision and
scale defaults to `MONEY(16,2)`.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

See [DECIMAL(p,s)](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") to learn other facts about the
`MONEY(p,s)` data type.
