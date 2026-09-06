---
title: "Understanding positioned update or delete"
source: "fgl-topics/c_fgl_positioned_updates_002.html"
breadcrumb: "SQL support > Positioned updates/deletes > Understanding positioned update or delete"
type: "concept"
---

# Understanding positioned update or delete

> This is an introduction to SQL positionned UPDATE/DELETE.

When declaring a database cursor with a `SELECT` statement using a unique table
and ending with the `FOR UPDATE` keywords, you can modify the current row pointed by
the `FOR UPDATE` cursor with [`UPDATE ... WHERE CURRENT
OF`](1159-update-where-current-of.md "Updates the current row in a result set of a database cursor declared for update."), or the current row with [`DELETE ... WHERE CURRENT
OF`](1160-delete-where-current-of.md "Deletes the current row in a result set of a database cursor declared for update.") statements. Such an operation is called positioned update or
positioned delete.

Do not confuse positioned update with the use of `SELECT FOR UPDATE` statements
that are not associated with a database cursor. Executing `SELECT FOR UPDATE`
statements is supported by the language, but you cannot perform positioned updates since there is no
cursor identifier associated with the result set.

Some database servers do not support hold cursors (`WITH HOLD`)
declared with a `SELECT` statement including the `FOR UPDATE`
keywords. The SQL standards require for update cursors to be automatically closed at
the end of a transaction. Therefore, it is strongly recommended that you use positioned updates in a
transaction block.

To perform a positioned update or delete, perform a [`DECLARE` instruction with a
`SELECT FOR UPDATE` statement](1158-declare-select-for-update.md "Associate a database cursor with a SELECT statement to perform positioned updates and deletes").

![SELECT FOR UPDATE statement diagram](../_images/PUPFig01.jpg)

*SELECT FOR UPDATE statement*

Then, [start a transaction](1106-database-transactions.md "Database transaction concepts and handling."), [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") the cursor and [`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.") a row.

![Open a cursor diagram](../_images/PUPFig02.jpg)

*Open a cursor*

Then, [`UPDATE`](1159-update-where-current-of.md "Updates the current row in a result set of a database cursor declared for update.") or [`DELETE`](1160-delete-where-current-of.md "Deletes the current row in a result set of a database cursor declared for update.") the
current row with the `WHERE CURRENT OF` clause, before ending the transaction.

![Delete the row diagram](../_images/PUPFig03.jpg)

*Delete the row*

## Related links

**Related concepts**  

[Result set processing](1148-result-set-processing.md "Shows how to fetch rows from a database query.")

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

[Example 1: Positioned UPDATE statement](1162-example-1-positioned-update-statement.md "Example 1: Positioned UPDATE statement")
