---
title: "Implicit and explicit JSON conversion in json.Serializer and util.JSON"
source: "fgl-topics/c_gws_JSONSerializer_implicit_explicit_casts.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > Implicit and explicit conversion"
type: "concept"
---

# Implicit and explicit JSON conversion in json.Serializer and util.JSON

> This topic provides an overview of implicit and explicit JSON conversion behavior in json.Serializer and util.JSON, highlighting the differences between strict and permissive conversion models.

Understanding the differences between [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") is important when
working with JSON conversion in Genero BDL. `json.Serializer` performs strict,
schema-driven conversion, while `util.JSON` applies permissive, dynamic conversion
rules.

The following sections compare implicit and explicit conversion behavior for primitive and
complex types, including type casting rules, failure scenarios, and the impact of options such as
[allowImplicitConversion](3686-allowimplicitconversion.md "Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another."). These sections also
describe how null and empty values are handled.

For additional options that influence conversion behavior, go to [json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

## Child topics

- [Implicit and explicit JSON conversion for primitive types](3664-primitive-type-conversion.md): This topic describes how primitive 4GL values are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.
- [Implicit and explicit JSON conversion for complex types](3665-complex-type-conversion.md): This topic describes how complex 4GL structures such as RECORD, DICTIONARY, arrays, and util.JSONObject are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.
- [Null and empty value handling in json.Serializer and util.JSON](3666-null-and-empty-value-handling-in-json-serializer-and-util-js.md): This topic compares how json.Serializer and util.JSON handle null and empty values during JSON serialization and deserialization.
- [Best practices](3669-best-practices.md): Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.
- [Limitations](3670-limitations.md): This section outlines the serialization features that are not supported or have limited support in json.Serializer.
