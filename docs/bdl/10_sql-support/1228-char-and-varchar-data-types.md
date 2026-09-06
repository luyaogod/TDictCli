---
title: "CHAR and VARCHAR data types"
source: "fgl-topics/c_fgl_odiagdmg_015.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > CHAR and VARCHAR data types"
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

## Dameng®

Dameng supports following data types to store character data:

| Dameng data type | Description |
| --- | --- |
| `CHAR(n)` | To store fixed-length character strings in the database encoding. The size is specified in byte units. |
| `VARCHAR(n)` | To store variable-length character strings in the database encoding. The size is specified in byte units. |
| `VARCHAR2(n)` | This is an alias for `VARCHAR`. |

The character set used by Dameng to store `CHAR` and `VARCHAR` data
is defined with the `CHARSET` parameter when creating a new database.
`CHARSET=1` defines a UTF-8 character set.

The `CASE_SENSITIVE` database parameter defines if character string comparison and
database object names are case sensitive (Y) or not case sensitive (N).

## Solution

Informix `CHAR(n)` columns can be
converted to Dameng `CHAR(n)` types, and
Informix `VARCHAR(n)` or `LVARCHAR(n)` columns can be converted to
Dameng `VARCHAR(n)`.

Configure the Dameng database for UTF-8 (`CHARSET=1`), using byte length semantics
and case insensitive character string comparison (`CASE_SENSITIVE=N`)

> **Important:**
>
> For `CHAR` and `VARCHAR` columns, the DPI client API function
> `dpi_desc_column()` returns a size always expressed in bytes. As result, a
> `CHAR(10)` column gives 10 bytes, while a column of a type
> `VARCHAR(10)` gives a size of 40 bytes, when using the database parameters
> `CHARSET=1` (UTF-8) and byte length semantics. For string literals, the size is also
> returned in bytes, and the SQL type is `DSQL_VARCHAR`, the same type id than for a
> `VARCHAR` columns. Consequently, it is not possible for the Dameng ODI driver to
> deduce a program variable size that would correspond to a number of UTF-8 characters, as with other
> DBs using CHAR length semantics with both `CHAR` and `VARCHAR`
> columns. Therefore, the only supported length semantics option (with UTF-8 encoding) is the BYTE
> length semantics, for Genero BDL (`FGL_LENGTH_SEMANTICS=BYTE`) and the Dameng
> database.

Consider using ASCII characters in `CHAR(n)` columns, to have the size matching
the number of characters, and store the UTF-8 data in `VARCHAR(n)` columns.

Use `CHAR(n)` and `VARCHAR(n)` in Dameng with and
`CHAR(n)` and `VARCHAR(n)` for program variables, where n is expressed
as an number of bytes in both context. In UTF-8, you can store n ASCII chars, n/2 Latin accute chars
and n/3 Asian chars.

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
