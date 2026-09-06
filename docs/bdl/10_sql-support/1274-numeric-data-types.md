---
title: "Numeric data types"
source: "fgl-topics/c_fgl_odiagmsv_027.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > Numeric data types"
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

## Microsoft™ SQL Server

Microsoft SQL Server supports the following numeric
data types:

| Microsoft SQL Server data type | Description |
| --- | --- |
| `SMALLINT` | 16 bit signed integer |
| `INTEGER` | 32 bit signed integer |
| `BIGINT` | 64 bit signed integer |
| `DECIMAL(p,s)` | Fixed point decimal. |
| `SMALLMONEY` | 32-bit floating point decimal with currency |
| `MONEY` | 64-bit floating point decimal with currency |
| `REAL` | 32-bit floating point decimal (C float) |
| `FLOAT[(n)] (DOUBLE)` | 64-bit floating point decimal (C double) |

Notes about SQL Server `DECIMAL` type:

- Without any decimal storage specification, the precision defaults to 18 and the scale defaults
  to zero:
  - `DECIMAL` in SQL Server = `DECIMAL(18,0)` in Genero BDL.
  - `DECIMAL(p)` in SQL Server = `DECIMAL(p,0)` in Genero BDL.
- The maximum precision is 38.

Notes about the SQL Server `MONEY` and `SMALLMONEY` types:

- `SMALLMONEY` values are encoded on 4 bytes (like Informix
  `SMALLFLOAT`) and `MONEY` values are encoded on 8 bytes (like Informix
  `FLOAT`). This implies different numeric encoding and rounding rules as with
  `DECIMAL(p,s)` types.
- The value range for SQL Server `SMALLMONEY` is `-214,748.3648` to
  `214,748.3647` and the value range of SQL Server `MONEY` is
  `-922,337,203,685,477.5808` to `922,337,203,685,477.5807`. As result,
  for example, a BDL variable defined as `MONEY(19,4)` that holds the value
  `923,000,000,000,000.0000` cannot be stored with an SQL Server
  `MONEY`.
- SQL Server `MONEY` and `SMALLMONEY` data types are accurate to a
  ten-thousandth of the monetary units that they represent: The number of decimal digits is 4. As
  result, it's for example not possible to store BDL `MONEY(15,5)` values in SQL Server
  `MONEY` columns.
- The currency symbol handling is different.

For these reasons, it's not recommended to use SQL Server `MONEY` or
`SMALLMONEY` data types to store Informix
`MONEY(p,s)` values.

## Solution

Use the following conversion rules to map Informix numeric types to SQL Server numeric types:

| Informix | Microsoft SQL SERVER |
| --- | --- |
| `SMALLINT` | `SMALLINT` |
| `INTEGER` | `INT / INTEGER` |
| `BIGINT` | `BIGINT` |
| `INT8` | `BIGINT` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p<=16)` | `DECIMAL(2*p,p)` |
| `DECIMAL(p>16)` | N/A |
| `DECIMAL` | `DECIMAL(32,16)` |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY` | `DECIMAL(16,2)` |
| `REAL / SMALLFLOAT` | `REAL` |
| `FLOAT[(n)] / DOUBLE PRECISION` | `FLOAT(n)` (Where n must be from 1 to 15) |

When creating tables from BDL programs, the database interface automatically converts Informix numeric data types to corresponding Microsoft SQL Server data types. In database creation scripts, apply the
conversion rules as described in the above table.

> **Important:**
>
> There is no SQL Server equivalent for the Informix
> `DECIMAL(p)` floating point decimal (i.e. without a scale). If your application is
> using such data types, you must review the database schema in order to use SQL Server compatible
> types.

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
