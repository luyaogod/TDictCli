---
title: "Example 3: Insert Cursor declared with 'hold' option"
source: "fgl-topics/c_fgl_InsertCursors_Example_003.html"
breadcrumb: "SQL support > SQL insert cursors > Examples > Example 3: Insert Cursor declared with 'hold' option"
type: "concept"
description: "MAIN DEFINE name CHAR(30) DATABASE stock DECLARE ic CURSOR WITH HOLD FOR INSERT INTO item VALUES (1,name) OPEN ic LET name = \"Item 1\" PUT ic BEGIN WORK UPDATE refs SET name=\"xyz\" WHERE key=123 COMMIT ..."
---

# Example 3: Insert Cursor declared with 'hold' option

```
MAIN
   DEFINE name CHAR(30)
   DATABASE stock 
   DECLARE ic CURSOR WITH HOLD FOR
     INSERT INTO item VALUES (1,name)
   OPEN ic 
   LET name = "Item 1"
   PUT ic 
   BEGIN WORK
     UPDATE refs SET name="xyz" WHERE key=123
   COMMIT WORK
   PUT ic 
   PUT ic 
   FLUSH ic 
   CLOSE ic 
   FREE ic 
END MAIN
```
