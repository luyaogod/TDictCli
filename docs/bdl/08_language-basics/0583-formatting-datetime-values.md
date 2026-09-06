---
title: "Formatting DATETIME values"
source: "fgl-topics/c_fgl_DataConversions_format_datetimes.html"
breadcrumb: "Language basics > Formatting data > Formatting DATETIME values"
type: "concept"
---

# Formatting DATETIME values

> Date-time values must be formatted when converted to strings.

## When does DATETIME formatting take place?

Datetime formatting occurs when converting a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") to a string, for example in a
`LET`, `DISPLAY` or `PRINT` instruction, and when
displaying datetime values in form fields.

By default, `DATETIME` values are formatted in the ISO
format:

```
yyyy-mm-dd hh:mm:ss.fffff
```

A `DATETIME` value can be formatted with the [`util.Datetime.format()`](../15_library-reference/3488-util-datetime-format.md "Formats a date/time value based on a specified format.")
method:

```
IMPORT util
MAIN
  DEFINE dt DATETIME YEAR TO SECOND
  LET dt = CURRENT
  DISPLAY util.Datetime.format(dt, "%Y-%m-%d %H:%M:%S")
END MAIN
```

This code example produces the following output:

```
2015-12-23 11:45:33
```

## Converting strings to DATETIME values

When a string represents a datetime value is ISO format, it can be directly converted to a
`DATETIME`:

```
DEFINE dt DATETIME YEAR TO FRACTION(5)
LET dt = "2015-12-24 11:34:56.82373"
```

If you need to convert a string that does not follow the ISO format, use the [`util.Datetime.parse()`](../15_library-reference/3491-util-datetime-parse.md "Converts a string to a DATETIME value based on a specified format.") method, by specifying a format
string:

```
DEFINE dt DATETIME YEAR TO MINUTE
LET dt = util.Datetime.parse( "2014-12-24 23:45", "%Y-%m-%d %H:%M" )
```

## Formatting symbols for DATETIME values

When formatting `DATETIME` values, the format-string of
the `util.Datetime.parse()` and `util.Datetime.format()` methods
consists of a set of place holders that represent the different parts of a datetime value
(year, month, day, hour, minute, second and fraction).

Table 1 shows the formatting symbols for `DATETIME` expressions. Any character different
from the placeholders described in this table is interpreted as a literal and will appear as-is in
the resulting string.

The calendar used for date formatting is the Gregorian calendar.

| Placeholder | Description |
| --- | --- |
| `%a` | The abbreviated name of the day of the weekWhen parsing a datetime string, `%a` and `%A` are equivalent to detect the name of the day of the week in abbreviated form or full day name. |
| `%A` | The full name of the day of the week |
| `%b` or `%h` | The abbreviated month nameWhen parsing a datetime string, `%b`/`%h` and `%B` are equivalent to detect the month name in abbreviated form or full month name. |
| `%B` | The full month name |
| `%c` | The date and time representation |
| `%C` | The century number (0-99) |
| `%D` | Equivalent to `%m/%d/%y` |
| `%d` | The day of month with 2 digits (01-31) |
| `%e` | The day of month with one or 2 digits (1-31) |
| `%F[n]` | The fractional part of a second (`1<n<5`, `%F=%F5`)`%F[n]` includes the dot: In the format string, there is not need to preceed this placeholder with the dot. |
| `%H` | The hour with 2 digits (00-23) |
| `%I` | The hour on a 12-hour clock (1-12) |
| `%y` | The year on 2 digits (91) |
| `%Y` | The year on 4 digits (1991) |
| `%m` | The month as 2 digits (01-12) |
| `%M` | The minutes (00-59) |
| `%n` | A newline character |
| `%p` | The locale's equivalent of AM or PM |
| `%r` | The 12-hour clock time. In the POSIX locale equivalent to `%I:%M:%S %p` |
| `%R` | Equivalent to `%H:%M` |
| `%S` | The seconds (00-59) |
| `%t` | A tab character |
| `%T` | Equivalent to `%H:%M:%S` |
| `%x` | The date, using the locale's date format |
| `%X` | The time, using the locale's time format |
| `%w` | The ordinal number of the day of the week (0-6), with Sunday = 0 |
| `%y` | The year within century (0-99) |
| `%Y` | The year, including the century (for example, 1991) |

| Formatting string | Datetime value | Corresponding string |
| --- | --- | --- |
| `%d/%m/%Y %H:%M:%S%F3` | `2011-10-24 11:23:45.98765` | `24/10/2011 11:23.45.987` |
| `%d/%m/%Y %H:%M` | `2011-10-24 11:23:45` | `24/10/2011 11:23` |
| `[%d %b %Y %H:%M:%S]` | `2011-10-24 11:23:45` | `[24 Oct 2011 11:23:45]` |
| `(%a.) %b. %d, %Y` | `1999-09-23` | `(Thu.) Sep. 23, 1999` |
