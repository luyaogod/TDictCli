---
title: "base.SqlHandle.prepare"
source: "fgl-topics/c_fgl_ClassSqlHandle_prepare.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.prepare"
type: "concept"
---

# base.SqlHandle.prepare

> Prepares an SQL statement for the SQL handle.

## Syntax

```
prepare(
   sql STRING )
```

1. sql is the SQL statement to be prepared.

## Usage

Call the `prepare()` method to prepare the SQL statement that will be executed
with either `execute()` or `open()`.

The SQL statement can contain `?` parameter place holders, to be filled with the
`setParameter()` method before executing the statement.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.prepare("INSERT INTO mytable VALUES (?,?)")
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
