---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagdmg_011.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Database concepts > Concurrency management"
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

## Dameng®

As in Informix, Dameng uses locks to manage data consistency and concurrency. The database manager sets
exclusive locks on the modified rows and shared locks when data is read,
based on the isolation level. The locks are held until the end of the transaction. When
several processes want to modify the same data, the latest processes must wait until the first
finishes its transaction. Readers do not have to wait for writers. The lock granularity
is at the row level. For more details, see Dameng's documentation.

Control:

- Lock wait mode is to wait by default, and can be controlled by the `NOWAIT`
  option with `SELECT ... FOR UPDATE`.
- Transaction isolation level is defined by the `SET TRANSACTION ISOLATION LEVEL`
  command.
- Locking granularity: Row level.
- Explicit locking: `SELECT ... FOR UPDATE`

Defaults:

- The default isolation level is `READ COMMITTED`.

## Solution

The `SET ISOLATION TO ...` instruction is converted to a `SET TRANSACTION
ISOLATION LEVEL` command with Dameng. The next table shows the isolation level mappings
applied by the database driver:

| SET ISOLATION instruction in program | Dameng TX isolation level |
| --- | --- |
| `SET ISOLATION TO DIRTY READ` | `READ UNCOMMITTED` |
| `SET ISOLATION TO COMMITTED READ [READ COMMITTED] [RETAIN UPDATE LOCKS]` | `READ COMMITTED` |
| `SET ISOLATION TO CURSOR STABILITY` | `READ COMMITTED` |
| `SET ISOLATION TO REPEATABLE READ` | `REPEATABLE READ` |

See Informix and Dameng documentation for more details about data consistency, concurrency
and locking mechanisms.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
