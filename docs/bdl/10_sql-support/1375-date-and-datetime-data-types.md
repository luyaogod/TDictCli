---
title: "DATE and DATETIME data types"
source: "fgl-topics/c_fgl_odiagora_005.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > DATE and DATETIME data types"
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

## ORACLE

Oracle® provides the following data
types to store date and time data:

- `DATE` = for year, month, day, hour, min, second storage.
- `TIMESTAMP`= for year, month, day, hour, min, second, fraction storage.

As in Informix, Oracle can convert quoted strings to `DATE` or
`TIMESTAMP` data if the contents of the string matches the NLS date format parameters
(NLS\_DATE\_FORMAT, NLS\_TIMESTAMP\_FORMAT).

The `TO_DATE()` and `TO_TIMESTAMP()` SQL functions convert strings
to dates or timestamps, based on a given format. The `TO_CHAR()` SQL function allows
you to convert dates or timestamps to strings, according to a given format.

In Oracle the result of an arithmetic
expression involving `DATE` values is a number of days as `NUMBER`
type; the decimal part is the fraction of the day ( `0.5 = 12H00`, `2.00694444
= (2 + (10/1440)) = 2 days and 10 minutes` ). The result of an expression involving Oracle `TIMESTAMP` data is of
type `INTERVAL`.

To compare dates that have time data in Oracle, you can use the `ROUND()` or `TRUNC()` SQL
functions.

Even if the keyword is the same, the Oracle `SYSDATE` operator (using second precision) is not the exact equivalent
of the Informix `SYSDATE` operator (using
subsecond precision). Use Oracle's `SYSTIMESTAMP` as equivalent for Informix
`SYSDATE` operator.

See the Oracle documentation for more
details.

## Solution

Use the following conversion rules to map Informix date/time types to Oracle date/time types:

| Informix data type | Oracle |
| --- | --- |
| `DATE` | `DATE` |
| `DATETIME YEAR TO YEAR` | `DATE` |
| `DATETIME YEAR TO MONTH` | `DATE` |
| `DATETIME YEAR TO DAY` | `DATE` |
| `DATETIME YEAR TO HOUR` | `DATE` |
| `DATETIME YEAR TO MINUTE` | `DATE` |
| `DATETIME YEAR TO SECOND` | `DATE` |
| `DATETIME YEAR TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME MONTH TO MONTH` | `DATE` |
| `DATETIME MONTH TO DAY` | `DATE` |
| `DATETIME MONTH TO HOUR` | `DATE` |
| `DATETIME MONTH TO MINUTE` | `DATE` |
| `DATETIME MONTH TO SECOND` | `DATE` |
| `DATETIME MONTH TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME DAY TO DAY` | `DATE` |
| `DATETIME DAY TO HOUR` | `DATE` |
| `DATETIME DAY TO MINUTE` | `DATE` |
| `DATETIME DAY TO SECOND` | `DATE` |
| `DATETIME DAY TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME HOUR TO HOUR` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `DATE` |
| `DATETIME HOUR TO SECOND` | `DATE` |
| `DATETIME HOUR TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME MINUTE TO MINUTE` | `DATE` |
| `DATETIME MINUTE TO SECOND` | `DATE` |
| `DATETIME MINUTE TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME SECOND TO SECOND` | `DATE` |
| `DATETIME SECOND TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME FRACTION TO FRACTION(n)` | `TIMESTAMP(n)` |

The `DATE` and
`DATETIME` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.date = { true | false }
dbi.database.dsname.ifxemul.datatype.datetime = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Storing BDL DATE values

The Oracle `DATE` type
is used to store Genero BDL `DATE` values. However, keep in mind that the Oracle `DATE` type stores also
time (hh:mm:ss) information. The database interface automatically sets the time part to midnight
(00:00:00) during input/output operations.

You must be very careful since manual modifications of the database might set the time part, for
example:

```
UPDATE table SET date_col = SYSDATE
```

