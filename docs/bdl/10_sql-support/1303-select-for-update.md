---
title: "SELECT … FOR UPDATE"
source: "fgl-topics/c_fgl_odiagmsv_014.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > SELECT … FOR UPDATE"
type: "concept"
description: "Informix® Legacy BDL programs typically use a cursor with SELECT FOR UPDATE to implement pessimistic locking and avoid several users editing the same rows: DECLARE cc CURSOR FOR SELECT ... FROM tab ..."
---

# SELECT … FOR UPDATE

## Informix®

Legacy BDL programs typically use a cursor with `SELECT FOR UPDATE` to implement
pessimistic locking and avoid several users editing the same rows:

```
DECLARE cc CURSOR FOR
SELECT ... FROM tab WHERE ... FOR UPDATE
OPEN cc
FETCH cc <-- lock is acquired
...
CLOSE cc <-- lock is released
```

The row must be fetched in order to set the lock.

If the cursor is local to a transaction, the lock is released when the transaction ends. If the
cursor is declared `WITH HOLD`, the lock is released when the cursor is closed.

Informix provides the `SET LOCK MODE`
instruction to define the lock wait
timeout:

```
SET LOCK MODE TO { WAIT | NOT WAIT | WAIT seconds }
```

The
default mode is `NOT WAIT`.

## Microsoft™ SQL Server

Microsoft SQL Server allows individual and exclusive
row locking by using server cursors with Scroll Locks option, combined with SQL hints
such as `UPDLOCK` in the `FROM`
clause::

```
SELECT ... FROM tab1 WITH(UPDLOCK) WHERE ...
```

> **Note:**
>
> In Transact-SQL, the `FOR UPDATE [OF col-list]` clause is
> not part of the `SELECT` syntax: It is part of T-SQL `DECLARE CURSOR`
> syntax.

Individual locks are acquired when fetching the rows.

When the cursor is opened outside a transaction (BDL `WITH HOLD` cursor option),
locks are released when the cursor is closed.

When the cursor is opened inside a transaction, locks are released when the transaction ends.

SQL Server's locking granularity is at the row level, page level or table level (the level is
automatically selected by the engine for optimization).

## Solution

When executing a `SELECT ... FOR UPDATE` in the program, the SQL Server database
drivers remove the `FOR UPDATE` clause from the SQL text and set the ODBC cursor
attribute `SQL_ATTR_CONCURRENCY` to `SQL_CONCUR_LOCK`. This enables
Scroll Locks concurrency in the server cursor.

The Scroll Locks option implements pessimistic concurrency control, in which the
application attempts to lock the underlying database rows at the time they are read into the cursor
result set.

SQL Server Transact-SQL hints such as `UPDLOCK` can be used to fine-tune the
locking semantics that will be used by SQL Server.

When using server cursors with Scroll Locks option, an update lock is placed on the
row when it is read into the cursor. If the cursor is opened within a transaction, the transaction
update lock is held until the transaction is either committed or rolled back; the cursor lock is
dropped when the next row is fetched.

If the cursor has been opened outside a transaction, the lock is dropped when the next row is
fetched.

Therefore, it is recommended that a cursor is opened in a transaction whenever the user wants
full pessimistic concurrency control.

An update lock prevents any other task from acquiring
an update or exclusive lock, which prevents any other task from updating the row.

An update lock, however, does not block a shared lock, so it does not prevent other tasks from
reading the row unless the second task is also requesting a read with an update lock.

`SELECT FOR UPDATE` statements are well supported in BDL as long as they are used
inside a transaction. Avoid cursors declared `WITH HOLD`.

> **Note:**
>
> The database interface is based on an emulation of an Informix engine using transaction logging. Therefore, opening a
> `SELECT ... FOR UPDATE` cursor declared outside a transaction will raise an SQL error
> -255 (not in transaction).

The `SELECT FOR UPDATE` statement cannot contain an `ORDER BY`
clause if you want to perform positioned updates/deletes with the [`WHERE CURRENT OF`](1304-update-delete-where-current-of.md) clause.

Cursors declared with `SELECT ... FOR UPDATE` using the `WITH HOLD`
clause cannot be supported with SQL Server.

Review the program logic if you use pessimistic locking, because it is based on the `NOT
WAIT` mode which is not supported by SQL Server.

## Related links

**Related concepts**  

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
