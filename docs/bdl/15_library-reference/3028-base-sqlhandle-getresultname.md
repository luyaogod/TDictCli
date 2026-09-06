---
title: "base.SqlHandle.getResultName"
source: "fgl-topics/c_fgl_ClassSqlHandle_getResultName.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.getResultName"
type: "concept"
---

# base.SqlHandle.getResultName

> Returns the name of a column in the result set produced by the SQL statement.

## Syntax

```
getResultName( index INTEGER )
 RETURNS STRING
```

1. index is the ordinal position of the result set column (starts at 1).

## Usage

Call the `getResultName()` method to query the name of a column in the result set,
after executing the SQL statement with the [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).") method and fetching a row with [`fetch()`](3020-base-sqlhandle-fetch.md "Fetches a new row from the SQL result set.").

The method takes the position of the column as the parameter.

## Example

```
DEFINE sh base.SqlHandle, i INT
...
FOR i=1 TO sh.getResultCount()
    DISPLAY sh.getResultName(i)
END FOR
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
