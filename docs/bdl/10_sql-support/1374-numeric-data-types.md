---
title: "Numeric data types"
source: "fgl-topics/c_fgl_odiagora_030.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > Numeric data types"
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

## ORACLE

Oracle® supports following data types to
store numbers:

| Oracle data type | Description |
| --- | --- |
| `NUMBER(p,s) (1<=p<= 38, -84<=s<=127)` | Fixed point decimal numbers. |
| `NUMBER(p) (1<=p<= 38)` | Integer numbers with a precision of p digits. |
| `NUMBER(*,s)` | Fixed point decimal numbers with a precision of 38 digits. |
| `NUMBER` | Floating point decimals with a precision of 38 digits. |
| `FLOAT(b) (1<=b<= 126)` | Floating point numbers with a binary precision b. This is a sub-type of NUMBER. |
| `BINARY_FLOAT` | 32-bit floating point number. |
| `BINARY_DOUBLE` | 64-bit floating point number. |

The type names `SMALLINT`, `INTEGER` are supported by Oracle. However, these will be converted to
the native `NUMBER(*,0)` type. When dividing `INTEGER` or
`SMALLINT` types, Informix rounds the
result ( 7 / 2 = 3 ), while Oracle
doesn't, because it does not have a native integer data type ( 7 / 2 = 3.5 )

The `DECIMAL` type name is also supported by Oracle, and is mapped to the native `NUMBER` type. When using
a precision and scale (`DECIMAL(p,s)`), the resulting `NUMBER(p,s)`
type is equivalent. However, `DECIMAL(p)` becomes a `NUMBER(p,0)`, and
`DECIMAL` without precision / scale becomes a `NUMBER(*,0)` both
storing whole numbers. This is different from Informix SQL, where a `DECIMAL` (with
or without precision) can store real numbers.
> **Important:**
>
> Oracle `NUMBER`
> without precision / scale can store real numbers, while Oracle `DECIMAL` will store
> whole numbers. For
> example:
>
> ```
> CREATE TABLE t1 ( num  NUMBER, dec DECIMAL );
> INSERT INTO t1 VALUES ( 123.456, 123.456 );
> SELECT * FROM t1;
>        NUM	  DEC
> ---------- ----------
>    123.456	  123
> ```

## Solution

Use the following conversion rules to map Informix numeric types to Oracle numeric types:

| Informix data type | Oracle data type |
| --- | --- |
| `SMALLINT` | `NUMBER(5,0)` |
| `INTEGER` | `NUMBER(10,0)` |
| `BIGINT` | `NUMBER(20,0)` |
| `INT8` | `NUMBER(20,0)` |
| `DECIMAL(p,s)` | `NUMBER(p,s)` |
| `DECIMAL(p)` | `FLOAT(p * 3.32193)` |
| `DECIMAL` (not recommended) | `FLOAT` |
| `MONEY(p,s)` | `NUMBER(p,s)` |
| `MONEY(p)` | `NUMBER(p,2)` |
| `MONEY` | `NUMBER(16,2)` |
| `SMALLFLOAT` | `BINARY_FLOAT` |
| `FLOAT[(p)]` | `BINARY_DOUBLE` |

Avoid dividing integers in SQL statements. If you do divide an integer, use the
`TRUNC()` function with Oracle.

When creating a table directly in Oracle's sqlplus with the
`INTEGER`, `SMALLINT` types, Oracle will create columns with the
native `NUMBER(38,0)` type. As result, it is not possible (for
fgldbsch) to distinguish the original type names used in `CREATE
TABLE` from the native `NUMBER(38,0)` type or `NUMBER(38)`
type (where scale defaults to zero). In the next example, all columns will be of type
`NUMBER(38,0)`:

```
$ sqlplus ...
sql> CREATE TABLE mytab (
   col1 INTEGER,
   col2 SMALLINT,
   col3 NUMBER(38),
   ...
```

When extracting the database schema with fgldbsch, `NUMBER`,
`NUMBER(p>32)` and `NUMBER(p>32,s)` types will by default give an
extraction error. However, these types can be converted to `DECIMAL(32)` and
`DECIMAL(32,s)` with the `-cv` option, by using the "B" character at
positions 22 (for `NUMBER`) and 23 (for
`NUMBER(p>32[,s])`).
> **Note:**
>
> When fetching a
> `NUMBER[(p>32,s)]` into a BDL
> `DECIMAL(32[,s])` type, if the value stored in the
> `NUMBER` column has more than 32 digits, it will be rounded to fit into a
> `DECIMAL(32)`, or the overflow error -1226 will occur when fetching into a
> `DECIMAL(32,s)`. Note that it must be allowed to fetch numeric expressions such as
> 1/3 (=0.333333333333....) into a `DECIMAL(p,s)`, even if such expression will produce
> more than 32 digits with Oracle.

When creating a table in a BDL program with `DECIMAL(p)`, this type is converted
to native Oracle
`FLOAT(p*3.32193)`. When creating a table in a BDL program with
`DECIMAL` (without precision) this type is converted to native Oracle
`FLOAT`. The native Oracle
`FLOAT[(p)]` type can be extracted by fgldbsch, but
Oracle's `FLOAT` has a higher precision than the BDL
`DECIMAL` type, which can lead to value rounding when fetching
rows.

When casting a numeric expression such as `CAST(SUM(col) AS DECIMAL)`, with
Informix SQL this results in a real numeric value. With Oracle the `DECIMAL` type becomes a
`NUMBER(*,0)` and the result is a whole number. The `DECIMAL` type
name in the `CAST()` expression is not translated by the Oracle drivers because this
type conversion is only done for DDL statement (`CREATE TABLE / ALTER TABLE`).

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

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
