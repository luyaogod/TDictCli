---
title: "base.SqlHandle.open"
source: "fgl-topics/c_fgl_ClassSqlHandle_open.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.open"
type: "concept"
---

# base.SqlHandle.open

> Opens the SQL handle (SELECT or INSERT cursor).

## Syntax

```
open()
```

## Usage

Call the `open()` method to execute the prepared SQL statement, and open the result
set cursor or insert cursor.

The SQL statement must have been prepared with a [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") call.

If the SQL statement contains `?` parameter placeholders, use a [`setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle.") call for each
parameter value:

- For a statement with a result set (`SELECT`), values must be provided before the
  `open()` call.
- For an insert cursor, values must be provided after the `open()` call and before
  each [`put()`](3036-base-sqlhandle-put.md "Put a new row in the insert cursor buffer.") call.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.open()
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
