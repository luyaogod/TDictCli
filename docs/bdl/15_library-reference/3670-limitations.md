---
title: "Limitations"
source: "fgl-topics/c_gws_JSONSerializer_limitations.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > Implicit and explicit conversion > Limitations"
type: "concept"
---

# Limitations

> This section outlines the serialization features that are not supported or have limited support in json.Serializer.

The following are the limitations of `json.Serializer` regarding serialization features:

- The ISO 8601 format for `DATETIME` (with the `T` separator) is not
  supported for serialization, but it is supported for deserialization.
- Converting objects or arrays to scalar types (such as `INTEGER` or
  `STRING`) always fails.
- `json.Serializer` does not accept null values by default; it requires the use of
  the attribute [json\_null attribute](3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.").
- `BOOLEAN` values do not auto-cast from arbitrary string values, such as
  "toto".

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")
