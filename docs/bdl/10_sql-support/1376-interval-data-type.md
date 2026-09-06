---
title: "INTERVAL data type"
source: "fgl-topics/c_fgl_odiagora_041.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > INTERVAL data type"
type: "concept"
description: "Informix® Informix provides the INTERVAL data type to store a value that represents a span of time. INTERVAL types are divided into two classes: Intervals of year-month class such as INTERVAL YEAR(5) ..."
---

# INTERVAL data type

## Informix®

Informix provides the `INTERVAL` data type to store a value that represents a
span of time.

`INTERVAL` types are divided into two classes:

- Intervals of year-month class such as `INTERVAL YEAR(5) TO
  MONTH`
- Intervals of day-time class such as `INTERVAL DAY(9) TO SECOND`

`INTERVAL` columns can be defined with various time units, by specifying a
start and end qualifier. For example, you can define an interval to store a number of hours and
minutes with `INTERVAL HOUR(n) TO MINUTE`, where
n defines the maximum number of digits for the hours unit.

The values of Informix `INTERVAL` can be represented with a
character string literal, or as `INTERVAL()`
literals:

```
'-9834 15:45:12.345'  -- an INTERVAL DAY(6) TO FRACTION(3)
'7623-11'   -- an INTERVAL YEAR(9) TO MONTH
INTERVAL(18734:45) HOUR(5) TO MINUTE
INTERVAL(-7634-11) YEAR(5) TO MONTH
```

## ORACLE

Oracle® provides an
`INTERVAL` data type similar to Informix,
that can be of year-month or day-time class.

However, Oracle's intervals cannot be defined with a time units different from the two interval
classes. For example, you cannot define an `INTERVAL HOUR TO MINUTE` in Oracle.

> **Note:**
>
> The ORACLE `INTERVAL DAY TO SECOND(n)` contains the fractional part of seconds
> and therefore is equivalent to the Informix
> `INTERVAL DAY TO FRACTION(n)` type.

## Solution

Informix `INTERVAL YEAR(n) TO MONTH`
and `INTERVAL MONTH(n) TO MONTH` data is stored in Oracle `INTERVAL YEAR(n) TO
MONTH` columns.

Informix `INTERVAL DAY(n) TO
FRACTION(p)` data and other forms of this interval class such as `INTERVAL HOUR(n) TO
MINUTE` is stored in Oracle `INTERVAL DAY(n) TO SECOND(p)` columns.

| Informix data type | Oracle |
| --- | --- |
| `INTERVAL YEAR[(p)] TO MONTH` | `INTERVAL YEAR[(p)] TO MONTH` |
| `INTERVAL MONTH[(p)] TO MONTH` | `INTERVAL YEAR(p-1) TO MONTH` |
| `INTERVAL DAY[(p)] TO FRACTION(n)` | `INTERVAL DAY[(p)] TO SECOND(n)` |
| `INTERVAL HOUR[(p)] TO HOUR` | `INTERVAL DAY(p-1) TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO MINUTE` | `INTERVAL DAY(p-1) TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO SECOND` | `INTERVAL DAY(p-1) TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO FRACTION(n)` | `INTERVAL DAY(p-1) TO SECOND(n)` |
| `INTERVAL MINUTE[(p)] TO MINUTE` | `INTERVAL DAY(p-3) TO SECOND(0)` |
| `INTERVAL MINUTE[(p)] TO SECOND` | `INTERVAL DAY(p-3) TO SECOND(0)` |
| `INTERVAL MINUTE[(p)] TO FRACTION(n)` | `INTERVAL DAY(p-3) TO SECOND(n)` |
| `INTERVAL SECOND[(p)] TO SECOND` | `INTERVAL DAY(p-4) TO SECOND(n)` |
| `INTERVAL SECOND[(p)] TO FRACTION(n)` | `INTERVAL DAY(p-4) TO SECOND(n)` |
| `INTERVAL FRACTION[(p)] TO FRACTION` | `INTERVAL DAY(0) TO SECOND(n)` |

When interval types do not match exactly, the Oracle ODI driver makes the required adjustments.
For example, when using an `INTERVAL HOUR(9) TO MINUTE` program variable as SQL
parameter, the number of hours is converted to a number of days and hours, to fit into an Oracle
`INTERVAL DAY(8) TO SECOND(0)`.

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
>    such data formatting with [`LOAD`/`UNLOAD`](1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.") and the type name returned by the [`base.SqlHandle.getResultType(n)`](../15_library-reference/3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.") method. For example, when creating a table
>    from a program with a column defined as `INTERVAL HOUR(6) TO FRACTION(4)` type, the
>    ODI converts the type name to Oracle's `INTERVAL DAY(5) TO SECOND(4)`. As result, the
>    data format for `LOAD`/`UNLOAD` will be `ddddd
>    hh:mm:ss.ffff`, and the type name returned by
>    `base.SqlHandle.getResultType()` will be `"INTERVAL DAY(5) TO
>    FRACTION(4)"`, based on the native Oracle interval type.

The
`INTERVAL` types translation can be controlled with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.datatype.interval = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
