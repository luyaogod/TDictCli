---
title: "DATE and DATETIME data types"
source: "fgl-topics/c_fgl_odiagmys_005.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > DATE and DATETIME data types"
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

## Oracle® MySQL and MariaDB

MySQL and MariaDB provide the following data type to store date and time data:

| MySQL data type | Description |
| --- | --- |
| `DATE` | for year, month, day storage. |
| `TIME[(n)]` | for hour, minute, second and fraction of second storage. |
| `DATETIME[(n)]` | for year, month, day, hour, minute, second and fraction of second storage.The range for MySQL `DATETIME` values is `1000-01-01 00:00:00` to `9999-12-31 23:59:59.999999`. |
| `TIMESTAMP` | Similar to `DATETIME`, but with smaller value range.The range for MySQL `TIMESTAMP` values is `1970-01-01 00:00:01.000000` to `2038-01-19 03:14:07.999999`. |

Like Informix, MySQL can convert quoted strings to
datetime data based on the ISO datetime format (`YYYY-MM-DD hh:mm:ss`).

In MySQL, the result of an arithmetic expression involving `DATE` values is an
`INTEGER` representing a number of days.

> **Important:**
>
> When the `ALLOW_INVALID_DATES` option is set in the
> `sql_mode` parameter, the database engine will allow invalid date values to be
> inserted, resulting as `0000-00-00` values in the table. It is not recommended to use
> this option.

## Solution

Use the following conversion rules to map Informix date/time types to MySQL date/time types:

| Informix data type | MySQL data type |
| --- | --- |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME` |
| `DATETIME HOUR TO SECOND` | `TIME` |
| `DATETIME HOUR TO FRACTION(p)` | `TIME(p)` |
| `DATETIME YEAR TO MONTH` | `DATETIME` |
| `DATETIME YEAR TO DAY` | `DATETIME` |
| `DATETIME YEAR TO HOUR` | `DATETIME` |
| `DATETIME YEAR TO MINUTE` | `DATETIME` |
| `DATETIME YEAR TO SECOND` | `DATETIME` |
| `DATETIME YEAR TO FRACTION(p)` | `DATETIME(p)` |

The `DATE` and
`DATETIME` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.date = { true | false }
dbi.database.dsname.ifxemul.datatype.datetime = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

MySQL and Informix `DATE` data types
both store year, month, day values.

With the `ALLOW_INVALID_DATES` option in `sql_mode` parameter,
invalid date values can be inserted. Fetching an invalid date value (`0000-00-00`)
will result as `NULL` in the target `DATE` BDL variable.

MySQL `TIME[(n)]` data type can be used to store Informix `DATETIME HOUR TO FRACTION(n)`, `DATETIME
HOUR TO SECOND`, `DATETIME HOUR TO MINUTE` values, and any other
`DATETIME` type with qualifiers `HOUR`, `MINUTE`,
`SECOND` and `FRACTION(n)`. Missing time parts
default to 00:00:00.0. For example, when using a `DATETIME MINUTE TO FRACTION(3)`
with the value of "45:23.999", the MySQL `TIME(3)` value will be "00:45:23.999".

Informix `DATETIME` values with any
qualifiers from `YEAR` to `FRACTION(5)` can be stored in MySQL
`DATETIME[(n)]` columns. Missing date or time parts default to
1900-01-01 00:00:00.0. For example, when using a `DATETIME DAY TO MINUTE` with the
value of "23 11:45", the MySQL `DATETIME` value will be "1900-01-23 11:45:00".

> **Important:**
>
> MySQL version older than 5.6.4 do not support fractional part in
> `TIME` and `DATETIME`. If you try to store a `DATETIME
> q1 TO FRACTION(p)` with such an old server version,
> the fractional part is lost.

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
