---
title: "Example 2: SqlHandle with result set SQL"
source: "fgl-topics/c_fgl_ClassSqlHandle_example_2.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > Examples > Example 2: SqlHandle with result set SQL"
type: "concept"
description: "The following code executes a simple SELECT statement with the base.SqlHandle API: MAIN DEFINE h base.SqlHandle, v VARCHAR(50), i INTEGER CONNECT TO \":memory:+driver='dbmsqt'\" CALL create_table() LET ..."
---

# Example 2: SqlHandle with result set SQL

The following code executes a simple SELECT statement with the base.SqlHandle
API:

```
MAIN
    DEFINE h base.SqlHandle,
           v VARCHAR(50),
           i INTEGER
    CONNECT TO ":memory:+driver='dbmsqt'"
    CALL create_table()
    LET h = base.SqlHandle.create()
    CALL h.prepare("SELECT * FROM t1 WHERE name>? ORDER BY pkey")

    LET v = "a"
    CALL h.setParameter(1, v)
    CALL h.open()
    WHILE TRUE
        CALL h.fetch()
        IF sqlca.sqlcode==NOTFOUND THEN
           EXIT WHILE
        END IF
        DISPLAY "-----------------"
        FOR i=1 TO h.getResultCount()
            DISPLAY i, ":", h.getResultName(i),
                    " / ", h.getResultType(i),
                    " = ", h.getResultValue(i)
        END FOR
    END WHILE
    CALL h.close()
END MAIN

FUNCTION create_table()
    CREATE TABLE t1 ( pkey INTEGER PRIMARY KEY,
                      name VARCHAR(50) )
    INSERT INTO t1 VALUES ( 101, 'aaaaa' )
    INSERT INTO t1 VALUES ( 102, 'bbbbbbbb' )
    INSERT INTO t1 VALUES ( 103, 'cccccc' )
    INSERT INTO t1 VALUES ( 104, 'ddddddd' )
END FUNCTION
```
