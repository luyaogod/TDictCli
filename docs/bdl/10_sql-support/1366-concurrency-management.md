---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagora_011.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Database concepts > Concurrency management"
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

## ORACLE

When data is modified, exclusive locks are set and held until the end of the
transaction. For data consistency, ORACLE uses a multi-version consistency model: a
copy of the original row is kept for readers before performing writer modifications. Readers do not
have to wait for writers as in Informix. The simplest way
to think of Oracle's implementation of read consistency is to imagine each user accessing a private
copy of the database, hence the multi-version consistency model. The lock wait mode
cannot be changed session wide as in Informix; the
waiting behavior can be controlled with a `SELECT FOR UPDATE NOWAIT` only. Locks are
set at the row level in ORACLE, and this cannot be changed.

Control:

- Lock wait mode (on `SELECT` only): `SELECT ... FOR UPDATE
  NOWAIT`
- Isolation level: `SET TRANSACTION ISOLATION LEVEL TO ...`
- Explicit exclusive lock: `SELECT ... FOR UPDATE [NOWAIT]`

Defaults:

- The default isolation level is Read Committed ( readers cannot see uncommitted data, no shared
  lock is set when reading data ).

The main difference between Informix and ORACLE is
that readers do not have to wait for writers in ORACLE.

## Solution

The `SET ISOLATION TO ...` instruction is converted to `ALTER SESSION SET
ISOLATION_LEVEL ...` in Oracle. The next table shows the isolation level mappings done by
the database driver:

| SET ISOLATION instruction in program | Native SQL command |
| --- | --- |
| `SET ISOLATION TO DIRTY READ` | `ALTER SESSION SET ISOLATION_LEVEL = READ COMMITTED` |
| `SET ISOLATION TO COMMITTED READ [READ COMMITTED] [RETAIN UPDATE LOCKS]` | `ALTER SESSION SET ISOLATION_LEVEL = READ COMMITTED` |
| `SET ISOLATION TO CURSOR STABILITY` | `ALTER SESSION SET ISOLATION_LEVEL = READ COMMITTED` |
| `SET ISOLATION TO REPEATABLE READ` | `ALTER SESSION SET ISOLATION_LEVEL = SERIALIZABLE` |

ORACLE does not provide a dirty read mode, the (session wide) lock wait mode cannot be changed
and the locking precision is always at the row level.

See the Informix and
ORACLE documentation for more details about data consistency, concurrency
and locking mechanisms.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
