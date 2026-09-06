---
title: "TEXT and BYTE (LOB) types"
source: "fgl-topics/c_fgl_odiagmsv_032.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > TEXT and BYTE (LOB) types"
type: "concept"
description: "Informix® Informix provides the TEXT , BYTE , CLOB and BLOB data types to store very large texts or binary data. Legacy Informix 4GL applications typically use the TEXT and BYTE types. Genero BDL does ..."
---

# TEXT and BYTE (LOB) types

## Informix®

Informix provides the `TEXT`,
`BYTE`, `CLOB` and `BLOB` data types to store very
large texts or binary data.

Legacy Informix 4GL applications typically use the
`TEXT` and `BYTE` types.

Genero BDL does not support the Informix
`CLOB` and `BLOB` types.

## Microsoft™ SQL Server

Microsoft SQL Server provides the
`VARCHAR(MAX)`, `NVARCHAR(MAX)` and `VARBINARY(MAX)`
data types to store large object data.

The SQL Server `TEXT`, `NTEXT` and `IMAGE` data
types still exist, but are considered as obsolete and will be removed in a future version.

In SQL Server, the `VARCHAR(MAX)`, `NVARCHAR(MAX)` and
`VARBINARY(MAX)` types have a limit of 2 gigabytes (2^31 -1 actually).

## Solution

SQL Server database drivers make the appropriate ODBC bindings for `TEXT` and
`BYTE` program variables, as SQL parameters and fetch buffers, to be used with SQL
Server `VARCHAR(MAX)`, `NVARCHAR(MAX)` and
`VARBINARY(MAX)` columns.

With Informix emulations enabled, in DDL statements such as `CREATE TABLE`:

- The `BYTE` data type name is converted to `VARBINARY(MAX)`.
- The `TEXT` data type name is converted to `VARCHAR(MAX)` or
  `NVARCHAR(MAX)`, depending on the applications locale, widechar mode
  usage and `dbi.database.dnname.ifxemul.nationalchars` FGLPROFILE
  setting.

To store `TEXT` data, use either `VARCHAR(MAX)` or
`NVARCHAR(MAX)`, by following the same rules as for `CHAR/VARCHAR`
data, based on the widemode mode. For more details, see [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md).
> **Note:**
>
> Starting with SQL Server 2019, the database collation can be defined with the `_UTF8`
> modifier, to store UTF-8 encoded data in `CHAR(n)`,
> `VARCHAR(n)` and `VARCHAR(MAX)` columns. In this
> case, UTF-8 data hold by `TEXT` variables can be stored in `VARCHAR(MAX)`
> columns instead of `NVARCHAR(MAX)` columns.

By default, in DDL statements executed by programs (such as `CREATE TABLE`), the
`TEXT` type name is left untouched. To force the ODI drivers to replace
`TEXT` by the `NVARCHAR(MAX)` type name (when the
widechar mode is enabled), set the [`dbi.database.dsname.ifxemul.nationalchars`](1079-ibm-informix-emulation-parameters-in-fglprofile.md) FGLPROFILE
entry to `true` (default is
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

Genero `TEXT/BYTE` program variables and the SQL Server large object types have
the same a limit of 2 gigabytes.

> **Note:**
>
> When using a stored procedure that has `SET/IF`
> statements and produces a result set with LOBs, the LOB columns must appear at the end of the
> `SELECT` list. If LOB columns are followed by other columns with regular types, the
> fetching rows will fail. Using `SET NOCOUNT ON` in the stored procedure does not
> help, because the cursor type is changed from a server cursor to a default result set cursor.

The `TEXT` and
`BYTE` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.text = { true | false }
dbi.database.dsname.ifxemul.datatype.byte = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1260-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
