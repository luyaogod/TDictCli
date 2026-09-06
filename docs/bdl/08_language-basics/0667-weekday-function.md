---
title: "WEEKDAY() [function]"
source: "fgl-topics/c_fgl_operators_WEEKDAY.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > WEEKDAY() [function]"
type: "concept"
---

# WEEKDAY() [function]

> The WEEKDAY() operator extracts the day of the week of a date time expression.

## Syntax

```
WEEKDAY ( expr )
```

1. expr is a date / time expression.

## Usage

Returns a positive whole number between
0 and 6 corresponding to the day of the week implied by its operand.

The
integer 0 (Zero) represents Sunday.

## Example

```
MAIN
  DISPLAY WEEKDAY( TODAY )
  DISPLAY WEEKDAY( CURRENT )
END MAIN
```

## Related links

**Related concepts**  

[YEAR() [function]](0664-year-function.md "The YEAR() operator extracts the year of a date time expression.")

[MONTH() [function]](0665-month-function.md "The MONTH() operator extracts the month of a date time expression.")

[DAY() [function]](0666-day-function.md "The DAY() operator extracts the day of the month of a date time expression.")

[MDY() [function]](0668-mdy-function.md "The MDY() operator creates a DATE from month, day and year units.")
