---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagifx_027.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Fully supported IBM® Informix® SQL features > Concurrency management"
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

## Locking granularity

With short transactions / lock mode wait / isolation committed read (as recommended in [Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")), the locking granularity does not have to be at the row
level.

To improve performance with IBM® Informix, use the `LOCK MODE PAGE` locking level, which is
the default.

However, if the application requires row-level locking, use the `LOCK MODE ROW`
clause, or define the `DEF_TABLES_LOCKMODE` configuration parameter to
`ROW`, in the onconfig file of the IDS server.

## `LAST COMMITTED` option

IBM Informix IDS
11 has introduced the `LAST COMMITTED` option for the `COMMITTED READ`
isolation level, which makes IDS behave like other database servers using row-versioning.

With the `LAST COMMITTED` option, when a lock is set on a row by another process,
a `SELECT` statement will return the most recently committed version of that row,
rather than wait for a lock to be released, or get an SQL error when the lock wait timeout
occurs.

> **Tip:**
>
> The `LAST COMMITTED` option can be enabled for all programs, with
> the `USELASTCOMMITTED` configuration parameter, saving a lot of code changes.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
