---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagpgs_011.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Database concepts > Concurrency management"
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

## PostgreSQL

When data is modified, exclusive locks are set and held until the end of the
transaction. For data consistency, PostgreSQL uses a multi-version consistency model: A
copy of the original row is kept for readers before performing writer modifications. Readers do not
have to wait for writers as in Informix. The simplest way
to think of the PostgreSQL implementation of read consistency is to imagine each user operating a
private copy of the database, hence the multi-version consistency model. Since PostgreSQL 9.4, the
lock wait mode for the current SQL session can be changed by updating the
`'lock_timeout'` parameter of the `pg_settings` system view. Locks are
set at the row level in PostgreSQL and this cannot be changed.

Control:

- Lock wait mode: `UPDATE pg_settings SET setting=ms WHERE
  name='lock_timeout'`
- Isolation level (session-wide): `SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL ...`
- Isolation level (transaction-block): `SET TRANSACTION ISOLATION LEVEL ...`
- Explicit exclusive lock: `SELECT ... FOR UPDATE`

Defaults:

- The default isolation level is Read Committed.

The main difference between Informix and PostgreSQL is
that readers do not have to wait for writers in PostgreSQL.

## Solution

The `SET ISOLATION TO ...` instruction is converted to `SET SESSION
CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL ...` in PostgreSQL. The next table shows the
isolation level mappings done by the PostgreSQL database driver:

| SET ISOLATION instruction in program | Native SQL command |
| --- | --- |
| `SET ISOLATION TO DIRTY READ` | `SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL READ COMMITTED` |
| `SET ISOLATION TO COMMITTED READ [READ COMMITTED] [RETAIN UPDATE LOCKS]` | `SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL READ COMMITTED` |
| `SET ISOLATION TO CURSOR STABILITY` | `SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL READ COMMITTED` |
| `SET ISOLATION TO REPEATABLE READ` | `SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL SERIALIZABLE` |

When using `SET LOCK MODE ...` in the programs, it will be converted to an
`UPDATE pg_setting` instruction for PostgreSQL:

| SET LOCK MODE instruction in program | Native SQL command |
| --- | --- |
| `SET LOCK MODE TO WAIT` | `UPDATE pg_settings SET setting=0 WHERE name='lock_timeout'` |
| `SET LOCK MODE TO WAIT seconds` | `UPDATE pg_settings SET setting= (seconds*1000) WHERE name='lock_timeout'` |
| `SET LOCK MODE TO NOT WAIT` | `UPDATE pg_settings SET setting=1 WHERE name='lock_timeout'` |

See the Informix and PostgreSQL documentation
for more details about data consistency, concurrency and locking mechanisms.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
