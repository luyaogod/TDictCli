---
title: "base.SqlHandle.create"
source: "fgl-topics/c_fgl_ClassSqlHandle_create.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.create"
type: "concept"
---

# base.SqlHandle.create

> Create a new base.SqlHandle object.

## Syntax

```
base.SqlHandle.create()
  RETURNS base.SqlHandle
```

## Usage

Use the `create()` method to create a `base.SqlHandle` object to
execute SQL statements.

The value returned by this method must be assigned to a variable defined with the
`base.SqlHandle` type.

As with other built-in classes, the SqlHandle object will be automatically destroyed if no longer
referenced.

The next step is to prepare an SQL statement with the [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") method.

## Example

```
DEFINE sh base.SqlHandle
LET sh = base.SqlHandle.create()
...
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
