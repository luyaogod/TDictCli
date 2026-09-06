---
title: "base.SqlHandle.fetch"
source: "fgl-topics/c_fgl_ClassSqlHandle_fetch.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.fetch"
type: "concept"
---

# base.SqlHandle.fetch

> Fetches a new row from the SQL result set.

## Syntax

```
fetch()
```

## Usage

Call the `fetch()` method to fetch a new row from the SQL result set.

When using a [dynamic scroll
cursor](3033-base-sqlhandle-openscrollcursor.md "Opens the SQL handle (with scrollable option)."), the `fetch()` method can be used to fetch to the next row.

The SQL statement must have been opened with one of the open cursor methods.

After performing the fetch call, you can query for column information
with the [`getResultCount()`](3027-base-sqlhandle-getresultcount.md "Returns the number of result set columns produced by the SQL statement."), [`getResultName(index)`](3028-base-sqlhandle-getresultname.md "Returns the name of a column in the result set produced by the SQL statement."), [`getResultType(index)`](3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.") and [`getResultValue(index)`](3030-base-sqlhandle-getresultvalue.md "Returns the value of a column in the result set produced by the SQL statement.") methods.

If no row is found (end of result set), `sqlca.sqlcode` is set to 100
(`NOTFOUND`).

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.fetch()
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
