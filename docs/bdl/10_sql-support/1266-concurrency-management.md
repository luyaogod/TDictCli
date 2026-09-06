---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagmsv_013.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Database concepts > Concurrency management"
type: "concept"
description: "Data consistency and concurrency concepts Data Consistency applies to situations when readers want to access data currently being modified by writers. Concurrent Data Access applies to situations when ..."
---

# Concurrency management

## Data consistency and concurrency concepts

- Data Consistency applies to situations when readers want to access data currently
  being modified by writers.
- Concurrent Data Access applies to situations when several writers are accessing the
  same data for modification.
- Locking Granularity defines the amount of data concerned when a lock is set (for
  example, row, page, table).

## Informix®

Informix uses a locking mechanism to handle data
consistency and concurrency. When a process changes database information with
`UPDATE`, `INSERT` or `DELETE`, an exclusive
lock is set on the touched rows. The lock remains active until the end of the transaction.
Statements performed outside a transaction are treated as a transaction containing a single
operation and therefore release the locks immediately after execution. `SELECT`
statements can set shared locks, depending on isolation level. In case of
locking conflicts (for example, when two processes want to acquire an exclusive lock on the same row
for modification, or when a writer is trying to modify data protected by a shared lock), the
behavior of a process can be changed by setting the lock wait mode.

Control:

- Lock wait mode: `SET LOCK MODE TO ...`
- Isolation level: `SET ISOLATION TO ...`
- Locking granularity: `CREATE TABLE ... LOCK MODE {PAGE|ROW}`
- Explicit exclusive lock: `SELECT ... FOR UPDATE`

Defaults:

- The default isolation level is `READ COMMITTED`.
- The default lock wait mode is `NOT WAIT`.
- The default locking granularity is `PAGE`.

## Microsoft™ SQL Server

As in Informix, SQL Server uses locks to manage data
consistency and concurrency. The database manager sets exclusive locks on the modified rows
and shared locks or  update locks when data is read, based on the isolation
level. The locks are held until the end of the transaction. When multiple processes want to
access the same data, the latest processes must wait until the first finishes its transaction or the
lock timeout occurs. The locking strategy of SQL Server is row locking with possible promotion to
page or table locking. SQL Server dynamically determines the appropriate level at which to place
locks for each Transact-SQL statement.

SQL Server supports snapshot isolation level, to force the current transaction to use a copy of
the original row, when it is changed by another transaction. With snapshot isolation enabled,
writers do not block readers, and readers do not block writers, as they would with the default
behavior of `READ COMMIT` isolation level. Snapshot isolation level significantly
reduces deadlocks for complex transactions. To enable snapshot isolation in SQL Server, use
`ALTER DATABASE ALLOW_SNAPSHOT_ISOLATION ON`. Furthermore, `ALTER DATABASE
READ_COMMITTED_SNAPSHOT ON` allows access to versioned rows under the default `READ
COMMITTED` isolation level (otherwise, snapshot isolation must be specified by every SQL
Session). For more details about snapshot isolation level, see SQL Server documentation.

Control:

- Lock wait mode: `SET LOCK_TIMEOUT milliseconds` (returns error
  1222 on time out).
- Isolation level: `SET TRANSACTION ISOLATION LEVEL ...`
- Locking granularity: Row, Page, or Table level (Automatic - See Dynamic Locking).
- Explicit locking: `SELECT ... FROM ... WITH (UPDLOCK)` (See Locking Hints)

Defaults:

- The default isolation level is `READ COMMITTED` (readers cannot see uncommitted
  data).
- The default `LOCK_TIMEOUT` is -1 (indicates no timeout period, wait
  forever).

## Solution

The `SET ISOLATION TO ...` in programs is converted to `SET TRANSACTION
ISOLATION LEVEL ...` for SQL Server. The table shows the isolation level mappings applied by
the database driver:

| SET ISOLATION instruction in program | Native SQL command |
| --- | --- |
| `SET ISOLATION TO DIRTY READ` | `SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED` |
| `SET ISOLATION TO COMMITTED READ [READ COMMITTED] [RETAIN UPDATE LOCKS]` | `SET TRANSACTION ISOLATION LEVEL READ COMMITTED` |
| `SET ISOLATION TO CURSOR STABILITY` | `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ` |
| `SET ISOLATION TO REPEATABLE READ` | `SET TRANSACTION ISOLATION LEVEL SERIALIZABLE` |

When
using `SET LOCK MODE ...` in the programs, it will
be converted to a `SET LOCK_TIMEOUT` instruction
for SQL Server:

| SET LOCK MODE instruction in program | Native SQL command |
| --- | --- |
| `SET LOCK MODE TO WAIT` | `SET LOCK_TIMEOUT -1` (wait forever) |
| `SET LOCK MODE TO WAIT seconds` | `SET LOCK_TIMEOUT (seconds*1000)` (wait N milliseconds) |
| `SET LOCK MODE TO NOT WAIT` | `SET LOCK_TIMEOUT 0` (do not wait) |

See Informix and
SQL Server documentation for more details about data consistency,
concurrency and locking mechanisms.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
