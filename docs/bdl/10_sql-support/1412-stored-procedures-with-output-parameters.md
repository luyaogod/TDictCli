---
title: "Stored procedures with output parameters"
source: "fgl-topics/c_fgl_sql_programming_018.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > Stored procedure calls > Stored procedures with output parameters"
type: "concept"
description: "Oracle® stored procedures or stored functions must be called with the input and output parameters specification in the USING clause of the EXECUTE , OPEN or FOREACH instruction. As in normal dynamic ..."
---

# Stored procedures with output parameters

Oracle® stored procedures or stored
functions must be called with the input and output parameters specification in the
`USING` clause of the `EXECUTE`, `OPEN` or
`FOREACH` instruction. As in normal dynamic SQL, parameters must correspond by
position, and the `IN/OUT/INOUT` options must match the parameter definition of the
stored procedure.

To execute the stored procedure, you must include the procedure
in an anonymous PL/SQL block with `BEGIN` and `END`
keywords:

```
PREPARE stmt FROM "begin proc1(?,?,?); end;"
```

Remark: Oracle stored procedures do not
specify the size of number and character parameters. The size of output values (especially character
strings) are defined by the calling context (i.e. the data type of the variable used when calling
the procedure). When you pass a CHAR(10) to the procedure, the returning value will be filled with
blanks to reach a size of 10 bytes.

Note that for technical reasons, the Oracle driver uses dynamic binding with `OCIBindDynamic()`. The Oracle Call Interface does not support stored
procedures parameters with the `CHAR` data type when using dynamic binding. You must
use `VARCHAR2` instead of `CHAR` to define character string parameters
for stored procedures.

Here is a complete example creating and calling a stored procedure
with output parameters:

```
MAIN
   DEFINE n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE
                   "create procedure proc1("
                || "         p1 in int,"
                || "         p2 in out number,"
                || "         p3 in out varchar2"
                || "      )"
                || " is begin"
                || "  p2:= p1 + 0.23;"
                || "  p3:= 'Value = ' || to_char(p1);"
                || "end;"
   PREPARE stmt FROM "begin proc1(?,?,?); end;"
   LET n = 111
   EXECUTE stmt USING n IN, d INOUT, c INOUT
   DISPLAY d
   DISPLAY c
END MAIN
```
