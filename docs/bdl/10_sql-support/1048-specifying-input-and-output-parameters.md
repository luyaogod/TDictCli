---
title: "Specifying input and output parameters"
source: "fgl-topics/c_fgl_sql_programming_007.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Stored procedures > Specifying input and output parameters"
type: "concept"
description: "Input and output parameters can be specified in SQL statement execution to pass and return values to/from stored procedures, depending on the database type: EXECUTE stmt USING param1 IN, param2 INOUT, ..."
---

# Specifying input and output parameters

Input and output parameters can be specified in SQL statement execution to pass and return values
to/from stored procedures, depending on the database type:

```
EXECUTE stmt USING param1 IN, param2 INOUT, param3 INOUT
```

## Related links

**Related concepts**  

[EXECUTE (SQL statement)](1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.")
