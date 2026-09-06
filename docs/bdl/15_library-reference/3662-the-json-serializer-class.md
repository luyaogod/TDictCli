---
title: "The json.Serializer class"
source: "fgl-topics/c_gws_JSONSerializer.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class"
type: "concept"
---

# The json.Serializer class

> The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.

The `json.Serializer` class is provided in the `json`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`json.Serializer` class, you must import the `json` package in your
program:

```
IMPORT json
```

The `json.Serializer` class provides methods that comply with schema-validation
standards ([JSON Schema](https://json-schema.org/)
and [OpenAPI](https://swagger.io/specification/))
and that:

- Serialize Genero BDL variables to JSON objects.
- Deserialize JSON objects to Genero BDL variables.

> **Note:**
>
> In contrast, [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") is designed to manipulate raw
> JSON data; it does not handle schema validation.

The `json.Serializer` class is a static class and does not have to be
instantiated. The status variable is set to zero after a successful method call.

## Implicit and explicit conversion of JSON to BDL

- `json.Serializer:` Implements strict, schema-driven conversion (explicit).
  Validates against JSON Schema/OpenAPI and performs controlled serialization/deserialization.
- [`util.JSON`](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") Implements permissive,
  dynamic conversion (implicit). Designed only to manipulate raw JSON data; it does not handle schema
  validation.

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [Implicit and explicit JSON conversion in json.Serializer and util.JSON](3663-implicit-and-explicit-conversion.md): This topic provides an overview of implicit and explicit JSON conversion behavior in json.Serializer and util.JSON, highlighting the differences between strict and permissive conversion models.
- [json.Serializer methods](3671-json-serializer-methods.md): Methods for the json.Serializer class.
