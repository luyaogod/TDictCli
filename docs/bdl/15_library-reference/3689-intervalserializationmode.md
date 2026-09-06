---
title: "intervalSerializationMode"
source: "fgl-topics/c_gws_JSONSerializer_option_intervalSerializationMode.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > intervalSerializationMode"
type: "concept"
---

# intervalSerializationMode

> Controls the JSON output format for INTERVAL values during serialization.

## Syntax

```
intervalSerializationMode = value
```

The value parameter is an [INTEGER](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers."):

- `0`: Genero default format (default)
- `1`: ISO 8601 duration format

A value outside the allowed range raises [error -15824](4483-genero-bdl-errors.md).

`intervalSerializationMode` is an optional option.

## Usage

Use this option to control how [INTERVAL](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") values
are written to JSON. By default (value `0`), the serializer writes
`INTERVAL` values in the Genero native format. Set to `1` to write in
ISO 8601 duration format:

- For year-month interval types, the output format is `[-]PnnnnnnnnnYnnM` (for example, `P36Y06M`).
- For day-second interval types, the output format is `[-]PnnnnnnnnnDTnnHnnMnn[.nnnnn]S` (for example, `P2DT10H30M00S`).

The leading field width reflects the declared interval precision (up to 9 digits); trailing fields are zero-padded to their minimum width.

Month, hour, minute, and second parts get a leading zero when only one digit is significant (for
example, `P36Y06M`, not `P36Y6M`).

For a description of the Genero native and ISO 8601 duration formats, see
[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

> **Note:**
>
> This option does not apply to `util.JSON`. For equivalent behavior,
> use [util.JSON.setIntervalSerializationMode](3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values.").

## Affects

Serialization (BDL to JSON)

## Example

To write `INTERVAL` values in ISO 8601 duration format:

```
CALL json.Serializer.setOption("intervalSerializationMode", 1)
```

For a complete example, see [Serializing DATETIME and INTERVAL values in standard formats](3695-datetime-and-interval-serialization-example.md "Use the datetimeSerializationMode and intervalSerializationMode options to output DATETIME values in RFC 3339 format and INTERVAL values in ISO 8601 duration format.").

## Related links

**Related concepts**  

[The json.Serializer class](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")

[Serializing DATETIME and INTERVAL values in standard formats](3695-datetime-and-interval-serialization-example.md "Use the datetimeSerializationMode and intervalSerializationMode options to output DATETIME values in RFC 3339 format and INTERVAL values in ISO 8601 duration format.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
