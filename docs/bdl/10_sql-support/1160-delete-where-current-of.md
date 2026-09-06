---
title: "DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_positioned_updates_DELETE_WHERE_CURRENT_OF.html"
breadcrumb: "SQL support > Positioned updates/deletes > DELETE … WHERE CURRENT OF"
type: "concept"
---

# DELETE … WHERE CURRENT OF

> Deletes the current row in a result set of a database cursor declared for update.

## Syntax

```
DELETE FROM table-specification
   WHERE CURRENT OF cid
```

1. table-specification identifies the target table
2. cid is the identifier of the database cursor
   declared for update.

## Usage

Use `DELETE ... WHERE CURRENT OF` to remove the row currently pointed by the
associated [`FOR UPDATE`](1158-declare-select-for-update.md "Associate a database cursor with a SELECT statement to perform positioned updates and deletes")
cursor.

The scope of reference of the cid cursor identifier is local to the module
where it is declared. Therefore, you must execute the `DECLARE`,
`UPDATE` or `DELETE` instructions in the same module.

There must be a current row in the result set. Make sure that the SQL status returned by the last
[`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.") is equal to zero.

After the deletion, no current row exists; you cannot use the cursor to delete or update a row
until you reposition the cursor with a `FETCH` statement.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

[UPDATE … WHERE CURRENT OF](1159-update-where-current-of.md "Updates the current row in a result set of a database cursor declared for update.")
