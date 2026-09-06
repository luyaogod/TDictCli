---
title: "CHAR and VARCHAR data types"
source: "fgl-topics/c_fgl_odiagmsv_017.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > CHAR and VARCHAR data types"
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

## Microsoft™ SQL Server

Microsoft SQL Server supports following data types to
store character data:

| Microsoft SQL Server data type | Description |
| --- | --- |
| `CHAR(n)` | Fixed character data type, where n is specified in bytes. Maximum storage size is 8000 bytes.Character encoding is defined by the database collation. Starting with SQL Server 2019, the encoding can be UTF-8, with a collation using the `_UTF8` option. |
| `VARCHAR(n)` | Variable character data type, where n is specified in bytes. Maximum storage size is 8000 bytes.Character encoding is defined by the database collation. Starting with SQL Server 2019, the encoding can be UTF-8, with a collation using the `_UTF8` option. |
| `VARCHAR(MAX)` | Large character data type. Maximum storage size is 2^31-1 bytes.Character encoding is defined by the database collation. Starting with SQL Server 2019, the encoding can be UTF-8, with a collation using the `_UTF8` option. |
| `NCHAR(n)` | Fixed character data type, where n is specified in byte-pairs. Maximum storage size is 4000 byte-pairs.Character encoding is UCS-2 or UTF-16. UTF-16 is enabled with a collation using the `_SC` option. |
| `NVARCHAR(n)` | Variable character data type, where n is specified in byte-pairs. Maximum storage size is 4000 byte-pairs.Character encoding is UCS-2 or UTF-16. UTF-16 is enabled with a collation using the `_SC` option. |
| `NVARCHAR(MAX)` | Large character type. Maximum storage size is 2^30-1 byte-pairs.Character encoding is UCS-2 or UTF-16. UTF-16 is enabled with a collation using the `_SC` option. |

To store large text data (LOBs), Microsoft SQL Server
(2005) provides the `VARCHAR(MAX)`/`NVARCHAR(MAX)` type as a
replacement for the old `TEXT`/`NTEXT` types. See [TEXT and BYTE (LOB) types](1282-text-and-byte-lob-types.md) for more details.

The use of `NCHAR`, `NVARCHAR` character types is the same as
`CHAR`, `VARCHAR` respectively, except:

- For `NCHAR` / `NVARCHAR`, the character encoding is UCS-2 or
  UTF-16, depending if the database collation is using the SC option.
- The length `N` in `N[VAR]CHAR(N)` defines a number of byte-pairs,
  not bytes. For UCS-2 this corresponds to a number of characters. But with UTF-16 there can be
  surrogate pairs using 4 bytes (2 byte-pairs).
- The maximum size of `NCHAR(N)` and `NVARCHAR(N)` column is 4000
  byte-pairs, compared to 8000 bytes for `CHAR/VARCHAR` using a single-byte character
  set.
- Unicode string literals are specified with a leading N. For example:
  N'日本語'
- The [LIKE statement](1295-matches-and-like.md) behaves
  differently with `CHAR` and `NCHAR` columns when using the N prefix
  before the search pattern.

Note that SQL Server uses Byte Length Semantics to define the size of
`CHAR/VARCHAR` columns, while `NCHAR` and `NVARCHAR`
sizes are expressed in byte-pair units, which for most cases (UCS-2) corresponds to a number of
characters (but this is not true for UTF-16 characters encoded in surrogate pairs).

SQL Server defines the character encoding for `CHAR` and `VARCHAR`
columns with the database collation, specified when creating a new database. When using the \_UTF8
modifier, character strings are UTF-8 encoded in `CHAR/VARCHAR` columns. In
`NCHAR/NVARCHAR` columns, character strings are always encoded in UCS-2 or UTF-16.
UTF-16 is used when the DB collation has the SC option.

Automatic charset conversion is supported by SQL Server between the client application and the
server. The client charset is defined by the Windows®
operating system, in the language settings for non-Unicode applications.

## Solution

Check that your database tables does not use `CHAR` or `VARCHAR`
types with a length exceeding the limits of SQL Server character types.

Table columns using a different character encoding than the database is
not supported with Genero: All table columns must use the same character encoding defined at the
database level.

