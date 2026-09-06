---
title: "ORACLE INTERVAL types"
source: "fgl-topics/c_fgl_Migrate_to_400_oracle_interval.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > ORACLE® INTERVAL types"
type: "concept"
---

# ORACLE INTERVAL types

> Better support for all kind of FGL INTERVAL types data storage with native ORACLE INTERVAL types.

Since first versions of Genero FGL, the [interval](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") types `INTERVAL[(n)] YEAR TO MONTH` and `INTERVAL DAY[(n)] TO
FRACTION(f)` are stored respectively in ORACLE `INTERVAL YEAR[(n)] TO
MONTH` and `INTERVAL DAY[(n)] TO SECOND(f)`. These native Informix and Oracle
interval types match exactly.

However, before BDL version 4.00, other FGL interval types such as `INTERVAL MONTH TO
MONTH` or `INTERVAL HOUR TO MINUTE`, and even `INTERVAL DAY TO
SECOND` were converted to ORACLE `CHAR(50)`. As a result, with interval data
in an Oracle `CHAR(50)` column, it was not possible to exploit the interval data
directly in the ORACLE engine. For example (without casting), interval arithmetic is impossible with
`CHAR(50)` strings representing interval values.

> **Important:**
>
> Starting with Genero BDL 4.00, the conversion from Informix `INTERVAL`
> types/values to `CHAR(50)` is no longer supported, and requires a database schema
> change, if you are using Informix `INTERVAL` that do not match exactly ORACLE
> `INTERVAL` types.

Starting with Genero BDL 4.00, all FGL interval types of both year-month and
day-time class are now stored in native ORACLE interval types.

For example, an FGL `INTERVAL HOUR(6) TO MINUTE` is converted to an ORACLE
`INTERVAL DAY(5) TO SECOND(0)`. When using such FGL interval variables in SQL, the
hours are split into days, to fit into the corresponding native ORACLE internal type. When fetching
the interval data from ORACLE, the number of days are converted to a number of hours to fit into the
FGL interval variable.

> **Important:**
>
> 1. When extracting a database schema from an Oracle database with the [fgldbsch](../09_advanced-features/0792-understanding-database-schemas.md "Database schemas hold the definition of the database tables and columns.") tool, native Oracle interval
>    types are converted to the corresponding Genero interval types. However, since Oracle only has 2
>    native interval types, the original Informix interval type definition is lost, if it does not match
>    exactly. For example, in a `CREATE TABLE` executed by a program, an Informix
>    `INTERVAL HOUR(9) TO MINUTE` type is converted to Oracle's `INTERVAL DAY(8) TO
>    SECOND(0)`, which is extracted by fgldbsch as an `INTERVAL DAY(8)
>    TO SECOND` type. To keep the original Informix interval types, you need to define the exact
>    type codes in the .sch file.
> 2. Native SQL data types are used by several Genero BDL instructions and will impact the behavior,
>    such data formatting with [`LOAD`/`UNLOAD`](../10_sql-support/1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.") and the type name returned by the [`base.SqlHandle.getResultType(n)`](../15_library-reference/3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.") method. For example, when creating a table
>    from a program with a column defined as `INTERVAL HOUR(6) TO FRACTION(4)` type, the
>    ODI converts the type name to Oracle's `INTERVAL DAY(5) TO SECOND(4)`. As result, the
>    data format for `LOAD`/`UNLOAD` will be `ddddd
>    hh:mm:ss.ffff`, and the type name returned by
>    `base.SqlHandle.getResultType()` will be `"INTERVAL DAY(5) TO
>    FRACTION(4)"`, based on the native Oracle interval type.

For more details, see [INTERVAL data type](../10_sql-support/1376-interval-data-type.md).
