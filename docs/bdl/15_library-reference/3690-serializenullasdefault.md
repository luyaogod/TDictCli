---
title: "serializeNullAsDefault"
source: "fgl-topics/c_gws_JSONSerializer_option_serializeNullAsDefault.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > serializeNullAsDefault"
type: "concept"
---

# serializeNullAsDefault

> Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null="null" is not defined.

## Syntax

```
serializeNullAsDefault = value
```

The value parameter uses the [BDL
boolean](../08_language-basics/0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE."): set to `1` to allow `NULL` values even when
constraints such as `JSONRequired` apply, or set to `0` to disallow
them (default).

`serializeNullAsDefault` is an optional option.

## Usage

Use this option when you want `NULL` values to be accepted during serialization
even if constraints are defined, such as [`JSONRequired`](3683-jsonrequired.md "Specify properties that are required in a JSON schema."), or when [`json_null="null"`](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.") is not specified.

By default, `NULL` values are rejected during serialization if they violate
constraints like `JSONRequired`. Enabling this option makes the behavior more
permissive by allowing `NULL` values to pass through even when constraints are
present.

## Affects

Serialization (BDL to JSON)

## Example

To allow `NULL` values during serialization when constraints are applied:

```
CALL json.Serializer.setOption("serializeNullAsDefault", 1)
```

For more examples, go to [Serializing JSON nulls with serializeNullAsDefault](3693-serializenullasdefault-example.md "Use the serializeNullAsDefault option to accept BDL NULL values during JSON serialization.").

## Related links

**Related concepts**  

[The json.Serializer class](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[JSONRequired](3683-jsonrequired.md "Specify properties that are required in a JSON schema.")

[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
