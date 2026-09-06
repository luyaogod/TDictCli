---
title: "CHAR and VARCHAR data types"
source: "fgl-topics/c_fgl_odiagpgs_016.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > CHAR and VARCHAR data types"
type: "concept"
description: "Informix® Informix supports the following character data types: Table 1. Informix character data types Informix data type Description CHAR(n) SBCS and MBCS character data (max is 32767 bytes) VARCHAR( ..."
---

# CHAR and VARCHAR data types

## Informix®

Informix supports the following character data types:

| Informix data type | Description |
| --- | --- |
| `CHAR(n)` | SBCS and MBCS character data (max is 32767 bytes) |
| `VARCHAR(n[,m])` | SBCS and MBCS character data (max is 255 bytes) |
| `NCHAR(n)` | Same as `CHAR`, with specific collation order |
| `NVARCHAR(n[,m])` | Same as `VARCHAR`, with specific collation order |
| `LVARCHAR(n)` | max size varies depending on the IDS version |

With Informix, both
`CHAR/VARCHAR` and `NCHAR/NVARCHAR` data types can be used to store
single-byte or multibyte encoded character strings. The only difference between
`CHAR/VARCHAR` and `NCHAR/NVARCHAR` is in how they use sorting:
`N[VAR]CHAR` types use the collation order, while `[VAR]CHAR` types
use the byte order.

The character set used to store strings in
`CHAR/VARCHAR/NCHAR/NVARCHAR` columns is defined by the DB\_LOCALE environment
variable.

The character set used by applications is defined by the CLIENT\_LOCALE environment
variable.

Informix uses Byte Length Semantics (the size
N that you specify in `[VAR]CHAR(N)` is expressed in bytes, not characters as in some
other databases)

## PostgreSQL

PostgreSQL supports following data types to store character data:

| PostgreSQL data type | Description |
| --- | --- |
| `CHAR(n)` | SBCS or MBCS character data using the database character set, where n is specified in characters (max is 10485760 characters) |
| `VARCHAR(n)` | SBCS or MBCS character data using the database character set, where n is specified in characters (max is 10485760 characters); The length specification is optional. |
| `TEXT` | SBCS or MBCS character data using the database character set (max is 1Gb) |

In PostgreSQL, `CHAR`, `VARCHAR` and `TEXT` types
store data in single byte or multibyte character sets. For `CHAR` and
`VARCHAR`, the size is specified in a number of characters, not bytes. The character
set used to store data for these types is defined by the database character set, which can be
specified when you create the database with the createdb tool or the
`CREATE DATABASE` SQL command.

> **Note:**
>
> The `VARCHAR` type of PostgreSQL can be used without a length
> specification. If no size is specified, the column accepts strings of any size. However, as Genero
> BDL needs to know the size of `CHAR` and `VARCHAR` columns to define
> fields and program variables from a schema file, it is not recommended to create tables in
> PostgreSQL having `VARCHAR` columns without size specification. If you try to extract
> a schema with fgldbsch, this tool will report that the `VARCHAR`
> column cannot be converted to a BDL type for [the .sch file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").

Automatic character set conversion between the PostgreSQL client and server is supported. With
PostgreSQL, the client charset can be defined in different ways, with the `SET
CLIENT_ENCODING TO` SQL command for example, or with configuration parameters. See the
PostgreSQL documentation for more details.

## Solution

Informix `CHAR(N)` types must be mapped
to PostgreSQL `CHAR(N)` types, and Informix `VARCHAR(N)` or `LVARCHAR(N)` columns must be mapped to
PostgreSQL `VARCHAR(N)`.

> **Note:**
>
> When creating a table from the BDL program with `NCHAR` or
> `NVARCHAR` types, the type names will be left as is and produce an SQL error because
> these types are not supported by PostgreSQL.

Table columns using a different character encoding than the database is
not supported with Genero: All table columns must use the same character encoding defined at the
database level.

Store single-byte or multibyte character strings in PostgreSQL `CHAR`,
`VARCHAR` and `TEXT` columns.

PostgreSQL uses character length semantics: When you define a `CHAR(20)` and the
database character set is multibyte, the column can hold more bytes/characters than the Informix `CHAR(20)` type, when using byte
length semantics.

When using a multibyte character set (such
as UTF-8), define database columns with the size in character
units, and use character length semantics in BDL programs with
FGL\_LENGTH\_SEMANTICS=CHAR.

When extracting a database schema from a PostgreSQL database, the [fgldbsch schema extractor](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") uses the
size of the column in characters, not the octet length. If you have created a `CHAR(10
(characters) )` column a in PostgreSQL database using the UTF-8 character set, the .sch file
will get a size of 10, that will be interpreted following FGL\_LENGTH\_SEMANTICS as a number of bytes
or characters.

Properly define the PostgreSQL database client character set, which must correspond to the
runtime system character set.

See also the section about [Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.").

The
`CHAR/VARCHAR` type translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.char = { true | false }
dbi.database.dsname.ifxemul.datatype.varchar = { true | false }
```

For
more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[CHAR and VARCHAR types](1012-char-and-varchar-types.md "Using the CHAR and VARCHAR data types with different databases.")

[Type conversion rules](1452-type-conversion-rules.md "Type conversion rules")
