---
title: "Data definition statements"
source: "fgl-topics/c_fgl_sql_programming_060.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Data definition statements"
type: "concept"
---

# Data definition statements

> It is recommended to avoid use of DDL in programs.

When using SQL Data Definition Language (DDL) statements such as `CREATE TABLE`,
`ALTER TABLE`, `DROP TABLE`, only a limited SQL syntax works on all
database servers. Most databases support `NOT NULL`, `CHECK`,
`PRIMARY KEY`, `UNIQUE`, `FOREIGN KEY` constraints, but
the syntax for naming constraints is different.

The following statement works with most database servers and creates
a table with equivalent properties in all cases:

```
 CREATE TABLE customer  (
     cust_id INTEGER NOT NULL,
     cust_name CHAR(50) NOT NULL,
     cust_lastorder DATE NOT NULL,
     cust_group INTEGER,
       PRIMARY KEY (cust_id),
       UNIQUE (cust_name),
       FOREIGN KEY (cust_group) REFERENCES group (group_id)
 )
```

Some engines like SQL Server have a different default behavior
for `NULL` columns when you create a table. You
may need to set up database properties to make sure that a column
allows nulls if the `NOT NULL` constraint is
not specified.

When executing `CREATE TABLE` or `ALTER TABLE` DDL statements, and
Informix emulation is enabled, the ODI drivers are able to convert Informix data types like
`DATETIME YEAR TO MINUTE` to a native SQL equivalent. This is especially required
when creating temporary tables during program execution. However, the ODI drivers will not convert
Informix specific table definition rules such as constraint naming.

To create tables in programs using non-standard clauses (for example to define storage options),
use `EXECUTE IMMEDIATE` with a string containing the SQL text, and adapt the
statement to the target database server.

| Database Server Type | Related topic |
| --- | --- |
| Microsoft™ SQL Server | [See details](1283-sql-table-definition.md) |
| Oracle® MySQL / MariadDB | [See details](1334-sql-table-definition.md) |
| Oracle Database Server | [See details](1381-sql-table-definition.md) |
| PostgreSQL | [See details](1437-sql-table-definition.md) |
| SQLite | [See details](1485-sql-table-definition.md) |
| Dameng® | [See details](1235-sql-table-definition.md) |

## Related links

**Related concepts**  

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
