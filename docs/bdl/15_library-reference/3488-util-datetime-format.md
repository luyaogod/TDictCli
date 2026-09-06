---
title: "util.Datetime.format"
source: "fgl-topics/c_fgl_ext_util_Datetime_format.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.format"
type: "concept"
---

# util.Datetime.format

> Formats a date/time value based on a specified format.

## Syntax

```
util.Datetime.format(
   t DATETIME q1 TO q2,
   format STRING
)
  RETURNS STRING
```

1. t is the date/time value to be formatted.
2. format is the format string, as described in [Formatting DATETIME values](../08_language-basics/0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.").

## Usage

The `util.Datetime.format()` method formats a [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") value based on the format
specification.

> **Note:**
>
> The `format()` method is designed to accept `DATETIME` values.
> However, since Genero BDL supports [implicit type
> conversion](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language."), it is also possible to pass a `DATE` value to the
> `format()` method.

The format string must be a combination of place holders such as `%Y`,
`%m`, `%d`, as described in [Formatting DATETIME values](../08_language-basics/0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.").

If the source value is `NULL` the result will be `NULL`.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Datetime.format( CURRENT, "%Y-%m-%d %H:%M" )
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.parse](3491-util-datetime-parse.md "Converts a string to a DATETIME value based on a specified format.")
