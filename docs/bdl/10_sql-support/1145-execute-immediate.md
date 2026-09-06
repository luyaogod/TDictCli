---
title: "EXECUTE IMMEDIATE"
source: "fgl-topics/c_fgl_DynamicSql_EXECUTE_IMMEDIATE.html"
breadcrumb: "SQL support > Dynamic SQL management > EXECUTE IMMEDIATE"
type: "concept"
---

# EXECUTE IMMEDIATE

> Performs a simple SQL execution without SQL parameters or result set.

## Syntax

```
EXECUTE IMMEDIATE sqltext
```

1. sqltext is a string expression containing the SQL statement to be
   executed.

## Usage

The `EXECUTE IMMEDIATE` instruction passes an SQL statement to the database server
for execution in the current database connection.

The SQL statement used by `EXECUTE IMMEDIATE` must be a single statement without
SQL parameters and must not produce a result set.

This instruction is equivalent to [`PREPARE`](1142-prepare-sql-statement.md "Prepares an SQL statement for execution."),
[`EXECUTE`](1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.") and [`FREE`](1144-free-sql-statement.md "Releases the resources allocated to a prepared statement.") done in one step.

Unlike the `EXECUTE` instruction, `EXECUTE IMMEDIATE` has no
`USING` clause (to specify SQL parameters), nor does it support an
`INTO` clause (to fetch values into variables). When specifying a
`USING` keyword after the string containing the SQL statement, the compiler will
consider this as the [`USING` formatting
operator](../08_language-basics/0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.").

## Example

```
MAIN
  DATABASE stores 
  EXECUTE IMMEDIATE "UPDATE tab SET col='aaa' WHERE key=345"
END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
