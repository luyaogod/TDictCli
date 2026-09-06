---
title: "util.Interval.format"
source: "fgl-topics/c_fgl_ext_util_Interval_format.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Interval class > util.Interval methods > util.Interval.format"
type: "concept"
---

# util.Interval.format

> Formats an interval value based on a specified format.

## Syntax

```
util.Interval.format(
   t INTERVAL q1 TO q2,
   format STRING
)
  RETURNS STRING
```

1. t is the interval value to be formatted.
2. format is the format string, as described in [Formatting INTERVAL values](../08_language-basics/0584-formatting-interval-values.md "Interval values must be formatted when converted to strings.").

## Usage

The `util.Interval.format()` method formats an `INTERVAL`
value based on the format specification.

The format string must be a combination of place holders such as `%Y`,
`%m`, `%d`, as described in [Formatting INTERVAL values](../08_language-basics/0584-formatting-interval-values.md "Interval values must be formatted when converted to strings.").

If the source value is `NULL` the result will be
`NULL`.

## Example

```
IMPORT util
MAIN
    DEFINE iv INTERVAL DAY(6) TO MINUTE
    LET iv = "-157 11:23"
    DISPLAY util.Interval.format(iv, "%d %H:%M")
END MAIN
```

## Related links

**Related concepts**  

[util.Interval.parse](3529-util-interval-parse.md "Converts a string to an INTERVAL value based on a specified format.")