Depending on the character set used by your application, either use `CHAR/VARCHAR`
or `NCHAR/NVARCHAR` columns with SQL Server:

- When the application charset is single-byte (like ISO-8859-15 or CP-1252), use
  `CHAR/VARCHAR` columns and define a database collation matching the application
  locale. In this configuration, the number of bytes corresponds to the number of characters, for both
  program `CHAR/VARCHAR` variables and database `CHAR/VARCHAR`
  columns.
- When the application locale is UTF-8, and the length semantics is in characters
  (FGL\_LENGTH\_SEMANTICS=CHAR), use `NCHAR/NVARCHAR` columns. In this case, the size of
  `CHAR/VARCHAR` program variables (as a number of chars) matches the size and unit of
  database columns.
- When the application locale is UTF-8, and the length semantics is in bytes
  (FGL\_LENGTH\_SEMANTICS=BYTE), use `CHAR/VARCHAR` columns with \_UTF8 collation. In this
  case, the size of `CHAR/VARCHAR` program variables (as a number of bytes) matches the
  size and unit of database columns.

Make sure that the regional language settings for non-Unicode applications corresponds to the
locale used by Genero programs. For more details about program locale settings, see [Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.").

The widechar mode of SQL Server drivers defines the ODBC SQL parameters and fetch buffers
bindings for character strings (SQL\_CHAR or SQL\_WCHAR), as well as automatic adding of N prefix in SQL
string literals. To get best behavior and performance of SQL Server, the widechar mode should
be enabled when using `NCHAR/NVARCHAR` columns, and it should be disabled when using
`CHAR/VARCHAR` columns (ref: FGL-315).

By Default, with all SQL Server ODI drivers, when not setting the FGLPROFILE
`*.widechar` option to false, the widechar mode is automatically
enabled, if the application locale is multibyte (UTF-8) and FGL\_LENGTH\_SEMANTICS=CHAR. Otherwise,
with SBCS locale (ISO-8859-15) or MBCS locale (UTF-8/BIG5) and BYTE length semantics and SQL Server
2019 + DB collation is \*\_UTF8, the widechar mode defaults to false. If the SQL Server
version is older than 2019 or the DB collation is NOT \*\_UTF8, for backward compatibility, the
widechar mode is enabled by default when the application locale is MBCS.

> **Important:**
>
> Setting FGLPROFILE `dbi.database.dbname.snc.widechar` to
> `false` is strongly discouraged and done at your own risk, when using a UTF-8
> application locale with a non-UTF-8 database collation for CHAR/VARCHAR column storage: Any UTF-8
> characters that cannot be represented in the database's collation will be converted to a question
> mark (`?`) when strings are sent to the database.

When extracting a database schema from an SQL Server database, the [fgldbsch schema extractor](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") uses the
size of the column in byte-pair units, not the byte length. If you have created an
`NCHAR(10)` column in an SQL Server database, the .sch file will
get a size of 10, that will be interpreted as a number of bytes or characters, depending on
FGL\_LENGTH\_SEMANTICS.

The
`CHAR/VARCHAR` type translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.char = { true | false }
dbi.database.dsname.ifxemul.datatype.varchar = { true | false }
```

For
more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Automatic conversion of CHAR/VARCHAR type names

By default, DDL statements (such as `CREATE TABLE`) executed by Genero programs
will not be scanned to convert the `CHAR`, `VARCHAR` and
`LVARCHAR` type names to `NCHAR/NVARCHAR` type names, even when the
widechar mode is used. Programs creating SQL tables in an SQL Server database using
`NCHAR/NVARCHAR` columns should create table columns with these national character
types.

To force the ODI drivers to replace `CHAR/VARCHAR/LVARCHAR` by
`NCHAR/NVARCHAR` type names (when the widechar mode is enabled), set the
[`dbi.database.dsname.ifxemul.nationalchars`](1079-ibm-informix-emulation-parameters-in-fglprofile.md) FGLPROFILE entry to
`true` (default is
`false`):

```
dbi.database.dsname.ifxemul.nationalchars = true
```

> **Note:**
>
> 1. The `dbi.database.dsname.ifxemul.nationalchars` will only take
>    effect, if the current application locale is multibyte (typically, UTF-8) and the
>    widechar mode is enabled (this is the default with UTF-8 and CHAR length
>    semantics).
> 2. The `dbi.database.dsname.ifxemul.nationalchars` parameter is
>    ignored, if the switch corresponding to the character type name
>    (`dbi.database.dsname.datatype.{char|varchar|text}`)
>    is set to `false`.

## Using the SNC driver

No specific character set configuration needs to be set with the SNC driver: The MS ODBC driver
detects the application locale automatically, by using `setlocale()` on Unix/Linux
and `GetACP()` on Windows.

According to the char or widechar mode, the SNC driver will do the
appropriate ODBC bindings, using `SQL_CHAR/SQL_VARCHAR` or
`SQL_WCHAR/SQL_WVARCHAR` SQL types, and let the MS ODBC driver do the charset
conversions when binding with `SQL_C_CHAR` for the C type, or do the appropriate
charset conversions when using `SQL_C_WCHAR` for the C type.

> **Important:**
>
> On Windows platforms, the ODBC data source option
> *"Perform translation for character data"* must be disabled, when using UTF-8 application
> locale defined by `LANG=.utf8`, NLS\_LENGTH\_SEMANTICS=BYTE, and the database has \_UTF8
> collation for CHAR/VARCHAR columns (the widechar mode is disabled).

The char / widechar modes can be forced with the following FGLPROFILE
entry:

```
dbi.database.dsname.snc.widechar= { true | false }
```

> **Important:**
>
> On Windows platforms, when not using the widechar mode, if the ODBC data source
> parameter *"Perform translation for character data"* is enabled, the MS ODBC driver makes
> charset conversions based on the system locale for non-unicode applications
> (`GetACP()`). The client encoding defined by `setlocale()` (set from
> LANG with Genero) does not apply. Therefore, it is not possible to force the application with LANG
> to another code page than the system locale, except with `LANG=.utf8`, with character
> data transformation option disabled in the ODBC source.

## Using the ESM driver

The database client application codeset needs to be defined when using the ESM driver.

The Easysoft client character set is specified with the "`Client_CSet`" parameter,
and the server character set is defined by "`Server_CSet`" or
"`Server_UCSet`" parameters. For example, to cover all UNICODE characters,
define:

```
Client_CSet   = UTF-8
Server_UCSet  = UTF-16LE
```

To support all UNICODE characters when using UTF-8 with `NCHAR/NVARCHAR` columns,
you need to define `Client_CSet=UTF-8` and
`Server_UCSet=UTF-16LE`.

According to the char or widechar mode, the ESM driver will do the
appropriate ODBC bindings, using `SQL_CHAR/SQL_VARCHAR` or
`SQL_WCHAR/SQL_WVARCHAR` as SQL type and `SQL_C_CHAR` as C type. The
Easysoft ODBC driver will then do the appropriate character set conversions for
`SQL_C_CHAR` data.

The char / widechar modes can be forced with the following FGLPROFILE
entry:

```
dbi.database.dsname.esm.widechar= { true | false }
```

## Using the FTM driver

The database client
application codeset needs to be defined when using the FTM driver.

The FreeTDS client
character set is defined with "`ClientCharset`" parameter in
odbc.ini:

```
ClientCharset = UTF-8
```

According to the
char or widechar mode, the FTM driver will do the appropriate ODBC
bindings, using `SQL_CHAR/SQL_VARCHAR` or `SQL_WCHAR/SQL_WVARCHAR` as
SQL type and `SQL_C_CHAR` as C type. The FreeTDS ODBC driver will then do the
appropriate character set conversions for `SQL_C_CHAR` data.

The char
/ widechar modes can be forced with the following FGLPROFILE
entry:

```
dbi.database.dsname.ftm.widechar= { true | false }
```

## Related links

**Related concepts**  

[Temporary tables](1291-temporary-tables.md "Temporary tables")

[CHAR and VARCHAR types](1012-char-and-varchar-types.md "Using the CHAR and VARCHAR data types with different databases.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1260-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
