---
title: "USING"
source: "fgl-topics/c_fgl_operators_USING.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > USING"
type: "concept"
---

# USING

> The USING operator converts date and numeric values to a string based on a formatting mask.

## Syntax

```
expr USING format
```

1. expr is a language expression.
2. format is a string expression that defines the formatting mask to be used.

## Usage

The `USING` operator applies a formatting string to the left operand.

The left operand must be a valid date, integer or decimal number.

The format string can be any valid string expression using formatting characters as
described in
[Formatting numeric values](0581-formatting-numeric-values.md "Numeric values must be formatted when converted to strings.")
and
[Formatting DATE values](0582-formatting-date-values.md "Date values must be formatted when converted to strings.").

[`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") and [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") expressions cannot be
formatted with the `USING` operator. Use the `util.Datetime` and
`util.Interval` methods instead. For more details, see [Formatting DATETIME values](0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.") and [Formatting INTERVAL values](0584-formatting-interval-values.md "Interval values must be formatted when converted to strings.").

The `USING` operator has a low order of precedence:
if you use operators with a higher precedence, the resulting string
might not be what you are expecting.

For example, the `||` concatenation operator is evaluated
before `USING`. As a result:

```
LET x = a || b USING "format"
```

will
first concatenate `a` and `b`, then
apply the `USING` format.

To solve this issue, use parentheses around the `USING` expression:

```
LET x = a || (b USING "format")
```

## Example

```
MAIN
  DEFINE d DECIMAL(12,2)
  LET d = -12345678.91
  DISPLAY d USING "$-##,###,##&.&&"
  DISPLAY TODAY USING "yyyy-mm-dd"
END MAIN
```

## Related links

**Related concepts**  

[Date, numeric and monetary formats](../09_advanced-features/0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.")

[Using the Ming Guo date format](../09_advanced-features/0891-using-the-ming-guo-date-format.md "Genero BDL can be configured to use the The Ming Guo calendar.")

[Concatenate (||)](0632-concatenate.md "The || operator makes a string concatenation.")
