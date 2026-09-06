---
title: "SELECT … FOR UPDATE"
source: "fgl-topics/c_fgl_odiagora_012.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > SELECT … FOR UPDATE"
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

## ORACLE

ORACLE allows individual and exclusive row locking
with:

```
SELECT ... FOR UPDATE [OF col-list]
```

A lock is acquired for each selected row when the cursor is opened, before the first fetch.

Cursors using `SELECT ... FOR UPDATE` are automatically closed when the
transaction ends.

Locks are not released when a cursor is closed.

ORACLE's locking granularity is at the row level.

The `NOWAIT` keyword can be used in `SELECT ... FOR UPDATE`
statement, the return immediately if the row is already locked by another
user:

```
SELECT ... FOR UPDATE [ OF col-list ] NOWAIT
```

## Solution

> **Important:**
>
> Cursors declared with `SELECT ... FOR UPDATE` using the
> `WITH HOLD` clause cannot be supported with ORACLE.

The database interface is based on an emulation of
an Informix engine using transaction logging.

Opening a `SELECT ... FOR UPDATE` cursor declared outside a transaction will
raise an SQL error -255 (not in transaction).

When using pessimistic locking with `DECLARE ... CURSOR FOR SELECT ...
FOR UPDATE`, review the program logic to have `OPEN` and
`CLOSE` instructions inside transactions (`BEGIN WORK` +
`COMMIT WORK` / `ROLLBACK WORK`).

See also [Cursors with Hold](1405-cursors-with-hold.md) and [UPDATE/DELETE WHERE CURRENT OF](1407-update-delete-where-current-of.md)
for more details.

## Related links

**Related concepts**  

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
