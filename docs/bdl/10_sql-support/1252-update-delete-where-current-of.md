---
title: "UPDATE/DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_odiagdmg_018.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > BDL programming > UPDATE/DELETE … WHERE CURRENT OF"
type: "concept"
description: "Informix® Informix allows positioned UPDATEs and DELETEs with the \"WHERE CURRENT OF cursor \" clause, if the cursor has been DECLARED with a SELECT ... FOR UPDATE statement. Dameng® UPDATE/DELETE ... ..."
---

# UPDATE/DELETE … WHERE CURRENT OF

## Informix®

Informix allows positioned UPDATEs and DELETEs with
the "WHERE CURRENT OF cursor" clause, if the cursor has been DECLARED with a
SELECT ... FOR UPDATE statement.

## Dameng®

`UPDATE/DELETE ... WHERE CURRENT OF` is supported by the Dameng DPI interface.

## Solution

With Dameng, `UPDATE/DELETE ... WHERE CURRENT OF`
instructions are executed as is without SQL translation.

As a replacement of `WHERE CURRENT OF`, if the database
table is defined with a primary key column, use the value fetched from the
`SELECT [FOR UPDATE]` cursor in the `WHERE` clause of the `UPDATE/DELETE`
statement.

## Related links

**Related concepts**  

[Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.")
