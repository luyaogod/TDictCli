---
title: "Data manipulation statements"
source: "fgl-topics/c_fgl_sql_programming_062.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Data manipulation statements"
type: "concept"
---

# Data manipulation statements

> Make sure that SQL statement syntaxes are supported by all target database engines.

Several SQL syntaxes for the `INSERT`, `UPDATE` and `DELETE`
statements are supported by the compiler. Some of the syntaxes are IBM®
Informix® specific, but will be converted to standard SQL at
compile time.

The following statements are standard SQL and work with all database servers:

```
(1) INSERT INTO table (column-list) VALUES (value-list)
(2) UPDATE table SET column = value, ... [WHERE condition]
(3) DELETE FROM table [WHERE condition]
```

The next statements are not standard SQL, but are converted by the compiler to standard SQL,
working with all database
servers:

```
(4) INSERT INTO table VALUES record.*
    -- where record is defined LIKE a table from db schema   
(5) UPDATE table SET (column-list) = (value-list) [WHERE condition]
(6) UPDATE table SET {[table.]*|(column-list)} = record.* ... [WHERE condition]
   -- where record is defined LIKE a table from db schema
(7) UPDATE table SET [table.]* = (value-list) [WHERE condition]
```

For maximum SQL portability, `INSERT` statements should be reviewed to ensure the
`SERIAL` column is excluded from the value list.

You can easily search for non-portable SQL statements in your sources by compiling with the
`-W stdsql` fglcomp option.

For example, the following statement:

```
INSERT INTO tab (col1, col2, ...) VALUES ( 0, p_value2, ... )
```

should be converted to:

```
INSERT INTO tab (col2, ...) VALUES ( p_value2, ... )
```

A static SQL `INSERT` statement using records defined from the schema file should
also be reviewed:

```
DEFINE rec LIKE tab.*
INSERT INTO tab VALUES ( rec.* )   -- will use the serial column
```

should be converted to:

```
INSERT INTO tab VALUES rec.* -- without parentheses, serial
column is removed
```

Using the `record.*` notation in static `INSERT`
and `UPDATE` syntax may not be compatible with database-specific features, where some
automatically-assigned columns must not be set or modified by the statement. For example, with
Microsoft SQL Server temporal tables, timestamp columns are automatically assigned and must not be
changed by `INSERT` or `UPDATE` statements. In such case, it is
mandatory to explicitly list all modifiable columns and corresponding program variables.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
