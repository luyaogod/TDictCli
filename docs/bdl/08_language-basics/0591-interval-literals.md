---
title: "Interval literals"
source: "fgl-topics/c_fgl_literals_INTERVAL.html"
breadcrumb: "Language basics > Literals > Interval literals"
type: "concept"
---

# Interval literals

> Interval literals define an interval value in an expression.

## Syntax 1: *year-month* class interval

```
 INTERVAL ( inrep ) YEAR[(precision)] TO MONTH
|INTERVAL ( inrep ) YEAR[(precision)] TO YEAR
|INTERVAL ( inrep ) MONTH[(precision)] TO MONTH
```

## Syntax 2: *day-time* class interval

```
 INTERVAL ( inrep ) DAY[(precision)] TO FRACTION[(scale)]
|INTERVAL ( inrep ) DAY[(precision)] TO SECOND
|INTERVAL ( inrep ) DAY[(precision)] TO MINUTE
|INTERVAL ( inrep ) DAY[(precision)] TO HOUR
|INTERVAL ( inrep ) DAY[(precision)] TO DAY

|INTERVAL ( inrep ) HOUR[(precision)] TO FRACTION[(scale)]
|INTERVAL ( inrep ) HOUR[(precision)] TO SECOND
|INTERVAL ( inrep ) HOUR[(precision)] TO MINUTE
|INTERVAL ( inrep ) HOUR[(precision)] TO HOUR

|INTERVAL ( inrep ) MINUTE[(precision)] TO FRACTION[(scale)]
|INTERVAL ( inrep ) MINUTE[(precision)] TO SECOND
|INTERVAL ( inrep ) MINUTE[(precision)] TO MINUTE

|INTERVAL ( inrep ) SECOND[(precision)] TO FRACTION[(scale)]
|INTERVAL ( inrep ) SECOND[(precision)] TO SECOND

|INTERVAL ( inrep ) FRACTION TO FRACTION[(scale)]
```

1. inrep is the representation of the interval value in normalized format
   (`YYYY-MM` or `DD hh:mm:ss.fffff`, depending on the interval
   class).
2. precision defines the number of significant digits of the first qualifier, it
   must be an integer from 1 to 9. For `YEAR`, the default is 4. For all other time
   units, the default is 2. For example, `YEAR(5)` indicates that the
   `INTERVAL` can store a number of years with up to 5 digits.
3. scale defines the scale of the fractional part, it can be 1, 2, 3, 4 or
   5.

## Usage

An interval literal is specified with the `INTERVAL()` notation, and is typically
assigned in interval or datetime expressions, or to assign an [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") variable.

## Example

```
MAIN
  DEFINE i1 INTERVAL YEAR TO MONTH
  DEFINE i2 INTERVAL HOUR(5) TO SECOND
  LET i1 = INTERVAL( 345-5 ) YEAR TO MONTH
  LET i2 = INTERVAL( 34562:22:33 ) HOUR(5) TO SECOND
END MAIN
```

## Related links

**Related concepts**  

[Interval expressions](0600-interval-expressions.md "This section covers interval expression evaluation rules.")
