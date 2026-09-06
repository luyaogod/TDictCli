---
title: "Stored procedures with return code"
source: "fgl-topics/c_fgl_sql_programming_029.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > Stored procedure calls > Stored procedures with return code"
type: "concept"
description: "SQL Server stored procedures can return integer values. To get the return value of a stored procedure, you must use an assignment expression in the ODBC call escape sequence: PREPARE stmt FROM \"{ ? = ..."
---

# Stored procedures with return code

SQL Server stored procedures can return integer values. To get the return value of a stored
procedure, you must use an assignment expression in the ODBC `call` escape
sequence:

```
PREPARE stmt FROM "{ ? = call proc3(?,?,?) }"
```

Then the statement can be executed with the `EXECUTE` instruction, by specifying
the output parameter receiving the returned value as first element in the `USING`
list, with the `OUT`
modifier:

```
MAIN
    DEFINE p1, p2, r INTEGER
    CONNECT TO ...
    WHENEVER ERROR CONTINUE
    EXECUTE IMMEDIATE "DROP PROCEDURE p_mul"
    WHENEVER ERROR STOP
    EXECUTE IMMEDIATE "CREATE PROCEDURE p_mul ( @p1 INT, @p2 INT )
     AS BEGIN
     DECLARE @r INT
     SET @r = @p1 * @p2;
     RETURN @r;
     END"
    LET p1 = 15
    LET p2 = 3
    PREPARE s1 FROM "{ ? = call p_mul(?,?) }"
    EXECUTE s1 USING r OUT, p1 IN, p2 IN
    DISPLAY "r = ", r
END MAIN
```
