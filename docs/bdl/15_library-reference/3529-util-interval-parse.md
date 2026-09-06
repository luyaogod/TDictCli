---
title: "util.Interval.parse"
source: "fgl-topics/c_fgl_ext_util_Interval_parse.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Interval class > util.Interval methods > util.Interval.parse"
type: "concept"
---

# util.Interval.parse

> Converts a string to an INTERVAL value based on a specified format.

## Syntax

```
util.Interval.parse(
  s STRING,
  format STRING
)
  RETURNS INTERVAL q1 TO q2
```

1. s is the source string to be parsed.
2. format is the format specification (see [Formatting INTERVAL values](../08_language-basics/0584-formatting-interval-values.md "Interval values must be formatted when converted to strings.")).

## Usage

The `util.Interval.parse()` method parses a string based on a format
specification, to produce an `INTERVAL` value.

The format specification must be a combination of place holders such as
`%Y`, `%m`, `%d`, etc.

The precision of the resulting `INTERVAL` value depends on the format
specification. For example, when using `"%Y-%m"`, the resulting value will be an
`INTERVAL YEAR(n) TO MONTH`.

The method returns `NULL`, if the source string cannot be converted to
an `INTERVAL` value based on the format specification.

For more details about the supported formats, see [Formatting INTERVAL values](../08_language-basics/0584-formatting-interval-values.md "Interval values must be formatted when converted to strings.").

## Example

```
IMPORT util
MAIN
    DEFINE iv1 INTERVAL YEAR(6) TO MONTH
    DEFINE iv2 INTERVAL DAY(6) TO FRACTION(5)
    LET iv1 = util.Interval.parse( "-999999-11", "%Y-%m" )
    DISPLAY iv1
    LET iv2 = util.Interval.parse( "-37467 + 23:45:34.12345", "%d + %H:%M:%S%F5" )
    DISPLAY iv2
END MAIN
```

## Related links

**Related concepts**  

[util.Interval.format](3528-util-interval-format.md "Formats an interval value based on a specified format.")
