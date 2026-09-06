---
title: "The json.JSONWriter class"
source: "fgl-topics/c_gws_JSONWriter.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class"
type: "concept"
---

# The json.JSONWriter class

> The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source.

The `json.JSONWriter` class is provided in the `json`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`json.JSONWriter` class, you must import the `json` package in your
program:

```
IMPORT json
```

The `json.JSONWriter` class implements a method to create a
`json.JSONWriter` object.

The purpose of the `json.JSONWriter` class is to process JSON text to interface
with software based on the JSON format by streaming it in a sequence of tokens, one token at a time,
instead of loading it into memory. The stream includes both literal values (strings, numbers,
booleans, and null) as well as the begin and end delimiters of JSON's two structured types (objects
and arrays).

The output format for `DATETIME` and `INTERVAL` values is
controlled by the `datetimeSerializationMode` and
`intervalSerializationMode` global options, set using [json.Serializer.setOption()](3675-json-serializer-setoption.md "Sets an option on the JSON serializer."). See [json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")

## Child topics

- [json.JSONWriter methods](3619-json-jsonwriter-methods.md): Methods for the json.JSONWriter class.
- [Examples](3637-examples.md): json.JSONWriter usage examples.
