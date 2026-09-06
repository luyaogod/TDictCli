---
title: "SAVEPOINT"
source: "fgl-topics/c_fgl_transactions_SAVEPOINT.html"
breadcrumb: "SQL support > Database transactions > SAVEPOINT"
type: "concept"
---

# SAVEPOINT

> Defines or resets the position of a rollback point in the current transaction.

## Syntax

```
SAVEPOINT spVname [UNIQUE]
```

1. spname is the savepoint identifier.

## Usage

The `SAVEPOINT` instruction declares a new rollback label at the current position
in the lexical order within the current transaction. After defining a savepoint, you can rollback to
the specified point in the transaction by using the [`ROLLBACK WORK TO SAVEPOINT`](1111-rollback-work.md "Cancels and terminates a database transaction in the current connection.")
instruction.

If the same savepoint
name was used in a prior `SAVEPOINT` instruction, the
previous savepoint is destroyed and the name is reused
to flag the new rollback position. The optional `UNIQUE` keyword
specifies that you do not want to reuse the same savepoint
name in a subsequent `SAVEPOINT` instruction.
Reusing the same name after a `SAVEPOINT
spname UNIQUE` will raise an SQL
error.

## Example

In this example, a first savepoint is defined before the `INSERT`
statement, then reset before the `UPDATE` statement.
The `ROLLBACK TO SAVEPOINT` instruction will cancel the
`UPDATE` statement only:

```
MAIN
  DATABASE stock
  BEGIN WORK
  DELETE FROM items
  SAVEPOINT sp1
  INSERT INTO items VALUES ( ... )
  SAVEPOINT sp1  -- releases previous savepoint named sp1
  UPDATE items SET ...
  ROLLBACK WORK TO SAVEPOINT sp1
  COMMIT WORK
END MAIN
```

## Related links

**Related concepts**  

[RELEASE SAVEPOINT](1112-release-savepoint.md "Destroys the specified savepoint in the current transaction.")

[ROLLBACK WORK](1111-rollback-work.md "Cancels and terminates a database transaction in the current connection.")
