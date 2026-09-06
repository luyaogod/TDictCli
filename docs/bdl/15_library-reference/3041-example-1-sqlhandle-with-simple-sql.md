---
title: "Example 1: SqlHandle with simple SQL"
source: "fgl-topics/c_fgl_ClassSqlHandle_example_1.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > Examples > Example 1: SqlHandle with simple SQL"
type: "concept"
description: "The following code executes a simple UPDATE statement with the base.SqlHandle API: MAIN DEFINE h base.SqlHandle DEFINE pkey INTEGER CONNECT TO \":memory:+driver='dbmsqt'\" CALL create_table() LET h = ..."
---

# Example 1: SqlHandle with simple SQL

The following code executes a simple UPDATE statement with the base.SqlHandle
API:

```
MAIN
    DEFINE h base.SqlHandle
    DEFINE pkey INTEGER
    CONNECT TO ":memory:+driver='dbmsqt'"
    CALL create_table()
    LET h = base.SqlHandle.create()
    CALL h.prepare("UPDATE t1 SET crea = ? WHERE pkey = ?")
    CALL h.setParameterType(1,"DATE")
    CALL h.setParameter(1,TODAY)
    CALL h.setParameter(2,pkey)
    TRY
        CALL h.execute()
    CATCH
        DISPLAY "SQL Error:", sqlca.sqlcode
    END TRY
END MAIN

FUNCTION create_table()
    DEFINE v_date DATE
    CREATE TABLE t1 ( pkey INTEGER PRIMARY KEY,
                      name VARCHAR(50),
                      crea DATE )
    LET v_date = TODAY
    INSERT INTO t1 VALUES ( 101, 'aaaaa', v_date )
    INSERT INTO t1 VALUES ( 102, 'bbbbbbbb', v_date )
END FUNCTION
```
