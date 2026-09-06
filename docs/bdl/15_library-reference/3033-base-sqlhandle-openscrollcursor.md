---
title: "base.SqlHandle.openScrollCursor"
source: "fgl-topics/c_fgl_ClassSqlHandle_openScrollCursor.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.openScrollCursor"
type: "concept"
---

# base.SqlHandle.openScrollCursor

> Opens the SQL handle (with scrollable option).

## Syntax

```
openScrollCursor()
```

## Usage

Call the `openScrollCursor()` method to execute a prepared SQL statement, and open
the result set for use as a scrollable SQL cursor.

The SQL statement must have been prepared with a [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") call.

If the SQL statement contains `?` parameter placeholders,
use a [`setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle.")
call for each parameter value. Values must be provided before opening the cursor.

After opening the scrollable cursor, use methods such as
`fetchFirst()`, `fetchPrevious()` and
`fetchAbsolute(n)` to move forwards and backwards in the SQL result set.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.openScrollCursor()
```
