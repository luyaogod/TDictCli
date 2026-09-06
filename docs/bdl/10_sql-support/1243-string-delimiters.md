---
title: "String delimiters"
source: "fgl-topics/c_fgl_odiagdmg_024.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data manipulation > String delimiters"
type: "concept"
description: "Informix® The ANSI SQL string delimiter character is the single quote ( 'string' ), while double quotes are used to delimit database object names: SELECT ... WHERE \"tabname\".\"colname\" = 'a string ..."
---

# String delimiters

## Informix®

The ANSI SQL string delimiter character is the single quote (`'string'`), while
double quotes are used to delimit database object names:

```
SELECT ... WHERE "tabname"."colname" = 'a string value'
```

In Informix databases created in native mode
(non-ANSI), you can use double quotes as string
delimiters:

```
SELECT ... WHERE tabname.colname = 'a string value'
```

This is important, since many BDL programs use that character to delimit the strings in SQL
commands.

> **Note:**
>
> This problem concerns only double quotes within SQL statements. Double quotes used in pure BDL
> string expressions are not subject to SQL compatibility problems.

## Dameng®

Dameng follows the ANSI SQL specification, using single
quotes for string delimiters and double quotes for database object names.

## Solution

When using Static SQL statements, the fglcomp compiler converts string
literals using double quotes to string literals with single
quotes:

```
$ cat s.4gl
MAIN
   DEFINE n INT
   SELECT COUNT(*) INTO n FROM tab1 WHERE col1 = "abc"
END MAIN

$ fglcomp -S s.4gl
s.4gl^3^SELECT COUNT(*) FROM tab1 WHERE col1 = 'abc'
```

However, SQL statements created dynamically are not modified by the Genero compiler.

The Genero database interface can automatically replace all double quotes by single quotes in SQL
statements. This applies to static and dynamic SQL statements.

The The translation of double quoted expression to single quoted expressions can be controlled
with the following FGLPROFILE
entry:

```
dbi.database.dbname.ifxemul.dblquotes = { true | false }
```

For
more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

However, database object names must not be delimited by double quotes, because the database
interface cannot determine the difference between a database object name and a quoted string. For
example, if the program executes the SQL
statement:

```
... WHERE "tabname"."colname" = "a string value"
```

replacing all double quotes by single quotes would
produce:

```
... WHERE 'tabname'.'colname' = 'a string value'
```

This would produce an error
since 'tabname'.'colname' is not allowed by ORACLE.

Escaped string delimiters can be used inside strings like the following:

```
'This is a single quote: '''
'This is a single quote: \''
"This is a double quote: """
"This is a double quote: \""
```

Although
double quotes are replaced automatically in SQL statements, it is recommended that you
use only single quotes to enforce portability.

## Related links

**Related concepts**  

[String literals in SQL statements](1031-string-literals-in-sql-statements.md "Single quotes is the standard for delimiting string literals in SQL.")
