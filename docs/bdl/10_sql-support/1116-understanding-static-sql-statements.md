---
title: "Understanding static SQL statements"
source: "fgl-topics/c_fgl_static_sql_002.html"
breadcrumb: "SQL support > Static SQL statements > Understanding static SQL statements"
type: "concept"
---

# Understanding static SQL statements

> This is an introduction to static SQL statements.

Static SQL statements are SQL instructions that are a part of the Genero BDL
language syntax. Static SQL statements can be used directly in the source code as a normal
procedural instruction. The static SQL statements are parsed and validated at compile time. At
runtime, these SQL statements are automatically prepared and executed by the runtime system.

Program variables can be used inside static SQL statements. Variables are detected by the
compiler and handled as SQL parameters at runtime.

The following example defines two variables that are directly used
in an `INSERT` statement:

```
MAIN
  DEFINE iref INTEGER, name CHAR(10)
  DATABASE stock 
  LET iref = 65345
  LET name = "Kartopia"
  INSERT INTO item ( item_ref, item_name ) VALUES ( iref, name )
  SELECT item_name INTO name 
    FROM item WHERE item_ref = iref 
END MAIN
```

As it is integrated in the language syntax, static SQL statement usage clarifies the source code,
but the SQL text is hard-coded and cannot be modified at runtime as is possible with [`PREPARE` / `EXECUTE`
instructions](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.") of dynamic SQL.

Limited SQL syntax is part of the language, only common SQL statements such as [`INSERT`](1120-insert.md "Creates a new row in a database table."), [`UPDATE`](1121-update.md "Modifies rows of a database table."), [`DELETE`](1122-delete.md "Removes rows from a database table."), [`SELECT`](1123-select.md "Produces a result set from a query on database tables.") are supported.

The compiler supports also [`SQL ... END
SQL`](1124-sql-end-sql.md "Performs an SQL that is not part of the static SQL syntax.") blocks to write free SQL text in your programs. The SQL syntax in SQL blocks is
not limited to the static SQL syntax.

## Related links

**Related concepts**  

[Database connections](1059-database-connections.md "Explains how to manage database connections in a program.")
