---
title: "Understanding dynamic SQL"
source: "fgl-topics/c_fgl_DynamicSql_002.html"
breadcrumb: "SQL support > Dynamic SQL management > Understanding dynamic SQL"
type: "concept"
---

# Understanding dynamic SQL

> This is an introduction to dynamic SQL programming.

Basic SQL instructions are part of the language syntax as [Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language."), but only a limited number of SQL instructions are supported this
way.

Dynamic SQL management allows you to execute any kind of SQL statement, hard coded
or created at runtime, with or without SQL parameters, returning or not returning a result set.

In order to execute an SQL statement dynamically, you must first [`PREPARE`](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") the SQL statement to initialize
a statement handle, then [`EXECUTE`](1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.") the prepared statement one or more times:

![Dynamic SQL management diagram](../_images/DNSFig01.jpg)

*Dynamic SQL management diagram*

When you no longer need the prepared statement, you can [`FREE`](1144-free-sql-statement.md "Releases the resources allocated to a prepared statement.") the statement handle to release allocated resources:

![FREE statement diagram](../_images/DNSFig02.jpg)

*FREE statement diagram*

When using insert cursors or SQL statements that produce a result set (like `SELECT`),
you must declare a cursor with a prepared statement handle.

Prepared SQL statements can contain SQL parameters by using `?` placeholders in
the SQL text. In this case, the `EXECUTE` or `OPEN` instruction
supplies input values in the `USING` clause.

To increase performance efficiency of SQL executed in a loop, use `PREPARE`
outside the loop, together with `EXECUTE` inside the loop, to eliminate overhead
caused by redundant parsing and optimizing.

The [`EXECUTE
IMMEDIATE`](1145-execute-immediate.md "Performs a simple SQL execution without SQL parameters or result set.") instruction prepares and executes an SQL statement in a single
instruction. SQL parameters and result sets cannot be used with `EXECUTE
IMMEDIATE`.

The [`base.SQLHandle`](1146-the-base-sqlhandle-built-in-class.md "Handle SQL queries with a 3GL API.")
built-in class is a 3GL API to execute SQL statements dynamically and perform SQL introspection (to
get result set column types).

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
