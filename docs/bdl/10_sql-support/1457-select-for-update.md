---
title: "SELECT … FOR UPDATE"
source: "fgl-topics/c_fgl_odiagpgs_012.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > SELECT … FOR UPDATE"
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

## PostgreSQL

With PostgreSQL, locks are released when closing the cursor or when the transaction ends.

PostgreSQL locking granularity is at the row level.

PostgreSQL supports the `NOWAIT` keywoard after `FOR UPDATE`, and
has the `lock_timeout` session parameter to define a lock timeout.

## Solution

The database interface is based on an emulation of an Informix engine using transaction logging. Therefore, opening a `SELECT ... FOR
UPDATE` cursor declared outside a transaction will raise an SQL error -255 (not in
transaction).

When using the `SET LOCK MODE TO .. WAIT` instruction, the PostgreSQL driver
converts this statement to an `UPDATE pg_settings` command, to change the
`lock_timeout` parameter. For more details, see [Concurrency management](1422-concurrency-management.md).

> **Important:**
>
> With PostgreSQL, if an SQL error occurs during a transaction block, the whole transaction is
> aborted. Since `FOR UPDATE` cursors are typically opened inside a transaction block,
> in case of concurrent lock error, the current transaction cannot be continued.

When using the [Informix-style (pessimistic)
locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.") with `DECLARE CURSOR … FOR UPDATE + BEGIN WORK + OPEN + FETCH + UPDATE +
COMMIT WORK` where the `FETCH` sets an exclusive lock on a row, it is
possible to adapt this code to PostgreSQL, by using savepoints to continue the transaction in case
of lock error.

By default, PostgreSQL uses `READ COMMITTED` isolation level, and waits for locks.
This configuration is best for regular SQL (batch) processing. However, to immediately detect locks
with an SQL error, we need to temporarily set the lock wait mode to not wait, and reset the lock
wait mode after fetching the row.

Define the following functions to set and release a dedicated
savepoint:

```
FUNCTION lock_1_init()
    IF fgl_db_driver_type() == "pgs" THEN
       SAVEPOINT sp_lock_1
       SET LOCK MODE TO NOT WAIT
    END IF
END FUNCTION

FUNCTION lock_1_fini(sqlcode)
    DEFINE sqlcode INTEGER
    IF fgl_db_driver_type() == "pgs" THEN
        IF sqlcode < 0 THEN
            ROLLBACK WORK TO SAVEPOINT sp_lock_1
        END IF
        RELEASE SAVEPOINT sp_lock_1
        SET LOCK MODE TO WAIT
    END IF
END FUNCTION
```

In the application code doing the row locking, add calls to `lock_1_init()`
function to set the savepoint before the `FETCH`, and reset the savepoint (with
potential rollback to savepoint in case of lock error), by calling
`lock_1_fini(sqlca.sqlcode)` after the `FETCH`. Note that the
functions execute SQL statements and reset `sqlca.sqlcode` to
zero:

```
DEFINE can_update BOOLEAN
...
DECLARE c_lock CURSOR FOR SELECT ... FROM ... FOR UPDATE
...
BEGIN WORK
...
OPEN c_lock USING p_key
CALL lock_1_init() -- resets sqlca.sqlcode to zero!
WHENEVER ERROR CONTINUE
FETCH c_lock INTO rec.* -- may give a lock error
WHENEVER ERROR STOP
LET can_update = (sqlca.sqlcode == 0)
CALL lock_1_fini(sqlca.sqlcode) -- resets sqlca.sqlcode to zero!
IF can_update THEN
    ... user edits the record in dialog ...
    UPDATE ... SET ... WHERE pkey = p_key
    COMMIT WORK
ELSE
    ERROR "Could not fetch row for locking..."
    ROLLBACK WORK
END IF
```

## Related links

**Related concepts**  

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
