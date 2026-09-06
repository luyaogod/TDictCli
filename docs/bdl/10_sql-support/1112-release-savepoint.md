---
title: "RELEASE SAVEPOINT"
source: "fgl-topics/c_fgl_transactions_RELEASE_SAVEPOINT.html"
breadcrumb: "SQL support > Database transactions > RELEASE SAVEPOINT"
type: "concept"
---

# RELEASE SAVEPOINT

> Destroys the specified savepoint in the current transaction.

## Syntax

```
RELEASE SAVEPOINT spname
```

- spname is the savepoint identifier.

## Usage

Use the `RELEASE SAVEPOINT` instruction to delete a savepoint defined by the
[`SAVEPOINT`](1109-savepoint.md "Defines or resets the position of a rollback point in the current transaction.") instruction.

See database documentation for more details about the behavior of this SQL statement.

Note for example that IBM® Informix® IDS will also release any savepoint that has been declared
between the specified savepoint and the [`RELEASE SAVEPOINT`](1112-release-savepoint.md "Destroys the specified savepoint in the current transaction.") instruction.

## Example

In the next example, the `RELEASE SAVEPOINT` instruction
cancels the `UPDATE` and `INSERT` statements
and destroys the sp1 and sp2 savepoints. Only the `DELETE`
statement will take effect at the end of the transaction:

```
MAIN
  DATABASE stock
  BEGIN WORK
  DELETE FROM items
  SAVEPOINT sp1
  INSERT INTO items VALUES ( ... )
  SAVEPOINT sp2
  UPDATE items SET ...
  RELEASE SAVEPOINT sp1
  ROLLBACK WORK TO SAVEPOINT
  COMMIT WORK
END MAIN
```

## Related links

**Related concepts**  

[Understanding database transactions](1107-understanding-database-transactions.md "This is an introduction to database transactions.")
