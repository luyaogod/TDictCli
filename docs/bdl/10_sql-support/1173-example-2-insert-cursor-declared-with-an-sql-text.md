---
title: "Example 2: Insert Cursor declared with an SQL text"
source: "fgl-topics/c_fgl_InsertCursors_Example_002.html"
breadcrumb: "SQL support > SQL insert cursors > Examples > Example 2: Insert Cursor declared with an SQL text"
type: "concept"
description: "MAIN DEFINE i INTEGER DEFINE rec RECORD key INTEGER, name CHAR(30) END RECORD DATABASE stock DECLARE ic CURSOR FROM \"INSERT INTO item VALUES (?,?)\" BEGIN WORK OPEN ic FOR i=1 TO 100 LET rec.key = i ..."
---

# Example 2: Insert Cursor declared with an SQL text

```
MAIN
   DEFINE i INTEGER
   DEFINE rec RECORD
           key INTEGER,
           name CHAR(30)
         END RECORD
   DATABASE stock
   DECLARE ic CURSOR FROM "INSERT INTO item VALUES (?,?)"
   BEGIN WORK
     OPEN ic
     FOR i=1 TO 100
         LET rec.key = i
         LET rec.name = "Item #" || i
         PUT ic FROM rec.*
         IF i MOD 50 = 0 THEN
             FLUSH ic
         END IF
     END FOR
     CLOSE ic
   COMMIT WORK
   FREE ic
END MAIN
```
