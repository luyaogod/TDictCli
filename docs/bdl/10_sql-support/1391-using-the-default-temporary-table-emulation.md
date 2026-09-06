---
title: "Using the default temporary table emulation"
source: "fgl-topics/c_fgl_odiagora_025.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Temporary tables > Using the default temporary table emulation"
type: "concept"
description: "The default temporary table emulation is provided to achieve a high level of Informix® compatibility, when minimal code change is required. In order to use the default temporary table emulation, ..."
---

# Using the default temporary table emulation

The default temporary table emulation is provided to achieve a high level of Informix® compatibility, when minimal code change
is required.

In order to use the default temporary table emulation, define the following FGLPROFILE
entry (or leave it
unset):

```
dbi.database.dbname.ifxemul.temptables.emulation = "default"
```

## How does the default temp table emulation work?

- Informix `CREATE TEMP TABLE` and
  `SELECT INTO TEMP` statements are automatically converted to Oracle® "`CREATE TABLE`". The name of the
  temporary table is converted to a unique table name.
- By default, Oracle tables are created in the current schema (i.e. no schema prefix is added to
  the generated table name). In order to control the schema where the Oracle tables are created, use
  the following FGLPROFILE
  parameters:

  ```
  dbi.database.dsname.ora.temptables.schema.source = { "login" | "command" }
  dbi.database.dsname.ora.temptables.schema.command = "SELECT ..."
  ```

  When
  specifying `"login"` in the `ora.temptables.schema.source` parameter,
  the user name passed to the connection instruction will be used as schema name. When specifying
  `"command"` as schema source, the schema name will be produced by the
  `SELECT` statement defined in the `ora.temptables.schema.command`
  parameter.

  Schema specification for temporary table emulations is mandatory when you have
  created a DB user for each end user, and you force common schema usage with the
  `ora.schema`
  parameter:

  ```
  dbi.database.dsname.ora.schema = "app_owner"
  ```

  In
  this case, it is better to add the current user as schema name for temporary tables, to get Oracle
  tables created in the current user schema, instead of creating tables in the common schema: Creating
  tables in another schema requires `CREATE ANY TABLE / DROP ANY TABLE` privileges for
  each DB user, which is not possible in an organization using strong security policy.
- By default, temporary tables are created in a dedicated tablespace named
  "`TEMPTABS`". Of course the `TEMPTABS` tablespace must exist before
  running programs, otherwise temporary table creation will fail. Using a specific tablespace for
  temporary tables allows you to specify storage options, for example to use a physical device which
  can be different from the disk drive used for real data storage. Additionally, backups of permanent
  application tables can be performed without the data of temporary tables.

  If required, you can specify another tablespace name for Oracle tables with the following
  FGLPROFILE
  entry:

  ```
  dbi.database.dsname.ora.temptables.tablespace = "mytemptabs"
  ```

  > **Note:**
  >
  > This tablespace must be a permanent tablespace.
- With Oracle, dropped tables are saved in the recycle bin by default. You may want to avoid the
  recycle bin feature at the database level or session level with:

  ```
  ALTER SYSTEM SET recyclebin = OFF scope=spfile
  ```

  or, at current session level only:

  ```
  ALTER SESSION SET recyclebin = OFF
  ```
- Once the temporary table has been created, all other SQL statements
  performed in the current SQL session are parsed to convert the original
  table name to the corresponding unique table name.
- When the BDL program disconnects from the database (for example, when it ends or when a
  `CLOSE DATABASE` instruction is executed), the tables which have not been removed
  with an explicit `DROP TABLE` are automatically removed by the database interface.
  However, if the program crashes, the tables will remain in the database, so you may need to cleanup
  the database from time to time.

## Prerequisites when using the default temp table emulation

- Application users must have sufficient privileges to create database tables in the schema used
  to create the Oracle table (usually, `CONNECT` and `RESOURCE`
  roles).
- Create a tablespace dedicated to the Oracle tables created by this emulation. The default
  tablespace is named "`TEMPTABS`", but it can be another tablespace defined by the
  `ora.temptables.tablespace` FGLPROFILE entry as described before.

  The tablespace must be of type "permanent", as it will hold permanent tables
  used to emulate Informix temp
  tables.

  Make sure the tablespace is big enough to hold all the data, and check for
  automatic extension.

  When using a PDB, the tablespace must be created in the context of the PDB.

  ```
  CREATE TABLESPACE temptabs
    DATAFILE 'file-path' SIZE 1M AUTOEXTEND ON;
  -- Give privileges on temptabs tablespace to other users
  ALTER USER dbuser QUOTA UNLIMITED ON TEMPTABS;
  ```

  For more details, see `CREATE TABLESPACE` in the Oracle documentation.

## Limitations of the default temp table emulation

- When using the default emulation, the real name of an emulated temporary table will have the
  following format:

  ```
  ttnumber_name
  ```

  Where
  name if the original table name, and number is by default the
  Oracle AUDSID session id returned by:

  ```
  SELECT USERENV('SESSIONID') FROM DUAL
  ```

  Since session ids are persistent
  over server shutdown, make sure that the table names generated for temporary table emulation do not
  exceed the maximum length of Oracle table names.

  If needed, you can provide a customized SQL
  command to generate a unique session id, with the following FGLPROFILE
  entry:

  ```
  dbi.database.dbname.ora.sid.command = "select ..."
  ```

  As
  an example, you can use the SID column value from V$SESSION:

  ```
  SELECT SID FROM V$SESSION WHERE AUDSID = USERENV('SESSIONID')
  ```
- Application tables and columns cannot use the same table name as the name generated for
  temporary tables. Make sure you are not using table or column names with the
  format:

  ```
  ttnumber_name
  ```
- Tokens matching the original table names are converted to unique names in all
  SQL statements. Make sure you are not using the temp table name for other database
  objects, like columns. The following example illustrates this limitation:

  ```
  CREATE TABLE tab1 ( key INTEGER, tmp1 CHAR(20) )
  CREATE TEMP TABLE tmp1 ( col1 INTEGER, col2 CHAR(20) )
  SELECT tmp1 FROM tab1 WHERE ...
  ```

## Maintenance of default temp table emulation

- If you want to list the tables created by a specific user, do this:

  ```
  SELECT * FROM ALL_TABLES WHERE OWNER = 'user_name'
  ```

  As with other database object names, the user name is stored in uppercase letters if it has
  been created without using double quotes (`create user scott ...` = stored name is
  "SCOTT").

## Creating indexes on temporary tables with default temp table emulation

- Indexes created on temporary tables must have unique names also. The database interface detects
  `CREATE INDEX` statements which are using temporary tables and converts the index
  name to unique names. Indexes created on temporary tables follow the same storage and schema
  settings as for the `CREATE TABLE` statements: Schema and tablespace specification of
  FGLPROFILE will also be applied to `CREATE INDEX` commands.
- `DROP INDEX` statements are also detected to replace the original index name by
  the real name.

## SERIALs in temporary table creation with default temp table emulation

You can
use `SERIAL/BIGSERIAL` data types when creating a temporary table.

Sequences and
triggers will be created in the current schema.

See issue about [`serial` types](1377-serial-and-bigserial-data-types.md) for more
details.
