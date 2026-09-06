---
title: "Stored functions returning a result set"
source: "fgl-topics/c_fgl_sql_programming_ifx_proc_resultset.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Fully supported IBM® Informix® SQL features > Stored procedure calls > Stored functions returning a result set"
type: "concept"
description: "To retrieve the rows of a result set produced by an IBM® Informix® stored function, you must create a cursor, as you would for a regular SELECT statement. This example shows how to execute a stored ..."
---

# Stored functions returning a result set

To retrieve the rows of a result set produced by an IBM®
Informix® stored function, you must create a cursor, as
you would for a regular `SELECT` statement.

This example shows how to execute a stored function producing a result set:

```
MAIN
    DEFINE m, p_pk INT, p_name VARCHAR(10)
    DATABASE test1
    CREATE TABLE t1 ( pk INT, name VARCHAR(10) )
    INSERT INTO t1 VALUES (1, 'aaaa')
    INSERT INTO t1 VALUES (2, 'bbbbbb')
    INSERT INTO t1 VALUES (3, 'cccc')
    EXECUTE IMMEDIATE "create function proc3(v_max INT)"
                    ||" returning int, lvarchar;"
                    ||" define r_pk integer;"
                    ||" define r_name lvarchar;"
                    ||" foreach c1 for select pk,name into r_pk, r_name from t1 where pk <= v_max"
                    ||" return r_pk,r_name with resume;"
                    ||" end foreach;"
                    ||" end function"
    DECLARE c CURSOR FROM "EXECUTE FUNCTION proc3(?)"
    LET m = 100
    FOREACH c USING m INTO p_pk, p_name
        DISPLAY p_pk, p_name
    END FOREACH
    EXECUTE IMMEDIATE "drop function proc3"
    DROP TABLE t1
END MAIN
```
