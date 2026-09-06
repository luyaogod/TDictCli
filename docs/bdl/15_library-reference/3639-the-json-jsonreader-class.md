---
title: "The json.JSONReader class"
source: "fgl-topics/c_gws_JSONReader.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class"
type: "concept"
---

# The json.JSONReader class

> The json.JSONReader class provides an interface compatible with JSON streaming that reads data in a JSON format from an input source.

The `json.JSONReader` class is provided in the `json`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`json.JSONReader` class, you must import the `json` package in your
program:

```
IMPORT json
```

The `json.JSONReader` class implements a method to create a
`json.JSONReader` object.

The purpose of the `json.JSONReader` class is to read JSON encoded text as a
stream of tokens. The stream includes both literal values (strings, numbers, booleans, and null) as
well as the begin and end delimiters of JSON's two structured types (objects and arrays). The tokens
are read in the same order that they appear in the JSON document.

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [json.JSONReader methods](3640-json-jsonreader-methods.md): Methods for the json.JSONReader class.
- [JSONReader Event Types](3659-jsonreader-event-types.md): Event types of the json.JSONReader class.
- [Examples](3660-examples.md): json.JSONReader usage examples.
