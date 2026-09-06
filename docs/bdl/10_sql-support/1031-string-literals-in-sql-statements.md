---
title: "String literals in SQL statements"
source: "fgl-topics/c_fgl_sql_programming_080.html"
breadcrumb: "SQL support > SQL programming > SQL portability > String literals in SQL statements"
type: "concept"
---

# String literals in SQL statements

> Single quotes is the standard for delimiting string literals in SQL.

Some database servers like IBM®
Informix® allow single and double quoted string
literals in SQL statements, both are equivalent:

```
SELECT COUNT(*) FROM table
 WHERE col1 = "abc'def""ghi"
   AND col1 = 'abc''def"ghi'
```

Most database servers do not support this specific feature.

| Database Server Type | Double quoted string literals |
| --- | --- |
| IBM Informix | Yes |
| Microsoft™ SQL Server | Yes |
| Oracle® MySQL / MariadDB | No |
| Oracle Database Server | No |
| PostgreSQL | No |
| SQLite | Yes |
| Dameng® | No |

The ANSI SQL standards define doubles quotes as database object names
delimiters, while single quotes are dedicated to string
literals:

```
CREATE TABLE "my table" ( "column 1" CHAR(10) ) 
SELECT COUNT(*) FROM "my table" WHERE "column 1" = 'abc'
```

If you want to write a single quote character inside a string literal,
you must write 2 single quotes:

```
... WHERE comment = 'John''s house'
```

When writing static SQL in your programs, the double quoted string
literals as converted to ANSI single quoted string literals by the fglcomp
compiler. However, dynamic SQL statements are not parsed by the compiler
and therefore need to use single quoted string literals.

We recommend that you always use single quotes for string literals
and, if needed, double quotes for database object names.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
