---
title: "util.Datetime.parse"
source: "fgl-topics/c_fgl_ext_util_Datetime_parse.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.parse"
type: "concept"
---

# util.Datetime.parse

> Converts a string to a DATETIME value based on a specified format.

## Syntax

```
util.Datetime.parse(
  s STRING,
  format STRING
)
  RETURNS DATETIME q1 TO q2
```

1. s is the source string to be parsed.
2. format is the format specification (see [Formatting DATETIME values](../08_language-basics/0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.")).

## Usage

The `util.Datetime.parse()` method parses a string based on a format
specification, to produce a [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") value.

> **Note:**
>
> The `parse()` method produces a `DATETIME` value. However, since
> Genero BDL supports [implicit type conversion](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language."), it
> is possible to assign a [`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") variable
> with the value returned from `parse()`, as long as the `DATETIME`
> contains a date part.

The format specification must be a combination of place holders such as
`%Y`, `%m`, `%d`, etc.

The precision of the resulting `DATETIME` value depends on the format
specification. For example, when using `"%Y-%m-%d %H:%M"`, the resulting value will
be a `DATETIME YEAR TO MINUTE`.

The method returns `NULL`, if the source string cannot be converted to a
`DATETIME` value based on the format specification.

For more details about the supported formats, see [Formatting DATETIME values](../08_language-basics/0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.").

## Example

```
IMPORT util
MAIN
    DEFINE dt DATETIME YEAR TO MINUTE
    LET dt = util.Datetime.parse( "2014-12-24 23:45", "%Y-%m-%d %H:%M" )
    DISPLAY dt
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.format](3488-util-datetime-format.md "Formats a date/time value based on a specified format.")
