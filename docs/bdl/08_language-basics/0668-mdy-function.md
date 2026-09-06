---
title: "MDY() [function]"
source: "fgl-topics/c_fgl_operators_MDY.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > MDY() [function]"
type: "concept"
---

# MDY() [function]

> The MDY() operator creates a DATE from month, day and year units.

## Syntax

```
MDY ( expr1, expr2, expr3 )
```

1. expr1 is an integer representing the month
   (from 1 to 12).
2. expr2 is an integer representing the day (from
   1 to 28, 29, 30 or 31 depending on the month).
3. expr3 is an integer representing the year (four
   digits).

## Usage

The `MDY()` operator builds a date value with 3 integers representing the month,
day and year. The result is a [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.")
value.

This function is sensitive to the `C1` modifier of the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment variable, defining a Ming Guo date
format.

## Example

```
MAIN
  DISPLAY MDY ( 12, 3+2, 1998 )
END MAIN
```

## Related links

**Related concepts**  

[YEAR() [function]](0664-year-function.md "The YEAR() operator extracts the year of a date time expression.")

[MONTH() [function]](0665-month-function.md "The MONTH() operator extracts the month of a date time expression.")

[DAY() [function]](0666-day-function.md "The DAY() operator extracts the day of the month of a date time expression.")

[WEEKDAY() [function]](0667-weekday-function.md "The WEEKDAY() operator extracts the day of the week of a date time expression.")

[Using the Ming Guo date format](../09_advanced-features/0891-using-the-ming-guo-date-format.md "Genero BDL can be configured to use the The Ming Guo calendar.")
