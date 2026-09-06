---
title: "Null and empty value handling in json.Serializer and util.JSON"
source: "fgl-topics/c_gws_JSONSerializer_null_handling.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > Implicit and explicit conversion > Null and empty value handling in json.Serializer and util.JSON"
type: "concept"
---

# Null and empty value handling in json.Serializer and util.JSON

> This topic compares how json.Serializer and util.JSON handle null and empty values during JSON serialization and deserialization.

For detailed comparisons and examples, refer to the tables below dedicated to null handling for
both APIs.

For more information about managing null with the `json_null` attribute, go to
[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.").

| Type | BDL value | JSON output | json.Serializer | util.json |
| --- | --- | --- | --- | --- |
| `BOOLEAN` | `NULL` | `null` | ERROR 1 | OK |
| `FGL-Integer/Number` | `NULL` | `null` | ERROR 1 | OK |
| `FGL-STRING` | `NULL` | `null` | ERROR 1 | OK |
| `DATE` | `NULL` | `null` | ERROR 1 | OK |
| `DATETIME` | `NULL` | `null` | ERROR 1 | OK |
| `FGL-Object``util.JSONObject` | `NULL` | {} | OK | OK |
| `FGL-Array``util.JSONArray` | `NULL` | [] | OK | OK |

**1.** Requires [json\_null attribute](3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") or [serializeNullAsDefault](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.")

| Type | JSON input | Result | json.Serializer | util.json |
| --- | --- | --- | --- | --- |
| `BOOLEAN` | `null` | `NULL` | ERROR 2 | OK |
| `FGL-Integer` | `null` | `NULL` | ERROR 2 | OK |
| `FGL-STRING` | `null` | `NULL` | ERROR 2 | OK |
| `DATE` | `null` | `NULL` | ERROR 2 | OK |
| `DATETIME` | `null` | `NULL` | ERROR 2 | OK |
| `FGL-Object``util.JSONObject` | `null` | - | ERROR 2 | OK |
| `FGL-Array``util.JSONArray` | `null` | - | ERROR 2 | OK |

**2.** Requires [json\_null attribute](3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") or [allowNullAsDefault](3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.") option.

## Related links

**Related concepts**  

[Best practices](3669-best-practices.md "Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.")

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")

## Child topics

- [json_null handling differences in json.Serializer and util.JSON](3667-json-null-handling-differences.md): This topic compares how json.Serializer and util.JSON interpret the json_null attribute during serialization and deserialization.
- [json_null attribute](3668-json-null-attribute.md): This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.
