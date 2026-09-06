---
title: "Stored functions with result set"
source: "fgl-topics/c_fgl_sql_programming_033.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > Stored procedure calls > Stored functions with result set"
type: "concept"
description: "With PostgreSQL, you can execute stored function returning a result set. To do so, you must declare a cursor and fetch the rows: MAIN DEFINE i, n INTEGER DEFINE d DECIMAL(6,2) DEFINE c VARCHAR(200) ..."
---

# Stored functions with result set

With PostgreSQL, you can execute stored function returning a result set. To do so, you must
declare a cursor and fetch the
rows:

```
MAIN
   DEFINE i, n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   CREATE TABLE tab1 ( c1 INTEGER, c2 DECIMAL(6,2), c3 VARCHAR(200) )
   INSERT INTO tab1 VALUES ( 1, 123.45, 'aaaaaa' )
   INSERT INTO tab1 VALUES ( 2, 123.66, 'bbbbbbbbb' )
   INSERT INTO tab1 VALUES ( 3, 444.77, 'cccccc' )
   EXECUTE IMMEDIATE "create function func2(integer)"
                || " returns setof tab1"
                || " as $$"
                || "  select * from tab1 where c1 > $1;"
                || " $$ language sql"
   DECLARE curs CURSOR FROM "select * from func2(?)"
   LET i = 1
   FOREACH curs USING i INTO n, d, c
       DISPLAY n, d, c
   END FOREACH
END MAIN
```
