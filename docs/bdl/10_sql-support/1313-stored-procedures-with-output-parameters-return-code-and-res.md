---
title: "Stored procedures with output parameters, return code and result set"
source: "fgl-topics/c_fgl_sql_programming_030.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > Stored procedure calls > Stored procedures with output parameters, return code and result set"
type: "concept"
description: "With SQL Server you can call stored procedures with a return code, output parameters and producing a result set. Return codes and output parameters are the last items returned to the application by ..."
---

# Stored procedures with output parameters, return code and result set

With SQL Server you can call stored procedures with a return code,
output parameters and producing a result set.

Return codes and output parameters are the last items returned
to the application by SQL Server; they are not returned until the
last row of the result set has been fetched, after the `SQLMoreResults()` ODBC
function is called. If output parameters are used, the SQL Server
driver executes a `SQLMoreResults()` call when closing the
cursor instead of `SQLCloseCursor()`, to get the
return code and output parameter values from SQL Server.

```
MAIN
   DEFINE r, i, n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   CREATE TABLE tab1 ( c1 INTEGER, c2 DECIMAL(6,2), c3 VARCHAR(200) )
   INSERT INTO tab1 VALUES ( 1, 123.45, 'aaaaaa' )
   INSERT INTO tab1 VALUES ( 2, 123.66, 'bbbbbbbbb' )
   INSERT INTO tab1 VALUES ( 3, 444.77, 'cccccc' )
   EXECUTE IMMEDIATE "create procedure proc3 @key integer output"
                || "  as begin"
                || "     set @key = @key - 1"
                || "     select * from tab1 where c1 > @key"
                || "     return (@key * 3)"
                || "  end"
   DECLARE curs CURSOR FROM "{ ? = call proc3(?) }"
   LET i = 1
   OPEN curs USING r INOUT, i INOUT
   DISPLAY r, i
   FETCH curs INTO n, d, c
   FETCH curs INTO n, d, c
   FETCH curs INTO n, d, c
   DISPLAY r, i
   CLOSE curs
   DISPLAY r, i -- Now the returned values are available
END MAIN
```

The return code and output parameter variables must be defined
as `INOUT` in the `OPEN` instruction.
