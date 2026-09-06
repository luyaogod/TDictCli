---
title: "json.JSONReader.getByteValue"
source: "fgl-topics/c_gws_jsonJSONReader_getByteValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getByteValue"
type: "concept"
---

# json.JSONReader.getByteValue

> Gets a byte value of a JSONReader object.

## Syntax

```
getByteValue(
   by BYTE)
```

1. by is of [BYTE](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") type.

## Usage

Use this method to get a byte value of a [JSONReader](3639-the-json-jsonreader-class.md "The json.JSONReader class provides an interface compatible with JSON streaming that reads data in a JSON format from an input source.")
object, where by can have a value in byte form.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
