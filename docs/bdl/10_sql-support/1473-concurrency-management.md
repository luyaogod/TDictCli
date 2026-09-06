---
title: "Concurrency management"
source: "fgl-topics/c_fgl_odiagsqt_010.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Database concepts > Concurrency management"
type: "concept"
description: "Informix® is a multiuser database engine, while SQLite is typically used for a single-user application. SQLite 3 supports multiuser access to the same database file, but it is not designed for large ..."
---

# Concurrency management

Informix® is a multiuser
database engine, while SQLite is typically used for a single-user
application. SQLite 3 supports multiuser access to the same database
file, but it is not designed for large multiuser applications.

SQLite 3 supports two isolation levels: `SERIALIZABLE` (the default), and
`READ UNCOMMITTED`. The isolation level can be changed with the
`PRAGMA` command.

By default in the `SERIALIZABLE` isolation level, SQLite will raise an SQL error
if a program tries to access a database resource in use by another program. To avoid the SQL error
and force programs to wait for each other, programs define the behavior when the SQLite database is
busy (`SQLITE_BUSY`), with a specific API call. No SQL command exists for this.

## Solution

We recommend that you use SQLite for single-user DB applications. If several programs must access
the same SQLite database, each program must perform a `SET LOCK MODE TO WAIT`
instruction after the connection: `SET LOCK MODE` will be mapped to a call to the
`sqlite3_busy_timeout()` SQLite API function to get the same behavior as Informix, while `SET ISOLATION` instructions
will be ignored.

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Optimistic locking](1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.")

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
