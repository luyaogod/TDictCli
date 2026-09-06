---
title: "ROLLBACK WORK"
source: "fgl-topics/c_fgl_transactions_ROLLBACK_WORK.html"
breadcrumb: "SQL support > Database transactions > ROLLBACK WORK"
type: "concept"
---

# ROLLBACK WORK

> Cancels and terminates a database transaction in the current connection.

## Syntax

```
ROLLBACK WORK [TO SAVEPOINT [spname]]
```

- spname is the savepoint identifier.

## Usage

Use `ROLLBACK WORK` to cancel the current transaction and invalidate all
changes since the beginning of the transaction started with [`BEGIN WORK`](1108-begin-work.md "Starts a database transaction in the current connection."). After the execution
of this instruction, the database is restored to the state that it was in before the transaction
began. All row and table locks that the canceled transaction holds are released. If you issue this
statement when no transaction is pending, an error occurs.

`ROLLBACK WORK` is part of the language syntax, the underlying database driver
executes the native SQL statement corresponding to this SQL instruction.

When specifying a [savepoint](1109-savepoint.md "Defines or resets the position of a rollback point in the current transaction.") with the
`TO SAVEPOINT` clause, all SQL statements executed since the specified savepoint
will be canceled. The transaction is not canceled, however, and you can continue to execute other
SQL statements.

## Example

This example checks for a potential SQL error after the
`DELETE` statement and cancels the complete transaction
with a `ROLLBACK` instruction:

```
MAIN
  DATABASE stock
  WHENEVER ERROR CONTINUE
  BEGIN WORK
  INSERT INTO orders_hist VALUES ( ... )
  DELETE FROM orders WHERE ...
  IF sqlca.sqlcode < 0 THEN
    ROLLBACK WORK
  ELSE
    COMMIT WORK
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Understanding database transactions](1107-understanding-database-transactions.md "This is an introduction to database transactions.")
