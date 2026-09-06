---
title: "DATE and DATETIME data types"
source: "fgl-topics/c_fgl_odiagsqt_005.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > DATE and DATETIME data types"
type: "concept"
description: "Informix® Informix provides two data types to store date and time information: DATE = for year, month and day storage. DATETIME = for year to fraction (1-5) storage. The DATE type is stored as an ..."
---

# DATE and DATETIME data types

## Informix®

Informix provides two data types to store date and
time information:

- `DATE` = for year, month and day storage.
- `DATETIME` = for year to fraction (1-5) storage.

The `DATE` type is stored as an `INTEGER` with the number of days
since 1899/12/31.

The `DATETIME` type can be defined with various time units, by specifying a start
and end qualifier. For example, you can define a datetime to store an hour-to-second time value with
`DATETIME HOUR TO SECOND`.

The values of Informix
`DATETIME` can be represented with a character string literal, or as
`DATETIME()`
literals:

```
'2017-12-24 15:45:12.345'  -- a DATETIME YEAR TO FRACTION(3)
'15:45'   -- a DATETIME HOUR TO MINUTE
DATETIME(2017-12-24 12:45) YEAR TO MINUTE
DATETIME(12:45:56.333) HOUR TO FRACTION(3)
```

Informix is able to convert quoted strings to
`DATE` / `DATETIME` data, if the string contains matching environment
parameters. The string to date conversion rules for `DATE` is defined by the DBDATE
environment variable. The string to datetime format for `DATETIME` is defined by the
GL\_DATETIME environment variable.
> **Note:**
>
> Within Genero programs, the string representation for
> `DATETIME` values is always ISO (`YYYY-MM-DD hh:mm:ss.fffff`)

Informix supports date arithmetic on
`DATE` and `DATETIME` values. The result of an arithmetic expression
involving dates/times is an `INTEGER` number of days when only DATE values are used,
and an `INTERVAL` value if a `DATETIME` is used in the expression.

Informix automatically converts an
`INTEGER` to a `DATE` when the integer is used to set a value of a
date column.

Informix provides the `CURRENT [ q1 TO
q2 ]` operator, to get the system date/time on the server where the current database
is located. When no qualifiers are specified, `CURRENT` returns a `DATETIME
YEAR TO FRACTION(3)`. Informix also supports the `SYSDATE` operator, which
returns the current system time as a `DATETIME YEAR TO FRACTION(5)`.
> **Note:**
>
> The
> `USEOSTIME` configuration parameter must be set to 1 in order to get the subsecond
> precision in `CURRENT` and `SYSDATE` operators. See Informix
> documentation for more details.

## SQLite

SQLite 3 does not have a native type for date/time storage, but you can use data/time type names
and functions based on the string representation of dates and times. The date/time values are stored
in the `TEXT` native type.

The date/time functions of SQLite are based on standard `DATE (YYYY-MM-DD)`,
`TIME (hh:mm:ss)` and `TIMESTAMP (YYYY-MM-DD hh:mm:ss)` concepts.

For maximum flexibility with other RDBMS SQL languages, SQLite allows you to define table columns
with your own type names. You can for example use the `SMALLDATETIME`,
`SMALLTIME`, `TIME(N)`, `DATETIME(N)` type names.

## Solution

All Informix - BDL date/time types can be stored in
SQLite date/time columns.

Since SQLite allows various data type names, the date/time type conversion rules define specific
type names such as `SMALLTIME`, `TINYDATETIME`, to map original Informix date/time types. This allows the SQLite ODI driver
and the fgldbsch tool detect the exact date/time type of a column. When a
`CREATE TABLE` statement in a BDL program uses `DATETIME HOUR TO
MINUTE`, it is mapped to a `SMALLTIME` by the ODI driver, and when extracting
the database schema, fgldbsch can recognized `SMALLTIME` as a BDL
/ Informix `DATETIME HOUR TO
MINUTE` column.

The storage format must follow the ISO date/time formatting style (`YYYY-MM-DD
hh:mm:ss.fffff`). Depending on the BDL date/time precision, some parts will be omitted. For
example a `DATETIME HOUR TO MINUTE` is stored as `hh:mm` (see
conversion table below for more details).

Use the following conversion rules to map Informix date/time types to SQLite date/time (pseudo) types:

