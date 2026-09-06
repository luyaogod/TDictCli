---
title: "json_null attribute"
source: "fgl-topics/c_gws_json_null_attribute.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > Implicit and explicit conversion > Null and empty value handling in json.Serializer and util.JSON > json_null attribute"
type: "concept"
---

# json_null attribute

> This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.

## Syntax

```
json_null = {"null" | "undefined"}
```

1. `"null"`: writes null or empty BDL elements as the JSON `null`
   keyword.
2. `"undefined"`: omits null or empty BDL elements from the JSON output (except for the
   root element).

## Usage

Use `json_null="null"` to represent null or empty BDL values using the JSON
`null` keyword during serialization. This also enables deserialization of JSON
`null` values into BDL variables.

Use `json_null="undefined"` to omit null or empty BDL elements from the resulting
JSON. The root element is never omitted, and null values inside arrays must still be written as
`null` to preserve index positions. For detailed behavior and examples, go to [JSON nulls and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.").

When you use [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."), null values are
rejected unless you enable the [`allowNullAsDefault` (deserialization)](3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.") or [`serializeNullAsDefault` (serialization)](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.") options, or define variables with
`json_null="null"`.

## Affects

The behavior of the `json_null` attribute differs between [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") and [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."):

- `util.JSON`: The attribute affects serialization (BDL to JSON) only. A missing root
  value is treated as `null`, and [schema-related attributes](3677-json-serializer-attributes.md) such as `JSONRequired` are ignored at runtime.
- `json.Serializer`: The attribute affects both serialization (BDL to JSON) and
  deserialization (JSON to BDL). A missing root value is considered an error, and schema-related
  attributes such as [JSONRequired](3683-jsonrequired.md "Specify properties that are required in a JSON schema.") are processed, validated,
  and enforced.

For a full comparison of `json_null` handling, go to
[json\_null handling differences](3667-json-null-handling-differences.md "This topic compares how json.Serializer and util.JSON interpret the json_null attribute during serialization and deserialization.").

## Example

To enable null value handling on individual fields:

```
TYPE myRecordType RECORD
  id INTEGER,
  name VARCHAR(50) ATTRIBUTES(json_null="null"),
  description VARCHAR(100) ATTRIBUTES(json_null="null")
END RECORD
```

For more information about null and undefined handling, go to [JSON nulls and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.").

## Related links

**Related concepts**  

[The json.Serializer class](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.")

[allowNullAsDefault](3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.")

[serializeNullAsDefault](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
