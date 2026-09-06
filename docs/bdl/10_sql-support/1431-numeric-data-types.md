---
title: "Numeric data types"
source: "fgl-topics/c_fgl_odiagpgs_027.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > Numeric data types"
type: "concept"
description: "Informix® Informix supports several data types to store numbers: Table 1. Informix numeric data types Informix data type Description SMALLINT 16 bit signed integer INTEGER 32 bit signed integer BIGINT ..."
---

# Numeric data types

## Informix®

Informix supports several data types to store
numbers:

| Informix data type | Description |
| --- | --- |
| `SMALLINT` | 16 bit signed integer |
| `INTEGER` | 32 bit signed integer |
| `BIGINT` | 64 bit signed integer |
| `INT8` | 64 bit signed integer (replaced by `BIGINT`) |
| `DECIMAL` | Equivalent to `DECIMAL(16)` |
| `DECIMAL(p)` | Floating-point decimal number (max precision is 32) |
| `DECIMAL(p,s)` | Fixed-point decimal number (max precision is 32) |
| `MONEY` | Equivalent to `DECIMAL(16,2)` |
| `MONEY(p)` | Equivalent to `DECIMAL(p,2)` (max precision is 32) |
| `MONEY(p,s)` | Equivalent to `DECIMAL(p,s)` (max precision is 32) |
| `REAL / SMALLFLOAT` | 32-bit floating point decimal (C float) |
| `DOUBLE PRECISION / FLOAT[(n)]` | 64-bit floating point decimal (C double) |

## PostgreSQL

PostgreSQL supports the following
data types to store numbers:

| PostgreSQL data type | Description |
| --- | --- |
| `INT2/SMALLINT` | 16 bit signed integer |
| `INT4/INTEGER` | 32 bit signed integer |
| `INT8/BIGINT` | 64 bit signed integer |
| `DECIMAL/NUMERIC(p,s)` | Decimals with precision and scale (fractional part) |
| `DECIMAL/NUMERIC(p)` | Integers with p digits (no fractional part) |
| `DECIMAL/NUMERIC` | Floating point numbers (no limit) |
| `DECIMAL/NUMERIC(p<s,s)` | Decimals with precision < scale to force rounding |
| `DECIMAL/NUMERIC(p,s<0)` | Decimals with negative scale to force rounding |
| `FLOAT4/REAL` | 16 bit variable precision, floating point type |
| `FLOAT8/DOUBLE PRECISION` | 32 bit variable precision, floating point type |
| `MONEY` | 64 bit fixed monetary type |

ANSI types like `SMALLINT`, `INTEGER`, `FLOAT` are
supported by PostgreSQL as aliases to `INT2`, `INT4` and
`FLOAT8` native types.

Informix `DECIMAL(p)` floating point
types are converted to `DECIMAL` without precision/scale, to store any floating point
number in PostgreSQL.

The PostgreSQL `MONEY` type is based on 8 bytes and has a smaler range of possible
values than the Informix `MONEY(p,s)` type. Because PostgreSQL `MONEY`
type depends on the `lc_monetary` database setting, database designers usually prefer
using the `DECIMAL/NUMERIC(p,s)` data type to store monetary data.

## Solution

Use the following conversion rules to map Informix numeric types to PostgreSQL numeric types:

| Informix data type | PostgreSQL data type |
| --- | --- |
| `SMALLINT` | `INT2` |
| `INTEGER` | `INT4` |
| `INT8 / BIGINT` | `INT8` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p)` | `DECIMAL` |
| `DECIMAL` | `DECIMAL` |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY` | `DECIMAL(16,2)` |
| `SMALLFLOAT` | `FLOAT4` |
| `FLOAT[(n)]` | `FLOAT8` |

The PostgreSQL `MONEY` type it is not the exact equivalent to the Informix
`MONEY(p,s)` type having same precision and storage possibilities than the
`DECIMAL(p,s)` type. Consequently, to support all possible Informix
`MONEY(p,s)` cases, we have to use PostgreSQL `DECIMAL(p,s)` for
original `MONEY(p,s)` types. In addition, the PostgreSQL `MONEY` type
has risky behavior, when changing the `lc_monetary` database setting.

> **Note:**
>
> Avoid using `DECIMAL[(p)]` type in FGL or SQL: Due to the implementation
> differences in Informix SQL / Genero BDL and the native SQL type, such data type is not recommended.
> Always specify a precision and scale with `DECIMAL(p,s)`.

The numeric types translation can be
controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.smallint = { true | false }
dbi.database.dsname.ifxemul.datatype.integer = { true | false }
dbi.database.dsname.ifxemul.datatype.bigint = { true | false }
dbi.database.dsname.ifxemul.datatype.int8 = { true | false }
dbi.database.dsname.ifxemul.datatype.decimal = { true | false }
dbi.database.dsname.ifxemul.datatype.money = { true | false }
dbi.database.dsname.ifxemul.datatype.float = { true | false }
dbi.database.dsname.ifxemul.datatype.smallfloat = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")

[Type conversion rules](1452-type-conversion-rules.md "Type conversion rules")
