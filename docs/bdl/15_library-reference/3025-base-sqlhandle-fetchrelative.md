---
title: "base.SqlHandle.fetchRelative"
source: "fgl-topics/c_fgl_ClassSqlHandle_fetchRelative.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.fetchRelative"
type: "concept"
---

# base.SqlHandle.fetchRelative

> Fetches a row relative to the current row in a scrollable SQL result set.

## Syntax

```
fetchRelative(
   position INTEGER)
```

1. position is the row offset in the result set. The offset can be negative, to
   fetch backwards.

## Usage

Call the `fetchRelative()` method to fetch the row at the specified offset,
relative to the current row in a scrollable SQL result set.

The SQL statement must have been opened with an [`openScrollCursor()`](3033-base-sqlhandle-openscrollcursor.md "Opens the SQL handle (with scrollable option).") or
[`openScrollCursorWithHold()`](3034-base-sqlhandle-openscrollcursorwithhold.md "Opens the SQL handle with scrollable and holdable option.").

After performing the fetch call, you can query for column information
with the [`getResultCount()`](3027-base-sqlhandle-getresultcount.md "Returns the number of result set columns produced by the SQL statement."), [`getResultName(index)`](3028-base-sqlhandle-getresultname.md "Returns the name of a column in the result set produced by the SQL statement."), [`getResultType(index)`](3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.") and [`getResultValue(index)`](3030-base-sqlhandle-getresultvalue.md "Returns the value of a column in the result set produced by the SQL statement.") methods.

If the result set is empty, or if no row exists at the specified offset relative to the current
row position, `sqlca.sqlcode` is set to 100 (`NOTFOUND`).

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.fetchRelative(-3)
```

For a complete example, see [Example 4: SqlHandle with scroll cursor](3044-example-4-sqlhandle-with-scroll-cursor.md).
