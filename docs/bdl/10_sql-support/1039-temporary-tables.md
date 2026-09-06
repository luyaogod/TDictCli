---
title: "Temporary tables"
source: "fgl-topics/c_fgl_sql_programming_087.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Temporary tables"
type: "concept"
---

# Temporary tables

> Syntax for temporary table creation is not unique across all database engines.

Not all database servers support temporary tables. The engines supporting this feature
often provide it with a specific table creation statement:

| Database Server Type | Native temp table creation syntax | Temp table support |
| --- | --- | --- |
| IBM® Informix® | CREATE TEMP TABLE tablename ( col-defs ) SELECT ... INTO TEMP tablename | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | CREATE TABLE #tablename ( col-defs ) SELECT select-list INTO #tablename FROM ... | Emulated, [see details](1291-temporary-tables.md) |
| Oracle® MySQL / MariadDB | CREATE TEMPORARY TABLE tablename ( col-defs ) CREATE TEMPORARY TABLE tablename LIKE other-table | Emulated, [see details](1341-temporary-tables.md) |
| Oracle Database Server | CREATE GLOBAL TEMPORARY TABLE tablename ( col-defs ) CREATE GLOBAL TEMPORARY TABLE tablename AS SELECT ...or, since Oracle 18c:CREATE PRIVATE TEMPORARY TABLE tablename ( col-defs ) CREATE PRIVATE TEMPORARY TABLE tablename AS SELECT ... | Emulated, [see details](1390-temporary-tables.md) |
| PostgreSQL | CREATE TEMP TABLE tablename ( col-defs ) SELECT select-list INTO TEMP tablename FROM ... | Emulated, [see details](1446-temporary-tables.md) |
| SQLite | CREATE TEMP TABLE tablename ( col-defs ) | Emulated, [see details](1491-temporary-tables.md) |
| Dameng® | CREATE GLOBAL TEMPORARY TABLE tablename ( col-defs ) CREATE GLOBAL TEMPORARY TABLE tablename AS SELECT ... | Emulated, [see details](1241-temporary-tables.md) |

The behavior and limitations of temporary tables varies with the type of database server. See
[database adaptation guides](1178-sql-database-guides.md "This section includes the SQL guides for various supported database servers.") for more
details.

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

Consider reviewing programs using temporary tables, and adapt the code to create temporary tables
with native SQL syntax.
