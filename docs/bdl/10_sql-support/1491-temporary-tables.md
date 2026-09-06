---
title: "Temporary tables"
source: "fgl-topics/c_fgl_odiagsqt_016.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data manipulation > Temporary tables"
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

## SQLite

SQLite supports temporary tables with the `CREATE TEMP TABLE`
statement:

```
CREATE TEMP TABLE mytt1 ( pkey INT, name VARCHAR(50) )
CREATE TEMP TABLE mytt2 AS SELECT * FROM source
```

## Solution

Informix `CREATE TEMP TABLE` statements
are kept as is, while `SELECT INTO TEMP` statements are converted to SQLite native
SQL `CREATE TEMP TABLE AS SELECT ...`

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

The general FGLPROFILE entry to
control temporary table emulation
is:

```
dbi.database.dsname.ifxemul.temptables = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Temporary tables](1039-temporary-tables.md "Syntax for temporary table creation is not unique across all database engines.")
