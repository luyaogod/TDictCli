---
title: "SQL identifiers case-sensitivity"
source: "fgl-topics/c_fgl_sql_programming_084.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Naming database objects > SQL identifiers case-sensitivity"
type: "concept"
---

# SQL identifiers case-sensitivity

> Handling case-sensitivity with different database engines.

In general, database engines use case-sensitive object identifiers. In most cases, when you do
not specify identifiers in double quotes, the SQL parser automatically converts names to uppercase
or lowercase. With this rule, identifiers will match when the objects are also created without
double quoted identifiers.

For example, when creating the following table in ORACLE DB:

```
CREATE TABLE Customer ( cust_ID INTEGER )
```

With the above SQL, ORACLE creates a table named "`CUSTOMER`" with a
"`CUST_ID`" column. Since DB object identifiers are converted to uppercase when
executing SQL statements, the next SELECT statement will match the "`CUSTOMER`" table
name and "`CUST_ID`" column
name:

```
SELECT * FROM customer WHERE CUST_ID = 219
```

However, the next statements
will fail:

```
SELECT * FROM "customer" WHERE "CUST_ID" = 219
```

The next table shows the behavior of each database engine regarding case sensitivity and double
quoted identifiers:

| Database Server Type | Un-quoted names | Double-quoted names |
| --- | --- | --- |
| IBM® Informix® | Converts to lowercase | Needs DELIMIDENT environment variable |
| Microsoft™ SQL Server | Not converted, kept as is | Needs `SET QUOTED_IDENTIFIER ON` (native delimiters are `[ ]` brackets) |
| Oracle® MySQL / MariadDB | Not converted, kept as is | Needs `ANSI_QUOTES` SQL mode |
| Oracle Database Server | Converts to uppercase | Yes |
| PostgreSQL | Converts to lowercase | Yes |
| SQLite | Not converted, kept as is | Warning: Case insensitive! |
| Dameng® | Not converted, kept as is | Yes, case sensitive when `CASE_SENSITIVE` database config parameter is Y |

> **Important:**
>
> When DB object identifiers are case sensitive and are not converted to
> uppercase or lowercase when not using double-quotes, it is possible to create two tables with a
> similar name (mind the uppercase C in the table name):
>
> ```
> CREATE TABLE customer ( cust_id INTEGER )  -- first table 
> CREATE TABLE Customer ( cust_id INTEGER )  -- second table
> ```

For maximum protability, define database object names in lowercase and do not use double quoted
identifiers.

When using [Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language."), the
fglcomp compiler converts table and column names to lowercase when not using
single quotes or double quotes. Check the output of fglcomp -S option.
