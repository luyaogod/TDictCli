---
title: "INSERT cursors"
source: "fgl-topics/c_fgl_odiagsqt_020.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > INSERT cursors"
type: "concept"
description: "Informix® Informix provides insert cursors to optimize row creation in a database. An insert cursor is declared as a cursor, and rows as added with the PUT instruction. The rows are buffered and sent ..."
---

# INSERT cursors

## Informix®

Informix provides insert cursors to
optimize row creation in a database. An insert cursor is declared as a cursor, and rows as added
with the `PUT` instruction. The rows are buffered and sent to the database server
when executing a `FLUSH` instruction, or when the cursor is closed with
`CLOSE`. When using transactions in Informix, the `OPEN`, `PUT` and `FLUSH`
instructions must be executed within a transaction block.

```
DECLARE c1 CURSOR FOR INSERT INTO tab1 ...
BEGIN WORK
OPEN c1
WHILE ...
   PUT c1 USING var-list
END WHILE
CLOSE c1
COMMIT WORK
```

## SQLite

SQLite does not support insert cursors.

## Solution

Insert cursors are emulated by the database interface, using basic `INSERT` SQL
instructions.

The performances might be not as good as with Informix, but the feature is fully supported.

## Related links

**Related concepts**  

[Insert cursors](1030-insert-cursors.md "Using insert cursors with non-Informix databases.")
