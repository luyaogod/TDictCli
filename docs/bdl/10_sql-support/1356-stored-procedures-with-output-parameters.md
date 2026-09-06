---
title: "Stored procedures with output parameters"
source: "fgl-topics/c_fgl_sql_programming_036.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > BDL programming > Stored procedure calls > Stored procedures with output parameters"
type: "concept"
description: "Since Oracle® MySQL C API (version 5.0) does not support an output parameter specification, the IN / OUT / INOUT technique cannot be used. In order to return values from a MySQL stored procedure or ..."
---

# Stored procedures with output parameters

Since Oracle® MySQL C API (version 5.0)
does not support an output parameter specification, the `IN`/`OUT`/`INOUT`
technique cannot be used.

In order to return values from a MySQL stored procedure or stored function, you must use SQL variables.
There are three steps to execute the procedure or function:

1. With the `SET` SQL statement, create and assign an SQL variables for each parameter.
2. `CALL` the stored procedure or stored function with the created SQL variables.
3. Perform a `SELECT` statement to return the SQL variables to the application.

In order to retrieve returning values into program variables, you must use an `INTO` clause
in the `EXECUTE` instruction.

The following example shows how to call a stored procedure with output parameters:

MySQL version 5.0 does not allow you to prepare the `CREATE PROCEDURE` statement; you may need
to execute this statement from the mysql command line tool.

MySQL version 5.0 cannot execute "`SELECT @variable`" with server-side cursors. Since the MySQL
driver uses server-side cursors to support multiple active result sets, it is not possible to execute the
`SELECT` statement to return output parameter values.

MySQL version >=5.0 evaluates "`@variable`" user variables assigned with a string as large
`text` (CLOB) expressions. That type of values must normally be fetched into `TEXT`
variable. To workaround this behavior, you can use the `substring(@var,1,255)` function to return
a `VARCHAR()` expression from MySQL and fetch into a `VARCHAR()` variable.

```
MAIN
   DEFINE n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE 
                   "create procedure proc1("
                || "   p1 integer,"
                || "   out p2 numeric(6,2),"
                || "   out p3 varchar(200)"
                || " )"
                || " no sql begin"
                || "    set p2 = p1 + 0.23;"
                || "    set p3 = concat( 'Value = ', p1 );"
                || "  end;"
   LET n = 111
   EXECUTE IMMEDIATE SFMT("set @p1 = %1", n)
   EXECUTE IMMEDIATE "set @p2 = NULL"
   EXECUTE IMMEDIATE "set @p3 = NULL"
   EXECUTE IMMEDIATE "call proc1(@p1, @p2, @p3)"
   PREPARE stmt FROM "select @p2, substring(@p3,1,200)"
   EXECUTE stmt INTO d, c
   DISPLAY d
   DISPLAY c
END MAIN
```
