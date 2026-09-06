---
title: "Example 3: SqlHandle with insert cursor"
source: "fgl-topics/c_fgl_ClassSqlHandle_example_3.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > Examples > Example 3: SqlHandle with insert cursor"
type: "concept"
description: "The following code implements an insert cursor with the base.SqlHandle API: MAIN DEFINE h base.SqlHandle DEFINE i INTEGER CONNECT TO \":memory:+driver='dbmsqt'\" CREATE TABLE t1 ( pkey INTEGER, name ..."
---

# Example 3: SqlHandle with insert cursor

The following code implements an insert cursor with the `base.SqlHandle`
API:

```
MAIN
    DEFINE h base.SqlHandle
    DEFINE i INTEGER
    CONNECT TO ":memory:+driver='dbmsqt'"
    CREATE TABLE t1 ( pkey INTEGER, name VARCHAR(50))
    LET h = base.SqlHandle.create()
    CALL h.prepare("INSERT INTO t1 VALUES (?, ?)")
    BEGIN WORK
    CALL h.open()
    FOR i=1 TO 10
        CALL h.setParameter(1, i)
        CALL h.setParameter(2, SFMT("item_%1",i))
        CALL h.put()
        IF i MOD 100 == 0 THEN
           CALL h.flush()
        END IF
    END FOR
    CALL h.close()
    COMMIT WORK
    SELECT COUNT(*) INTO i FROM t1
    DISPLAY i
END MAIN
```
