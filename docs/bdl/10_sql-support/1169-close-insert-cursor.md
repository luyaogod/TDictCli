---
title: "CLOSE (insert cursor)"
source: "fgl-topics/c_fgl_InsertCursors_CLOSE.html"
breadcrumb: "SQL support > SQL insert cursors > CLOSE (insert cursor)"
type: "concept"
---

# CLOSE (insert cursor)

> Flushes and closes an insert cursor.

## Syntax

```
CLOSE cid
```

1. cid is the identifier of the insert cursor.

## Usage

Closing the insert cursor flushes automatically the rows remaining in the insert buffer, and
releases the resources allocated for the insert buffer on the database server.

After using the `CLOSE` instruction, you must reopen the cursor with [`OPEN`](1166-open-insert-cursor.md "Initializes an insert cursor.") before adding new rows with [`PUT`](1167-put-insert-cursor.md "Adds a new row to the insert cursor buffer.")/[`FLUSH`](1168-flush-insert-cursor.md "Flushes the buffer of an insert cursor.").
