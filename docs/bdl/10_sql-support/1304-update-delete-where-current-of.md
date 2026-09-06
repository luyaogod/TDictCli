---
title: "UPDATE/DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_odiagmsv_019.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > UPDATE/DELETE … WHERE CURRENT OF"
type: "concept"
description: "Informix® Informix allows positioned UPDATEs and DELETEs with the \"WHERE CURRENT OF cursor \" clause, if the cursor has been DECLARED with a SELECT ... FOR UPDATE statement. Microsoft™ SQL Server ..."
---

# UPDATE/DELETE … WHERE CURRENT OF

## Informix®

Informix allows positioned UPDATEs and DELETEs with
the "WHERE CURRENT OF cursor" clause, if the cursor has been DECLARED with a
SELECT ... FOR UPDATE statement.

## Microsoft™ SQL Server

`UPDATE/DELETE ... WHERE CURRENT OF` is supported with Microsoft SQL Server by using server side cursors.

## Solution

With Microsoft SQL Server, `UPDATE/DELETE ...
WHERE CURRENT OF` instructions are executed as is without any SQL translation: Since BDL
cursors for the [`SELECT FOR UPDATE`](1303-select-for-update.md)
statements are implemented with ODBC server cursors, native positioned update/delete can take place
in SQL Server.

As a replacement of `WHERE CURRENT OF`, if the database
table is defined with a primary key column, use the value fetched from the
`SELECT [FOR UPDATE]` cursor in the `WHERE` clause of the `UPDATE/DELETE`
statement.

## Related links

**Related concepts**  

[Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.")
