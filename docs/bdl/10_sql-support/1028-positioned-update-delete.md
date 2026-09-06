---
title: "Positioned UPDATE/DELETE"
source: "fgl-topics/c_fgl_sql_programming_079.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Positioned UPDATE/DELETE"
type: "concept"
---

# Positioned UPDATE/DELETE

> Using positioned updates/deletes with named database cursors.

The "`WHERE CURRENT OF cursor-name`" clause in `UPDATE`
and `DELETE` statements is not supported by all database engines.

| Database Server Type | WHERE CURRENT OF support |
| --- | --- |
| IBM® Informix® | Yes, [see details](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Yes, [see details](1304-update-delete-where-current-of.md) |
| Oracle® MySQL / MariadDB | No, [see details](1351-update-delete-where-current-of.md) |
| Oracle Database Server | Emulated, [see details](1407-update-delete-where-current-of.md) |
| PostgreSQL | Yes, [see details](1458-update-delete-where-current-of.md) |
| SQLite | No, [see details](1500-update-delete-where-current-of.md) |
| Dameng® | Yes, [see details](1252-update-delete-where-current-of.md) |

Some database drivers can emulate `WHERE CURRENT OF` mechanisms
by using rowids, but this requires additional processing. You should review the
code to disable this option.

The standard SQL solution is to use primary keys in all tables and write
`UPDATE` / `DELETE` statements with a
`WHERE` clause based on the primary key:

```
DEFINE rec RECORD
           id    INTEGER,
           name  CHAR(100)
     END RECORD
BEGIN WORK
  UPDATE CUSTOMER SET CUSTNAME = rec.name
        WHERE CUSTID = rec.id
  ...
COMMIT WORK
```
