---
title: "Stored procedures with output parameters"
source: "fgl-topics/c_fgl_sql_programming_026.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > Stored procedure calls > Stored procedures with output parameters"
type: "concept"
description: "SQL Server stored procedures must be called with the input and output parameters specification in the USING clause of the EXECUTE , OPEN or FOREACH instruction. As in normal dynamic SQL, parameters ..."
---

# Stored procedures with output parameters

SQL Server stored procedures must be called with the input and
output parameters specification in the `USING` clause
of the `EXECUTE`, `OPEN` or
`FOREACH` instruction. As in normal dynamic SQL,
parameters must correspond by position and the `IN/OUT/INOUT` options
must match the parameter definition of the stored procedure.

To execute the stored procedure, you must use an ODBC `call` escape
sequence:

```
PREPARE stmt FROM "{ call proc1(?,?,?) }"
```

Here is a complete example creating and calling a stored procedure
with output parameters:

```
MAIN
   DEFINE n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE
                   "create procedure proc1"
                || "             @v1 integer,"
                || "             @v2 decimal(6,2) output,"
                || "             @v3 varchar(20) output"
                || " as begin"
                || "  set @v2 = @v1 + 0.23"
                || "  set @v3 = 'Value = ' || cast(@v1 as varchar)"
                || "end"
   PREPARE stmt FROM "{ call proc1(?,?,?) }"
   LET n = 111
   EXECUTE stmt USING n IN, d OUT, c OUT
   DISPLAY d
   DISPLAY c
END MAIN
```
