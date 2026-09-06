---
title: "DECLARE (insert cursor)"
source: "fgl-topics/c_fgl_InsertCursors_DECLARE.html"
breadcrumb: "SQL support > SQL insert cursors > DECLARE (insert cursor)"
type: "concept"
---

# DECLARE (insert cursor)

> The DECLARE with an INSERT instruction defines an insert cursor.

## Syntax

```
DECLARE cid CURSOR [WITH HOLD] FOR { insert-statement | sid }
```

1. cid is the identifier of the insert cursor.
2. insert-statement is an `INSERT` statement defined in static
   SQL.
3. sid is the identifier of a prepared `INSERT` statement.

## Usage

Use the `DECLARE` instruction with an `INSERT` instruction to
define a new insert cursor in the current database session.

The `INSERT` statement is parsed, validated
and the execution plan is created.

`DECLARE` must
precede any other statement that refers to the cursor during program
execution.

The scope of reference of the cid cursor identifier is local to the module
where it is declared.

After declaring the insert cursor, it must be opened with the [`OPEN`](1166-open-insert-cursor.md "Initializes an insert cursor.") instruction.

The static insert-statement statement can include a list of variables in
the `VALUES` clause. These variables are automatically read by the [`PUT`](1167-put-insert-cursor.md "Adds a new row to the insert cursor buffer.") statement; you do not have to
provide the list of variables in that statement. As an alternative, use the ? (question mark) SQL
parameter placeholder in the `VALUE` clause to bind program variables provided in the
`FROM` clause of the `PUT` instruction.

When declaring a cursor with a prepared sid statement, the statement can
include ? (question mark) placeholders for SQL parameters. In this case you must
provide a list of variables in the `FROM` clause of the
`PUT` statement.

Use the `WITH HOLD` option to declare cursors that have uninterrupted inserts
across multiple transactions.

Resources allocated by the `DECLARE` can be released later by the [`FREE`](1170-free-insert-cursor.md "Releases resources allocated for an insert cursor.") instruction.

The number of declared cursors in a single program is limited by the database server and the
available memory. Make sure that you free the resources when you no longer need the
declared insert cursor.

The
identifier of a cursor that was declared in one module cannot be referenced
from another module.

## Related links

**Related concepts**  

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
