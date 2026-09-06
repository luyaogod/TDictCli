---
title: "util.Date.parse"
source: "fgl-topics/c_fgl_ext_util_Date_parse.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Date class > util.Date methods > util.Date.parse"
type: "concept"
---

# util.Date.parse

> Converts a string to a DATE value based on a format specification.

## Syntax

```
util.Date.parse(
  s STRING,
  format STRING
 )
  RETURNS DATE
```

1. s is the source string to be parsed.
2. format is the format specification (see [Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.")).

## Usage

The `util.Date.parse()` method parses a string based on a format specification, to
produce a [`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") value.

The format specification must be a combination of `dd`, `mm`,
`yyyy` place holders as used in the `USING` operator.

The method returns `NULL`, if the source string cannot be converted to a
`DATE` value based on the format specification.

For more details about the supported formats, see [Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.").

## Example

```
IMPORT util
MAIN
    DISPLAY util.Date.parse( "2014-03-15", "yyyy-mm-dd" )
END MAIN
```

## Related links

**Related concepts**  

[USING](../08_language-basics/0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.")
