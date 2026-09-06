---
title: "Best practices"
source: "fgl-topics/c_gws_JSONSerializer_best_practices.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > Implicit and explicit conversion > Best practices"
type: "concept"
---

# Best practices

> Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.

The following best practices can help ensure effective use of
`json.Serializer`:

- Disable `allowImplicitConversion` in production to avoid silent errors.
- Use native JSON booleans and numbers rather than quoted values (for example, use
  `true` instead of `"true"`).
- Favor strict formats for `DATE` (`YYYY-MM-DD`); for
  `DATETIME` and `INTERVAL`, you have options to use RFC 3339 and
  ISO 8601 standards, see [json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").
- Use `json_null` attribute when handling
  `NULL` values. For more information about managing null with the
  `json_null` attribute, go to [NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.").
- Use the [`json_name`](../09_advanced-features/0958-bdl-names-and-json-element-names.md "To identify elements, JSON standards use different format as Genero BDL variable names.") attribute
  to map a BDL variable to a JSON element name that cannot be represented as a valid BDL identifier,
  such as names with spaces or special characters. For details, see [Type attributes](../08_language-basics/0756-type-attributes.md "Types can be defined with meta-data information.").

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[Attributes on record definitions](../08_language-basics/0718-attributes-on-record-definitions.md "Records can be defined with attributes, to complete the type description.")

[Example 5: Type attributes](../08_language-basics/0708-example-5-type-attributes.md "Example 5: Type attributes")
