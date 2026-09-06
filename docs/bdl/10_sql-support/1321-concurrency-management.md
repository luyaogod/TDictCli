---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagmys_012.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Database concepts > Concurrency management"
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

## Oracle® MySQL and MariaDB

When data is modified, exclusive locks are set and held until the end of the
transaction. For data consistency, MySQL uses a locking mechanism. Readers must wait
for writers as in Informix.

Control:

- No SQL instruction for lock wait timeout control is provided. The lock wait timeout can be
  defined at the database engine level. If the storage engine is InnoDB, see the
  `innodb_lock_wait_timetout` configuration parameter.
- Isolation level: `SET TRANSACTION ISOLATION LEVEL ...`
- Explicit exclusive lock: `SELECT ... FOR UPDATE`

Defaults:

- The default isolation level is Read Committed.
- The default locking granularity is per table (per page when using BDB tables).

## Solution

Since there is no SQL instruction to define the lock wait timeout for the current session with
Oracle MySQL and MariaDB, executing the
`SET LOCK MODE` instruction will produce an SQL error [-6370](../15_library-reference/4483-genero-bdl-errors.md). Avoid `SET LOCK
MODE` with this kind of database engine.

The `SET ISOLATION TO ...` instruction is converted to `SET SESSION
TRANSACTION ISOLATION LEVEL ...` in MySQL. The table shows the isolation level mappings
applied by the MySQL database driver:

| SET ISOLATION instruction in program | Native SQL command |
| --- | --- |
| `SET ISOLATION TO DIRTY READ` | `SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED` |
| `SET ISOLATION TO COMMITTED READ [READ COMMITTED] [RETAIN UPDATE LOCKS]` | `SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED` |
| `SET ISOLATION TO CURSOR STABILITY` | `SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED` |
| `SET ISOLATION TO REPEATABLE READ` | `SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ` |

See Informix and MySQL documentation
for more details about data consistency, concurrency and locking mechanisms.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
