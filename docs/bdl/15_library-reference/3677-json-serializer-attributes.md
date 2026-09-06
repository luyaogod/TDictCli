---
title: "json.Serializer attributes"
source: "fgl-topics/r_gws_JSON_attributes.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer attributes"
type: "reference"
description: "Important: Most of the attributes in this table apply only to json.Serializer . Attributes that also work with util.JSON , such as json_name and json_null , are identified in their descriptions. Table ..."
---

# json.Serializer attributes

> **Important:**
>
> Most of the attributes in this table apply only to [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."). Attributes that also work with [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data."), such as `json_name` and `json_null`, are identified in their descriptions.

| Attribute | Description |
| --- | --- |
| JSONAdditionalProperties = { true \| false } | Allows a record to accept JSON properties not explicitly listed in its schema definition. |
| JSONAllOf | Combine multiple record schemas into a single type using the JSONAllOf attribute in Genero BDL. |
| JSONEnum = '{ "enum-value" } [,...]' | Defines an explicit, typed list of acceptable values for a field and maps to the `enum` keyword in JSON Schema. |
| JSONOneOf | Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization. |
| JSONSelector | Identifies the variant member of a JSONOneOf record to use during serialization or deserialization. |
| JSONPattern = "pattern" | Specify a regular expression pattern for string values in a JSON schema. |
| JSONRequired | Specify properties that are required in a JSON schema. |
| json_null = {"null" \| "undefined"} | This attribute controls the representation of null or empty values. `json.Serializer` uses it during both serialization and deserialization, but `util.JSON` applies it only when serializing. |
| json_name="JSON element name" | This attribute maps a BDL variable to a JSON element name that cannot be represented as a valid BDL identifier, including names with spaces or special characters. This attribute is supported by both [`json.Serializer`](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and [`util.JSON`](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data."). |
