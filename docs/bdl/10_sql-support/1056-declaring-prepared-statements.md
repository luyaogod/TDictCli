---
title: "Declaring prepared statements"
source: "fgl-topics/c_fgl_sql_programming_054.html"
breadcrumb: "SQL support > SQL programming > SQL performance > Declaring prepared statements"
type: "concept"
---

# Declaring prepared statements

> Optimize prepared cursor statements by using the FROM clause of DECLARE CURSOR.

Line 2 of this example shows a cursor declared with a prepared
statement:

```
PREPARE s FROM "SELECT * FROM table WHERE ", condition
DECLARE c CURSOR FOR s
```

While this has no performance impact with IBM® Informix® database
drivers, it can become a bottleneck when using non-IBM Informix database servers:

Statement preparation consumes a lot of memory and processor resources.
Declaring a cursor with a prepared statement is a native IBM Informix feature, which consumes only one
real statement preparation. Non-IBM Informix database servers do not support this
feature, so the statement is prepared twice (once for the `PREPARE`,
and once for the `DECLARE`). When used in a big loop,
this code can cause performance problems.

To optimize the code, use the `FROM` clause in the `DECLARE`
statement:

```
DECLARE c CURSOR FROM "SELECT * FROM table WHERE " || condition
```

By using this solution only one statement preparation will be done
by the database server.

This performance problem does not occur with `DECLARE` statements using static
SQL.

## Related links

**Related concepts**  

[DECLARE (result set cursor)](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.")
