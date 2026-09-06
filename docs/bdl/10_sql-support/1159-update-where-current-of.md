---
title: "UPDATE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_positioned_updates_UPDATE_WHERE_CURRENT_OF.html"
breadcrumb: "SQL support > Positioned updates/deletes > UPDATE … WHERE CURRENT OF"
type: "concept"
---

# UPDATE … WHERE CURRENT OF

> Updates the current row in a result set of a database cursor declared for update.

## Syntax

```
UPDATE table-specification
   SET
       column = { variable | sql-expression }      
       [,...] 
   WHERE CURRENT OF cid
```

1. table-specification identifies the target table
   (see UPDATE for more details).
2. column is a name of a table column.
3. variable is a program variable, a record member
   or an array member used as a parameter buffer
   to provide values.
4. sql-expression is an expression supported by
   the database server, this can be a literal or `NULL` for
   example.
5. cid is the identifier of the database cursor
   declared for update.

## Usage

Use `UPDATE ... WHERE CURRENT OF` to modify the values of the row currently
pointed by the associated [`FOR
UPDATE`](1158-declare-select-for-update.md "Associate a database cursor with a SELECT statement to perform positioned updates and deletes") cursor.

The scope of reference of the cid cursor
identifier is local to the module where it is declared.
Therefore, you must execute the `DECLARE`,`UPDATE` or `DELETE`
instructions in the same module.

There must be a current row in the result set. Make sure that the SQL status returned by the
last [`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.") is equal to zero.

If the `DECLARE` statement that created the cursor
specified one or more columns in the `FOR UPDATE` clause,
you are restricted to updating only those columns in a subsequent `UPDATE
... WHERE CURRENT OF` statement.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

[DELETE … WHERE CURRENT OF](1160-delete-where-current-of.md "Deletes the current row in a result set of a database cursor declared for update.")

[Example 1: Positioned UPDATE statement](1162-example-1-positioned-update-statement.md "Example 1: Positioned UPDATE statement")
