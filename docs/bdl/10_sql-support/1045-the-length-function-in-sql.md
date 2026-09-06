---
title: "The LENGTH() function in SQL"
source: "fgl-topics/c_fgl_sql_programming_093.html"
breadcrumb: "SQL support > SQL programming > SQL portability > The LENGTH() function in SQL"
type: "concept"
---

# The LENGTH() function in SQL

> The semantics of the LENGTH() SQL function differs according to the database engine.

The SQL LENGTH() function must be used with care: Each database server has different semantics
for this function, regarding length and trailing blanks handling.

The BDL language provides a LENGTH() built-in function, which is different from the SQL LENGTH()
function, used in SQL statements. The LENGTH() function of BDL returns zero when the string
expression is NULL.

| Database Server Type | Function name | Counting unit | Significant trailing blanks for CHAR() columns | Return value when NULL | Related topic |
| --- | --- | --- | --- | --- | --- |
| IBM® Informix® | LENGTH(expr) | Octets | No | NULL | [Native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | LEN(expr) | Characters | No | NULL | [See details](1296-the-length-function.md) |
| Oracle® MySQL / MariadDB | LENGTH(expr) | Characters | No | NULL | [See details](1343-the-length-function.md) |
| Oracle Database Server | LENGTH(expr) | Characters | Yes | NULL | [See details](1395-the-length-function.md) |
| PostgreSQL | LENGTH(expr) | Characters | No | NULL | [See details](1450-the-length-function.md) |
| SQLite | LENGTH(expr) | Characters | Yes | NULL | [See details](1494-the-length-function.md) |
| Dameng® | LENGTH(expr) | Characters | Yes | NULL | [See details](1246-the-length-function.md) |

Search for LENGTH() usage in your SQL statements and review the code of the database-specific
function.
> **Tip:**
>
> Create a user-defined SQL function that implements the Informix SQL
> LENGTH() function. For example, with Oracle
> PL/SQL:
>
> ```
> CREATE OR REPLACE FUNCTION vlength(
>     value IN VARCHAR2
>     )
>     RETURN INTEGER
> AUTHID CURRENT_USER
> IS
> BEGIN
>     RETURN NVL(LENGTH(RTRIM(value)),0);
> END;
> /
> ```
