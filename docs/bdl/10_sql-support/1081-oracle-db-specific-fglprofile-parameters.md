---
title: "Oracle DB specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_ora.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > Oracle DB specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .ora.schema Name of the database schema to be selected after connection is established. dbi.database.stores.ora.schema = \"store2\" Set this parameter to a specific schema in order ..."
---

# Oracle DB specific FGLPROFILE parameters

## `dbi.database.dsname.ora.schema`

Name of the database schema to be selected after connection is established.

```
dbi.database.stores.ora.schema = "store2"
```

Set this parameter to a specific schema in order to share the same table with
all users.

For more details, see [Database users](1368-database-users.md).

## `dbi.database.dsname.ora.prefetch.rows`

Maximum number of rows to prefetch.

```
dbi.database.stores.ora.prefetch.rows = 50
```

Use this parameter to increase performance by defining the maximum number of
rows to be fetched into the db client buffer. However, the bigger this parameter
is, the more memory that is used by each program. This parameter applies to all
cursors in the program.

The default is 10 rows.

## `dbi.database.dsname.ora.prefetch.memory`

Maximum buffer size for prefetching (in bytes).

```
dbi.database.stores.ora.prefetch.memory = 4096
```

This parameter is equivalent to `prefetch.rows`, but here you
can specify the memory size instead of the number of rows.
Like `prefetch.rows`, this parameter applies to all cursors in
the program.

The default is 0, which means that memory size is not included in computing
the number of rows to prefetch.

## `dbi.database.dsname.ora.sid.command`

SQL command (SELECT) to generate a unique session id (used for temp table
names).

```
dbi.database.stores.ora.sid.command =
 "SELECT TO_CHAR(SID)||'_'||TO_CHAR(SERIAL#)
      FROM V$SESSION WHERE AUDSID=USERENV('SESSIONID')"
```

By default, the driver uses "`SELECT USERENV('SESSIONID') FROM DUAL`".
This is the standard session identifier in Oracle®, but it can become a very large number and cannot
be reset.

This parameter gives you the freedom to provide your own way to generate a
session id.

The SELECT statement must return a single row with one single column.

Value can be an integer or an identifier such as `"pre999"`.

## `dbi.database.dsname.ora.boolean.aschar`

With Oracle 23c and the `dbmora_23` ODI driver, forces to use the SQL
`CHAR(1)` type must be used to store FGL `BOOLEAN` data to ensure
backward compatibility:

```
dbi.database.stores.ora.boolean.aschar = true
```

By default, when using Oracle 23c + `dbmora_23`, the native SQL
`BOOLEAN` type is used for FGL `BOOLEAN`.

For more details, see [BOOLEAN data type](1372-boolean-data-type.md).

## `dbi.database.dsname.ora.date.ifxfetch`

Controls the way an Oracle
DATE is fetched into program variables, especially CHAR/VARCHAR targets.

```
dbi.database.stores.ora.date.ifxfetch = true
```

By default, since Oracle
DATE type is equivalent to `DATETIME YEAR TO SECOND`, values are fetched
into `CHAR`/`VARCHAR` with time information and are
formatted with the style `YYYY-MM-DD hh:mm:ss`.
If you need to get the IBM®
Informix® behavior, to fetch DATEs only
with the YMD part following the DBDATE environment variable, set this parameter to
true. However, this parameter is useless when fetching Oracle DATEs into `DATE` or
`DATETIME` variables, which is the recommended way to hold date
and time values in programs.

Default is `false` (with
time information, using normalized format).

## `dbi.database.dsname.ora.temptables.schema.source`

Defines the source of the schema name used for [temporary table emulation](1390-temporary-tables.md).

Possible values are:

- `"login"` : The user name specified in the connection parameters.
- `"command"` : The value returned by a SELECT statement specified
  in the `ora.temptables.schema.command` parameter.

```
dbi.database.stores.ora.temptables.schema.source = "login"
```

By default, when using the default temporary table emulation, the driver uses no
schema at all. When using global temporary table emulation, the driver uses the
TEMPTABS schema by default.

When specifying a `"command"` source for the schema, you can provide
your own SELECT statement to produce the schema name to be used when creating a
table.

## `dbi.database.dsname.ora.temptables.schema.command`

SQL command (SELECT) to get the schema name to be used for [temporary table emulation](1390-temporary-tables.md).

```
dbi.database.stores.ora.temptables.schema.source = "command"
dbi.database.stores.ora.temptables.schema.command = "SELECT SYS_CONTEXT('USERENV','SESSION_USER') FROM DUAL"
```

This configuration parameter is only taken into account if the
`ora.temptables.schema.source` parameter is defined as
`"command"`.

The SELECT statement must return a single row with one single column.

Value must be a character string.

## `dbi.database.dsname.ora.temptables.tablespace`

Defines the tablespace to be used for temporary table emulation.

```
dbi.database.stores.ora.temptables.tablespace = "mytemptabs"
```

By default:

1. When using the [default temporary table emulation](1391-using-the-default-temporary-table-emulation.md),
   the driver uses the TEMPTABS tablespace. The tablespace specified in the
   `ora.temptables.tablespace` entry must be a permanent tablespace.
2. When using [global temporary table emulation](1392-using-the-global-temporary-table-emulation.md), the
   driver uses no tablespace by default. The tablespace specified in the
   `ora.temptables.tablespace` entry must be a temporary tablespace.
3. When using [private temporary table emulation](1393-using-the-private-temporary-table-emulation.md), the
   driver uses no tablespace by default. The tablespace specified in the
   `ora.temptables.tablespace` entry must be a temporary tablespace.

For more details, see [Temporary tables](1390-temporary-tables.md).

## `dbi.database.dsname.ora.temptables.private.prefix.source`

Defines the method to find out the prefix for Oracle private temporary tables.

Possible values are:

- `"value"` : The PTT is specified with the
  `dbi.database.dsname.ora.temptables.private.prefix.value` parameter.
- `"command"` : The PTT is found with the SELECT statement specified in the
  `dbi.database.dsname.ora.temptables.private.prefix.command` parameter.

For more details, see [Using the private temporary table emulation](1393-using-the-private-temporary-table-emulation.md).

## `dbi.database.dsname.ora.temptables.private.prefix.value`

Defines the prefix for Oracle private temporary tables when
`dbi.database.dsname.ora.temptables.private.prefix.source = "value"`.

For example:

```
dbi.database.stores.ora.temptables.private.prefix.source = "value"
dbi.database.stores.ora.temptables.private.prefix.value = "myprefix"
```

For more details, see [Using the private temporary table emulation](1393-using-the-private-temporary-table-emulation.md).

## `dbi.database.dsname.ora.temptables.private.prefix.source`

Defines the SELECT statement to find theprefix for Oracle private temporary tables when
`dbi.database.dsname.ora.temptables.private.prefix.source = "command"`.

For example:

```
dbi.database.stores.ora.temptables.private.prefix.source = "command"
dbi.database.stores.ora.temptables.private.prefix.command = "select 'myprefix' from dual"
```

For more details, see [Using the private temporary table emulation](1393-using-the-private-temporary-table-emulation.md).
