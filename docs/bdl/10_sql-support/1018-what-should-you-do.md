---
title: "What should you do?"
source: "fgl-topics/c_fgl_sql_programming_069.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types > What should you do?"
type: "concept"
---

# What should you do?

> This section contains facts and tips to consider regarding character data types and locale settings.

Make sure that you have correctly defined the locale and length semantics for your character
string data types.

When designing your database tables, consider using `CHAR(N)` for fixed-length
string data (such as codes) and `VARCHAR(N)` for variable-length string data,
such as names, address and comments.

Use `VARCHAR` variables for `VARCHAR` columns, and
`CHAR` variables for `CHAR` columns, to achieve portability across all
kinds of database servers.

Note that with SQL Server, the column types can be `NCHAR/NVARCHAR` National
Character types, to Unicode character types to be combined with `CHAR/VARCHAR`
program variables. SQL Server 2019 surrports UTF-8 collation for `CHAR/VARCHAR`
columns.

Avoid storing empty strings in `VARCHAR` columns, or make sure that your program
is prepared to get nulls while the database stores empty strings.

Using byte or character length semantics depends mainly on the character set of your application:

- When using a single-byte character set, keep the default byte length semantics.
- When using a multibyte character set such as UTF-8, use character length semantics in both the
  database and the programs.

The database column definition and the program variable definition must match, this can be
simplified by using a database schema.

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")

[The LENGTH() function in SQL](1045-the-length-function-in-sql.md "The semantics of the LENGTH() SQL function differs according to the database engine.")
