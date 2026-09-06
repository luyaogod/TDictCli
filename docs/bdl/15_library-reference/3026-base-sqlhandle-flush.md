---
title: "base.SqlHandle.flush"
source: "fgl-topics/c_fgl_ClassSqlHandle_flush.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.flush"
type: "concept"
---

# base.SqlHandle.flush

> Flushes the rows from the insert cursor buffer.

## Syntax

```
flush()
```

## Usage

With an insert cursor, call the `flush()` method to force
the buffered rows to the database server.

The SQL statement must have been opened with an [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).") call.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.flush()
```

For a complete example, see [Example 3: SqlHandle with insert cursor](3043-example-3-sqlhandle-with-insert-cursor.md).
