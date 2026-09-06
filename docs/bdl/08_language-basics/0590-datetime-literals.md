---
title: "Datetime literals"
source: "fgl-topics/c_fgl_literals_DATETIME.html"
breadcrumb: "Language basics > Literals > Datetime literals"
type: "concept"
---

# Datetime literals

> Datetime literals define date/time value in an expression.

## Syntax

```
DATETIME ( dtrep ) qual1 TO qual2[(scale)]
```

where qual1 can be one of:

```
  YEAR
  MONTH
  DAY
  HOUR
  MINUTE
  SECOND
  FRACTION
```

and qual2 can be one of:

```
  YEAR
  MONTH
  DAY
  HOUR
  MINUTE
  SECOND
  FRACTION
  FRACTION(1)
  FRACTION(2)
  FRACTION(3)
  FRACTION(4)
  FRACTION(5)
```

1. dtrep is the datetime value representation
   in normalized format (YYYY-MM-DD hh:mm:ss.fffff).
2. scale defines the number of significant digits
   of the fractions of a second.
3. qual1 and qual2 qualifiers
   define the precision of the `DATETIME` literal.

## Usage

A datetime literal is specified with the `DATETIME()` notation, and is typically
used in interval or datetime expressions, or to assign a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") variable.

In order to get the current date and time, use the [`CURRENT`](0659-current-function.md "The CURRENT function operator returns the current system date and time.") operator.

## Example

```
MAIN
  DEFINE d1 DATETIME YEAR TO SECOND
  DEFINE d2 DATETIME HOUR TO FRACTION(5)
  LET d1 = DATETIME( 2002-12-24 23:55:56 ) YEAR TO SECOND
  LET d2 = DATETIME( 23:44:55.34532 ) HOUR TO FRACTION(5)
END MAIN
```

## Related links

**Related concepts**  

[Datetime expressions](0599-datetime-expressions.md "This section covers date-time expression evaluation rules.")
