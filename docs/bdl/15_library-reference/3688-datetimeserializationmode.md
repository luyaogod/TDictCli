---
title: "datetimeSerializationMode"
source: "fgl-topics/c_gws_JSONSerializer_option_datetimeSerializationMode.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > datetimeSerializationMode"
type: "concept"
---

# datetimeSerializationMode

> Controls the JSON output format for DATETIME values during serialization.

## Syntax

```
datetimeSerializationMode = value
```

The value parameter is an [INTEGER](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers."):

- `0`: Genero default format (default)
- `1`: RFC 3339, normalized to UTC
- `2`: RFC 3339, with the local time zone offset

A value outside the allowed range raises [error -15824](4483-genero-bdl-errors.md).

`datetimeSerializationMode` is an optional option.

## Usage

Use this option to control how [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") values are written to JSON. By default (value `0`),
the serializer writes `DATETIME` values in the Genero native format. Set to
`1` to write in RFC 3339 format normalized to UTC, or to `2` to write
in RFC 3339 format with the local time zone offset of the server process.

When set to `2`, the offset reflects the time zone of the server process
at runtime. The same `DATETIME` value can produce different output on servers
configured with different time zones.

For a description of the Genero native and RFC 3339 formats, see
[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

> **Note:**
>
> This option does not apply to `util.JSON`. For equivalent behavior,
> use [util.JSON.setDatetimeSerializationMode](3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values.").

## Affects

Serialization (BDL to JSON)

## Example

To write `DATETIME` values in RFC 3339 format normalized to UTC:

```
CALL json.Serializer.setOption("datetimeSerializationMode", 1)
```

For a complete example, see [Serializing DATETIME and INTERVAL values in standard formats](3695-datetime-and-interval-serialization-example.md "Use the datetimeSerializationMode and intervalSerializationMode options to output DATETIME values in RFC 3339 format and INTERVAL values in ISO 8601 duration format.").

## Related links

**Related concepts**  

[The json.Serializer class](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")

[Serializing DATETIME and INTERVAL values in standard formats](3695-datetime-and-interval-serialization-example.md "Use the datetimeSerializationMode and intervalSerializationMode options to output DATETIME values in RFC 3339 format and INTERVAL values in ISO 8601 duration format.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
