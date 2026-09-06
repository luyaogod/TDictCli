---
title: "IBM Informix emulation parameters in FGLPROFILE"
source: "fgl-topics/c_fgl_Connections_015.html"
breadcrumb: "SQL support > Database connections > IBM® Informix® emulation parameters in FGLPROFILE"
type: "concept"
description: "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries."
---

# IBM Informix emulation parameters in FGLPROFILE

> Emulation of Informix® specific SQL features can be controlled with FGLPROFILE entries.

## What are Informix SQL emulation settings used for?

To simplify the migration process to other database servers such as IBM®
Informix, the database drivers can emulate some IBM Informix-specific features like SERIAL columns and temporary
tables; the drivers can also do some SQL syntax translation.

Avoid using IBM Informix
emulations; write portable SQL code instead.
IBM Informix emulations
are only provided to help you in the migration process. Disabling IBM
Informix emulations improves performance, because SQL
statements do not have to be parsed to search for IBM
Informix-specific syntax.

Emulations can be controlled with [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
parameters. You can disable all possible switches one-by-one, in order to test your programs for SQL
compatibility.

## `dbi.database.dsname.ifxemul`

This is a global switch to enable or disable IBM
Informix emulations.

Values can be `true` or `false`. Default is
`true`.

```
dbi.database.stores.ifxemul = false
```

## `dbi.database.dsname.ifxemul.datatype.typename`

The '`ifxemul.datatype`' switches define whether the specified data type
must be converted to a native type (for example, when creating a table with the CREATE
TABLE statement).

The typename can be one of `char, varchar, date, datetime, decimal, money,
float, real, integer, interval, smallint, serial, text, byte, bigint, bigserial, int8, serial8,
boolean`.

Default is `true` for all types.

```
dbi.database.stores.ifxemul.datatype.serial = false
```

## `dbi.database.dsname.ifxemul.datatype.serial.emulation`

This parameter can be used to control the serial generation technique used by the
driver to generate auto-incremented values.

The value can be one of the following:

- `native` uses the database's native sequence generator directly in the table
  definitions (depends on the db type).
- `native2` uses a secondary native sequence generator directly in the table
  definitions (depends on the db type).
- `regtable` uses the SERIALREG table with triggers. It is slower than the
  `native` emulation.
- `trigseq`", uses database sequence generator with triggers (not supported by all
  drivers).

Default is "`native`".

```
dbi.database.stores.ifxemul.datatype.serial.emulation = "native"
```

Serial emulations depend on the type of database server used. See [SQL database guides](1178-sql-database-guides.md "This section includes the SQL guides for various supported database servers.") for more details.

## `dbi.database.dsname.ifxemul.datatype.serial.sqlerrd2`

Use this parameter to disable the automatic serial retrieval done by the drivers, to fill the
`sqlca.sqlerrd[2]` register. When this feature is enabled, depending on the type of
serial emulation, drivers need to execute an additional SQL query to fetch the last generated
serial. If you have implemented your own function to fetch the last serial and you cannot disable
the whole serial emulation, set this parameter to false.

Default is `true`.

```
dbi.database.stores.ifxemul.datatype.serial.sqlerrd2 = false
```

See [Disabling automatic serial retrieval for sqlca.sqlerrd[2]](1377-serial-and-bigserial-data-types.md) for more
details.

## `dbi.database.dsname.ifxemul.temptables`

This switch can be used to control temporary table emulation.

Defaults is `true`.

```
dbi.database.stores.ifxemul.temptables = false
```

## `dbi.database.dsname.ifxemul.temptables.emulation`

This parameter can be used to specify what technique must be used to emulate temporary tables
in the database server.

Possible values are "`default`", "`private`" and
"`global`".

```
dbi.database.stores.ifxemul.temptables.emulation = "global"
```

See [SQL database guides](1178-sql-database-guides.md "This section includes the SQL guides for various supported database servers.") for more details.

## `dbi.database.dsname.ifxemul.dblquotes`

This switch can be used to define whether double quoted strings must be converted to single
quoted strings.

Default is `true`.

```
dbi.database.stores.ifxemul.dblquotes = false
```

If this emulation is enabled, all double quoted strings are converted, including database
object names.

## `dbi.database.dsname.ifxemul.outers`

This switch can be used to control IBM
Informix `OUTER` translation to
native SQL outer join syntax.

