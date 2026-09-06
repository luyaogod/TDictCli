---
title: "base.SqlHandle.put"
source: "fgl-topics/c_fgl_ClassSqlHandle_put.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.put"
type: "concept"
---

# base.SqlHandle.put

> Put a new row in the insert cursor buffer.

## Syntax

```
put()
```

## Usage

Call the `put()` method to create a new row for the insert cursor.

The SQL statement must have been prepared with a [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") call.

All SQL parameter values must be provided with a [`setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle.") call, before
performing the `put()` call.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.setParameter(1,7623)
CALL sh.setParameter(2,"Scott")
CALL sh.put()
...

CALL sh.setParameter(1,7624)
CALL sh.setParameter(2,"Mike")
CALL sh.put()
```

For a complete example, see [Example 3: SqlHandle with insert cursor](3043-example-3-sqlhandle-with-insert-cursor.md).
