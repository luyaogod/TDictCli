---
title: "json.Serializer options"
source: "fgl-topics/r_gws_jsonserializer_options_2.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization > json.Serializer options"
type: "reference"
---

# json.Serializer options

> Options for controlling serialization and deserialization behavior, set using the json.Serializer class.

## Global options

> **Important:**
>
> These options apply to all JSON write paths and affect both [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and [json.JSONWriter](../15_library-reference/3618-the-json-jsonwriter-class.md "The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source."). They do not work with [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data."); to configure equivalent
> serialization modes for `util.JSON`, use
> [util.JSON.setDatetimeSerializationMode](../15_library-reference/3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values.") and
> [util.JSON.setIntervalSerializationMode](../15_library-reference/3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values.").

| Option | Description | Affects |
| --- | --- | --- |
| datetimeSerializationMode = value | Controls the JSON output format for `DATETIME` values during serialization. | Serialization (BDL to JSON) |
| intervalSerializationMode = value | Controls the JSON output format for `INTERVAL` values during serialization. | Serialization (BDL to JSON) |

## Serializer-specific options

> **Important:**
>
> These options are supported only by [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and do not work with [json.JSONWriter](../15_library-reference/3618-the-json-jsonwriter-class.md "The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source.") or [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

| Option | Description | Affects |
| --- | --- | --- |
| allowImplicitConversion = value | Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another. | Deserialization (JSON to BDL) |
| allowNullAsDefault = value | Allow NULL values to be accepted during deserialization when the `json_null="null"` attribute is not explicitly specified. | Deserialization (JSON to BDL) |
| serializeNullAsDefault = value | Allow NULL values during serialization even when constraints are set, such as when `JSONRequired` is defined or when `json_null="null"` is not defined. | Serialization (BDL to JSON) |

## Related links

**Related concepts**  

[The json.Serializer class](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[The json.JSONWriter class](../15_library-reference/3618-the-json-jsonwriter-class.md "The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source.")

[Implicit and explicit JSON conversion in json.Serializer and util.JSON](../15_library-reference/3663-implicit-and-explicit-conversion.md "This topic provides an overview of implicit and explicit JSON conversion behavior in json.Serializer and util.JSON, highlighting the differences between strict and permissive conversion models.")

[Best practices](../15_library-reference/3669-best-practices.md "Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.")

[json.Serializer.getOption](../15_library-reference/3674-json-serializer-getoption.md "Returns the current value of a JSON serializer option.")

[json.Serializer.setOption](../15_library-reference/3675-json-serializer-setoption.md "Sets an option on the JSON serializer.")
