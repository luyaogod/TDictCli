---
title: "GROUP BY clause"
source: "fgl-topics/c_fgl_sql_programming_092.html"
breadcrumb: "SQL support > SQL programming > SQL portability > GROUP BY clause"
type: "concept"
description: "Some databases allow you to specify a column index in the GROUP BY clause: SELECT a, b, sum(c) FROM table GROUP BY 1,2 This is not possible with all database servers: Table 1. Database server support ..."
---

# GROUP BY clause

Some databases allow you to specify a column index in the GROUP
BY clause:

```
SELECT a, b, sum(c) FROM table GROUP BY 1,2
```

This is not possible with all database servers:

| Database Server Type | GROUP BY colindex, ... ? |
| --- | --- |
| IBM® Informix® | Yes |
| Microsoft™ SQL Server | No |
| Oracle® MySQL / MariadDB | Yes |
| Oracle Database Server | Yes, since Oracle 23aiSee `GROUP_BY_POSITION_ENABLED` |
| PostgreSQL | Yes |
| SQLite | Yes |
| Dameng® | No |

Search for GROUP BY in your SQL statements and use explicit column names.
