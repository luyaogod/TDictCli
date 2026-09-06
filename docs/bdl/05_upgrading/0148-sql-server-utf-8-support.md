---
title: "SQL Server UTF-8 support"
source: "fgl-topics/c_fgl_Migrate_to_400_sqlserver_utf8.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > SQL Server UTF-8 support"
type: "concept"
---

# SQL Server UTF-8 support

> Support for UTF-8 collation in CHAR/VARCHAR columns with SQL Server 2019.

Microsoft™ SQL Server 2019 introduced support for UTF-8
database collations: When using the `_UTF8` modifier in the database collation name,
the CHAR and VARCHAR columns will store UTF-8 encoded character strings.
> **Note:**
>
> National character
> types NCHAR/NVARCHAR use the UCS-2 or UTF-16 encoding, and are supported by Genero as well, to store
> character string data from UTF-8 applications.

With UTF-8 support in CHAR/VARCHAR column types, SQL Server 2019 offers the choice to store
UNICODE data in CHAR/VARCHAR (with length unit in bytes) or NCHAR/NVARCHAR columns (with length unit
in double-bytes).

Starting with Genero version 4.00, SQL Server ODI drivers are able to deal with SQL Server
CHAR/VARCHAR columns using UTF-8 enabled collation.

> **Note:**
>
> This applies also to BDL variables defined as `TEXT`, with values stored in
> `VARCHAR(MAX)` or `NVARCHAR(MAX)` columns.

For all SQL Server ODI drivers, the
`dbi.database.dsname.???.widechar` FGLPROFILE parameter defines
how character strings are converted and bound with the ODBC API.

The application locale (LANG/LC\_ALL), the length semantincs (FGL\_LENGTH\_SEMANTICS) and the SQL
Server database collation, define the default setting for the widechar mode.

Options for UTF-8 Genero applications:

- When application locale is UTF-8 and has BYTE length semantics, use UTF-8 in CHAR/VARCHAR
  columns, and widechar mode must be set to `false`.
- When application locale is UTF-8 and has CHAR length semantics, use UTF-16 in NCHAR/NVARCHAR
  columns, and widechar mode must be set to `true`.

In all cases, when the FGLPROFILE entry is not defined, the ODI drivers for SQL Server select the
appropriate widechar mode by default.

> **Important:**
>
> Minimum ODBC driver versions required to use UTF-8 collation for CHAR/VARCHAR
> columns are:
>
> - MS ODBC for SQL Server: 17.6.1.1
> - Easysoft ODBC for SQL Server: 2.0.19
> - FreeTDS: 1.3.x (not yet GA)

For more details, see [CHAR and VARCHAR data types](../10_sql-support/1273-char-and-varchar-data-types.md) and [TEXT and BYTE (LOB) types](../10_sql-support/1282-text-and-byte-lob-types.md).
