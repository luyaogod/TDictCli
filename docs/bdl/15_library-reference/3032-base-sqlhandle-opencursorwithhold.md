---
title: "base.SqlHandle.openCursorWithHold"
source: "fgl-topics/c_fgl_ClassSqlHandle_openCursorWithHold.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.openCursorWithHold"
type: "concept"
---

# base.SqlHandle.openCursorWithHold

> Opens the SQL handle with holdable option.

## Syntax

```
openCursorWithHold()
```

## Usage

Call the `openCursorWithHold()` method to execute the prepared SQL statement, and
open the result set cursor with hold option, to keep the result set available across transaction
boundaries.

The SQL statement must have been prepared with a [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") call.

If the SQL statement contains `?` parameter placeholders,
use a [`setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle.")
call for each parameter value. Values must be provided before opening the cursor.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.openCursorWithHold()
BEGIN WORK
CALL sh.fetch()
...
COMMIT WORK
CALL sh.fetch()
```
