---
title: "Formatting INTERVAL values"
source: "fgl-topics/c_fgl_DataConversions_format_intervals.html"
breadcrumb: "Language basics > Formatting data > Formatting INTERVAL values"
type: "concept"
---

# Formatting INTERVAL values

> Interval values must be formatted when converted to strings.

## When does INTERVAL formatting take place?

Interval formatting occurs when converting an [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") to a string, for example in a
`LET`, `DISPLAY` or `PRINT` instruction, and when
displaying interval values in form fields.

By default, `INTERVAL` values are formatted in the ISO format, and
depending on the interval class and type.

For example, an `INTERVAL YEAR(4) TO MONTH` will be formatted
as:

```
[+|-]yyyy-mm
```

For an `INTERVAL DAY(n) TO FRACTION(5)`, the default format
is:

```
[+|-]dddd hh:mm:ss.fffff
```

An `INTERVAL` value can be formatted with the [`util.Interval.format()`](../15_library-reference/3528-util-interval-format.md "Formats an interval value based on a specified format.")
method:

```
IMPORT util
MAIN
  DEFINE iv INTERVAL DAY(6) TO MINUTE
  LET iv = "-157 11:23"
  DISPLAY util.Interval.format(iv, "%d %H:%M")
END MAIN
```

This code example produces the following output:

```
-157 11:23
```

## Converting strings to INTERVAL values

When a string represents a interval value is ISO format, it can be directly converted
to a `INTERVAL`:

```
DEFINE iv INTERVAL HOUR(6) TO FRACTION(5)
LET iv = "20234:34:56.82373"
```

If you need to convert a string that does not follow the ISO format, use the [`util.Interval.parse()`](../15_library-reference/3529-util-interval-parse.md "Converts a string to an INTERVAL value based on a specified format.")
method, by specifying a format string:

```
DEFINE iv INTERVAL DAY(6) TO FRACTION(5)
LET iv = util.Interval.parse( "-7467 + 23:45:34.12345", "%d + %H:%M:%S%F5" )
```

## Formatting symbols for INTERVAL values

When formatting `INTERVAL` values, the format-string
of the `util.Interval.parse()` and `util.Interval.format()`
methods consists of a set of place holders that represent the different parts of a interval
value (year, month, day, hour, minute, second and fraction).

Table 1 shows the formatting symbols for `INTERVAL` expressions.
Any character different from the
placeholders described in this table is interpreted as a literal and will appear as-is in
the resulting string.

| Placeholder | Description |
| --- | --- |
| `%Y` | Years (0-999999999) |
| `%m` | Months (0-999999999 if highest INTERVAL qualifier is MONTH(n), 0-11 otherwise) |
| `%d` | Days (0-999999999) |
| `%H` | Hours (0-999999999 if highest INTERVAL qualifier is HOUR(n), 00-23 otherwise) |
| `%M` | Minutes (0-999999999 if highest INTERVAL qualifier is MINUTE(n), 00-59 otherwise) |
| `%S` | Seconds (0-999999999 if highest INTERVAL qualifier is SECOND(n), 00-59 otherwise) |
| `%F[n]` | The fractional part of a second, where n specifies the number of digits in the fractional part (1 to 5) |
| `%t` | A tab character |
| `%n` | A newline character |

| Format String | Interval value | Result string |
| --- | --- | --- |
| `%d days %H:%M` | `54561 11:23` | `54561 days 11:23` |
| `%d days %H:%M:%S%F5` | `54561 11:23:45.12345` | `54561 days 11:23:45.12345` |
| `[%Y years and %m months]` | `1023-03` | `[1023 years and 03 months]` |
