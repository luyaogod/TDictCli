---
title: "CHAR and VARCHAR data types"
source: "fgl-topics/c_fgl_odiagora_016.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > CHAR and VARCHAR data types"
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

## ORACLE

Oracle® supports the following data
types to store character data:

| Oracle data type | Description |
| --- | --- |
| `CHAR(n)` | SBCS or MBCS character data using the database character set, where n is specified in bytes or characters, based on the length semantics (max is 2000 bytes) |
| `VARCHAR2(n)` | SBCS or MBCS character data using the database character set, where n is specified in bytes or characters, based on the length semantics (max is 4000 bytes) |
| `NCHAR(n)` | SBCS or MBCS character data using the national character set, where n is specified in bytes or characters, based on the length semantics (max is 2000 bytes) |
| `NVARCHAR2(n)` | SBCS or MBCS character data using the national character set, where n is specified in bytes or characters, based on the length semantics (max is 4000 bytes) |

> **Note:**
>
> Oracle supports extended character
> types with the `MAX_STRING_SIZE=EXTENDED` server parameter. Use
> `VARCHAR2` type can get a size up to 32Kb when
> `MAX_STRING_SIZE=EXTENDED` is set. However, the storage technique used by Oracle for such a large string type is
> different from the native/standard `VARCHAR2(4000)` type. Large character strings
> will be stored as LOBs. Extended character types are not supported by Genero's Oracle database driver.

In Oracle
`CHAR(N)/VARCHAR2(N)` types, the size N can be specified in character or byte units,
depending on length semantics settings. See [Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md) for more
details.

When comparing `CHAR` and `VARCHAR2` values in Oracle, the trailing
blanks are significant; this is not the case when using Informix
`VARCHAR` columns. However, before comparing string values, Oracle blank-pads `CHAR(N)` data to the maximum
length of both operands. As result, it looks like trailing blanks are not significant in
`CHAR(N)` comparison. For example, a column defined as `CHAR(5)` with
the value `'abc '` (with 2 trailing blanks) will not be equal to
`'abc'`, but when comparing `(col = 'abc')`, Oracle will add 2 blanks to the right operand and values will
match. Blank padding does not occur for `VARCHAR2()` data, as result, the expression
`(col = 'abc')` will be false, if col `VARCHAR2` does not exactly
contain the value `'abc'`. For more details, see blank-padded and non-padded
comparison semantics in Oracle
documentation.

Oracle treats empty strings like
`NULL` values; Informix doesn't. See issue
[Empty Character Strings](1396-empty-character-strings.md) for more
details.

With Oracle, you can define a Database Character Set and a National Character Set.Oracle uses the Database Character Set to
store string data in the `CHAR/VARCHAR2` columns, and uses the National Character Set
for `NCHAR/NVARCHAR2` columns.

## Solution

Informix
`CHAR(N)` types must be mapped to Oracle
`CHAR(N)` types, and Informix
`VARCHAR(N)` or `LVARCHAR(N)` columns must be mapped to Oracle
`VARCHAR2(N)`.

Table columns using a different character encoding than the database is
not supported with Genero: All table columns must use the same character encoding defined at the
database level.

Check that your database tables do not use `CHAR`, `VARCHAR` or
`LVARCHAR` types with a length exceeding the Oracle limits of `CHAR/VARCHAR2`.

When using a multibyte character set (such as UTF-8), configure Oracle to use character length semantics, define
`CHAR/VARCHAR2` database columns with a size in character units, and use character
length semantics in BDL programs with FGL\_LENGTH\_SEMANTICS=CHAR. See [Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md) for more details.

When extracting a database schema from an Oracle database, the [fgldbsch
schema extractor](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") uses the size of the column in characters, not the octet length. If you have
created a `CHAR(10 (characters) )` column a in the database, the
.sch file will get a size of 10, that will be interpreted according to
FGL\_LENGTH\_SEMANTICS as a number of bytes or characters.

The Oracle client character set must correspond to the Genero runtime system
locale (LANG/LC\_ALL). You can define the Oracle client character set with the NLS\_LANG environment variable.

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
