---
title: "DATE and DATETIME data types"
source: "fgl-topics/c_fgl_odiagmsv_005.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > DATE and DATETIME data types"
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

## Microsoft™ SQL Server

Microsoft SQL Server (2008+) provides the following
data type to store date and time data:

| Microsoft SQL Server data type | Description |
| --- | --- |
| `DATE` | for year, month, day storage |
| `DATETIME` | for year, month, day, hour, min, second, fraction(3) storage (from January 1, 1753 through December 31, 9999). Values are rounded to increments of .000, .003, or .007 seconds |
| `SMALLDATETIME` | for year, month, day, hour, minutes storage (from January 1, 1900, through June 6, 2079). Values with 29.998 seconds or lower are rounded down to the nearest minute; values with 29.999 seconds or higher are rounded up to the nearest minute |
| `TIME(n)` | for hour, minute, second and fraction(7) storage. Where n defines the precision of fractional seconds |
| `DATETIME2(n)` | for year, month, day, hour, minute, second and fraction(7) storage. Where n defines the precision of fractional seconds |
| `DATETIMEOFFSET(n)` | for year, month, day, hour, minute, second, fraction(7) and time zone information storage. Where n defines the precision of fractional seconds |

Like Informix, Microsoft SQL Server can convert quoted strings to `DATETIME` data. The
`CONVERT()` SQL function allows you to convert strings to dates.

Microsoft SQL Server does not allow direct arithmetic
operations on datetimes; the date handling SQL functions must be used instead
(`DATEADD` and `DATEDIFF`).

SQL Server provides equivalent functions for Informix
`YEAR()`, `MONTH()` and `DAY()`. Take care with the
`DAY()` function on SQL Server because it begins from January 1, 1900 while Informix begins from December 31, 1899.

| Informix | Microsoft SQL SERVER |
| --- | --- |
| SELECT day(0), month(0), year(0) FROM systables WHERE tabid=1; ------ ------ ------ 31 12 1899 | SELECT day(0), month(0), year(0) ----------- ----------- ----------- 1 1 1900 |

The SQL Server equivalent for Informix
`WEEKDAY()` is the `DATEPART(dw,date-value)`
function.

The weekday date part depends on the value set by `SET DATEFIRST n`,
which sets the first day of the week (1=Monday ... 7=Sunday (default)).

## Solution

Use the following conversion rules to map Informix date/time types to Microsoft SQL Server date/time types:

| Informix data type | Microsoft SQL Server data type |
| --- | --- |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME(0)` |
| `DATETIME HOUR TO SECOND` | `TIME(0)` |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` |
| `DATETIME YEAR TO MONTH` | `DATETIME2(0)` |
| `DATETIME YEAR TO DAY` | `DATETIME2(0)` |
| `DATETIME YEAR TO HOUR` | `DATETIME2(0)` |
| `DATETIME YEAR TO MINUTE` | `DATETIME2(0)` |
| `DATETIME YEAR TO SECOND` | `DATETIME2(0)` |
| `DATETIME YEAR TO FRACTION(n)` | `DATETIME2(n)` |

The `DATE` and
`DATETIME` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.date = { true | false }
dbi.database.dsname.ifxemul.datatype.datetime = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

SQL Server and Informix `DATE` data
type are equivalent and store year, month, day values.

SQL Server `TIME(n)` data type can be used to store Informix `DATETIME HOUR TO FRACTION(n)`, `DATETIME
HOUR TO SECOND`, `DATETIME HOUR TO MINUTE` values, and any other
`DATETIME` type with qualifiers `HOUR`, `MINUTE`,
`SECOND` and `FRACTION(n)`. Missing time parts
default to 00:00:00.0. For example, when using a `DATETIME MINUTE TO FRACTION(3)`
with the value of "45:23.999", the SQL Server `TIME(3)` value will be
"00:45:23.999".

Informix `DATETIME` values with any
qualifiers from `YEAR` to `FRACTION(5)` can be stored in SQL Server
`DATETIME2(n)` columns. Missing date or time parts default to 1900-01-01 00:00:00.0.
For example, when using a `DATETIME DAY TO MINUTE` with the value of "23 11:45", the
SQL Server `DATETIME(0)` value will be "1900-01-23 11:45:00".

SQL Server `DATETIMEOFFSET(n)` data type is not supported by
Genero: This type includes timezone information, that would be lost when fetching the value into a
`DATETIME` variable. If timezone information needs to be stored, use a dedicated
column for this data, or create a view on the base table, that splits the
`DATETIMEOFFSET` column into a `DATETIME2`
(`cast(colname as datetime2)`) and timezome data
(`datepart(tz, colname)`).

> **Important:**
>
> - When fetching a `TIME` or `DATETIME2` with a precision that is
>   greater than 5 (the `DATETIME` precision limit), the database interface will allocate
>   a buffer of `VARCHAR(16)` for the `TIME` and
>   `VARCHAR(27)` for the `DATETIME2` column. As a result, you can fetch
>   such data into a `CHAR` or `VARCHAR` variable.
> - Review the program logic if you are using the Informix `WEEKDAY()` function because SQL Server uses a different basis for the
>   days numbers ( Monday = 1 ).
> - Use the SQL Server's `GETDATE()` function to get the system current date.

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
