---
title: "Numeric data types"
source: "fgl-topics/c_fgl_odiagmys_023.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > Numeric data types"
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

## Oracle® MySQL and MariaDB

MySQL and MariadDB support the following data types to store numbers:

| MySQL data type | Description |
| --- | --- |
| `SMALLINT` | 16 bit signed integer |
| `INTEGER` | 32 bit signed integer |
| `BIGINT` | 64 bit signed integer |
| `DECIMAL(p,s)` | Fixed point decimal. Maximum precision of recent versions is 65. |
| `DECIMAL(p)` | Stores whole numeric numbers up to p digits |
| `FLOAT[(M,D)]` | 32 bit floating point number |
| `DOUBLE[(M,D)]` | 64 bit floating point number |

The `STRICT_TRANS_TABLES` option in the `sql_mode` parameter
defines if numeric data truncation/overflow should produce an SQL error -1264, or just an SQL
warning -1265. To avoid problems you want to use the `STRICT_TRANS_TABLES`
option.

## Solution

Use the following conversion rules to map Informix numeric types to MySQL numeric types:

| Informix data type | MySQL equivalent |
| --- | --- |
| `SMALLINT` | `SMALLINT` |
| `INTEGER` | `INTEGER` |
| `INT8 / BIGINT` | `BIGINT` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p<=16)` | `DECIMAL(p*2,p)` |
| `DECIMAL(p>16)` | N/A |
| `MONEY` | `DECIMAL(16,2)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `SMALLFLOAT` | `REAL` |
| `FLOAT[(n)]` | `FLOAT[(n)] (DOUBLE)` |

> **Important:**
>
> If the `STRICT_TRANS_TABLES`
> option is not defined in the `sql_mode` parameter, MySQL truncates character
> strings, when the value is too large for the target column. However, the
> `STRICT_TRANS_TABLES` option controls also numeric data truncation/overflow. This
> option should be used, to avoid numeric data truncation/overflow being ignored (with only an SQL
> warning), and to produce an SQL error instead when the numeric value does not fit into the target
> column type.

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
