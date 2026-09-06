---
title: "base.SqlHandle.close"
source: "fgl-topics/c_fgl_ClassSqlHandle_close.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.close"
type: "concept"
---

# base.SqlHandle.close

> Closes the SQL handle (cursor).

## Syntax

```
close()
```

## Usage

Call the `close()` method when you are finished using the SQL handle as a cursor
opened with the [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).")
method.

The statement can be re-opened after it has been closed.

A SqlHandle object is automatically closed when the object is destroyed.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.close()
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