Default is `true`.

```
dbi.database.stores.ifxemul.outers = false
```

Consider using standard ISO outer joins in your SQL statements (LEFT OUTER).

## `dbi.database.dsname.ifxemul.today`

This switch can be used to convert the `TODAY` keyword to a native
expression returning the current date.

Default is `true`.

```
dbi.database.stores.ifxemul.today = false
```

## `dbi.database.dsname.ifxemul.current`

This switch can be used to convert the `CURRENT X TO Y` expressions to
a native expression returning the current time.

Default is `true`.

```
dbi.database.stores.ifxemul.current = false
```

## `dbi.database.dsname.ifxemul.selectunique`

This switch can be used to convert the `SELECT UNIQUE` to
`SELECT DISTINCT`.

Default is `true`.

```
dbi.database.stores.ifxemul.selectunique = false
```

> **Note:**
>
> Consider replacing all `UNIQUE` keywords by `DISTINCT`.

## `dbi.database.dsname.ifxemul.colsubs`

This switch can be used to control column substrings expressions (col[x,y]) to native
substring expressions.

Default is `true`.

```
dbi.database.stores.ifxemul.colsubs = false
```

> **Note:**
>
> Consider using substring SQL functions instead of [x,y] expressions in SQL.

## `dbi.database.dsname.ifxemul.matches`

This switch can be used to define whether `MATCHES` expressions must
be converted to LIKE expressions.

Default is `true`.

```
dbi.database.stores.ifxemul.matches = false
```

> **Note:**
>
> Consider using LIKE expressions instead of MATCHES in SQL.

## `dbi.database.dsname.ifxemul.length`

This switch can be used to define whether `LENGTH()` function names
have to be converted to the native equivalent.

Default is `true`.

```
dbi.database.stores.ifxemul.length = true
```

## `dbi.database.dsname.ifxemul.rowid`

This switch can be used to define whether `ROWID` keywords have to be converted to
native equivalent.

Default is `true`.

```
dbi.database.stores.ifxemul.rowid = false
```

> **Note:**
>
> Consider using primary keys instead of ROWIDs.

## `dbi.database.dsname.ifxemul.listupdate`

This switch can be used to convert the `UPDATE` statements using
non-ANSI syntax.

Default is `true`.

```
dbi.database.stores.ifxemul.listupdate = false
```

## `dbi.database.dsname.ifxemul.extend`

This switch can be used to convert simple `EXTEND()` expressions
to native date/time expressions.

Default is `true`.

```
dbi.database.stores.ifxemul.extend = true
```

## `dbi.database.dsname.ifxemul.rowlimiting`

This switch can be used to convert Informix SQL row limiting clause `SELECT [SKIP n] FIRST
m` to the native SQL equivalent.
> **Note:**
>
> - Row limiting clauses using SQL parameters will not be converted: The `SKIP` and
>   `FIRST` keywords must be followed by an integer constaint.
> - When using nested SQL queries, only the row limiting clause of the main `SELECT`
>   will be converted. The row limiting clauses in subqueries will not be converted.

Default is
`true`.

```
dbi.database.stores.ifxemul.rowlimiting = true
```

## `dbi.database.dsname.ifxemul.nationalchars`

When creating tables in programs, with some database brands, this switch can be used to convert
`CHAR(n)`, `VARCHAR(n)`, `LVARCHAR(n)` and
`TEXT` data type names to the corresponding native national character types
(`NCHAR(n)`, `NVARCHAR(n)`, etc).

Default is
`false`.

```
dbi.database.stores.ifxemul.nationalchars = true
```

> **Note:**
>
> The `dbi.database.dsname.ifxemul.nationalchars` switch is
> ignored, when the data type emulation parameter corresponding to the type (such as
> `dbi.database.dsname.datatype.varchar`) is set to
> `false`.

The `dbi.database.dsname.ifxemul.nationalchars` parameter can
be used with the following database drivers:

- SQL Server ODI drivers, when:
  - The current locale is multibyte (typically UTF-8)
  - The widechar mode is enabledFor more details, see [`CHAR/VARCHAR` types
  with SQL Server](1273-char-and-varchar-data-types.md) and [`TEXT` type with SQL
  Server](1282-text-and-byte-lob-types.md).

## Related links

**Related concepts**  

[SQL programming](0984-sql-programming.md "Covers topics about interacting with a database server using SQL.")
