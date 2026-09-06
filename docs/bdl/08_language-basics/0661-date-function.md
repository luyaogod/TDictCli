---
title: "DATE [()] [function]"
source: "fgl-topics/c_fgl_operators_DATE.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > DATE [()] [function]"
type: "concept"
---

# DATE [()] [function]

> The DATE() operator converts an expression to a DATE value.

## Syntax

```
DATE [(expr)]
```

1. expr is the expression to be converted to a
   date.

## Usage

`DATE()` converts a character string, an integer or datetime expression to a [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") value.

When expr is a character string expression, it must properly formatted by
following the datetime format settings like [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.").

If expr is an integer expression, it is used as
the number of days since December 31, 1899.

If you supply no operand, it returns a character representation of
the current date in the format "weekday month day year".

## Example

```
MAIN
  DISPLAY DATE ( 34000 )
  DISPLAY DATE ( "12/04/1978" )
  DISPLAY DATE ( CURRENT )
  DISPLAY DATE
END MAIN
```

## Related links

**Related concepts**  

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")
