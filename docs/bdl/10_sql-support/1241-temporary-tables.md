---
title: "Temporary tables"
source: "fgl-topics/c_fgl_odiagdmg_021.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data manipulation > Temporary tables"
type: "concept"
description: "Informix® Informix temporary tables are created with the CREATE TEMP TABLE DDL instruction or with SELECT ... INTO TEMP statement: CREATE TEMP TABLE tt1 ( pkey INT, name VARCHAR(50) ) CREATE TEMP ..."
---

# Temporary tables

## Informix®

Informix temporary tables are created with the
`CREATE TEMP TABLE` DDL instruction or with `SELECT ... INTO TEMP`
statement:

```
CREATE TEMP TABLE tt1 ( pkey INT, name VARCHAR(50) )
CREATE TEMP TABLE tt2 ( pkey INT, name VARCHAR(50) ) WITH NO LOG 
SELECT * FROM tab1 WHERE pkey > 100 INTO TEMP tt2
```

Temporary tables are automatically dropped when the SQL session ends, but they can also be
dropped with the `DROP TABLE` command. There is no name conflict when several users
create temporary tables with the same name.

BDL reports can create a temporary table when the rows are not sorted externally (by the source
SQL statement).

Informix allows you to create indexes on temporary
tables. No name conflict occurs when several users create an index on a temporary table by using the
same index identifier.

When creating temporary tables in Informix, the
`WITH NO LOG` clause can be used to avoid the overhead of recording DML operations in
transaction logs.

## Dameng®

Dameng supports the `DECLARE GLOBAL TEMPORARY
TABLE` instruction.

Dameng does not provide an equivalent of the Informix CREATE TEMP TABLE instruction, that creates
a temporary table dedicated to the current SQL session.

For more details, see the Dameng documentation.

## Solution

To emulate Informix temp table instructions `CREATE TEMP TABLE` and `SELECT
… INTO TEMP`, the Dameng ODI driver will use regular permanent tables, by converting the
table name to a unique name. These tables are created in the database user schema. The connected
user must have sufficient privileges to create permanent tables in its schema.

The general FGLPROFILE entry to
control temporary table emulation
is:

```
dbi.database.dsname.ifxemul.temptables = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## How does it work ?

- Informix-specific statements involving temporary table creation are automatically converted to
  Dameng `CREATE TABLE` statements in the "temptabs" tablespace, and the table name is
  converted to a unique name in the
  form:

  ```
  user-schema.ttunique-id_original-table-name
  ```

  For
  example, in the user schema `"USER1"`, a `CREATE TEMP TABLE tab1` will
  be converted
  to:

  ```
  CREATE TABLE USER1.tt140166177805368_652332_tab1 ...  TABLESPACE temptabs
  ```
- The unique-id part for the real table name is by default generated with a
  SELECT statement that builds a string from the current session id and the fraction of seconds part
  of the current systimestamp. If needed, you can define a different select statement with the
  following FGLPROFILE
  entry:

  ```
  dbi.database.dsname.dmg.sid.command = "SELECT ... FROM dual"
  ```
- Once the temporary table has been created, all other SQL statements performed in the current SQL
  session are parsed to replace the original table name by the unique table name.

## Prerequisites

A dedicated regular tablespace needs to be dedicated for temporary table emulation. The default
name for this tablespace is `"temptabs"`, and can be changed by the following
FGLPROFILE
setting:

```
dbi.database.dbname.dmg.temptables.tablespace = "tsname"
```

where
dbname is the database resource name used by the program, and
tsname is the dedicated tablespace.

To create a tablespace in Dameng, use the following SQL
commands:

```
CREATE TABLESPACE temptabs
  DATAFILE '/home/dmdba/dmdbms/data/TEST1/TEMPTABS1.DBF'
      SIZE 50 AUTOEXTEND ON;
```

## Related links

**Related concepts**  

[Temporary tables](1039-temporary-tables.md "Syntax for temporary table creation is not unique across all database engines.")
