---
title: "JSON serialization rules and customization"
source: "fgl-topics/c_gws_json_serialization_overview.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization"
type: "concept"
---

# JSON serialization rules and customization

> This section describes the rules, attributes, and options used when converting BDL values to and from JSON in GWS applications.

The topics in this section cover the default BDL/JSON type mappings, the attributes that modify
these mappings, and the options that control JSON serialization and deserialization. These
attributes and options are also referenced from the [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") class page.

For permissive JSON parsing rules implemented by the [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") class, refer to [JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.") and
[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")

## Related links

**Related concepts**  

[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")

## Child topics

- [Implicit and explicit JSON conversion in json.Serializer and util.JSON](5049-implicit-and-explicit-conversion.md): This topic provides an overview of implicit and explicit JSON conversion behavior in json.Serializer and util.JSON, highlighting the differences between strict and permissive conversion models.
- [json.Serializer options](5063-json-serializer-options.md): Options for controlling serialization and deserialization behavior, set using the json.Serializer class.
