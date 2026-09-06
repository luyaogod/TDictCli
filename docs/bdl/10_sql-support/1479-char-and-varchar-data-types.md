---
title: "CHAR and VARCHAR data types"
source: "fgl-topics/c_fgl_odiagsqt_014.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > CHAR and VARCHAR data types"
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

## SQLite

SQLite 3 provides the `TEXT` native data type with no strict size limitation.
SQLite allows the `CHAR(n)`, `VARCHAR(n)`, `NCHAR(n)`
and `NVARCHAR(n)` type names to be used, but actually stores the data in a
`TEXT` native type.

SQLite treats empty strings as `NOT NULL` values like Informix.

> **Note:**
>
> With the default `BINARY` collation, SQLite compares
> `VARCHAR` and `CHAR` values by taking trailing blanks into
> account. Informix always ignores trailing blanks when
> comparing `CHAR/VARCHAR` values.

SQLite supports only the UTF-8 character encoding. Thus, client
applications must provide UTF-8 encoded strings.

## Solution

The database interface supports character string variables in SQL statements for input (BDL
`USING`) and output (BDL `INTO`).

> **Important:**
>
> With the default `BINARY` collation, `CHAR` and
> `VARCHAR` comparison in SQLite takes trailing blanks into account. As result, some
> queries returning rows with Informix may not return the
> same result set with SQLite. When creating a table in SQLite, you can change the default collation
> rule to force the database engine to trim trailing blanks before comparing
> `CHAR/VARCHAR` values, by specifying `COLLATION RTRIM` in the column
> definitions. When creating a table from a Genero program, if [Informix emulation](1059-database-connections.md "Explains how to manage database connections in a program.") is enabled for the `CHAR/VARCHAR` types, the SQLite
> database driver adds automatically `COLLATE RTRIM` after the `CHAR(N)`
> or `VARCHAR(N)` type, to get the same comparison semantics as Informix.

Table columns using a different character encoding than the database is
not supported with Genero: All table columns must use the same character encoding defined at the
database level.

Regarding character sets, the SQLite database driver automatically converts character strings
used in the programs to/from UTF-8 for SQLite.

SQLite uses character length semantics: When you define a `CHAR(20)` and the database
character set is multibyte, the column can hold more bytes/characters than the Informix `CHAR(20)` type, when using byte length semantics.

When using a multibyte character set (such as UTF-8), define database columns
with the size in character units, and use character
length semantics in BDL programs with FGL\_LENGTH\_SEMANTICS=CHAR.

When extracting a database schema from a SQLite database, the [fgldbsch schema extractor](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") uses the
size of the column in characters, not the octet length. If you have created a `CHAR(10
(characters) )` column a in SQLite database using the UTF-8 character set, the .sch file
will get a size of 10, that will be interpreted following FGL\_LENGTH\_SEMANTICS, as a number of bytes
or characters.

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
