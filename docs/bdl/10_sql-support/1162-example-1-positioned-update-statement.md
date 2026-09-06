---
title: "Example 1: Positioned UPDATE statement"
source: "fgl-topics/c_fgl_positioned_updates_example_1.html"
breadcrumb: "SQL support > Positioned updates/deletes > Examples > Example 1: Positioned UPDATE statement"
type: "concept"
description: "MAIN DEFINE pname CHAR(30) DATABASE stock DECLARE uc CURSOR FOR SELECT name FROM item WHERE key=123 FOR UPDATE BEGIN WORK OPEN uc FETCH uc INTO pname IF sqlca.sqlcode=0 THEN LET pname = \"Dummy\" UPDATE ..."
---

# Example 1: Positioned UPDATE statement

```
MAIN
  DEFINE pname CHAR(30)
  DATABASE stock 
  DECLARE uc CURSOR FOR
    SELECT name FROM item WHERE key=123 FOR UPDATE
  BEGIN WORK
    OPEN uc
    FETCH uc INTO pname
    IF sqlca.sqlcode=0 THEN
      LET pname = "Dummy"
      UPDATE item SET name=pname WHERE CURRENT OF uc
    END IF
    CLOSE uc
  COMMIT WORK
  FREE uc
END MAIN
```
