---
title: "allowImplicitConversion"
source: "fgl-topics/c_gws_JSONSerializer_option_allowImplicitConversion.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > allowImplicitConversion"
type: "concept"
---

# allowImplicitConversion

> Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another.

## Syntax

```
allowImplicitConversion = value
```

The value parameter is a BDL boolean (`1` to enable,
`0` to disable). The default is `0`.

## Usage

This option modifies deserialization behavior by allowing implicit type conversions when the JSON property specifies one type, but the provided value can be interpreted as another.

By default, `json.Serializer` performs strict schema validation. If a JSON property is expected to be a certain type (for example, boolean), but the value has a different type (for example, a string), a deserialization error occurs.

When `allowImplicitConversion` is enabled:

- **Primitive types**: Implicit conversions for primitive BDL types (see [Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")) are permitted. When the option is disabled, implicit conversions are not allowed.
- **Arrays**: Arrays of `INTEGER`, `NUMBER(DECIMAL)`, `BOOLEAN`, or `STRING` can accept values that can be converted implicitly. For example, boolean values are accepted as `1`, `0`, `true`, `false`, `"true"`, `"false"`, `"1"`, or `"0"`. Dictionary elements follow the same rules.
- **Records/Objects**: Record or object fields are not affected by this option.

> **Important:**
>
> This option does not apply to simple-type `oneOf` JSON schemas. Allowing implicit conversion can break schema validation by creating ambiguous matches.

## Affects

Deserialization (JSON to BDL)

## Example

To enable implicit type conversions during deserialization:

```
CALL json.Serializer.setOption("allowImplicitConversion", 1)
```

For a detailed example, see [Deserializing JSON values with allowImplicitConversion](3694-allowimplicitconversion-example.md "Use the allowImplicitConversion option to accept JSON values that require implicit type conversion during deserialization.").

## Related links

**Related concepts**  

[The json.Serializer class](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

[Implicit and explicit JSON conversion in json.Serializer and util.JSON](3663-implicit-and-explicit-conversion.md "This topic provides an overview of implicit and explicit JSON conversion behavior in json.Serializer and util.JSON, highlighting the differences between strict and permissive conversion models.")

[Best practices](3669-best-practices.md "Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
