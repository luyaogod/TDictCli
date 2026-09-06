---
title: "SQL ... END SQL"
source: "fgl-topics/c_fgl_static_sql_SQL_blocks.html"
breadcrumb: "SQL support > Static SQL statements > SQL ... END SQL"
type: "concept"
---

# SQL ... END SQL

> Performs an SQL that is not part of the static SQL syntax.

## Syntax

```
SQL
 sql-statement
END SQL
```

where sql-statement
is:

```
  sql-keyword
| identifier
| INTO $host-variable [,...]
| $host-variable
| {+ sql-directive }
| --+ sql-directive
| --# fgl-comment
[...]
```

1. sql-keyword is any keyword of the SQL language.
2. identifier is a regular SQL identifier such
   as a table or column name.
3. host-variable is a program variable defined
   in the current scope.
4. sql-directive is a special comment to be kept
   in the SQL statement.
5. fgl-comment defines a comment that will be
   interpreted as a regular syntax element.

## Usage

SQL blocks provide a convenient
way to execute specific SQL statements that are not supported
in the language as static SQL statements.

SQL blocks start with the `SQL` keyword and end with the `END
SQL` keywords. The content of the SQL block is parsed by the fglcomp
compiler to extract host variables, but the SQL statement syntax is not checked. This is actually
the main purpose of SQL blocks, compared to regular static SQL statements; with SQL blocks, you can
use any recent SQL statement introduced by the latest version of your database server. Note,
however, that you can achieve the same result using dynamic SQL instructions.

Only one SQL statement
can be included in an SQL block. Using the `;` semicolon
statement separator is forbidden.

Program variables can be used inside the SQL statement. However, unlike static SQL
statements, each host variable must be identified with a `$` dollar prefix. The list
of fetch targets must be preceded by the `INTO` keyword, as in static
`SELECT` statements. Complete records can be used in SQL blocks by using the dot star
notation (`$record.*`), you can also use the `THROUGH` or
`THRU` keywords), as well as array elements.

SQL blocks can also be used to
declare a cursor with the `DECLARE mycursor CURSOR FOR SQL
... END SQL` syntax.

SQL directives
can be used inside SQL blocks as special comments with the `{+}` or
`--+` syntax. The SQL directives will
be kept in the SQL text that will be executed by the database
server. You typically write optimizer hints with the SQL
directives syntax.

The `--#` specific comment
is supported for backward compatibility. The SQL text following
this marker will be parsed as regular SQL text, but will be ignored
by other compilers. It is not recommended to use this feature.

You can check the resulting SQL statement after parsing by using the `-S` option
of fglcomp.

## Example

```
MAIN
   DEFINE myrec RECORD
            key INTEGER,
            name CHAR(10)
         END RECORD
   DATABASE stock 
   LET myrec.key = 123
   SQL
     SELECT (+EXPLAIN) items.* INTO $myrec.*
        FROM items WHERE key=$myrec.key 
   END SQL
END MAIN
```

## Related links

**Related concepts**  

[Result set processing](1148-result-set-processing.md "Shows how to fetch rows from a database query.")
