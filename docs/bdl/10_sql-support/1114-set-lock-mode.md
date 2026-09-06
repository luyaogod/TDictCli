---
title: "SET LOCK MODE"
source: "fgl-topics/c_fgl_transactions_SET_LOCK_MODE.html"
breadcrumb: "SQL support > Database transactions > SET LOCK MODE"
type: "concept"
---

# SET LOCK MODE

> Defines the behavior of the program that tries to access a locked row or table.

## Syntax

```
SET LOCK MODE TO { NOT WAIT | WAIT[ seconds ] }
```

1. seconds defines the number of seconds to wait for lock acquisition.

## Usage

The `SET LOCK MODE` instruction
defines the timeout for lock acquisition for the current connection.

When
possible, the underlying database driver sets the corresponding connection
parameter to define the timeout for lock acquisition. But
some database servers may not support setting the lock timeout
parameter. In this case, the runtime system generates an exception.

When
using the `NOT WAIT` clause, the timeout is set to
zero. If the resource is locked, the database server ends
the operation immediately and raises an exception with the
SQL error.

When using the `WAIT` clause with a number of seconds, if the resource is locked,
the database server waits for the specified seconds, and then raises an exception with an SQL
error.

When using the `WAIT` clause
without a number of seconds, the database server waits for lock acquisition
for an infinite time.

With most database servers, the default
is to wait for locks to be released.

Make sure that the database
server and corresponding database driver both support a lock
acquisition timeout option, otherwise the program will raise an
exception.

## Example

```
MAIN
  DATABASE stock
  SET LOCK MODE TO WAIT 20
  ...
END MAIN
```

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
