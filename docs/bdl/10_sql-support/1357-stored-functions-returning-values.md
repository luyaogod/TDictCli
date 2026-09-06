---
title: "Stored functions returning values"
source: "fgl-topics/c_fgl_sql_programming_037.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > BDL programming > Stored procedure calls > Stored functions returning values"
type: "concept"
description: "The following example shows how to retrieve the return value of a stored function with Oracle® MySQL: MySQL version 5.0 does not allow you to prepare the CREATE FUNCTION statement; you may need to ..."
---

# Stored functions returning values

The following example shows how to retrieve the return value of
a stored function with Oracle® MySQL:

MySQL version 5.0 does not allow you to prepare the `CREATE
FUNCTION` statement; you may need to execute this statement
from the mysql command line tool.

```
MAIN
   DEFINE n INTEGER
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE "create function func1(p1 integer)"
                || " no sql begin"
                || "    return concat( 'Value = ', p1 );"
                || "  end;"
   PREPARE stmt FROM "select func1(?)"
   LET n = 111
   EXECUTE stmt USING n INTO c
   DISPLAY c
END MAIN
```
