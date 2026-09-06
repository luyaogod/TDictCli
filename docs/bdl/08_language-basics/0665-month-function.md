---
title: "MONTH() [function]"
source: "fgl-topics/c_fgl_operators_MONTH.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > MONTH() [function]"
type: "concept"
---

# MONTH() [function]

> The MONTH() operator extracts the month of a date time expression.

## Syntax

```
MONTH ( expr )
```

1. expr is a date / time expression.

## Usage

Returns a positive whole number between
1 and 12 corresponding to the month of its operand.

## Example

```
MAIN
  DISPLAY MONTH ( TODAY )
  DISPLAY MONTH ( CURRENT )
END MAIN
```

## Related links

**Related concepts**  

[YEAR() [function]](0664-year-function.md "The YEAR() operator extracts the year of a date time expression.")

[DAY() [function]](0666-day-function.md "The DAY() operator extracts the day of the month of a date time expression.")

[WEEKDAY() [function]](0667-weekday-function.md "The WEEKDAY() operator extracts the day of the week of a date time expression.")

[MDY() [function]](0668-mdy-function.md "The MDY() operator creates a DATE from month, day and year units.")
