---
title: "SET ISOLATION"
source: "fgl-topics/c_fgl_transactions_SET_ISOLATION.html"
breadcrumb: "SQL support > Database transactions > SET ISOLATION"
type: "concept"
---

# SET ISOLATION

> Defines the transaction isolation level for the current connection.

## Syntax

```
SET ISOLATION TO
  { DIRTY READ
  | COMMITTED READ [LAST COMMITTED] [RETAIN UPDATE LOCKS]
  | CURSOR STABILITY
  | REPEATABLE READ }
```

## Usage

The `SET ISOLATION` instruction
sets the transaction isolation level for the current connection.
See database concepts in your database server documentation for more
details about isolation levels and concurrency management.

When
possible, the underlying database driver sets the corresponding transaction
isolation level. If the isolation level cannot be set, the
runtime system generates an exception.

When using the `DIRTY
READ` isolation level, the database server might return a
phantom row, which is an uncommitted row that was inserted or modified
within a transaction that has subsequently rolled back. No other isolation
level allows access to a phantom row.

On most database servers,
the default isolation level is `COMMITTED READ`, which
is appropriate to portable database programming.

The `LAST
COMMITTED` and `RETAIN UPDATE LOCKS` options
have been added to the language syntax for conformance with IBM® Informix® IDS 11. The `LAST
COMMITTED` option can be turned on implicitly with a server
configuration parameter, saving unnecessary code changes.

## Example

```
MAIN
  DATABASE stock
  SET ISOLATION TO COMMITTED READ
  ...
END MAIN
```

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
