---
title: "Row limiting clause (SELECT)"
source: "fgl-topics/c_fgl_sql_programming_097.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Row limiting clause (SELECT)"
type: "concept"
---

# Row limiting clause (SELECT)

> How to use the right clause to limit the number of rows produced by a SELECT statement?

IBM® Informix® SQL
supports the `SKIP` and `FIRST/LIMIT` keywords to limit the number of
rows of a result set.

For
example:

```
SELECT SKIP 10 FIRST 20 customer.* FROM customer ... ORDER BY cust_name
```

It is strongly recommended to use an `ORDER BY` clause when limiting the result
set rows.

The above Informix SQL syntax is not portable.

Recent database engines support the row limiting clause syntax defined by the SQL
standard:

```
SELECT ... OFFSET n ROWS FETCH FIRST m ROWS ONLY
```

This should be the prefered syntax to be used, if all target database types support this
`SELECT` clause.

The ODI database drivers can convert
the Informix SQL `SKIP/FIRST` row limiting clause to a native SQL equivalent, if the
row limiting clause parameters are simple integer literals (the clause is not translated when using
SQL parameters / program variables).
> **Important:**
>
> In addition to the
> `SKIP/FIRST` clause of the projection clause, Informix SQL supports also a
> `LIMIT` clause after the `ORDER BY`
> clause:
>
> ```
> SELECT customer.* FROM customer ... ORDER BY cust_name LIMIT 10
> ```
>
> This
> Informix SQL syntax construction is not converted by the ODI drivers. To benefit from the
> conversion, review the code to use the Informix SQL `SKIP/FIRST` clause
> instead.

| Database Server Type | Native equivalent | Informix emulation |
| --- | --- | --- |
| IBM Informix 12+ | `SELECT [SKIP n] FIRST m ...` | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server 2012+ | `SELECT ... ORDER BY ... OFFSET n ROWS [ FETCH FIRST m ROWS ONLY ]` | Emulated, [see details](1297-row-limiting-clause.md) |
| Oracle® MySQL 5.5-8.0 / MariadDB 10.x | `SELECT ... ORDER BY ... LIMIT m [OFFSET n]` | Emulated, [see details](1345-row-limiting-clause.md) |
| Oracle Database Server 19c+ | `SELECT ... ORDER BY ... [OFFSET n ROWS] FETCH FIRST m ROWS ONLY` | Emulated, [see details](1401-row-limiting-clause.md) |
| PostgreSQL 9.4+ | `SELECT ... ORDER BY ... [OFFSET n ROWS] FETCH FIRST m ROWS ONLY` | Emulated, [see details](1451-row-limiting-clause.md) |
| SQLite 3.8+ | `SELECT ... ORDER BY ... LIMIT m [OFFSET n]` | Emulated, [see details](1495-row-limiting-clause.md) |
| Dameng® 8+ | `SELECT ... ORDER BY ... [OFFSET n ROWS] FETCH FIRST m ROWS ONLY` | Emulated, [see details](1247-row-limiting-clause.md) |
