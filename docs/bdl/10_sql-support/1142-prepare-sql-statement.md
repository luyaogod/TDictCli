---
title: "PREPARE (SQL statement)"
source: "fgl-topics/c_fgl_DynamicSql_PREPARE.html"
breadcrumb: "SQL support > Dynamic SQL management > PREPARE (SQL statement)"
type: "concept"
---

# PREPARE (SQL statement)

> Prepares an SQL statement for execution.

## Syntax

```
PREPARE sid FROM sqltext
```

1. sid is an identifier to handle the prepared
   SQL statement.
2. sqltext is a string expression containing the
   SQL statement to be prepared.

## Usage

The `PREPARE` instruction
allocates resources for an SQL statement handle, in the context
of the current database connection. The SQL text is sent to the database
server for parsing, validation and to generate the execution
plan.

Prepared SQL statements can be executed with the [`EXECUTE`](1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.") instruction, or, when the SQL statement generates a result set, the
prepared statement can be used to declare cursors with the `DECLARE` instruction.

A statement identifier (sid) can represent only one SQL statement at a time.
You can execute a new `PREPARE` instruction with an existing statement identifier if
you wish to assign the text of a different SQL statement to the statement identifier. The scope of
reference of the sid statement identifier is local to the module where it is
declared. That is, the identifier of a statement that was prepared in one module cannot be
referenced from another module.

The SQL statement can have parameter
placeholders, identified by the question mark (`?`)
character. You cannot directly reference a variable in the text of
a prepared SQL statement. You cannot use question mark (`?`)
placeholders for SQL identifiers such as a table name or a
column name; you must specify these identifiers in the statement
text when you prepare it.

Resources allocated by `PREPARE` can be released later by the [`FREE`](1144-free-sql-statement.md "Releases the resources allocated to a prepared statement.") instruction.

The number of prepared statements in a single program is limited
by the database server and the available memory. Make sure
that you free the resources when you no longer need the prepared
statement.

Some database servers support multiple SQL statement preparation in a unique
`PREPARE` instruction, but most database servers deny multiple statements.
It is recommended that you only prepare one SQL statement at a time.

## Example

```
FUNCTION deleteOrder(n)
  DEFINE n INTEGER
  PREPARE s1 FROM "DELETE FROM order WHERE key=?"
  EXECUTE s1 USING n 
  FREE s1
END FUNCTION
```

See
[`EXECUTE`](1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.") for more code
examples.

## Related links

**Related concepts**  

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