| Informix data type | SQLite (pseudo data type) | Storage format |
| --- | --- | --- |
| `DATE` | `DATE` | `YYYY-MM-DD` |
| `DATETIME HOUR TO HOUR` | `SMALLTIME` | `hh:00` |
| `DATETIME HOUR TO MINUTE` | `SMALLTIME` | `hh:mm` |
| `DATETIME HOUR TO SECOND` | `TIME` | `hh:mm:ss` |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` | `hh:mm:ss.fffff` |
| `DATETIME MINUTE TO MINUTE` | `SMALLTIME` | `00:mm` |
| `DATETIME MINUTE TO SECOND` | `TIME` | `00:mm:ss` |
| `DATETIME MINUTE TO FRACTION(n)` | `TIME(n)` | `00:mm:ss.fffff` |
| `DATETIME SECOND TO SECOND` | `TIME` | `00:00:ss` |
| `DATETIME SECOND TO FRACTION(n)` | `TIME(n)` | `00:00:ss.fffff` |
| `DATETIME FRACTION TO FRACTION(n)` | `TIME(n)` | `00:00:00.fffff` |
| `DATETIME YEAR TO YEAR` | `TINYDATETIME` | `YYYY-01-01` |
| `DATETIME YEAR TO MONTH` | `TINYDATETIME` | `YYYY-MM-01` |
| `DATETIME YEAR TO DAY` | `TINYDATETIME` | `YYYY-MM-DD` |
| `DATETIME YEAR TO HOUR` | `SMALLDATETIME` | `YYYY-MM-DD hh:00` |
| `DATETIME YEAR TO MINUTE` | `SMALLDATETIME` | `YYYY-MM-DD hh:mm` |
| `DATETIME YEAR TO SECOND` | `DATETIME` | `YYYY-MM-DD hh:mm:ss` |
| `DATETIME YEAR TO FRACTION(n)` | `DATETIME(n)` | `YYYY-MM-DD hh:mm:ss.fffff` |
| `DATETIME MONTH TO MONTH` | `TINYDATETIME` | `1900-MM-01` |
| `DATETIME MONTH TO DAY` | `TINYDATETIME` | `1900-MM-DD` |
| `DATETIME MONTH TO HOUR` | `SMALLDATETIME` | `1900-MM-DD hh:00` |
| `DATETIME MONTH TO MINUTE` | `SMALLDATETIME` | `1900-MM-DD hh:mm` |
| `DATETIME MONTH TO SECOND` | `DATETIME` | `1900-MM-DD hh:mm:ss` |
| `DATETIME MONTH TO FRACTION(n)` | `DATETIME(n)` | `1900-MM-DD hh:mm:ss.fffff` |
| `DATETIME DAY TO DAY` | `TINYDATETIME` | `1900-01-DD` |
| `DATETIME DAY TO HOUR` | `SMALLDATETIME` | `1900-01-DD hh:00` |
| `DATETIME DAY TO MINUTE` | `SMALLDATETIME` | `1900-01-DD hh:mm` |
| `DATETIME DAY TO SECOND` | `DATETIME` | `1900-01-DD hh:mm:ss` |
| `DATETIME DAY TO FRACTION(n)` | `DATETIME(n)` | `1900-01-DD hh:mm:ss.fffff` |

The `DATE` and
`DATETIME` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.date = { true | false }
dbi.database.dsname.ifxemul.datatype.datetime = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

In SQL statements, `CURRENT [q1 TO q2]`
expressions are converted to SQLite `strftime('%Y-%m-%d %H:%M:%S','now')`. The SQLite
`'now'` option returns the current date/time in UTC, while the FGL runtime system
`CURRENT` instruction returns the current local time. Both values can be different.
Always consider using SQL parameters with program variables assigned by the `CURRENT`
instruction of Genero BDL, instead of using `CURRENT` instructions in SQL
statements.

When using the `CURRENT` keyword without qualifiers, the SQL
translator in ODI drivers will not convert this expression to a native equivalent, because it is
difficult to find the required date/time precision. Depending on the context, the Informix `CURRENT` expression will be converted to match
the target `DATETIME` or `DATE` type. In SQL statements, always use
qualifiers after the `CURRENT` keyword, or use a `DATETIME` variable
assigned with the current date/time set by program, and use this variable as SQL
parameter.

```
-- Next CURRENT keyword will not be converted!
SELECT COUNT(*) INTO cnt FROM customer WHERE creadate < CURRENT
-- Best practice:
DEFINE dtcurr DATETIME YEAR TO FRACTION(3)
LET drcurr = CURRENT -- OK: it's the built-in language CURRENT expression!
SELECT COUNT(*) INTO cnt FROM customer WHERE creadate < dtcurr
```

## Related links

**Related concepts**  

[Date/time literals in SQL statements](1033-date-time-literals-in-sql-statements.md "Good practices for date and time handling in SQL.")

[SQL LOAD and UNLOAD](1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.")
