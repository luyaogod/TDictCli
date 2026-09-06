---
title: "TIME() [function]"
source: "fgl-topics/c_fgl_operators_TIME.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > TIME() [function]"
type: "concept"
---

# TIME() [function]

> The TIME() operator returns a time part of the date time expression.

## Syntax

```
TIME [(datetime-expr)]
```

1. datetime-expr is a datetime expression.

## Usage

`TIME()` converts the time-of-day portion of its datetime
operand to a character string.

This operator converts a date time expression to a character string
representing the time-of-day part of its operand.

The format of the returned string is always `"hh:mm:ss"`.

If you supply no operand, it returns a character representation of the current time. You can use
the [`CURRENT`](0659-current-function.md "The CURRENT function operator returns the current system date and time.") operator to get a
datetime result of the current system time.

## Example

```
MAIN
  DISPLAY TIME ( CURRENT )
END MAIN
```

## Related links

**Related concepts**  

[Datetime expressions](0599-datetime-expressions.md "This section covers date-time expression evaluation rules.")
