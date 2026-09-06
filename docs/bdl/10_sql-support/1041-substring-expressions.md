---
title: "Substring expressions"
source: "fgl-topics/c_fgl_sql_programming_089.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Substring expressions"
type: "concept"
---

# Substring expressions

> Handle substrings expressions with different database engines.

Only IBM® Informix® supports substring specification
with square brackets:

```
SELECT * FROM item WHERE item_code[1,4] = "XBFG"
```

This syntax is specific to Informix SQL. Other database server types provide a function that
extracts substrings from a character string.

| Database Server Type | Substring function | col[x,y] support |
| --- | --- | --- |
| IBM Informix | SUBSTR(expr,start,length) | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | SUBSTRING(expr,start,length) | Emulated, [see details](1292-substrings-in-sql.md) |
| Oracle® MySQL / MariadDB | SUBSTR(expr,start,length) | Emulated, [see details](1342-substrings-in-sql.md) |
| Oracle Database Server | SUBSTR(expr,start,length) | Emulated, [see details](1394-substrings-in-sql.md) |
| PostgreSQL | SUBSTRING(expr FROM start FOR length) SUBSTR(expr,start,length) | Emulated, [see details](1447-substrings-in-sql.md) |
| SQLite | SUBSTR(expr,start,length) | Emulated, [see details](1492-substrings-in-sql.md) |
| Dameng® | SUBSTR(expr,start,length) | Emulated, [see details](1242-substrings-in-sql.md) |

Informix allows you to update some parts
of a CHAR and VARCHAR column by using the substring specification ( `UPDATE tab SET col[1,2]
='ab'` ). This is not possible with other databases.

Review the SQL statements using substring expressions and use the database specific substring
function. Depending on the database types to be used, use SUBSTR() or SUBSTRING() as provided by the
database engines.

You may also want to create your own user function with the databases engines that do not support
the same syntax as most of your database engines, to have a common way to extract substrings. Note
that with Microsoft SQL Server, a user function call must
specify the owner as prefix. Therefore, when using SQL Server, instead of creating a SUBSTR()
function, it is recommended that you create a SUBSTRING() function in other database types, since
SQL Server provides SUBSTRING().
