---
title: "DATE and DATETIME data types"
source: "fgl-topics/c_fgl_odiagpgs_005.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > DATE and DATETIME data types"
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

## PostgreSQL

PostgreSQL supports the following data types to store date/time values:

| PostgreSQL data type | Description |
| --- | --- |
| `DATE` | For year, month, day storage |
| `TIME [(p)] [WITHOUT TIME ZONE]` | For hour, minute, second and fraction of second storage (with `0<=p<=6`) |
| `TIME [(p)] WITH TIME ZONE` | For hour, minute, second and fraction of second storage (with `0<=p<=6`), applying time zone conversions. |
| `TIMESTAMP [(p)] [WITHOUT TIME ZONE]` | For year, month, day, hour, minute, second and fraction of second storage (with `0<=p<=6`) |
| `TIMESTAMP [(p)] WITH TIME ZONE` | For year, month, day, hour, minute, second and fraction of second storage (with `0<=p<=6`), applying time zone conversions. |

The result of an arithmetic expression involving PostgreSQL `DATE` values is an
`INTEGER` representing a number of days.

When the precision `p` is not specified in a PostgreSQL
`TIME`/`TIMESTAMP` type, the precision of the fractional part is not
fixed.

PostgreSQL can convert quoted strings to date/time data depending on the
`DateStyle` session parameter. PostgreSQL always accepts ISO date time strings. The
date format can be changed with `SET DateStyle` PostgreSQL command. However, it is
better to use the default ISO `DateStyle` format.

The `WITH TIME ZONE` option enables automatic storage of date/time data in UTC.
This option is not to store time zone offset along with the date/time value: It enables conversion
from/to UTC to/from a date/time in current PostgreSQL system TimeZone or a specific time zone: The
value stored in the PostgreSQL database is always in UTC. See PostgreSQL documentation for more
details.

## Solution

Use the following conversion rules to map Informix date/time types to PostgreSQL date/time types:

| Informix data type | PostgreSQL data type |
| --- | --- |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME(0) WITHOUT TIME ZONE` |
| `DATETIME HOUR TO SECOND` | `TIME(0) WITHOUT TIME ZONE` |
| `DATETIME HOUR TO FRACTION(p)` | `TIME(p) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO MONTH` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO DAY` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO HOUR` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO MINUTE` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO SECOND` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO FRACTION(p)` | `TIMESTAMP(p) WITHOUT TIME ZONE` |

When the PostgreSQL `TIME`/`TIMESTAMP` precision `p`
is not specified, the number of digits in the fractional part is variable. This prevents a proper
type mapping, for example when extracting column types with [fgldbsch](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") from a PostgreSQL database.
When creating table columns in PostgreSQL, always use
`TIME(p)`/`TIMESTAMP(p)` types, and specify 0 (zero) if there are no
fraction of seconds to be stored.

The `DATE` and
`DATETIME` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.date = { true | false }
dbi.database.dsname.ifxemul.datatype.datetime = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

PostgreSQL and Informix `DATE` types
are equivalent and store year, month, day values.

The PostgreSQL ODI driver can detect the current `DateSyle` used in the SQL
session. It can be ISO, SQL and German. The `DateStyle` `"Postgres"`
is not supported for `TIMESTAMP` types. Consider using the ISO default
`DateStyle` format.

PostgreSQL `TIME(n) WITHOUT TIME ZONE` data type can be used to store Informix `DATETIME HOUR TO FRACTION(n)`,
`DATETIME HOUR TO SECOND`, `DATETIME HOUR TO MINUTE` values, and any
other `DATETIME` type with qualifiers `HOUR`, `MINUTE`,
`SECOND` and `FRACTION(n)`. Missing time parts
default to 00:00:00.0. For example, when using a `DATETIME MINUTE TO FRACTION(3)`
with the value of "45:23.999", the PostgreSQL `TIME(3) WITHOUT TIME ZONE` value will
be "00:45:23.999".

Informix `DATETIME` values with any
qualifiers from `YEAR` to `FRACTION(5)` can be stored in PostgreSQL
`TIMESTAMP(n) WITHOUT TIME ZONE` columns. Missing date or time parts default to
1900-01-01 00:00:00.0. For example, when using a `DATETIME DAY TO MINUTE` with the
value of "23 11:45", the PostgreSQL `TIMESTAMP(0) WITHOUT TIME ZONE` value will be
"1900-01-23 11:45:00".

The `TIME WITH TIME ZONE` or `TIMESTAMP WITH TIME ZONE` PostgreSQL
data types must be used with care and are not recommended for SQL portability: When passing
date/time data from FGL variables for `TIME WITH TIME ZONE` or `TIMESTAMP WITH
TIME ZONE` columns, the value provided by the ODI driver has no time zone offset. However,
it is considered by PostgreSQL as a date/time in the current PostgreSQL system
`TimeZone` and will be converted to UTC before it is stored. When fetching the value
back from the SQL column into an FGL variable, PostgreSQL will automatically convert from UTC to a
date/time in the PostgreSQL system `TimeZone`. For example, insert of FGL value
`2023-08-16 17:45:33` in the time zone `GMT+1H` and with
`+1H` for Daylight Saving Time (DST), will be stored as UTC `2023-08-16
15:45:33` because `17H - 2H = 15H`. When fetching the value back, UTC
`2023-08-16 15:45:33` is converted to a date/time in the current system
`TimeZone`: `15H + 2H = 17H`. PostgreSQL adds the TZ offset
`+02` in the raw data string, but it is ignored by the ODI driver. As result, the FGL
variable value is set to `2023-08-16 17:45:33`. If the PostgreSQL system
`TimeZone` is changed, the value in the FGL variable will be different. If you need
to store date/time data in UTC, do the conversions at the BDL level before using the UTC date/time
values in SQL statements. See [`util.Datetime.toUTC()`](../15_library-reference/3494-util-datetime-toutc.md "Converts a date/time value to the UTC date/time.").

Complex `DATETIME` expressions (involving `INTERVAL` values for
example) are Informix-specific and have no equivalent in PostgreSQL.

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

[Type conversion rules](1452-type-conversion-rules.md "Type conversion rules")
