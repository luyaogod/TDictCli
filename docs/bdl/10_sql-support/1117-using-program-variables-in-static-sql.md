---
title: "Using program variables in static SQL"
source: "fgl-topics/c_fgl_static_sql_003.html"
breadcrumb: "SQL support > Static SQL statements > Using program variables in static SQL"
type: "concept"
---

# Using program variables in static SQL

> Static SQL syntax supports the usage of program variables as SQL parameters.

## Understanding SQL host variables

Using [program variables](../08_language-basics/0686-variables.md "Explains how to define program variables.") directly in static SQL
statements gives a better understanding of the source code and requires less lines as when using SQL
parameters in dynamic SQL
statements.

```
MAIN
  DEFINE c_num INTEGER
  DEFINE c_name CHAR(10)
  DATABASE stock 
  SELECT cust_name INTO c_name FROM customer WHERE cust_num = c_num
END MAIN
```

> **Important:**
>
> When using program variables in a static SQL statement of a [`DECLARE name CURSOR`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.")
> instruction, make sure that these variables are available when executing the `OPEN`
> or `FOREACH` instruction. Otherwise, the behavior is unexpected and can make the
> runtime system crash. Do not mix `?` SQL parameter placeholders with program
> variables: This would lead to a runtime error at `OPEN` or at
> `FOREACH` time, because the number of SQL parameters will not match the number of
> variable provided in the `USING` clause. The fglcomp -S option
> extracts the static SQL statements showing `?` placeholders instead of programs
> variables. This can help to check the actual number of SQL parameters in a static SQL statement.

## Using the $ prefix for host variables

The purpose of the `$` prefix for program variables is to make the code more
readable, and force compilation errors whenever a variable is not defined. This feature was inspired
from ESQL/C. It applies only to static SQL statements (parsed by the compiler).

When using the `$` in static SQL statements, the compiler treats the identifier
following the `$` as a program variable:

```
SCHEMA stores
MAIN
    DEFINE rec RECORD LIKE customer.*
    DEFINE cust_id INTEGER
    SELECT * INTO rec.* FROM customer
        WHERE customer_num = $cust_id
END MAIN
```

The `$` sign will not figure in the resulting SQL statement stored in the
.42m compiled module.

In a static `SELECT` statement, the `INTO
var-list` clause is not part of the final SQL statement. There is no need
for a `$` dollar sign before variables that appear after the `INTO`
keyword.

When not marking program variables with `$`, the compiler performs a fuzzy look up
for each symbol that can be a potential program variable: If a variable exists for a given symbol,
the compiler uses that variable. Otherwise, the compiler assumes that the symbol is a column
name.

Using `$` dollar signs before program variables in SQL statements prevents
spelling
errors:

```
DEFINE cid INTEGER
SELECT ... WHERE cust_id = cid     -- ok, cid is a program variable
SELECT ... WHERE cust_id = $cid    -- ok, cid is a (marked) program variable
SELECT ... WHERE cust_id = c_id    -- fglcomp assumes column => gives SQL error at runtime!
SELECT ... WHERE cust_id = cust_id -- fglcomp assumes column => wrong SQL results at runtime!
SELECT ... WHERE cust_id = $c_id   -- fglcomp assumes variable => error -4369 (undefined)
```

See also:

- [Mark SQL host variables with $](../13_programming-tools/2644-mark-sql-host-variables-with.md "The fglcomp compiler can mark SQL host variables with a $ dollar sign.")
- [Complete column and variable names in static SQL](../13_programming-tools/2543-source-code-completion.md)
- [Checking column names](1118-table-and-column-names-in-static-sql.md)

## Using the @ prefix for database object names

When using the `@` prefix, the compiler treats the identifier following the
`@` as an SQL object name (this syntax is supported for backward compatibility with
Informix
4GL):

```
MAIN
  DEFINE cust_name CHAR(10)
  DEFINE cnt INTEGER
  DATABASE stock 
  SELECT COUNT(*) INTO cnt FROM customer WHERE @cust_name = cust_name 
END MAIN
```

The `@` sign will not figure in the resulting SQL statement stored in the
.42m compiled module.

A database object name may also conflict with another symbol of the program code, such as a
module name. In this case, you can also use the `@` sign before the conflicting
database object name to solve the issue.

In the next example, the imported module name conflicts with a database table name
"account":

```
IMPORT FGL account
MAIN
    DEFINE rec RECORD
               pkey INTEGER,
               name VARCHAR(50)
           END RECORD
    SELECT * INTO rec.* FROM account
      WHERE @account.pkey = rec.pkey
END MAIN
```

Without the `@` sign, the compiler would produce the following error, expecting
that `account.pkey` is a public variable of the imported `account`
module:

```
IMPORT FGL account
MAIN
    DEFINE rec RECORD
               pkey INTEGER,
               name VARCHAR(50)
           END RECORD
    SELECT * INTO rec.* FROM account
      WHERE account.pkey = rec.pkey
| The symbol 'pkey' does not represent a defined variable.
| See error number -4369.
END MAIN
```

## Checking for invalid column names

The fglcomp compiler provides the [`-W colname`](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") option, to check for SQL column names in static SQL statements,
based on the SQL table definitions of the current schema file.

For example, it the stores.sch schema file defines a
"`customer`" table with "`fname`" and "`lname`"
columns, the following code will produce a warning when using the `-W colname`
option:

```
SCHEMA stores
MAIN
    DEFINE cnt INTEGER
    SELECT COUNT(*) INTO cnt FROM customer
       WHERE xfname IS NULL
END MAIN
```

Using the `-W colname` option:

```
$ fglcomp -W colname main.4gl
main.4gl:5:14:5:19:warning:(-4322) The symbol 'xfname' is not the name of a column
   in the specified database.
```

## Checking risky SQL cursor usage

SQL cursors have a modular scope. The [`DECLARE`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") instruction can be implemented in a function, while the
`OPEN`, `FETCH`, `FOREACH` or `PUT`
instructions can be used in another function.

When the `DECLARE` statement uses a static SQL (`SELECT` or
`INSERT`) with local function variables as SQL parameters or `INTO`
targets, these variables will no longer be on the FGL stack in the scope of the other function doing
the `OPEN`, `FETCH`, `FOREACH` or `PUT`.
This can lead to unexpected behavior, when the last instructions do not have `USING`,
`INTO` or `FROM` (`PUT`) clauses to bind variables
available in the current scope.

For example:

```
FUNCTION func1()
    DEFINE v1, v2 INT
    DECLARE c1 CURSOR FOR
      SELECT col1 INTO v1 -- local function variable
        FROM tab
       WHERE col1 = v2 -- local function variable
END FUNCTION

FUNCTION func2()
    OPEN c1 -- Missing USING clause
    FETCH c1 -- Missing INTO clause
END FUNCTION
```

The fglcomp compiler provides the [`-W
fragile-cursor`](../13_programming-tools/2516-fglcomp.md) option, to detect such risky SQL cursor usage. With the above code,
you will get following warnings:

```
$ fglcomp -W all sample.4gl
sample.4gl:10:5:10:11:warning:(-8460) Cursor c1 is used in a risky manner.
sample.4gl:11:5:11:12:warning:(-8460) Cursor c1 is used in a risky manner.
```

The `-W fragile-cursor` option can also detect old legacy `OPEN
USING` followed by `FOREACH` without `USING` clause. This was
mandatory in early versions of Informix 4GL because [`FOREACH`](1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor.") had no `USING` clause. In addition, when the
Informix OPTOFC option is enabled, successive `OPEN USING / FETCH / OPEN / FOREACH`
can return invalid rows.

For example:

```
FUNCTION func1()
    DEFINE v1, v2 INT
    DECLARE c1 CURSOR FOR SELECT col2 FROM tab1 WHERE col1 = v1
    OPEN c1 USING v1
    FOREACH c1 INTO v2 -- missing USING clause
        DISPLAY "v2 = ", v2
    END FOREACH
END FUNCTION
```

With above sample, fglcomp -W fragile-cursor will produce the following
warning:

```
sample2.4gl:5:5:7:15:warning:(-8460) Cursor c1 is used in a risky manner.
```

## Related links

**Related concepts**  

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
