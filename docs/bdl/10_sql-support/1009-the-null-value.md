---
title: "The NULL value"
source: "fgl-topics/c_fgl_sql_programming_nulls.html"
breadcrumb: "SQL support > SQL programming > SQL portability > The NULL value"
type: "concept"
---

# The NULL value

> Several considerations need to be taken regarding SQL columns allowing NULL values.

## SQL Basics about NULL

The SQL concept of `NULL` means "no value" or "undefined value". This is used in
SQL to insert or update table rows and omit values for some fields that are not mandatory.

The `NULL` value is not zero: Zero is a real, numeric and non-null value.

Depending on the database engine, an empty string is null, or non-null. For more details, see
[Empty strings and NULLs](1016-empty-strings-and-nulls.md "Depending on the context, an empty string ( '' ) can be considered as NULL or NOT NULL.").

## The [NOT] NULL column constraint

SQL table columns can be defined with the constraint `NULL` or `NOT
NULL` (to deny `NULL` values):

```
CREATE TABLE customer (
   cust_id INTEGER NOT NULL PRIMARY KEY,
   cust_name VARCHAR(50) NOT NULL,
   cust_creation DATETIME YEAR TO SECOND NOT NULL,
   cust_address VARCHAR(100) NULL,
   ...
)
```

When omitting the `NULL` or `NOT NULL` constraint in `CREATE
TABLE`, most database engines define the column as nullable. However, when no constraint is
specified, some database brands such as Microsoft™ SQL
Server deny `NULL` values in their default configuration. There is an SQL Server
database option for that: See [SQL table definition](1283-sql-table-definition.md) in the SQL Server adaptation
guide.

Mandatory information should always be defined with a `NOT NULL` constraint. For
example, in an SQL table storing customer information, columns such as the customer name, the row
creation date/timestamp, columns referencing other tables rows with foreign keys, and obviously the
primary key column, should be defined with the `NOT NULL` constraint. It is also a
good practice to deny `NULL` values for flag/option columns like a Yes/No status, or
gender information.

On the other hand, secondary data such as a comment field, or fields that can only be assigned
after the first row creation (such as update timestamps), should accept `NULL`
values.

## Comparing NULL values

When `NULL` values enter in the game, it is important to understand that in means
"no value".

When used in comparison, arithmetic expressions, the SQL engine can't do anything with
`NULL`. In such case, the expression usually evaluates to `NULL` as
well.
> **Important:**
>
> Do not use the `=` or `!=` operator to
> check for `NULL` values!

The standard SQL operator to check for `NULL` values is the `IS [NOT]
NULL`
operator.

```
SELECT COUNT(*) FROM customers WHERE cust_comment IS NULL
```

Furthermore, when manipulating non-null values, keep in mind what happens with an operator, when
`NULL` values can show up. In the next SQL example, without knowing the semantics of
the `!=` comparison operator, it's not obvious to guess if the last
`SELECT` statement is supposed to return only the row with value
200:

```
CREATE TABLE tab1 ( col1 INTEGER );
INSERT INTO tab1 VALUES (NULL);
INSERT INTO tab1 VALUES (100);
INSERT INTO tab1 VALUES (300);
SELECT SUM(col1) FROM tab1;
SELECT AVG(col1) FROM tab1;
SELECT * FROM tab1 WHERE col1 != 100;
```

To simplify the usage of SQL columns that can be `NULL`, the most database engines
provide SQL functions to replace a potential `NULL` value by a default, non-null value:

| Database Server Type | NVL() | ISNULL() | IFNULL() | COALESCE() |
| --- | --- | --- | --- | --- |
| IBM® Informix® | Yes | No | No | Yes |
| Microsoft SQL Server | No | Yes | No | Yes |
| Oracle® MySQL | No | No | Yes | Yes |
| MariadDB | Yes | No | Yes | Yes |
| Oracle Database Server | Yes | No | Yes | Yes |
| PostgreSQL | No | No | No | Yes |
| SQLite | No | No | Yes | Yes |
| Dameng® | Yes | Yes | No | Yes |

As you can see, the `COALESCE()` function, which is standard ANSI, is
supported by all database engine brands supported by Genero, and since it follows simple SQL
function syntax, it can be used in static SQL.

## Sorting NULL values

When sorting SQL rows with the `ORDER BY` clause on a column that accepts
`NULL` values, pay attention to the way the database engine considers nulls versus
non-null values when sorting.

For example, with Informix, when using `ORDER BY colname
[ASC]`, rows with `NULL` appear at the top of the result set. This is
different with PostgreSQL, where null values by default sort as if larger than other non-null
value.

| Database Server Type | NULL rows position (ASC) |
| --- | --- |
| IBM Informix | First |
| Microsoft SQL Server | First |
| Oracle MySQL | Last |
| MariadDB | First |
| Oracle Database Server | Last |
| PostgreSQL | Last |
| SQLite | First |

Consider using `NOT NULL` constraints for `ORDER BY` sort columns,
when the position of the rows with `NULL` values matters.
