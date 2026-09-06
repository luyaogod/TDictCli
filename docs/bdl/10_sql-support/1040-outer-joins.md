---
title: "Outer joins"
source: "fgl-topics/c_fgl_sql_programming_088.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Outer joins"
type: "concept"
description: "Use standard ISO outer join syntax instead of the old IBM Informix OUTER() syntax."
---

# Outer joins

> Use standard ISO outer join syntax instead of the old IBM® Informix® OUTER() syntax.

Old IBM Informix
SQL outer joins specified with the `OUTER` keyword in the `FROM` part
are not standard:

```
SELECT * FROM master, OUTER ( detail )
  WHERE master.mid = detail.mid
    AND master.cdate IS NOT NULL
```

| Database Server Type | Informix OUTER join support |
| --- | --- |
| IBM Informix | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Emulated, [see details](1289-outer-joins.md) |
| Oracle® MySQL / MariadDB | Emulated, [see details](1289-outer-joins.md) |
| Oracle Database Server | Emulated, [see details](1388-outer-joins.md) |
| PostgreSQL | Emulated, [see details](1444-outer-joins.md) |
| SQLite | Emulated, [see details](1489-outer-joins.md) |
| Dameng® | Emulated, [see details](1239-outer-joins.md) |

Most recent database servers now support the standard ANSI outer
join specification:

```
SELECT * FROM master LEFT OUTER JOIN detail ON (master.mid = detail.mid)
  WHERE master.cdate IS NOT NULL
```

it is recommended that you use recent database servers and use ANSI outer joins
only.
