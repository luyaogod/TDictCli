---
title: "util.JSON.setIntervalSerializationMode"
source: "fgl-topics/c_fgl_ext_util_JSON_setIntervalSerializationMode.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods > util.JSON.setIntervalSerializationMode"
type: "concept"
---

# util.JSON.setIntervalSerializationMode

> Defines the JSON formatting mode for INTERVAL values.

## Syntax

```
util.JSON.setIntervalSerializationMode(
     s STRING
   )
```

1. s is a string value that contains the format specification. Values can be:
   - `"FGL"`: Legacy FGL `INTERVAL` formatting (default).
   - `"ISO8601"`: ISO 8601 duration formatting such as `"P560Y05M"` or
     `"-P416DT15H22M33.009S"`.

## Usage

The `util.JSON.setIntervalSerializationMode()` class method globally defines the
formatting mode for `INTERVAL` values when using `util.JSON`
classes.

> **Note:**
>
> RFC 3339 does not define a format specification for durations. RFC 3339 only defines the format
> for date/time values. For durations, JSON follows the ISO 8601 specification.

By default, the serialization format is the legacy `"FGL"`
format, which is also used when converting an `INTERVAL` to a
`STRING`, or when displaying an `INTERVAL` value to a form field.

The `"ISO8601"` format mode can be used to conform to JSON standards, in order to
produce duration strings that conform to the ISO 8601 specification:

- For `INTERVAL` types of class year-month, the
  `ISO8601` serialization produces strings as `[-]PnnnnnnnnnYnnM`.
- For `INTERVAL` types of class day-second, the
  `ISO8601` serialization produces strings as
  `[-]PnnnnnnnnnDTnnHnnMnn[.nnnnn]S`.

Month, hour, minute and seconds parts get a leading zero when only one digit is significant. For
example, with 36 years and 6 months you get `P36Y06M`, not `P36Y6M`.
Both are equivalent and valid ISO 8601.

The leading `+` sign is accepted by the JSON to FGL conversion, for example:
`+P36Y6M`. However, the serialization methods will no produce a leading
`+` for positive intervals.

Plus/minus signs must appear after P in unit parts (`P-44DT-9H30M5.678S`) are not
accepted (common JSON convention rejects this)

For more details, read also [BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

## Example

```
IMPORT util
MAIN
    DEFINE iv1 INTERVAL YEAR(5) TO MONTH
    DEFINE iv2 INTERVAL DAY(9) TO FRACTION(5)
    LET iv1 = INTERVAL(-788-05) YEAR(5) TO MONTH
    LET iv2 = INTERVAL(8000 14:44:56.99999) DAY(9) TO FRACTION(5)
    CALL util.JSON.setIntervalSerializationMode("FGL")
    DISPLAY "FGL (default) iv1  : ", util.JSON.stringify(iv1)
    DISPLAY "FGL (default) iv2  : ", util.JSON.stringify(iv2)
    CALL util.JSON.setIntervalSerializationMode("ISO8601")
    DISPLAY "ISO8601 iv1        : ", util.JSON.stringify(iv1)
    DISPLAY "ISO8601 iv2        : ", util.JSON.stringify(iv2)
END MAIN
```

Output:

```
FGL (default) iv1  : "  -788-05"
FGL (default) iv2  : "      8000 14:44:56.99999"
ISO8601 iv1        : "-P788Y05M"
ISO8601 iv2        : "P8000DT14H44M56.99999S"
```

## Related links

**Related concepts**  

[util.JSON.setDatetimeSerializationMode](3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values.")
