---
title: "base.SqlHandle.getResultCount"
source: "fgl-topics/c_fgl_ClassSqlHandle_getResultCount.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.getResultCount"
type: "concept"
---

# base.SqlHandle.getResultCount

> Returns the number of result set columns produced by the SQL statement.

## Syntax

```
getResultCount()
 RETURNS INTEGER
```

## Usage

Call the `getResultCount()` method to query the number of columns in the result
set, after executing the SQL statement with the [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).") method and fetching a row with [`fetch()`](3020-base-sqlhandle-fetch.md "Fetches a new row from the SQL result set.").

## Example

```
DEFINE sh base.SqlHandle, i INT
...
FOR i=1 TO sh.getResultCount()
    DISPLAY sh.getResultName(i)
END FOR
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
