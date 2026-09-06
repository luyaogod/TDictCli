---
title: "Temporary tables"
source: "fgl-topics/c_fgl_odiagmsv_023.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > Temporary tables"
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

## Microsoft™ SQL Server

Microsoft SQL Server provides local (SQL session wide)
or global (database wide) temporary tables by using the '#' or '##' characters as table name
prefix.

The following SQL Server statement creates a session temporary table with the `CREATE
TABLE` statement:

```
CREATE TABLE #temp1 ( kcol INTEGER, .... )
```

To create and fill a session temporary table in one step, use the SQL Server `SELECT ...
INTO` statement:

```
SELECT * INTO #temp2 FROM customers WHERE ...
```

> **Important:**
>
> When not specifying the collation in the column definition, `CHAR/VARCHAR` columns
> of SQL Server temporary tables created with `CREATE TABLE
> #tabname` get by default the collation of the `tempdb`
> system database. If the collation of the temp table columns does not match the current database
> collation, SQL Server will return the error `468 "Cannot resolve the collation conflict
> between A and B in the ..."` when comparing `CHAR/VARCHAR` columns from
> permanent and temporary tables, as in the next
> query:
>
> ```
> SELECT COUNT(*) FROM #temptab1 WHERE name IN (SELECT cust_name FROM customers)
> ```
>
> To create columns with the database collation and avoid the 468 collation mismatch error, one
> must add the `COLLATE DATABASE_DEFAULT` clause after `CHAR/VARCHAR`
> and `NCHAR/NVARCHAR` types in the `CREATE TABLE
> #tabname` statement. Alternatively, you can change the collation of the
> `model` system database, which is used to create `tempdb`, each time
> SQL Server starts up.
>
> When using the `SELECT * INTO #tabname` SQL Server statement, the
> character type columns of the temp table get the current database collation automatically.

## Solution

In BDL, Informix temporary tables instructions
are converted to generate native SQL Server temporary tables.

When Informix emulations are enabled, `CREATE TEMP TABLE
tabname` statements executed by Genero programs are converted to
`CREATE TABLE #tabname` for SQL Server, and `SELECT
<select-list> <select-body> INTO TEMP tabname` statements are
converted to SQL Server `SELECT <select-list> INTO #tabname
<select-body>`.

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

To avoid 468 SQL errors when comparing character type columns of temp tables and permanent
tables, the collation of the temporary table columns must match the collation of the permanent
table.
> **Tip:**
>
> At SQL Server startup, the `tempdb` system database is recreated from the
> `model` system database. Database options such as the collation are inherited from
> the `model` database. While changing the options of the model system database is
> quite challenging (to be done in SQL Server single-user mode), this can be a good option to get a
> database collation in `tempdb` that matches your application database collation.

When using a converted Informix `SELECT ... INTO TEMP tabname`
statement, or a native SQL Server `SELECT ... INTO #tabname` statement, the character
columns of the temporary table will automatically get the default collation of the current database,
and the error 468 will not occur when comparing character columns.

When using a converted Informix `CREATE TEMP TABLE tabname`
statement, or the native SQL Server `CREATE TABLE #tabname` statement, the temporary
table columns get by default the collation of the `tempdb` system database, inherited
from the SQL Server instance collation. To avoid 468 errors, only for `CREATE TEMP
TABLE` statements, the ODI driver will add the `COLLATE DATABASE_DEFAULT`
clause to SQL Server character types corresponding to Informix `CHAR`,
`VARCHAR`, `LVARCHAR` and `TEXT` type names. For
example, a `CREATE TEMP TABLE tt1 ( col1 VARCHAR(20) NOT NULL )` will be converted to
`CREATE TABLE #tt1 ( col1 VARCHAR(20) COLLATE DATABASE_DEFAULT NOT NULL )`

> **Tip:**
>
> To avoid having the ODI drivers add the `COLLATE DATABASE_DEFAULT`
> clause, disable Informix emulation switches for character types, and replace the orignal Informix types
> by the required SQL Server types:
>
> ```
> dbi.database.dsname.datatype.char = false
> dbi.database.dsname.datatype.varchar = false
> dbi.database.dsname.datatype.text = false
> ```
>
> Another
> option is to use dynamic SQL and add the [fglhint](1147-fglhint-sql-comments.md "Using special SQL comment hints to control statement execution.")
> comment to disable Informix emulations and use directly the native SQL Server
> syntax:
>
> ```
> EXECUTE IMMEDIATE "/* fglhint_no_ifxemul */ CREATE TABLE #tt1 ... "
> ```

The general FGLPROFILE entry to
control temporary table emulation
is:

```
dbi.database.dsname.ifxemul.temptables = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

> **Note:**
>
> Microsoft SQL Server does not support scroll
> cursors based on a temporary table.

## Related links

**Related concepts**  

[Temporary tables](1039-temporary-tables.md "Syntax for temporary table creation is not unique across all database engines.")
