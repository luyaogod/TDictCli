---
title: "Usage"
source: "fgl-topics/c_fgl_ClassSqlHandle_usage.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > Usage"
type: "concept"
description: "The base.SqlHandle class is a built-in class providing dynamic SQL support with a 3GL API. Compared to regular SQL cursor instructions, the main purpose of the base.SqlHandle class is to provide ..."
---

# Usage

The `base.SqlHandle` class is a built-in class providing dynamic SQL support with
a 3GL API.

Compared to regular SQL cursor instructions, the main purpose of the
`base.SqlHandle` class is to provide column name and SQL data type information with
the `getResultName()` and `getResultType()` methods. It is also
possible to write generic code for parameterized queries with the `setParameter()`
method.

A database connection must exist in order to use SqlHandle objects.

Unlike regular Genero cursors, SQL handle objects are created dynamically, and can be passed as
parameter or returned from
functions:

```
MAIN
   DEFINE h base.SqlHandle
   CONNECT TO "mydb"
   LET h = base.SqlHandle.create()
   CALL my_prepare(h)
   CALL my_execute(h)
END MAIN

FUNCTION my_prepare(h)
   DEFINE h base.SqlHandle
   CALL h.prepare("INSERT INTO cust VALUES ( ?, ?, ? )")
END FUNCTION

FUNCTION my_execute(h)
   DEFINE h base.SqlHandle
   CALL h.execute()
END FUNCTION
```

## Executing a simple SQL statement without a result set

Perform the following steps, to execute a SQL statement without a result set:

1. Define the SQL handle variable as `base.SqlHandle`
2. Create a SQL handle object `base.SqlHandle.create()`
3. `prepare(sql-text)`
4. For each SQL parameter:
   1. `setParameter(index, value)`
5. `execute()` -- test for `sqlca.sqlcode`
6. Repeat from (5), (4), or (3)

## Executing a SQL statement returning a result set

Perform the following steps, to execute a SQL statement with a result set:

1. Define the SQL handle variable as `base.SqlHandle`
2. Create a SQL handle object `base.SqlHandle.create()`
3. `prepare(sql-text)`
4. For each SQL parameter:
   1. `setParameter(index,
      value)`
5. `open()`
6. `fetch()` -- test for `sqlca.sqlcode == 100`
7. `getResultCount()` -- for each column index:
   1. `getResultName(index)`
   2. `getResultType(index)`
   3. `getResultValue(index)`
8. `close()`
9. Repeat from (6), (4), (5), or (3)

## Executing a SQL statement returning a result set, as scrollable cursor

Perform the following steps, to execute a SQL statement with a result set and scroll
forwards and backwards in the rows:

1. Define the SQL handle variable as `base.SqlHandle`
2. Create a SQL handle object `base.SqlHandle.create()`
3. `prepare(sql-text)`
4. For each SQL parameter:
   1. `setParameter(index,
      value)`
5. `openScrollCursor()`
6. `fetch()` (next row), `fetchLast()`,
   `fetchFirst()`, `fetchPrevious()`,
   `fetchRelative(n)` or `fetchAbsolute(n)` --
   test for `sqlca.sqlcode == 100`
7. `getResultCount()` -- for each column index:
   1. `getResultName(index)`
   2. `getResultType(index)`
   3. `getResultValue(index)`
8. `close()`
9. Repeat from (6), (4), (5), or (3)

## Creating rows with an insert cursor

Perform the following steps, to insert many rows with an SQL handle insert cursor:

1. Define the SQL handle variable as `base.SqlHandle`
2. Create an SQL handle object `base.SqlHandle.create()`
3. `prepare(insert-stmt-with-params)`
4. `BEGIN WORK`
5. `open()`
6. For each row to insert:
   1. For each SQL parameter:
      1. `setParameter(index, value)`
   2. `put()`
7. `close()`
8. `COMMIT WORK`
9. Repeat from (4) or (3)

## SQL error handling with SqlHandle

Handling SQL error and status information (such as `NOTFOUND`) can be
done with SqlHandle objects as with regular SQL instruction, by testing the
`sqlca.sqlcode` register, and by using `TRY/CATCH`
blocks or `WHENEVER ERROR`.

```
MAIN
   DEFINE h base.SqlHandle
   CONNECT TO "mydb"
   LET h = base.SqlHandle.create()
   TRY
       CALL h.prepare("SELECT * FROM mytab")
       CALL h.open()
       CALL h.fetch()
       DISPLAY h.getResultValue(1)
       CALL h.close()
   CATCH
       DISPLAY "SQL ERROR:", sqlca.sqlcode
   END TRY
END MAIN
```

## Related links

**Related concepts**  

[Result set processing](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query.")

[Dynamic SQL management](../10_sql-support/1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")

[SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
