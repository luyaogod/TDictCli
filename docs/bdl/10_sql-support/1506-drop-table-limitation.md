---
title: "Drop table limitation"
source: "fgl-topics/c_fgl_odiagsqt_040.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > Drop table limitation"
type: "concept"
description: "SQLite As of version 3.47, SQLite does not allow to drop an SQL table, when there are open SQL cursors, even if the cursors are not using the table to be dropped. In such case, SQLite will return the ..."
---

# Drop table limitation

## SQLite

As of version 3.47, SQLite does not allow to drop an SQL table, when there are open SQL cursors,
even if the cursors are not using the table to be dropped.

In such case, SQLite will return the error `6 "database table is locked"`.

For example, the following code will fail with SQLite:

```
MAIN
    CONNECT TO ":memory:+driver='dbmsqt'"
    CREATE TABLE tab1 ( pk INT, name VARCHAR(50) )
    INSERT INTO tab1 VALUES ( 101, 'aaa' )
    INSERT INTO tab1 VALUES ( 102, 'bbbb' )
    DECLARE c1 CURSOR FOR SELECT * FROM tab1
    OPEN c1
    FETCH c1
    CREATE TEMP TABLE tt1 ( pk INT, name VARCHAR(50) )
    DROP TABLE tt1    -- Fails
    CLOSE c1
END MAIN
```

Note that SQLite does not complain when creating the table (`CREATE TABLE`).

This SQLite limitation prevents the execution of [reports with aggregate functions](../12_reports/2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically."), since the report engine must create and drop temporary
tables while the report driver cursor is still fetching rows from the database with an active SQL
cursor.

## Solution

As a general pattern, Data Definition Language statements (`CREATE TABLE` or
`DROP TABLE`) should be executed without an active SQL cursor.

With reports, compute the aggregrates in the report driver and pass the values as parameters to
the report routine.

This known SQLite limitation may be lifted in a future version. Whatch the enhancements in latest
SQLite releases.
