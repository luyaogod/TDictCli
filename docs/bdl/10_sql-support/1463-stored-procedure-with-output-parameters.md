---
title: "Stored procedure with output parameters"
source: "fgl-topics/c_fgl_sql_programming_032.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > Stored procedure calls > Stored procedure with output parameters"
type: "concept"
description: "To execute a stored procedure with PostgreSQL, you must use CALL procname , as shown in this line: PREPARE stmt FROM \"call proc1(?,?,?)\" In order to retrieve returning values into program variables, ..."
---

# Stored procedure with output parameters

To execute a stored procedure with PostgreSQL, you must use `CALL procname`, as
shown in this line:

```
PREPARE stmt FROM "call proc1(?,?,?)"
```

In order to retrieve returning values into program variables, you
must use an `INTO` clause in the `EXECUTE` instruction.

The following example shows how to call a stored procedure with
PostgreSQL:

```
MAIN
   DEFINE n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE
                   "create procedure proc1("
                || "          p1 integer,"
                || "          inout p2 numeric(6,2),"
                || "          inout p3 varchar(200)"
                || "          )"
                || " as $$"
                || "  begin"
                || "    p2:= p1 + 0.23;"
                || "    p3:= 'Value = ' || cast(p1 as text);"
                || "  end;"
                || " $$ language plpgsql"
   PREPARE stmt FROM "call proc1(?,?,?)"
   LET n = 111
   EXECUTE stmt USING n, d, c INTO d, c
   DISPLAY d
   DISPLAY c
END MAIN
```
