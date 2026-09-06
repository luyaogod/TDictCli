---
title: "util.JSON.setDatetimeSerializationMode"
source: "fgl-topics/c_fgl_ext_util_JSON_setDatetimeSerializationMode.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods > util.JSON.setDatetimeSerializationMode"
type: "concept"
---

# util.JSON.setDatetimeSerializationMode

> Defines the JSON formatting mode for DATETIME values.

## Syntax

```
util.JSON.setDatetimeSerializationMode(
     s STRING
   )
```

1. s is a string value that contains the format specification. Values can be:
   - `"FGL"`: Legacy FGL `DATETIME` formatting (default).
   - `"RFC3339/UTC"`: RFC 3339 format, with UTC conversion and trailing
     `Z` indicator.
   - `"RFC3339/TZO"`: RFC 3339 format, keeping value as local date/time, followed by
     `±hh:mm` timezone offset.

## Usage

The `util.JSON.setDatetimeSerializationMode()` class method globally defines the
formatting mode for `DATETIME` values when using `util.JSON`
classes.

The `"RFC3339/*"` format modes can be used to conform to JSON standards, in order
to produce date/time strings that conform to the RFC 3339 specification.

> **Note:**
>
> RFC 3339 does not define a format specification for durations. RFC 3339 only defines the format
> for date/time values. For durations, JSON follows the ISO 8601 specification.

**Default FGL mode**

By default, the serialization format is the legacy `"FGL"` format, which is also
used when converting a `DATETIME` to a `STRING`, or when displaying a
`DATETIME` value to a form field.

**The RFC3339/TZO mode**

`DATETIME YEAR TO MINUTE/SECOND/FRACTION` values convert into RFC 3339 format as
local date/time, by adding the `T` separator and the `±hh:mm` time
zone offset (automatically computed from the current system `TZ` settings). Since
RFC-3339 requires a complete time part with seconds, then resulting string will always be in the
form `YYYY-MM-DDThh:mm:ss[.fffff]±hh:mm`: If the `DATETIME` lower
qualifier is `MINUTE`, the resulting string will be completed with
`:00` seconds to conform to RFC 3339.

`DATETIME YEAR TO HOUR` cannot be converted to RFC 3339, because it’s impossible
to compute a proper timezone offset, since there can be timezones with half-hour offsets like
`TZ=Asia/Kolkata (UTC+5:30)`. Consequently, such DATETIME is serialized by using the
default FGL format.

With `DATETIME HOUR TO MINUTE/SECOND/FRACTION`, the result will be a local time
without any timezone indicator, because the timezone offset cannot be computed when date part is
missing. The resulting string will be in the form `hh:mm[:ss[.fffff]]`, as a local
wall-clock time, depending on the precision of the last qualifier. `DATETIME HOUR TO
MINUTE` will produce `hh:mm` (accepted by ISO 8601, RFC 3339 defines only
date+time data format)

For `DATETIME YEAR TO DAY` (or `TO MONTH/YEAR`) with only date
info, the `±hh:mm` offset makes no sense, and is not RFC 3339 anyway (there must be a
time part)

`DATETIME` with partial date information (such as `DATETIME MONTH TO
MINUTE`) cannot have a valid RFC 3339 representation. For example,
`10-15T12:00` is ambiguous and invalid in RFC 3339. Consequently, it must be
serialized with the default FGL format rules, as a local date/time value.

**The RFC3339/UTC mode**

`DATETIME YEAR TO MINUTE/SECOND/FRACTION` values are converted to a UTC date/time
value according to system settings (`TZ`), and formatted by adding the
`T` separator between the date and the time parts, plus the trailing
`Z` indicator. The date/time value is converted from local date/time to UTC .

Other `DATETIME` types will be serialized in the default FGL format, because the
precision of the `DATETIME` is not sufficient to make proper UTC conversion. Even
`DATETIME YEAR TO HOUR` can’t be properly converted, because of half-hour timezones
like `+05:30`.

For more details, read also [BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

## Example

```
IMPORT util
MAIN
    DEFINE dt1 DATETIME YEAR TO SECOND
    DEFINE dt2 DATETIME YEAR TO MINUTE
    LET dt1 = DATETIME(2026-04-24 14:35:44) YEAR TO SECOND
    LET dt2 = DATETIME(2026-04-24 09:11) YEAR TO MINUTE
    DISPLAY "Timezone TZ=", fgl_getenv("TZ")
    CALL util.JSON.setDatetimeSerializationMode("FGL")
    DISPLAY "FGL (default) dt1  : ", util.JSON.stringify(dt1)
    DISPLAY "FGL (default) dt2  : ", util.JSON.stringify(dt2)
    CALL util.JSON.setDatetimeSerializationMode("RFC3339/TZO")
    DISPLAY "RFC3339/TZO dt1    : ", util.JSON.stringify(dt1)
    DISPLAY "RFC3339/TZO dt2    : ", util.JSON.stringify(dt2)
    CALL util.JSON.setDatetimeSerializationMode("RFC3339/UTC")
    DISPLAY "RFC3339/UTC dt1    : ", util.JSON.stringify(dt1)
    DISPLAY "RFC3339/UTC dt2    : ", util.JSON.stringify(dt2)
END MAIN
```

Output, with TZ=Europe/Paris:

```
Timezone TZ=Europe/Paris
FGL (default) dt1  : "2026-04-24 14:35:44"
FGL (default) dt2  : "2026-04-24 09:11"
RFC3339/TZO dt1    : "2026-04-24T14:35:44+02:00"
RFC3339/TZO dt2    : "2026-04-24T09:11:00+02:00"
RFC3339/UTC dt1    : "2026-04-24T12:35:44Z"
RFC3339/UTC dt2    : "2026-04-24T07:11:00Z"
```

## Related links

**Related concepts**  

[util.JSON.setIntervalSerializationMode](3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values.")
