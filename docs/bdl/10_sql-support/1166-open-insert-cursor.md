---
title: "OPEN (insert cursor)"
source: "fgl-topics/c_fgl_InsertCursors_OPEN.html"
breadcrumb: "SQL support > SQL insert cursors > OPEN (insert cursor)"
type: "concept"
---

# OPEN (insert cursor)

> Initializes an insert cursor.

## Syntax

```
OPEN cid
```

1. cid is the identifier of the insert cursor.

## Usage

The `OPEN` statement initializes the insert cursor if the specified cursor was
created with a [`DECLARE …
INSERT`](1165-declare-insert-cursor.md "The DECLARE with an INSERT instruction defines an insert cursor.") statement.

Once the insert cursor is opened, you can add rows with the [`PUT`](1167-put-insert-cursor.md "Adds a new row to the insert cursor buffer.") statement.

When
used with an insert cursor, the `OPEN` instruction
cannot include a `USING` clause.

A subsequent `OPEN` statement
closes the cursor and then reopens it.

If the insert cursor was not declared `WITH HOLD` option, the`OPEN`
instruction generates an SQL error if there is no current transaction started.

If you release cursor resources with a [`FREE`](1170-free-insert-cursor.md "Releases resources allocated for an insert cursor.") instruction, you cannot use the cursor unless you declare the cursor
again.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
