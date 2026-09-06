---
title: "FREE (SQL statement)"
source: "fgl-topics/c_fgl_DynamicSql_FREE.html"
breadcrumb: "SQL support > Dynamic SQL management > FREE (SQL statement)"
type: "concept"
---

# FREE (SQL statement)

> Releases the resources allocated to a prepared statement.

## Syntax

```
FREE sid
```

1. sid is the identifier of the prepared SQL statement.

## Usage

The `FREE` instruction takes the name of a statement as parameter and releases
the resources allocated by the [`PREPARE`](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") instruction.

After using `FREE`, the statement identifier cannot be referenced by a cursor
declaration ([`DECLARE`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.")), or by the
[`EXECUTE`](1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.") statement, until you
prepare the statement again.

Free the statement if
it is not needed anymore, this saves resources on the database client
and database server side.

## Example

```
FUNCTION update_customer_name( key, name )
  DEFINE key INTEGER
  DEFINE name CHAR(10)
  PREPARE s1 FROM "UPDATE customer SET name=? WHERE customer_num=?"
  EXECUTE s1 USING name, key 
  FREE s1
END FUNCTION
```
