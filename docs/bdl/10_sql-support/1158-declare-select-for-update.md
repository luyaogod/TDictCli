---
title: "DECLARE (SELECT … FOR UPDATE)"
source: "fgl-topics/c_fgl_positioned_updates_DECLARE.html"
breadcrumb: "SQL support > Positioned updates/deletes > DECLARE (SELECT … FOR UPDATE)"
type: "concept"
---

# DECLARE (SELECT … FOR UPDATE)

> Associate a database cursor with a SELECT statement to perform positioned updates and deletes

## Syntax

```
DECLARE cid [SCROLL] CURSOR [WITH HOLD]
    FOR { select-statement | sid }
```

1. cid is the identifier of the database cursor.
2. select-statement is a `SELECT` statement defined in static
   SQL, with the `FOR UPDATE` keywords.
3. sid is the identifier of a prepared `SELECT` statement
   including the `FOR UPDATE` keywords.

## Usage

`DECLARE ... FOR UPDATE` will define a cursor that can be used to do positioned
[updates](1159-update-where-current-of.md "Updates the current row in a result set of a database cursor declared for update.") and [deletes](1160-delete-where-current-of.md "Deletes the current row in a result set of a database cursor declared for update.") with the `WHERE
CURRENT OF` clause, and/or to set an exclusive lock on fetched rows in a transaction.

`DECLARE` must precede any other statement that refers to the cursor during
program execution.

To perform positioned updates, the select-statement must include the
`FOR UPDATE` keywords.

The scope of reference of the cid cursor identifier is local to the module
where it is declared. Therefore, you must execute the `DECLARE`,
`UPDATE` or `DELETE` instructions in the same module.

The static select-statement used in the `DECLARE` can contain ?
(question mark) parameter placeholders, that can be bound to program variables with the
`USING` clause of the `OPEN` instruction.

Use the `WITH HOLD` option carefully, because this feature is specific to IBM® Informix® servers.
Other database servers do not behave as Informix does
with such cursors. For example, if the `SELECT` is not declared `FOR
UPDATE`, most database servers keep cursors open after the end of a transaction. However,
some database servers automatically closes all cursors when the transaction is rolled back.

## Related links

**Related concepts**  

[Result set processing](1148-result-set-processing.md "Shows how to fetch rows from a database query.")

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
