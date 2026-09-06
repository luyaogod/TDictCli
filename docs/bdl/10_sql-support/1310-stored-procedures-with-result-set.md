---
title: "Stored procedures with result set"
source: "fgl-topics/c_fgl_sql_programming_027.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > Stored procedure calls > Stored procedures with result set"
type: "concept"
description: "With SQL Server, you can execute stored procedures returning a result set. To do so, you must declare a cursor and fetch the rows. The next example uses a stored procedure with a simple SELECT ..."
---

# Stored procedures with result set

With SQL Server, you can execute stored procedures returning a
result set. To do so, you must declare a cursor and fetch the rows.

The next example uses a stored procedure with a simple `SELECT` statement.
If the stored procedure contains additional Transact-SQL
statements such as `SET` or `IF` (which
is the case in complex stored procedures), SQL Server
generates multiple result sets. By default the Genero SQL Server
driver uses "server cursors" to support multiple active
SQL statements. But SQL Server stored procedures generating
multiple result sets cannot be used with server cursors:
The server cursor is silently converted to a "default result set"
cursor by the ODBC driver. Since Default result set cursors
do not support multiple active statements, you cannot
use another SQL statement while processing the results of such stored
procedure. You must `CLOSE` the cursor
created for the stored procedure before continuing with
other SQL statements.

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
   EXECUTE IMMEDIATE "create procedure proc2 @key integer"
                || "  as select * from tab1 where c1 > @key"
   DECLARE curs CURSOR FROM "{ call proc2(?) }"
   LET i = 1
   FOREACH curs USING i INTO n, d, c
       DISPLAY n, d, c
   END FOREACH
END MAIN
```

It is possible to fetch large objects (text/image) from stored
procedure generating a result set. However, if the stored
procedure executes other statements as the `SELECT` (like `SET`/`IF` commands),
the SQL Server ODBC driver will convert the server cursor
to a regular default result set cursor, requiring the
LOB columns to appear at the end of the select list. Thus, in
most cases (stored procedures typically use `SET` / `IF`
statements), you will have to move the LOB columns and
the end of the column list.
