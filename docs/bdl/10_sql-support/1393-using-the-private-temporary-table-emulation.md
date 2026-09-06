---
title: "Using the private temporary table emulation"
source: "fgl-topics/c_fgl_odiagora_ptt.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Temporary tables > Using the private temporary table emulation"
type: "concept"
description: "The private temporary table emulation feature can be used to benefit from the native Oracle® 18c private temporary tables. Important: Oracle private temporary tables have several limitations compared ..."
---

# Using the private temporary table emulation

The private temporary table emulation feature can be used to benefit from the native Oracle® 18c private temporary tables.
> **Important:**
>
> Oracle private temporary
> tables have several limitations compared to Informix® temporary
> tables. Using this feature requires code review to overcome these limitations.

In order to use private temporary table emulation, define the following FGLPROFILE
entry:

```
dbi.database.dbname.ifxemul.temptables.emulation = "private"
```

## How does the private temp table emulation work?

- Informix `CREATE TEMP TABLE` and
  `SELECT INTO TEMP` statements are automatically converted to Oracle "`CREATE PRIVATE TEMPORARY TABLE`".
- The original name of the temporary table must be converted to an Oracle private temporary table
  name (PTT) with a specific prefix, which is defined by the
  `private_temp_table_prefix` system parameter (default is
  `ORA$PTT_`).

  The detection of the PTT prefix can be controlled with FGLPROFILE entries, to get the value of
  `private_temp_table_prefix`. For more details, see [Determine the private temporary table prefix](1393-using-the-private-temporary-table-emulation.md)
- Since Oracle private temporary tables are only visible to the current user and SQL session, it is
  not possible to specify a schema prefix.
- To mimic the behavior of Informix temporary tables, Oracle private temporary tables are created
  with the `ON COMMIT PRESERVE DEFINITION` clause. The temp table will then remain when
  a transaction ends.
- By default, private temporary tables are created in the default temporary table space. If needed,
  you can specify another tablespace name for Oracle private temporary tables with the following
  FGLPROFILE
  entry:

  ```
  dbi.database.dsname.ora.temptables.tablespace = "mytemptabs"
  ```

  > **Note:**
  >
  > This tablespace must be a temporary tablespace.
- Once the temporary table has been created, all other SQL statements performed in the current SQL
  session are parsed to convert the original table name to the corresponding private temporary table
  name.
- When the BDL program disconnects from the database, Oracle will automatically drop private
  temporary tables created during the SQL session.

## Prerequisites when using the private temp table emulation

- By default (with Oracle 18.3), the maximum number of private temporary tables allowed per user
  is `16`. Consider increasing this number, if your programs create more temporary
  tables. Otherwise, Oracle will produce the error "`ORA-32460: maximum number of private
  temporary tables per session exceeded`". The max number of private temp tables can be
  changed with the `_ptt_max_num` hidden system
  parameter:

  ```
  ALTER SYSTEM SET "_ptt_max_num"=32 SCOPE = BOTH;
  ```
- For more details, check also "`CREATE PRIVATE TEMPORARY TABLE`" in the Oracle
  documentation.

## Limitations of the private temp table emulation

- By default (in Oracle 18.3), the max number of private temp tables per user is 16. This can
  be changed by a system parameter.
- It is not possible to define table constraints such as `NOT NULL` or
  `PRIMARY KEY`, etc.
- Index creation is not allowed on private temp tables. Therefore, reports with internal
  `ORDER BY` cannot work.
- Columns can't have `DEFAULT` values. Thus, the `native2` serial
  emulation cannot work.
- Triggers are not supported. Therefore, `native` and `regtable`
  serial emulations cannot work with private temp tables.
- There is no such concept as Informix `CREATE TEMP TABLE ... WITH NO LOG`. By
  default, Oracle private temporary table data is always logged for transaction rollback.
- Tokens matching the original table names are converted to unique names in all SQL statements.
  Make sure you are not using the temp table name for other database objects, like columns. The
  following example illustrates this limitation:

  ```
  CREATE TABLE tab1 ( key INTEGER, tmp1 CHAR(20) )
  CREATE TEMP TABLE tmp1 ( col1 INTEGER, col2 CHAR(20) )
  SELECT tmp1 FROM tab1 WHERE ...
  ```

## Serial emulation with private temporary tables

> **Important:**
>
> Because Oracle private temporary tables do not support default value
> specifications (used for `native2` serial emulation) and triggers (used for
> `native` and `regtable` serial emulations), it is not possible to
> emulate Informix serial types with Oracle private temporary tables.

## Determine the private temporary table prefix

The
default prefix for Oracle private temporary tables in `"ORA$PTT_"`. This prefix can
be changed by database administrators.

To get the current PTT prefix, the Oracle ODI driver
can read the `private_temp_table_prefix` system parameter from the
`v$parameter` system
view:

```
SELECT value FROM v$parameter WHERE name = 'private_temp_table_prefix';
```

However, regular DB users do not have the permission to read this system view.

Therefore, the
detection of the PTT prefix can be controlled with FGLPROFILE entries, to avoid the need to access the
`v$parameter` system view.

The next FGLPROFILE entry defines the way to get the
prefix of PTTs:

```
dbi.database.mydb.ora.temptables.private.prefix.source = { "value" | "command" }
```

When the source is defined as `"value"`, the next FGLPROFILE entry defines the value of
the PTT prefix:

```
dbi.database.mydb.ora.temptables.private.prefix.value = "myprefix"
```

When the source is defined as `"value"`, but no
`dbi.database.dsname.ora.private.prefix.value` entry is defined, the ODI
driver uses `"ORA$PTT_"`.

When the source is defined as
`"command"`, the next FGLPROFILE entry defines the SQL command to be executed to get
the PTT
prefix:

```
dbi.database.mydb.ora.temptables.private.prefix.command = "select 'myprefix' from dual"
```

When
the source is defined as `"command"`, but no
`dbi.database.dsname.ora.private.prefix.command` entry is defined, the ODI
driver will try to read from `v$parameter` with:

```
select value from v$parameter where name = 'private_temp_table_prefix'
```

> **Note:**
>
> Read privileges on the `v$parameter` system view is not allowed for regular DB
> users by default (`GRANT SELECT` permission must be performed for all DB
> users)

Otherwise, when none of the above entries are defined, the Oracle ODI driver uses
`"ORA$PTT_"` by default. If the defaut PTT prefix was not changed in the database,
regular DB users do not need `GRANT SELECT` permission on
`v$parameter`.

## Related links

**Related concepts**  

[SERIAL and BIGSERIAL data types](1377-serial-and-bigserial-data-types.md "SERIAL and BIGSERIAL data types")
