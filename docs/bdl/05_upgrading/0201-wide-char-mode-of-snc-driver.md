---
title: "Wide Char mode of SNC driver"
source: "fgl-topics/c_fgl_Migrate_to_310_snc_widechar.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Wide Char mode of SNC driver"
type: "concept"
---

# Wide Char mode of SNC driver

> The snc.widechar FGLPROFILE entry defaults to the right setting for the current application locale.

SQL Server performs better, when using the correct ODBC types in
`SQLBindParameter()`, depending on the database columns types:

- When using CHAR/VARCHAR/TEXT SQL types (typically, when the application locale uses a
  single-byte encoding), the recommendation is that ODBC SQL types SQL\_[VAR]CHAR (single-char ODBC SQL
  types) are used
- When using NCHAR/NVARCHAR/NTEXT SQL types (when the application locale uses a multibyte
  encoding such as UTF-8), it is recommended to use the ODBC SQL types SQL\_W[VAR]CHAR (these are
  wide-char ODBC SQL types)

The `dbi.database.dbname.snc.widechar` FGLPROFILE entry
controls the ODBC SQL types used by the SNC driver.

Before version 3.10, the `snc.widechar` option was set to `true` by
default, assuming that the application locale uses a multibyte encoding like UTF-8. If the database
columns are defined with CHAR/VARCHAR/TEXT types, this parameter had to be set to
`false`.

Starting with 3.10, to simplify application configurations when using a single-byte encoding, the
`snc.widechar` option defaults to the expected setting, based on the current
application locale (assuming that the database column types fit the application locale):

- If the application locale defines a single-byte encoding (such as ISO8859-1), we assume that the
  database columns are defined with CHAR/VARCHAR/TEXT types to store single-byte characters, and the
  SNC driver will use SQL\_[VAR]CHAR.
- If the application locale defines a multibyte encoding (such as UTF-8 or BIG5), we assume that
  the database columns are defined with NCHAR/NVARCHAR/NTEXT types to store UNICODE characters, and
  the SNC driver will use SQL\_W[VAR]CHAR.

> **Note:**
>
> Set the `dbi.database.dbname.snc.widechar` to
> `false`, only if you are using a multibyte encoding such as BIG5, with
> CHAR/VARCHAR/TEXT column types in the database.

> **Note:**
>
> Easysoft (ESM) and FreeTDS (FTM) drivers do not support the `snc.widechar`
> option: The SQL char type binding mode is automatic, depending on the current application locale,
> and it cannot be changed.

## Related links

**Related concepts**  

[CHAR and VARCHAR data types](../10_sql-support/1273-char-and-varchar-data-types.md "CHAR and VARCHAR data types")

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")

[Multibyte character sets (MBCS)](../09_advanced-features/0873-multibyte-character-sets-mbcs.md "Multibyte character sets (MBCS)")
