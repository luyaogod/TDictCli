---
title: "Numeric data types"
source: "fgl-topics/c_fgl_odiagsqt_017.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > Numeric data types"
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

## SQLite

SQLite 3 supports `INTEGER` (8 byte integer) and `REAL` (8 byte
floating point) as native types to store numbers, but allows also synonyms:

| Supported synonyms | SQLite type affinity |
| --- | --- |
| `INT, INTEGER, TINYINT, SMALLINT, MEDIUMINT, BIGINT, UNSIGNED BIG INT, INT2, INT8` | `INTEGER` (8 bytes!) |
| `REAL, DOUBLE, DOUBLE PRECISION, FLOAT` | `REAL` (8 bytes!) |
| `DECIMAL(p,s), NUMERIC` | `NUMERIC` (based on `REAL`) |

> **Important:**
>
> Exact decimal types like `DECIMAL(p,s)` may be stored as
> floating point numbers (`REAL`), `INTEGER` or `TEXT`
> types, according to the type affinity selected by SQLite. When converted to floating point type,
> data loss and rounding rule differences are possible with SQLite.

## Solution

Informix numeric types are not translated by the
SQLite database driver: The numeric types are used as is when creating tables, since SQLite supports
a wide range of type synonyms.

Since SQLite 3 does not have exact decimal types like `DECIMAL(p,s)`, you must pay
attention to the rounding rules and data loss when using numbers with many significant digits.
Arithmetic operations like division have different results than with Informix. It is better to fetch the original column value into a
`DECIMAL` variable, and do arithmetic operations in the application program.

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
