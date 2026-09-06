---
title: "Table and column names in static SQL"
source: "fgl-topics/c_fgl_static_sql_004.html"
breadcrumb: "SQL support > Static SQL statements > Table and column names in static SQL"
type: "concept"
---

# Table and column names in static SQL

> How are SQL object names and keywords converted in static SQL?

## Case sensitivity with SQL table and column names

In static SQL statements, table and column names will be converted to lowercase by the [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler. The SQL keywords are
always converted to uppercase.

For
example:

```
UPDATE CUSTOMER set CUST_name = 'undef' WHERE cust_name is null
```

Will be converted
to:

```
UPDATE customer SET cust_name = 'undef' WHERE cust_name IS NULL
```

While SQL keywords are not case sensitive for database servers, table names and column names can
be case-sensitive.

> **Tip:**
>
> Dump the static SQL statement texts with the [`-S` option](1119-sql-texts-generated-by-the-compiler.md "The Genero BDL compiler provides an option to extract static SQL statements from .4gl sources.") of fglcomp.

## Checking column names

When compiling sources with fglcomp, the warning option [`-W
colname`](../13_programming-tools/2516-fglcomp.md) can be used to detect column names in static SQL statements, that are not
referenced in the current [database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") used for
compilation. This helps to prevent SQL errors at runtime.

The compiler performs a lookup of column names in the context of the SQL statement, from the
specified tables names. The compiler can not print a warning for an invalid column name, if the
tables used in the SQL statement are not referenced in the database schema.

Example of `-W colname` option usage:

```
$ cat main.4gl
SCHEMA stores
MAIN
    DEFINE cust_id, cnt INTEGER
    SELECT COUNT(*) INTO cnt FROM customer
       WHERE custmer_num = $cust_id
END MAIN

$ fglcomp -W colname main.4gl
main.4gl:5:14:5:24:warning:(-4322) The symbol 'custmer_num' is not the name of a column
 in the specified database.
```

> **Tip:**
>
> Turn fglcomp warnings
> into errors with the -W error option:
>
> ```
> $ fglcomp -M -W colname -W error main.4gl
> main.4gl:5:14:5:24:error:(-4322) The symbol 'custmer_num' is not the name of a column
>  in the specified database.
> ```

## Related links

**Related concepts**  

[SQL basics](0985-sql-basics.md "This section contains fundamental information to know about SQL programming with Genero BDL.")
