---
title: "Stored function with output parameters"
source: "fgl-topics/c_fgl_sql_programming_034.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > Stored procedure calls > Stored function with output parameters"
type: "concept"
description: "To execute a stored function with PostgreSQL that returns several values, you must use SELECT * FROM function , as shown in this line: PREPARE stmt FROM \"select * from proc1(?)\" In order to retrieve ..."
---

# Stored function with output parameters

To execute a stored function with PostgreSQL that returns several values, you must use
`SELECT * FROM function`, as shown in this
line:

```
PREPARE stmt FROM "select * from proc1(?)"
```

In order to retrieve returning values into program variables, you must use an
`INTO` clause in the `EXECUTE` instruction.

The following example shows how to call a stored function with
PostgreSQL:

```
MAIN
   DEFINE n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE 
                   "create function func1("
                || "          p1 integer,"
                || "          out p2 numeric(6,2),"
                || "          out p3 varchar(200)"
                || "          )"
                || " as $$"
                || "  begin"
                || "    p2:= p1 + 0.23;"
                || "    p3:= 'Value = ' || cast(p1 as text);"
                || "  end;"
                || " $$ language plpgsql"
   PREPARE stmt FROM "select * from func1(?)"
   LET n = 111
   EXECUTE stmt USING n INTO d, c
   DISPLAY d
   DISPLAY c
END MAIN
```
