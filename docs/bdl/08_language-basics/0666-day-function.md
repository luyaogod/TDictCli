---
title: "DAY() [function]"
source: "fgl-topics/c_fgl_operators_DAY.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > DAY() [function]"
type: "concept"
---

# DAY() [function]

> The DAY() operator extracts the day of the month of a date time expression.

## Syntax

```
DAY ( expr )
```

1. expr is a date / time expression.

## Usage

Returns a positive whole number between
1 and 31 corresponding to the day of the month of its operand.

## Example

```
MAIN
  DISPLAY DAY ( TODAY )
  DISPLAY DAY ( CURRENT )
END MAIN
```

## Related links

**Related concepts**  

[YEAR() [function]](0664-year-function.md "The YEAR() operator extracts the year of a date time expression.")

[MONTH() [function]](0665-month-function.md "The MONTH() operator extracts the month of a date time expression.")

[WEEKDAY() [function]](0667-weekday-function.md "The WEEKDAY() operator extracts the day of the week of a date time expression.")

[MDY() [function]](0668-mdy-function.md "The MDY() operator creates a DATE from month, day and year units.")
