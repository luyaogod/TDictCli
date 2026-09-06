---
title: "MDY(m,d,y) literals"
source: "fgl-topics/c_fgl_literals_MDY.html"
breadcrumb: "Language basics > Literals > MDY(m,d,y) literals"
type: "concept"
---

# MDY(m,d,y) literals

> MDY() literals define a DATE literal in an expression.

## Syntax

```
MDY( month-integer, day-integer, year-integer )
```

1. month-integer is an integer literal representing the month (1-12).
2. day-integer is an integer literal representing the day in the month (1-31).
3. year-integer is an integer literal representing a four-digit year.

## Usage

Date `MDY()` literals are used to define the initializer value for a [variable](0686-variables.md "Explains how to define program variables.") or the value of a [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of type [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.").

The `MDY()` literal is a different language elements than the [`MDY()` operator](0668-mdy-function.md "The MDY() operator creates a DATE from month, day and year units.").

## Example

```
MAIN
  DEFINE d DATE = MDY(12,24,2018)
  DISPLAY "Date = ", d
END MAIN
```

## Related links

**Related concepts**  

[Variable default values](0697-variable-default-values.md "Variables get a default value when defined.")
