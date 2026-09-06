---
title: "Concurrent data access"
source: "fgl-topics/c_fgl_sql_programming_070.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Concurrent data access"
type: "concept"
---

# Concurrent data access

> Understanding concurrent data access and data consistency.

Data concurrency is the simultaneous access of the same data by many users.
Data consistency means that each user sees a consistent view of the data, while other
users modify data through SQL transactions. The database systems are responsible for data
concurrency and data consistency. Each database server has its own behavior and specifics regarding
concurrency and consistency. Before writing inter-operable database applications, carefully read the
documentation of all database engine you want to support.

The program code must be prepared to support the behavior of the database server regarding
concurrency and consistency management. Programmers must have a good knowledge of multiuser
application programming, transactions, locking mechanisms, isolation levels and wait mode.

Application processes accessing the database can change transaction parameters such as the
isolation level and lock wait mode. Existing programs might have to be adapted in order to work with
this feature, as it changes the behavior of the SQL statement execution.

Basically, when a process modifies data during a transaction, other (writing) processes will have
to wait, until the first process terminates its transaction. The longer the other processes have to
wait, the worse the user experience becomes. Keep this in mind when programming with SQL.

The best recommendation to ensure consistent behavior with different types of database engines is
to adhere to the following practice:

- The database must support transactions; this is usually the case.
- Transactions must be as short as possible (under a second is fine, 3 or more seconds is a
  considered as a long transaction).
- The isolation level should be set to `COMMITTED READ` or `CURSOR
  STABILITY`.
- The wait mode for locks must be `WAIT` or `WAIT n` (timeout). Wait
  mode can be adapted to wait for the longest transaction.

| Database Server Type | Concurrency topic |
| --- | --- |
| IBM® Informix® | [Concurrency in IBM Informix](1187-concurrency-management.md) |
| Microsoft™ SQL Server | [Concurrency in SQL Server](1266-concurrency-management.md) |
| Oracle® MySQL | [Concurrency in Oracle MySQL](1321-concurrency-management.md) |
| Oracle Database Server | [Concurrency in Oracle DB](1366-concurrency-management.md) |
| PostgreSQL | [Concurrency in PostgreSQL](1422-concurrency-management.md) |
| SQLite | [Concurrency in SQLite](1473-concurrency-management.md) |
| Dameng® | [Concurrency in Dameng](1221-concurrency-management.md) |

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
