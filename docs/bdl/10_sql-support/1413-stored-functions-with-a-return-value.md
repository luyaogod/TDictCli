---
title: "Stored functions with a return value"
source: "fgl-topics/c_fgl_sql_programming_019.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > Stored procedure calls > Stored functions with a return value"
type: "concept"
description: "To execute the stored function returning a value, you must include the function in an anonymous PL/SQL block with BEGIN and END keywords, and use an assignment expression to specify the place holder ..."
---

# Stored functions with a return value

To execute the stored function returning a value, you must include the function in an anonymous
PL/SQL block with `BEGIN` and `END` keywords, and use an assignment
expression to specify the place holder for the returning value:

```
PREPARE stmt FROM "begin ?:= func1(?,?,?); end;"
```
