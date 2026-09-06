---
title: "PUT (insert cursor)"
source: "fgl-topics/c_fgl_InsertCursors_PUT.html"
breadcrumb: "SQL support > SQL insert cursors > PUT (insert cursor)"
type: "concept"
---

# PUT (insert cursor)

> Adds a new row to the insert cursor buffer.

## Syntax

```
PUT cid FROM pvar [,...]
```

1. cid is the identifier of the insert cursor.
2. pvar is a variable containing an input value for the new row.

## Usage

The `PUT` instruction
adds a row to the insert cursor buffer.

If the insert cursor was not declared `WITH HOLD` option, the `PUT`
instruction generates an SQL error if there is no current transaction started.

If
the insert buffer has no room for the new row when the statement executes,
the buffered rows are written to the database in a block, and the
buffer is emptied. As a result, some `PUT` statement
executions cause rows to be written to the database, and some do not.

## Related links

**Related concepts**  

[OPEN (insert cursor)](1166-open-insert-cursor.md "Initializes an insert cursor.")

[FLUSH (insert cursor)](1168-flush-insert-cursor.md "Flushes the buffer of an insert cursor.")

[Variables](../08_language-basics/0686-variables.md "Explains how to define program variables.")

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
