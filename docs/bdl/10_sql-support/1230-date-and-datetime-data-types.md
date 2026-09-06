---
title: "DATE and DATETIME data types"
source: "fgl-topics/c_fgl_odiagdmg_005.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > DATE and DATETIME data types"
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

## Dameng®

Dameng provides the following data type to store date and time data:

| Dameng data type | Description |
| --- | --- |
| `DATE` | for year, month, day storage |
| `TIME(p)` | for hour, minutes, seconds and fraction of second storage |
| `TIMESTAMP(n)` | for year, month, day, hour, minutes, seconds and fraction of second storage |

Like Informix, Dameng can convert quoted strings to dates, times or timestamps.

Date and time literals can be expressed in character strings using the following format:
`'yyyy-mm-dd'` for dates, `'hh:mm:ss[.fff...]'` for times and
`'yyyy-mm-dd hh:mm:ss[.fff...]'` for timestamps.

Dameng supports `DATE` arithmetics like Informix: It is possible to add or
substract a number of days to/from a date. When substracting two date expressions, the result is a
number of days. A date value can be represented by a simple integer, as a number of days since epoch
(zero = 1900-01-01)

## Solution

Use the following conversion rules to map Informix date/time types to Dameng date/time types:

| Informix data type | Dameng data type |
| --- | --- |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME(0)` |
| `DATETIME HOUR TO SECOND` | `TIME(0)` |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` |
| `DATETIME YEAR TO MONTH` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO DAY` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO HOUR` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO MINUTE` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO SECOND` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO FRACTION(n)` | `TIMESTAMP(n)` |

The `DATE` and
`DATETIME` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.date = { true | false }
dbi.database.dsname.ifxemul.datatype.datetime = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

Dameng and Informix `DATE` data type are equivalent and store year, month, day values.

Dameng `TIME(p)` data type can be used to store Informix `DATETIME HOUR TO
SECOND/MINUTE/FRACTION(p)` values. Missing time parts default to 00:00:00.

`DATETIME` values with any qualifiers from `YEAR` to
`FRACTION(5)` can be stored in Dameng `TIMESTAMP(p)` columns. Missing
date or time parts default to 1900-01-01 00:00:00.0. For example, when using a `DATETIME DAY
TO MINUTE` with the value of "23 11:45", the Dameng `TIMESTAMP` value will be
"1900-01-23 11:45:00.0".

> **Important:**
>
> - Literal `DATETIME` and `INTERVAL` expressions (i.e.
>   `DATETIME (1999-10-12) YEAR TO DAY)` are not converted.
> - It is strongly recommended that you use BDL variables in dynamic SQL statements, instead of
>   quoted strings representing `DATE` values. For
>   example:
>
>   ```
>   LET stmt = "SELECT ... FROM customer WHERE creat_date >'", adate,"'"
>   ```
>
>   is
>   not portable, use a question mark place holder instead and OPEN the cursor USING
>   adate:
>
>   ```
>   LET stmt = "SELECT ... FROM customer WHERE creat_date > ?"
>   ```
> - SQL Statements using expressions with `TODAY` / `CURRENT` /
>   `EXTEND` must be reviewed and adapted to the native syntax.

## CURRENT expressions

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

## Date/time SQL functions

Dameng provides two kind of date/time functions to get the current date and time: Function that
return the date/time with time zone offset, and date/time returning the local time without time zone
offset. You must use the second set of function to get the equivalent purpose of Informix
`TODAY` and `CURRENT` expressions.

| Informix | Dameng |
| --- | --- |
| `today` | `curdate()` |
| `current hour to second` | `localtime(0)` |
| `current hour to fraction(n)` | `localtime(n)` |
| `current year to second` | `localtimestamp(0)` |
| `current year to fraction(n)` | `localtimestamp(n)` |

## Related links

**Related concepts**  

[Date/time literals in SQL statements](1033-date-time-literals-in-sql-statements.md "Good practices for date and time handling in SQL.")
