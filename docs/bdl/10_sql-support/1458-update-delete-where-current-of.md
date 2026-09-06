---
title: "UPDATE/DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_odiagpgs_034.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > UPDATE/DELETE … WHERE CURRENT OF"
type: "concept"
description: "Informix® Informix allows positioned UPDATEs and DELETEs with the \"WHERE CURRENT OF cursor \" clause, if the cursor has been DECLARED with a SELECT ... FOR UPDATE statement. PostgreSQL UPDATE/DELETE ..."
---

# UPDATE/DELETE … WHERE CURRENT OF

## Informix®

Informix allows positioned UPDATEs and DELETEs with
the "WHERE CURRENT OF cursor" clause, if the cursor has been DECLARED with a
SELECT ... FOR UPDATE statement.

## PostgreSQL

`UPDATE/DELETE ... WHERE CURRENT OF` is supported by PostgreSQL with server-side
cursors created by a `DECLARE` statement.

## Solution

With PostgreSQL, `UPDATE/DELETE ... WHERE CURRENT OF` instructions are executed as
is without any SQL translation: Since [`SELECT FOR
UPDATE`](1457-select-for-update.md) statements are implemented with a `DECLARE` statement to get a
server cursor, native positioned update/delete can take place in PostgreSQL.

As a replacement of `WHERE CURRENT OF`, if the database
table is defined with a primary key column, use the value fetched from the
`SELECT [FOR UPDATE]` cursor in the `WHERE` clause of the `UPDATE/DELETE`
statement.

## Related links

**Related concepts**  

[Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.")
