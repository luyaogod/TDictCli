---
title: "Understanding SQL insert cursors"
source: "fgl-topics/c_fgl_InsertCursors_002.html"
breadcrumb: "SQL support > SQL insert cursors > Understanding SQL insert cursors"
type: "concept"
---

# Understanding SQL insert cursors

> This is an introduction to SQL insert cursors.

An insert cursor is a database cursor declared with a restricted form of the
`INSERT` statement, designed to perform buffered row insertion in database
tables.

The insert cursor simply inserts rows of data; it cannot be used to fetch data. When an insert
cursor is opened, a buffer is created in memory to hold a block of rows. The buffer receives rows
of data as the program executes `PUT` statements. The rows are written to disk
only when the buffer is full. You can use the `CLOSE` , `FLUSH` ,
or `COMMIT WORK` statement to flush the buffer when it is less than full. You must
close an insert cursor to insert any buffered rows into the database before the program ends. You
can lose data if you do not close the cursor properly.

When the database server supports buffered inserts, an insert cursor
increases processing efficiency (compared with embedding the `INSERT` statement
directly). This process reduces communication between the program
and the database server and also increases the speed of the insertions.

Before using the insert cursor, you must declare it with the [`DECLARE` instruction using an
`INSERT` statement](1165-declare-insert-cursor.md "The DECLARE with an INSERT instruction defines an insert cursor.").

![Declare a cursor diagram](../_images/INCFig01.jpg)

*Declaring a cursor*

Once declared, you can open the insert cursor with the [`OPEN`](1166-open-insert-cursor.md "Initializes an insert cursor.") instruction. This instruction prepares the insert buffer. When the
insert cursor is opened, you can add rows to the insert buffer with the [`PUT`](1167-put-insert-cursor.md "Adds a new row to the insert cursor buffer.") statement.

![OPEN and PUT statements diagram](../_images/INCFig02.jpg)

*OPEN and PUT statements*

Rows are automatically added to the database table when the insert buffer is full. To force row
insertion in the table, you can use the [`FLUSH`](1168-flush-insert-cursor.md "Flushes the buffer of an insert cursor.") instruction.

![FLUSH statement diagram](../_images/INCFig03.jpg)

*FLUSH statement*

Finally, when all rows are added, you can [`CLOSE`](1169-close-insert-cursor.md "Flushes and closes an insert cursor.") the cursor and if you no longer need it, you can deallocate
resources with the [`FREE`](1170-free-insert-cursor.md "Releases resources allocated for an insert cursor.")
instruction.

![CLOSE and FREE statements diagram](../_images/INCFig04.jpg)

*CLOSE and FREE statements*

By default, insert cursors must be opened inside a transaction block, with `BEGIN
WORK` and `COMMIT WORK`, and they are automatically closed at the end of
the transaction. If needed, you can declare insert cursors with the `WITH HOLD`
clause, to allow uninterrupted row insertion across multiple transactions.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
