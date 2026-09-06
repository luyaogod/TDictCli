---
title: "json_null handling differences in json.Serializer and util.JSON"
source: "fgl-topics/c_gws_JSONSerializer_json_null_handling_2.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization > Implicit and explicit conversion > Null and empty value handling in json.Serializer and util.JSON > json_null handling differences"
type: "concept"
---

# json_null handling differences in json.Serializer and util.JSON

> This topic compares how json.Serializer and util.JSON interpret the json_null attribute during serialization and deserialization.

The handling of the [json\_null attribute](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") attribute and schema-related
features differs significantly between [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") and
[json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."):

- **util.json**

  - The `json_null` attribute only affects serialization (BDL to JSON).
  - At the root level, a missing value is implicitly treated as `null`.
  - Ignores schema-related attributes such as `JSONRequired` at runtime.
- **json.Serializer**

  - The `json_null` attribute affects both serialization (BDL to JSON) and
    deserialization (JSON to BDL).
  - At the root level, a missing value is considered an error, even if
    `json_null="undefined"` is set.
  - Performs field presence validation and handles default values during deserialization.
  - Processes and enforces schema-related attributes such as `JSONRequired`.

For more information about managing null with the `json_null` attribute, go to
[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.").

## Recommended practice

When using `json.Serializer`:

- Avoid relying on [`json_null="undefined"`](../09_advanced-features/0959-nulls-and-empty-structures.md) at the root level.
- If the root value may be omitted, wrap the primitive type inside a `RECORD` to
  ensure proper control and prevent deserialization errors.

For detailed comparisons and examples, refer to the tables below that are dedicated to null
handling for both APIs.

| Value is NULL | Target type | json.Serializer | util.JSON | Notes |
| --- | --- | --- | --- | --- |
| (default) |  |  |  |  |
|  | Primitive type | Omitted | Omitted | `json.Serializer: ERROR` if root |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `{}` / `[]` | `{}` / `[]` | `json.Serializer: ERROR` if root |
| [json\_null="null"](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") |  |  |  |  |
|  | Primitive type | `null` | `null` |  |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `null` | `null` |  |
| [`json_null="undefined"`](../09_advanced-features/0959-nulls-and-empty-structures.md) |  |  |  |  |
|  | Primitive type | Omitted | Omitted | `undefined` at root is treated as `null``json.Serializer`: `ERROR` if root |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | Omitted | Omitted | `undefined` at root is treated as `null``json.Serializer`: `ERROR` if root |
| `JSONRequired` |  |  |  |  |
|  | Primitive type | ERROR | Omitted | `util.JSON` ignores attribute `JSONRequired` |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `{}` / `[]` | `{}` / `[]` | `util.JSON` ignores attribute `JSONRequired` |
| `JSONRequired + json_null="null"` |  |  |  |  |
|  | Primitive type | `null` | `null` | `util.JSON` ignores attribute `JSONRequired` |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `null` | `null` | `util.JSON` ignores attribute `JSONRequired` |
| `JSONRequired + json_null="undefined"` |  |  |  |  |
|  | Primitive type | ERROR | `null` | `util.JSON` ignores attribute `JSONRequired` |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | ERROR | `null` | `util.JSON` ignores attribute `JSONRequired` |

| Incoming Value | Target type | Attribute | json.Serializer | util.JSON | Notes |
| --- | --- | --- | --- | --- | --- |
| Omitted |  |  |  |  |  |
|  | Primitive type |  | OK | OK | ERROR if root |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` |  | OK | OK | ERROR if root |
| `null` |  |  |  |  |  |
|  | Primitive type |  | ERROR | OK |  |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` |  | ERROR | OK |  |
| Omitted |  |  |  |  |  |
|  | Primitive type | [json\_null="null"](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") | OK | OK | ERROR if root |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | [json\_null="null"](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") | OK | OK | ERROR if root |
| `{}` / `[]` | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` |  | OK | OK |  |
| `null` |  |  |  |  |  |
|  | Primitive type | [json\_null="null"](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") | OK | OK |  |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | [json\_null="null"](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") | OK | OK |  |
| Omitted |  |  |  |  |  |
|  | Primitive type | [`json_null = "undefined"`](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") | OK | OK | ERROR if root |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | [`json_null = "undefined"`](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") | OK | OK | ERROR if root |
| Omitted |  |  |  |  |  |
|  | Primitive type | `JSONRequired` | ERROR | OK | `util.JSON` ignores attribute `JSONRequired` |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `JSONRequired` | ERROR | OK | `util.JSON` ignores attribute `JSONRequired` |
| `null` |  |  |  |  |  |
|  | Primitive type | `JSONRequired` | ERROR | OK |  |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `JSONRequired` | ERROR | OK |  |
| Omitted |  |  |  |  |  |
|  | Primitive type | `JSONRequired + json_null="null"` | OK | OK |  |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `JSONRequired + json_null="null"` | OK | OK |  |
| Omitted |  |  |  |  |  |
|  | Primitive type | `JSONRequired + json_null = "undefined"` | ERROR | OK | `json.Serializer`: "undefined" should be treated as `json_null="null"` at the root level |
|  | `FGL-Object``util.JSONObject``FGL-Array``util.JSONArray` | `JSONRequired + json_null = "undefined"` | ERROR | OK | `json.Serializer`: "undefined" should be treated as `json_null="null"`at the root level |

## Related links

**Related concepts**  

[Best practices](../15_library-reference/3669-best-practices.md "Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.")

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

**Related reference**  

[json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
