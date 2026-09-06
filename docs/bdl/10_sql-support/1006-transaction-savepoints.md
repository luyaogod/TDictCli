---
title: "Transaction savepoints"
source: "fgl-topics/c_fgl_sql_programming_095.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Transaction savepoints"
type: "concept"
---

# Transaction savepoints

> Using transaction savepoints with different database engines.

IBM® Informix® IDS
11.50 introduced transaction savepoints, following the ANSI SQL standards. While most recent
database severs support savepoints, you must pay attention and avoid Informix specific features. For example, Oracle® (11), SQL Server (2008 R2) do not support the `RELEASE
SAVEPOINT` instruction.

| Database Server Type | SAVEPOINT & ROLLBACK WORK TO SAVEPOINT | RELEASE SAVEPOINT | SAVEPOINT UNIQUE | Related topic |
| --- | --- | --- | --- | --- |
| IBM Informix | Yes | Yes | Yes | [See details](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server (Only 2005+ with SNC driver) | Yes | No | No | [See details](1267-transactions-handling.md) |
| Oracle MySQL / MariadDB | Yes | Yes | No | [See details](1322-transactions-handling.md) |
| Oracle Database Server | Yes | No | No | [See details](1367-transactions-handling.md) |
| PostgreSQL | Yes | Yes | No | [See details](1423-transactions-handling.md) |
| SQLite | Yes | Yes | No | [See details](1474-transactions-handling.md) |
| Dameng® | Yes | No | No | [See details](1222-transactions-handling.md) |

## Related links

**Related concepts**  

[SAVEPOINT](1109-savepoint.md "Defines or resets the position of a rollback point in the current transaction.")
