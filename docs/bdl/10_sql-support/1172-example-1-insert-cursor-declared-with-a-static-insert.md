---
title: "Example 1: Insert Cursor declared with a Static INSERT"
source: "fgl-topics/c_fgl_InsertCursors_Example_001.html"
breadcrumb: "SQL support > SQL insert cursors > Examples > Example 1: Insert Cursor declared with a Static INSERT"
type: "concept"
description: "MAIN DEFINE i INTEGER DEFINE rec RECORD key INTEGER, name CHAR(30) END RECORD DATABASE stock DECLARE ic CURSOR FOR INSERT INTO item VALUES (rec.*) BEGIN WORK OPEN ic FOR i=1 TO 100 LET rec.key = i LET ..."
---

# Example 1: Insert Cursor declared with a Static INSERT

```
MAIN
   DEFINE i INTEGER
   DEFINE rec RECORD
           key INTEGER,
           name CHAR(30)
         END RECORD
   DATABASE stock 
   DECLARE ic CURSOR FOR
     INSERT INTO item VALUES (rec.*)
   BEGIN WORK
     OPEN ic 
     FOR i=1 TO 100
         LET rec.key = i 
         LET rec.name = "Item #" || i 
         PUT ic 
         IF i MOD 50 = 0 THEN
             FLUSH ic 
         END IF
     END FOR
     CLOSE ic 
   COMMIT WORK
   FREE ic
END MAIN
```
