---
title: "UNITS"
source: "fgl-topics/c_fgl_operators_UNITS.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > UNITS"
type: "concept"
---

# UNITS

> The UNITS operator converts an integer to an interval.

## Syntax

```
expr UNITS qual
```

where qual can be one of:

```
  YEAR
  MONTH
  DAY
  HOUR
  MINUTE
  SECOND
  FRACTION(1-6)
```

1. expr is an integer expression.

## Usage

The `UNITS` operator converts an integer expression to an [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") value expressed in a single
unit of time that you specify after the `UNITS` keyword.

For the qualifiers `YEAR`, `MONTH`, `DAY`,
`HOUR` and `SECOND`, if the left-hand expression evaluates to a
decimal number, any fractional part is discarded before the `UNITS` operator is
applied. However, when using `UNITS FRACTION`, the expression can be a decimal number
where the integer part is interpreted as a number of seconds, and the decimal part as the fraction
of a
second:

```
MAIN
  DEFINE iv INTERVAL SECOND(9) TO FRACTION(5)
  LET iv = 76242.77999 UNITS FRACTION
  DISPLAY iv  -- Displays "     76242.77999"
END MAIN
```

The `UNITS` operator can be used to compare `INTERVAL` values. For
example, to check if an `INERVAL SECOND(9) TO SECOND` is
negative:

```
FUNCTION is_negative( iv INTERVAL SECOND(9) TO SECOND )
    RETURN (iv < 0 UNITS SECOND )
END FUNCTION
```

`UNITS` has
a higher precedence than any arithmetic or boolean operator. As a
result, a left-hand arithmetic expression that uses a `UNITS` operator
must be enclosed in parentheses. For example, `10 + 20 UNITS
MINUTES` will be evaluated as `10 + (20 UNITS MINUTES)` and
give a conversion error. It must be written `(10 + 20) UNITS
MINUTES` to get the expected result.

Because the difference between two [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") values is an integer count of days rather than an
`INTERVAL` data type, you might want to use the `UNITS` operator to
convert such differences explicitly to `INTERVAL`
values:

```
MAIN
  DEFINE d DATE
  LET d = TODAY + 200
  DISPLAY (d - TODAY) UNITS DAY
END MAIN
```

## Related links

**Related concepts**  

[DATETIME qual1 TO qual2](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")

[EXTEND() [function]](0660-extend-function.md "The EXTEND() operator adjusts a date time value depending on the qualifier.")
