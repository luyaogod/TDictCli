---
title: "FLUSH (insert cursor)"
source: "fgl-topics/c_fgl_InsertCursors_FLUSH.html"
breadcrumb: "SQL support > SQL insert cursors > FLUSH (insert cursor)"
type: "concept"
---

# FLUSH (insert cursor)

> Flushes the buffer of an insert cursor.

## Syntax

```
FLUSH cid
```

1. cid is the identifier of the insert cursor.

## Usage

When flushing an insert cursor, all buffered rows are inserted into the target database table and
the insert buffer is cleared.

The insert buffer may be automatically flushed by the runtime system if there no room when a
new row is added with the [`PUT`](1167-put-insert-cursor.md "Adds a new row to the insert cursor buffer.")
instruction.

## Related links

**Related concepts**  

[CLOSE (insert cursor)](1169-close-insert-cursor.md "Flushes and closes an insert cursor.")
