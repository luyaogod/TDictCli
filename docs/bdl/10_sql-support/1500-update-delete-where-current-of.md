---
title: "UPDATE/DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_odiagsqt_022.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > UPDATE/DELETE … WHERE CURRENT OF"
type: "concept"
description: "Informix® Informix allows positioned UPDATEs and DELETEs with the \"WHERE CURRENT OF cursor \" clause, if the cursor has been DECLARED with a SELECT ... FOR UPDATE statement. SQLite SQLite does not ..."
---

# UPDATE/DELETE … WHERE CURRENT OF

## Informix®

Informix allows positioned UPDATEs and DELETEs with
the "WHERE CURRENT OF cursor" clause, if the cursor has been DECLARED with a
SELECT ... FOR UPDATE statement.

## SQLite

SQLite does not support the `SELECT ... FOR UPDATE` and `UPDATE/DELETE
WHERE CURRENT OF` SQL instruction.

## Solution

`UPDATE/DELETE ... WHERE CURRENT OF` is not supported with SQLite.

As a replacement of `WHERE CURRENT OF`, if the database
table is defined with a primary key column, use the value fetched from the
`SELECT [FOR UPDATE]` cursor in the `WHERE` clause of the `UPDATE/DELETE`
statement.

## Related links

**Related concepts**  

[Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.")
