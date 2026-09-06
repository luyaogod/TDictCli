---
title: "Using portable data types"
source: "fgl-topics/c_fgl_sql_programming_061.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Using portable data types"
type: "concept"
---

# Using portable data types

> Only a limited set of data types are really portable across several database engines.

The ANSI SQL specification defines standard data types, but for historical
reasons most databases vendors have implemented native (non-standard) data
types. You can usually use a synonym for ANSI types, but the database server
will uses the native types behind the scenes. For example, when you create
a table with an `INTEGER` column in Oracle, the native
`NUMBER` data type is used.

In your programs, avoid data types that do not have a native equivalent
in the target database. This includes simple types like floating point numbers,
as well as complex data types like `INTERVAL`. Numbers may
cause rounding or overflow problems, because the values stored in the database
have different limits. For the `DECIMAL` types, always use the
same precision and scale for the program variables and the database columns.

To write portable applications, we strongly recommend using the following data types only:

- `CHAR(n)`
- `VARCHAR(n)`
- `BIGINT`
- `INTEGER`
- `SMALLINT`
- `DECIMAL(p,s)`
- `DATE`
- `DATETIME HOUR TO MINUTE`
- `DATETIME HOUR TO SECOND`
- `DATETIME HOUR TO FRACTION(n)`
- `DATETIME YEAR TO MINUTE`
- `DATETIME YEAR TO SECOND`
- `DATETIME YEAR TO FRACTION(n)`

Only if the target DB engine provides an equivalent native SQL type:

- `TEXT`/`BYTE` (for LOBs)
- `INTERVAL YEAR(p) TO MONTH`
- `INTERVAL DAY(p) TO MINUTE`
- `INTERVAL DAY(p) TO SECOND`
- `INTERVAL DAY(p) TO FRACTION(n)`

| Database Server Type | Data type topic |
| --- | --- |
| IBM® Informix® | Genero BDL is based on Informix SQL data types... |
| Microsoft™ SQL Server | [SQL types mapping: SQL Server](1271-sql-types-mapping-sql-server.md) |
| Oracle® MySQL / MariadDB | [SQL types mapping: Oracle MySQL](1325-sql-types-mapping-oracle-mysql.md) |
| Oracle Database Server | [SQL types mapping: Oracle database](1371-sql-types-mapping-oracle-database.md) |
| PostgreSQL | [SQL types mapping: PostgreSQL](1427-sql-types-mapping-postgresql.md) |
| SQLite | [SQL types mapping: SQLite](1477-sql-types-mapping-sqlite.md) |
| Dameng® | [SQL types mapping: Dameng](1226-sql-types-mapping-dameng.md) |

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
