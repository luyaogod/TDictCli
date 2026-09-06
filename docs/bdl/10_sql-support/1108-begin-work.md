---
title: "BEGIN WORK"
source: "fgl-topics/c_fgl_transactions_BEGIN_WORK.html"
breadcrumb: "SQL support > Database transactions > BEGIN WORK"
type: "concept"
---

# BEGIN WORK

> Starts a database transaction in the current connection.

## Syntax

```
BEGIN WORK
```

## Usage

Use the `BEGIN
WORK` instruction to indicate where the database transaction
starts in your program. Each row that an `UPDATE`, `DELETE`,
or `INSERT` statement affects during a transaction
is locked and remains locked throughout the transaction.

`BEGIN WORK` is part of the language syntax,
the underlying database driver executes the native SQL statement
corresponding to this SQL instruction.

The transaction block is ended by a [`COMMIT WORK`](1110-commit-work.md "Validates and terminates a database transaction in the current connection.") or [`ROLLBACK WORK`](1111-rollback-work.md "Cancels and terminates a database transaction in the current connection.") instruction.

## Example

The next code example starts a transaction block, inserts a row and updates
the row, then commits the transaction.
To other users, the `INSERT` and `UPDATE`
instruction will be seen as an single atomic database modification:

```
MAIN
  DATABASE stock
  BEGIN WORK
  INSERT INTO items VALUES ( ... )
  UPDATE items SET ...
  COMMIT WORK
END MAIN
```

## Related links

**Related concepts**  

[Understanding database transactions](1107-understanding-database-transactions.md "This is an introduction to database transactions.")
