---
title: "base.SqlHandle.openScrollCursorWithHold"
source: "fgl-topics/c_fgl_ClassSqlHandle_openScrollCursorWithHold.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.openScrollCursorWithHold"
type: "concept"
---

# base.SqlHandle.openScrollCursorWithHold

> Opens the SQL handle with scrollable and holdable option.

## Syntax

```
openScrollCursorWithHold()
```

## Usage

Call the `openScrollCursorWithHold()` method to execute a prepared SQL statement,
and open the result set for use as a scrollable SQL cursor, with hold option, to keep the result set
available across transaction boundaries.

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
CALL sh.openScrollCursorWithHold()
BEGIN WORK
CALL sh.fetchFirst()
...
COMMIT WORK
CALL sh.fetchLast()
```
