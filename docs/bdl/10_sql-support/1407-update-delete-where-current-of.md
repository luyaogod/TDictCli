---
title: "UPDATE/DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_odiagora_038.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > UPDATE/DELETE … WHERE CURRENT OF"
type: "concept"
description: "Informix® Informix allows positioned UPDATEs and DELETEs with the \"WHERE CURRENT OF cursor \" clause, if the cursor has been DECLARED with a SELECT ... FOR UPDATE statement. ORACLE The UPDATE/DELETE ..."
---

# UPDATE/DELETE … WHERE CURRENT OF

## Informix®

Informix allows positioned UPDATEs and DELETEs with
the "WHERE CURRENT OF cursor" clause, if the cursor has been DECLARED with a
SELECT ... FOR UPDATE statement.

## ORACLE

The `UPDATE/DELETE ... WHERE CURRENT OF cursor` statements are
not support by the Oracle database API.

However, ROWIDs can be used for positioned updates and deletes.

## Solution

`UPDATE/DELETE ... WHERE CURRENT OF` instructions
are emulated by the ORACLE database interface by using ROWIDs:

The ORACLE database interface
replaces `WHERE CURRENT OF cursor` by `WHERE
ROWID=:rid` and sets the value of the ROWID returned by the last `FETCH` done
with the given cursor.

As a replacement of `WHERE CURRENT OF`, if the database
table is defined with a primary key column, use the value fetched from the
`SELECT [FOR UPDATE]` cursor in the `WHERE` clause of the `UPDATE/DELETE`
statement.

## Related links

**Related concepts**  

[Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.")
