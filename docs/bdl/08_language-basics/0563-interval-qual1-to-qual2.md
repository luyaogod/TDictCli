---
title: "INTERVAL qual1 TO qual2"
source: "fgl-topics/c_fgl_datatypes_INTERVAL.html"
breadcrumb: "Language basics > Primitive Data types > INTERVAL qual1 TO qual2"
type: "concept"
---

# INTERVAL qual1 TO qual2

> The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.

## Syntax 1: year-month class interval

```
  INTERVAL YEAR[(precision)] TO MONTH
| INTERVAL YEAR[(precision)] TO YEAR
| INTERVAL MONTH[(precision)] TO MONTH
```

## Syntax 2: day-time class interval

```
  INTERVAL DAY[(precision)] TO FRACTION[(scale)]
| INTERVAL DAY[(precision)] TO SECOND
| INTERVAL DAY[(precision)] TO MINUTE
| INTERVAL DAY[(precision)] TO HOUR
| INTERVAL DAY[(precision)] TO DAY

| INTERVAL HOUR[(precision)] TO FRACTION[(scale)]
| INTERVAL HOUR[(precision)] TO SECOND
| INTERVAL HOUR[(precision)] TO MINUTE
| INTERVAL HOUR[(precision)] TO HOUR

| INTERVAL MINUTE[(precision)] TO FRACTION[(scale)]
| INTERVAL MINUTE[(precision)] TO SECOND
| INTERVAL MINUTE[(precision)] TO MINUTE

| INTERVAL SECOND[(precision)] TO FRACTION[(scale)]
| INTERVAL SECOND[(precision)] TO SECOND

| INTERVAL FRACTION TO FRACTION[(scale)]
```

1. precision defines the number of significant digits of the first qualifier, it
   must be an integer from 1 to 9. For `YEAR`, the default is 4. For all other time
   units, the default is 2. For example, `YEAR(5)` indicates that the
   `INTERVAL` can store a number of years with up to 5 digits.
2. scale defines the precision of the fractional part of a second, from 1 to 5 digits:
   - `FRACTION(1)` : `10-1 s` (deciseconds)
   - `FRACTION(2)` : `10-2 s` (centiseconds)
   - `FRACTION(3)` : `10-3 s` (milliseconds)
   - `FRACTION(4)` : `10-4 s` (100 microseconds)
   - `FRACTION(5)` : `10-5 s` (10 microseconds)

## Usage

The `INTERVAL` data type stores a span of time, the difference between two points
in time. It can also be used to store quantities that are measured in units of time, such as ages or
times.

The `INTERVAL` data type falls into two classes. These are mutually exclusive
because year and month are not fixed-length units of time, and therefore incompatible with
`INTERVAL` data types whose time units are smaller than month:

- Year-Time intervals store a span of years, months or both.
- Day-Time intervals store a span of days, hours, minutes, seconds and fraction of
  seconds, or a contiguous subset of those units.

`INTERVAL` variables are initialized to `NULL` in functions,
modules and globals.

`INTERVAL` variables can be assigned from [interval literals](0591-interval-literals.md "Interval literals define an interval value in an expression."), by using the `INTERVAL() q1 TO
q2`
notation:

```
DEFINE iv INTERVAL DAY(5) TO SECOND
LET iv = INTERVAL(-7634 14:23:55) DAY(5) TO SECOND
```

`INTERVAL` variables can be assigned from string literals, by using the format
`YYYY-MM` or `DD hh:mm:ss.fffff`, depending on the interval
class:

```
DEFINE iv INTERVAL DAY(5) TO SECOND
LET iv = "-7634 14:23:55"
```

The ISO 8601 duration format is also supported: For `INTERVAL` of class
year-month, the ISO 8601 string must be formatted as `[-]PnnnnnnnnnY[-nnM]`. For
`INTERVAL` of class day-second, the ISO 8601 string must be formatted as
"`[-]PnnnnnnnnnDTnnHnnMnn[.nnnnn]S`":

```
DEFINE iv_y9_m INTERVAL YEAR(9) TO MONTH
DEFINE iv_d5_f INTERVAL DAY(5) TO FRACTION
LET iv_y9_m = "-P7634Y10M"
LET iv_d5_f = "-P456DT14H23M55.999S"
```

For more details, read [Character string to date time types](0578-data-type-conversion-reference.md).

`INTERVAL` variables defined with a single time unit can be assigned from integer
values, by using the [`UNITS`](0669-units.md "The UNITS operator converts an integer to an interval.")
operator:

```
DEFINE iv INTERVAL SECOND(5) TO SECOND
LET iv = 567 UNITS SECOND
```

The
`INTERVAL` type is used for [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") computation.

Depending on the data type of the operands, `DATETIME` or
`DECIMAL`, the arithmetic operations give different resulting types:

| Left Operand Type | Operator | Right Operand Type | Result Type |
| --- | --- | --- | --- |
| `INTERVAL` | `*` | `DECIMAL` | `INTERVAL` |
| `INTERVAL` | `/` | `DECIMAL` | `INTERVAL` |
| `INTERVAL` | `-` | `INTERVAL` | `INTERVAL` |
| `INTERVAL` | `+` | `INTERVAL` | `INTERVAL` |
| `DATETIME` | `-` | `INTERVAL` | `DATETIME` |
| `DATETIME` | `+` | `INTERVAL` | `DATETIME` |
| `DATETIME` | `-` | `DATETIME` | `INTERVAL` |

The next example shows how to use `INTERVAL` with `DATETIME`
variables:

```
MAIN
  DEFINE iym1, iym2 INTERVAL YEAR TO MONTH,
         dt1, dt2 DATETIME YEAR TO MINUTE,
         diff INTERVAL DAY(5) TO MINUTE
  LET iym1 = "2342-4"
  LET iym2 = "-55-11"
  DISPLAY iym1 + iym2
  LET dt1 = CURRENT
  LET dt2 = "2010-12-24 00:00"
  LET diff = dt1 - dt2
  DISPLAY diff
  LET diff = INTERVAL(-7634 14:23) DAY(5) TO MINUTE
  DISPLAY diff
END MAIN
```

For example, in the expression above `DISPLAY iym1 + iym2`, both values are from
the same `INTERVAL` class, that is both are year-month, and the result of the
`DATETIME`+`INTERVAL` calculation is a `DATETIME`
value:

```
Result: DATETIME 2286-05 YEAR TO MONTH
```

`INTERVAL` values can be negative.

In order to check if an `INTERVAL` is negative, use the [`UNITS`](0669-units.md "The UNITS operator converts an integer to an interval.") operator, to produce an interval
constant for the comparison. Using numeric constants will not
work:

```
MAIN
    DEFINE start, end DATETIME YEAR TO SECOND
    DEFINE diff INTERVAL SECOND(9) TO SECOND
    LET start = CURRENT + 100 UNITS SECOND
    LET end   = CURRENT - 200 UNITS SECOND
    LET diff = end - start
    IF diff < 0 THEN
       DISPLAY "this will not display!"
    END IF
    IF diff < 0 UNITS SECOND THEN
       DISPLAY "negative interval"
    ELSE
       DISPLAY "positive interval"
    END IF
END MAIN
```

Data type conversion can be controlled by catching the runtime exceptions. For more details, see
[Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[DATE](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.")

[DATETIME qual1 TO qual2](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")

[Data type conversion reference](0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")
