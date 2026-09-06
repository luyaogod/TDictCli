---
title: "Using the global temporary table emulation"
source: "fgl-topics/c_fgl_odiagora_026.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Temporary tables > Using the global temporary table emulation"
type: "concept"
description: "The global temporary table emulation is provided to get the benefit of Oracle's® global temporary tables, by sharing the same table structure with multiple SQL sessions, reducing the cost of the ..."
---

# Using the global temporary table emulation

The global temporary table emulation is provided to get the benefit of Oracle's® global temporary tables, by sharing the same table
structure with multiple SQL sessions, reducing the cost of the `CREATE TABLE`
statement execution. However, this emulation does not provide the same level of Informix® compatibility as the default emulation, and must be used
carefully.

In order to use the global temporary table emulation, define the following FGLPROFILE
entry:

```
dbi.database.dbname.ifxemul.temptables.emulation = "global"
```

## How does the global temp table emulation work?

- Informix `CREATE TEMP TABLE` and
  `SELECT INTO TEMP` statements are automatically converted to Oracle "`CREATE GLOBAL TEMPORARY TABLE`"
  statements. The original table name is kept, but it gets by default a "`TEMPTABS`"
  schema prefix, to share the underlying table structure with other database users.

  In order to
  control the schema where the Oracle global
  temporary tables are created, use the following FGLPROFILE parameters:

  ```
  dbi.database.dsname.ora.temptables.schema.source = { "login" | "command" }
  dbi.database.dsname.ora.temptables.schema.command = "SELECT ..."
  ```

  > **Note:**
  >
  > The purpose of global temporary tables is to create the table once in Oracle and have all users share the same table structure.
  > Therefore the default "`TEMPTABS`" schema is considered sufficient for this emulation
  > method.
- The Global Temporary Tables are created with the "`ON COMMIT PRESERVE ROWS`"
  option, to keep the rows in the table when a transaction ends.
- If the global temporary table to be created exists already, error ORA-00955 will be ignored by
  the database driver. This allows to do several `CREATE TEMP TABLE`
  statements in your programs with no SQL error, to emulate the Informix behavior. This works fine as long as the table name is unique for a given structure
  (column count and data types must match).
  > **Note:**
  >
  > Since `DROP TABLE` statements are
  > converted to `DELETE` statements, you might want to use Oracle sqlplus to issue a real `DROP TABLE` of
  > global temporary tables, to make a real cleanup from time to time.
- By default, global temporary table is created without any tablespace option, and thus will be
  created in the default tablespace assigned to the current user.

  If required, you can specify another tablespace name for Oracle tables with the following FGLPROFILE entry:

  ```
  dbi.database.dsname.ora.temptables.tablespace = "mytemptabs"
  ```

  > **Note:**
  >
  > This tablespace must be a temporary tablespace.
- Once the Global Temporary Table has been created, all other SQL statements performed in the
  current SQL session are parsed to convert the original table name to
  schema.original-tablename.
- When doing a `DROP TABLE temp-table` statement in the program,
  the database driver converts it to a `DELETE` statement, to remove all data added by
  the current session. A next `CREATE TEMP TABLE` or `SELECT INTO TEMP`
  will fail with error ORA-00955 but since this error is ignored, it will be transparent for the
  program. We cannot use `TRUNCATE TABLE` because that would require at least
  `DROP ANY TABLE` privileges for all users.
- When the BDL program disconnects from the database (for example, when it ends or when a
  `CLOSE DATABASE` instruction is executed), the tables that have not been dropped by
  the program with an explicit `DROP TABLE` statement will be automatically cleaned by
  Oracle.

## Prerequisites when using the global temp table emulation

- You must create a database user (schema) dedicated to this emulation, the default name is
  "`TEMPTABS`":

  ```
  CREATE USER temptabs IDENTIFIED BY pswd;
  ```
- If your programs need to create temporary tables on the fly with this method, you must grant
  `CREATE ANY TABLE + CREATE ANY INDEX` system privilege to all DB users. This is not a
  good practice for security reasons. You better "prepare" the database by creating the Global
  Temporary Table (when using the default schema, create it with the `TEMPTABS` user).
  Do not forget to specify `ON COMMIT PRESERVE ROWS` option. Then `GRANT
  INSERT`, `UPDATE`, `DELETE` and `SELECT` object
  privileges to `PUBLIC`, for example:

  ```
  CREATE GLOBAL TEMPORARY TABLE temptabs.mytable
    ( k INT PRIMARY KEY, c CHAR(10) ) ON COMMIT PRESERVE ROWS;
  CREATE UNIQUE INDEX temptabs.ix1 ON temptabs.mytable ( c );
  GRANT SELECT, UPDATE, INSERT, DELETE ON temptabs.mytable TO PUBLIC;
  ```

  For testing purpose, consider using a user with DBA privileges, to simplify the
  configuration.

## Limitations of the global temp table emulation

- Global Temporary Tables are shared by multiple users/sessions. In order to have the global
  emulation working properly, each temporary table name must be unique for a given
  table structure, for all programs. Avoid using temp tables names such as "tmp1".
  It is recommended to use table names as follows:

  ```
  CREATE TEMP TABLE custinfo_1 (
      cust_id INTEGER,
      cust_name VARCHAR(50)
   );
  CREATE TEMP TABLE custinfo_2 (
      cust_id INTEGER,
      cust_name VARCHAR(50),
      cust_addr VARCHAR(200)
   );
  ```
- Tokens matching the original table names are converted to unique
  names in all SQL statements. Make sure you are not using the
  temp table name for other database objects, like columns. The
  following example illustrates this limitation:

  ```
  CREATE TABLE tab1 ( key INTEGER, tmp1 CHAR(20) );
  CREATE TEMP TABLE tmp1 ( col1 INTEGER, col2 CHAR(20) );
  SELECT tmp1 FROM tab1 WHERE ...
  ```

## Creating indexes on temporary tables with global temp table emulation

- Indexes created on temporary tables get also the `TEMPTABS` schema prefix by
  default. Indexes created on temporary tables follow the same storage and schema settings as for the
  `CREATE GLOBAL TEMPORARY TABLE` statements: Schema and tablespace specification of
  FGLPROFILE will also be applied to `CREATE INDEX` commands.
- When executing a `DROP INDEX` statement on a temporary table in a program, the
  database driver just ignores the statement.

## SERIALs in temporary table creation with global temp table emulation

You can
use `SERIAL/BIGSERIAL` data types when creating a temporary table.

Sequences and
triggers will be created in the `TEMPTABS` schema too.

See issue about [`serial` types](1377-serial-and-bigserial-data-types.md) for more
details.
