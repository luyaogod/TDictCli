---
title: "Temporary tables"
source: "fgl-topics/c_fgl_odiagora_024.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Temporary tables"
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

## Oracle

Oracle® supports global temporary tables
that can be shared among several processes: Only data is temporary and local to an SQL process, the
table structure is common to all programs.

Oracle 18c provides private temporary tables, similar to Informix temporary tables: The table
structure and data is private to the program. However, private temporary tables in Oracle 18c have
several limitations that need to be considered before usage.

## Solution

In accordance with some prerequisites, temporary table creation in BDL programs can be supported
by the database interface.

> **Important:**
>
> When creating a temporary table, you perform a Data Definition Language
> statement. Oracle automatically commits
> the current transaction when executing a DDL statement. Therefore, you must avoid temp table
> creation/destruction in transactions.

The general FGLPROFILE entry to
control temporary table emulation
is:

```
dbi.database.dsname.ifxemul.temptables = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

> **Important:**
>
> Simple Informix-style SQL statement creating temporary tables can be
> converted to a native SQL equivalent instruction. However, complex SQL statements such as
> `SELECT .. INTO TEMP` with subqueries may fail. In such cases, create a view from
> the complex query and then create the temp table from the view. Or, disable Informix emulation
> and use the native SQL syntax to create the temporary table (`EXECUTE IMMEDIATE "/*
> fglhint_no_ifxemul */ …"`)
>
> With Informix SQL, if the source table has a column defined as `SERIAL` or
> `BIGSERIAL`, a `SELECT ... INTO TEMP` will produce a new temp table
> with an auto-incremented serial column. With the `SELECT … INTO TEMP` emulation for
> non-Informix databases, not using the native sequence generators (such as `IDENTITY`
> columns in SQL Server), the resulting temporary table will get a simple `INTEGER` or
> `BIGINT` column, instead of an auto-incremented column.

The temporary table emulation can use permanent tables, global temporary tables or (since Oracle
18c) private temporary tables.

Use the following FGLPROFILE entry, to define the type of temporary table emulation:

```
dbi.database.dbname.ifxemul.temptables.emulation = { "default"
                                                   | "private"
                                                   | "global" }
```

By default, the database driver uses regular permanent tables. The default emulation provides
maximum compatibility with Informix temporary tables, but
requires real table creation which can be a significant overhead with Oracle. For more details, see [Using the default temporary table emulation](1391-using-the-default-temporary-table-emulation.md).

The "private" temp table emulation uses native Oracle Private Temporary Tables, which are only visible for the current SQL
session. However, the private emulation mode has some limitations and constraints requiring code
review. For more details, see [Using the private temporary table emulation](1393-using-the-private-temporary-table-emulation.md).

The "global" temp table emulation uses native Oracle Global Temporary Tables, requiring only one initial table creation
and thus making programs run faster. However, the global emulation mode has some limitations
requiring code review. For more details, see [Using the global temporary table emulation](1392-using-the-global-temporary-table-emulation.md).

## Related links

**Related concepts**  

[Temporary tables](1039-temporary-tables.md "Syntax for temporary table creation is not unique across all database engines.")

[Oracle DB specific FGLPROFILE parameters](1081-oracle-db-specific-fglprofile-parameters.md "Oracle DB specific FGLPROFILE parameters")

## Child topics

- [Using the default temporary table emulation](1391-using-the-default-temporary-table-emulation.md)
- [Using the global temporary table emulation](1392-using-the-global-temporary-table-emulation.md)
- [Using the private temporary table emulation](1393-using-the-private-temporary-table-emulation.md)
