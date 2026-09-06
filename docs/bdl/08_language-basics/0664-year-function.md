---
title: "YEAR() [function]"
source: "fgl-topics/c_fgl_operators_YEAR.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > YEAR() [function]"
type: "concept"
---

# YEAR() [function]

> The YEAR() operator extracts the year of a date time expression.

## Syntax

```
YEAR ( expr )
```

1. expr is a date / time expression.

## Usage

Returns an integer corresponding
to the year portion of its operand.

## Example

```
MAIN
  DISPLAY YEAR ( TODAY )
  DISPLAY YEAR ( CURRENT )
END MAIN
```

## Related links

**Related concepts**  

[MONTH() [function]](0665-month-function.md "The MONTH() operator extracts the month of a date time expression.")

[DAY() [function]](0666-day-function.md "The DAY() operator extracts the day of the month of a date time expression.")

[WEEKDAY() [function]](0667-weekday-function.md "The WEEKDAY() operator extracts the day of the week of a date time expression.")

[MDY() [function]](0668-mdy-function.md "The MDY() operator creates a DATE from month, day and year units.")
