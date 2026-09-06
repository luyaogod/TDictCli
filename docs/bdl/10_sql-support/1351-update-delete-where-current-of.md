---
title: "UPDATE/DELETE … WHERE CURRENT OF"
source: "fgl-topics/c_fgl_odiagmys_028.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > BDL programming > UPDATE/DELETE … WHERE CURRENT OF"
type: "concept"
description: "Informix® Informix allows positioned UPDATEs and DELETEs with the \"WHERE CURRENT OF cursor \" clause, if the cursor has been DECLARED with a SELECT ... FOR UPDATE statement. Oracle® MySQL and MariaDB ..."
---

# UPDATE/DELETE … WHERE CURRENT OF

## Informix®

Informix allows positioned UPDATEs and DELETEs with
the "WHERE CURRENT OF cursor" clause, if the cursor has been DECLARED with a
SELECT ... FOR UPDATE statement.

## Oracle® MySQL and MariaDB

MySQL and MariaDB do not support `UPDATE/DELETE` with the `WHERE CURRENT OF
cursor` clause.

## Solution

`UPDATE/DELETE ... WHERE CURRENT OF` is not supported with MySQL/MariaDB.

As a replacement of `WHERE CURRENT OF`, if the database
table is defined with a primary key column, use the value fetched from the
`SELECT [FOR UPDATE]` cursor in the `WHERE` clause of the `UPDATE/DELETE`
statement.

## Related links

**Related concepts**  

[Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.")
