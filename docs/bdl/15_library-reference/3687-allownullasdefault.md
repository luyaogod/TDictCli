---
title: "allowNullAsDefault"
source: "fgl-topics/c_gws_JSONSerializer_option_allowNullAsDefault.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > allowNullAsDefault"
type: "concept"
---

# allowNullAsDefault

> Allow NULL values to be accepted during deserialization when the json_null="null" attribute is not explicitly specified.

## Syntax

```
allowNullAsDefault = value
```

The value parameter is a BDL boolean (`1` to allow NULL values
without the `json_null="null"` attribute, or `0` to disallow them).
The default is `0`.

## Usage

This option allows `NULL` values to be accepted during deserialization when the [`json_null="null"`](../09_advanced-features/0959-nulls-and-empty-structures.md) attribute is not explicitly specified on the target field.

By default, `NULL` values are rejected during deserialization unless the
`json_null="null"` attribute is explicitly set on the affected field. Enabling
`allowNullAsDefault` provides a more permissive behavior by allowing
`NULL` values as defaults for fields without this attribute.

## Affects

Deserialization (JSON to BDL)

## Example

To allow NULL values during deserialization without requiring the `json_null="null"` attribute:

```
CALL json.Serializer.setOption("allowNullAsDefault", 1)
```

For more examples, go to [Deserializing JSON nulls with allowNullAsDefault](3692-allownullasdefault-example.md "Use the allowNullAsDefault option to accept JSON null values during JSON to BDL deserialization.")

## Related links

**Related concepts**  

[The json.Serializer class](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