(`SYSDATE` is
equivalent to `CURRENT YEAR TO SECOND` in Informix).

After this type of update, when columns have date values with a time part different from
midnight, some `SELECT` statements might not return all the expected rows.

When fetching Oracle
`DATE` values into Genero BDL `DATE` or `DATETIME`
variables, the date and time information is directly set for the individual date/time parts and the
conversion is straight forward. But when fetching an Oracle `DATE` into a `CHAR` or
`VARCHAR` variable, date to string conversion occurs. Since Oracle dates are equivalent of Informix `DATETIME YEAR TO SECOND`, the values are by
default converted with the ISO format (YYYY-MM-DD hh:mm:ss), which is not the typical Informix behavior where dates are formatted from the DBDATE
environment variable. If your application fetches `DATE` values into
`CHAR/VARCHAR` and you want to get the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") conversion, you must set the following FGLPROFILE
entry:

```
dbi.database.dbname.ora.date.ifxfetch = true
```

> **Note:**
>
> Oracle does not support
> `INTEGER` to `DATE` automatic conversion.

## Storing BDL DATETIME values

Informix `DATETIME` data with any
precision from `YEAR` to `SECOND` is stored in Oracle `DATE` columns. The database interface
makes the conversion automatically. Missing date or time parts default to 1900-01-01 00:00:00. For
example, when using a `DATETIME HOUR TO MINUTE` with the value of "11:45", the Oracle `DATE` value will be
"1900-01-01 11:45:00".

Informix `DATETIME YEAR TO FRACTION(n)`
data is stored in Oracle
`TIMESTAMP` columns. The `TIMESTAMP` data type can store up to 9
digits in the fractional part, and therefore can store all precisions of Informix `DATETIME`.

> **Important:**
>
> - Most arithmetic expressions involving dates ( for example, to add or remove a number of days
>   from a date ) will produce the same result with Oracle. But keep in mind that Oracle
>   evaluates date arithmetic expressions to `NUMBER` ( days.fraction
>   ) while Informix evaluates to `INTEGER`
>   when only `DATE` values are used in the expression, or to `INTERVAL`
>   values if at least one `DATETIME` is used in the expression.
> - Even if a configuration parameter exists to get the Informix behavior, avoid fetching date values into `CHAR` or
>   `VARCHAR`, to bypass the DBDATE / ISO format conversion difference with Oracle.

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

Informix SQL provides a set of date/time functions
that do not exist in Oracle, but can be
easily implemented as user-defined SQL function in your Oracle database.

For example, to implement the `YEAR()` function, create the following user
function:

```
CREATE OR REPLACE FUNCTION year( adate in date )
RETURN NUMBER
IS
   v_year NUMBER;
BEGIN
   v_year := to_number( to_char( adate, 'YYYY' ) );
   RETURN (v_year);
END YEAR;
/
```

Similar SQL functions can be created, based on the following correspondance table:

| Informix | Oracle |
| --- | --- |
| today | trunc( sysdate ) |
| current year to second | sysdate |
| sysdate | systimestamp |
| day( value ) | to_number( to_char( value, 'dd' ) ) |
| extend( dtvalue, first to last ) | to_date( nvl( to_char( dtvalue, 'fmt-mask-1' ), '19000101000000' ), 'YYYYMMDDHH24MISS' ) |
| mdy( m, d, y ) | to_date( to_char(m,'09') \|\| to_char(d,'09') \|\| to_char(y,'0009'), 'MMDDYYYY' ) |
| month( value ) | to_number( to_char( value, 'mm' ) ) |
| weekday( value ) | to_number( to_char( value, 'd' ) ) - 1 |
| year( value ) | to_number( to_char( value, 'yyyy' ) ) |
| date( "string" \| integer ) | No equivalent: Depends from DBDATE with Informix |

## Related links

**Related concepts**  

[Date/time literals in SQL statements](1033-date-time-literals-in-sql-statements.md "Good practices for date and time handling in SQL.")
