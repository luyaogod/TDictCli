---
title: "SELECT … FOR UPDATE"
source: "fgl-topics/c_fgl_odiagmys_013.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > BDL programming > SELECT … FOR UPDATE"
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

## Oracle® MySQL and MariaDB

MySQL and MariaDB support the `FOR UPDATE` clause in `SELECT`.

MySQL and MariaDB locking mechanism depends upon the transaction manager.

The default locking granularity is per table when you use the default non-transactional
configuration.

Use the InnoDB Storage Engine to get transactions and locking mechanisms.

Locks are released at the end of the transaction.

## Solution

Check if the MySQL storage engine supports `SELECT FOR UPDATE`, otherwise review
the program logic.

## Related links

**Related concepts**  

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
