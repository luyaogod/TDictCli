---
title: "EXTEND() [function]"
source: "fgl-topics/c_fgl_operators_EXTEND.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > EXTEND() [function]"
type: "concept"
---

# EXTEND() [function]

> The EXTEND() operator adjusts a date time value depending on the qualifier.

## Syntax

```
EXTEND ( dt-expr, qual1 TO qual2 [(scale)])
```

1. dt-expr is a date / time expression.
2. qual1 and qual2 define
   the date time qualifiers. The qualifiers can be a combination of `YEAR`,
   `MONTH`, `DAY`, `HOUR`, `MINUTE`,
   `SECOND` and `FRACTION[(scale)]` as in the
   definition of a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")
   type.
3. scale defines the precision of the fractional part of a second, from 1 to 5 digits:
   - `FRACTION(1)` : `10-1 s` (deciseconds)
   - `FRACTION(2)` : `10-2 s` (centiseconds)
   - `FRACTION(3)` : `10-3 s` (milliseconds)
   - `FRACTION(4)` : `10-4 s` (100 microseconds)
   - `FRACTION(5)` : `10-5 s` (10 microseconds)

## Usage

The `EXTEND()` operator is used to convert a date time expression to a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") value with a different
precision.

The possible qualifiers are the same than in a `DATETIME` data type
definition.

If the datetime qualifiers are not specified after in the `EXTEND()` expression,
the precision defaults to `YEAR TO FRACTION(3)`.

The expressions passed as first parameter must be a valid datetime value. If it is a character
string, it must consist of valid and unambiguous time-unit values and separators, but with these
restrictions:

- It cannot be a character string in date format, such as `"12/12/99"`.
- It cannot be an ambiguous numeric datetime value, such as `"05:06"` or
  `"05"`.
- It cannot be a time expression that returns an [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") value.

## Example

```
MAIN
    DEFINE dt_y2f5 DATETIME YEAR TO FRACTION(5)
    LET dt_y2f5 = DATETIME(2024-12-24 11:45:59.98765) YEAR TO FRACTION(5)
    DISPLAY dt_y2f5
    DISPLAY EXTEND(dt_y2f5) -- defaults to YEAR TO FRACTION(3)
    DISPLAY EXTEND(dt_y2f5,YEAR TO DAY)
    DISPLAY EXTEND(dt_y2f5,HOUR TO SECOND)
    DISPLAY EXTEND(dt_y2f5,HOUR TO FRACTION) -- defaults to FRACTION(3)
    DISPLAY EXTEND(dt_y2f5,HOUR TO FRACTION(5))
    DISPLAY EXTEND(dt_y2f5,SECOND TO FRACTION(5))
    DISPLAY EXTEND(dt_y2f5,FRACTION TO FRACTION(5))
END MAIN
```

Output:

```
2024-12-24 11:45:59.98765
2024-12-24 11:45:59.987
2024-12-24
11:45:59
11:45:59.987
11:45:59.98765
59.98765
.98765
```

## Related links

**Related concepts**  

[CURRENT [function]](0659-current-function.md "The CURRENT function operator returns the current system date and time.")
