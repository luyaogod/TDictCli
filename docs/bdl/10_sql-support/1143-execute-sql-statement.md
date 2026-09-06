---
title: "EXECUTE (SQL statement)"
source: "fgl-topics/c_fgl_DynamicSql_EXECUTE.html"
breadcrumb: "SQL support > Dynamic SQL management > EXECUTE (SQL statement)"
type: "concept"
---

# EXECUTE (SQL statement)

> This instruction runs an SQL statement previously prepared.

## Syntax

```
EXECUTE sid
     [ USING pvar {IN|OUT|INOUT} [,...] ]
     [ INTO fvar [,...] ]
```

1. sid is an identifier to handle the prepared SQL statement.
2. pvar is a variable containing an input value for an SQL parameter.
3. fvar is a variable used as fetch buffer.

## Usage

The `EXECUTE` instruction performs the execution of an SQL statement initiated
by a [`PREPARE`](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") instruction.

Once prepared, an SQL statement can be executed several times with different SQL parameters.

If the SQL statement has (`?`) parameter placeholders, you must specify the
`USING` clause to provide a list of variables as parameter buffers. Parameter
values are assigned by position.

If the SQL statement returns a result set with one row, you can specify the
`INTO` clause to provide a list of variables to receive the result set column
values. Fetched values are assigned by position. If the SQL statement returns a result
set with more than one row, the instruction raises an exception.

The `IN`, `OUT` or `INOUT` options can only be
used for simple variables, you cannot specify those options for a complete record with the
`record.*` notation.

The `IN`, `OUT` or `INOUT` options can be used
to call stored procedures having input / output parameters. Use the `IN`,
`OUT` or `INOUT` options to indicate if a parameter is respectively
for input, output or both.

You cannot execute a prepared SQL statement based on database tables if the table structure
has changed (`ALTER TABLE`) since the `PREPARE` instruction; you
must re-prepare the SQL statement.

## Example

```
MAIN
  DEFINE var1 CHAR(20)
  DEFINE var2 INTEGER

  DATABASE stores 

  PREPARE s1 FROM "UPDATE tab SET col=? WHERE key=?"
  LET var1 = "aaaa"
  LET var2 = 345
  EXECUTE s1 USING var1, var2

  PREPARE s2 FROM "SELECT col FROM tab WHERE key=?"
  LET var2 = 564
  EXECUTE s2 USING var2 INTO var1

  PREPARE s3 FROM "CALL myproc(?,?)"
  LET var1 = 'abc'
  EXECUTE s3 USING var1 IN, var2 OUT

END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")

[Stored procedures](1047-stored-procedures.md "Executing stored procedures with different database engine types.")
