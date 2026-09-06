---
title: "CHAR and VARCHAR data types"
source: "fgl-topics/c_fgl_odiagmys_016.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > CHAR and VARCHAR data types"
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

## Oracle® MySQL and MariaDB

MySQL and MariaDB support the following data types to store character data:

| MySQL data type | Description |
| --- | --- |
| `CHAR(n)` | SBCS or MBCS character data using the database character set, where n is specified in characters (max is 255 characters) |
| `VARCHAR(n)` | SBCS or MBCS character data using the database character set, where n is specified in characters (max is 65535 characters) |
| `NCHAR(n)` | SBCS or MBCS character data using the national character set, where n is specified in characters (max is 255 characters) |
| `NVARCHAR(n)` | SBCS or MBCS character data using the national character set, where n is specified in characters (max is 65535 characters) |
| `TEXT` | The Large Object type to store SBCS or MBCS character data using the database character set |

MySQL uses character length semantics to define the size of `CHAR/VARCHAR`
columns, while Informix and Genero use Byte Length
Semantics.

MySQL can support multiple character sets, you can run the `SHOW CHARACTER SET`
statement to list supported encoding. There are different configuration levels to define the character
set used by MySQL to store data. The server character set defines the default for
database character sets if not specified in the `CREATE DATABASE` command.
You can even define a specific character set at the table and column level, but this is not
recommended with Genero applications. The database character set is used to store `CHAR`
and `VARCHAR` columns. The `NCHAR` and `NATIONAL VARCHAR`
types use a predefined character set which can be different from the database character set.
In MySQL the national character set is UTF-8.

MySQL can automatically convert from/to the client and server characters sets. In the client
applications, the character set can be defined with the `default-character-set`
configuration option, or with the `SET NAMES` SQL instruction (but that instruction
is not supported with Genero see below).

> **Important:**
>
> If the `STRICT_TRANS_TABLES`
> option is not defined in the `sql_mode` parameter, MySQL truncates character
> strings, when the value is too large for the target column. However, the
> `STRICT_TRANS_TABLES` option controls also numeric data truncation/overflow. This
> option should be used, to avoid numeric data truncation/overflow being ignored (with only an SQL
> warning), and to produce an SQL error instead when the numeric value does not fit into the target
> column type.

Note that by default, when fetching `CHAR` columns from MySQL, trailing blanks are
trimmed. This does not matter as long as you fetch `CHAR` columns into
`CHAR` variables, but this non-standard behavior will impact `CHAR`
fetch into `VARCHAR`, or other SQL areas such as string concatenation for example.
You can control the behavior of `CHAR` trailing blanks trimming with the
`PAD_CHAR_TO_FULL_LENGTH` option of the `sql_mode` parameter. But when
this mode is used, the result of the SQL `LENGTH()` function will be different since
trailing blanks are significant for that function in MySQL.

## Solution

Informix `CHAR(N)` types
must be mapped to MySQL `CHAR(N)` types. Informix `VARCHAR(N)` or `LVARCHAR(N)` types must be mapped to
MySQL `VARCHAR(N)`.

Table columns using a different character encoding than the database is
not supported with Genero: All table columns must use the same character encoding defined at the
database level.

Store single-byte or multibyte character strings in MySQL `CHAR`,
`VARCHAR` and `TEXT` columns.

Define the DB client application character set, to make it correspond to the runtime system
character set.
> **Important:**
>
> MySQL/MariaDB support the `SET NAMES` SQL instruction,
> to change the character set after connecting to the database. However, this is not supported with
> Genero: The driver needs to know the character set at connection initialization. Use the
> `default-character-set` configuration option.

With MySQL/MariaDB, the client application character set must be defined with the
`default-character-set` option, under the `[client]` section of the
configuration file, for
example:

```
[client]
default-character-set="utf8mb4"
```

> **Note:**
>
> A specific MySQL/MariaDB configuration file can be defined with an FGLPROFILE entry, see [Oracle MySQL specific FGLPROFILE parameters](1082-oracle-mysql-specific-fglprofile-parameters.md), [MariaDB specific FGLPROFILE parameters](1083-mariadb-specific-fglprofile-parameters.md).

MySQL uses character length semantics: When you define a `CHAR(20)` and the database
character set is multibyte, the column can hold more bytes/characters than the Informix `CHAR(20)` type, when using byte length semantics.
When using a multibyte character set (such as UTF-8), define database columns with the size in
character units, and use character length semantics in BDL programs with FGL\_LENGTH\_SEMANTICS=CHAR.

When extracting a database schema from a MySQL database, the [fgldbsch schema extractor](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") uses the
size of the column in characters, not the octet length. If you have created a `CHAR(10
(characters) )` column in a MySQL database using the UTF-8 character set, the
.sch file will get a size of 10, which will be interpreted depending on
FGL\_LENGTH\_SEMANTICS as a number of bytes or characters.

Review your database schema when using `CHAR` columns with a size exceeding the
MySQL limits: If you need to store `CHAR` character strings larger as the MySQL
`CHAR` limit, you can use the MySQL `TEXT` type. However, as
of MySQL version 5.0.3 (supporting large `VARCHAR` sizes), as long as you use short
sizes for `CHAR` (<100c), the character types can be used as is in MySQL.

The `CHAR(N>255)` types are converted by the SQL Translator to a MySQL
`TEXT` type, because MySQL `CHAR` type has a limit of 255 characters.
When designing a database, consider using `CHAR` only for short character string data
storage (<50c), and use `VARCHAR` for larger character string data storage (name,
address, comments).

> **Note:**
>
> For each `TEXT` column fetched from MySQL, the MySQL database driver needs
> to allocate a temporary string buffer of 65535 bytes. The memory used by this temporary buffer is
> freed when freeing the cursor.

When using `VARCHAR` types, the SQL Translator leaves the type definition as is,
even for N > 255, assuming that the target MySQL server version is at least 5.0.3 (supporting
`VARCHAR(N)` up to 65535 characters).

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
