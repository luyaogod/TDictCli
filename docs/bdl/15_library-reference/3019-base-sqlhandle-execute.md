---
title: "base.SqlHandle.execute"
source: "fgl-topics/c_fgl_ClassSqlHandle_execute.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.execute"
type: "concept"
---

# base.SqlHandle.execute

> Executes a simple SQL statement (without result set).

## Syntax

```
execute()
```

## Usage

Call the `execute()` method to perform the SQL statement prepared by a [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") call.

The `execute()` method is typically used for SQL statements that do not produce a
result set, such as `INSERT`, `UPDATE`, `DELETE` or DDL
statements like `CREATE TABLE`.

The `execute()` method can also be used to fetch a single row from a
`SELECT` statement: After executing a `SELECT` statement, get the
column values with [`getResultValue()`](3030-base-sqlhandle-getresultvalue.md "Returns the value of a column in the result set produced by the SQL statement."). However, it is better practice to use [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).") + [`fetch()`](3020-base-sqlhandle-fetch.md "Fetches a new row from the SQL result set.") for SQL statements returning
a result set.

If the SQL statement contains `?` parameter place holders, issue a [`setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle.") call for each
parameter, before executing the SQL statement.

As with standard Genero SQL instructions, SQL errors can be trapped
with `WHENEVER ERROR` or `TRY / CATCH` blocks and by testing
`sqlca.sqlcode`.

## Example

```
DEFINE sh base.SqlHandle
...
CALL sh.execute()
```

For a complete example, see [Example 1: SqlHandle with simple SQL](3041-example-1-sqlhandle-with-simple-sql.md).
