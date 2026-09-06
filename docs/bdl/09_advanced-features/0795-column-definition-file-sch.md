---
title: "Column Definition File (.sch)"
source: "fgl-topics/c_fgl_DatabaseSchema_018.html"
breadcrumb: "Advanced features > Database schema > Structure of database schema files > Column Definition File (.sch)"
type: "concept"
---

# Column Definition File (.sch)

> The .sch database schema file contains the data types of database table columns.

## Description

The [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") of program [variables](../08_language-basics/0686-variables.md "Explains how to define program variables.") or [form fields](../11_user-interface/1670-form-fields.md "Form fields are form elements designed for data input and/or data display.") used to hold data of a given database
column must match the data type used in the database. The definition of these elements is simplified
by centralizing the information in external .sch files, which contain column
data types.

The .sch file can be produced by the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, from the system views of the target database.

In form files, you can directly specify the table and column name in the field definition in the
[`ATTRIBUTES`](../11_user-interface/1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.") section
of forms.

In programs, you can define variables with the data type of a database column by using the [`LIKE`](../08_language-basics/0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.") keyword.

As column data types are extracted from the database system tables, you may get different results
with different database servers. For example, Informix®
provides the `DATE` data type to store simple dates in year, month, and day format (=
`DATE` FGL type), while Oracle® stores dates as year to second ( = `DATETIME YEAR TO SECOND` FGL
type).

The table describes the fields you will find in a row of the .sch file:

| Pos | Type | Description |
| --- | --- | --- |
| 1 | `STRING` | Database table name. |
| 2 | `STRING` | Column name. |
| 3 | `SMALLINT` | Coded column data type. If the column is `NOT NULL`, you must add 256 to the value. |
| 4 | `SMALLINT` | Coded data type length. |
| 5 | `SMALLINT` | Ordinal position of the column in the table. |

This table shows the data types and their corresponding type code that can be present
in a schema file:

| Data type name | Data type code (field #3) | Data type length (field #4) |
| --- | --- | --- |
| `CHAR` | 0 | Maximum number of characters or bytes (see note) |
| `SMALLINT` | 1 | Fixed length of 2 |
| `INTEGER` | 2 | Fixed length of 4 |
| `FLOAT` / `DOUBLE` `PRECISION` | 3 | Fixed length of 8 |
| `SMALLFLOAT` / `REAL` | 4 | Fixed length of 4 |
| `DECIMAL` | 5 | If the decimal is defined with a precision and scale, the length is computed using this formula:`length = (precision * 256) + scale`If the decimal is defined as a floating point decimal (i.e. with no scale), the length is computed as follows:`length = (precision * 256) + 255` |
| `SERIAL` | 6 | Fixed length of 4 |
| `DATE` | 7 | Fixed length of 4 |
| `MONEY` | 8 | The length is computed using this formula:`length = (precision * 256) + scale`A `MONEY` cannot be defined with a floating point, is has always a scale. |
| *Unused* | 9 | *N/A* |
| `DATETIME` | 10 | For `DATETIME` types, the length is determined using the formula:`length = ( digits * 256 ) + ( qual1 * 16 ) + qual2`where digits is the total number of digits used when displaying the datetime value. For example, a `DATETIME YEAR TO MINUTE` (YYYY-MM-DD hh:mm) uses 12 digits.The qual1 and qual2 elements identify datetime qualifiers using the following codes:`0 = YEAR``2 = MONTH``4 = DAY``6 = HOUR``8 = MINUTE``10 = SECOND``11 = FRACTION(1)``12 = FRACTION(2)``13 = FRACTION(3)``14 = FRACTION(4)``15 = FRACTION(5)`For example, a DATETIME YEAR TO MINUTE size length is computed as follows:`(12 * 256) + (0 * 16) + 8 = 3080` |
| `BYTE` | 11 | Length of descriptor |
| `TEXT` | 12 | Length of descriptor |
| `VARCHAR` | 13 | Maximum number of characters or bytes (see note)If the length is positive:`length = ( min_space * 256 ) + max_size`If length is negative:`length + 65536 = ( min_space * 256 ) + max_size` |
| `INTERVAL` | 14 | For `INTERVAL` types, the length is determined using the following formula:`length = ( digits * 256 ) + ( qual1 * 16 ) + qual2`where digits is the total number of digits used when displaying the interval value. For example, a `INTERVAL HOUR(5) TO FRACTION(3)` (hhhhh:mm:ss.fff) uses 12 digits.The qual1 and qual2 elements identify datetime qualifiers according to this list:`0 = YEAR``2 = MONTH``4 = DAY``6 = HOUR``8 = MINUTE``10 = SECOND``11 = FRACTION(1)``12 = FRACTION(2)``13 = FRACTION(3)``14 = FRACTION(4)``15 = FRACTION(5)`For example, an `INTERVAL HOUR(5) TO FRACTION(3)` size length is computed as follows:`(12 * 256) + (6 * 16) + 13 = 3181` |
| `NCHAR` | 15 | Maximum number of characters or bytes (see note) |
| `NVARCHAR` | 16 | Maximum number of characters or bytes (see note) |
| `INT8` | 17 | Fixed length of 10 (size of int8 structure)In programs, will be converted to a `BIGINT` type. |
| `SERIAL8` | 18 | Fixed length of 10 (size of int8 structure)In programs, will be converted to `BIGINT` type. |
| `BOOLEAN` (`SQLBOOL`) | 45 | Boolean type, in the meaning of Informix front-end `SQLBOOL` (sqltype.h) |
| `BIGINT` | 52 | Fixed length of 8 (bytes) |
| `BIGSERIAL` | 53 | Fixed length of 8 (bytes) |
| `VARCHAR2` | 201 | Maximum number of charactersIn programs, will be converted to a `VARCHAR` type. |
| `NVARCHAR2` | 202 | Maximum number of charactersIn programs, will be converted to a `VARCHAR` type. |

> **Note:**
>
> Data type length (field #4) is a `SMALLINT` value encoding the
> length or composite length of the type. For character string types, the unit of the length used to
> define character program variables and form fields depends on the [length semantics](0881-length-semantics-settings.md).

## Informix SERIAL types

When the database schema defines `SERIAL`, `BIGSERIAL` or
`SERIAL8` types, form fields referencing the serial column will get the [`NOENTRY`](../11_user-interface/1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog.") attribute automatically,
except if defined with the `TYPE LIKE` syntax.

## Informix CHAR/VARCHAR types and SQL\_LOGICAL\_CHAR

The Informix `SQL_LOCICAL_CHARS`
onconfig parameter is taken into account when extracting
`CHAR/VARCHAR` column types, to convert the size in bytes from
`syscolumns.collength` to a number of characters for the .sch
file.

It is then possible to use FGL\_LENGTH\_SEMANTICS=CHAR when the application locale is UTF-8.

## Informix DISTINCT types

Informix IDS version 9.x and higher allow you to
define `DISTINCT` types from a base type with the `CREATE DISTINCT
TYPE` instruction. In the syscolumns table, Informix identifies distinct types in the coltype column by adding the 0x0800 bit (2048) to
the base type code. For example, a distinct type defined with the `VARCHAR`
built-in type (code 13) will be identified with the code 2061 (13 + 2048). Informix sets additional bits when the distinct type is based on the
`LVARCHAR` or `BOOLEAN` opaque types. If the base type is an
`LVARCHAR`, the type code used in coltype gets the 0x2000 bit set (8192) and when the
base type is `BOOLEAN`, the type code gets the 0x4000 bit (16384).

When extracting a schema from an Informix database
defining columns with `DISTINCT` types, the schema extractor will keep the original
type code of the distinct type in the .sch file for columns using distinct
types based on built-in types (with the 0x0800 bit set). Regarding the exception of opaque
types, `BOOLEAN`-based distinct types get the code 45 ( + 256 if `NOT
NULL`), and `LVARCHAR`-based distinct types are mapped to the code 201 (+ 256
if `NOT NULL`) if the `-cv` option enables conversion from
`LVARCHAR` to `VARCHAR2`.

The fglcomp and fglform compilers understand the distinct
type code bit 0x0800, so you can define program variables with a `DEFINE LIKE`
instruction based on a column that was created with a distinct Informix type.

## Example

```
customer^customer_num^258^4^1^
customer^customer_name^256^50^2^
customer^customer_address^0^100^3^
order^order_num^258^4^1^
order^order_custnum^258^4^2^
order^order_date^263^4^3^
order^order_total^261^1538^4^
```
