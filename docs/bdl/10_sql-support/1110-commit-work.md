---
title: "COMMIT WORK"
source: "fgl-topics/c_fgl_transactions_COMMIT_WORK.html"
breadcrumb: "SQL support > Database transactions > COMMIT WORK"
type: "concept"
---

# COMMIT WORK

> Validates and terminates a database transaction in the current connection.

## Syntax

```
COMMIT WORK
```

## Usage

Use the `COMMIT WORK` instruction to commit all modifications made to the database
from the beginning of a transaction started with [`BEGIN WORK`](1108-begin-work.md "Starts a database transaction in the current connection."). The database server takes the required steps to make sure that
all modifications that the transaction makes are completed correctly and saved to disk.

`COMMIT WORK` is
part of the language syntax, the underlying database driver executes
the native SQL statement corresponding to this SQL instruction.

The `COMMIT WORK` statement releases all exclusive
locks that have been set during the transaction. With some databases,
shared locks are not released if the `FOR UPDATE` cursor
is declared `WITH HOLD` option. However, the `COMMIT
WORK` statement closes all cursors not declared with the `WITH
HOLD` option.

## Related links

**Related concepts**  

[Understanding database transactions](1107-understanding-database-transactions.md "This is an introduction to database transactions.")
