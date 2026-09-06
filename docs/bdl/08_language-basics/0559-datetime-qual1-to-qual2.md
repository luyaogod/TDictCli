---
title: "DATETIME qual1 TO qual2"
source: "fgl-topics/c_fgl_datatypes_DATETIME.html"
breadcrumb: "Language basics > Primitive Data types > DATETIME qual1 TO qual2"
type: "concept"
---

# DATETIME qual1 TO qual2

> The DATETIME data type stores date and time data with time units from the year to fractions of a second.

## Syntax

```
  DATETIME YEAR TO FRACTION [ ( scale ) ]
| DATETIME YEAR TO SECOND
| DATETIME YEAR TO MINUTE
| DATETIME YEAR TO HOUR
| DATETIME YEAR TO DAY
| DATETIME YEAR TO MONTH
| DATETIME YEAR TO YEAR
| DATETIME MONTH TO FRACTION [ ( scale ) ]
| DATETIME MONTH TO SECOND
| DATETIME MONTH TO MINUTE
| DATETIME MONTH TO HOUR
| DATETIME MONTH TO DAY
| DATETIME MONTH TO MONTH
| DATETIME DAY TO FRACTION [ ( scale ) ]
| DATETIME DAY TO SECOND
| DATETIME DAY TO MINUTE
| DATETIME DAY TO HOUR
| DATETIME DAY TO DAY
| DATETIME HOUR TO FRACTION [ ( scale ) ]
| DATETIME HOUR TO SECOND
| DATETIME HOUR TO MINUTE
| DATETIME HOUR TO HOUR
| DATETIME MINUTE TO FRACTION [ ( scale ) ]
| DATETIME MINUTE TO SECOND
| DATETIME MINUTE TO MINUTE
| DATETIME SECOND TO FRACTION [ ( scale ) ]
| DATETIME SECOND TO SECOND
| DATETIME FRACTION TO FRACTION [ ( scale ) ]
```

1. scale defines the precision of the fractional part of a second, from 1 to 5 digits:
   - `FRACTION(1)` : `10-1 s` (deciseconds)
   - `FRACTION(2)` : `10-2 s` (centiseconds)
   - `FRACTION(3)` : `10-3 s` (milliseconds)
   - `FRACTION(4)` : `10-4 s` (100 microseconds)
   - `FRACTION(5)` : `10-5 s` (10 microseconds)

## Usage

The `DATETIME` data type stores an instance in time, expressed as a calendar
date and time-of-day.

The qualifiers following the `DATETIME` keyword define the precision of the
`DATETIME` type. While many sort of datetime types can be defined with all
possible qualifier combinations, only a limited set of `DATETIME` types are
typical used in applications:

- `DATETIME HOUR TO MINUTE, DATETIME HOUR TO SECOND, DATETIME HOUR TO FRACTION(scale)`:
  To hold a time value.
- `DATETIME YEAR TO MINUTE, DATETIME YEAR TO SECOND, DATETIME YEAR TO FRACTION(scale)`:
  To hold a date with time value.

`DATETIME YEAR TO DAY` is equivalent to `DATE`, consider using
`DATE` instead.

The year of a `DATETIME` can range from 1 to 9999.

When the `FRACTION` qualifier is specified without a precision, the precision
defaults to 3.

`DATETIME` arithmetic is based on the [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") data type, and can be combined
with `DATE` values:

| Left Operand Type | Operator | Right Operand Type | Result Type |
| --- | --- | --- | --- |
| `DATETIME` | `-` | `DATETIME` | `INTERVAL` |
| `DATETIME` | `-` | `DATE` | `INTERVAL` |
| `DATETIME` | `-` | `INTERVAL` | `DATETIME` |
| `DATETIME` | `+` | `INTERVAL` | `DATETIME` |

`DATETIME` variables are initialized to `NULL` in functions,
modules and globals.

The [`CURRENT`](0659-current-function.md "The CURRENT function operator returns the current system date and time.") operator
provides current system
date/time:

```
DEFINE dt DATETIME YEAR TO SECOND
LET dt = CURRENT
```

`DATETIME` variables can be assigned with [datetime literals](0590-datetime-literals.md "Datetime literals define date/time value in an expression."), by using the
`DATETIME() q1 TO q2` notation:

```
DEFINE dt DATETIME YEAR TO SECOND
LET dt = DATETIME(2014-02-21 13:45:34) YEAR TO SECOND
```

`DATETIME` variables can be assigned from string literals in the 4GL legacy format
`YYYY-MM-DD hh:mm:ss.fffff`, or the RFC 3339 format (ISO 8601 sub-set) using the T
separator between the date and time part, with optional `+/-hh:ss` timezone offset or
`Z` UTC indicator:

```
DEFINE dt DATETIME YEAR TO FRACTION(5)
LET dt = "2012-10-05 11:34:56.99999"
LET dt = "2012-10-05T11:34:56.99999+02:00"
LET dt = "2012-10-05T09:34:56.99999Z"
```

When a `DATETIME` value is assigned from a string using the RFC 3339 format, and
the string ends with a "`+/-hh:mm`" timezone offset or `Z` UTC
indicator, the value is adjusted to the local timezone when assigned.

The RFC 3339 specification requires full time part including the seconds (with optional
fractions). When the target variable is a `DATETIME YEAR TO HOUR` or `DATETIME
YEAR TO MINUTE`, the last units of the source string will be
ignored:

```
DEFINE dt DATETIME YEAR TO MINUTE
LET dt = "2012-10-05T11:34:56.99999+02:00"
DISPLAY dt -- shows 2012-10-05 11:34
```

To read an RFS 3339 without data loss, use a
`DATETIME YEAR TO SECOND` or `DATETIME YEAR TO FRACTION(5)`. For more
details, read [Character string to date time types](0578-data-type-conversion-reference.md).

When converting a `DATETIME` to a string, the format `YYYY-MM-DD
hh:mm:ss.fffff` is used.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

A `DATETIME` value can be converted to a different `DATETIME`
(or `DATE`) with a different precision by using the [`EXTEND()`](0660-extend-function.md "The EXTEND() operator adjusts a date time value depending on the qualifier.") operator:

```
MAIN
    DEFINE dt1 DATETIME YEAR TO MONTH
    DEFINE dt2 DATETIME YEAR TO FRACTION(5)
    LET dt1 = CURRENT
    LET dt2 = EXTEND(dt1, YEAR TO FRACTION(5))
END MAIN
```

Datetime conversion functions are provided in the [`util.Datetime`](../15_library-reference/3486-the-util-datetime-class.md "The util.Datetime class provides DATETIME data-type related utility methods.") class, for example
to convert local datetime to UTC datetime values:

```
IMPORT util
MAIN
    DEFINE dt DATETIME YEAR TO FRACTION(5)
    LET dt = "2012-10-05 11:34:56.99999"
    DISPLAY util.Datetime.toUTC( dt )
END MAIN
```

## Related links

**Related concepts**  

[DATE](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.")

[INTERVAL qual1 TO qual2](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.")

[Data type conversion reference](0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")
